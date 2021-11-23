package rlr

import (
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"gitlab.nekotal.tech/lachain/crosschain/bridge-backend-service/src/service/storage"
	workers "gitlab.nekotal.tech/lachain/crosschain/bridge-backend-service/src/service/workers"
)

// emitRegistreted ...
func (r *BridgeSRV) emitFelony(worker workers.IWorker) {
	for {
		events := r.storage.GetEventsByTypeAndStatuses([]storage.EventStatus{storage.EventStatusFelonyInit, storage.EventStatusFelonySent})
		for _, event := range events {
			println(worker.GetChain(), event.ChainID)
			if event.Status == storage.EventStatusFelonyInit {
				r.logger.Infoln("attempting to send felony")
				if _, err := r.sendFelony(worker, event); err != nil {
					r.logger.Errorf("submit claim failed: %s", err)
				}
			} else {
				r.handleTxSent(worker.GetChain(), event, storage.TxTypeFelony,
					storage.EventStatusFelonyConfirmed, storage.EventStatusFelonyFailed)
			}
		}

		time.Sleep(2 * time.Second)
	}
}

// ethSendClaim ...
func (r *BridgeSRV) sendFelony(worker workers.IWorker, event *storage.Event) (string, error) {
	txSent := &storage.TxSent{
		Chain:      worker.GetChain(),
		Type:       storage.TxTypeFelony,
		CreateTime: time.Now().Unix(),
	}
	r.logger.Infof("felony parameters: relayer(%s)\n", event.RelayerAddress)
	txHash, err := worker.Felony(common.HexToAddress(event.RelayerAddress), event.ChainID)
	if err != nil {
		txSent.ErrMsg = err.Error()
		txSent.Status = storage.TxSentStatusFailed
		r.storage.UpdateEventStatus(event, storage.EventStatusFelonySentFailed)
		r.storage.CreateTxSent(txSent)
		return "", fmt.Errorf("could not send claim tx: %w", err)
	}
	txSent.TxHash = txHash
	r.storage.UpdateEventStatus(event, storage.EventStatusFelonySent)
	r.logger.Infof("send register tx success | chain=%s, tx_hash=%s", worker.GetChain(), txSent.TxHash)
	// create new tx(claimed)
	r.storage.CreateTxSent(txSent)

	return txSent.TxHash, nil
}
