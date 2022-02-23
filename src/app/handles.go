package app

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
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
			"/gas-price/{chain}",
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

func (a *App) GasPriceHandler(w http.ResponseWriter, r *http.Request) {
	msg := mux.Vars(r)["chain"]

	if msg == "" {
		a.logger.Errorf("Empty request(gas-price/{chain})")
		common.ResponJSON(w, http.StatusInternalServerError, createNewError("empty request", ""))
		return
	}
	gasPrice := a.relayer.GetGasPrice(msg)
	common.ResponJSON(w, http.StatusOK, gasPrice)
}
