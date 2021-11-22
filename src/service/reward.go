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
func (r *BridgeSRV) emitReward(worker workers.IWorker) {
	for {
		events := r.storage.GetEventsByTypeAndStatuses([]storage.EventStatus{storage.EventStatusRewardInit, storage.EventStatusRewardSent})
		for _, event := range events {
			if event.Status == storage.EventStatusRewardInit {
				r.logger.Infoln("attempting to send Reward")
				println(event.ChainID, event.Status, event.EventID)
				if _, err := r.sendReward(worker, event); err != nil {
					r.logger.Errorf("submit claim failed: %s", err)
				}
			} else {
				r.handleTxSent(event.ChainID, event, storage.TxTypeReward,
					storage.EventStatusRewardConfirmed, storage.EventStatusRewardFailed)
			}
		}

		time.Sleep(2 * time.Second)
	}
}

// ethSendClaim ...
func (r *BridgeSRV) sendReward(worker workers.IWorker, event *storage.Event) (string, error) {
	txSent := &storage.TxSent{
		Chain:      worker.GetChain(),
		Type:       storage.TxTypeReward,
		CreateTime: time.Now().Unix(),
	}

	r.logger.Infof("Reward parameters: relayer(%s) Reward(%s)\n", event.RelayerAddress, event.Penalty)

	txHash, err := worker.Reward(common.HexToAddress(event.RelayerAddress), event.Penalty)
	if err != nil {
		txSent.ErrMsg = err.Error()
		txSent.Status = storage.TxSentStatusFailed
		r.storage.UpdateEventStatus(event, storage.EventStatusRewardSentFailed)
		r.storage.CreateTxSent(txSent)
		return "", fmt.Errorf("could not send claim tx: %w", err)
	}
	txSent.TxHash = txHash
	r.storage.UpdateEventStatus(event, storage.EventStatusRewardSent)
	r.logger.Infof("send Reward tx success | chain=%s, tx_hash=%s", worker.GetChain(), txSent.TxHash)
	// create new tx(claimed)
	r.storage.CreateTxSent(txSent)

	return txSent.TxHash, nil

}
