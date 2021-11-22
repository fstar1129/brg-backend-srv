package eth

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"time"

	hubEth "gitlab.nekotal.tech/lachain/crosschain/bridge-backend-service/src/service/workers/eth-compatible/abi/relayer-hub/eth"
	hubLA "gitlab.nekotal.tech/lachain/crosschain/bridge-backend-service/src/service/workers/eth-compatible/abi/relayer-hub/la"
	"gitlab.nekotal.tech/lachain/crosschain/bridge-backend-service/src/service/workers/utils"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"
	jsonrpc "github.com/ybbus/jsonrpc/v2"

	"gitlab.nekotal.tech/lachain/crosschain/bridge-backend-service/src/models"
	"gitlab.nekotal.tech/lachain/crosschain/bridge-backend-service/src/service/storage"
)

// Erc20Worker ...
type Erc20Worker struct {
	provider     string
	chainID      string
	logger       *logrus.Entry // logger
	config       *models.WorkerConfig
	client       *ethclient.Client
	contractAddr common.Address
}

// NewErc20Worker ...
func NewErc20Worker(logger *logrus.Logger, cfg *models.WorkerConfig) *Erc20Worker {
	client, err := ethclient.Dial(cfg.Provider)
	if err != nil {
		panic("new eth client error")
	}

	privKey, err := utils.GetPrivateKey(cfg)
	if err != nil {
		panic(fmt.Sprintf("generate private key error, err=%s", err.Error()))
	}

	publicKey := privKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		panic(fmt.Sprintf("get public key error, err=%s", err.Error()))
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	if !bytes.Equal(cfg.WorkerAddr.Bytes(), fromAddress.Bytes()) {
		panic(fmt.Sprintf(
			"relayer address supplied in config (%s) does not match mnemonic (%s)",
			cfg.WorkerAddr, fromAddress,
		))
	}

	// init token addresses
	return &Erc20Worker{
		chainID:      cfg.ChainID,
		logger:       logger.WithField("worker", cfg.ChainID),
		provider:     cfg.Provider,
		config:       cfg,
		client:       client,
		contractAddr: cfg.ContractAddr,
	}
}

// GetChain returns chain ID
func (w *Erc20Worker) GetChain() string {
	return string(w.chainID)
}

// GetStartHeight returns start blockchain height from config
func (w *Erc20Worker) GetStartHeight() (int64, error) {
	if w.config.StartBlockHeight == 0 {
		return w.GetHeight()
	}
	return w.config.StartBlockHeight, nil
}

// GetConfirmNum returns numbers of blocks after them tx will be confirmed
func (w *Erc20Worker) GetConfirmNum() int64 {
	return w.config.ConfirmNum
}

// Register ...
func (w *Erc20Worker) Register(relayerAddress common.Address) (string, error) {
	auth, err := w.getTransactor()
	if err != nil {
		return "", err
	}

	instance, err := hubEth.NewHubEth(w.contractAddr, w.client)
	if err != nil {
		return "", err
	}

	tx, err := instance.Register(auth, relayerAddress)
	if err != nil {
		return "", err
	}

	return tx.Hash().String(), nil
}

// Unregister ...
func (w *Erc20Worker) Unregister(relayerAddress common.Address) (string, error) {
	auth, err := w.getTransactor()
	if err != nil {
		return "", err
	}

	instance, err := hubEth.NewHubEth(w.contractAddr, w.client)
	if err != nil {
		return "", err
	}

	tx, err := instance.Unregister(auth, relayerAddress)
	if err != nil {
		return "", err
	}

	return tx.Hash().String(), nil
}

// Felony ...
func (w *Erc20Worker) Felony(relayerAddress common.Address, chainID string) (string, error) {
	auth, err := w.getTransactor()
	if err != nil {
		return "", err
	}
	if chainID != storage.LaChain {
		instance, err := hubEth.NewHubEth(w.contractAddr, w.client)
		if err != nil {
			return "", err
		}

		tx, err := instance.Felony(auth, relayerAddress)
		if err != nil {
			return "", err
		}

		return tx.Hash().String(), nil

	} else if chainID != storage.EthChain {
		instance, err := hubLA.NewHubLA(w.contractAddr, w.client)
		if err != nil {
			return "", err
		}

		tx, err := instance.Felony(auth, relayerAddress)
		if err != nil {
			return "", err
		}

		return tx.Hash().String(), nil

	}

	return "", nil
}

//Penalty
func (w *Erc20Worker) Penalty(relayerAddress common.Address, amount string) (string, error) {
	auth, err := w.getTransactor()
	if err != nil {
		return "", err
	}

	instance, err := hubLA.NewHubLA(w.contractAddr, w.client)
	if err != nil {
		return "", err
	}
	value, _ := new(big.Int).SetString(amount, 10)
	tx, err := instance.Slash(auth, relayerAddress, value)
	if err != nil {
		return "", err
	}

	return tx.Hash().String(), nil
}

// // GetStatus returns status of relayer account(balance eg)
// func (w *Erc20Worker) GetStatus(symbol string) (interface{}, error) {
// 	ethStatus := &EthStatus{}

// 	allowance, err := w.Allowance(symbol)
// 	if err != nil {
// 		return nil, err
// 	}
// 	ethStatus.Allowance = QuoBigInt(allowance, GetBigIntForDecimal(w.Config.ChainDecimal)).String()

// 	balance, err := w.Erc20Balance(w.Config.WorkerAddr, symbol)
// 	if err != nil {
// 		return nil, err
// 	}
// 	ethStatus.Erc20Balance = QuoBigInt(balance, GetBigIntForDecimal(w.Config.ChainDecimal)).String()

// 	ethBalance, err := w.EthBalance(w.Config.WorkerAddr)
// 	if err != nil {
// 		return nil, err
// 	}
// 	ethStatus.EthBalance = QuoBigInt(ethBalance, GetBigIntForDecimal(18)).String()

// 	return ethStatus, nil
// }

// GetBlockAndTxs ...
func (w *Erc20Worker) GetBlockAndTxs(height int64) (*models.BlockAndTxLogs, error) {
	var head *Header
	rpcClient := jsonrpc.NewClient(w.provider)

	resp, err := rpcClient.Call("eth_getBlockByNumber", fmt.Sprintf("0x%x", height), false)
	if err != nil {
		w.logger.Errorln("while call eth_getBlockByNumber, err = ", err)
		return nil, err
	}

	if err := resp.GetObject(&head); err != nil {
		w.logger.Errorln("while GetObject, err = ", err)
		return nil, err
	}

	if head == nil {
		return nil, fmt.Errorf("not found")
	}

	logs, err := w.getLogs(head.Hash)
	if err != nil {
		w.logger.Errorf("while getEvents(blockhash = %s), err = %v", head.Hash, err)
		return nil, err
	}

	return &models.BlockAndTxLogs{
		Height:          height,
		BlockHash:       head.Hash.String(),
		ParentBlockHash: head.ParentHash.Hex(),
		BlockTime:       int64(head.Time),
		TxLogs:          logs,
	}, nil
}

// GetFetchInterval ...
func (w *Erc20Worker) GetFetchInterval() time.Duration {
	return time.Duration(w.config.FetchInterval) * time.Second
}

// getLogs ...
func (w *Erc20Worker) getLogs(blockHash common.Hash) ([]*storage.TxLog, error) {
	//	topics := [][]common.Hash{{DepositEventHash, ProposalEventHash, ProposalVoteHash}}
	logs, err := w.client.FilterLogs(context.Background(), ethereum.FilterQuery{
		BlockHash: &blockHash,
		//	Topics:    topics,
		Addresses: []common.Address{w.contractAddr},
	})
	if err != nil {
		return nil, err
	}

	models := make([]*storage.TxLog, 0, len(logs))
	for _, log := range logs {
		w.logger.Infof("WORKER(%s) NEW EVENT: %v\n\n", w.chainID, log)
		event, err := ParseEvent(&log)
		if err != nil {
			w.logger.WithFields(logrus.Fields{"function": "GetLogs()"}).Errorf("parse event log error, err=%s", err)
			continue
		}
		if event == nil {
			continue
		}

		txLog := event.ToTxLog(w.chainID)
		txLog.Chain = w.chainID
		txLog.Height = int64(log.BlockNumber)
		txLog.EventID = log.TxHash.Hex()
		txLog.BlockHash = log.BlockHash.Hex()
		txLog.TxHash = log.TxHash.Hex()
		txLog.Status = storage.TxStatusInit

		models = append(models, txLog)
	}

	return models, nil
}

// GetHeight ..
func (w *Erc20Worker) GetHeight() (int64, error) {
	header, err := w.client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return 0, nil
	}
	return header.Number.Int64(), nil
}

// GetSentTxStatus ...
func (w *Erc20Worker) GetSentTxStatus(hash string) storage.TxStatus {
	// _, isPending, err := w.client.TransactionByHash(context.Background(), common.HexToHash(hash))
	// if err != nil {
	// 	w.logger.Errorln("GetSentTxStatus, err = ", err)
	// 	return storage.TxSentStatusNotFound
	// }

	// if isPending {
	// 	return storage.TxSentStatusPending
	// }

	txReceipt, err := w.client.TransactionReceipt(context.Background(), common.HexToHash(hash))
	if err != nil {
		return storage.TxSentStatusNotFound
	}

	if txReceipt.Status == types.ReceiptStatusFailed {
		return storage.TxSentStatusFailed
	}

	return storage.TxSentStatusSuccess
}

func (w *Erc20Worker) GetTxCountLatest() (uint64, error) {
	var result uint64
	rpcClient := jsonrpc.NewClient(w.provider)

	resp, err := rpcClient.Call("eth_getTransactionCount", w.config.WorkerAddr.Hex(), "latest")
	if err != nil {
		return 0, err
	}

	if err := resp.GetObject(&result); err != nil {
		return 0, err
	}

	return result, nil
}

// GetTransactor ...
func (w *Erc20Worker) getTransactor() (auth *bind.TransactOpts, err error) {
	privateKey, err := utils.GetPrivateKey(w.config)
	if err != nil {
		return nil, err
	}

	var nonce uint64
	// if w.chainID == storage.LaChain {
	// 	nonce, err = w.GetTxCountLatest()
	// 	if err != nil {
	// 		return nil, err
	// 	}

	// 	auth, err = bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(int64(26)))
	// 	if err != nil {
	// 		return nil, err
	// 	}

	// } else {
	nonce, err = w.client.PendingNonceAt(context.Background(), w.config.WorkerAddr)
	if err != nil {
		return nil, err
	}
	auth = bind.NewKeyedTransactor(privateKey)
	// }

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)                // in wei
	auth.GasLimit = uint64(w.config.GasLimit) // in units
	auth.GasPrice = w.config.GasPrice

	return auth, nil
}

// EthBalance ...
func (w *Erc20Worker) EthBalance(address common.Address) (*big.Int, error) {
	return w.client.BalanceAt(context.Background(), address, nil)
}

// GetWorkerAddress ...
func (w *Erc20Worker) GetWorkerAddress() string {
	return w.config.WorkerAddr.String()
}

// GetColdWalletAddress ...
func (w *Erc20Worker) GetColdWalletAddress() string {
	return w.config.ColdWalletAddr.String()
}

// IsSameAddress ...
func (w *Erc20Worker) IsSameAddress(addrA string, addrB string) bool {
	return bytes.Equal(common.FromHex(addrA), common.FromHex(addrB))
}

// SendAmount ...
func (w *Erc20Worker) SendAmount(address string, amount *big.Int) (string, error) {
	return "", fmt.Errorf("not implemented") // TODO
}
