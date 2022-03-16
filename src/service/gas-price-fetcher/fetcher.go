package fetcher

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
	"gitlab.nekotal.tech/lachain/crosschain/bridge-backend-service/src/models"
	"gitlab.nekotal.tech/lachain/crosschain/bridge-backend-service/src/service/storage"
)

//FetcherSrv
type FetcherSrv struct {
	logger     *logrus.Entry
	storage    *storage.DataBase
	posFetCfg  *models.FetcherConfig
	bscFetCfg  *models.FetcherConfig
	ethFetCfg  *models.FetcherConfig
	avaxFetCfg *models.FetcherConfig
}

//CreateFetcherSrv
func CreateFetcherSrv(logger *logrus.Logger, db *storage.DataBase, posFetCfg, bscFetCfg, ethFetCfg, avaxFetCfg *models.FetcherConfig) *FetcherSrv {
	return &FetcherSrv{
		logger:     logger.WithField("layer", "fetcher"),
		storage:    db,
		posFetCfg:  posFetCfg,
		bscFetCfg:  bscFetCfg,
		ethFetCfg:  ethFetCfg,
		avaxFetCfg: avaxFetCfg,
	}
}

func (f *FetcherSrv) Run() {
	f.logger.Infoln("Fetcher srv started")
	go f.collector()
}

func (f *FetcherSrv) collector() {
	for {
		f.getAllGasPrice()
		time.Sleep(30 * time.Second)
	}
}

func (f *FetcherSrv) getAllGasPrice() {
	bscGasPrice := f.getBscGasPrice()
	ethGasPrice := f.getEthGasPrice()
	posGasPrice := f.getPosGasPrice()
	avaxGasPrice := f.getAvaxGasPrice()
	gasPrices := make([]*storage.GasPrice, 3)
	gasPrices[0] = bscGasPrice
	gasPrices[1] = ethGasPrice
	gasPrices[2] = posGasPrice
	gasPrices[3] = avaxGasPrice
	f.storage.SaveGasPriceInfo(gasPrices)
	f.logger.Infoln("New gas prices fetched")
}

func (f *FetcherSrv) getBscGasPrice() *storage.GasPrice {
	// httpClient := &http.Client{
	// 	Timeout: time.Second * 10,
	// }

	// resp, err := f.makeReq(f.bscFetCfg.URL, httpClient)
	// if err != nil {
	// 	logrus.Warnf("fetch BSC gas price error = %s", err)
	// 	return &storage.GasPrice{}
	// }
	// var gasPrice = (*resp)["standard"].(float64)
	return &storage.GasPrice{ChainName: "BSC", Price: fmt.Sprintf("%f", 20.0), UpdateTime: time.Now().Unix()}
}

func (f *FetcherSrv) getEthGasPrice() *storage.GasPrice {
	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}

	resp, err := f.makeReq(f.ethFetCfg.URL, httpClient)
	if err != nil {
		logrus.Warnf("fetch ETH gas price error = %s", err)
		return &storage.GasPrice{}
	}
	var gasPrice = (*resp)["average"].(float64) / 10
	return &storage.GasPrice{ChainName: "ETH", Price: fmt.Sprintf("%f", gasPrice), UpdateTime: time.Now().Unix()}
}

func (f *FetcherSrv) getPosGasPrice() *storage.GasPrice {
	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}

	resp, err := f.makeReq(f.posFetCfg.URL, httpClient)
	if err != nil {
		logrus.Warnf("fetch ETH gas price error = %s", err)
		return &storage.GasPrice{}
	}
	var gasPrice = (*resp)["standard"].(float64)
	return &storage.GasPrice{ChainName: "POS", Price: fmt.Sprintf("%f", gasPrice), UpdateTime: time.Now().Unix()}
}

func (f *FetcherSrv) getAvaxGasPrice() *storage.GasPrice {
	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}

	resp, err := f.makeReq(f.avaxFetCfg.URL, httpClient)
	if err != nil {
		logrus.Warnf("fetch ETH gas price error = %s", err)
		return &storage.GasPrice{}
	}
	var gasPrice = (*resp)["standard"].(float64)
	return &storage.GasPrice{ChainName: "AVAX", Price: fmt.Sprintf("%f", gasPrice), UpdateTime: time.Now().Unix()}
}

// MakeReq HTTP request helper
func (f *FetcherSrv) makeReq(url string, c *http.Client) (*map[string]interface{}, error) {
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}
	resp, err := f.doReq(req, c)
	if err != nil {
		return nil, err
	}

	t := make(map[string]interface{})
	er := json.Unmarshal(resp, &t)
	if er != nil {
		return nil, er
	}

	return &t, err
}

// helper
// doReq HTTP client
func (f *FetcherSrv) doReq(req *http.Request, client *http.Client) ([]byte, error) {
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if 200 != resp.StatusCode {
		return nil, fmt.Errorf("%s", body)
	}
	return body, nil
}
