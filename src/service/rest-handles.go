package rlr

import (
	"gitlab.nekotal.tech/lachain/crosschain/bridge-backend-service/src/models"
)

// Status ...
func (r *BridgeSRV) Status() (*models.RelayerStatus, error) {
	relayerStatus := &models.RelayerStatus{
		Mode:    "",
		Workers: make(map[string]models.WorkerStatus),
	}

	// for k, w := range r.Workers {
	// 	fmt.Println("----- NAME: ", k)
	// 	var worker models.WorkerStatus
	// 	// get block logs
	// 	currentBlockLog, err := r.storage.GetCurrentBlockLog(w.GetChain())
	// 	if err != nil {
	// 		return nil, err
	// 	}

	// 	worker.SyncHeight = currentBlockLog.Height
	// 	worker.LastBlockFetchedAt = time.Unix(currentBlockLog.CreateTime, 0)

	// 	height, err := w.GetHeight()
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	worker.Height = height

	// 	status, err := w.GetStatus("LA") //TODO
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	worker.Status = status

	// 	relayerStatus.Workers[k] = worker
	// }

	// fmt.Println(relayerStatus)

	return relayerStatus, nil
}

// CreateNewBindRequest ...
func (r *BridgeSRV) CreateNewBindRequest() {}

// GetRelayerAccountBalance ...
func (r *BridgeSRV) GetRelayerAccountBalance() {

}

// GetWorkers ...
func (r *BridgeSRV) GetWorkers() {
	//	r.Workers
}
