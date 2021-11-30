package eth

import (
	"bytes"
	"fmt"
	"math/big"
	"strings"

	"gitlab.nekotal.tech/lachain/crosschain/bridge-backend-service/src/service/storage"
	ethBr "gitlab.nekotal.tech/lachain/crosschain/bridge-backend-service/src/service/workers/eth-compatible/abi/bridge/eth"
	laBr "gitlab.nekotal.tech/lachain/crosschain/bridge-backend-service/src/service/workers/eth-compatible/abi/bridge/la"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"gitlab.nekotal.tech/lachain/crosschain/bridge-backend-service/src/service/workers/utils"
)

////
// EVENTS HASHES  | web3.utils.sha3('HTLT(types,...)');
////

const (
	ProposalEventName = "ProposalEvent"
)

var (
	ProposalEventHash = common.HexToHash("0x8dc49847a011c3b316cd0f50cf982e0fd5b3ddb7fdf970fc81a25557f0923a73")
)

// ProposalEvent represents a ProposalEvent event raised by the Bridge.sol contract.
type ProposalEvent struct {
	DestinationChainID [8]byte
	RecipientAddress   common.Address
	Amount             *big.Int
	DepositNonce       uint64
	Status             uint8
	ResourceID         [32]byte
	DataHash           [32]byte
	Raw                types.Log // Blockchain specific contextual infos
}

func ParseEthProposalEvent(abi *abi.ABI, log *types.Log) (ContractEvent, error) {
	var ev ProposalEvent
	// abi, _ := abi.JSON(strings.NewReader(ethBr.EthBrABI))
	if err := abi.UnpackIntoInterface(&ev, ProposalEventName, log.Data); err != nil {
		return nil, err
	}
	ev.DestinationChainID = utils.BytesToBytes8(log.Topics[1].Bytes())
	ev.RecipientAddress = common.HexToAddress(log.Topics[1].Hex())
	ev.Amount = big.NewInt(0).SetBytes(log.Topics[3].Bytes())

	fmt.Printf("ProposalEvent\n")
	fmt.Printf("destination chain ID: 0x%s\n", common.Bytes2Hex(ev.DestinationChainID[:]))
	fmt.Printf("deposit nonce: %d\n", ev.DepositNonce)
	fmt.Printf("status: %d\n", ev.Status)
	fmt.Printf("resource ID: 0x%s\n", common.Bytes2Hex(ev.ResourceID[:]))
	fmt.Printf("DataHash: 0x%s\n\n", common.Bytes2Hex(ev.DataHash[:]))

	return ev, nil
}

func ParseLAProposalEvent(abi *abi.ABI, log *types.Log) (ContractEvent, error) {
	var ev ProposalEvent
	// abi, _ := abi.JSON(strings.NewReader(ethBr.EthBrABI))
	if err := abi.UnpackIntoInterface(&ev, ProposalEventName, log.Data); err != nil {
		return nil, err
	}

	fmt.Printf("ProposalEvent\n")
	fmt.Printf("destination chain ID: 0x%s\n", common.Bytes2Hex(ev.DestinationChainID[:]))
	fmt.Printf("deposit nonce: %d\n", ev.DepositNonce)
	fmt.Printf("status: %d\n", ev.Status)
	fmt.Printf("resource ID: 0x%s\n", common.Bytes2Hex(ev.ResourceID[:]))
	fmt.Printf("DataHash: 0x%s\n\n", common.Bytes2Hex(ev.DataHash[:]))

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
		SwapID:             fmt.Sprintf("0x%s", common.Bytes2Hex(ev.DataHash[:])),
		DestinationChainID: common.Bytes2Hex(ev.DestinationChainID[:]),
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
		if w.GetChain() == storage.EthChain {
			abi, _ := abi.JSON(strings.NewReader(ethBr.EthBrABI))
			return ParseEthProposalEvent(&abi, log)
		} else if w.GetChain() == storage.LaChain {
			abi, _ := abi.JSON(strings.NewReader(laBr.LaBrABI))
			return ParseLAProposalEvent(&abi, log)
		}
	}
	return nil, nil
}

// ContractEvent ...
type ContractEvent interface {
	ToTxLog(chain string) *storage.TxLog
}
