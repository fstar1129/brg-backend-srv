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
func (r *BridgeSRV) emitUnregistered(worker workers.IWorker) {
	for {
		events := r.storage.GetEventsByTypeAndStatuses([]storage.EventStatus{storage.EventStatusUnregisterInit, storage.EventStatusUnregisterSent})
		for _, event := range events {
			if event.Status == storage.EventStatusUnregisterInit {
				r.logger.Infoln("attempting to send unregister")
				if _, err := r.sendUnregister(worker, event); err != nil {
					r.logger.Errorf("submit claim failed: %s", err)
				}
			} else {
				r.handleTxSent(r.laWorker.GetChain(), event, storage.TxTypeUnregister,
					storage.EventStatusUnregisterConfirmed, storage.EventStatusUnregisterFailed)
			}
		}

		time.Sleep(2 * time.Second)
	}
}

// ethSendClaim ...
func (r *BridgeSRV) sendUnregister(worker workers.IWorker, event *storage.Event) (string, error) {
	txSent := &storage.TxSent{
		Chain:      worker.GetChain(),
		Type:       storage.TxTypeUnregister,
		CreateTime: time.Now().Unix(),
	}

	r.logger.Infof("unregister parameters: relayer(%s)\n", event.RelayerAddress)

	txHash, err := worker.Unregister(common.HexToAddress(event.RelayerAddress))
	if err != nil {
		txSent.ErrMsg = err.Error()
		txSent.Status = storage.TxSentStatusFailed
		r.storage.UpdateEventStatus(event, storage.EventStatusUnregisterSentFailed)
		r.storage.CreateTxSent(txSent)
		return "", fmt.Errorf("could not send claim tx: %w", err)
	}
	txSent.TxHash = txHash

	r.logger.Infof("send unregister tx success | chain=%s, tx_hash=%s", worker.GetChain(), txSent.TxHash)
	// create new tx(claimed)
	r.storage.CreateTxSent(txSent)

	return txSent.TxHash, nil

}
