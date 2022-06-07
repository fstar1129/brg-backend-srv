package rlr

import (
	"time"

	"github.com/latoken/bridge-backend-service/src/service/storage"
	"github.com/latoken/bridge-backend-service/src/service/workers/utils"
)

func (b *BridgeSRV) UpdateFailedTxOnLachain() {
	for {
		events := b.storage.GetEventsByTypeAndStatuses([]storage.EventStatus{storage.EventStatusPassedFailed, storage.EventStatusPassedSent})
		for _, event := range events {
			b.logger.Infoln("attempting to send confirmation tx")
			txHash, err := b.SendConfirmationLA(event)
			if err != nil {
				b.logger.Errorf("confirmation failed: %s | txHash: %s", err, txHash)
			}
			b.logger.Infoln("confirmation tx success")
		}
		time.Sleep(time.Minute)
	}
}

func (b *BridgeSRV) SendConfirmationLA(event *storage.Event) (string, error) {
	var liquidity []byte
	var status uint8 = 3
	if event.Status == storage.EventStatusPassedFailed {
		status = 4
	}

	//no need to update on chain for deposit to lachain tx
	if event.DestinationChainID == b.laWorker.GetDestinationID() {
		if status == 3 {
			b.storage.UpdateEventStatus(event, storage.EventStatusUpdateConfirmed)
		} else {
			b.storage.UpdateEventStatus(event, storage.EventStatusUpdateFailed)
		}
		return "", nil
	}

	b.logger.Infof("Update status parameters:  depositNonce(%d) | sender(%s) | outAmount(%s) | resourceID(%s) \n",
		event.DepositNonce, event.ReceiverAddr, event.OutAmount, event.ResourceID)

	//required to update liquidity index for amTokens
	if event.ResourceID == b.storage.FetchResourceIDByName("amToken").ID {
		wor := b.Workers["POS"]
		liquidity, _ = wor.GetLiquidityIndex(wor.GetConfig().AmTokenHandlerAddress, wor.GetConfig().AMUSDTContractAddr)
	}

	txHash, err := b.laWorker.UpdateSwapStatusOnChain(event.DepositNonce, utils.StringToBytes8(event.OriginChainID), utils.StringToBytes8(event.DestinationChainID), utils.StringToBytes32(event.ResourceID), event.ReceiverAddr, event.OutAmount, liquidity, status)
	if err != nil {
		b.storage.UpdateEventStatus(event, storage.EventStatusUpdateFailed)
		return "", err
	}
	b.logger.Infof("Update status tx success: %s", txHash)
	b.storage.UpdateEventStatus(event, storage.EventStatusUpdateConfirmed)
	return txHash, nil
}
