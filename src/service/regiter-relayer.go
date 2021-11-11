package rlr

import (
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"gitlab.nekotal.tech/lachain/crosschain/bridge-backend-service/src/service/storage"
	workers "gitlab.nekotal.tech/lachain/crosschain/bridge-backend-service/src/service/workers"
)

// emitRegistreted ...
func (r *BridgeSRV) emitRegistreted(worker workers.IWorker) {
	for {
		events := r.storage.GetEventsByTypeAndStatuses([]storage.EventStatus{storage.EventStatusRegisterInitConfrimed, storage.EventStatusRegisterSent})
		for _, event := range events {
			if event.Status == storage.EventStatusRegisterInitConfrimed {
				r.logger.Infoln("attempting to send registrate")
				if _, err := r.sendRegister(worker, event); err != nil {
					r.logger.Errorf("submit claim failed: %s", err)
				}
			} else {
				r.handleTxSent(worker.GetChain(), event, storage.TxTypeRegister,
					storage.EventStatusRegisterConfirmed, storage.EventStatusRegisterFailed)
			}
		}

		time.Sleep(2 * time.Second)
	}
}

// ethSendClaim ...
func (r *BridgeSRV) sendRegister(worker workers.IWorker, event *storage.Event) (string, error) {
	txSent := &storage.TxSent{
		Chain:      worker.GetChain(),
		Type:       storage.TxTypeRegister,
		CreateTime: time.Now().Unix(),
	}

	r.logger.Infof("register parameters: relayer(%s)\n", event.RelayerAddress)

	txHash, err := worker.Register(common.HexToAddress(event.RelayerAddress))
	if err != nil {
		txSent.ErrMsg = err.Error()
		txSent.Status = storage.TxSentStatusFailed
		r.storage.UpdateEventStatus(event, storage.EventStatusRegisterSentFailed)
		r.storage.CreateTxSent(txSent)
		return "", fmt.Errorf("could not send claim tx: %w", err)
	}
	txSent.TxHash = txHash
	r.storage.UpdateEventStatus(event, storage.EventStatusRegisterSent)
	r.logger.Infof("send register tx success | chain=%s, tx_hash=%s", worker.GetChain(), txSent.TxHash)
	// create new tx(claimed)
	r.storage.CreateTxSent(txSent)

	return txSent.TxHash, nil
}
