package app

import (
	"encoding/json"
	"net/http"

	"gitlab.nekotal.tech/lachain/crosschain/bridge-backend-service/src/common"
)

const numPerPage = 100

// Endpoints ...
func (a *App) Endpoints(w http.ResponseWriter, r *http.Request) {
	endpoints := struct {
		Endpoints []string `json:"endpoints"`
	}{
		Endpoints: []string{
			"/status",
			"/failed_swaps/{page}",
			"/resend_tx/{id}",
			"/set_mode/{mode}",
		},
	}

	jsonBytes, err := json.MarshalIndent(endpoints, "", "    ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

// StatusHandler ...
func (a *App) StatusHandler(w http.ResponseWriter, r *http.Request) {
	status, err := a.relayer.Status()
	if err != nil {
		common.ResponError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// jsonBytes, err := json.MarshalIndent(status, "", "    ")
	// if err != nil {
	// 	common.ResponError(w, http.StatusInternalServerError, err.Error())
	// 	return
	// }

	common.ResponJSON(w, http.StatusOK, status)
}

// // FailedSwapsHandler ...
// func (a *App) FailedSwapsHandler(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)
// 	pageStr := params["page"]
// 	if pageStr == "" {
// 		http.Error(w, "required parameter 'page' is missing", http.StatusBadRequest)
// 		return
// 	}

// 	page, err := strconv.Atoi(pageStr)
// 	if err != nil {
// 		http.Error(w, fmt.Sprintf("parse page error, err=%s", err.Error()), http.StatusBadRequest)
// 		return
// 	}

// 	if page < 1 {
// 		http.Error(w, "page should be no less than 1", http.StatusBadRequest)
// 		return
// 	}

// 	failedSwaps, totalCount, err := a.Deputy.FailedSwaps((page-1)*numPerPage, numPerPage)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	res := &a.FailedSwaps{
// 		TotalCount: totalCount,
// 		CurPage:    page,
// 		NumPerPage: numPerPage,
// 		Swaps:      failedSwaps,
// 	}

// 	jsonBytes, err := json.MarshalIndent(res, "", "    ")
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(jsonBytes)
// }

// // ResendTxHandler ...
// func (a *App) ResendTxHandler(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)
// 	idStr := params["id"]
// 	if idStr == "" {
// 		http.Error(w, "required parameter 'id' is missing", http.StatusBadRequest)
// 		return
// 	}

// 	id, err := strconv.Atoi(idStr)
// 	if err != nil {
// 		http.Error(w, fmt.Sprintf("parse id error, err=%s", err.Error()), http.StatusBadRequest)
// 		return
// 	}

// 	txHash, err := a.Deputy.ResendTx(int64(id))

// 	response := struct {
// 		TxHash string `json:"tx_hash"`
// 		ErrMsg string `json:"err_msg"`
// 	}{
// 		TxHash: txHash,
// 	}
// 	if err != nil {
// 		response.ErrMsg = err.Error()
// 	}

// 	jsonBytes, err := json.MarshalIndent(response, "", "    ")
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(jsonBytes)
// }

// // SetModeHandler ...
// func (a *App) SetModeHandler(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)
// 	modeStr := params["mode"]
// 	if modeStr == "" {
// 		http.Error(w, "required parameter 'mode' is missing", http.StatusBadRequest)
// 		return
// 	}

// 	mode, err := strconv.Atoi(modeStr)
// 	if err != nil {
// 		http.Error(w, fmt.Sprintf("parse mode error, err=%s", err.Error()), http.StatusBadRequest)
// 		return
// 	}

// 	if dc.DeputyMode(mode) != dc.DeputyModeNormal && dc.DeputyMode(mode) != dc.DeputyModeStopSendHTLT {
// 		http.Error(w, fmt.Sprintf("mode only supports %d(%s) and %d(%s)",
// 			dc.DeputyModeNormal, dc.DeputyModeNormal, dc.DeputyModeStopSendHTLT, dc.DeputyModeStopSendHTLT), http.StatusBadRequest)
// 		return
// 	}

// 	a.Deputy.SetMode(dc.DeputyMode(mode))

// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// }
