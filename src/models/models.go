package models

import (
	"math/big"
	"time"

	"github.com/latoken/bridge-backend-service/src/service/storage"

	"github.com/ethereum/go-ethereum/common"
)

// ServiceConfig contains configurations for rest-api service.
type ServiceConfig struct {
	Host string // Service Host
	Port string // Service port
}

// RelayerStatus ...
type RelayerStatus struct {
	Mode    string                  `json:"mode"`
	Workers map[string]WorkerStatus `json:"workers"`
}

// WorkerStatus ...
type WorkerStatus struct {
	Height             int64       `json:"height"`
	SyncHeight         int64       `json:"sync_height"`
	LastBlockFetchedAt time.Time   `json:"last_block_fetched_at"`
	Status             interface{} `json:"status"`
}

// BlockAndTxLogs ...
type BlockAndTxLogs struct {
	Height          int64
	BlockHash       string
	ParentBlockHash string
	BlockTime       int64
	TxLogs          []*storage.TxLog
}

// SwapRequest ...
type SwapRequest struct {
	ID                   common.Hash
	RandomNumberHash     common.Hash
	ExpireHeight         int64
	SenderAddress        string
	RecipientAddress     string
	RecipientWorkerChain string
	OutAmount            *big.Int
}

// StorageConfig contains configurations for storage, postgreSQL
type StorageConfig struct {
	URL        string // DataBase URL for connection
	DBDriver   string // DataBase driver
	DBHOST     string // DataBase host
	DBPORT     int64
	DBSSL      string // DataBase sslmode
	DBName     string // DataBase name
	DBUser     string // DataBase's user
	DBPassword string // User's password
}

// WorkerConfig ...
type WorkerConfig struct {
	NetworkType        string         `json:"type"`
	ChainID            int64          `json:"chain_id"`
	ChainName          string         `json:"chain_id"`
	PrivateKey         string         `json:"private_key"`
	Provider           string         `json:"provider"`
	ContractAddr       common.Address `json:"contract_addr"`
	AaveLPAddress      common.Address `json:"aave_lp_addr"`
	USDTContractAddr   common.Address `json:"USDT_token_addr"`
	WorkerAddr         common.Address `json:"worker_addr"`
	ColdWalletAddr     common.Address `json:"cold_wallet_addr"`
	FetchInterval      int64          `json:"fetch_interval"`
	GasLimit           int64          `json:"gas_limit"`
	GasPrice           *big.Int       `json:"gas_price"`
	ChainDecimal       int            `json:"chain_decimal"`
	ConfirmNum         int64          `json:"confirm_num"`
	StartBlockHeight   int64          `json:"start_block_height"`
	DestinationChainID string         `json:"dest_id"`
}

type TssConfig struct {
	Address    string
	BaseFolder string
}

// RelayerConfig ...
type RelayerConfig struct {
	Address common.Address
}

// FetcherConfig
type FetcherConfig struct {
	ChainName string
	URL       string
}
