package rlr

import (
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"gitlab.nekotal.tech/lachain/crosschain/bridge-backend-service/src/service/storage"
	workers "gitlab.nekotal.tech/lachain/crosschain/bridge-backend-service/src/service/workers"
)

// !!! TODO !!!

// emitRegistreted ...
func (r *BridgeSRV) emitPenalty(worker workers.IWorker) {
	for {
		events := r.storage.GetEventsByTypeAndStatuses([]storage.EventStatus{storage.EventStatusPenaltyInit, storage.EventStatusPenaltySent})
		for _, event := range events {
			if event.Status == storage.EventStatusPenaltyInit {
				r.logger.Infoln("attempting to send Penalty")
				if _, err := r.sendPenalty(worker, event); err != nil {
					r.logger.Errorf("submit claim failed: %s", err)
				}
			} else {
				r.handleTxSent(event.ChainID, event, storage.TxTypePenalty,
					storage.EventStatusPenaltyConfirmed, storage.EventStatusPenaltyFailed)
			}
		}

		time.Sleep(2 * time.Second)
	}
}

// ethSendClaim ...
func (r *BridgeSRV) sendPenalty(worker workers.IWorker, event *storage.Event) (string, error) {
	txSent := &storage.TxSent{
		Chain:      worker.GetChain(),
		Type:       storage.TxTypePenalty,
		CreateTime: time.Now().Unix(),
	}

	r.logger.Infof("Penalty parameters: relayer(%s) penalty(%s)\n", event.RelayerAddress, event.Penalty)

	txHash, err := worker.Penalty(common.HexToAddress(event.RelayerAddress), event.Penalty)
	if err != nil {
		txSent.ErrMsg = err.Error()
		txSent.Status = storage.TxSentStatusFailed
		r.storage.UpdateEventStatus(event, storage.EventStatusPenaltySentFailed)
		r.storage.CreateTxSent(txSent)
		return "", fmt.Errorf("could not send claim tx: %w", err)
	}
	txSent.TxHash = txHash
	r.storage.UpdateEventStatus(event, storage.EventStatusPenaltySent)
	r.logger.Infof("send Penalty tx success | chain=%s, tx_hash=%s", worker.GetChain(), txSent.TxHash)
	// create new tx(claimed)
	r.storage.CreateTxSent(txSent)

	return txSent.TxHash, nil

}
