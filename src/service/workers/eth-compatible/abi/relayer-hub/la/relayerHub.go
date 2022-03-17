// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package hubLA

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

// HubLAMetaData contains all meta data concerning the HubLA contract.
var HubLAMetaData = &bind.MetaData{
	ABI: "[{\"name\":\"owner\",\"type\":\"function\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"name\":\"renounceOwnership\",\"type\":\"function\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"name\":\"transferOwnership\",\"type\":\"function\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"name\":\"_relayerBlocked\",\"type\":\"function\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"name\":\"_penalty\",\"type\":\"function\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"name\":\"_penaltyPercentage\",\"type\":\"function\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"name\":\"_relayerRewards\",\"type\":\"function\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"name\":\"_rewardPercentage\",\"type\":\"function\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"name\":\"_systemRewardAddress\",\"type\":\"function\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"name\":\"_bridgeAddress\",\"type\":\"function\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"name\":\"_requiredDeposit\",\"type\":\"function\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"name\":\"_dues\",\"type\":\"function\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"name\":\"INIT_REQUIRED_DEPOSIT\",\"type\":\"function\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"name\":\"INIT_DUES\",\"type\":\"function\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"name\":\"INIT_PENALTY_PERCENTAGE\",\"type\":\"function\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"name\":\"INIT_REWARD_PERCENTAGE\",\"type\":\"function\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"name\":\"init\",\"type\":\"function\",\"inputs\":[{\"name\":\"initBridgeAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"name\":\"register\",\"type\":\"function\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"name\":\"unregister\",\"type\":\"function\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"name\":\"isRelayer\",\"type\":\"function\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"name\":\"updateSystemRewardAddr\",\"type\":\"function\",\"inputs\":[{\"name\":\"newAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"name\":\"alreadyExecuted\",\"type\":\"function\",\"inputs\":[{\"name\":\"_relayer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"swapValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"name\":\"alreadyVoted\",\"type\":\"function\",\"inputs\":[{\"name\":\"_relayer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"swapValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"name\":\"slash\",\"type\":\"function\",\"inputs\":[{\"name\":\"_relayer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"penalty\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"name\":\"felony\",\"type\":\"function\",\"inputs\":[{\"name\":\"_relayer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"name\":\"updatePenaltyPercentage\",\"type\":\"function\",\"inputs\":[{\"name\":\"penaltyPercentage\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"name\":\"swapReward\",\"type\":\"function\",\"inputs\":[{\"name\":\"_relayer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"swapValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"name\":\"updateRewardPercentage\",\"type\":\"function\",\"inputs\":[{\"name\":\"rewardPercent\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"name\":\"addReward\",\"type\":\"function\",\"inputs\":[{\"name\":\"_relayer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reward\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"name\":\"getTotalReward\",\"type\":\"function\",\"inputs\":[{\"name\":\"_relayer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"totalReward\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"name\":\"updateReward\",\"type\":\"function\",\"inputs\":[{\"name\":\"_relayer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reward\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"name\":\"OwnershipTransferred\",\"type\":\"event\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"internalType\":\"address\",\"indexed\":true},{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\",\"indexed\":true}],\"anonymous\":false},{\"name\":\"RelayerBlocked\",\"type\":\"event\",\"inputs\":[{\"name\":\"relayer\",\"type\":\"address\",\"internalType\":\"address\"}],\"anonymous\":false},{\"name\":\"PenaltyPercentChanged\",\"type\":\"event\",\"inputs\":[{\"name\":\"newPercent\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"anonymous\":false},{\"name\":\"RewardPercentageChanged\",\"type\":\"event\",\"inputs\":[{\"name\":\"newPercentage\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"anonymous\":false},{\"name\":\"RelayerRegister\",\"type\":\"event\",\"inputs\":[{\"name\":\"_relayer\",\"type\":\"address\",\"internalType\":\"address\"}],\"anonymous\":false},{\"name\":\"RelayerUnRegister\",\"type\":\"event\",\"inputs\":[{\"name\":\"_relayer\",\"type\":\"address\",\"internalType\":\"address\"}],\"anonymous\":false},{\"name\":\"SystemRewardAddressChanged\",\"type\":\"event\",\"inputs\":[{\"name\":\"newAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"anonymous\":false},{\"name\":\"PenaltyForRelayer\",\"type\":\"event\",\"inputs\":[{\"name\":\"reason\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"relayer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"penalty\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"anonymous\":false},{\"name\":\"RewardForRelayer\",\"type\":\"event\",\"inputs\":[{\"name\":\"relayer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reward\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"anonymous\":false}]",
}

// HubLAABI is the input ABI used to generate the binding from.
// Deprecated: Use HubLAMetaData.ABI instead.
var HubLAABI = HubLAMetaData.ABI

// HubLA is an auto generated Go binding around an Ethereum contract.
type HubLA struct {
	HubLACaller     // Read-only binding to the contract
	HubLATransactor // Write-only binding to the contract
	HubLAFilterer   // Log filterer for contract events
}

// HubLACaller is an auto generated read-only Go binding around an Ethereum contract.
type HubLACaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HubLATransactor is an auto generated write-only Go binding around an Ethereum contract.
type HubLATransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HubLAFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HubLAFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HubLASession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HubLASession struct {
	Contract     *HubLA            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HubLACallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HubLACallerSession struct {
	Contract *HubLACaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// HubLATransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HubLATransactorSession struct {
	Contract     *HubLATransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HubLARaw is an auto generated low-level Go binding around an Ethereum contract.
type HubLARaw struct {
	Contract *HubLA // Generic contract binding to access the raw methods on
}

// HubLACallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HubLACallerRaw struct {
	Contract *HubLACaller // Generic read-only contract binding to access the raw methods on
}

// HubLATransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HubLATransactorRaw struct {
	Contract *HubLATransactor // Generic write-only contract binding to access the raw methods on
}

// NewHubLA creates a new instance of HubLA, bound to a specific deployed contract.
func NewHubLA(address common.Address, backend bind.ContractBackend) (*HubLA, error) {
	contract, err := bindHubLA(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HubLA{HubLACaller: HubLACaller{contract: contract}, HubLATransactor: HubLATransactor{contract: contract}, HubLAFilterer: HubLAFilterer{contract: contract}}, nil
}

// NewHubLACaller creates a new read-only instance of HubLA, bound to a specific deployed contract.
func NewHubLACaller(address common.Address, caller bind.ContractCaller) (*HubLACaller, error) {
	contract, err := bindHubLA(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HubLACaller{contract: contract}, nil
}

// NewHubLATransactor creates a new write-only instance of HubLA, bound to a specific deployed contract.
func NewHubLATransactor(address common.Address, transactor bind.ContractTransactor) (*HubLATransactor, error) {
	contract, err := bindHubLA(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HubLATransactor{contract: contract}, nil
}

// NewHubLAFilterer creates a new log filterer instance of HubLA, bound to a specific deployed contract.
func NewHubLAFilterer(address common.Address, filterer bind.ContractFilterer) (*HubLAFilterer, error) {
	contract, err := bindHubLA(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HubLAFilterer{contract: contract}, nil
}

// bindHubLA binds a generic wrapper to an already deployed contract.
func bindHubLA(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HubLAABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HubLA *HubLARaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HubLA.Contract.HubLACaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HubLA *HubLARaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HubLA.Contract.HubLATransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HubLA *HubLARaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HubLA.Contract.HubLATransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HubLA *HubLACallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HubLA.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HubLA *HubLATransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HubLA.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HubLA *HubLATransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HubLA.Contract.contract.Transact(opts, method, params...)
}

// INITDUES is a free data retrieval call binding the contract method 0x95468d26.
//
// Solidity: function INIT_DUES() view returns(uint256)
func (_HubLA *HubLACaller) INITDUES(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _HubLA.contract.Call(opts, &out, "INIT_DUES")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// INITDUES is a free data retrieval call binding the contract method 0x95468d26.
//
// Solidity: function INIT_DUES() view returns(uint256)
func (_HubLA *HubLASession) INITDUES() (*big.Int, error) {
	return _HubLA.Contract.INITDUES(&_HubLA.CallOpts)
}

// INITDUES is a free data retrieval call binding the contract method 0x95468d26.
//
// Solidity: function INIT_DUES() view returns(uint256)
func (_HubLA *HubLACallerSession) INITDUES() (*big.Int, error) {
	return _HubLA.Contract.INITDUES(&_HubLA.CallOpts)
}

// INITPENALTYPERCENTAGE is a free data retrieval call binding the contract method 0xdf9c203f.
//
// Solidity: function INIT_PENALTY_PERCENTAGE() view returns(uint256)
func (_HubLA *HubLACaller) INITPENALTYPERCENTAGE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _HubLA.contract.Call(opts, &out, "INIT_PENALTY_PERCENTAGE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// INITPENALTYPERCENTAGE is a free data retrieval call binding the contract method 0xdf9c203f.
//
// Solidity: function INIT_PENALTY_PERCENTAGE() view returns(uint256)
func (_HubLA *HubLASession) INITPENALTYPERCENTAGE() (*big.Int, error) {
	return _HubLA.Contract.INITPENALTYPERCENTAGE(&_HubLA.CallOpts)
}

// INITPENALTYPERCENTAGE is a free data retrieval call binding the contract method 0xdf9c203f.
//
// Solidity: function INIT_PENALTY_PERCENTAGE() view returns(uint256)
func (_HubLA *HubLACallerSession) INITPENALTYPERCENTAGE() (*big.Int, error) {
	return _HubLA.Contract.INITPENALTYPERCENTAGE(&_HubLA.CallOpts)
}

// INITREQUIREDDEPOSIT is a free data retrieval call binding the contract method 0x7ae23088.
//
// Solidity: function INIT_REQUIRED_DEPOSIT() view returns(uint256)
func (_HubLA *HubLACaller) INITREQUIREDDEPOSIT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _HubLA.contract.Call(opts, &out, "INIT_REQUIRED_DEPOSIT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// INITREQUIREDDEPOSIT is a free data retrieval call binding the contract method 0x7ae23088.
//
// Solidity: function INIT_REQUIRED_DEPOSIT() view returns(uint256)
func (_HubLA *HubLASession) INITREQUIREDDEPOSIT() (*big.Int, error) {
	return _HubLA.Contract.INITREQUIREDDEPOSIT(&_HubLA.CallOpts)
}

// INITREQUIREDDEPOSIT is a free data retrieval call binding the contract method 0x7ae23088.
//
// Solidity: function INIT_REQUIRED_DEPOSIT() view returns(uint256)
func (_HubLA *HubLACallerSession) INITREQUIREDDEPOSIT() (*big.Int, error) {
	return _HubLA.Contract.INITREQUIREDDEPOSIT(&_HubLA.CallOpts)
}

// INITREWARDPERCENTAGE is a free data retrieval call binding the contract method 0xa49909fb.
//
// Solidity: function INIT_REWARD_PERCENTAGE() view returns(uint256)
func (_HubLA *HubLACaller) INITREWARDPERCENTAGE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _HubLA.contract.Call(opts, &out, "INIT_REWARD_PERCENTAGE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// INITREWARDPERCENTAGE is a free data retrieval call binding the contract method 0xa49909fb.
//
// Solidity: function INIT_REWARD_PERCENTAGE() view returns(uint256)
func (_HubLA *HubLASession) INITREWARDPERCENTAGE() (*big.Int, error) {
	return _HubLA.Contract.INITREWARDPERCENTAGE(&_HubLA.CallOpts)
}

// INITREWARDPERCENTAGE is a free data retrieval call binding the contract method 0xa49909fb.
//
// Solidity: function INIT_REWARD_PERCENTAGE() view returns(uint256)
func (_HubLA *HubLACallerSession) INITREWARDPERCENTAGE() (*big.Int, error) {
	return _HubLA.Contract.INITREWARDPERCENTAGE(&_HubLA.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0x318c136e.
//
// Solidity: function _bridgeAddress() view returns(address)
func (_HubLA *HubLACaller) BridgeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HubLA.contract.Call(opts, &out, "_bridgeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeAddress is a free data retrieval call binding the contract method 0x318c136e.
//
// Solidity: function _bridgeAddress() view returns(address)
func (_HubLA *HubLASession) BridgeAddress() (common.Address, error) {
	return _HubLA.Contract.BridgeAddress(&_HubLA.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0x318c136e.
//
// Solidity: function _bridgeAddress() view returns(address)
func (_HubLA *HubLACallerSession) BridgeAddress() (common.Address, error) {
	return _HubLA.Contract.BridgeAddress(&_HubLA.CallOpts)
}

// Dues is a free data retrieval call binding the contract method 0x5bc89bb7.
//
// Solidity: function _dues() view returns(uint256)
func (_HubLA *HubLACaller) Dues(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _HubLA.contract.Call(opts, &out, "_dues")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Dues is a free data retrieval call binding the contract method 0x5bc89bb7.
//
// Solidity: function _dues() view returns(uint256)
func (_HubLA *HubLASession) Dues() (*big.Int, error) {
	return _HubLA.Contract.Dues(&_HubLA.CallOpts)
}

// Dues is a free data retrieval call binding the contract method 0x5bc89bb7.
//
// Solidity: function _dues() view returns(uint256)
func (_HubLA *HubLACallerSession) Dues() (*big.Int, error) {
	return _HubLA.Contract.Dues(&_HubLA.CallOpts)
}

// Penalty is a free data retrieval call binding the contract method 0xe722ec53.
//
// Solidity: function _penalty(address ) view returns(uint256)
func (_HubLA *HubLACaller) Penalty(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _HubLA.contract.Call(opts, &out, "_penalty", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Penalty is a free data retrieval call binding the contract method 0xe722ec53.
//
// Solidity: function _penalty(address ) view returns(uint256)
func (_HubLA *HubLASession) Penalty(arg0 common.Address) (*big.Int, error) {
	return _HubLA.Contract.Penalty(&_HubLA.CallOpts, arg0)
}

// Penalty is a free data retrieval call binding the contract method 0xe722ec53.
//
// Solidity: function _penalty(address ) view returns(uint256)
func (_HubLA *HubLACallerSession) Penalty(arg0 common.Address) (*big.Int, error) {
	return _HubLA.Contract.Penalty(&_HubLA.CallOpts, arg0)
}

// PenaltyPercentage is a free data retrieval call binding the contract method 0x93a55f92.
//
// Solidity: function _penaltyPercentage() view returns(uint256)
func (_HubLA *HubLACaller) PenaltyPercentage(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _HubLA.contract.Call(opts, &out, "_penaltyPercentage")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PenaltyPercentage is a free data retrieval call binding the contract method 0x93a55f92.
//
// Solidity: function _penaltyPercentage() view returns(uint256)
func (_HubLA *HubLASession) PenaltyPercentage() (*big.Int, error) {
	return _HubLA.Contract.PenaltyPercentage(&_HubLA.CallOpts)
}

// PenaltyPercentage is a free data retrieval call binding the contract method 0x93a55f92.
//
// Solidity: function _penaltyPercentage() view returns(uint256)
func (_HubLA *HubLACallerSession) PenaltyPercentage() (*big.Int, error) {
	return _HubLA.Contract.PenaltyPercentage(&_HubLA.CallOpts)
}

// RelayerBlocked is a free data retrieval call binding the contract method 0x0537b4d6.
//
// Solidity: function _relayerBlocked(address ) view returns(bool)
func (_HubLA *HubLACaller) RelayerBlocked(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _HubLA.contract.Call(opts, &out, "_relayerBlocked", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// RelayerBlocked is a free data retrieval call binding the contract method 0x0537b4d6.
//
// Solidity: function _relayerBlocked(address ) view returns(bool)
func (_HubLA *HubLASession) RelayerBlocked(arg0 common.Address) (bool, error) {
	return _HubLA.Contract.RelayerBlocked(&_HubLA.CallOpts, arg0)
}

// RelayerBlocked is a free data retrieval call binding the contract method 0x0537b4d6.
//
// Solidity: function _relayerBlocked(address ) view returns(bool)
func (_HubLA *HubLACallerSession) RelayerBlocked(arg0 common.Address) (bool, error) {
	return _HubLA.Contract.RelayerBlocked(&_HubLA.CallOpts, arg0)
}

// RelayerRewards is a free data retrieval call binding the contract method 0x13bb6094.
//
// Solidity: function _relayerRewards(address ) view returns(uint256)
func (_HubLA *HubLACaller) RelayerRewards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _HubLA.contract.Call(opts, &out, "_relayerRewards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RelayerRewards is a free data retrieval call binding the contract method 0x13bb6094.
//
// Solidity: function _relayerRewards(address ) view returns(uint256)
func (_HubLA *HubLASession) RelayerRewards(arg0 common.Address) (*big.Int, error) {
	return _HubLA.Contract.RelayerRewards(&_HubLA.CallOpts, arg0)
}

// RelayerRewards is a free data retrieval call binding the contract method 0x13bb6094.
//
// Solidity: function _relayerRewards(address ) view returns(uint256)
func (_HubLA *HubLACallerSession) RelayerRewards(arg0 common.Address) (*big.Int, error) {
	return _HubLA.Contract.RelayerRewards(&_HubLA.CallOpts, arg0)
}

// RequiredDeposit is a free data retrieval call binding the contract method 0x003ee503.
//
// Solidity: function _requiredDeposit() view returns(uint256)
func (_HubLA *HubLACaller) RequiredDeposit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _HubLA.contract.Call(opts, &out, "_requiredDeposit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RequiredDeposit is a free data retrieval call binding the contract method 0x003ee503.
//
// Solidity: function _requiredDeposit() view returns(uint256)
func (_HubLA *HubLASession) RequiredDeposit() (*big.Int, error) {
	return _HubLA.Contract.RequiredDeposit(&_HubLA.CallOpts)
}

// RequiredDeposit is a free data retrieval call binding the contract method 0x003ee503.
//
// Solidity: function _requiredDeposit() view returns(uint256)
func (_HubLA *HubLACallerSession) RequiredDeposit() (*big.Int, error) {
	return _HubLA.Contract.RequiredDeposit(&_HubLA.CallOpts)
}

// RewardPercentage is a free data retrieval call binding the contract method 0xa1444a4a.
//
// Solidity: function _rewardPercentage() view returns(uint256)
func (_HubLA *HubLACaller) RewardPercentage(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _HubLA.contract.Call(opts, &out, "_rewardPercentage")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardPercentage is a free data retrieval call binding the contract method 0xa1444a4a.
//
// Solidity: function _rewardPercentage() view returns(uint256)
func (_HubLA *HubLASession) RewardPercentage() (*big.Int, error) {
	return _HubLA.Contract.RewardPercentage(&_HubLA.CallOpts)
}

// RewardPercentage is a free data retrieval call binding the contract method 0xa1444a4a.
//
// Solidity: function _rewardPercentage() view returns(uint256)
func (_HubLA *HubLACallerSession) RewardPercentage() (*big.Int, error) {
	return _HubLA.Contract.RewardPercentage(&_HubLA.CallOpts)
}

// SystemRewardAddress is a free data retrieval call binding the contract method 0x1b3566c8.
//
// Solidity: function _systemRewardAddress() view returns(address)
func (_HubLA *HubLACaller) SystemRewardAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HubLA.contract.Call(opts, &out, "_systemRewardAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SystemRewardAddress is a free data retrieval call binding the contract method 0x1b3566c8.
//
// Solidity: function _systemRewardAddress() view returns(address)
func (_HubLA *HubLASession) SystemRewardAddress() (common.Address, error) {
	return _HubLA.Contract.SystemRewardAddress(&_HubLA.CallOpts)
}

// SystemRewardAddress is a free data retrieval call binding the contract method 0x1b3566c8.
//
// Solidity: function _systemRewardAddress() view returns(address)
func (_HubLA *HubLACallerSession) SystemRewardAddress() (common.Address, error) {
	return _HubLA.Contract.SystemRewardAddress(&_HubLA.CallOpts)
}

// GetTotalReward is a free data retrieval call binding the contract method 0x34381749.
//
// Solidity: function getTotalReward(address _relayer) view returns(uint256 totalReward)
func (_HubLA *HubLACaller) GetTotalReward(opts *bind.CallOpts, _relayer common.Address) (*big.Int, error) {
	var out []interface{}
	err := _HubLA.contract.Call(opts, &out, "getTotalReward", _relayer)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalReward is a free data retrieval call binding the contract method 0x34381749.
//
// Solidity: function getTotalReward(address _relayer) view returns(uint256 totalReward)
func (_HubLA *HubLASession) GetTotalReward(_relayer common.Address) (*big.Int, error) {
	return _HubLA.Contract.GetTotalReward(&_HubLA.CallOpts, _relayer)
}

// GetTotalReward is a free data retrieval call binding the contract method 0x34381749.
//
// Solidity: function getTotalReward(address _relayer) view returns(uint256 totalReward)
func (_HubLA *HubLACallerSession) GetTotalReward(_relayer common.Address) (*big.Int, error) {
	return _HubLA.Contract.GetTotalReward(&_HubLA.CallOpts, _relayer)
}

// IsRelayer is a free data retrieval call binding the contract method 0x541d5548.
//
// Solidity: function isRelayer(address sender) view returns(bool)
func (_HubLA *HubLACaller) IsRelayer(opts *bind.CallOpts, sender common.Address) (bool, error) {
	var out []interface{}
	err := _HubLA.contract.Call(opts, &out, "isRelayer", sender)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRelayer is a free data retrieval call binding the contract method 0x541d5548.
//
// Solidity: function isRelayer(address sender) view returns(bool)
func (_HubLA *HubLASession) IsRelayer(sender common.Address) (bool, error) {
	return _HubLA.Contract.IsRelayer(&_HubLA.CallOpts, sender)
}

// IsRelayer is a free data retrieval call binding the contract method 0x541d5548.
//
// Solidity: function isRelayer(address sender) view returns(bool)
func (_HubLA *HubLACallerSession) IsRelayer(sender common.Address) (bool, error) {
	return _HubLA.Contract.IsRelayer(&_HubLA.CallOpts, sender)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_HubLA *HubLACaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HubLA.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_HubLA *HubLASession) Owner() (common.Address, error) {
	return _HubLA.Contract.Owner(&_HubLA.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_HubLA *HubLACallerSession) Owner() (common.Address, error) {
	return _HubLA.Contract.Owner(&_HubLA.CallOpts)
}

// AddReward is a paid mutator transaction binding the contract method 0x9feb8f50.
//
// Solidity: function addReward(address _relayer, uint256 reward) returns()
func (_HubLA *HubLATransactor) AddReward(opts *bind.TransactOpts, _relayer common.Address, reward *big.Int) (*types.Transaction, error) {
	return _HubLA.contract.Transact(opts, "addReward", _relayer, reward)
}

// AddReward is a paid mutator transaction binding the contract method 0x9feb8f50.
//
// Solidity: function addReward(address _relayer, uint256 reward) returns()
func (_HubLA *HubLASession) AddReward(_relayer common.Address, reward *big.Int) (*types.Transaction, error) {
	return _HubLA.Contract.AddReward(&_HubLA.TransactOpts, _relayer, reward)
}

// AddReward is a paid mutator transaction binding the contract method 0x9feb8f50.
//
// Solidity: function addReward(address _relayer, uint256 reward) returns()
func (_HubLA *HubLATransactorSession) AddReward(_relayer common.Address, reward *big.Int) (*types.Transaction, error) {
	return _HubLA.Contract.AddReward(&_HubLA.TransactOpts, _relayer, reward)
}

// AlreadyExecuted is a paid mutator transaction binding the contract method 0x0f03ee59.
//
// Solidity: function alreadyExecuted(address _relayer, uint256 swapValue) returns()
func (_HubLA *HubLATransactor) AlreadyExecuted(opts *bind.TransactOpts, _relayer common.Address, swapValue *big.Int) (*types.Transaction, error) {
	return _HubLA.contract.Transact(opts, "alreadyExecuted", _relayer, swapValue)
}

// AlreadyExecuted is a paid mutator transaction binding the contract method 0x0f03ee59.
//
// Solidity: function alreadyExecuted(address _relayer, uint256 swapValue) returns()
func (_HubLA *HubLASession) AlreadyExecuted(_relayer common.Address, swapValue *big.Int) (*types.Transaction, error) {
	return _HubLA.Contract.AlreadyExecuted(&_HubLA.TransactOpts, _relayer, swapValue)
}

// AlreadyExecuted is a paid mutator transaction binding the contract method 0x0f03ee59.
//
// Solidity: function alreadyExecuted(address _relayer, uint256 swapValue) returns()
func (_HubLA *HubLATransactorSession) AlreadyExecuted(_relayer common.Address, swapValue *big.Int) (*types.Transaction, error) {
	return _HubLA.Contract.AlreadyExecuted(&_HubLA.TransactOpts, _relayer, swapValue)
}

// AlreadyVoted is a paid mutator transaction binding the contract method 0xe499dacf.
//
// Solidity: function alreadyVoted(address _relayer, uint256 swapValue) returns()
func (_HubLA *HubLATransactor) AlreadyVoted(opts *bind.TransactOpts, _relayer common.Address, swapValue *big.Int) (*types.Transaction, error) {
	return _HubLA.contract.Transact(opts, "alreadyVoted", _relayer, swapValue)
}

// AlreadyVoted is a paid mutator transaction binding the contract method 0xe499dacf.
//
// Solidity: function alreadyVoted(address _relayer, uint256 swapValue) returns()
func (_HubLA *HubLASession) AlreadyVoted(_relayer common.Address, swapValue *big.Int) (*types.Transaction, error) {
	return _HubLA.Contract.AlreadyVoted(&_HubLA.TransactOpts, _relayer, swapValue)
}

// AlreadyVoted is a paid mutator transaction binding the contract method 0xe499dacf.
//
// Solidity: function alreadyVoted(address _relayer, uint256 swapValue) returns()
func (_HubLA *HubLATransactorSession) AlreadyVoted(_relayer common.Address, swapValue *big.Int) (*types.Transaction, error) {
	return _HubLA.Contract.AlreadyVoted(&_HubLA.TransactOpts, _relayer, swapValue)
}

// Felony is a paid mutator transaction binding the contract method 0x35409f7f.
//
// Solidity: function felony(address _relayer) returns()
func (_HubLA *HubLATransactor) Felony(opts *bind.TransactOpts, _relayer common.Address) (*types.Transaction, error) {
	return _HubLA.contract.Transact(opts, "felony", _relayer)
}

// Felony is a paid mutator transaction binding the contract method 0x35409f7f.
//
// Solidity: function felony(address _relayer) returns()
func (_HubLA *HubLASession) Felony(_relayer common.Address) (*types.Transaction, error) {
	return _HubLA.Contract.Felony(&_HubLA.TransactOpts, _relayer)
}

// Felony is a paid mutator transaction binding the contract method 0x35409f7f.
//
// Solidity: function felony(address _relayer) returns()
func (_HubLA *HubLATransactorSession) Felony(_relayer common.Address) (*types.Transaction, error) {
	return _HubLA.Contract.Felony(&_HubLA.TransactOpts, _relayer)
}

// Init is a paid mutator transaction binding the contract method 0x19ab453c.
//
// Solidity: function init(address initBridgeAddress) returns()
func (_HubLA *HubLATransactor) Init(opts *bind.TransactOpts, initBridgeAddress common.Address) (*types.Transaction, error) {
	return _HubLA.contract.Transact(opts, "init", initBridgeAddress)
}

// Init is a paid mutator transaction binding the contract method 0x19ab453c.
//
// Solidity: function init(address initBridgeAddress) returns()
func (_HubLA *HubLASession) Init(initBridgeAddress common.Address) (*types.Transaction, error) {
	return _HubLA.Contract.Init(&_HubLA.TransactOpts, initBridgeAddress)
}

// Init is a paid mutator transaction binding the contract method 0x19ab453c.
//
// Solidity: function init(address initBridgeAddress) returns()
func (_HubLA *HubLATransactorSession) Init(initBridgeAddress common.Address) (*types.Transaction, error) {
	return _HubLA.Contract.Init(&_HubLA.TransactOpts, initBridgeAddress)
}

// Register is a paid mutator transaction binding the contract method 0x1aa3a008.
//
// Solidity: function register() payable returns()
func (_HubLA *HubLATransactor) Register(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HubLA.contract.Transact(opts, "register")
}

// Register is a paid mutator transaction binding the contract method 0x1aa3a008.
//
// Solidity: function register() payable returns()
func (_HubLA *HubLASession) Register() (*types.Transaction, error) {
	return _HubLA.Contract.Register(&_HubLA.TransactOpts)
}

// Register is a paid mutator transaction binding the contract method 0x1aa3a008.
//
// Solidity: function register() payable returns()
func (_HubLA *HubLATransactorSession) Register() (*types.Transaction, error) {
	return _HubLA.Contract.Register(&_HubLA.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_HubLA *HubLATransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HubLA.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_HubLA *HubLASession) RenounceOwnership() (*types.Transaction, error) {
	return _HubLA.Contract.RenounceOwnership(&_HubLA.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_HubLA *HubLATransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _HubLA.Contract.RenounceOwnership(&_HubLA.TransactOpts)
}

// Slash is a paid mutator transaction binding the contract method 0x02fb4d85.
//
// Solidity: function slash(address _relayer, uint256 penalty) returns()
func (_HubLA *HubLATransactor) Slash(opts *bind.TransactOpts, _relayer common.Address, penalty *big.Int) (*types.Transaction, error) {
	return _HubLA.contract.Transact(opts, "slash", _relayer, penalty)
}

// Slash is a paid mutator transaction binding the contract method 0x02fb4d85.
//
// Solidity: function slash(address _relayer, uint256 penalty) returns()
func (_HubLA *HubLASession) Slash(_relayer common.Address, penalty *big.Int) (*types.Transaction, error) {
	return _HubLA.Contract.Slash(&_HubLA.TransactOpts, _relayer, penalty)
}

// Slash is a paid mutator transaction binding the contract method 0x02fb4d85.
//
// Solidity: function slash(address _relayer, uint256 penalty) returns()
func (_HubLA *HubLATransactorSession) Slash(_relayer common.Address, penalty *big.Int) (*types.Transaction, error) {
	return _HubLA.Contract.Slash(&_HubLA.TransactOpts, _relayer, penalty)
}

// SwapReward is a paid mutator transaction binding the contract method 0xadaa71cc.
//
// Solidity: function swapReward(address _relayer, uint256 swapValue) returns()
func (_HubLA *HubLATransactor) SwapReward(opts *bind.TransactOpts, _relayer common.Address, swapValue *big.Int) (*types.Transaction, error) {
	return _HubLA.contract.Transact(opts, "swapReward", _relayer, swapValue)
}

// SwapReward is a paid mutator transaction binding the contract method 0xadaa71cc.
//
// Solidity: function swapReward(address _relayer, uint256 swapValue) returns()
func (_HubLA *HubLASession) SwapReward(_relayer common.Address, swapValue *big.Int) (*types.Transaction, error) {
	return _HubLA.Contract.SwapReward(&_HubLA.TransactOpts, _relayer, swapValue)
}

// SwapReward is a paid mutator transaction binding the contract method 0xadaa71cc.
//
// Solidity: function swapReward(address _relayer, uint256 swapValue) returns()
func (_HubLA *HubLATransactorSession) SwapReward(_relayer common.Address, swapValue *big.Int) (*types.Transaction, error) {
	return _HubLA.Contract.SwapReward(&_HubLA.TransactOpts, _relayer, swapValue)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_HubLA *HubLATransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _HubLA.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_HubLA *HubLASession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _HubLA.Contract.TransferOwnership(&_HubLA.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_HubLA *HubLATransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _HubLA.Contract.TransferOwnership(&_HubLA.TransactOpts, newOwner)
}

// Unregister is a paid mutator transaction binding the contract method 0xe79a198f.
//
// Solidity: function unregister() returns()
func (_HubLA *HubLATransactor) Unregister(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HubLA.contract.Transact(opts, "unregister")
}

// Unregister is a paid mutator transaction binding the contract method 0xe79a198f.
//
// Solidity: function unregister() returns()
func (_HubLA *HubLASession) Unregister() (*types.Transaction, error) {
	return _HubLA.Contract.Unregister(&_HubLA.TransactOpts)
}

// Unregister is a paid mutator transaction binding the contract method 0xe79a198f.
//
// Solidity: function unregister() returns()
func (_HubLA *HubLATransactorSession) Unregister() (*types.Transaction, error) {
	return _HubLA.Contract.Unregister(&_HubLA.TransactOpts)
}

// UpdatePenaltyPercentage is a paid mutator transaction binding the contract method 0xeb961693.
//
// Solidity: function updatePenaltyPercentage(uint256 penaltyPercentage) returns()
func (_HubLA *HubLATransactor) UpdatePenaltyPercentage(opts *bind.TransactOpts, penaltyPercentage *big.Int) (*types.Transaction, error) {
	return _HubLA.contract.Transact(opts, "updatePenaltyPercentage", penaltyPercentage)
}

// UpdatePenaltyPercentage is a paid mutator transaction binding the contract method 0xeb961693.
//
// Solidity: function updatePenaltyPercentage(uint256 penaltyPercentage) returns()
func (_HubLA *HubLASession) UpdatePenaltyPercentage(penaltyPercentage *big.Int) (*types.Transaction, error) {
	return _HubLA.Contract.UpdatePenaltyPercentage(&_HubLA.TransactOpts, penaltyPercentage)
}

// UpdatePenaltyPercentage is a paid mutator transaction binding the contract method 0xeb961693.
//
// Solidity: function updatePenaltyPercentage(uint256 penaltyPercentage) returns()
func (_HubLA *HubLATransactorSession) UpdatePenaltyPercentage(penaltyPercentage *big.Int) (*types.Transaction, error) {
	return _HubLA.Contract.UpdatePenaltyPercentage(&_HubLA.TransactOpts, penaltyPercentage)
}

// UpdateReward is a paid mutator transaction binding the contract method 0x1a399125.
//
// Solidity: function updateReward(address _relayer, uint256 reward) returns()
func (_HubLA *HubLATransactor) UpdateReward(opts *bind.TransactOpts, _relayer common.Address, reward *big.Int) (*types.Transaction, error) {
	return _HubLA.contract.Transact(opts, "updateReward", _relayer, reward)
}

// UpdateReward is a paid mutator transaction binding the contract method 0x1a399125.
//
// Solidity: function updateReward(address _relayer, uint256 reward) returns()
func (_HubLA *HubLASession) UpdateReward(_relayer common.Address, reward *big.Int) (*types.Transaction, error) {
	return _HubLA.Contract.UpdateReward(&_HubLA.TransactOpts, _relayer, reward)
}

// UpdateReward is a paid mutator transaction binding the contract method 0x1a399125.
//
// Solidity: function updateReward(address _relayer, uint256 reward) returns()
func (_HubLA *HubLATransactorSession) UpdateReward(_relayer common.Address, reward *big.Int) (*types.Transaction, error) {
	return _HubLA.Contract.UpdateReward(&_HubLA.TransactOpts, _relayer, reward)
}

// UpdateRewardPercentage is a paid mutator transaction binding the contract method 0xa4f209b0.
//
// Solidity: function updateRewardPercentage(uint256 rewardPercent) returns()
func (_HubLA *HubLATransactor) UpdateRewardPercentage(opts *bind.TransactOpts, rewardPercent *big.Int) (*types.Transaction, error) {
	return _HubLA.contract.Transact(opts, "updateRewardPercentage", rewardPercent)
}

// UpdateRewardPercentage is a paid mutator transaction binding the contract method 0xa4f209b0.
//
// Solidity: function updateRewardPercentage(uint256 rewardPercent) returns()
func (_HubLA *HubLASession) UpdateRewardPercentage(rewardPercent *big.Int) (*types.Transaction, error) {
	return _HubLA.Contract.UpdateRewardPercentage(&_HubLA.TransactOpts, rewardPercent)
}

// UpdateRewardPercentage is a paid mutator transaction binding the contract method 0xa4f209b0.
//
// Solidity: function updateRewardPercentage(uint256 rewardPercent) returns()
func (_HubLA *HubLATransactorSession) UpdateRewardPercentage(rewardPercent *big.Int) (*types.Transaction, error) {
	return _HubLA.Contract.UpdateRewardPercentage(&_HubLA.TransactOpts, rewardPercent)
}

// UpdateSystemRewardAddr is a paid mutator transaction binding the contract method 0x4a910b85.
//
// Solidity: function updateSystemRewardAddr(address newAddr) returns()
func (_HubLA *HubLATransactor) UpdateSystemRewardAddr(opts *bind.TransactOpts, newAddr common.Address) (*types.Transaction, error) {
	return _HubLA.contract.Transact(opts, "updateSystemRewardAddr", newAddr)
}

// UpdateSystemRewardAddr is a paid mutator transaction binding the contract method 0x4a910b85.
//
// Solidity: function updateSystemRewardAddr(address newAddr) returns()
func (_HubLA *HubLASession) UpdateSystemRewardAddr(newAddr common.Address) (*types.Transaction, error) {
	return _HubLA.Contract.UpdateSystemRewardAddr(&_HubLA.TransactOpts, newAddr)
}

// UpdateSystemRewardAddr is a paid mutator transaction binding the contract method 0x4a910b85.
//
// Solidity: function updateSystemRewardAddr(address newAddr) returns()
func (_HubLA *HubLATransactorSession) UpdateSystemRewardAddr(newAddr common.Address) (*types.Transaction, error) {
	return _HubLA.Contract.UpdateSystemRewardAddr(&_HubLA.TransactOpts, newAddr)
}

// HubLAOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the HubLA contract.
type HubLAOwnershipTransferredIterator struct {
	Event *HubLAOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *HubLAOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HubLAOwnershipTransferred)
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
		it.Event = new(HubLAOwnershipTransferred)
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
func (it *HubLAOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HubLAOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HubLAOwnershipTransferred represents a OwnershipTransferred event raised by the HubLA contract.
type HubLAOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_HubLA *HubLAFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*HubLAOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _HubLA.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &HubLAOwnershipTransferredIterator{contract: _HubLA.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_HubLA *HubLAFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *HubLAOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _HubLA.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HubLAOwnershipTransferred)
				if err := _HubLA.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_HubLA *HubLAFilterer) ParseOwnershipTransferred(log types.Log) (*HubLAOwnershipTransferred, error) {
	event := new(HubLAOwnershipTransferred)
	if err := _HubLA.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HubLAPenaltyForRelayerIterator is returned from FilterPenaltyForRelayer and is used to iterate over the raw logs and unpacked data for PenaltyForRelayer events raised by the HubLA contract.
type HubLAPenaltyForRelayerIterator struct {
	Event *HubLAPenaltyForRelayer // Event containing the contract specifics and raw log

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
func (it *HubLAPenaltyForRelayerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HubLAPenaltyForRelayer)
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
		it.Event = new(HubLAPenaltyForRelayer)
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
func (it *HubLAPenaltyForRelayerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HubLAPenaltyForRelayerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HubLAPenaltyForRelayer represents a PenaltyForRelayer event raised by the HubLA contract.
type HubLAPenaltyForRelayer struct {
	Reason  string
	Relayer common.Address
	Penalty *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPenaltyForRelayer is a free log retrieval operation binding the contract event 0x73f262bdf0c3145b25a03d5c75d5989bdab8ecb77968d7ff6de8a3bd83b8b13b.
//
// Solidity: event PenaltyForRelayer(string reason, address relayer, uint256 penalty)
func (_HubLA *HubLAFilterer) FilterPenaltyForRelayer(opts *bind.FilterOpts) (*HubLAPenaltyForRelayerIterator, error) {

	logs, sub, err := _HubLA.contract.FilterLogs(opts, "PenaltyForRelayer")
	if err != nil {
		return nil, err
	}
	return &HubLAPenaltyForRelayerIterator{contract: _HubLA.contract, event: "PenaltyForRelayer", logs: logs, sub: sub}, nil
}

// WatchPenaltyForRelayer is a free log subscription operation binding the contract event 0x73f262bdf0c3145b25a03d5c75d5989bdab8ecb77968d7ff6de8a3bd83b8b13b.
//
// Solidity: event PenaltyForRelayer(string reason, address relayer, uint256 penalty)
func (_HubLA *HubLAFilterer) WatchPenaltyForRelayer(opts *bind.WatchOpts, sink chan<- *HubLAPenaltyForRelayer) (event.Subscription, error) {

	logs, sub, err := _HubLA.contract.WatchLogs(opts, "PenaltyForRelayer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HubLAPenaltyForRelayer)
				if err := _HubLA.contract.UnpackLog(event, "PenaltyForRelayer", log); err != nil {
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

// ParsePenaltyForRelayer is a log parse operation binding the contract event 0x73f262bdf0c3145b25a03d5c75d5989bdab8ecb77968d7ff6de8a3bd83b8b13b.
//
// Solidity: event PenaltyForRelayer(string reason, address relayer, uint256 penalty)
func (_HubLA *HubLAFilterer) ParsePenaltyForRelayer(log types.Log) (*HubLAPenaltyForRelayer, error) {
	event := new(HubLAPenaltyForRelayer)
	if err := _HubLA.contract.UnpackLog(event, "PenaltyForRelayer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HubLAPenaltyPercentChangedIterator is returned from FilterPenaltyPercentChanged and is used to iterate over the raw logs and unpacked data for PenaltyPercentChanged events raised by the HubLA contract.
type HubLAPenaltyPercentChangedIterator struct {
	Event *HubLAPenaltyPercentChanged // Event containing the contract specifics and raw log

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
func (it *HubLAPenaltyPercentChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HubLAPenaltyPercentChanged)
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
		it.Event = new(HubLAPenaltyPercentChanged)
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
func (it *HubLAPenaltyPercentChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HubLAPenaltyPercentChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HubLAPenaltyPercentChanged represents a PenaltyPercentChanged event raised by the HubLA contract.
type HubLAPenaltyPercentChanged struct {
	NewPercent *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPenaltyPercentChanged is a free log retrieval operation binding the contract event 0xa7422d81f8725b29775f8eaafb874b2d997bb6e6d2a60eda98f5fb2eaeb58f3e.
//
// Solidity: event PenaltyPercentChanged(uint256 newPercent)
func (_HubLA *HubLAFilterer) FilterPenaltyPercentChanged(opts *bind.FilterOpts) (*HubLAPenaltyPercentChangedIterator, error) {

	logs, sub, err := _HubLA.contract.FilterLogs(opts, "PenaltyPercentChanged")
	if err != nil {
		return nil, err
	}
	return &HubLAPenaltyPercentChangedIterator{contract: _HubLA.contract, event: "PenaltyPercentChanged", logs: logs, sub: sub}, nil
}

// WatchPenaltyPercentChanged is a free log subscription operation binding the contract event 0xa7422d81f8725b29775f8eaafb874b2d997bb6e6d2a60eda98f5fb2eaeb58f3e.
//
// Solidity: event PenaltyPercentChanged(uint256 newPercent)
func (_HubLA *HubLAFilterer) WatchPenaltyPercentChanged(opts *bind.WatchOpts, sink chan<- *HubLAPenaltyPercentChanged) (event.Subscription, error) {

	logs, sub, err := _HubLA.contract.WatchLogs(opts, "PenaltyPercentChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HubLAPenaltyPercentChanged)
				if err := _HubLA.contract.UnpackLog(event, "PenaltyPercentChanged", log); err != nil {
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

// ParsePenaltyPercentChanged is a log parse operation binding the contract event 0xa7422d81f8725b29775f8eaafb874b2d997bb6e6d2a60eda98f5fb2eaeb58f3e.
//
// Solidity: event PenaltyPercentChanged(uint256 newPercent)
func (_HubLA *HubLAFilterer) ParsePenaltyPercentChanged(log types.Log) (*HubLAPenaltyPercentChanged, error) {
	event := new(HubLAPenaltyPercentChanged)
	if err := _HubLA.contract.UnpackLog(event, "PenaltyPercentChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HubLARelayerBlockedIterator is returned from FilterRelayerBlocked and is used to iterate over the raw logs and unpacked data for RelayerBlocked events raised by the HubLA contract.
type HubLARelayerBlockedIterator struct {
	Event *HubLARelayerBlocked // Event containing the contract specifics and raw log

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
func (it *HubLARelayerBlockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HubLARelayerBlocked)
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
		it.Event = new(HubLARelayerBlocked)
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
func (it *HubLARelayerBlockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HubLARelayerBlockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HubLARelayerBlocked represents a RelayerBlocked event raised by the HubLA contract.
type HubLARelayerBlocked struct {
	Relayer common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRelayerBlocked is a free log retrieval operation binding the contract event 0xa5ea068b2ed8e05403d639cea236649975a9712487c612f9a6945b6bad00d81d.
//
// Solidity: event RelayerBlocked(address relayer)
func (_HubLA *HubLAFilterer) FilterRelayerBlocked(opts *bind.FilterOpts) (*HubLARelayerBlockedIterator, error) {

	logs, sub, err := _HubLA.contract.FilterLogs(opts, "RelayerBlocked")
	if err != nil {
		return nil, err
	}
	return &HubLARelayerBlockedIterator{contract: _HubLA.contract, event: "RelayerBlocked", logs: logs, sub: sub}, nil
}

// WatchRelayerBlocked is a free log subscription operation binding the contract event 0xa5ea068b2ed8e05403d639cea236649975a9712487c612f9a6945b6bad00d81d.
//
// Solidity: event RelayerBlocked(address relayer)
func (_HubLA *HubLAFilterer) WatchRelayerBlocked(opts *bind.WatchOpts, sink chan<- *HubLARelayerBlocked) (event.Subscription, error) {

	logs, sub, err := _HubLA.contract.WatchLogs(opts, "RelayerBlocked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HubLARelayerBlocked)
				if err := _HubLA.contract.UnpackLog(event, "RelayerBlocked", log); err != nil {
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

// ParseRelayerBlocked is a log parse operation binding the contract event 0xa5ea068b2ed8e05403d639cea236649975a9712487c612f9a6945b6bad00d81d.
//
// Solidity: event RelayerBlocked(address relayer)
func (_HubLA *HubLAFilterer) ParseRelayerBlocked(log types.Log) (*HubLARelayerBlocked, error) {
	event := new(HubLARelayerBlocked)
	if err := _HubLA.contract.UnpackLog(event, "RelayerBlocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HubLARelayerRegisterIterator is returned from FilterRelayerRegister and is used to iterate over the raw logs and unpacked data for RelayerRegister events raised by the HubLA contract.
type HubLARelayerRegisterIterator struct {
	Event *HubLARelayerRegister // Event containing the contract specifics and raw log

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
func (it *HubLARelayerRegisterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HubLARelayerRegister)
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
		it.Event = new(HubLARelayerRegister)
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
func (it *HubLARelayerRegisterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HubLARelayerRegisterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HubLARelayerRegister represents a RelayerRegister event raised by the HubLA contract.
type HubLARelayerRegister struct {
	Relayer common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRelayerRegister is a free log retrieval operation binding the contract event 0x03f64c5eb6e7ea7290123c03dbf1d28d676a583815f5f6a8ebd7f153cfb45215.
//
// Solidity: event RelayerRegister(address _relayer)
func (_HubLA *HubLAFilterer) FilterRelayerRegister(opts *bind.FilterOpts) (*HubLARelayerRegisterIterator, error) {

	logs, sub, err := _HubLA.contract.FilterLogs(opts, "RelayerRegister")
	if err != nil {
		return nil, err
	}
	return &HubLARelayerRegisterIterator{contract: _HubLA.contract, event: "RelayerRegister", logs: logs, sub: sub}, nil
}

// WatchRelayerRegister is a free log subscription operation binding the contract event 0x03f64c5eb6e7ea7290123c03dbf1d28d676a583815f5f6a8ebd7f153cfb45215.
//
// Solidity: event RelayerRegister(address _relayer)
func (_HubLA *HubLAFilterer) WatchRelayerRegister(opts *bind.WatchOpts, sink chan<- *HubLARelayerRegister) (event.Subscription, error) {

	logs, sub, err := _HubLA.contract.WatchLogs(opts, "RelayerRegister")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HubLARelayerRegister)
				if err := _HubLA.contract.UnpackLog(event, "RelayerRegister", log); err != nil {
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

// ParseRelayerRegister is a log parse operation binding the contract event 0x03f64c5eb6e7ea7290123c03dbf1d28d676a583815f5f6a8ebd7f153cfb45215.
//
// Solidity: event RelayerRegister(address _relayer)
func (_HubLA *HubLAFilterer) ParseRelayerRegister(log types.Log) (*HubLARelayerRegister, error) {
	event := new(HubLARelayerRegister)
	if err := _HubLA.contract.UnpackLog(event, "RelayerRegister", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HubLARelayerUnRegisterIterator is returned from FilterRelayerUnRegister and is used to iterate over the raw logs and unpacked data for RelayerUnRegister events raised by the HubLA contract.
type HubLARelayerUnRegisterIterator struct {
	Event *HubLARelayerUnRegister // Event containing the contract specifics and raw log

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
func (it *HubLARelayerUnRegisterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HubLARelayerUnRegister)
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
		it.Event = new(HubLARelayerUnRegister)
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
func (it *HubLARelayerUnRegisterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HubLARelayerUnRegisterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HubLARelayerUnRegister represents a RelayerUnRegister event raised by the HubLA contract.
type HubLARelayerUnRegister struct {
	Relayer common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRelayerUnRegister is a free log retrieval operation binding the contract event 0x013c40a2b0c93c7ba3e8b43e7529d1d6f020670e66817894c8f86708afb3e62c.
//
// Solidity: event RelayerUnRegister(address _relayer)
func (_HubLA *HubLAFilterer) FilterRelayerUnRegister(opts *bind.FilterOpts) (*HubLARelayerUnRegisterIterator, error) {

	logs, sub, err := _HubLA.contract.FilterLogs(opts, "RelayerUnRegister")
	if err != nil {
		return nil, err
	}
	return &HubLARelayerUnRegisterIterator{contract: _HubLA.contract, event: "RelayerUnRegister", logs: logs, sub: sub}, nil
}

// WatchRelayerUnRegister is a free log subscription operation binding the contract event 0x013c40a2b0c93c7ba3e8b43e7529d1d6f020670e66817894c8f86708afb3e62c.
//
// Solidity: event RelayerUnRegister(address _relayer)
func (_HubLA *HubLAFilterer) WatchRelayerUnRegister(opts *bind.WatchOpts, sink chan<- *HubLARelayerUnRegister) (event.Subscription, error) {

	logs, sub, err := _HubLA.contract.WatchLogs(opts, "RelayerUnRegister")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HubLARelayerUnRegister)
				if err := _HubLA.contract.UnpackLog(event, "RelayerUnRegister", log); err != nil {
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

// ParseRelayerUnRegister is a log parse operation binding the contract event 0x013c40a2b0c93c7ba3e8b43e7529d1d6f020670e66817894c8f86708afb3e62c.
//
// Solidity: event RelayerUnRegister(address _relayer)
func (_HubLA *HubLAFilterer) ParseRelayerUnRegister(log types.Log) (*HubLARelayerUnRegister, error) {
	event := new(HubLARelayerUnRegister)
	if err := _HubLA.contract.UnpackLog(event, "RelayerUnRegister", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HubLARewardForRelayerIterator is returned from FilterRewardForRelayer and is used to iterate over the raw logs and unpacked data for RewardForRelayer events raised by the HubLA contract.
type HubLARewardForRelayerIterator struct {
	Event *HubLARewardForRelayer // Event containing the contract specifics and raw log

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
func (it *HubLARewardForRelayerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HubLARewardForRelayer)
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
		it.Event = new(HubLARewardForRelayer)
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
func (it *HubLARewardForRelayerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HubLARewardForRelayerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HubLARewardForRelayer represents a RewardForRelayer event raised by the HubLA contract.
type HubLARewardForRelayer struct {
	Relayer common.Address
	Reward  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRewardForRelayer is a free log retrieval operation binding the contract event 0x63dd553807ea27cb996173133b29a47f29c6b3bec34043271f5a5f7dd1f949c9.
//
// Solidity: event RewardForRelayer(address relayer, uint256 reward)
func (_HubLA *HubLAFilterer) FilterRewardForRelayer(opts *bind.FilterOpts) (*HubLARewardForRelayerIterator, error) {

	logs, sub, err := _HubLA.contract.FilterLogs(opts, "RewardForRelayer")
	if err != nil {
		return nil, err
	}
	return &HubLARewardForRelayerIterator{contract: _HubLA.contract, event: "RewardForRelayer", logs: logs, sub: sub}, nil
}

// WatchRewardForRelayer is a free log subscription operation binding the contract event 0x63dd553807ea27cb996173133b29a47f29c6b3bec34043271f5a5f7dd1f949c9.
//
// Solidity: event RewardForRelayer(address relayer, uint256 reward)
func (_HubLA *HubLAFilterer) WatchRewardForRelayer(opts *bind.WatchOpts, sink chan<- *HubLARewardForRelayer) (event.Subscription, error) {

	logs, sub, err := _HubLA.contract.WatchLogs(opts, "RewardForRelayer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HubLARewardForRelayer)
				if err := _HubLA.contract.UnpackLog(event, "RewardForRelayer", log); err != nil {
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

// ParseRewardForRelayer is a log parse operation binding the contract event 0x63dd553807ea27cb996173133b29a47f29c6b3bec34043271f5a5f7dd1f949c9.
//
// Solidity: event RewardForRelayer(address relayer, uint256 reward)
func (_HubLA *HubLAFilterer) ParseRewardForRelayer(log types.Log) (*HubLARewardForRelayer, error) {
	event := new(HubLARewardForRelayer)
	if err := _HubLA.contract.UnpackLog(event, "RewardForRelayer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HubLARewardPercentageChangedIterator is returned from FilterRewardPercentageChanged and is used to iterate over the raw logs and unpacked data for RewardPercentageChanged events raised by the HubLA contract.
type HubLARewardPercentageChangedIterator struct {
	Event *HubLARewardPercentageChanged // Event containing the contract specifics and raw log

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
func (it *HubLARewardPercentageChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HubLARewardPercentageChanged)
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
		it.Event = new(HubLARewardPercentageChanged)
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
func (it *HubLARewardPercentageChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HubLARewardPercentageChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HubLARewardPercentageChanged represents a RewardPercentageChanged event raised by the HubLA contract.
type HubLARewardPercentageChanged struct {
	NewPercentage *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRewardPercentageChanged is a free log retrieval operation binding the contract event 0x23ad5751be9629264e3725ea971a625cff4033b5a3df9f4526e8a32159c0fa89.
//
// Solidity: event RewardPercentageChanged(uint256 newPercentage)
func (_HubLA *HubLAFilterer) FilterRewardPercentageChanged(opts *bind.FilterOpts) (*HubLARewardPercentageChangedIterator, error) {

	logs, sub, err := _HubLA.contract.FilterLogs(opts, "RewardPercentageChanged")
	if err != nil {
		return nil, err
	}
	return &HubLARewardPercentageChangedIterator{contract: _HubLA.contract, event: "RewardPercentageChanged", logs: logs, sub: sub}, nil
}

// WatchRewardPercentageChanged is a free log subscription operation binding the contract event 0x23ad5751be9629264e3725ea971a625cff4033b5a3df9f4526e8a32159c0fa89.
//
// Solidity: event RewardPercentageChanged(uint256 newPercentage)
func (_HubLA *HubLAFilterer) WatchRewardPercentageChanged(opts *bind.WatchOpts, sink chan<- *HubLARewardPercentageChanged) (event.Subscription, error) {

	logs, sub, err := _HubLA.contract.WatchLogs(opts, "RewardPercentageChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HubLARewardPercentageChanged)
				if err := _HubLA.contract.UnpackLog(event, "RewardPercentageChanged", log); err != nil {
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

// ParseRewardPercentageChanged is a log parse operation binding the contract event 0x23ad5751be9629264e3725ea971a625cff4033b5a3df9f4526e8a32159c0fa89.
//
// Solidity: event RewardPercentageChanged(uint256 newPercentage)
func (_HubLA *HubLAFilterer) ParseRewardPercentageChanged(log types.Log) (*HubLARewardPercentageChanged, error) {
	event := new(HubLARewardPercentageChanged)
	if err := _HubLA.contract.UnpackLog(event, "RewardPercentageChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HubLASystemRewardAddressChangedIterator is returned from FilterSystemRewardAddressChanged and is used to iterate over the raw logs and unpacked data for SystemRewardAddressChanged events raised by the HubLA contract.
type HubLASystemRewardAddressChangedIterator struct {
	Event *HubLASystemRewardAddressChanged // Event containing the contract specifics and raw log

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
func (it *HubLASystemRewardAddressChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HubLASystemRewardAddressChanged)
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
		it.Event = new(HubLASystemRewardAddressChanged)
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
func (it *HubLASystemRewardAddressChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HubLASystemRewardAddressChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HubLASystemRewardAddressChanged represents a SystemRewardAddressChanged event raised by the HubLA contract.
type HubLASystemRewardAddressChanged struct {
	NewAddr common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSystemRewardAddressChanged is a free log retrieval operation binding the contract event 0x422c8cd0e76296d4cdfe4d22844f444eb533e74612cf59490295221860ad6441.
//
// Solidity: event SystemRewardAddressChanged(address newAddr)
func (_HubLA *HubLAFilterer) FilterSystemRewardAddressChanged(opts *bind.FilterOpts) (*HubLASystemRewardAddressChangedIterator, error) {

	logs, sub, err := _HubLA.contract.FilterLogs(opts, "SystemRewardAddressChanged")
	if err != nil {
		return nil, err
	}
	return &HubLASystemRewardAddressChangedIterator{contract: _HubLA.contract, event: "SystemRewardAddressChanged", logs: logs, sub: sub}, nil
}

// WatchSystemRewardAddressChanged is a free log subscription operation binding the contract event 0x422c8cd0e76296d4cdfe4d22844f444eb533e74612cf59490295221860ad6441.
//
// Solidity: event SystemRewardAddressChanged(address newAddr)
func (_HubLA *HubLAFilterer) WatchSystemRewardAddressChanged(opts *bind.WatchOpts, sink chan<- *HubLASystemRewardAddressChanged) (event.Subscription, error) {

	logs, sub, err := _HubLA.contract.WatchLogs(opts, "SystemRewardAddressChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HubLASystemRewardAddressChanged)
				if err := _HubLA.contract.UnpackLog(event, "SystemRewardAddressChanged", log); err != nil {
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

// ParseSystemRewardAddressChanged is a log parse operation binding the contract event 0x422c8cd0e76296d4cdfe4d22844f444eb533e74612cf59490295221860ad6441.
//
// Solidity: event SystemRewardAddressChanged(address newAddr)
func (_HubLA *HubLAFilterer) ParseSystemRewardAddressChanged(log types.Log) (*HubLASystemRewardAddressChanged, error) {
	event := new(HubLASystemRewardAddressChanged)
	if err := _HubLA.contract.UnpackLog(event, "SystemRewardAddressChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
