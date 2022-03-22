package config

import (
	"fmt"
	"math/big"
	"strings"

	"gitlab.nekotal.tech/lachain/crosschain/bridge-backend-service/src/models"
	"gitlab.nekotal.tech/lachain/crosschain/bridge-backend-service/src/service/storage"

	"github.com/ethereum/go-ethereum/common"
)

// Reads Service params from config.json
func (v *viperConfig) ReadServiceConfig() string {
	return fmt.Sprintf("%s:%s", v.GetString("service.host"), v.GetString("service.port"))
}

//reads ethereum chain apis to fetch gas price
func (v *viperConfig) ReadFetcherConfig() (pos, bsc, eth, avax *models.FetcherConfig) {
	return v.readFetcherConfig(storage.PosChain), v.readFetcherConfig(storage.BscChain), v.readFetcherConfig(storage.EthChain), v.readFetcherConfig(storage.AvaxChain)
}

func (v *viperConfig) readFetcherConfig(name string) *models.FetcherConfig {
	return &models.FetcherConfig{
		ChainName: name,
		URL:       v.GetString(fmt.Sprintf("gas_price_api.%s", name)),
	}
}

// ReadLachainConfig reads lachain chain params from config.json
func (v *viperConfig) ReadLachainConfig() *models.WorkerConfig {
	return v.readWorkerConfig(storage.LaChain)
}

// ReadEthWorkerConfig reads ethereum chain worker params from config.json
func (v *viperConfig) ReadWorkersConfig() (pos *models.WorkerConfig, bsc *models.WorkerConfig, eth *models.WorkerConfig, avax *models.WorkerConfig) {
	return v.readWorkerConfig(storage.PosChain), v.readWorkerConfig(storage.BscChain), v.readWorkerConfig(storage.EthChain), v.readWorkerConfig(storage.AvaxChain)
}

// readETHWorkerConfig reads ethereum chain worker params from config.json
func (v *viperConfig) readWorkerConfig(name string) *models.WorkerConfig {
	return &models.WorkerConfig{
		NetworkType:        v.GetString(fmt.Sprintf("workers.%s.type", name)),
		ChainID:            v.GetInt64(fmt.Sprintf("workers.%s.chain_id", name)),
		ChainName:          strings.ToUpper(name),
		User:               v.GetString(fmt.Sprintf("workers.%s.user", name)),
		Password:           v.GetString(fmt.Sprintf("workers.%s.password", name)),
		WorkerAddr:         common.HexToAddress(v.GetString(fmt.Sprintf("workers.%s.worker_addr", name))),
		PrivateKey:         v.GetString(fmt.Sprintf("workers.%s.private_key", name)),
		Provider:           v.GetString(fmt.Sprintf("workers.%s.provider", name)),
		ContractAddr:       common.HexToAddress(v.GetString(fmt.Sprintf("workers.%s.contract_addr", name))),
		TokenContractAddr:  common.HexToAddress(v.GetString(fmt.Sprintf("workers.%s.token_addr", name))),
		GasLimit:           v.GetInt64(fmt.Sprintf("workers.%s.gas_limit", name)),
		GasPrice:           big.NewInt(v.GetInt64(fmt.Sprintf("workers.%s.gas_price", name))),
		FetchInterval:      v.GetInt64(fmt.Sprintf("workers.%s.fetch_interval", name)),
		ConfirmNum:         v.GetInt64(fmt.Sprintf("workers.%s.confirm_num", name)),
		StartBlockHeight:   v.GetInt64(fmt.Sprintf("workers.%s.start_block_height", name)),
		DestinationChainID: v.GetString(fmt.Sprintf("workers.%s.dest_id", name)),
	}
}

// Reads storage params from config.json
func (v *viperConfig) ReadDBConfig() *models.StorageConfig {
	return &models.StorageConfig{
		URL:        v.GetString("storage.url"),
		DBDriver:   v.GetString("storage.driver"),
		DBHOST:     v.GetString("storage.host"),
		DBPORT:     v.GetInt64("storage.port"),
		DBSSL:      v.GetString("storage.ssl_mode"),
		DBName:     v.GetString("storage.db_name"),
		DBUser:     v.GetString("storage.user"),
		DBPassword: v.GetString("storage.password"),
	}
}

func (v *viperConfig) ReadResourceIDs() []*storage.ResourceId {
	tokens := [6]string{"tether", "matic-network", "latoken", "binancecoin", "ethereum", "avalanche"}
	resouceIDs := make([]*storage.ResourceId, len(tokens))
	for index, name := range tokens {
		resouceIDs[index] = &storage.ResourceId{
			Name: name,
			ID:   v.GetString(fmt.Sprintf("resourceIDs.%s", name)),
		}
	}
	return resouceIDs
}
