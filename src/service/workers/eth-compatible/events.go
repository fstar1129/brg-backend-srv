package eth

import (
	"bytes"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/latoken/bridge-backend-service/src/service/storage"
	ethBr "github.com/latoken/bridge-backend-service/src/service/workers/eth-compatible/abi/bridge/eth"
	laBr "github.com/latoken/bridge-backend-service/src/service/workers/eth-compatible/abi/bridge/la"
	"github.com/latoken/bridge-backend-service/src/service/workers/utils"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

////
// EVENTS HASHES  | web3.utils.sha3('HTLT(types,...)');
////

const (
	ProposalEventName = "ProposalEvent"
)

var (
	ProposalEventHash = common.HexToHash("0x9686dcabd0450cad86a88df15a9d35b08b35d1b08a19008df37cf8538c467516")
)

var txStatus = make(map[string]uint8)

// ProposalEvent represents a ProposalEvent event raised by the Bridge.sol contract.
type ProposalEvent struct {
	OriginChainID      [8]byte
	DestinationChainID [8]byte
	RecipientAddress   common.Address
	Amount             *big.Int
	DepositNonce       uint64
	Status             uint8
	ResourceID         [32]byte
	DataHash           [32]byte
	Raw                types.Log // Blockchain specific contextual infos
}

func setTxMonitor(SwapID string, Status uint8) {
	if Status == 1 {
		txStatus[SwapID] = Status
		go func(SwapID string, Status uint8) {
			time.Sleep(5 * 60 * time.Second)
			if NewStatus, ok := txStatus[SwapID]; ok {
				if NewStatus != 3 {
					fmt.Printf("ERROR[%s] SwapID %s stuck in status %d\n\n", time.Now().Format(time.RFC3339Nano), SwapID, NewStatus)
				}
				delete(txStatus, SwapID)
			}
		}(SwapID, Status)
	} else {
		if HashStatus, ok := txStatus[SwapID]; ok {
			if HashStatus < Status {
				txStatus[SwapID] = Status
			}
		}
	}
}

func ParseLAProposalEvent(abi *abi.ABI, log *types.Log) (ContractEvent, error) {
	var ev ProposalEvent
	if err := abi.UnpackIntoInterface(&ev, ProposalEventName, log.Data); err != nil {
		return nil, err
	}

	var event_time = time.Now().Format(time.RFC3339Nano)
	var SwapID = ev.CalcutateSwapID()
	fmt.Printf("INFO[%s] SwapID: %s\n", event_time, SwapID)
	fmt.Printf("INFO[%s] ProposalEvent\n", event_time)
	fmt.Printf("INFO[%s] origin chain ID: 0x%s\n", event_time, common.Bytes2Hex(ev.OriginChainID[:]))
	fmt.Printf("INFO[%s] destination chain ID: 0x%s\n", event_time, common.Bytes2Hex(ev.DestinationChainID[:]))
	fmt.Printf("INFO[%s] deposit nonce: %d\n", event_time, ev.DepositNonce)
	fmt.Printf("INFO[%s] status: %d\n", event_time, ev.Status)
	fmt.Printf("INFO[%s] resource ID: 0x%s\n", event_time, common.Bytes2Hex(ev.ResourceID[:]))
	fmt.Printf("INFO[%s] recipient: 0x%s\n", event_time, common.Bytes2Hex(ev.RecipientAddress[:]))
	fmt.Printf("INFO[%s] DataHash: 0x%s\n", event_time, common.Bytes2Hex(ev.DataHash[:]))
	fmt.Printf("INFO[%s] amount: %s\n\n", event_time, ev.Amount.String())

	setTxMonitor(SwapID, ev.Status)

	return ev, nil
}

func ParseETHProposalEvent(abi *abi.ABI, log *types.Log) (ContractEvent, error) {
	var ev ProposalEvent
	if err := abi.UnpackIntoInterface(&ev, ProposalEventName, log.Data); err != nil {
		return nil, err
	}
	ev.OriginChainID = utils.BytesToBytes8(log.Topics[1].Bytes())
	ev.DestinationChainID = utils.BytesToBytes8(log.Topics[2].Bytes())
	ev.RecipientAddress = common.BytesToAddress(log.Topics[3].Bytes())

	var event_time = time.Now().Format(time.RFC3339Nano)
	var SwapID = ev.CalcutateSwapID()
	fmt.Printf("INFO[%s] ProposalEvent\n", event_time)
	fmt.Printf("INFO[%s] SwapID: %s\n", event_time, SwapID)
	fmt.Printf("INFO[%s] origin chain ID: 0x%s\n", event_time, common.Bytes2Hex(ev.OriginChainID[:]))
	fmt.Printf("INFO[%s] destination chain ID: 0x%s\n", event_time, common.Bytes2Hex(ev.DestinationChainID[:]))
	fmt.Printf("INFO[%s] deposit nonce: %d\n", event_time, ev.DepositNonce)
	fmt.Printf("INFO[%s] status: %d\n", event_time, ev.Status)
	fmt.Printf("INFO[%s] resource ID: 0x%s\n", event_time, common.Bytes2Hex(ev.ResourceID[:]))
	fmt.Printf("INFO[%s] DataHash: 0x%s\n", event_time, common.Bytes2Hex(ev.DataHash[:]))
	fmt.Printf("INFO[%s] amount: %s\n\n", event_time, ev.Amount.String())

	setTxMonitor(SwapID, ev.Status)

	return ev, nil
}

// !!! TODO !!!
func (ev ProposalEvent) CalcutateSwapID() string {
	return utils.CalcutateSwapID(common.Bytes2Hex(ev.OriginChainID[:]), common.Bytes2Hex(ev.DestinationChainID[:]), fmt.Sprint(ev.DepositNonce))
}

// ToTxLog ...
func (ev ProposalEvent) ToTxLog(chain string) *storage.TxLog {
	// if status == 2 -> already claimed -> mint
	// if status == 3-> already minted(executed)
	txlog := &storage.TxLog{
		Chain:              chain,
		TxType:             storage.TxTypeClaim,
		ReceiverAddr:       ev.RecipientAddress.String(),
		OutAmount:          ev.Amount.String(),
		SwapID:             ev.CalcutateSwapID(),
		DestinationChainID: common.Bytes2Hex(ev.DestinationChainID[:]),
		OriginСhainID:      common.Bytes2Hex(ev.OriginChainID[:]),
		DepositNonce:       ev.DepositNonce,
		SwapStatus:         ev.Status,
		ResourceID:         common.Bytes2Hex(ev.ResourceID[:]),
	}

	if ev.Status == uint8(2) {
		txlog.TxType = storage.TxTypePassed
	} else if ev.Status == uint8(3) {
		txlog.TxType = storage.TxTypeSpend
	}

	return txlog
}

// ParseEvent ...
func (w *Erc20Worker) parseEvent(log *types.Log) (ContractEvent, error) {
	if bytes.Equal(log.Topics[0][:], ProposalEventHash[:]) {
		if w.GetChainName() == "LA" {
			abi, _ := abi.JSON(strings.NewReader(laBr.LaBrABI))
			return ParseLAProposalEvent(&abi, log)
		} else {
			abi, _ := abi.JSON(strings.NewReader(ethBr.EthBrABI))
			return ParseETHProposalEvent(&abi, log)
		}
	}
	return nil, nil
}

// ContractEvent ...
type ContractEvent interface {
	ToTxLog(chain string) *storage.TxLog
}
