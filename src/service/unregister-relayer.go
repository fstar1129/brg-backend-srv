package rlr

import (
	"gitlab.nekotal.tech/lachain/crosschain/bridge-backend-service/src/service/storage"
	workers "gitlab.nekotal.tech/lachain/crosschain/bridge-backend-service/src/service/workers"
)

// !!! TODO !!!

// emitRegistreted ...
func (r *BridgeSRV) emitUnregistering() {
	for {
		// regs := r.storage.GetRegistrationByStatus([]storage.EventStatus{storage.EventStatusRegisterReceived, storage.EventStatusRegisterSent})
		// for _, reg := range regs {
		// 	if reg.Status == storage.EventStatusRegisterReceived {
		// 		r.logger.Infoln("attempting to send registrate")
		// 		if _, err := r.sendRegister(r.laWorker, reg); err != nil {
		// 			r.logger.Errorf("submit claim failed: %s", err)
		// 		}
		// 	} else {
		// 		r.handleTxSent(r.laWorker.GetChain(), storage.TxTypeRegister,
		// 			storage.EventStatusRegisterConfirmed, storage.EventStatusRegisterFailed)
		// 	}
		// }

		// time.Sleep(2 * time.Second)
	}
}

// ethSendClaim ...
func (r *BridgeSRV) sendUnregister(worker workers.IWorker, reg *storage.Registration) (string, error) {
	// txSent := &storage.TxSent{
	// 	Chain:      worker.GetChain(),
	// 	Type:       storage.TxTypeRegister,
	// 	CreateTime: time.Now().Unix(),
	// }

	// r.logger.Infoln("unegister parameters: relayer(%d)\n", reg.RelayerAddress)

	// txHash, err := worker.Unregister(common.HexToAddress(reg.RelayerAddress))
	// if err != nil {
	// 	txSent.ErrMsg = err.Error()
	// 	txSent.Status = storage.TxSentStatusFailed
	// 	r.storage.CreateTxSent(txSent)
	// 	return "", fmt.Errorf("could not send claim tx: %w", err)
	// }
	// txSent.TxHash = txHash

	// r.logger.Infof("send unregister tx success | chain=%s, tx_hash=%s", worker.GetChain(), txSent.TxHash)
	// // create new tx(claimed)
	// r.storage.CreateTxSent(txSent)

	// return txSent.TxHash, nil

	return "", nil
}
