// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package labr

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// BridgeProposal is an auto generated low-level Go binding around an user-defined struct.
type BridgeProposal struct {
	ResourceID    [32]byte
	DataHash      [32]byte
	YesVotes      []common.Address
	NoVotes       []common.Address
	Status        uint8
	ProposedBlock *big.Int
}

// LabrMetaData contains all meta data concerning the Labr contract.
var LabrMetaData = &bind.MetaData{
	ABI: "[{\"name\":\"paused\",\"type\":\"function\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"name\":\"sub\",\"type\":\"function\",\"inputs\":[{\"name\":\"a\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"b\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"pure\"},{\"name\":\"_isInitialised\",\"type\":\"function\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"name\":\"ownableInit\",\"type\":\"function\",\"inputs\":[{\"name\":\"owner_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"name\":\"owner\",\"type\":\"function\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"name\":\"renounceOwnership\",\"type\":\"function\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"name\":\"transferOwnership\",\"type\":\"function\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"name\":\"getEthSignedMessageHash\",\"type\":\"function\",\"inputs\":[{\"name\":\"_messageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"name\":\"verify\",\"type\":\"function\",\"inputs\":[{\"name\":\"_ethSignedMessageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_signer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"name\":\"getSigner\",\"type\":\"function\",\"inputs\":[{\"name\":\"messageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"pure\"},{\"name\":\"createMesssageHash\",\"type\":\"function\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"bytes8\",\"internalType\":\"bytes8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"name\":\"implementation\",\"type\":\"function\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"name\":\"_chainID\",\"type\":\"function\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes8\",\"internalType\":\"bytes8\"}],\"stateMutability\":\"view\"},{\"name\":\"_relayerThreshold\",\"type\":\"function\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"name\":\"_totalProposals\",\"type\":\"function\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"name\":\"_fee\",\"type\":\"function\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"name\":\"_expiry\",\"type\":\"function\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"name\":\"_relayerHubAddress\",\"type\":\"function\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"name\":\"_backendSrvAddress\",\"type\":\"function\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"name\":\"_balancerAddress\",\"type\":\"function\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"name\":\"_nativeResourceID\",\"type\":\"function\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"name\":\"_depositCounts\",\"type\":\"function\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes8\",\"internalType\":\"bytes8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"name\":\"_resourceIDToHandlerAddress\",\"type\":\"function\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"name\":\"_depositRecords\",\"type\":\"function\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"bytes8\",\"internalType\":\"bytes8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"name\":\"_proposals\",\"type\":\"function\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structBridge.Proposal\",\"components\":[{\"name\":\"_resourceID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_dataHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_yesVotes\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"_noVotes\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"_status\",\"type\":\"uint8\",\"internalType\":\"enumBridge.ProposalStatus\"},{\"name\":\"_proposedBlock\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"name\":\"extraLATransferred\",\"type\":\"function\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"name\":\"_hasVotedOnProposal\",\"type\":\"function\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"name\":\"initialize\",\"type\":\"function\",\"inputs\":[{\"name\":\"chainID_\",\"type\":\"bytes8\",\"internalType\":\"bytes8\"},{\"name\":\"relayerThreshold_\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fee_\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"expiry_\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"ownerAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"initBackendSrvAddress_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"initBalancerAddress_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"name\":\"adminSetRelayerHub\",\"type\":\"function\",\"inputs\":[{\"name\":\"newRelayerHub\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"name\":\"adminSetBackendSrvAddress\",\"type\":\"function\",\"inputs\":[{\"name\":\"newBackendSrv\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"name\":\"adminSetBalancerAddress\",\"type\":\"function\",\"inputs\":[{\"name\":\"newBalancer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"name\":\"adminPauseTransfers\",\"type\":\"function\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"name\":\"adminUnpauseTransfers\",\"type\":\"function\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"name\":\"adminChangeRelayerThreshold\",\"type\":\"function\",\"inputs\":[{\"name\":\"newThreshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"name\":\"adminSetResource\",\"type\":\"function\",\"inputs\":[{\"name\":\"handlerAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"resourceID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"name\":\"adminSetBurnable\",\"type\":\"function\",\"inputs\":[{\"name\":\"handlerAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"name\":\"adminSetNativeResourceID\",\"type\":\"function\",\"inputs\":[{\"name\":\"resourceID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"name\":\"getProposal\",\"type\":\"function\",\"inputs\":[{\"name\":\"originChainID\",\"type\":\"bytes8\",\"internalType\":\"bytes8\"},{\"name\":\"destinationChainID\",\"type\":\"bytes8\",\"internalType\":\"bytes8\"},{\"name\":\"depositNonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"recipientAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"resourceID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structBridge.Proposal\",\"components\":[{\"name\":\"_resourceID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_dataHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_yesVotes\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"_noVotes\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"_status\",\"type\":\"uint8\",\"internalType\":\"enumBridge.ProposalStatus\"},{\"name\":\"_proposedBlock\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"name\":\"adminChangeFee\",\"type\":\"function\",\"inputs\":[{\"name\":\"newFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"name\":\"adminWithdraw\",\"type\":\"function\",\"inputs\":[{\"name\":\"handlerAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amountOrTokenID\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"name\":\"deposit\",\"type\":\"function\",\"inputs\":[{\"name\":\"destinationChainID\",\"type\":\"bytes8\",\"internalType\":\"bytes8\"},{\"name\":\"resourceID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"recipientAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"params\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"name\":\"voteProposal\",\"type\":\"function\",\"inputs\":[{\"name\":\"originChainID\",\"type\":\"bytes8\",\"internalType\":\"bytes8\"},{\"name\":\"destinationChainID\",\"type\":\"bytes8\",\"internalType\":\"bytes8\"},{\"name\":\"depositNonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"resourceID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"recipientAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"name\":\"cancelProposal\",\"type\":\"function\",\"inputs\":[{\"name\":\"originChainID\",\"type\":\"bytes8\",\"internalType\":\"bytes8\"},{\"name\":\"destinationChainID\",\"type\":\"bytes8\",\"internalType\":\"bytes8\"},{\"name\":\"depositNonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"recipientAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"resourceID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"name\":\"executeProposal\",\"type\":\"function\",\"inputs\":[{\"name\":\"originChainID\",\"type\":\"bytes8\",\"internalType\":\"bytes8\"},{\"name\":\"destinationChainID\",\"type\":\"bytes8\",\"internalType\":\"bytes8\"},{\"name\":\"depositNonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"resourceID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"recipientAddress\",\"type\":\"address\",\"internalType\":\"addresspayable\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"params\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"name\":\"adminCollectFees\",\"type\":\"function\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"addresspayable\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"name\":\"adminApprove\",\"type\":\"function\",\"inputs\":[{\"name\":\"resourceIDOwner\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"resourceIDSpender\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"amountOrTokenID\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"name\":\"relayerCollectReward\",\"type\":\"function\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"name\":\"getExtraLATransferred\",\"type\":\"function\",\"inputs\":[{\"name\":\"originChainID\",\"type\":\"bytes8\",\"internalType\":\"bytes8\"},{\"name\":\"destinationChainID\",\"type\":\"bytes8\",\"internalType\":\"bytes8\"},{\"name\":\"depositNonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"resourceID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"recipientAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"name\":\"transferExtraFee\",\"type\":\"function\",\"inputs\":[{\"name\":\"originChainID\",\"type\":\"bytes8\",\"internalType\":\"bytes8\"},{\"name\":\"destinationChainID\",\"type\":\"bytes8\",\"internalType\":\"bytes8\"},{\"name\":\"depositNonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"resourceID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"name\":\"depositFunds\",\"type\":\"function\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"name\":\"Paused\",\"type\":\"event\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"anonymous\":false},{\"name\":\"Unpaused\",\"type\":\"event\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"anonymous\":false},{\"name\":\"OwnershipTransferred\",\"type\":\"event\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"internalType\":\"address\",\"indexed\":true},{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\",\"indexed\":true}],\"anonymous\":false},{\"name\":\"RelayerThresholdChanged\",\"type\":\"event\",\"inputs\":[{\"name\":\"newThreshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"anonymous\":false},{\"name\":\"Deposit\",\"type\":\"event\",\"inputs\":[{\"name\":\"originChainID\",\"type\":\"bytes8\",\"internalType\":\"bytes8\"},{\"name\":\"destinationChainID\",\"type\":\"bytes8\",\"internalType\":\"bytes8\"},{\"name\":\"resourceID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"depositNonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"depositor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"recipientAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"dataHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"name\":\"ProposalEvent\",\"type\":\"event\",\"inputs\":[{\"name\":\"originChainID\",\"type\":\"bytes8\",\"internalType\":\"bytes8\"},{\"name\":\"destinationChainID\",\"type\":\"bytes8\",\"internalType\":\"bytes8\"},{\"name\":\"recipientAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"depositNonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumBridge.ProposalStatus\"},{\"name\":\"resourceID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"dataHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"name\":\"ProposalVote\",\"type\":\"event\",\"inputs\":[{\"name\":\"originChainID\",\"type\":\"bytes8\",\"internalType\":\"bytes8\"},{\"name\":\"depositNonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumBridge.ProposalStatus\"},{\"name\":\"resourceID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"name\":\"RewardCollected\",\"type\":\"event\",\"inputs\":[{\"name\":\"relayer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"anonymous\":false},{\"name\":\"ExtraFeeTransferred\",\"type\":\"event\",\"inputs\":[{\"name\":\"originChainID\",\"type\":\"bytes8\",\"internalType\":\"bytes8\"},{\"name\":\"destinationChainID\",\"type\":\"bytes8\",\"internalType\":\"bytes8\"},{\"name\":\"depositNonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"resouceID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"anonymous\":false}]",
}

// LabrABI is the input ABI used to generate the binding from.
// Deprecated: Use LabrMetaData.ABI instead.
var LabrABI = LabrMetaData.ABI

// Labr is an auto generated Go binding around an Ethereum contract.
type Labr struct {
	LabrCaller     // Read-only binding to the contract
	LabrTransactor // Write-only binding to the contract
	LabrFilterer   // Log filterer for contract events
}

// LabrCaller is an auto generated read-only Go binding around an Ethereum contract.
type LabrCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LabrTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LabrTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LabrFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LabrFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LabrSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LabrSession struct {
	Contract     *Labr             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LabrCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LabrCallerSession struct {
	Contract *LabrCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// LabrTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LabrTransactorSession struct {
	Contract     *LabrTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LabrRaw is an auto generated low-level Go binding around an Ethereum contract.
type LabrRaw struct {
	Contract *Labr // Generic contract binding to access the raw methods on
}

// LabrCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LabrCallerRaw struct {
	Contract *LabrCaller // Generic read-only contract binding to access the raw methods on
}

// LabrTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LabrTransactorRaw struct {
	Contract *LabrTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLabr creates a new instance of Labr, bound to a specific deployed contract.
func NewLabr(address common.Address, backend bind.ContractBackend) (*Labr, error) {
	contract, err := bindLabr(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Labr{LabrCaller: LabrCaller{contract: contract}, LabrTransactor: LabrTransactor{contract: contract}, LabrFilterer: LabrFilterer{contract: contract}}, nil
}

// NewLabrCaller creates a new read-only instance of Labr, bound to a specific deployed contract.
func NewLabrCaller(address common.Address, caller bind.ContractCaller) (*LabrCaller, error) {
	contract, err := bindLabr(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LabrCaller{contract: contract}, nil
}

// NewLabrTransactor creates a new write-only instance of Labr, bound to a specific deployed contract.
func NewLabrTransactor(address common.Address, transactor bind.ContractTransactor) (*LabrTransactor, error) {
	contract, err := bindLabr(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LabrTransactor{contract: contract}, nil
}

// NewLabrFilterer creates a new log filterer instance of Labr, bound to a specific deployed contract.
func NewLabrFilterer(address common.Address, filterer bind.ContractFilterer) (*LabrFilterer, error) {
	contract, err := bindLabr(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LabrFilterer{contract: contract}, nil
}

// bindLabr binds a generic wrapper to an already deployed contract.
func bindLabr(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LabrABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Labr *LabrRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Labr.Contract.LabrCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Labr *LabrRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Labr.Contract.LabrTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Labr *LabrRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Labr.Contract.LabrTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Labr *LabrCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Labr.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Labr *LabrTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Labr.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Labr *LabrTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Labr.Contract.contract.Transact(opts, method, params...)
}

// BackendSrvAddress is a free data retrieval call binding the contract method 0x7f5d6a5e.
//
// Solidity: function _backendSrvAddress() view returns(address)
func (_Labr *LabrCaller) BackendSrvAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Labr.contract.Call(opts, &out, "_backendSrvAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BackendSrvAddress is a free data retrieval call binding the contract method 0x7f5d6a5e.
//
// Solidity: function _backendSrvAddress() view returns(address)
func (_Labr *LabrSession) BackendSrvAddress() (common.Address, error) {
	return _Labr.Contract.BackendSrvAddress(&_Labr.CallOpts)
}

// BackendSrvAddress is a free data retrieval call binding the contract method 0x7f5d6a5e.
//
// Solidity: function _backendSrvAddress() view returns(address)
func (_Labr *LabrCallerSession) BackendSrvAddress() (common.Address, error) {
	return _Labr.Contract.BackendSrvAddress(&_Labr.CallOpts)
}

// BalancerAddress is a free data retrieval call binding the contract method 0x8c682bbf.
//
// Solidity: function _balancerAddress() view returns(address)
func (_Labr *LabrCaller) BalancerAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Labr.contract.Call(opts, &out, "_balancerAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BalancerAddress is a free data retrieval call binding the contract method 0x8c682bbf.
//
// Solidity: function _balancerAddress() view returns(address)
func (_Labr *LabrSession) BalancerAddress() (common.Address, error) {
	return _Labr.Contract.BalancerAddress(&_Labr.CallOpts)
}

// BalancerAddress is a free data retrieval call binding the contract method 0x8c682bbf.
//
// Solidity: function _balancerAddress() view returns(address)
func (_Labr *LabrCallerSession) BalancerAddress() (common.Address, error) {
	return _Labr.Contract.BalancerAddress(&_Labr.CallOpts)
}

// ChainID is a free data retrieval call binding the contract method 0xbeab7131.
//
// Solidity: function _chainID() view returns(bytes8)
func (_Labr *LabrCaller) ChainID(opts *bind.CallOpts) ([8]byte, error) {
	var out []interface{}
	err := _Labr.contract.Call(opts, &out, "_chainID")

	if err != nil {
		return *new([8]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([8]byte)).(*[8]byte)

	return out0, err

}

// ChainID is a free data retrieval call binding the contract method 0xbeab7131.
//
// Solidity: function _chainID() view returns(bytes8)
func (_Labr *LabrSession) ChainID() ([8]byte, error) {
	return _Labr.Contract.ChainID(&_Labr.CallOpts)
}

// ChainID is a free data retrieval call binding the contract method 0xbeab7131.
//
// Solidity: function _chainID() view returns(bytes8)
func (_Labr *LabrCallerSession) ChainID() ([8]byte, error) {
	return _Labr.Contract.ChainID(&_Labr.CallOpts)
}

// DepositCounts is a free data retrieval call binding the contract method 0x174e300b.
//
// Solidity: function _depositCounts(bytes8 ) view returns(uint64)
func (_Labr *LabrCaller) DepositCounts(opts *bind.CallOpts, arg0 [8]byte) (uint64, error) {
	var out []interface{}
	err := _Labr.contract.Call(opts, &out, "_depositCounts", arg0)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// DepositCounts is a free data retrieval call binding the contract method 0x174e300b.
//
// Solidity: function _depositCounts(bytes8 ) view returns(uint64)
func (_Labr *LabrSession) DepositCounts(arg0 [8]byte) (uint64, error) {
	return _Labr.Contract.DepositCounts(&_Labr.CallOpts, arg0)
}

// DepositCounts is a free data retrieval call binding the contract method 0x174e300b.
//
// Solidity: function _depositCounts(bytes8 ) view returns(uint64)
func (_Labr *LabrCallerSession) DepositCounts(arg0 [8]byte) (uint64, error) {
	return _Labr.Contract.DepositCounts(&_Labr.CallOpts, arg0)
}

// DepositRecords is a free data retrieval call binding the contract method 0xa48ace38.
//
// Solidity: function _depositRecords(uint64 , bytes8 ) view returns(bytes)
func (_Labr *LabrCaller) DepositRecords(opts *bind.CallOpts, arg0 uint64, arg1 [8]byte) ([]byte, error) {
	var out []interface{}
	err := _Labr.contract.Call(opts, &out, "_depositRecords", arg0, arg1)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// DepositRecords is a free data retrieval call binding the contract method 0xa48ace38.
//
// Solidity: function _depositRecords(uint64 , bytes8 ) view returns(bytes)
func (_Labr *LabrSession) DepositRecords(arg0 uint64, arg1 [8]byte) ([]byte, error) {
	return _Labr.Contract.DepositRecords(&_Labr.CallOpts, arg0, arg1)
}

// DepositRecords is a free data retrieval call binding the contract method 0xa48ace38.
//
// Solidity: function _depositRecords(uint64 , bytes8 ) view returns(bytes)
func (_Labr *LabrCallerSession) DepositRecords(arg0 uint64, arg1 [8]byte) ([]byte, error) {
	return _Labr.Contract.DepositRecords(&_Labr.CallOpts, arg0, arg1)
}

// Expiry is a free data retrieval call binding the contract method 0xc5ec8970.
//
// Solidity: function _expiry() view returns(uint256)
func (_Labr *LabrCaller) Expiry(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Labr.contract.Call(opts, &out, "_expiry")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Expiry is a free data retrieval call binding the contract method 0xc5ec8970.
//
// Solidity: function _expiry() view returns(uint256)
func (_Labr *LabrSession) Expiry() (*big.Int, error) {
	return _Labr.Contract.Expiry(&_Labr.CallOpts)
}

// Expiry is a free data retrieval call binding the contract method 0xc5ec8970.
//
// Solidity: function _expiry() view returns(uint256)
func (_Labr *LabrCallerSession) Expiry() (*big.Int, error) {
	return _Labr.Contract.Expiry(&_Labr.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xc5b37c22.
//
// Solidity: function _fee() view returns(uint256)
func (_Labr *LabrCaller) Fee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Labr.contract.Call(opts, &out, "_fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fee is a free data retrieval call binding the contract method 0xc5b37c22.
//
// Solidity: function _fee() view returns(uint256)
func (_Labr *LabrSession) Fee() (*big.Int, error) {
	return _Labr.Contract.Fee(&_Labr.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xc5b37c22.
//
// Solidity: function _fee() view returns(uint256)
func (_Labr *LabrCallerSession) Fee() (*big.Int, error) {
	return _Labr.Contract.Fee(&_Labr.CallOpts)
}

// HasVotedOnProposal is a free data retrieval call binding the contract method 0x86c99f0e.
//
// Solidity: function _hasVotedOnProposal(bytes32 , bytes32 , address ) view returns(bool)
func (_Labr *LabrCaller) HasVotedOnProposal(opts *bind.CallOpts, arg0 [32]byte, arg1 [32]byte, arg2 common.Address) (bool, error) {
	var out []interface{}
	err := _Labr.contract.Call(opts, &out, "_hasVotedOnProposal", arg0, arg1, arg2)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasVotedOnProposal is a free data retrieval call binding the contract method 0x86c99f0e.
//
// Solidity: function _hasVotedOnProposal(bytes32 , bytes32 , address ) view returns(bool)
func (_Labr *LabrSession) HasVotedOnProposal(arg0 [32]byte, arg1 [32]byte, arg2 common.Address) (bool, error) {
	return _Labr.Contract.HasVotedOnProposal(&_Labr.CallOpts, arg0, arg1, arg2)
}

// HasVotedOnProposal is a free data retrieval call binding the contract method 0x86c99f0e.
//
// Solidity: function _hasVotedOnProposal(bytes32 , bytes32 , address ) view returns(bool)
func (_Labr *LabrCallerSession) HasVotedOnProposal(arg0 [32]byte, arg1 [32]byte, arg2 common.Address) (bool, error) {
	return _Labr.Contract.HasVotedOnProposal(&_Labr.CallOpts, arg0, arg1, arg2)
}

// IsInitialised is a free data retrieval call binding the contract method 0xdd2e8ec3.
//
// Solidity: function _isInitialised() view returns(bool)
func (_Labr *LabrCaller) IsInitialised(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Labr.contract.Call(opts, &out, "_isInitialised")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsInitialised is a free data retrieval call binding the contract method 0xdd2e8ec3.
//
// Solidity: function _isInitialised() view returns(bool)
func (_Labr *LabrSession) IsInitialised() (bool, error) {
	return _Labr.Contract.IsInitialised(&_Labr.CallOpts)
}

// IsInitialised is a free data retrieval call binding the contract method 0xdd2e8ec3.
//
// Solidity: function _isInitialised() view returns(bool)
func (_Labr *LabrCallerSession) IsInitialised() (bool, error) {
	return _Labr.Contract.IsInitialised(&_Labr.CallOpts)
}

// NativeResourceID is a free data retrieval call binding the contract method 0xa21b952e.
//
// Solidity: function _nativeResourceID() view returns(bytes32)
func (_Labr *LabrCaller) NativeResourceID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Labr.contract.Call(opts, &out, "_nativeResourceID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// NativeResourceID is a free data retrieval call binding the contract method 0xa21b952e.
//
// Solidity: function _nativeResourceID() view returns(bytes32)
func (_Labr *LabrSession) NativeResourceID() ([32]byte, error) {
	return _Labr.Contract.NativeResourceID(&_Labr.CallOpts)
}

// NativeResourceID is a free data retrieval call binding the contract method 0xa21b952e.
//
// Solidity: function _nativeResourceID() view returns(bytes32)
func (_Labr *LabrCallerSession) NativeResourceID() ([32]byte, error) {
	return _Labr.Contract.NativeResourceID(&_Labr.CallOpts)
}

// Proposals is a free data retrieval call binding the contract method 0x7bda741c.
//
// Solidity: function _proposals(bytes32 , bytes32 ) view returns((bytes32,bytes32,address[],address[],uint8,uint256))
func (_Labr *LabrCaller) Proposals(opts *bind.CallOpts, arg0 [32]byte, arg1 [32]byte) (BridgeProposal, error) {
	var out []interface{}
	err := _Labr.contract.Call(opts, &out, "_proposals", arg0, arg1)

	if err != nil {
		return *new(BridgeProposal), err
	}

	out0 := *abi.ConvertType(out[0], new(BridgeProposal)).(*BridgeProposal)

	return out0, err

}

// Proposals is a free data retrieval call binding the contract method 0x7bda741c.
//
// Solidity: function _proposals(bytes32 , bytes32 ) view returns((bytes32,bytes32,address[],address[],uint8,uint256))
func (_Labr *LabrSession) Proposals(arg0 [32]byte, arg1 [32]byte) (BridgeProposal, error) {
	return _Labr.Contract.Proposals(&_Labr.CallOpts, arg0, arg1)
}

// Proposals is a free data retrieval call binding the contract method 0x7bda741c.
//
// Solidity: function _proposals(bytes32 , bytes32 ) view returns((bytes32,bytes32,address[],address[],uint8,uint256))
func (_Labr *LabrCallerSession) Proposals(arg0 [32]byte, arg1 [32]byte) (BridgeProposal, error) {
	return _Labr.Contract.Proposals(&_Labr.CallOpts, arg0, arg1)
}

// RelayerHubAddress is a free data retrieval call binding the contract method 0xb6d97623.
//
// Solidity: function _relayerHubAddress() view returns(address)
func (_Labr *LabrCaller) RelayerHubAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Labr.contract.Call(opts, &out, "_relayerHubAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RelayerHubAddress is a free data retrieval call binding the contract method 0xb6d97623.
//
// Solidity: function _relayerHubAddress() view returns(address)
func (_Labr *LabrSession) RelayerHubAddress() (common.Address, error) {
	return _Labr.Contract.RelayerHubAddress(&_Labr.CallOpts)
}

// RelayerHubAddress is a free data retrieval call binding the contract method 0xb6d97623.
//
// Solidity: function _relayerHubAddress() view returns(address)
func (_Labr *LabrCallerSession) RelayerHubAddress() (common.Address, error) {
	return _Labr.Contract.RelayerHubAddress(&_Labr.CallOpts)
}

// RelayerThreshold is a free data retrieval call binding the contract method 0xd7a9cd79.
//
// Solidity: function _relayerThreshold() view returns(uint256)
func (_Labr *LabrCaller) RelayerThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Labr.contract.Call(opts, &out, "_relayerThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RelayerThreshold is a free data retrieval call binding the contract method 0xd7a9cd79.
//
// Solidity: function _relayerThreshold() view returns(uint256)
func (_Labr *LabrSession) RelayerThreshold() (*big.Int, error) {
	return _Labr.Contract.RelayerThreshold(&_Labr.CallOpts)
}

// RelayerThreshold is a free data retrieval call binding the contract method 0xd7a9cd79.
//
// Solidity: function _relayerThreshold() view returns(uint256)
func (_Labr *LabrCallerSession) RelayerThreshold() (*big.Int, error) {
	return _Labr.Contract.RelayerThreshold(&_Labr.CallOpts)
}

// ResourceIDToHandlerAddress is a free data retrieval call binding the contract method 0x84db809f.
//
// Solidity: function _resourceIDToHandlerAddress(bytes32 ) view returns(address)
func (_Labr *LabrCaller) ResourceIDToHandlerAddress(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _Labr.contract.Call(opts, &out, "_resourceIDToHandlerAddress", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ResourceIDToHandlerAddress is a free data retrieval call binding the contract method 0x84db809f.
//
// Solidity: function _resourceIDToHandlerAddress(bytes32 ) view returns(address)
func (_Labr *LabrSession) ResourceIDToHandlerAddress(arg0 [32]byte) (common.Address, error) {
	return _Labr.Contract.ResourceIDToHandlerAddress(&_Labr.CallOpts, arg0)
}

// ResourceIDToHandlerAddress is a free data retrieval call binding the contract method 0x84db809f.
//
// Solidity: function _resourceIDToHandlerAddress(bytes32 ) view returns(address)
func (_Labr *LabrCallerSession) ResourceIDToHandlerAddress(arg0 [32]byte) (common.Address, error) {
	return _Labr.Contract.ResourceIDToHandlerAddress(&_Labr.CallOpts, arg0)
}

// TotalProposals is a free data retrieval call binding the contract method 0x9d5773e0.
//
// Solidity: function _totalProposals() view returns(uint256)
func (_Labr *LabrCaller) TotalProposals(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Labr.contract.Call(opts, &out, "_totalProposals")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalProposals is a free data retrieval call binding the contract method 0x9d5773e0.
//
// Solidity: function _totalProposals() view returns(uint256)
func (_Labr *LabrSession) TotalProposals() (*big.Int, error) {
	return _Labr.Contract.TotalProposals(&_Labr.CallOpts)
}

// TotalProposals is a free data retrieval call binding the contract method 0x9d5773e0.
//
// Solidity: function _totalProposals() view returns(uint256)
func (_Labr *LabrCallerSession) TotalProposals() (*big.Int, error) {
	return _Labr.Contract.TotalProposals(&_Labr.CallOpts)
}

// CreateMesssageHash is a free data retrieval call binding the contract method 0xc79564f1.
//
// Solidity: function createMesssageHash(uint256 amount, address recipient, bytes8 chainId) pure returns(bytes32)
func (_Labr *LabrCaller) CreateMesssageHash(opts *bind.CallOpts, amount *big.Int, recipient common.Address, chainId [8]byte) ([32]byte, error) {
	var out []interface{}
	err := _Labr.contract.Call(opts, &out, "createMesssageHash", amount, recipient, chainId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CreateMesssageHash is a free data retrieval call binding the contract method 0xc79564f1.
//
// Solidity: function createMesssageHash(uint256 amount, address recipient, bytes8 chainId) pure returns(bytes32)
func (_Labr *LabrSession) CreateMesssageHash(amount *big.Int, recipient common.Address, chainId [8]byte) ([32]byte, error) {
	return _Labr.Contract.CreateMesssageHash(&_Labr.CallOpts, amount, recipient, chainId)
}

// CreateMesssageHash is a free data retrieval call binding the contract method 0xc79564f1.
//
// Solidity: function createMesssageHash(uint256 amount, address recipient, bytes8 chainId) pure returns(bytes32)
func (_Labr *LabrCallerSession) CreateMesssageHash(amount *big.Int, recipient common.Address, chainId [8]byte) ([32]byte, error) {
	return _Labr.Contract.CreateMesssageHash(&_Labr.CallOpts, amount, recipient, chainId)
}

// ExtraLATransferred is a free data retrieval call binding the contract method 0x51fc8c58.
//
// Solidity: function extraLATransferred(bytes32 , bytes32 ) view returns(bool)
func (_Labr *LabrCaller) ExtraLATransferred(opts *bind.CallOpts, arg0 [32]byte, arg1 [32]byte) (bool, error) {
	var out []interface{}
	err := _Labr.contract.Call(opts, &out, "extraLATransferred", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ExtraLATransferred is a free data retrieval call binding the contract method 0x51fc8c58.
//
// Solidity: function extraLATransferred(bytes32 , bytes32 ) view returns(bool)
func (_Labr *LabrSession) ExtraLATransferred(arg0 [32]byte, arg1 [32]byte) (bool, error) {
	return _Labr.Contract.ExtraLATransferred(&_Labr.CallOpts, arg0, arg1)
}

// ExtraLATransferred is a free data retrieval call binding the contract method 0x51fc8c58.
//
// Solidity: function extraLATransferred(bytes32 , bytes32 ) view returns(bool)
func (_Labr *LabrCallerSession) ExtraLATransferred(arg0 [32]byte, arg1 [32]byte) (bool, error) {
	return _Labr.Contract.ExtraLATransferred(&_Labr.CallOpts, arg0, arg1)
}

// GetEthSignedMessageHash is a free data retrieval call binding the contract method 0xfa540801.
//
// Solidity: function getEthSignedMessageHash(bytes32 _messageHash) pure returns(bytes32)
func (_Labr *LabrCaller) GetEthSignedMessageHash(opts *bind.CallOpts, _messageHash [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Labr.contract.Call(opts, &out, "getEthSignedMessageHash", _messageHash)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetEthSignedMessageHash is a free data retrieval call binding the contract method 0xfa540801.
//
// Solidity: function getEthSignedMessageHash(bytes32 _messageHash) pure returns(bytes32)
func (_Labr *LabrSession) GetEthSignedMessageHash(_messageHash [32]byte) ([32]byte, error) {
	return _Labr.Contract.GetEthSignedMessageHash(&_Labr.CallOpts, _messageHash)
}

// GetEthSignedMessageHash is a free data retrieval call binding the contract method 0xfa540801.
//
// Solidity: function getEthSignedMessageHash(bytes32 _messageHash) pure returns(bytes32)
func (_Labr *LabrCallerSession) GetEthSignedMessageHash(_messageHash [32]byte) ([32]byte, error) {
	return _Labr.Contract.GetEthSignedMessageHash(&_Labr.CallOpts, _messageHash)
}

// GetExtraLATransferred is a free data retrieval call binding the contract method 0xd4ae814c.
//
// Solidity: function getExtraLATransferred(bytes8 originChainID, bytes8 destinationChainID, uint64 depositNonce, bytes32 resourceID, address recipientAddress, uint256 amount) view returns(bool)
func (_Labr *LabrCaller) GetExtraLATransferred(opts *bind.CallOpts, originChainID [8]byte, destinationChainID [8]byte, depositNonce uint64, resourceID [32]byte, recipientAddress common.Address, amount *big.Int) (bool, error) {
	var out []interface{}
	err := _Labr.contract.Call(opts, &out, "getExtraLATransferred", originChainID, destinationChainID, depositNonce, resourceID, recipientAddress, amount)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetExtraLATransferred is a free data retrieval call binding the contract method 0xd4ae814c.
//
// Solidity: function getExtraLATransferred(bytes8 originChainID, bytes8 destinationChainID, uint64 depositNonce, bytes32 resourceID, address recipientAddress, uint256 amount) view returns(bool)
func (_Labr *LabrSession) GetExtraLATransferred(originChainID [8]byte, destinationChainID [8]byte, depositNonce uint64, resourceID [32]byte, recipientAddress common.Address, amount *big.Int) (bool, error) {
	return _Labr.Contract.GetExtraLATransferred(&_Labr.CallOpts, originChainID, destinationChainID, depositNonce, resourceID, recipientAddress, amount)
}

// GetExtraLATransferred is a free data retrieval call binding the contract method 0xd4ae814c.
//
// Solidity: function getExtraLATransferred(bytes8 originChainID, bytes8 destinationChainID, uint64 depositNonce, bytes32 resourceID, address recipientAddress, uint256 amount) view returns(bool)
func (_Labr *LabrCallerSession) GetExtraLATransferred(originChainID [8]byte, destinationChainID [8]byte, depositNonce uint64, resourceID [32]byte, recipientAddress common.Address, amount *big.Int) (bool, error) {
	return _Labr.Contract.GetExtraLATransferred(&_Labr.CallOpts, originChainID, destinationChainID, depositNonce, resourceID, recipientAddress, amount)
}

// GetProposal is a free data retrieval call binding the contract method 0xa241c0dd.
//
// Solidity: function getProposal(bytes8 originChainID, bytes8 destinationChainID, uint64 depositNonce, address recipientAddress, uint256 amount, bytes32 resourceID) view returns((bytes32,bytes32,address[],address[],uint8,uint256))
func (_Labr *LabrCaller) GetProposal(opts *bind.CallOpts, originChainID [8]byte, destinationChainID [8]byte, depositNonce uint64, recipientAddress common.Address, amount *big.Int, resourceID [32]byte) (BridgeProposal, error) {
	var out []interface{}
	err := _Labr.contract.Call(opts, &out, "getProposal", originChainID, destinationChainID, depositNonce, recipientAddress, amount, resourceID)

	if err != nil {
		return *new(BridgeProposal), err
	}

	out0 := *abi.ConvertType(out[0], new(BridgeProposal)).(*BridgeProposal)

	return out0, err

}

// GetProposal is a free data retrieval call binding the contract method 0xa241c0dd.
//
// Solidity: function getProposal(bytes8 originChainID, bytes8 destinationChainID, uint64 depositNonce, address recipientAddress, uint256 amount, bytes32 resourceID) view returns((bytes32,bytes32,address[],address[],uint8,uint256))
func (_Labr *LabrSession) GetProposal(originChainID [8]byte, destinationChainID [8]byte, depositNonce uint64, recipientAddress common.Address, amount *big.Int, resourceID [32]byte) (BridgeProposal, error) {
	return _Labr.Contract.GetProposal(&_Labr.CallOpts, originChainID, destinationChainID, depositNonce, recipientAddress, amount, resourceID)
}

// GetProposal is a free data retrieval call binding the contract method 0xa241c0dd.
//
// Solidity: function getProposal(bytes8 originChainID, bytes8 destinationChainID, uint64 depositNonce, address recipientAddress, uint256 amount, bytes32 resourceID) view returns((bytes32,bytes32,address[],address[],uint8,uint256))
func (_Labr *LabrCallerSession) GetProposal(originChainID [8]byte, destinationChainID [8]byte, depositNonce uint64, recipientAddress common.Address, amount *big.Int, resourceID [32]byte) (BridgeProposal, error) {
	return _Labr.Contract.GetProposal(&_Labr.CallOpts, originChainID, destinationChainID, depositNonce, recipientAddress, amount, resourceID)
}

// GetSigner is a free data retrieval call binding the contract method 0xf7b2ec0d.
//
// Solidity: function getSigner(bytes32 messageHash, bytes signature) pure returns(address)
func (_Labr *LabrCaller) GetSigner(opts *bind.CallOpts, messageHash [32]byte, signature []byte) (common.Address, error) {
	var out []interface{}
	err := _Labr.contract.Call(opts, &out, "getSigner", messageHash, signature)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetSigner is a free data retrieval call binding the contract method 0xf7b2ec0d.
//
// Solidity: function getSigner(bytes32 messageHash, bytes signature) pure returns(address)
func (_Labr *LabrSession) GetSigner(messageHash [32]byte, signature []byte) (common.Address, error) {
	return _Labr.Contract.GetSigner(&_Labr.CallOpts, messageHash, signature)
}

// GetSigner is a free data retrieval call binding the contract method 0xf7b2ec0d.
//
// Solidity: function getSigner(bytes32 messageHash, bytes signature) pure returns(address)
func (_Labr *LabrCallerSession) GetSigner(messageHash [32]byte, signature []byte) (common.Address, error) {
	return _Labr.Contract.GetSigner(&_Labr.CallOpts, messageHash, signature)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_Labr *LabrCaller) Implementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Labr.contract.Call(opts, &out, "implementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_Labr *LabrSession) Implementation() (common.Address, error) {
	return _Labr.Contract.Implementation(&_Labr.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_Labr *LabrCallerSession) Implementation() (common.Address, error) {
	return _Labr.Contract.Implementation(&_Labr.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Labr *LabrCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Labr.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Labr *LabrSession) Owner() (common.Address, error) {
	return _Labr.Contract.Owner(&_Labr.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Labr *LabrCallerSession) Owner() (common.Address, error) {
	return _Labr.Contract.Owner(&_Labr.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Labr *LabrCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Labr.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Labr *LabrSession) Paused() (bool, error) {
	return _Labr.Contract.Paused(&_Labr.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Labr *LabrCallerSession) Paused() (bool, error) {
	return _Labr.Contract.Paused(&_Labr.CallOpts)
}

// Sub is a free data retrieval call binding the contract method 0xb67d77c5.
//
// Solidity: function sub(uint256 a, uint256 b) pure returns(uint256)
func (_Labr *LabrCaller) Sub(opts *bind.CallOpts, a *big.Int, b *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Labr.contract.Call(opts, &out, "sub", a, b)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Sub is a free data retrieval call binding the contract method 0xb67d77c5.
//
// Solidity: function sub(uint256 a, uint256 b) pure returns(uint256)
func (_Labr *LabrSession) Sub(a *big.Int, b *big.Int) (*big.Int, error) {
	return _Labr.Contract.Sub(&_Labr.CallOpts, a, b)
}

// Sub is a free data retrieval call binding the contract method 0xb67d77c5.
//
// Solidity: function sub(uint256 a, uint256 b) pure returns(uint256)
func (_Labr *LabrCallerSession) Sub(a *big.Int, b *big.Int) (*big.Int, error) {
	return _Labr.Contract.Sub(&_Labr.CallOpts, a, b)
}

// Verify is a free data retrieval call binding the contract method 0xf5cf2ce1.
//
// Solidity: function verify(bytes32 _ethSignedMessageHash, address _signer, bytes _signature) pure returns(bool)
func (_Labr *LabrCaller) Verify(opts *bind.CallOpts, _ethSignedMessageHash [32]byte, _signer common.Address, _signature []byte) (bool, error) {
	var out []interface{}
	err := _Labr.contract.Call(opts, &out, "verify", _ethSignedMessageHash, _signer, _signature)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Verify is a free data retrieval call binding the contract method 0xf5cf2ce1.
//
// Solidity: function verify(bytes32 _ethSignedMessageHash, address _signer, bytes _signature) pure returns(bool)
func (_Labr *LabrSession) Verify(_ethSignedMessageHash [32]byte, _signer common.Address, _signature []byte) (bool, error) {
	return _Labr.Contract.Verify(&_Labr.CallOpts, _ethSignedMessageHash, _signer, _signature)
}

// Verify is a free data retrieval call binding the contract method 0xf5cf2ce1.
//
// Solidity: function verify(bytes32 _ethSignedMessageHash, address _signer, bytes _signature) pure returns(bool)
func (_Labr *LabrCallerSession) Verify(_ethSignedMessageHash [32]byte, _signer common.Address, _signature []byte) (bool, error) {
	return _Labr.Contract.Verify(&_Labr.CallOpts, _ethSignedMessageHash, _signer, _signature)
}

// AdminApprove is a paid mutator transaction binding the contract method 0x8e0747be.
//
// Solidity: function adminApprove(bytes32 resourceIDOwner, bytes32 resourceIDSpender, uint256 amountOrTokenID) returns()
func (_Labr *LabrTransactor) AdminApprove(opts *bind.TransactOpts, resourceIDOwner [32]byte, resourceIDSpender [32]byte, amountOrTokenID *big.Int) (*types.Transaction, error) {
	return _Labr.contract.Transact(opts, "adminApprove", resourceIDOwner, resourceIDSpender, amountOrTokenID)
}

// AdminApprove is a paid mutator transaction binding the contract method 0x8e0747be.
//
// Solidity: function adminApprove(bytes32 resourceIDOwner, bytes32 resourceIDSpender, uint256 amountOrTokenID) returns()
func (_Labr *LabrSession) AdminApprove(resourceIDOwner [32]byte, resourceIDSpender [32]byte, amountOrTokenID *big.Int) (*types.Transaction, error) {
	return _Labr.Contract.AdminApprove(&_Labr.TransactOpts, resourceIDOwner, resourceIDSpender, amountOrTokenID)
}

// AdminApprove is a paid mutator transaction binding the contract method 0x8e0747be.
//
// Solidity: function adminApprove(bytes32 resourceIDOwner, bytes32 resourceIDSpender, uint256 amountOrTokenID) returns()
func (_Labr *LabrTransactorSession) AdminApprove(resourceIDOwner [32]byte, resourceIDSpender [32]byte, amountOrTokenID *big.Int) (*types.Transaction, error) {
	return _Labr.Contract.AdminApprove(&_Labr.TransactOpts, resourceIDOwner, resourceIDSpender, amountOrTokenID)
}

// AdminChangeFee is a paid mutator transaction binding the contract method 0x91c404ac.
//
// Solidity: function adminChangeFee(uint256 newFee) returns()
func (_Labr *LabrTransactor) AdminChangeFee(opts *bind.TransactOpts, newFee *big.Int) (*types.Transaction, error) {
	return _Labr.contract.Transact(opts, "adminChangeFee", newFee)
}

// AdminChangeFee is a paid mutator transaction binding the contract method 0x91c404ac.
//
// Solidity: function adminChangeFee(uint256 newFee) returns()
func (_Labr *LabrSession) AdminChangeFee(newFee *big.Int) (*types.Transaction, error) {
	return _Labr.Contract.AdminChangeFee(&_Labr.TransactOpts, newFee)
}

// AdminChangeFee is a paid mutator transaction binding the contract method 0x91c404ac.
//
// Solidity: function adminChangeFee(uint256 newFee) returns()
func (_Labr *LabrTransactorSession) AdminChangeFee(newFee *big.Int) (*types.Transaction, error) {
	return _Labr.Contract.AdminChangeFee(&_Labr.TransactOpts, newFee)
}

// AdminChangeRelayerThreshold is a paid mutator transaction binding the contract method 0x4e056005.
//
// Solidity: function adminChangeRelayerThreshold(uint256 newThreshold) returns()
func (_Labr *LabrTransactor) AdminChangeRelayerThreshold(opts *bind.TransactOpts, newThreshold *big.Int) (*types.Transaction, error) {
	return _Labr.contract.Transact(opts, "adminChangeRelayerThreshold", newThreshold)
}

// AdminChangeRelayerThreshold is a paid mutator transaction binding the contract method 0x4e056005.
//
// Solidity: function adminChangeRelayerThreshold(uint256 newThreshold) returns()
func (_Labr *LabrSession) AdminChangeRelayerThreshold(newThreshold *big.Int) (*types.Transaction, error) {
	return _Labr.Contract.AdminChangeRelayerThreshold(&_Labr.TransactOpts, newThreshold)
}

// AdminChangeRelayerThreshold is a paid mutator transaction binding the contract method 0x4e056005.
//
// Solidity: function adminChangeRelayerThreshold(uint256 newThreshold) returns()
func (_Labr *LabrTransactorSession) AdminChangeRelayerThreshold(newThreshold *big.Int) (*types.Transaction, error) {
	return _Labr.Contract.AdminChangeRelayerThreshold(&_Labr.TransactOpts, newThreshold)
}

// AdminCollectFees is a paid mutator transaction binding the contract method 0x7a7eed77.
//
// Solidity: function adminCollectFees(address recipient, uint256 amount) returns()
func (_Labr *LabrTransactor) AdminCollectFees(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Labr.contract.Transact(opts, "adminCollectFees", recipient, amount)
}

// AdminCollectFees is a paid mutator transaction binding the contract method 0x7a7eed77.
//
// Solidity: function adminCollectFees(address recipient, uint256 amount) returns()
func (_Labr *LabrSession) AdminCollectFees(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Labr.Contract.AdminCollectFees(&_Labr.TransactOpts, recipient, amount)
}

// AdminCollectFees is a paid mutator transaction binding the contract method 0x7a7eed77.
//
// Solidity: function adminCollectFees(address recipient, uint256 amount) returns()
func (_Labr *LabrTransactorSession) AdminCollectFees(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Labr.Contract.AdminCollectFees(&_Labr.TransactOpts, recipient, amount)
}

// AdminPauseTransfers is a paid mutator transaction binding the contract method 0x80ae1c28.
//
// Solidity: function adminPauseTransfers() returns()
func (_Labr *LabrTransactor) AdminPauseTransfers(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Labr.contract.Transact(opts, "adminPauseTransfers")
}

// AdminPauseTransfers is a paid mutator transaction binding the contract method 0x80ae1c28.
//
// Solidity: function adminPauseTransfers() returns()
func (_Labr *LabrSession) AdminPauseTransfers() (*types.Transaction, error) {
	return _Labr.Contract.AdminPauseTransfers(&_Labr.TransactOpts)
}

// AdminPauseTransfers is a paid mutator transaction binding the contract method 0x80ae1c28.
//
// Solidity: function adminPauseTransfers() returns()
func (_Labr *LabrTransactorSession) AdminPauseTransfers() (*types.Transaction, error) {
	return _Labr.Contract.AdminPauseTransfers(&_Labr.TransactOpts)
}

// AdminSetBackendSrvAddress is a paid mutator transaction binding the contract method 0x23328f39.
//
// Solidity: function adminSetBackendSrvAddress(address newBackendSrv) returns()
func (_Labr *LabrTransactor) AdminSetBackendSrvAddress(opts *bind.TransactOpts, newBackendSrv common.Address) (*types.Transaction, error) {
	return _Labr.contract.Transact(opts, "adminSetBackendSrvAddress", newBackendSrv)
}

// AdminSetBackendSrvAddress is a paid mutator transaction binding the contract method 0x23328f39.
//
// Solidity: function adminSetBackendSrvAddress(address newBackendSrv) returns()
func (_Labr *LabrSession) AdminSetBackendSrvAddress(newBackendSrv common.Address) (*types.Transaction, error) {
	return _Labr.Contract.AdminSetBackendSrvAddress(&_Labr.TransactOpts, newBackendSrv)
}

// AdminSetBackendSrvAddress is a paid mutator transaction binding the contract method 0x23328f39.
//
// Solidity: function adminSetBackendSrvAddress(address newBackendSrv) returns()
func (_Labr *LabrTransactorSession) AdminSetBackendSrvAddress(newBackendSrv common.Address) (*types.Transaction, error) {
	return _Labr.Contract.AdminSetBackendSrvAddress(&_Labr.TransactOpts, newBackendSrv)
}

// AdminSetBalancerAddress is a paid mutator transaction binding the contract method 0x14ae4a39.
//
// Solidity: function adminSetBalancerAddress(address newBalancer) returns()
func (_Labr *LabrTransactor) AdminSetBalancerAddress(opts *bind.TransactOpts, newBalancer common.Address) (*types.Transaction, error) {
	return _Labr.contract.Transact(opts, "adminSetBalancerAddress", newBalancer)
}

// AdminSetBalancerAddress is a paid mutator transaction binding the contract method 0x14ae4a39.
//
// Solidity: function adminSetBalancerAddress(address newBalancer) returns()
func (_Labr *LabrSession) AdminSetBalancerAddress(newBalancer common.Address) (*types.Transaction, error) {
	return _Labr.Contract.AdminSetBalancerAddress(&_Labr.TransactOpts, newBalancer)
}

// AdminSetBalancerAddress is a paid mutator transaction binding the contract method 0x14ae4a39.
//
// Solidity: function adminSetBalancerAddress(address newBalancer) returns()
func (_Labr *LabrTransactorSession) AdminSetBalancerAddress(newBalancer common.Address) (*types.Transaction, error) {
	return _Labr.Contract.AdminSetBalancerAddress(&_Labr.TransactOpts, newBalancer)
}

// AdminSetBurnable is a paid mutator transaction binding the contract method 0x8c0c2631.
//
// Solidity: function adminSetBurnable(address handlerAddress, address tokenAddress) returns()
func (_Labr *LabrTransactor) AdminSetBurnable(opts *bind.TransactOpts, handlerAddress common.Address, tokenAddress common.Address) (*types.Transaction, error) {
	return _Labr.contract.Transact(opts, "adminSetBurnable", handlerAddress, tokenAddress)
}

// AdminSetBurnable is a paid mutator transaction binding the contract method 0x8c0c2631.
//
// Solidity: function adminSetBurnable(address handlerAddress, address tokenAddress) returns()
func (_Labr *LabrSession) AdminSetBurnable(handlerAddress common.Address, tokenAddress common.Address) (*types.Transaction, error) {
	return _Labr.Contract.AdminSetBurnable(&_Labr.TransactOpts, handlerAddress, tokenAddress)
}

// AdminSetBurnable is a paid mutator transaction binding the contract method 0x8c0c2631.
//
// Solidity: function adminSetBurnable(address handlerAddress, address tokenAddress) returns()
func (_Labr *LabrTransactorSession) AdminSetBurnable(handlerAddress common.Address, tokenAddress common.Address) (*types.Transaction, error) {
	return _Labr.Contract.AdminSetBurnable(&_Labr.TransactOpts, handlerAddress, tokenAddress)
}

// AdminSetNativeResourceID is a paid mutator transaction binding the contract method 0xa2b5b654.
//
// Solidity: function adminSetNativeResourceID(bytes32 resourceID) returns()
func (_Labr *LabrTransactor) AdminSetNativeResourceID(opts *bind.TransactOpts, resourceID [32]byte) (*types.Transaction, error) {
	return _Labr.contract.Transact(opts, "adminSetNativeResourceID", resourceID)
}

// AdminSetNativeResourceID is a paid mutator transaction binding the contract method 0xa2b5b654.
//
// Solidity: function adminSetNativeResourceID(bytes32 resourceID) returns()
func (_Labr *LabrSession) AdminSetNativeResourceID(resourceID [32]byte) (*types.Transaction, error) {
	return _Labr.Contract.AdminSetNativeResourceID(&_Labr.TransactOpts, resourceID)
}

// AdminSetNativeResourceID is a paid mutator transaction binding the contract method 0xa2b5b654.
//
// Solidity: function adminSetNativeResourceID(bytes32 resourceID) returns()
func (_Labr *LabrTransactorSession) AdminSetNativeResourceID(resourceID [32]byte) (*types.Transaction, error) {
	return _Labr.Contract.AdminSetNativeResourceID(&_Labr.TransactOpts, resourceID)
}

// AdminSetRelayerHub is a paid mutator transaction binding the contract method 0x5d699c5d.
//
// Solidity: function adminSetRelayerHub(address newRelayerHub) returns()
func (_Labr *LabrTransactor) AdminSetRelayerHub(opts *bind.TransactOpts, newRelayerHub common.Address) (*types.Transaction, error) {
	return _Labr.contract.Transact(opts, "adminSetRelayerHub", newRelayerHub)
}

// AdminSetRelayerHub is a paid mutator transaction binding the contract method 0x5d699c5d.
//
// Solidity: function adminSetRelayerHub(address newRelayerHub) returns()
func (_Labr *LabrSession) AdminSetRelayerHub(newRelayerHub common.Address) (*types.Transaction, error) {
	return _Labr.Contract.AdminSetRelayerHub(&_Labr.TransactOpts, newRelayerHub)
}

// AdminSetRelayerHub is a paid mutator transaction binding the contract method 0x5d699c5d.
//
// Solidity: function adminSetRelayerHub(address newRelayerHub) returns()
func (_Labr *LabrTransactorSession) AdminSetRelayerHub(newRelayerHub common.Address) (*types.Transaction, error) {
	return _Labr.Contract.AdminSetRelayerHub(&_Labr.TransactOpts, newRelayerHub)
}

// AdminSetResource is a paid mutator transaction binding the contract method 0xcb10f215.
//
// Solidity: function adminSetResource(address handlerAddress, bytes32 resourceID, address tokenAddress) returns()
func (_Labr *LabrTransactor) AdminSetResource(opts *bind.TransactOpts, handlerAddress common.Address, resourceID [32]byte, tokenAddress common.Address) (*types.Transaction, error) {
	return _Labr.contract.Transact(opts, "adminSetResource", handlerAddress, resourceID, tokenAddress)
}

// AdminSetResource is a paid mutator transaction binding the contract method 0xcb10f215.
//
// Solidity: function adminSetResource(address handlerAddress, bytes32 resourceID, address tokenAddress) returns()
func (_Labr *LabrSession) AdminSetResource(handlerAddress common.Address, resourceID [32]byte, tokenAddress common.Address) (*types.Transaction, error) {
	return _Labr.Contract.AdminSetResource(&_Labr.TransactOpts, handlerAddress, resourceID, tokenAddress)
}

// AdminSetResource is a paid mutator transaction binding the contract method 0xcb10f215.
//
// Solidity: function adminSetResource(address handlerAddress, bytes32 resourceID, address tokenAddress) returns()
func (_Labr *LabrTransactorSession) AdminSetResource(handlerAddress common.Address, resourceID [32]byte, tokenAddress common.Address) (*types.Transaction, error) {
	return _Labr.Contract.AdminSetResource(&_Labr.TransactOpts, handlerAddress, resourceID, tokenAddress)
}

// AdminUnpauseTransfers is a paid mutator transaction binding the contract method 0xffaac0eb.
//
// Solidity: function adminUnpauseTransfers() returns()
func (_Labr *LabrTransactor) AdminUnpauseTransfers(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Labr.contract.Transact(opts, "adminUnpauseTransfers")
}

// AdminUnpauseTransfers is a paid mutator transaction binding the contract method 0xffaac0eb.
//
// Solidity: function adminUnpauseTransfers() returns()
func (_Labr *LabrSession) AdminUnpauseTransfers() (*types.Transaction, error) {
	return _Labr.Contract.AdminUnpauseTransfers(&_Labr.TransactOpts)
}

// AdminUnpauseTransfers is a paid mutator transaction binding the contract method 0xffaac0eb.
//
// Solidity: function adminUnpauseTransfers() returns()
func (_Labr *LabrTransactorSession) AdminUnpauseTransfers() (*types.Transaction, error) {
	return _Labr.Contract.AdminUnpauseTransfers(&_Labr.TransactOpts)
}

// AdminWithdraw is a paid mutator transaction binding the contract method 0x780cf004.
//
// Solidity: function adminWithdraw(address handlerAddress, address tokenAddress, address recipient, uint256 amountOrTokenID) returns()
func (_Labr *LabrTransactor) AdminWithdraw(opts *bind.TransactOpts, handlerAddress common.Address, tokenAddress common.Address, recipient common.Address, amountOrTokenID *big.Int) (*types.Transaction, error) {
	return _Labr.contract.Transact(opts, "adminWithdraw", handlerAddress, tokenAddress, recipient, amountOrTokenID)
}

// AdminWithdraw is a paid mutator transaction binding the contract method 0x780cf004.
//
// Solidity: function adminWithdraw(address handlerAddress, address tokenAddress, address recipient, uint256 amountOrTokenID) returns()
func (_Labr *LabrSession) AdminWithdraw(handlerAddress common.Address, tokenAddress common.Address, recipient common.Address, amountOrTokenID *big.Int) (*types.Transaction, error) {
	return _Labr.Contract.AdminWithdraw(&_Labr.TransactOpts, handlerAddress, tokenAddress, recipient, amountOrTokenID)
}

// AdminWithdraw is a paid mutator transaction binding the contract method 0x780cf004.
//
// Solidity: function adminWithdraw(address handlerAddress, address tokenAddress, address recipient, uint256 amountOrTokenID) returns()
func (_Labr *LabrTransactorSession) AdminWithdraw(handlerAddress common.Address, tokenAddress common.Address, recipient common.Address, amountOrTokenID *big.Int) (*types.Transaction, error) {
	return _Labr.Contract.AdminWithdraw(&_Labr.TransactOpts, handlerAddress, tokenAddress, recipient, amountOrTokenID)
}

// CancelProposal is a paid mutator transaction binding the contract method 0x8d39b457.
//
// Solidity: function cancelProposal(bytes8 originChainID, bytes8 destinationChainID, uint64 depositNonce, address recipientAddress, uint256 amount, bytes32 resourceID) returns()
func (_Labr *LabrTransactor) CancelProposal(opts *bind.TransactOpts, originChainID [8]byte, destinationChainID [8]byte, depositNonce uint64, recipientAddress common.Address, amount *big.Int, resourceID [32]byte) (*types.Transaction, error) {
	return _Labr.contract.Transact(opts, "cancelProposal", originChainID, destinationChainID, depositNonce, recipientAddress, amount, resourceID)
}

// CancelProposal is a paid mutator transaction binding the contract method 0x8d39b457.
//
// Solidity: function cancelProposal(bytes8 originChainID, bytes8 destinationChainID, uint64 depositNonce, address recipientAddress, uint256 amount, bytes32 resourceID) returns()
func (_Labr *LabrSession) CancelProposal(originChainID [8]byte, destinationChainID [8]byte, depositNonce uint64, recipientAddress common.Address, amount *big.Int, resourceID [32]byte) (*types.Transaction, error) {
	return _Labr.Contract.CancelProposal(&_Labr.TransactOpts, originChainID, destinationChainID, depositNonce, recipientAddress, amount, resourceID)
}

// CancelProposal is a paid mutator transaction binding the contract method 0x8d39b457.
//
// Solidity: function cancelProposal(bytes8 originChainID, bytes8 destinationChainID, uint64 depositNonce, address recipientAddress, uint256 amount, bytes32 resourceID) returns()
func (_Labr *LabrTransactorSession) CancelProposal(originChainID [8]byte, destinationChainID [8]byte, depositNonce uint64, recipientAddress common.Address, amount *big.Int, resourceID [32]byte) (*types.Transaction, error) {
	return _Labr.Contract.CancelProposal(&_Labr.TransactOpts, originChainID, destinationChainID, depositNonce, recipientAddress, amount, resourceID)
}

// Deposit is a paid mutator transaction binding the contract method 0x98ce7a18.
//
// Solidity: function deposit(bytes8 destinationChainID, bytes32 resourceID, uint256 amount, address recipientAddress, bytes signature, bytes params) payable returns()
func (_Labr *LabrTransactor) Deposit(opts *bind.TransactOpts, destinationChainID [8]byte, resourceID [32]byte, amount *big.Int, recipientAddress common.Address, signature []byte, params []byte) (*types.Transaction, error) {
	return _Labr.contract.Transact(opts, "deposit", destinationChainID, resourceID, amount, recipientAddress, signature, params)
}

// Deposit is a paid mutator transaction binding the contract method 0x98ce7a18.
//
// Solidity: function deposit(bytes8 destinationChainID, bytes32 resourceID, uint256 amount, address recipientAddress, bytes signature, bytes params) payable returns()
func (_Labr *LabrSession) Deposit(destinationChainID [8]byte, resourceID [32]byte, amount *big.Int, recipientAddress common.Address, signature []byte, params []byte) (*types.Transaction, error) {
	return _Labr.Contract.Deposit(&_Labr.TransactOpts, destinationChainID, resourceID, amount, recipientAddress, signature, params)
}

// Deposit is a paid mutator transaction binding the contract method 0x98ce7a18.
//
// Solidity: function deposit(bytes8 destinationChainID, bytes32 resourceID, uint256 amount, address recipientAddress, bytes signature, bytes params) payable returns()
func (_Labr *LabrTransactorSession) Deposit(destinationChainID [8]byte, resourceID [32]byte, amount *big.Int, recipientAddress common.Address, signature []byte, params []byte) (*types.Transaction, error) {
	return _Labr.Contract.Deposit(&_Labr.TransactOpts, destinationChainID, resourceID, amount, recipientAddress, signature, params)
}

// DepositFunds is a paid mutator transaction binding the contract method 0xe2c41dbc.
//
// Solidity: function depositFunds() payable returns()
func (_Labr *LabrTransactor) DepositFunds(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Labr.contract.Transact(opts, "depositFunds")
}

// DepositFunds is a paid mutator transaction binding the contract method 0xe2c41dbc.
//
// Solidity: function depositFunds() payable returns()
func (_Labr *LabrSession) DepositFunds() (*types.Transaction, error) {
	return _Labr.Contract.DepositFunds(&_Labr.TransactOpts)
}

// DepositFunds is a paid mutator transaction binding the contract method 0xe2c41dbc.
//
// Solidity: function depositFunds() payable returns()
func (_Labr *LabrTransactorSession) DepositFunds() (*types.Transaction, error) {
	return _Labr.Contract.DepositFunds(&_Labr.TransactOpts)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0x8d225943.
//
// Solidity: function executeProposal(bytes8 originChainID, bytes8 destinationChainID, uint64 depositNonce, bytes32 resourceID, address recipientAddress, uint256 amount, bytes params) returns()
func (_Labr *LabrTransactor) ExecuteProposal(opts *bind.TransactOpts, originChainID [8]byte, destinationChainID [8]byte, depositNonce uint64, resourceID [32]byte, recipientAddress common.Address, amount *big.Int, params []byte) (*types.Transaction, error) {
	return _Labr.contract.Transact(opts, "executeProposal", originChainID, destinationChainID, depositNonce, resourceID, recipientAddress, amount, params)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0x8d225943.
//
// Solidity: function executeProposal(bytes8 originChainID, bytes8 destinationChainID, uint64 depositNonce, bytes32 resourceID, address recipientAddress, uint256 amount, bytes params) returns()
func (_Labr *LabrSession) ExecuteProposal(originChainID [8]byte, destinationChainID [8]byte, depositNonce uint64, resourceID [32]byte, recipientAddress common.Address, amount *big.Int, params []byte) (*types.Transaction, error) {
	return _Labr.Contract.ExecuteProposal(&_Labr.TransactOpts, originChainID, destinationChainID, depositNonce, resourceID, recipientAddress, amount, params)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0x8d225943.
//
// Solidity: function executeProposal(bytes8 originChainID, bytes8 destinationChainID, uint64 depositNonce, bytes32 resourceID, address recipientAddress, uint256 amount, bytes params) returns()
func (_Labr *LabrTransactorSession) ExecuteProposal(originChainID [8]byte, destinationChainID [8]byte, depositNonce uint64, resourceID [32]byte, recipientAddress common.Address, amount *big.Int, params []byte) (*types.Transaction, error) {
	return _Labr.Contract.ExecuteProposal(&_Labr.TransactOpts, originChainID, destinationChainID, depositNonce, resourceID, recipientAddress, amount, params)
}

// Initialize is a paid mutator transaction binding the contract method 0x37d8d3a8.
//
// Solidity: function initialize(bytes8 chainID_, uint256 relayerThreshold_, uint256 fee_, uint256 expiry_, address ownerAddress, address initBackendSrvAddress_, address initBalancerAddress_) returns()
func (_Labr *LabrTransactor) Initialize(opts *bind.TransactOpts, chainID_ [8]byte, relayerThreshold_ *big.Int, fee_ *big.Int, expiry_ *big.Int, ownerAddress common.Address, initBackendSrvAddress_ common.Address, initBalancerAddress_ common.Address) (*types.Transaction, error) {
	return _Labr.contract.Transact(opts, "initialize", chainID_, relayerThreshold_, fee_, expiry_, ownerAddress, initBackendSrvAddress_, initBalancerAddress_)
}

// Initialize is a paid mutator transaction binding the contract method 0x37d8d3a8.
//
// Solidity: function initialize(bytes8 chainID_, uint256 relayerThreshold_, uint256 fee_, uint256 expiry_, address ownerAddress, address initBackendSrvAddress_, address initBalancerAddress_) returns()
func (_Labr *LabrSession) Initialize(chainID_ [8]byte, relayerThreshold_ *big.Int, fee_ *big.Int, expiry_ *big.Int, ownerAddress common.Address, initBackendSrvAddress_ common.Address, initBalancerAddress_ common.Address) (*types.Transaction, error) {
	return _Labr.Contract.Initialize(&_Labr.TransactOpts, chainID_, relayerThreshold_, fee_, expiry_, ownerAddress, initBackendSrvAddress_, initBalancerAddress_)
}

// Initialize is a paid mutator transaction binding the contract method 0x37d8d3a8.
//
// Solidity: function initialize(bytes8 chainID_, uint256 relayerThreshold_, uint256 fee_, uint256 expiry_, address ownerAddress, address initBackendSrvAddress_, address initBalancerAddress_) returns()
func (_Labr *LabrTransactorSession) Initialize(chainID_ [8]byte, relayerThreshold_ *big.Int, fee_ *big.Int, expiry_ *big.Int, ownerAddress common.Address, initBackendSrvAddress_ common.Address, initBalancerAddress_ common.Address) (*types.Transaction, error) {
	return _Labr.Contract.Initialize(&_Labr.TransactOpts, chainID_, relayerThreshold_, fee_, expiry_, ownerAddress, initBackendSrvAddress_, initBalancerAddress_)
}

// OwnableInit is a paid mutator transaction binding the contract method 0xea439b2b.
//
// Solidity: function ownableInit(address owner_) returns()
func (_Labr *LabrTransactor) OwnableInit(opts *bind.TransactOpts, owner_ common.Address) (*types.Transaction, error) {
	return _Labr.contract.Transact(opts, "ownableInit", owner_)
}

// OwnableInit is a paid mutator transaction binding the contract method 0xea439b2b.
//
// Solidity: function ownableInit(address owner_) returns()
func (_Labr *LabrSession) OwnableInit(owner_ common.Address) (*types.Transaction, error) {
	return _Labr.Contract.OwnableInit(&_Labr.TransactOpts, owner_)
}

// OwnableInit is a paid mutator transaction binding the contract method 0xea439b2b.
//
// Solidity: function ownableInit(address owner_) returns()
func (_Labr *LabrTransactorSession) OwnableInit(owner_ common.Address) (*types.Transaction, error) {
	return _Labr.Contract.OwnableInit(&_Labr.TransactOpts, owner_)
}

// RelayerCollectReward is a paid mutator transaction binding the contract method 0xf408da91.
//
// Solidity: function relayerCollectReward() returns()
func (_Labr *LabrTransactor) RelayerCollectReward(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Labr.contract.Transact(opts, "relayerCollectReward")
}

// RelayerCollectReward is a paid mutator transaction binding the contract method 0xf408da91.
//
// Solidity: function relayerCollectReward() returns()
func (_Labr *LabrSession) RelayerCollectReward() (*types.Transaction, error) {
	return _Labr.Contract.RelayerCollectReward(&_Labr.TransactOpts)
}

// RelayerCollectReward is a paid mutator transaction binding the contract method 0xf408da91.
//
// Solidity: function relayerCollectReward() returns()
func (_Labr *LabrTransactorSession) RelayerCollectReward() (*types.Transaction, error) {
	return _Labr.Contract.RelayerCollectReward(&_Labr.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Labr *LabrTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Labr.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Labr *LabrSession) RenounceOwnership() (*types.Transaction, error) {
	return _Labr.Contract.RenounceOwnership(&_Labr.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Labr *LabrTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Labr.Contract.RenounceOwnership(&_Labr.TransactOpts)
}

// TransferExtraFee is a paid mutator transaction binding the contract method 0xc2a43f17.
//
// Solidity: function transferExtraFee(bytes8 originChainID, bytes8 destinationChainID, uint64 depositNonce, bytes32 resourceID, address recipient, uint256 amount) returns()
func (_Labr *LabrTransactor) TransferExtraFee(opts *bind.TransactOpts, originChainID [8]byte, destinationChainID [8]byte, depositNonce uint64, resourceID [32]byte, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Labr.contract.Transact(opts, "transferExtraFee", originChainID, destinationChainID, depositNonce, resourceID, recipient, amount)
}

// TransferExtraFee is a paid mutator transaction binding the contract method 0xc2a43f17.
//
// Solidity: function transferExtraFee(bytes8 originChainID, bytes8 destinationChainID, uint64 depositNonce, bytes32 resourceID, address recipient, uint256 amount) returns()
func (_Labr *LabrSession) TransferExtraFee(originChainID [8]byte, destinationChainID [8]byte, depositNonce uint64, resourceID [32]byte, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Labr.Contract.TransferExtraFee(&_Labr.TransactOpts, originChainID, destinationChainID, depositNonce, resourceID, recipient, amount)
}

// TransferExtraFee is a paid mutator transaction binding the contract method 0xc2a43f17.
//
// Solidity: function transferExtraFee(bytes8 originChainID, bytes8 destinationChainID, uint64 depositNonce, bytes32 resourceID, address recipient, uint256 amount) returns()
func (_Labr *LabrTransactorSession) TransferExtraFee(originChainID [8]byte, destinationChainID [8]byte, depositNonce uint64, resourceID [32]byte, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Labr.Contract.TransferExtraFee(&_Labr.TransactOpts, originChainID, destinationChainID, depositNonce, resourceID, recipient, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Labr *LabrTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Labr.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Labr *LabrSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Labr.Contract.TransferOwnership(&_Labr.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Labr *LabrTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Labr.Contract.TransferOwnership(&_Labr.TransactOpts, newOwner)
}

// VoteProposal is a paid mutator transaction binding the contract method 0x776860da.
//
// Solidity: function voteProposal(bytes8 originChainID, bytes8 destinationChainID, uint64 depositNonce, bytes32 resourceID, address recipientAddress, uint256 amount) returns()
func (_Labr *LabrTransactor) VoteProposal(opts *bind.TransactOpts, originChainID [8]byte, destinationChainID [8]byte, depositNonce uint64, resourceID [32]byte, recipientAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Labr.contract.Transact(opts, "voteProposal", originChainID, destinationChainID, depositNonce, resourceID, recipientAddress, amount)
}

// VoteProposal is a paid mutator transaction binding the contract method 0x776860da.
//
// Solidity: function voteProposal(bytes8 originChainID, bytes8 destinationChainID, uint64 depositNonce, bytes32 resourceID, address recipientAddress, uint256 amount) returns()
func (_Labr *LabrSession) VoteProposal(originChainID [8]byte, destinationChainID [8]byte, depositNonce uint64, resourceID [32]byte, recipientAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Labr.Contract.VoteProposal(&_Labr.TransactOpts, originChainID, destinationChainID, depositNonce, resourceID, recipientAddress, amount)
}

// VoteProposal is a paid mutator transaction binding the contract method 0x776860da.
//
// Solidity: function voteProposal(bytes8 originChainID, bytes8 destinationChainID, uint64 depositNonce, bytes32 resourceID, address recipientAddress, uint256 amount) returns()
func (_Labr *LabrTransactorSession) VoteProposal(originChainID [8]byte, destinationChainID [8]byte, depositNonce uint64, resourceID [32]byte, recipientAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Labr.Contract.VoteProposal(&_Labr.TransactOpts, originChainID, destinationChainID, depositNonce, resourceID, recipientAddress, amount)
}

// LabrDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Labr contract.
type LabrDepositIterator struct {
	Event *LabrDeposit // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LabrDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LabrDeposit)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LabrDeposit)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LabrDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LabrDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LabrDeposit represents a Deposit event raised by the Labr contract.
type LabrDeposit struct {
	OriginChainID      [8]byte
	DestinationChainID [8]byte
	ResourceID         [32]byte
	DepositNonce       uint64
	Depositor          common.Address
	RecipientAddress   common.Address
	TokenAddress       common.Address
	Amount             *big.Int
	DataHash           [32]byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x370525803ffa9a7c0e6adb3868e393dca45d8b42b2f62fd1f23ecfe99f6ce8fc.
//
// Solidity: event Deposit(bytes8 originChainID, bytes8 destinationChainID, bytes32 resourceID, uint64 depositNonce, address depositor, address recipientAddress, address tokenAddress, uint256 amount, bytes32 dataHash)
func (_Labr *LabrFilterer) FilterDeposit(opts *bind.FilterOpts) (*LabrDepositIterator, error) {

	logs, sub, err := _Labr.contract.FilterLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return &LabrDepositIterator{contract: _Labr.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x370525803ffa9a7c0e6adb3868e393dca45d8b42b2f62fd1f23ecfe99f6ce8fc.
//
// Solidity: event Deposit(bytes8 originChainID, bytes8 destinationChainID, bytes32 resourceID, uint64 depositNonce, address depositor, address recipientAddress, address tokenAddress, uint256 amount, bytes32 dataHash)
func (_Labr *LabrFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *LabrDeposit) (event.Subscription, error) {

	logs, sub, err := _Labr.contract.WatchLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LabrDeposit)
				if err := _Labr.contract.UnpackLog(event, "Deposit", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDeposit is a log parse operation binding the contract event 0x370525803ffa9a7c0e6adb3868e393dca45d8b42b2f62fd1f23ecfe99f6ce8fc.
//
// Solidity: event Deposit(bytes8 originChainID, bytes8 destinationChainID, bytes32 resourceID, uint64 depositNonce, address depositor, address recipientAddress, address tokenAddress, uint256 amount, bytes32 dataHash)
func (_Labr *LabrFilterer) ParseDeposit(log types.Log) (*LabrDeposit, error) {
	event := new(LabrDeposit)
	if err := _Labr.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LabrExtraFeeTransferredIterator is returned from FilterExtraFeeTransferred and is used to iterate over the raw logs and unpacked data for ExtraFeeTransferred events raised by the Labr contract.
type LabrExtraFeeTransferredIterator struct {
	Event *LabrExtraFeeTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LabrExtraFeeTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LabrExtraFeeTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LabrExtraFeeTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LabrExtraFeeTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LabrExtraFeeTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LabrExtraFeeTransferred represents a ExtraFeeTransferred event raised by the Labr contract.
type LabrExtraFeeTransferred struct {
	OriginChainID      [8]byte
	DestinationChainID [8]byte
	DepositNonce       uint64
	ResouceID          [32]byte
	Recipient          common.Address
	Amount             *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterExtraFeeTransferred is a free log retrieval operation binding the contract event 0x11d9d6d82ced6158185f9c4a4ab3c7768ed3c14cbd759491ca5d2fb42b7935fd.
//
// Solidity: event ExtraFeeTransferred(bytes8 originChainID, bytes8 destinationChainID, uint64 depositNonce, bytes32 resouceID, address recipient, uint256 amount)
func (_Labr *LabrFilterer) FilterExtraFeeTransferred(opts *bind.FilterOpts) (*LabrExtraFeeTransferredIterator, error) {

	logs, sub, err := _Labr.contract.FilterLogs(opts, "ExtraFeeTransferred")
	if err != nil {
		return nil, err
	}
	return &LabrExtraFeeTransferredIterator{contract: _Labr.contract, event: "ExtraFeeTransferred", logs: logs, sub: sub}, nil
}

// WatchExtraFeeTransferred is a free log subscription operation binding the contract event 0x11d9d6d82ced6158185f9c4a4ab3c7768ed3c14cbd759491ca5d2fb42b7935fd.
//
// Solidity: event ExtraFeeTransferred(bytes8 originChainID, bytes8 destinationChainID, uint64 depositNonce, bytes32 resouceID, address recipient, uint256 amount)
func (_Labr *LabrFilterer) WatchExtraFeeTransferred(opts *bind.WatchOpts, sink chan<- *LabrExtraFeeTransferred) (event.Subscription, error) {

	logs, sub, err := _Labr.contract.WatchLogs(opts, "ExtraFeeTransferred")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LabrExtraFeeTransferred)
				if err := _Labr.contract.UnpackLog(event, "ExtraFeeTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseExtraFeeTransferred is a log parse operation binding the contract event 0x11d9d6d82ced6158185f9c4a4ab3c7768ed3c14cbd759491ca5d2fb42b7935fd.
//
// Solidity: event ExtraFeeTransferred(bytes8 originChainID, bytes8 destinationChainID, uint64 depositNonce, bytes32 resouceID, address recipient, uint256 amount)
func (_Labr *LabrFilterer) ParseExtraFeeTransferred(log types.Log) (*LabrExtraFeeTransferred, error) {
	event := new(LabrExtraFeeTransferred)
	if err := _Labr.contract.UnpackLog(event, "ExtraFeeTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LabrOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Labr contract.
type LabrOwnershipTransferredIterator struct {
	Event *LabrOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LabrOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LabrOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LabrOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LabrOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LabrOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LabrOwnershipTransferred represents a OwnershipTransferred event raised by the Labr contract.
type LabrOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Labr *LabrFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*LabrOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Labr.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &LabrOwnershipTransferredIterator{contract: _Labr.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Labr *LabrFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *LabrOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Labr.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LabrOwnershipTransferred)
				if err := _Labr.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Labr *LabrFilterer) ParseOwnershipTransferred(log types.Log) (*LabrOwnershipTransferred, error) {
	event := new(LabrOwnershipTransferred)
	if err := _Labr.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LabrPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Labr contract.
type LabrPausedIterator struct {
	Event *LabrPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LabrPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LabrPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LabrPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LabrPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LabrPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LabrPaused represents a Paused event raised by the Labr contract.
type LabrPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Labr *LabrFilterer) FilterPaused(opts *bind.FilterOpts) (*LabrPausedIterator, error) {

	logs, sub, err := _Labr.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &LabrPausedIterator{contract: _Labr.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Labr *LabrFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *LabrPaused) (event.Subscription, error) {

	logs, sub, err := _Labr.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LabrPaused)
				if err := _Labr.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Labr *LabrFilterer) ParsePaused(log types.Log) (*LabrPaused, error) {
	event := new(LabrPaused)
	if err := _Labr.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LabrProposalEventIterator is returned from FilterProposalEvent and is used to iterate over the raw logs and unpacked data for ProposalEvent events raised by the Labr contract.
type LabrProposalEventIterator struct {
	Event *LabrProposalEvent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LabrProposalEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LabrProposalEvent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LabrProposalEvent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LabrProposalEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LabrProposalEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LabrProposalEvent represents a ProposalEvent event raised by the Labr contract.
type LabrProposalEvent struct {
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

// FilterProposalEvent is a free log retrieval operation binding the contract event 0x9686dcabd0450cad86a88df15a9d35b08b35d1b08a19008df37cf8538c467516.
//
// Solidity: event ProposalEvent(bytes8 originChainID, bytes8 destinationChainID, address recipientAddress, uint256 amount, uint64 depositNonce, uint8 status, bytes32 resourceID, bytes32 dataHash)
func (_Labr *LabrFilterer) FilterProposalEvent(opts *bind.FilterOpts) (*LabrProposalEventIterator, error) {

	logs, sub, err := _Labr.contract.FilterLogs(opts, "ProposalEvent")
	if err != nil {
		return nil, err
	}
	return &LabrProposalEventIterator{contract: _Labr.contract, event: "ProposalEvent", logs: logs, sub: sub}, nil
}

// WatchProposalEvent is a free log subscription operation binding the contract event 0x9686dcabd0450cad86a88df15a9d35b08b35d1b08a19008df37cf8538c467516.
//
// Solidity: event ProposalEvent(bytes8 originChainID, bytes8 destinationChainID, address recipientAddress, uint256 amount, uint64 depositNonce, uint8 status, bytes32 resourceID, bytes32 dataHash)
func (_Labr *LabrFilterer) WatchProposalEvent(opts *bind.WatchOpts, sink chan<- *LabrProposalEvent) (event.Subscription, error) {

	logs, sub, err := _Labr.contract.WatchLogs(opts, "ProposalEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LabrProposalEvent)
				if err := _Labr.contract.UnpackLog(event, "ProposalEvent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseProposalEvent is a log parse operation binding the contract event 0x9686dcabd0450cad86a88df15a9d35b08b35d1b08a19008df37cf8538c467516.
//
// Solidity: event ProposalEvent(bytes8 originChainID, bytes8 destinationChainID, address recipientAddress, uint256 amount, uint64 depositNonce, uint8 status, bytes32 resourceID, bytes32 dataHash)
func (_Labr *LabrFilterer) ParseProposalEvent(log types.Log) (*LabrProposalEvent, error) {
	event := new(LabrProposalEvent)
	if err := _Labr.contract.UnpackLog(event, "ProposalEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LabrProposalVoteIterator is returned from FilterProposalVote and is used to iterate over the raw logs and unpacked data for ProposalVote events raised by the Labr contract.
type LabrProposalVoteIterator struct {
	Event *LabrProposalVote // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LabrProposalVoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LabrProposalVote)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LabrProposalVote)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LabrProposalVoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LabrProposalVoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LabrProposalVote represents a ProposalVote event raised by the Labr contract.
type LabrProposalVote struct {
	OriginChainID [8]byte
	DepositNonce  uint64
	Status        uint8
	ResourceID    [32]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterProposalVote is a free log retrieval operation binding the contract event 0x85f41114efc645854a10eef33ef4dec54341cb3ec3ab32386c92c881f3b1b505.
//
// Solidity: event ProposalVote(bytes8 originChainID, uint64 depositNonce, uint8 status, bytes32 resourceID)
func (_Labr *LabrFilterer) FilterProposalVote(opts *bind.FilterOpts) (*LabrProposalVoteIterator, error) {

	logs, sub, err := _Labr.contract.FilterLogs(opts, "ProposalVote")
	if err != nil {
		return nil, err
	}
	return &LabrProposalVoteIterator{contract: _Labr.contract, event: "ProposalVote", logs: logs, sub: sub}, nil
}

// WatchProposalVote is a free log subscription operation binding the contract event 0x85f41114efc645854a10eef33ef4dec54341cb3ec3ab32386c92c881f3b1b505.
//
// Solidity: event ProposalVote(bytes8 originChainID, uint64 depositNonce, uint8 status, bytes32 resourceID)
func (_Labr *LabrFilterer) WatchProposalVote(opts *bind.WatchOpts, sink chan<- *LabrProposalVote) (event.Subscription, error) {

	logs, sub, err := _Labr.contract.WatchLogs(opts, "ProposalVote")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LabrProposalVote)
				if err := _Labr.contract.UnpackLog(event, "ProposalVote", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseProposalVote is a log parse operation binding the contract event 0x85f41114efc645854a10eef33ef4dec54341cb3ec3ab32386c92c881f3b1b505.
//
// Solidity: event ProposalVote(bytes8 originChainID, uint64 depositNonce, uint8 status, bytes32 resourceID)
func (_Labr *LabrFilterer) ParseProposalVote(log types.Log) (*LabrProposalVote, error) {
	event := new(LabrProposalVote)
	if err := _Labr.contract.UnpackLog(event, "ProposalVote", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LabrRelayerThresholdChangedIterator is returned from FilterRelayerThresholdChanged and is used to iterate over the raw logs and unpacked data for RelayerThresholdChanged events raised by the Labr contract.
type LabrRelayerThresholdChangedIterator struct {
	Event *LabrRelayerThresholdChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LabrRelayerThresholdChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LabrRelayerThresholdChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LabrRelayerThresholdChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LabrRelayerThresholdChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LabrRelayerThresholdChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LabrRelayerThresholdChanged represents a RelayerThresholdChanged event raised by the Labr contract.
type LabrRelayerThresholdChanged struct {
	NewThreshold *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRelayerThresholdChanged is a free log retrieval operation binding the contract event 0xa20d6b84cd798a24038be305eff8a45ca82ef54a2aa2082005d8e14c0a4746c8.
//
// Solidity: event RelayerThresholdChanged(uint256 newThreshold)
func (_Labr *LabrFilterer) FilterRelayerThresholdChanged(opts *bind.FilterOpts) (*LabrRelayerThresholdChangedIterator, error) {

	logs, sub, err := _Labr.contract.FilterLogs(opts, "RelayerThresholdChanged")
	if err != nil {
		return nil, err
	}
	return &LabrRelayerThresholdChangedIterator{contract: _Labr.contract, event: "RelayerThresholdChanged", logs: logs, sub: sub}, nil
}

// WatchRelayerThresholdChanged is a free log subscription operation binding the contract event 0xa20d6b84cd798a24038be305eff8a45ca82ef54a2aa2082005d8e14c0a4746c8.
//
// Solidity: event RelayerThresholdChanged(uint256 newThreshold)
func (_Labr *LabrFilterer) WatchRelayerThresholdChanged(opts *bind.WatchOpts, sink chan<- *LabrRelayerThresholdChanged) (event.Subscription, error) {

	logs, sub, err := _Labr.contract.WatchLogs(opts, "RelayerThresholdChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LabrRelayerThresholdChanged)
				if err := _Labr.contract.UnpackLog(event, "RelayerThresholdChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRelayerThresholdChanged is a log parse operation binding the contract event 0xa20d6b84cd798a24038be305eff8a45ca82ef54a2aa2082005d8e14c0a4746c8.
//
// Solidity: event RelayerThresholdChanged(uint256 newThreshold)
func (_Labr *LabrFilterer) ParseRelayerThresholdChanged(log types.Log) (*LabrRelayerThresholdChanged, error) {
	event := new(LabrRelayerThresholdChanged)
	if err := _Labr.contract.UnpackLog(event, "RelayerThresholdChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LabrRewardCollectedIterator is returned from FilterRewardCollected and is used to iterate over the raw logs and unpacked data for RewardCollected events raised by the Labr contract.
type LabrRewardCollectedIterator struct {
	Event *LabrRewardCollected // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LabrRewardCollectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LabrRewardCollected)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LabrRewardCollected)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LabrRewardCollectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LabrRewardCollectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LabrRewardCollected represents a RewardCollected event raised by the Labr contract.
type LabrRewardCollected struct {
	Relayer common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRewardCollected is a free log retrieval operation binding the contract event 0xe8354b169cd993d5cdfad1036a9a3f1ea7ed77e430bccb279200fd088243f595.
//
// Solidity: event RewardCollected(address relayer, uint256 amount)
func (_Labr *LabrFilterer) FilterRewardCollected(opts *bind.FilterOpts) (*LabrRewardCollectedIterator, error) {

	logs, sub, err := _Labr.contract.FilterLogs(opts, "RewardCollected")
	if err != nil {
		return nil, err
	}
	return &LabrRewardCollectedIterator{contract: _Labr.contract, event: "RewardCollected", logs: logs, sub: sub}, nil
}

// WatchRewardCollected is a free log subscription operation binding the contract event 0xe8354b169cd993d5cdfad1036a9a3f1ea7ed77e430bccb279200fd088243f595.
//
// Solidity: event RewardCollected(address relayer, uint256 amount)
func (_Labr *LabrFilterer) WatchRewardCollected(opts *bind.WatchOpts, sink chan<- *LabrRewardCollected) (event.Subscription, error) {

	logs, sub, err := _Labr.contract.WatchLogs(opts, "RewardCollected")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LabrRewardCollected)
				if err := _Labr.contract.UnpackLog(event, "RewardCollected", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRewardCollected is a log parse operation binding the contract event 0xe8354b169cd993d5cdfad1036a9a3f1ea7ed77e430bccb279200fd088243f595.
//
// Solidity: event RewardCollected(address relayer, uint256 amount)
func (_Labr *LabrFilterer) ParseRewardCollected(log types.Log) (*LabrRewardCollected, error) {
	event := new(LabrRewardCollected)
	if err := _Labr.contract.UnpackLog(event, "RewardCollected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LabrUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Labr contract.
type LabrUnpausedIterator struct {
	Event *LabrUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LabrUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LabrUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LabrUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LabrUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LabrUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LabrUnpaused represents a Unpaused event raised by the Labr contract.
type LabrUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Labr *LabrFilterer) FilterUnpaused(opts *bind.FilterOpts) (*LabrUnpausedIterator, error) {

	logs, sub, err := _Labr.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &LabrUnpausedIterator{contract: _Labr.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Labr *LabrFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *LabrUnpaused) (event.Subscription, error) {

	logs, sub, err := _Labr.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LabrUnpaused)
				if err := _Labr.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Labr *LabrFilterer) ParseUnpaused(log types.Log) (*LabrUnpaused, error) {
	event := new(LabrUnpaused)
	if err := _Labr.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
