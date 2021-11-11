package eth

import (
	"bytes"
	"fmt"
	"strings"

	"gitlab.nekotal.tech/lachain/crosschain/bridge-backend-service/src/service/storage"
	hub "gitlab.nekotal.tech/lachain/crosschain/bridge-backend-service/src/service/workers/eth-compatible/abi/relayer-hub"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

////
// EVENTS HASHES  | web3.utils.sha3('HTLT(types,...)');
////

const (
	// RelayerRegisterEventName ...
	RelayerRegisterEventName = "RelayerRegister"
)

// RegisterRelayer = 0x03f64c5eb6e7ea7290123c03dbf1d28d676a583815f5f6a8ebd7f153cfb45215
// RelayerUnRegister  = 0x013c40a2b0c93c7ba3e8b43e7529d1d6f020670e66817894c8f86708afb3e62c
// PenaltyForRelayer = 0xf771f6f0a64d5cbff2d92da9bcc1b2ff0c4ba763dec1a9a50a7b06a60429a93a
// RewardForRelayer = 0x823083e86e79d611ae5083c45dd88f8f38ed42a14f0884540d09a3273c6100e5
// RelayerBlocked = 0xa5ea068b2ed8e05403d639cea236649975a9712487c612f9a6945b6bad00d81d

var (
	// RelayerRegisterEventHash ...
	RelayerRegisterEventHash = common.HexToHash("0x03f64c5eb6e7ea7290123c03dbf1d28d676a583815f5f6a8ebd7f153cfb45215")
)

// DepositEvent represents a Deposit event raised by the Bridge.sol contract.
type RelayerRegisterEvent struct {
	Relayer common.Address
}

// ParseRelayerRegisterEvent ...
func ParseRelayerRegisterEvent(abi *abi.ABI, log *types.Log) (ContractEvent, error) {
	var ev RelayerRegisterEvent
	ev.Relayer = common.HexToAddress(log.Topics[1].Hex())
	fmt.Printf("RelayerRegister\n")
	fmt.Printf("relayer address: %s\n", ev.Relayer.Hex())

	return ev, nil
}

// !!! TODO !!!

// ToTxLog ...
func (ev RelayerRegisterEvent) ToTxLog(chain string) *storage.TxLog {
	return &storage.TxLog{
		Chain:  chain,
		TxType: storage.TxTypeRegister,
		Data:   ev.Relayer.Hex(),
	}
}

// ParseEvent ...
func ParseEvent(log *types.Log) (ContractEvent, error) {
	if bytes.Equal(log.Topics[0][:], RelayerRegisterEventHash[:]) {
		abi, _ := abi.JSON(strings.NewReader(hub.HubABI))
		return ParseRelayerRegisterEvent(&abi, log)
	}

	return nil, nil
}

// ContractEvent ...
type ContractEvent interface {
	ToTxLog(chain string) *storage.TxLog
}
