// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package hubEth

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

// HubEthMetaData contains all meta data concerning the HubEth contract.
var HubEthMetaData = &bind.MetaData{
	ABI: "[{\"name\":\"owner\",\"type\":\"function\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"name\":\"renounceOwnership\",\"type\":\"function\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"name\":\"transferOwnership\",\"type\":\"function\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"name\":\"_relayerBlocked\",\"type\":\"function\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"name\":\"_penaltyPercentage\",\"type\":\"function\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"name\":\"_rewardPercentage\",\"type\":\"function\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"name\":\"_relayersExistMap\",\"type\":\"function\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"constructor\",\"inputs\":[{\"name\":\"bridgeAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"penaltyPercentage\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"rewardPercentage\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"backendSrvAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"name\":\"register\",\"type\":\"function\",\"inputs\":[{\"name\":\"relayerToAdd\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"name\":\"unregister\",\"type\":\"function\",\"inputs\":[{\"name\":\"relayerToRemove\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"name\":\"isRelayer\",\"type\":\"function\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"name\":\"alreadyExecuted\",\"type\":\"function\",\"inputs\":[{\"name\":\"_relayer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"swapValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"name\":\"felony\",\"type\":\"function\",\"inputs\":[{\"name\":\"relayer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"name\":\"updatePenaltyPercentage\",\"type\":\"function\",\"inputs\":[{\"name\":\"penaltyPercentage\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"name\":\"swapReward\",\"type\":\"function\",\"inputs\":[{\"name\":\"_relayer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"swapValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"name\":\"updateRewardPercentage\",\"type\":\"function\",\"inputs\":[{\"name\":\"rewardPercent\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"name\":\"OwnershipTransferred\",\"type\":\"event\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"internalType\":\"address\",\"indexed\":true},{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\",\"indexed\":true}],\"anonymous\":false},{\"name\":\"RelayerBlocked\",\"type\":\"event\",\"inputs\":[{\"name\":\"relayer\",\"type\":\"address\",\"internalType\":\"address\",\"indexed\":true}],\"anonymous\":false},{\"name\":\"PenaltyPercentChanged\",\"type\":\"event\",\"inputs\":[{\"name\":\"newPercentage\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"anonymous\":false},{\"name\":\"RewardForRelayer\",\"type\":\"event\",\"inputs\":[{\"name\":\"relayer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reward\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"anonymous\":false},{\"name\":\"RewardPercentageChanged\",\"type\":\"event\",\"inputs\":[{\"name\":\"newPercentage\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"anonymous\":false},{\"name\":\"RelayerRegister\",\"type\":\"event\",\"inputs\":[{\"name\":\"_relayer\",\"type\":\"address\",\"internalType\":\"address\",\"indexed\":true}],\"anonymous\":false},{\"name\":\"RelayerUnRegister\",\"type\":\"event\",\"inputs\":[{\"name\":\"_relayer\",\"type\":\"address\",\"internalType\":\"address\",\"indexed\":true}],\"anonymous\":false},{\"name\":\"PenaltyForRelayer\",\"type\":\"event\",\"inputs\":[{\"name\":\"reason\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"relayer\",\"type\":\"address\",\"internalType\":\"address\",\"indexed\":true},{\"name\":\"penalty\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"anonymous\":false}]",
}

// HubEthABI is the input ABI used to generate the binding from.
// Deprecated: Use HubEthMetaData.ABI instead.
var HubEthABI = HubEthMetaData.ABI

// HubEth is an auto generated Go binding around an Ethereum contract.
type HubEth struct {
	HubEthCaller     // Read-only binding to the contract
	HubEthTransactor // Write-only binding to the contract
	HubEthFilterer   // Log filterer for contract events
}

// HubEthCaller is an auto generated read-only Go binding around an Ethereum contract.
type HubEthCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HubEthTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HubEthTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HubEthFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HubEthFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HubEthSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HubEthSession struct {
	Contract     *HubEth           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HubEthCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HubEthCallerSession struct {
	Contract *HubEthCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// HubEthTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HubEthTransactorSession struct {
	Contract     *HubEthTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HubEthRaw is an auto generated low-level Go binding around an Ethereum contract.
type HubEthRaw struct {
	Contract *HubEth // Generic contract binding to access the raw methods on
}

// HubEthCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HubEthCallerRaw struct {
	Contract *HubEthCaller // Generic read-only contract binding to access the raw methods on
}

// HubEthTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HubEthTransactorRaw struct {
	Contract *HubEthTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHubEth creates a new instance of HubEth, bound to a specific deployed contract.
func NewHubEth(address common.Address, backend bind.ContractBackend) (*HubEth, error) {
	contract, err := bindHubEth(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HubEth{HubEthCaller: HubEthCaller{contract: contract}, HubEthTransactor: HubEthTransactor{contract: contract}, HubEthFilterer: HubEthFilterer{contract: contract}}, nil
}

// NewHubEthCaller creates a new read-only instance of HubEth, bound to a specific deployed contract.
func NewHubEthCaller(address common.Address, caller bind.ContractCaller) (*HubEthCaller, error) {
	contract, err := bindHubEth(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HubEthCaller{contract: contract}, nil
}

// NewHubEthTransactor creates a new write-only instance of HubEth, bound to a specific deployed contract.
func NewHubEthTransactor(address common.Address, transactor bind.ContractTransactor) (*HubEthTransactor, error) {
	contract, err := bindHubEth(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HubEthTransactor{contract: contract}, nil
}

// NewHubEthFilterer creates a new log filterer instance of HubEth, bound to a specific deployed contract.
func NewHubEthFilterer(address common.Address, filterer bind.ContractFilterer) (*HubEthFilterer, error) {
	contract, err := bindHubEth(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HubEthFilterer{contract: contract}, nil
}

// bindHubEth binds a generic wrapper to an already deployed contract.
func bindHubEth(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HubEthABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HubEth *HubEthRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HubEth.Contract.HubEthCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HubEth *HubEthRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HubEth.Contract.HubEthTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HubEth *HubEthRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HubEth.Contract.HubEthTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HubEth *HubEthCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HubEth.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HubEth *HubEthTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HubEth.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HubEth *HubEthTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HubEth.Contract.contract.Transact(opts, method, params...)
}

// PenaltyPercentage is a free data retrieval call binding the contract method 0x93a55f92.
//
// Solidity: function _penaltyPercentage() view returns(uint256)
func (_HubEth *HubEthCaller) PenaltyPercentage(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _HubEth.contract.Call(opts, &out, "_penaltyPercentage")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PenaltyPercentage is a free data retrieval call binding the contract method 0x93a55f92.
//
// Solidity: function _penaltyPercentage() view returns(uint256)
func (_HubEth *HubEthSession) PenaltyPercentage() (*big.Int, error) {
	return _HubEth.Contract.PenaltyPercentage(&_HubEth.CallOpts)
}

// PenaltyPercentage is a free data retrieval call binding the contract method 0x93a55f92.
//
// Solidity: function _penaltyPercentage() view returns(uint256)
func (_HubEth *HubEthCallerSession) PenaltyPercentage() (*big.Int, error) {
	return _HubEth.Contract.PenaltyPercentage(&_HubEth.CallOpts)
}

// RelayerBlocked is a free data retrieval call binding the contract method 0x0537b4d6.
//
// Solidity: function _relayerBlocked(address ) view returns(bool)
func (_HubEth *HubEthCaller) RelayerBlocked(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _HubEth.contract.Call(opts, &out, "_relayerBlocked", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// RelayerBlocked is a free data retrieval call binding the contract method 0x0537b4d6.
//
// Solidity: function _relayerBlocked(address ) view returns(bool)
func (_HubEth *HubEthSession) RelayerBlocked(arg0 common.Address) (bool, error) {
	return _HubEth.Contract.RelayerBlocked(&_HubEth.CallOpts, arg0)
}

// RelayerBlocked is a free data retrieval call binding the contract method 0x0537b4d6.
//
// Solidity: function _relayerBlocked(address ) view returns(bool)
func (_HubEth *HubEthCallerSession) RelayerBlocked(arg0 common.Address) (bool, error) {
	return _HubEth.Contract.RelayerBlocked(&_HubEth.CallOpts, arg0)
}

// RelayersExistMap is a free data retrieval call binding the contract method 0xbd59a093.
//
// Solidity: function _relayersExistMap(address ) view returns(bool)
func (_HubEth *HubEthCaller) RelayersExistMap(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _HubEth.contract.Call(opts, &out, "_relayersExistMap", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// RelayersExistMap is a free data retrieval call binding the contract method 0xbd59a093.
//
// Solidity: function _relayersExistMap(address ) view returns(bool)
func (_HubEth *HubEthSession) RelayersExistMap(arg0 common.Address) (bool, error) {
	return _HubEth.Contract.RelayersExistMap(&_HubEth.CallOpts, arg0)
}

// RelayersExistMap is a free data retrieval call binding the contract method 0xbd59a093.
//
// Solidity: function _relayersExistMap(address ) view returns(bool)
func (_HubEth *HubEthCallerSession) RelayersExistMap(arg0 common.Address) (bool, error) {
	return _HubEth.Contract.RelayersExistMap(&_HubEth.CallOpts, arg0)
}

// RewardPercentage is a free data retrieval call binding the contract method 0xa1444a4a.
//
// Solidity: function _rewardPercentage() view returns(uint256)
func (_HubEth *HubEthCaller) RewardPercentage(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _HubEth.contract.Call(opts, &out, "_rewardPercentage")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardPercentage is a free data retrieval call binding the contract method 0xa1444a4a.
//
// Solidity: function _rewardPercentage() view returns(uint256)
func (_HubEth *HubEthSession) RewardPercentage() (*big.Int, error) {
	return _HubEth.Contract.RewardPercentage(&_HubEth.CallOpts)
}

// RewardPercentage is a free data retrieval call binding the contract method 0xa1444a4a.
//
// Solidity: function _rewardPercentage() view returns(uint256)
func (_HubEth *HubEthCallerSession) RewardPercentage() (*big.Int, error) {
	return _HubEth.Contract.RewardPercentage(&_HubEth.CallOpts)
}

// IsRelayer is a free data retrieval call binding the contract method 0x541d5548.
//
// Solidity: function isRelayer(address sender) view returns(bool)
func (_HubEth *HubEthCaller) IsRelayer(opts *bind.CallOpts, sender common.Address) (bool, error) {
	var out []interface{}
	err := _HubEth.contract.Call(opts, &out, "isRelayer", sender)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRelayer is a free data retrieval call binding the contract method 0x541d5548.
//
// Solidity: function isRelayer(address sender) view returns(bool)
func (_HubEth *HubEthSession) IsRelayer(sender common.Address) (bool, error) {
	return _HubEth.Contract.IsRelayer(&_HubEth.CallOpts, sender)
}

// IsRelayer is a free data retrieval call binding the contract method 0x541d5548.
//
// Solidity: function isRelayer(address sender) view returns(bool)
func (_HubEth *HubEthCallerSession) IsRelayer(sender common.Address) (bool, error) {
	return _HubEth.Contract.IsRelayer(&_HubEth.CallOpts, sender)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_HubEth *HubEthCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HubEth.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_HubEth *HubEthSession) Owner() (common.Address, error) {
	return _HubEth.Contract.Owner(&_HubEth.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_HubEth *HubEthCallerSession) Owner() (common.Address, error) {
	return _HubEth.Contract.Owner(&_HubEth.CallOpts)
}

// AlreadyExecuted is a paid mutator transaction binding the contract method 0x0f03ee59.
//
// Solidity: function alreadyExecuted(address _relayer, uint256 swapValue) returns()
func (_HubEth *HubEthTransactor) AlreadyExecuted(opts *bind.TransactOpts, _relayer common.Address, swapValue *big.Int) (*types.Transaction, error) {
	return _HubEth.contract.Transact(opts, "alreadyExecuted", _relayer, swapValue)
}

// AlreadyExecuted is a paid mutator transaction binding the contract method 0x0f03ee59.
//
// Solidity: function alreadyExecuted(address _relayer, uint256 swapValue) returns()
func (_HubEth *HubEthSession) AlreadyExecuted(_relayer common.Address, swapValue *big.Int) (*types.Transaction, error) {
	return _HubEth.Contract.AlreadyExecuted(&_HubEth.TransactOpts, _relayer, swapValue)
}

// AlreadyExecuted is a paid mutator transaction binding the contract method 0x0f03ee59.
//
// Solidity: function alreadyExecuted(address _relayer, uint256 swapValue) returns()
func (_HubEth *HubEthTransactorSession) AlreadyExecuted(_relayer common.Address, swapValue *big.Int) (*types.Transaction, error) {
	return _HubEth.Contract.AlreadyExecuted(&_HubEth.TransactOpts, _relayer, swapValue)
}

// Felony is a paid mutator transaction binding the contract method 0x35409f7f.
//
// Solidity: function felony(address relayer) returns()
func (_HubEth *HubEthTransactor) Felony(opts *bind.TransactOpts, relayer common.Address) (*types.Transaction, error) {
	return _HubEth.contract.Transact(opts, "felony", relayer)
}

// Felony is a paid mutator transaction binding the contract method 0x35409f7f.
//
// Solidity: function felony(address relayer) returns()
func (_HubEth *HubEthSession) Felony(relayer common.Address) (*types.Transaction, error) {
	return _HubEth.Contract.Felony(&_HubEth.TransactOpts, relayer)
}

// Felony is a paid mutator transaction binding the contract method 0x35409f7f.
//
// Solidity: function felony(address relayer) returns()
func (_HubEth *HubEthTransactorSession) Felony(relayer common.Address) (*types.Transaction, error) {
	return _HubEth.Contract.Felony(&_HubEth.TransactOpts, relayer)
}

// Register is a paid mutator transaction binding the contract method 0x4420e486.
//
// Solidity: function register(address relayerToAdd) returns()
func (_HubEth *HubEthTransactor) Register(opts *bind.TransactOpts, relayerToAdd common.Address) (*types.Transaction, error) {
	return _HubEth.contract.Transact(opts, "register", relayerToAdd)
}

// Register is a paid mutator transaction binding the contract method 0x4420e486.
//
// Solidity: function register(address relayerToAdd) returns()
func (_HubEth *HubEthSession) Register(relayerToAdd common.Address) (*types.Transaction, error) {
	return _HubEth.Contract.Register(&_HubEth.TransactOpts, relayerToAdd)
}

// Register is a paid mutator transaction binding the contract method 0x4420e486.
//
// Solidity: function register(address relayerToAdd) returns()
func (_HubEth *HubEthTransactorSession) Register(relayerToAdd common.Address) (*types.Transaction, error) {
	return _HubEth.Contract.Register(&_HubEth.TransactOpts, relayerToAdd)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_HubEth *HubEthTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HubEth.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_HubEth *HubEthSession) RenounceOwnership() (*types.Transaction, error) {
	return _HubEth.Contract.RenounceOwnership(&_HubEth.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_HubEth *HubEthTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _HubEth.Contract.RenounceOwnership(&_HubEth.TransactOpts)
}

// SwapReward is a paid mutator transaction binding the contract method 0xadaa71cc.
//
// Solidity: function swapReward(address _relayer, uint256 swapValue) returns()
func (_HubEth *HubEthTransactor) SwapReward(opts *bind.TransactOpts, _relayer common.Address, swapValue *big.Int) (*types.Transaction, error) {
	return _HubEth.contract.Transact(opts, "swapReward", _relayer, swapValue)
}

// SwapReward is a paid mutator transaction binding the contract method 0xadaa71cc.
//
// Solidity: function swapReward(address _relayer, uint256 swapValue) returns()
func (_HubEth *HubEthSession) SwapReward(_relayer common.Address, swapValue *big.Int) (*types.Transaction, error) {
	return _HubEth.Contract.SwapReward(&_HubEth.TransactOpts, _relayer, swapValue)
}

// SwapReward is a paid mutator transaction binding the contract method 0xadaa71cc.
//
// Solidity: function swapReward(address _relayer, uint256 swapValue) returns()
func (_HubEth *HubEthTransactorSession) SwapReward(_relayer common.Address, swapValue *big.Int) (*types.Transaction, error) {
	return _HubEth.Contract.SwapReward(&_HubEth.TransactOpts, _relayer, swapValue)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_HubEth *HubEthTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _HubEth.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_HubEth *HubEthSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _HubEth.Contract.TransferOwnership(&_HubEth.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_HubEth *HubEthTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _HubEth.Contract.TransferOwnership(&_HubEth.TransactOpts, newOwner)
}

// Unregister is a paid mutator transaction binding the contract method 0x2ec2c246.
//
// Solidity: function unregister(address relayerToRemove) returns()
func (_HubEth *HubEthTransactor) Unregister(opts *bind.TransactOpts, relayerToRemove common.Address) (*types.Transaction, error) {
	return _HubEth.contract.Transact(opts, "unregister", relayerToRemove)
}

// Unregister is a paid mutator transaction binding the contract method 0x2ec2c246.
//
// Solidity: function unregister(address relayerToRemove) returns()
func (_HubEth *HubEthSession) Unregister(relayerToRemove common.Address) (*types.Transaction, error) {
	return _HubEth.Contract.Unregister(&_HubEth.TransactOpts, relayerToRemove)
}

// Unregister is a paid mutator transaction binding the contract method 0x2ec2c246.
//
// Solidity: function unregister(address relayerToRemove) returns()
func (_HubEth *HubEthTransactorSession) Unregister(relayerToRemove common.Address) (*types.Transaction, error) {
	return _HubEth.Contract.Unregister(&_HubEth.TransactOpts, relayerToRemove)
}

// UpdatePenaltyPercentage is a paid mutator transaction binding the contract method 0xeb961693.
//
// Solidity: function updatePenaltyPercentage(uint256 penaltyPercentage) returns()
func (_HubEth *HubEthTransactor) UpdatePenaltyPercentage(opts *bind.TransactOpts, penaltyPercentage *big.Int) (*types.Transaction, error) {
	return _HubEth.contract.Transact(opts, "updatePenaltyPercentage", penaltyPercentage)
}

// UpdatePenaltyPercentage is a paid mutator transaction binding the contract method 0xeb961693.
//
// Solidity: function updatePenaltyPercentage(uint256 penaltyPercentage) returns()
func (_HubEth *HubEthSession) UpdatePenaltyPercentage(penaltyPercentage *big.Int) (*types.Transaction, error) {
	return _HubEth.Contract.UpdatePenaltyPercentage(&_HubEth.TransactOpts, penaltyPercentage)
}

// UpdatePenaltyPercentage is a paid mutator transaction binding the contract method 0xeb961693.
//
// Solidity: function updatePenaltyPercentage(uint256 penaltyPercentage) returns()
func (_HubEth *HubEthTransactorSession) UpdatePenaltyPercentage(penaltyPercentage *big.Int) (*types.Transaction, error) {
	return _HubEth.Contract.UpdatePenaltyPercentage(&_HubEth.TransactOpts, penaltyPercentage)
}

// UpdateRewardPercentage is a paid mutator transaction binding the contract method 0xa4f209b0.
//
// Solidity: function updateRewardPercentage(uint256 rewardPercent) returns()
func (_HubEth *HubEthTransactor) UpdateRewardPercentage(opts *bind.TransactOpts, rewardPercent *big.Int) (*types.Transaction, error) {
	return _HubEth.contract.Transact(opts, "updateRewardPercentage", rewardPercent)
}

// UpdateRewardPercentage is a paid mutator transaction binding the contract method 0xa4f209b0.
//
// Solidity: function updateRewardPercentage(uint256 rewardPercent) returns()
func (_HubEth *HubEthSession) UpdateRewardPercentage(rewardPercent *big.Int) (*types.Transaction, error) {
	return _HubEth.Contract.UpdateRewardPercentage(&_HubEth.TransactOpts, rewardPercent)
}

// UpdateRewardPercentage is a paid mutator transaction binding the contract method 0xa4f209b0.
//
// Solidity: function updateRewardPercentage(uint256 rewardPercent) returns()
func (_HubEth *HubEthTransactorSession) UpdateRewardPercentage(rewardPercent *big.Int) (*types.Transaction, error) {
	return _HubEth.Contract.UpdateRewardPercentage(&_HubEth.TransactOpts, rewardPercent)
}

// HubEthOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the HubEth contract.
type HubEthOwnershipTransferredIterator struct {
	Event *HubEthOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *HubEthOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HubEthOwnershipTransferred)
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
		it.Event = new(HubEthOwnershipTransferred)
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
func (it *HubEthOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HubEthOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HubEthOwnershipTransferred represents a OwnershipTransferred event raised by the HubEth contract.
type HubEthOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_HubEth *HubEthFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*HubEthOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _HubEth.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &HubEthOwnershipTransferredIterator{contract: _HubEth.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_HubEth *HubEthFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *HubEthOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _HubEth.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HubEthOwnershipTransferred)
				if err := _HubEth.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_HubEth *HubEthFilterer) ParseOwnershipTransferred(log types.Log) (*HubEthOwnershipTransferred, error) {
	event := new(HubEthOwnershipTransferred)
	if err := _HubEth.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HubEthPenaltyForRelayerIterator is returned from FilterPenaltyForRelayer and is used to iterate over the raw logs and unpacked data for PenaltyForRelayer events raised by the HubEth contract.
type HubEthPenaltyForRelayerIterator struct {
	Event *HubEthPenaltyForRelayer // Event containing the contract specifics and raw log

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
func (it *HubEthPenaltyForRelayerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HubEthPenaltyForRelayer)
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
		it.Event = new(HubEthPenaltyForRelayer)
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
func (it *HubEthPenaltyForRelayerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HubEthPenaltyForRelayerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HubEthPenaltyForRelayer represents a PenaltyForRelayer event raised by the HubEth contract.
type HubEthPenaltyForRelayer struct {
	Reason  string
	Relayer common.Address
	Penalty *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPenaltyForRelayer is a free log retrieval operation binding the contract event 0x73f262bdf0c3145b25a03d5c75d5989bdab8ecb77968d7ff6de8a3bd83b8b13b.
//
// Solidity: event PenaltyForRelayer(string reason, address indexed relayer, uint256 penalty)
func (_HubEth *HubEthFilterer) FilterPenaltyForRelayer(opts *bind.FilterOpts, relayer []common.Address) (*HubEthPenaltyForRelayerIterator, error) {

	var relayerRule []interface{}
	for _, relayerItem := range relayer {
		relayerRule = append(relayerRule, relayerItem)
	}

	logs, sub, err := _HubEth.contract.FilterLogs(opts, "PenaltyForRelayer", relayerRule)
	if err != nil {
		return nil, err
	}
	return &HubEthPenaltyForRelayerIterator{contract: _HubEth.contract, event: "PenaltyForRelayer", logs: logs, sub: sub}, nil
}

// WatchPenaltyForRelayer is a free log subscription operation binding the contract event 0x73f262bdf0c3145b25a03d5c75d5989bdab8ecb77968d7ff6de8a3bd83b8b13b.
//
// Solidity: event PenaltyForRelayer(string reason, address indexed relayer, uint256 penalty)
func (_HubEth *HubEthFilterer) WatchPenaltyForRelayer(opts *bind.WatchOpts, sink chan<- *HubEthPenaltyForRelayer, relayer []common.Address) (event.Subscription, error) {

	var relayerRule []interface{}
	for _, relayerItem := range relayer {
		relayerRule = append(relayerRule, relayerItem)
	}

	logs, sub, err := _HubEth.contract.WatchLogs(opts, "PenaltyForRelayer", relayerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HubEthPenaltyForRelayer)
				if err := _HubEth.contract.UnpackLog(event, "PenaltyForRelayer", log); err != nil {
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
// Solidity: event PenaltyForRelayer(string reason, address indexed relayer, uint256 penalty)
func (_HubEth *HubEthFilterer) ParsePenaltyForRelayer(log types.Log) (*HubEthPenaltyForRelayer, error) {
	event := new(HubEthPenaltyForRelayer)
	if err := _HubEth.contract.UnpackLog(event, "PenaltyForRelayer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HubEthPenaltyPercentChangedIterator is returned from FilterPenaltyPercentChanged and is used to iterate over the raw logs and unpacked data for PenaltyPercentChanged events raised by the HubEth contract.
type HubEthPenaltyPercentChangedIterator struct {
	Event *HubEthPenaltyPercentChanged // Event containing the contract specifics and raw log

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
func (it *HubEthPenaltyPercentChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HubEthPenaltyPercentChanged)
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
		it.Event = new(HubEthPenaltyPercentChanged)
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
func (it *HubEthPenaltyPercentChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HubEthPenaltyPercentChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HubEthPenaltyPercentChanged represents a PenaltyPercentChanged event raised by the HubEth contract.
type HubEthPenaltyPercentChanged struct {
	NewPercentage *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterPenaltyPercentChanged is a free log retrieval operation binding the contract event 0xa7422d81f8725b29775f8eaafb874b2d997bb6e6d2a60eda98f5fb2eaeb58f3e.
//
// Solidity: event PenaltyPercentChanged(uint256 newPercentage)
func (_HubEth *HubEthFilterer) FilterPenaltyPercentChanged(opts *bind.FilterOpts) (*HubEthPenaltyPercentChangedIterator, error) {

	logs, sub, err := _HubEth.contract.FilterLogs(opts, "PenaltyPercentChanged")
	if err != nil {
		return nil, err
	}
	return &HubEthPenaltyPercentChangedIterator{contract: _HubEth.contract, event: "PenaltyPercentChanged", logs: logs, sub: sub}, nil
}

// WatchPenaltyPercentChanged is a free log subscription operation binding the contract event 0xa7422d81f8725b29775f8eaafb874b2d997bb6e6d2a60eda98f5fb2eaeb58f3e.
//
// Solidity: event PenaltyPercentChanged(uint256 newPercentage)
func (_HubEth *HubEthFilterer) WatchPenaltyPercentChanged(opts *bind.WatchOpts, sink chan<- *HubEthPenaltyPercentChanged) (event.Subscription, error) {

	logs, sub, err := _HubEth.contract.WatchLogs(opts, "PenaltyPercentChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HubEthPenaltyPercentChanged)
				if err := _HubEth.contract.UnpackLog(event, "PenaltyPercentChanged", log); err != nil {
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
// Solidity: event PenaltyPercentChanged(uint256 newPercentage)
func (_HubEth *HubEthFilterer) ParsePenaltyPercentChanged(log types.Log) (*HubEthPenaltyPercentChanged, error) {
	event := new(HubEthPenaltyPercentChanged)
	if err := _HubEth.contract.UnpackLog(event, "PenaltyPercentChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HubEthRelayerBlockedIterator is returned from FilterRelayerBlocked and is used to iterate over the raw logs and unpacked data for RelayerBlocked events raised by the HubEth contract.
type HubEthRelayerBlockedIterator struct {
	Event *HubEthRelayerBlocked // Event containing the contract specifics and raw log

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
func (it *HubEthRelayerBlockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HubEthRelayerBlocked)
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
		it.Event = new(HubEthRelayerBlocked)
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
func (it *HubEthRelayerBlockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HubEthRelayerBlockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HubEthRelayerBlocked represents a RelayerBlocked event raised by the HubEth contract.
type HubEthRelayerBlocked struct {
	Relayer common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRelayerBlocked is a free log retrieval operation binding the contract event 0xa5ea068b2ed8e05403d639cea236649975a9712487c612f9a6945b6bad00d81d.
//
// Solidity: event RelayerBlocked(address indexed relayer)
func (_HubEth *HubEthFilterer) FilterRelayerBlocked(opts *bind.FilterOpts, relayer []common.Address) (*HubEthRelayerBlockedIterator, error) {

	var relayerRule []interface{}
	for _, relayerItem := range relayer {
		relayerRule = append(relayerRule, relayerItem)
	}

	logs, sub, err := _HubEth.contract.FilterLogs(opts, "RelayerBlocked", relayerRule)
	if err != nil {
		return nil, err
	}
	return &HubEthRelayerBlockedIterator{contract: _HubEth.contract, event: "RelayerBlocked", logs: logs, sub: sub}, nil
}

// WatchRelayerBlocked is a free log subscription operation binding the contract event 0xa5ea068b2ed8e05403d639cea236649975a9712487c612f9a6945b6bad00d81d.
//
// Solidity: event RelayerBlocked(address indexed relayer)
func (_HubEth *HubEthFilterer) WatchRelayerBlocked(opts *bind.WatchOpts, sink chan<- *HubEthRelayerBlocked, relayer []common.Address) (event.Subscription, error) {

	var relayerRule []interface{}
	for _, relayerItem := range relayer {
		relayerRule = append(relayerRule, relayerItem)
	}

	logs, sub, err := _HubEth.contract.WatchLogs(opts, "RelayerBlocked", relayerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HubEthRelayerBlocked)
				if err := _HubEth.contract.UnpackLog(event, "RelayerBlocked", log); err != nil {
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
// Solidity: event RelayerBlocked(address indexed relayer)
func (_HubEth *HubEthFilterer) ParseRelayerBlocked(log types.Log) (*HubEthRelayerBlocked, error) {
	event := new(HubEthRelayerBlocked)
	if err := _HubEth.contract.UnpackLog(event, "RelayerBlocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HubEthRelayerRegisterIterator is returned from FilterRelayerRegister and is used to iterate over the raw logs and unpacked data for RelayerRegister events raised by the HubEth contract.
type HubEthRelayerRegisterIterator struct {
	Event *HubEthRelayerRegister // Event containing the contract specifics and raw log

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
func (it *HubEthRelayerRegisterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HubEthRelayerRegister)
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
		it.Event = new(HubEthRelayerRegister)
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
func (it *HubEthRelayerRegisterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HubEthRelayerRegisterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HubEthRelayerRegister represents a RelayerRegister event raised by the HubEth contract.
type HubEthRelayerRegister struct {
	Relayer common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRelayerRegister is a free log retrieval operation binding the contract event 0x03f64c5eb6e7ea7290123c03dbf1d28d676a583815f5f6a8ebd7f153cfb45215.
//
// Solidity: event RelayerRegister(address indexed _relayer)
func (_HubEth *HubEthFilterer) FilterRelayerRegister(opts *bind.FilterOpts, _relayer []common.Address) (*HubEthRelayerRegisterIterator, error) {

	var _relayerRule []interface{}
	for _, _relayerItem := range _relayer {
		_relayerRule = append(_relayerRule, _relayerItem)
	}

	logs, sub, err := _HubEth.contract.FilterLogs(opts, "RelayerRegister", _relayerRule)
	if err != nil {
		return nil, err
	}
	return &HubEthRelayerRegisterIterator{contract: _HubEth.contract, event: "RelayerRegister", logs: logs, sub: sub}, nil
}

// WatchRelayerRegister is a free log subscription operation binding the contract event 0x03f64c5eb6e7ea7290123c03dbf1d28d676a583815f5f6a8ebd7f153cfb45215.
//
// Solidity: event RelayerRegister(address indexed _relayer)
func (_HubEth *HubEthFilterer) WatchRelayerRegister(opts *bind.WatchOpts, sink chan<- *HubEthRelayerRegister, _relayer []common.Address) (event.Subscription, error) {

	var _relayerRule []interface{}
	for _, _relayerItem := range _relayer {
		_relayerRule = append(_relayerRule, _relayerItem)
	}

	logs, sub, err := _HubEth.contract.WatchLogs(opts, "RelayerRegister", _relayerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HubEthRelayerRegister)
				if err := _HubEth.contract.UnpackLog(event, "RelayerRegister", log); err != nil {
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
// Solidity: event RelayerRegister(address indexed _relayer)
func (_HubEth *HubEthFilterer) ParseRelayerRegister(log types.Log) (*HubEthRelayerRegister, error) {
	event := new(HubEthRelayerRegister)
	if err := _HubEth.contract.UnpackLog(event, "RelayerRegister", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HubEthRelayerUnRegisterIterator is returned from FilterRelayerUnRegister and is used to iterate over the raw logs and unpacked data for RelayerUnRegister events raised by the HubEth contract.
type HubEthRelayerUnRegisterIterator struct {
	Event *HubEthRelayerUnRegister // Event containing the contract specifics and raw log

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
func (it *HubEthRelayerUnRegisterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HubEthRelayerUnRegister)
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
		it.Event = new(HubEthRelayerUnRegister)
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
func (it *HubEthRelayerUnRegisterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HubEthRelayerUnRegisterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HubEthRelayerUnRegister represents a RelayerUnRegister event raised by the HubEth contract.
type HubEthRelayerUnRegister struct {
	Relayer common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRelayerUnRegister is a free log retrieval operation binding the contract event 0x013c40a2b0c93c7ba3e8b43e7529d1d6f020670e66817894c8f86708afb3e62c.
//
// Solidity: event RelayerUnRegister(address indexed _relayer)
func (_HubEth *HubEthFilterer) FilterRelayerUnRegister(opts *bind.FilterOpts, _relayer []common.Address) (*HubEthRelayerUnRegisterIterator, error) {

	var _relayerRule []interface{}
	for _, _relayerItem := range _relayer {
		_relayerRule = append(_relayerRule, _relayerItem)
	}

	logs, sub, err := _HubEth.contract.FilterLogs(opts, "RelayerUnRegister", _relayerRule)
	if err != nil {
		return nil, err
	}
	return &HubEthRelayerUnRegisterIterator{contract: _HubEth.contract, event: "RelayerUnRegister", logs: logs, sub: sub}, nil
}

// WatchRelayerUnRegister is a free log subscription operation binding the contract event 0x013c40a2b0c93c7ba3e8b43e7529d1d6f020670e66817894c8f86708afb3e62c.
//
// Solidity: event RelayerUnRegister(address indexed _relayer)
func (_HubEth *HubEthFilterer) WatchRelayerUnRegister(opts *bind.WatchOpts, sink chan<- *HubEthRelayerUnRegister, _relayer []common.Address) (event.Subscription, error) {

	var _relayerRule []interface{}
	for _, _relayerItem := range _relayer {
		_relayerRule = append(_relayerRule, _relayerItem)
	}

	logs, sub, err := _HubEth.contract.WatchLogs(opts, "RelayerUnRegister", _relayerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HubEthRelayerUnRegister)
				if err := _HubEth.contract.UnpackLog(event, "RelayerUnRegister", log); err != nil {
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
// Solidity: event RelayerUnRegister(address indexed _relayer)
func (_HubEth *HubEthFilterer) ParseRelayerUnRegister(log types.Log) (*HubEthRelayerUnRegister, error) {
	event := new(HubEthRelayerUnRegister)
	if err := _HubEth.contract.UnpackLog(event, "RelayerUnRegister", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HubEthRewardForRelayerIterator is returned from FilterRewardForRelayer and is used to iterate over the raw logs and unpacked data for RewardForRelayer events raised by the HubEth contract.
type HubEthRewardForRelayerIterator struct {
	Event *HubEthRewardForRelayer // Event containing the contract specifics and raw log

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
func (it *HubEthRewardForRelayerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HubEthRewardForRelayer)
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
		it.Event = new(HubEthRewardForRelayer)
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
func (it *HubEthRewardForRelayerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HubEthRewardForRelayerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HubEthRewardForRelayer represents a RewardForRelayer event raised by the HubEth contract.
type HubEthRewardForRelayer struct {
	Relayer common.Address
	Reward  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRewardForRelayer is a free log retrieval operation binding the contract event 0x63dd553807ea27cb996173133b29a47f29c6b3bec34043271f5a5f7dd1f949c9.
//
// Solidity: event RewardForRelayer(address relayer, uint256 reward)
func (_HubEth *HubEthFilterer) FilterRewardForRelayer(opts *bind.FilterOpts) (*HubEthRewardForRelayerIterator, error) {

	logs, sub, err := _HubEth.contract.FilterLogs(opts, "RewardForRelayer")
	if err != nil {
		return nil, err
	}
	return &HubEthRewardForRelayerIterator{contract: _HubEth.contract, event: "RewardForRelayer", logs: logs, sub: sub}, nil
}

// WatchRewardForRelayer is a free log subscription operation binding the contract event 0x63dd553807ea27cb996173133b29a47f29c6b3bec34043271f5a5f7dd1f949c9.
//
// Solidity: event RewardForRelayer(address relayer, uint256 reward)
func (_HubEth *HubEthFilterer) WatchRewardForRelayer(opts *bind.WatchOpts, sink chan<- *HubEthRewardForRelayer) (event.Subscription, error) {

	logs, sub, err := _HubEth.contract.WatchLogs(opts, "RewardForRelayer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HubEthRewardForRelayer)
				if err := _HubEth.contract.UnpackLog(event, "RewardForRelayer", log); err != nil {
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
func (_HubEth *HubEthFilterer) ParseRewardForRelayer(log types.Log) (*HubEthRewardForRelayer, error) {
	event := new(HubEthRewardForRelayer)
	if err := _HubEth.contract.UnpackLog(event, "RewardForRelayer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HubEthRewardPercentageChangedIterator is returned from FilterRewardPercentageChanged and is used to iterate over the raw logs and unpacked data for RewardPercentageChanged events raised by the HubEth contract.
type HubEthRewardPercentageChangedIterator struct {
	Event *HubEthRewardPercentageChanged // Event containing the contract specifics and raw log

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
func (it *HubEthRewardPercentageChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HubEthRewardPercentageChanged)
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
		it.Event = new(HubEthRewardPercentageChanged)
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
func (it *HubEthRewardPercentageChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HubEthRewardPercentageChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HubEthRewardPercentageChanged represents a RewardPercentageChanged event raised by the HubEth contract.
type HubEthRewardPercentageChanged struct {
	NewPercentage *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRewardPercentageChanged is a free log retrieval operation binding the contract event 0x23ad5751be9629264e3725ea971a625cff4033b5a3df9f4526e8a32159c0fa89.
//
// Solidity: event RewardPercentageChanged(uint256 newPercentage)
func (_HubEth *HubEthFilterer) FilterRewardPercentageChanged(opts *bind.FilterOpts) (*HubEthRewardPercentageChangedIterator, error) {

	logs, sub, err := _HubEth.contract.FilterLogs(opts, "RewardPercentageChanged")
	if err != nil {
		return nil, err
	}
	return &HubEthRewardPercentageChangedIterator{contract: _HubEth.contract, event: "RewardPercentageChanged", logs: logs, sub: sub}, nil
}

// WatchRewardPercentageChanged is a free log subscription operation binding the contract event 0x23ad5751be9629264e3725ea971a625cff4033b5a3df9f4526e8a32159c0fa89.
//
// Solidity: event RewardPercentageChanged(uint256 newPercentage)
func (_HubEth *HubEthFilterer) WatchRewardPercentageChanged(opts *bind.WatchOpts, sink chan<- *HubEthRewardPercentageChanged) (event.Subscription, error) {

	logs, sub, err := _HubEth.contract.WatchLogs(opts, "RewardPercentageChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HubEthRewardPercentageChanged)
				if err := _HubEth.contract.UnpackLog(event, "RewardPercentageChanged", log); err != nil {
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
func (_HubEth *HubEthFilterer) ParseRewardPercentageChanged(log types.Log) (*HubEthRewardPercentageChanged, error) {
	event := new(HubEthRewardPercentageChanged)
	if err := _HubEth.contract.UnpackLog(event, "RewardPercentageChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
