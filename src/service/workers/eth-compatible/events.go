package eth

import (
	"bytes"
	"fmt"
	"math/big"
	"strings"

	"gitlab.nekotal.tech/lachain/crosschain/bridge-backend-service/src/service/storage"
	hubEth "gitlab.nekotal.tech/lachain/crosschain/bridge-backend-service/src/service/workers/eth-compatible/abi/relayer-hub/eth"
	hubLA "gitlab.nekotal.tech/lachain/crosschain/bridge-backend-service/src/service/workers/eth-compatible/abi/relayer-hub/la"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

////
// EVENTS HASHES  | web3.utils.sha3('HTLT(types,...)');
////

const (
	// RelayerRegisterEventName ...
	RelayerRegisterEventName   = "RelayerRegister"
	RelayerUnregisterEventName = "RelayerUnRegister"
	RelayerBlockedEventName    = "RelayerBlocked"
	PenaltyForRelayerEventName = "PenaltyForRelayer"
	RewardForRelayerEventName  = "RewardForRelayer"
)

// RegisterRelayer = 0x03f64c5eb6e7ea7290123c03dbf1d28d676a583815f5f6a8ebd7f153cfb45215
// RelayerUnRegister  = 0x013c40a2b0c93c7ba3e8b43e7529d1d6f020670e66817894c8f86708afb3e62c
// PenaltyForRelayer = 0xf771f6f0a64d5cbff2d92da9bcc1b2ff0c4ba763dec1a9a50a7b06a60429a93a
// RewardForRelayer = 0x63dd553807ea27cb996173133b29a47f29c6b3bec34043271f5a5f7dd1f949c9
// RelayerBlocked = 0xa5ea068b2ed8e05403d639cea236649975a9712487c612f9a6945b6bad00d81d

var (
	RelayerRegisterEventHash   = common.HexToHash("0x03f64c5eb6e7ea7290123c03dbf1d28d676a583815f5f6a8ebd7f153cfb45215")
	RelayerUnregisterEventHash = common.HexToHash("0x013c40a2b0c93c7ba3e8b43e7529d1d6f020670e66817894c8f86708afb3e62c")
	RelayerBlockedEventHash    = common.HexToHash("0xa5ea068b2ed8e05403d639cea236649975a9712487c612f9a6945b6bad00d81d")
	PenaltyForRelayerEventHash = common.HexToHash("0x73f262bdf0c3145b25a03d5c75d5989bdab8ecb77968d7ff6de8a3bd83b8b13b")
	RewardForRelayerEventHash  = common.HexToHash("0x63dd553807ea27cb996173133b29a47f29c6b3bec34043271f5a5f7dd1f949c9")
)

// DepositEvent represents a Deposit event raised by the Bridge.sol contract.
type RelayerRegisterEvent struct {
	Relayer common.Address
	Raw     types.Log
}

type RelayerUnregisterEvent struct {
	Relayer common.Address
	Raw     types.Log
}

type RelayerBlockedEvent struct {
	Relayer common.Address
	Raw     types.Log
}

type PenaltyForRelayerEvent struct {
	Reason  string
	Relayer common.Address
	Penalty *big.Int
	Raw     types.Log
}

type RewardForRelayerEvent struct {
	Relayer common.Address
	Reward  *big.Int
	Raw     types.Log
}

// ParseRelayerRegisterEvent ...
func ParseRelayerRegisterEvent(abi *abi.ABI, log *types.Log) (ContractEvent, error) {
	var ev RelayerRegisterEvent
	if err := abi.UnpackIntoInterface(&ev, RelayerRegisterEventName, log.Data); err != nil {
		fmt.Println("AA")
		return nil, err
	}
	fmt.Printf("RelayerRegister\n")
	fmt.Printf("relayer address: %s\n", ev.Relayer)

	return ev, nil
}

func ParseRelayerUnregisterEvent(abi *abi.ABI, log *types.Log) (ContractEvent, error) {
	var ev RelayerUnregisterEvent
	if err := abi.UnpackIntoInterface(&ev, RelayerUnregisterEventName, log.Data); err != nil {
		fmt.Println("AA", err)
		return nil, err
	}
	fmt.Printf("RelayerUnregister\n")
	fmt.Printf("relayer address: %s\n", ev.Relayer.Hex())

	return ev, nil
}

func ParseRelayerBlockedEvent(abi *abi.ABI, log *types.Log) (ContractEvent, error) {
	var ev RelayerBlockedEvent
	if err := abi.UnpackIntoInterface(&ev, RelayerBlockedEventName, log.Data); err != nil {
		fmt.Println("AA", err)
		return nil, err
	}
	fmt.Printf("RelayerBlocked\n")
	fmt.Printf("relayer address: %s\n", ev.Relayer.Hex())

	return ev, nil
}

func ParsePenaltyForRelayerEvent(abi *abi.ABI, log *types.Log) (ContractEvent, error) {
	var ev PenaltyForRelayerEvent
	if err := abi.UnpackIntoInterface(&ev, PenaltyForRelayerEventName, log.Data); err != nil {
		fmt.Println("AA", err)
		return nil, err
	}
	ev.Relayer = common.BytesToAddress(log.Topics[1][:])
	fmt.Printf("PenaltyForRelayer\n")
	fmt.Printf("relayer address: %s with penalty %s\n", ev.Relayer.Hex(), ev.Penalty.String())

	return ev, nil
}

func ParseRewardForRelayerEvent(abi *abi.ABI, log *types.Log) (ContractEvent, error) {
	var ev RewardForRelayerEvent
	if err := abi.UnpackIntoInterface(&ev, RewardForRelayerEventName, log.Data); err != nil {
		fmt.Println("AA", err)
		return nil, err
	}
	fmt.Printf("RewardForRelayer\n")
	fmt.Printf("relayer address: %s with reward %s\n", ev.Relayer.Hex(), ev.Reward.String())

	return ev, nil
}

// !!! TODO !!!

// ToTxLog ...
func (ev RelayerRegisterEvent) ToTxLog(chain string) *storage.TxLog {
	return &storage.TxLog{
		Chain:       chain,
		TxType:      storage.TxTypeRegister,
		Data:        ev.Relayer.Hex(),
		Status:      storage.TxStatusInit,
		EventStatus: storage.EventStatusRegisterInit,
	}
}

func (ev RelayerUnregisterEvent) ToTxLog(chain string) *storage.TxLog {
	return &storage.TxLog{
		Chain:       chain,
		TxType:      storage.TxTypeUnregister,
		Data:        ev.Relayer.Hex(),
		EventStatus: storage.EventStatusUnregisterInit,
	}
}

func (ev RelayerBlockedEvent) ToTxLog(chain string) *storage.TxLog {
	return &storage.TxLog{
		Chain:       chain,
		TxType:      storage.TxTypeFelony,
		Data:        ev.Relayer.Hex(),
		EventStatus: storage.EventStatusFelonyInit,
	}
}

func (ev PenaltyForRelayerEvent) ToTxLog(chain string) *storage.TxLog {
	return &storage.TxLog{
		Chain:       chain,
		TxType:      storage.TxTypePenalty,
		Data:        ev.Relayer.Hex(),
		Penalty:     ev.Penalty.String(),
		EventStatus: storage.EventStatusPenaltyInit,
	}
}

func (ev RewardForRelayerEvent) ToTxLog(chain string) *storage.TxLog {
	return &storage.TxLog{
		Chain:       chain,
		TxType:      storage.TxTypeReward,
		Data:        ev.Relayer.Hex(),
		Penalty:     ev.Reward.String(),
		EventStatus: storage.EventStatusRewardInit,
	}
}

// ParseEvent ...
func ParseEvent(log *types.Log) (ContractEvent, error) {
	if bytes.Equal(log.Topics[0][:], RelayerRegisterEventHash[:]) {
		abi, _ := abi.JSON(strings.NewReader(hubLA.HubLAABI))
		return ParseRelayerRegisterEvent(&abi, log)
	} else if bytes.Equal(log.Topics[0][:], RelayerUnregisterEventHash[:]) {
		abi, _ := abi.JSON(strings.NewReader(hubLA.HubLAABI))
		return ParseRelayerUnregisterEvent(&abi, log)
	} else if bytes.Equal(log.Topics[0][:], RelayerBlockedEventHash[:]) {
		abi, _ := abi.JSON(strings.NewReader(hubLA.HubLAABI))
		return ParseRelayerBlockedEvent(&abi, log)
	} else if bytes.Equal(log.Topics[0][:], PenaltyForRelayerEventHash[:]) {
		abi, _ := abi.JSON(strings.NewReader(hubEth.HubEthABI))
		return ParsePenaltyForRelayerEvent(&abi, log)
	} else if bytes.Equal(log.Topics[0][:], RewardForRelayerEventHash[:]) {
		abi, _ := abi.JSON(strings.NewReader(hubEth.HubEthABI))
		return ParseRewardForRelayerEvent(&abi, log)
	}

	return nil, nil
}

// ContractEvent ...
type ContractEvent interface {
	ToTxLog(chain string) *storage.TxLog
}
