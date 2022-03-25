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
	logger       *logrus.Entry
	storage      *storage.DataBase
	chainFetCfgs []*models.FetcherConfig
}

//CreateFetcherSrv
func CreateFetcherSrv(logger *logrus.Logger, db *storage.DataBase, chainFetCfgs []*models.FetcherConfig) *FetcherSrv {
	return &FetcherSrv{
		logger:       logger.WithField("layer", "fetcher"),
		storage:      db,
		chainFetCfgs: chainFetCfgs,
	}
}

func (f *FetcherSrv) Run() {
	f.logger.Infoln("Fetcher srv started")
	go f.collector()
}

func (f *FetcherSrv) collector() {
	for {
		f.getAllGasPrice()
		time.Sleep(60 * time.Second)
	}
}

func (f *FetcherSrv) getAllGasPrice() {
	gasPrices := make([]*storage.GasPrice, 0, len(f.chainFetCfgs))

	for _, cfg := range f.chainFetCfgs {
		switch cfg.ChainName {
		case "BSC":
			gasPrices = append(gasPrices, f.getBscGasPrice(cfg))
		case "ETH":
			gasPrices = append(gasPrices, f.getEthGasPrice(cfg))
		case "POS":
			gasPrices = append(gasPrices, f.getPosGasPrice(cfg))
		case "AVAX":
			gasPrices = append(gasPrices, f.getAvaxGasPrice(cfg))
		case "FTM":
			gasPrices = append(gasPrices, f.getFtmGasPrice(cfg))
		default:
			logrus.Warnf("Gas price getter not implemented for ", cfg.ChainName)
		}
	}

	f.storage.SaveGasPriceInfo(gasPrices)
	f.logger.Infoln("New gas prices fetched")
}

func (f *FetcherSrv) getBscGasPrice(cfg *models.FetcherConfig) *storage.GasPrice {
	// httpClient := &http.Client{
	// 	Timeout: time.Second * 10,
	// }

	// resp, err := f.makeReq(f.chainFetCfgs[storage.BscChain].URL, httpClient)
	// if err != nil {
	// 	logrus.Warnf("fetch BSC gas price error = %s", err)
	// 	return &storage.GasPrice{}
	// }
	// var gasPrice = (*resp)["standard"].(float64)
	return &storage.GasPrice{ChainName: "BSC", Price: fmt.Sprintf("%f", 20.0), UpdateTime: time.Now().Unix()}
}

func (f *FetcherSrv) getEthGasPrice(cfg *models.FetcherConfig) *storage.GasPrice {
	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}

	resp, err := f.makeReq(cfg.URL, httpClient)
	if err != nil {
		logrus.Warnf("fetch ETH gas price error = %s", err)
		return &storage.GasPrice{}
	}
	var gasPrice = (*resp)["average"].(float64) / 10
	return &storage.GasPrice{ChainName: "ETH", Price: fmt.Sprintf("%f", gasPrice), UpdateTime: time.Now().Unix()}
}

func (f *FetcherSrv) getPosGasPrice(cfg *models.FetcherConfig) *storage.GasPrice {
	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}

	resp, err := f.makeReq(cfg.URL, httpClient)
	if err != nil {
		logrus.Warnf("fetch POS gas price error = %s", err)
		return &storage.GasPrice{}
	}
	var gasPrice = (*resp)["fast"].(float64)
	return &storage.GasPrice{ChainName: "POS", Price: fmt.Sprintf("%f", gasPrice), UpdateTime: time.Now().Unix()}
}

func (f *FetcherSrv) getAvaxGasPrice(cfg *models.FetcherConfig) *storage.GasPrice {
	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}

	resp, err := f.makeReq(cfg.URL, httpClient)
	if err != nil {
		logrus.Warnf("fetch AVAX gas price error = %s", err)
		return &storage.GasPrice{}
	}
	var gasPrice = (*resp)["standard"].(float64)
	return &storage.GasPrice{ChainName: "AVAX", Price: fmt.Sprintf("%f", gasPrice), UpdateTime: time.Now().Unix()}
}

func (f *FetcherSrv) getFtmGasPrice(cfg *models.FetcherConfig) *storage.GasPrice {
	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}

	resp, err := f.makeReq(cfg.URL, httpClient)
	if err != nil {
		logrus.Warnf("fetch FTM gas price error = %s", err)
		return &storage.GasPrice{}
	}
	var gasPrice = (*resp)["standard"].(float64)
	return &storage.GasPrice{ChainName: "FTM", Price: fmt.Sprintf("%f", gasPrice), UpdateTime: time.Now().Unix()}
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
