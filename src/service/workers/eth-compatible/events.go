package eth

import (
	"bytes"
	"fmt"
	"math/big"
	"strings"

	"gitlab.nekotal.tech/lachain/crosschain/bridge-backend-service/src/service/storage"
	ethBr "gitlab.nekotal.tech/lachain/crosschain/bridge-backend-service/src/service/workers/eth-compatible/abi/bridge/eth"
	laBr "gitlab.nekotal.tech/lachain/crosschain/bridge-backend-service/src/service/workers/eth-compatible/abi/bridge/la"
	"gitlab.nekotal.tech/lachain/crosschain/bridge-backend-service/src/service/workers/utils"

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

func ParseLAProposalEvent(abi *abi.ABI, log *types.Log) (ContractEvent, error) {
	var ev ProposalEvent
	if err := abi.UnpackIntoInterface(&ev, ProposalEventName, log.Data); err != nil {
		return nil, err
	}

	fmt.Printf("ProposalEvent\n")
	fmt.Printf("origin chain ID: 0x%s\n", common.Bytes2Hex(ev.OriginChainID[:]))
	fmt.Printf("destination chain ID: 0x%s\n", common.Bytes2Hex(ev.DestinationChainID[:]))
	fmt.Printf("deposit nonce: %d\n", ev.DepositNonce)
	fmt.Printf("status: %d\n", ev.Status)
	fmt.Printf("resource ID: 0x%s\n", common.Bytes2Hex(ev.ResourceID[:]))
	fmt.Printf("DataHash: 0x%s\n\n", common.Bytes2Hex(ev.DataHash[:]))
	fmt.Printf("amount: ", ev.Amount.String())

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

	fmt.Printf("ProposalEvent\n")
	fmt.Printf("origin chain ID: 0x%s\n", common.Bytes2Hex(ev.OriginChainID[:]))
	fmt.Printf("destination chain ID: 0x%s\n", common.Bytes2Hex(ev.DestinationChainID[:]))
	fmt.Printf("deposit nonce: %d\n", ev.DepositNonce)
	fmt.Printf("status: %d\n", ev.Status)
	fmt.Printf("resource ID: 0x%s\n", common.Bytes2Hex(ev.ResourceID[:]))
	fmt.Printf("DataHash: 0x%s\n\n", common.Bytes2Hex(ev.DataHash[:]))
	fmt.Printf("amount: ", ev.Amount.String())

	return ev, nil
}

// !!! TODO !!!

// ToTxLog ...
func (ev ProposalEvent) ToTxLog(chain string) *storage.TxLog {
	// if status == 2 -> already claimed -> mint
	// if status == 3-> already minted(executed)
	txlog := &storage.TxLog{
		Chain:              chain,
		TxType:             storage.TxTypeClaim,
		ReceiverAddr:       ev.RecipientAddress.String(),
		OutAmount:          ev.Amount.String(),
		SwapID:             utils.CalcutateSwapID(string(ev.DataHash[:]), string(ev.DepositNonce)),
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
		if w.GetChainName() == storage.LaChain {
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
