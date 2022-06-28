package rlr

import (
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/latoken/bridge-backend-service/src/service/storage"
	"github.com/latoken/bridge-backend-service/src/service/workers"
	"github.com/latoken/bridge-backend-service/src/service/workers/utils"
)

// Updates withdraw swap status on lachain
func (b *BridgeSRV) UpdateTxOnLachain() {
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
	txSent := &storage.TxSent{
		Chain:      "LA",
		Type:       storage.TxTypeSpend,
		SwapID:     event.SwapID,
		CreateTime: time.Now().Unix(),
	}
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

	var originWorker workers.IWorker
	var destWorker workers.IWorker
	for _, wrkr := range b.Workers {
		if strings.ToLower(wrkr.GetDestinationID()) == strings.ToLower(event.OriginChainID) {
			originWorker = wrkr
		}
		if strings.ToLower(wrkr.GetDestinationID()) == strings.ToLower(event.DestinationChainID) {
			destWorker = wrkr
		}
	}

	if originWorker == nil || destWorker == nil {
		err := "Missing worker"
		println(err)
		txSent.ErrMsg = err
		txSent.Status = storage.TxSentStatusFailed
		b.storage.UpdateEventStatus(event, storage.EventStatusUpdateFailed)
		b.storage.CreateTxSent(txSent)
		return "", fmt.Errorf("could not send update tx: %s", err)
	}

	originDecimals, err := originWorker.GetDecimalsFromResourceID(event.ResourceID)
	if err != nil {
		println("error in decimals", err.Error())
		txSent.ErrMsg = err.Error()
		txSent.Status = storage.TxSentStatusFailed
		b.storage.UpdateEventStatus(event, storage.EventStatusUpdateFailed)
		b.storage.CreateTxSent(txSent)
		return "", fmt.Errorf("could not send update tx: %w", err)
	}

	destDecimals, err := destWorker.GetDecimalsFromResourceID(event.ResourceID)
	if err != nil {
		println("error in decimals", err.Error())
		txSent.ErrMsg = err.Error()
		txSent.Status = storage.TxSentStatusFailed
		b.storage.UpdateEventStatus(event, storage.EventStatusUpdateFailed)
		b.storage.CreateTxSent(txSent)
		return "", fmt.Errorf("could not send update tx: %w", err)
	}

	//using inAmount to check for decimals of other chain
	var inAmount *big.Int
	outAmount, _ := new(big.Int).SetString(event.OutAmount, 10)

	if originDecimals == destDecimals {
		inAmount = outAmount
	} else if originDecimals == 0 || destDecimals == 0 || originDecimals > 63 || destDecimals > 63 {
		err = fmt.Errorf("One of decimals is zero or greater than 63")
		println("error in decimals", err.Error())
		txSent.ErrMsg = err.Error()
		txSent.Status = storage.TxSentStatusFailed
		b.storage.UpdateEventStatus(event, storage.EventStatusUpdateFailed)
		b.storage.CreateTxSent(txSent)
		return "", fmt.Errorf("could not send update tx: %w", err)
	} else {
		inAmount = utils.ConvertDecimalsForInput(originDecimals, destDecimals, event.OutAmount)
	}

	b.logger.Infof("Update status parameters:  depositNonce(%d) | sender(%s) | outAmount(%s) | resourceID(%s) | inAmount(%s) \n",
		event.DepositNonce, event.ReceiverAddr, event.OutAmount, event.ResourceID, inAmount.String())

	//required to update liquidity index for amTokens
	if event.ResourceID == b.storage.FetchResourceIDByName("amToken").ID {
		wor := b.Workers["POS"]
		liquidity, _ = wor.GetLiquidityIndex(wor.GetConfig().AmTokenHandlerAddress, wor.GetConfig().AMUSDTContractAddr)
	}

	txHash, err := b.laWorker.UpdateSwapStatusOnChain(event.DepositNonce, utils.StringToBytes8(event.OriginChainID), utils.StringToBytes8(event.DestinationChainID), utils.StringToBytes32(event.ResourceID), event.ReceiverAddr, outAmount, inAmount, liquidity, status)
	if err != nil {
		txSent.ErrMsg = err.Error()
		txSent.Status = storage.TxSentStatusFailed
		b.storage.CreateTxSent(txSent)
		b.storage.UpdateEventStatus(event, storage.EventStatusUpdateFailed)
		return "", err
	}
	b.logger.Infof("Update status tx success: %s", txHash)
	b.storage.UpdateEventStatus(event, storage.EventStatusUpdateConfirmed)
	b.storage.CreateTxSent(txSent)
	return txHash, nil
}
