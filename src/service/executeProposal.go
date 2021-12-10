package rlr

import (
	"fmt"
	"time"

	"gitlab.nekotal.tech/lachain/crosschain/bridge-backend-service/src/service/storage"
	workers "gitlab.nekotal.tech/lachain/crosschain/bridge-backend-service/src/service/workers"
	"gitlab.nekotal.tech/lachain/crosschain/bridge-backend-service/src/service/workers/utils"
)

// !!! TODO !!!

// emitRegistreted ...
func (r *BridgeSRV) emitProposal(worker workers.IWorker) {
	for {
		events := r.storage.GetEventsByTypeAndStatuses([]storage.EventStatus{storage.EventStatusPassedInit, storage.EventStatusPassedSent})
		for _, event := range events {
			if event.Status == storage.EventStatusPassedInit &&
				worker.GetDestinationID() == event.DestinationChainID {
				r.logger.Infoln("attempting to send execute proposal")
				if _, err := r.sendExecuteProposal(worker, event); err != nil {
					r.logger.Errorf("submit claim failed: %s", err)
				}
			} else {
				r.handleTxSent(event.ChainID, event, storage.TxTypePassed,
					storage.EventStatusPassedConfirmed, storage.EventStatusPassedFailed)
			}
		}

		time.Sleep(2 * time.Second)
	}
}

// ethSendClaim ...
func (r *BridgeSRV) sendExecuteProposal(worker workers.IWorker, event *storage.Event) (txHash string, err error) {
	txSent := &storage.TxSent{
		Chain:      worker.GetChainName(),
		Type:       storage.TxTypePassed,
		CreateTime: time.Now().Unix(),
	}

	r.logger.Infof("Execute parameters:  depositNonce(%d) | sender(%s) | outAmount(%s) | resourceID(%s) | chainID(%s)\n",
		event.DepositNonce, event.SenderAddr, event.OutAmount, event.ResourceID, worker.GetChainName())
	if worker.GetChainName() == storage.LaChain {
		txHash, err = worker.ExecuteProposalLa(event.DepositNonce, utils.StringToBytes8(event.OriginChainID), utils.StringToBytes8(event.DestinationChainID), utils.StringToBytes32(event.ResourceID),
			event.ReceiverAddr, event.OutAmount)
	} else {
		txHash, err = worker.ExecuteProposalEth(event.DepositNonce, utils.StringToBytes8(event.OriginChainID), utils.StringToBytes8(event.DestinationChainID), utils.StringToBytes32(event.ResourceID),
			event.ReceiverAddr, event.OutAmount)
	}
	if err != nil {
		txSent.ErrMsg = err.Error()
		txSent.Status = storage.TxSentStatusFailed
		r.storage.UpdateEventStatus(event, storage.EventStatusPassedSentFailed)
		r.storage.CreateTxSent(txSent)
		return "", fmt.Errorf("could not send claim tx: %w", err)
	}
	txSent.TxHash = txHash
	r.storage.UpdateEventStatus(event, storage.EventStatusPassedSent)
	r.logger.Infof("send execute proposal tx success | chain=%s, tx_hash=%s", worker.GetChainName(), txSent.TxHash)
	// create new tx(claimed)
	r.storage.CreateTxSent(txSent)

	return txSent.TxHash, nil

}
