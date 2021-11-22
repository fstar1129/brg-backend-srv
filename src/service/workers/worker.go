package workers

import (
	"time"

	"github.com/ethereum/go-ethereum/common"
	"gitlab.nekotal.tech/lachain/crosschain/bridge-backend-service/src/models"
	"gitlab.nekotal.tech/lachain/crosschain/bridge-backend-service/src/service/storage"
)

// IWorker ...
type IWorker interface {
	// GetChain returns unique name of the chain(like LA, ETH and etc)
	GetChain() string
	// GetWokrerAddress returns worker address
	GetWorkerAddress() string
	// GetStartHeight returns blockchain start height for watcher
	GetStartHeight() (int64, error)
	// GetConfirmNum returns numbers of blocks after them tx will be confirmed
	GetConfirmNum() int64
	// GetHeight returns current height of chain
	GetHeight() (int64, error)
	// GetBlockAndTxs returns block info and txs included in this block
	GetBlockAndTxs(height int64) (*models.BlockAndTxLogs, error)
	// GetFetchInterval returns fetch interval of the chain like average blocking time, it is used in observer
	GetFetchInterval() time.Duration
	// GetWorkerAddress returns relayer account address
	// GetWorkerAddress() string
	// // GetColdWalletAddress returns the address of the relayer's cold wallet
	// GetColdWalletAddress() string
	// GetSentTxStatus returns status of tx sent
	GetSentTxStatus(hash string) storage.TxStatus
	// // GetBalance returns balance of swap token for any address
	// GetBalance(address, tokenSymbol string) (*big.Int, error)
	// // GetStatus returns status of relayer account(balance eg)
	// GetStatus(tokenSymbol string) (interface{}, error)
	// // GetBalanceAlertMsg returns balance alert message if necessary, like account balance is less than amount in config
	// GetBalanceAlertMsg(tokenSymbol string) (string, error)
	// IsSameAddress returns is addrA the same with addrB
	IsSameAddress(addrA string, addrB string) bool
	// CalcSwapID calculate swap id for each chain
	//CalcSwapID(randomNumberHash common.Hash, sender string, senderOtherChain string) ([]byte, error)
	// Refundable returns is swap refundable
	//Refundable(swapID common.Hash) (bool, error)
	// GetSwap returns swap request detail
	//	GetSwap(swapID common.Hash) (*models.SwapRequest, error)
	// HasSwap returns does swap exist
	// HasSwap(swapID common.Hash) (bool, error)
	// HTLT sends htlt tx

	// Register
	Register(relayerAddress common.Address) (string, error)

	// Unregister
	Unregister(relayerAddress common.Address) (string, error)

	//Felony
	Felony(relayerAddress common.Address, chainID string) (string, error)

	// Penalty
	Penalty(relayerAddress common.Address, amount string) (string, error)

	//Reward
	Reward(relayerAddress common.Address, amount string) (string, error)

	//HTLT(erc20TokenAddr, lrc20TokenAddr, recipientAddr, otherChainRecipientAddr string, timestamp int64,
	//	heightSpan int64, outAmount *big.Int) (string, error)
	// CreateRequest sends wrapped tokens tx
	//	CreateRequest(swapID common.Hash) (string, error)
	// SendClaim
	//	SendClaim(swapID common.Hash) (string, error)
	// Refund sends refund tx
	//Refund(swapID common.Hash) (string, error)
	// SendAmount
	//SendAmount(address string, amount *big.Int) (string, error)
}
