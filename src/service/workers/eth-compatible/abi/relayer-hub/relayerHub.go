// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package hub

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// HubABI is the input ABI used to generate the binding from.
const HubABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bridgeAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"penaltyPercentage\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardPercentage\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"penalty\",\"type\":\"uint256\"}],\"name\":\"PenaltyForRelayer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPercentage\",\"type\":\"uint256\"}],\"name\":\"PenaltyPercentChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"}],\"name\":\"RelayerBlocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_relayer\",\"type\":\"address\"}],\"name\":\"RelayerRegister\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_relayer\",\"type\":\"address\"}],\"name\":\"RelayerUnRegister\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"RewardForRelayer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPercentage\",\"type\":\"uint256\"}],\"name\":\"RewardPercentageChanged\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"_penaltyPercentage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_relayerBlocked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_relayersExistMap\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_rewardPercentage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_relayer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"swapValue\",\"type\":\"uint256\"}],\"name\":\"alreadyExecuted\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"}],\"name\":\"felony\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"isRelayer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"relayerToAdd\",\"type\":\"address\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_relayer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"swapValue\",\"type\":\"uint256\"}],\"name\":\"swapReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"relayerToRemove\",\"type\":\"address\"}],\"name\":\"unregister\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"penaltyPercentage\",\"type\":\"uint256\"}],\"name\":\"updatePenaltyPercentage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"rewardPercent\",\"type\":\"uint256\"}],\"name\":\"updateRewardPercentage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Hub is an auto generated Go binding around an Ethereum contract.
type Hub struct {
	HubCaller     // Read-only binding to the contract
	HubTransactor // Write-only binding to the contract
	HubFilterer   // Log filterer for contract events
}

// HubCaller is an auto generated read-only Go binding around an Ethereum contract.
type HubCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HubTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HubTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HubFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HubFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HubSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HubSession struct {
	Contract     *Hub              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HubCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HubCallerSession struct {
	Contract *HubCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// HubTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HubTransactorSession struct {
	Contract     *HubTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HubRaw is an auto generated low-level Go binding around an Ethereum contract.
type HubRaw struct {
	Contract *Hub // Generic contract binding to access the raw methods on
}

// HubCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HubCallerRaw struct {
	Contract *HubCaller // Generic read-only contract binding to access the raw methods on
}

// HubTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HubTransactorRaw struct {
	Contract *HubTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHub creates a new instance of Hub, bound to a specific deployed contract.
func NewHub(address common.Address, backend bind.ContractBackend) (*Hub, error) {
	contract, err := bindHub(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Hub{HubCaller: HubCaller{contract: contract}, HubTransactor: HubTransactor{contract: contract}, HubFilterer: HubFilterer{contract: contract}}, nil
}

// NewHubCaller creates a new read-only instance of Hub, bound to a specific deployed contract.
func NewHubCaller(address common.Address, caller bind.ContractCaller) (*HubCaller, error) {
	contract, err := bindHub(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HubCaller{contract: contract}, nil
}

// NewHubTransactor creates a new write-only instance of Hub, bound to a specific deployed contract.
func NewHubTransactor(address common.Address, transactor bind.ContractTransactor) (*HubTransactor, error) {
	contract, err := bindHub(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HubTransactor{contract: contract}, nil
}

// NewHubFilterer creates a new log filterer instance of Hub, bound to a specific deployed contract.
func NewHubFilterer(address common.Address, filterer bind.ContractFilterer) (*HubFilterer, error) {
	contract, err := bindHub(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HubFilterer{contract: contract}, nil
}

// bindHub binds a generic wrapper to an already deployed contract.
func bindHub(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HubABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Hub *HubRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Hub.Contract.HubCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Hub *HubRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hub.Contract.HubTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Hub *HubRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Hub.Contract.HubTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Hub *HubCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Hub.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Hub *HubTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hub.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Hub *HubTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Hub.Contract.contract.Transact(opts, method, params...)
}

// PenaltyPercentage is a free data retrieval call binding the contract method 0x93a55f92.
//
// Solidity: function _penaltyPercentage() view returns(uint256)
func (_Hub *HubCaller) PenaltyPercentage(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Hub.contract.Call(opts, &out, "_penaltyPercentage")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PenaltyPercentage is a free data retrieval call binding the contract method 0x93a55f92.
//
// Solidity: function _penaltyPercentage() view returns(uint256)
func (_Hub *HubSession) PenaltyPercentage() (*big.Int, error) {
	return _Hub.Contract.PenaltyPercentage(&_Hub.CallOpts)
}

// PenaltyPercentage is a free data retrieval call binding the contract method 0x93a55f92.
//
// Solidity: function _penaltyPercentage() view returns(uint256)
func (_Hub *HubCallerSession) PenaltyPercentage() (*big.Int, error) {
	return _Hub.Contract.PenaltyPercentage(&_Hub.CallOpts)
}

// RelayerBlocked is a free data retrieval call binding the contract method 0x0537b4d6.
//
// Solidity: function _relayerBlocked(address ) view returns(bool)
func (_Hub *HubCaller) RelayerBlocked(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Hub.contract.Call(opts, &out, "_relayerBlocked", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// RelayerBlocked is a free data retrieval call binding the contract method 0x0537b4d6.
//
// Solidity: function _relayerBlocked(address ) view returns(bool)
func (_Hub *HubSession) RelayerBlocked(arg0 common.Address) (bool, error) {
	return _Hub.Contract.RelayerBlocked(&_Hub.CallOpts, arg0)
}

// RelayerBlocked is a free data retrieval call binding the contract method 0x0537b4d6.
//
// Solidity: function _relayerBlocked(address ) view returns(bool)
func (_Hub *HubCallerSession) RelayerBlocked(arg0 common.Address) (bool, error) {
	return _Hub.Contract.RelayerBlocked(&_Hub.CallOpts, arg0)
}

// RelayersExistMap is a free data retrieval call binding the contract method 0xbd59a093.
//
// Solidity: function _relayersExistMap(address ) view returns(bool)
func (_Hub *HubCaller) RelayersExistMap(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Hub.contract.Call(opts, &out, "_relayersExistMap", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// RelayersExistMap is a free data retrieval call binding the contract method 0xbd59a093.
//
// Solidity: function _relayersExistMap(address ) view returns(bool)
func (_Hub *HubSession) RelayersExistMap(arg0 common.Address) (bool, error) {
	return _Hub.Contract.RelayersExistMap(&_Hub.CallOpts, arg0)
}

// RelayersExistMap is a free data retrieval call binding the contract method 0xbd59a093.
//
// Solidity: function _relayersExistMap(address ) view returns(bool)
func (_Hub *HubCallerSession) RelayersExistMap(arg0 common.Address) (bool, error) {
	return _Hub.Contract.RelayersExistMap(&_Hub.CallOpts, arg0)
}

// RewardPercentage is a free data retrieval call binding the contract method 0xa1444a4a.
//
// Solidity: function _rewardPercentage() view returns(uint256)
func (_Hub *HubCaller) RewardPercentage(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Hub.contract.Call(opts, &out, "_rewardPercentage")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardPercentage is a free data retrieval call binding the contract method 0xa1444a4a.
//
// Solidity: function _rewardPercentage() view returns(uint256)
func (_Hub *HubSession) RewardPercentage() (*big.Int, error) {
	return _Hub.Contract.RewardPercentage(&_Hub.CallOpts)
}

// RewardPercentage is a free data retrieval call binding the contract method 0xa1444a4a.
//
// Solidity: function _rewardPercentage() view returns(uint256)
func (_Hub *HubCallerSession) RewardPercentage() (*big.Int, error) {
	return _Hub.Contract.RewardPercentage(&_Hub.CallOpts)
}

// IsRelayer is a free data retrieval call binding the contract method 0x541d5548.
//
// Solidity: function isRelayer(address sender) view returns(bool)
func (_Hub *HubCaller) IsRelayer(opts *bind.CallOpts, sender common.Address) (bool, error) {
	var out []interface{}
	err := _Hub.contract.Call(opts, &out, "isRelayer", sender)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRelayer is a free data retrieval call binding the contract method 0x541d5548.
//
// Solidity: function isRelayer(address sender) view returns(bool)
func (_Hub *HubSession) IsRelayer(sender common.Address) (bool, error) {
	return _Hub.Contract.IsRelayer(&_Hub.CallOpts, sender)
}

// IsRelayer is a free data retrieval call binding the contract method 0x541d5548.
//
// Solidity: function isRelayer(address sender) view returns(bool)
func (_Hub *HubCallerSession) IsRelayer(sender common.Address) (bool, error) {
	return _Hub.Contract.IsRelayer(&_Hub.CallOpts, sender)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Hub *HubCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Hub.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Hub *HubSession) Owner() (common.Address, error) {
	return _Hub.Contract.Owner(&_Hub.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Hub *HubCallerSession) Owner() (common.Address, error) {
	return _Hub.Contract.Owner(&_Hub.CallOpts)
}

// AlreadyExecuted is a paid mutator transaction binding the contract method 0x0f03ee59.
//
// Solidity: function alreadyExecuted(address _relayer, uint256 swapValue) returns()
func (_Hub *HubTransactor) AlreadyExecuted(opts *bind.TransactOpts, _relayer common.Address, swapValue *big.Int) (*types.Transaction, error) {
	return _Hub.contract.Transact(opts, "alreadyExecuted", _relayer, swapValue)
}

// AlreadyExecuted is a paid mutator transaction binding the contract method 0x0f03ee59.
//
// Solidity: function alreadyExecuted(address _relayer, uint256 swapValue) returns()
func (_Hub *HubSession) AlreadyExecuted(_relayer common.Address, swapValue *big.Int) (*types.Transaction, error) {
	return _Hub.Contract.AlreadyExecuted(&_Hub.TransactOpts, _relayer, swapValue)
}

// AlreadyExecuted is a paid mutator transaction binding the contract method 0x0f03ee59.
//
// Solidity: function alreadyExecuted(address _relayer, uint256 swapValue) returns()
func (_Hub *HubTransactorSession) AlreadyExecuted(_relayer common.Address, swapValue *big.Int) (*types.Transaction, error) {
	return _Hub.Contract.AlreadyExecuted(&_Hub.TransactOpts, _relayer, swapValue)
}

// Felony is a paid mutator transaction binding the contract method 0x35409f7f.
//
// Solidity: function felony(address relayer) returns()
func (_Hub *HubTransactor) Felony(opts *bind.TransactOpts, relayer common.Address) (*types.Transaction, error) {
	return _Hub.contract.Transact(opts, "felony", relayer)
}

// Felony is a paid mutator transaction binding the contract method 0x35409f7f.
//
// Solidity: function felony(address relayer) returns()
func (_Hub *HubSession) Felony(relayer common.Address) (*types.Transaction, error) {
	return _Hub.Contract.Felony(&_Hub.TransactOpts, relayer)
}

// Felony is a paid mutator transaction binding the contract method 0x35409f7f.
//
// Solidity: function felony(address relayer) returns()
func (_Hub *HubTransactorSession) Felony(relayer common.Address) (*types.Transaction, error) {
	return _Hub.Contract.Felony(&_Hub.TransactOpts, relayer)
}

// Register is a paid mutator transaction binding the contract method 0x4420e486.
//
// Solidity: function register(address relayerToAdd) returns()
func (_Hub *HubTransactor) Register(opts *bind.TransactOpts, relayerToAdd common.Address) (*types.Transaction, error) {
	return _Hub.contract.Transact(opts, "register", relayerToAdd)
}

// Register is a paid mutator transaction binding the contract method 0x4420e486.
//
// Solidity: function register(address relayerToAdd) returns()
func (_Hub *HubSession) Register(relayerToAdd common.Address) (*types.Transaction, error) {
	return _Hub.Contract.Register(&_Hub.TransactOpts, relayerToAdd)
}

// Register is a paid mutator transaction binding the contract method 0x4420e486.
//
// Solidity: function register(address relayerToAdd) returns()
func (_Hub *HubTransactorSession) Register(relayerToAdd common.Address) (*types.Transaction, error) {
	return _Hub.Contract.Register(&_Hub.TransactOpts, relayerToAdd)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Hub *HubTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hub.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Hub *HubSession) RenounceOwnership() (*types.Transaction, error) {
	return _Hub.Contract.RenounceOwnership(&_Hub.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Hub *HubTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Hub.Contract.RenounceOwnership(&_Hub.TransactOpts)
}

// SwapReward is a paid mutator transaction binding the contract method 0xadaa71cc.
//
// Solidity: function swapReward(address _relayer, uint256 swapValue) returns()
func (_Hub *HubTransactor) SwapReward(opts *bind.TransactOpts, _relayer common.Address, swapValue *big.Int) (*types.Transaction, error) {
	return _Hub.contract.Transact(opts, "swapReward", _relayer, swapValue)
}

// SwapReward is a paid mutator transaction binding the contract method 0xadaa71cc.
//
// Solidity: function swapReward(address _relayer, uint256 swapValue) returns()
func (_Hub *HubSession) SwapReward(_relayer common.Address, swapValue *big.Int) (*types.Transaction, error) {
	return _Hub.Contract.SwapReward(&_Hub.TransactOpts, _relayer, swapValue)
}

// SwapReward is a paid mutator transaction binding the contract method 0xadaa71cc.
//
// Solidity: function swapReward(address _relayer, uint256 swapValue) returns()
func (_Hub *HubTransactorSession) SwapReward(_relayer common.Address, swapValue *big.Int) (*types.Transaction, error) {
	return _Hub.Contract.SwapReward(&_Hub.TransactOpts, _relayer, swapValue)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Hub *HubTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Hub.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Hub *HubSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Hub.Contract.TransferOwnership(&_Hub.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Hub *HubTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Hub.Contract.TransferOwnership(&_Hub.TransactOpts, newOwner)
}

// Unregister is a paid mutator transaction binding the contract method 0x2ec2c246.
//
// Solidity: function unregister(address relayerToRemove) returns()
func (_Hub *HubTransactor) Unregister(opts *bind.TransactOpts, relayerToRemove common.Address) (*types.Transaction, error) {
	return _Hub.contract.Transact(opts, "unregister", relayerToRemove)
}

// Unregister is a paid mutator transaction binding the contract method 0x2ec2c246.
//
// Solidity: function unregister(address relayerToRemove) returns()
func (_Hub *HubSession) Unregister(relayerToRemove common.Address) (*types.Transaction, error) {
	return _Hub.Contract.Unregister(&_Hub.TransactOpts, relayerToRemove)
}

// Unregister is a paid mutator transaction binding the contract method 0x2ec2c246.
//
// Solidity: function unregister(address relayerToRemove) returns()
func (_Hub *HubTransactorSession) Unregister(relayerToRemove common.Address) (*types.Transaction, error) {
	return _Hub.Contract.Unregister(&_Hub.TransactOpts, relayerToRemove)
}

// UpdatePenaltyPercentage is a paid mutator transaction binding the contract method 0xeb961693.
//
// Solidity: function updatePenaltyPercentage(uint256 penaltyPercentage) returns()
func (_Hub *HubTransactor) UpdatePenaltyPercentage(opts *bind.TransactOpts, penaltyPercentage *big.Int) (*types.Transaction, error) {
	return _Hub.contract.Transact(opts, "updatePenaltyPercentage", penaltyPercentage)
}

// UpdatePenaltyPercentage is a paid mutator transaction binding the contract method 0xeb961693.
//
// Solidity: function updatePenaltyPercentage(uint256 penaltyPercentage) returns()
func (_Hub *HubSession) UpdatePenaltyPercentage(penaltyPercentage *big.Int) (*types.Transaction, error) {
	return _Hub.Contract.UpdatePenaltyPercentage(&_Hub.TransactOpts, penaltyPercentage)
}

// UpdatePenaltyPercentage is a paid mutator transaction binding the contract method 0xeb961693.
//
// Solidity: function updatePenaltyPercentage(uint256 penaltyPercentage) returns()
func (_Hub *HubTransactorSession) UpdatePenaltyPercentage(penaltyPercentage *big.Int) (*types.Transaction, error) {
	return _Hub.Contract.UpdatePenaltyPercentage(&_Hub.TransactOpts, penaltyPercentage)
}

// UpdateRewardPercentage is a paid mutator transaction binding the contract method 0xa4f209b0.
//
// Solidity: function updateRewardPercentage(uint256 rewardPercent) returns()
func (_Hub *HubTransactor) UpdateRewardPercentage(opts *bind.TransactOpts, rewardPercent *big.Int) (*types.Transaction, error) {
	return _Hub.contract.Transact(opts, "updateRewardPercentage", rewardPercent)
}

// UpdateRewardPercentage is a paid mutator transaction binding the contract method 0xa4f209b0.
//
// Solidity: function updateRewardPercentage(uint256 rewardPercent) returns()
func (_Hub *HubSession) UpdateRewardPercentage(rewardPercent *big.Int) (*types.Transaction, error) {
	return _Hub.Contract.UpdateRewardPercentage(&_Hub.TransactOpts, rewardPercent)
}

// UpdateRewardPercentage is a paid mutator transaction binding the contract method 0xa4f209b0.
//
// Solidity: function updateRewardPercentage(uint256 rewardPercent) returns()
func (_Hub *HubTransactorSession) UpdateRewardPercentage(rewardPercent *big.Int) (*types.Transaction, error) {
	return _Hub.Contract.UpdateRewardPercentage(&_Hub.TransactOpts, rewardPercent)
}

// HubOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Hub contract.
type HubOwnershipTransferredIterator struct {
	Event *HubOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *HubOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HubOwnershipTransferred)
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
		it.Event = new(HubOwnershipTransferred)
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
func (it *HubOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HubOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HubOwnershipTransferred represents a OwnershipTransferred event raised by the Hub contract.
type HubOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Hub *HubFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*HubOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Hub.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &HubOwnershipTransferredIterator{contract: _Hub.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Hub *HubFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *HubOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Hub.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HubOwnershipTransferred)
				if err := _Hub.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Hub *HubFilterer) ParseOwnershipTransferred(log types.Log) (*HubOwnershipTransferred, error) {
	event := new(HubOwnershipTransferred)
	if err := _Hub.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HubPenaltyForRelayerIterator is returned from FilterPenaltyForRelayer and is used to iterate over the raw logs and unpacked data for PenaltyForRelayer events raised by the Hub contract.
type HubPenaltyForRelayerIterator struct {
	Event *HubPenaltyForRelayer // Event containing the contract specifics and raw log

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
func (it *HubPenaltyForRelayerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HubPenaltyForRelayer)
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
		it.Event = new(HubPenaltyForRelayer)
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
func (it *HubPenaltyForRelayerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HubPenaltyForRelayerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HubPenaltyForRelayer represents a PenaltyForRelayer event raised by the Hub contract.
type HubPenaltyForRelayer struct {
	Reason  string
	Relayer common.Address
	Penalty *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPenaltyForRelayer is a free log retrieval operation binding the contract event 0x73f262bdf0c3145b25a03d5c75d5989bdab8ecb77968d7ff6de8a3bd83b8b13b.
//
// Solidity: event PenaltyForRelayer(string reason, address indexed relayer, uint256 penalty)
func (_Hub *HubFilterer) FilterPenaltyForRelayer(opts *bind.FilterOpts, relayer []common.Address) (*HubPenaltyForRelayerIterator, error) {

	var relayerRule []interface{}
	for _, relayerItem := range relayer {
		relayerRule = append(relayerRule, relayerItem)
	}

	logs, sub, err := _Hub.contract.FilterLogs(opts, "PenaltyForRelayer", relayerRule)
	if err != nil {
		return nil, err
	}
	return &HubPenaltyForRelayerIterator{contract: _Hub.contract, event: "PenaltyForRelayer", logs: logs, sub: sub}, nil
}

// WatchPenaltyForRelayer is a free log subscription operation binding the contract event 0x73f262bdf0c3145b25a03d5c75d5989bdab8ecb77968d7ff6de8a3bd83b8b13b.
//
// Solidity: event PenaltyForRelayer(string reason, address indexed relayer, uint256 penalty)
func (_Hub *HubFilterer) WatchPenaltyForRelayer(opts *bind.WatchOpts, sink chan<- *HubPenaltyForRelayer, relayer []common.Address) (event.Subscription, error) {

	var relayerRule []interface{}
	for _, relayerItem := range relayer {
		relayerRule = append(relayerRule, relayerItem)
	}

	logs, sub, err := _Hub.contract.WatchLogs(opts, "PenaltyForRelayer", relayerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HubPenaltyForRelayer)
				if err := _Hub.contract.UnpackLog(event, "PenaltyForRelayer", log); err != nil {
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
func (_Hub *HubFilterer) ParsePenaltyForRelayer(log types.Log) (*HubPenaltyForRelayer, error) {
	event := new(HubPenaltyForRelayer)
	if err := _Hub.contract.UnpackLog(event, "PenaltyForRelayer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HubPenaltyPercentChangedIterator is returned from FilterPenaltyPercentChanged and is used to iterate over the raw logs and unpacked data for PenaltyPercentChanged events raised by the Hub contract.
type HubPenaltyPercentChangedIterator struct {
	Event *HubPenaltyPercentChanged // Event containing the contract specifics and raw log

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
func (it *HubPenaltyPercentChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HubPenaltyPercentChanged)
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
		it.Event = new(HubPenaltyPercentChanged)
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
func (it *HubPenaltyPercentChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HubPenaltyPercentChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HubPenaltyPercentChanged represents a PenaltyPercentChanged event raised by the Hub contract.
type HubPenaltyPercentChanged struct {
	NewPercentage *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterPenaltyPercentChanged is a free log retrieval operation binding the contract event 0xa7422d81f8725b29775f8eaafb874b2d997bb6e6d2a60eda98f5fb2eaeb58f3e.
//
// Solidity: event PenaltyPercentChanged(uint256 newPercentage)
func (_Hub *HubFilterer) FilterPenaltyPercentChanged(opts *bind.FilterOpts) (*HubPenaltyPercentChangedIterator, error) {

	logs, sub, err := _Hub.contract.FilterLogs(opts, "PenaltyPercentChanged")
	if err != nil {
		return nil, err
	}
	return &HubPenaltyPercentChangedIterator{contract: _Hub.contract, event: "PenaltyPercentChanged", logs: logs, sub: sub}, nil
}

// WatchPenaltyPercentChanged is a free log subscription operation binding the contract event 0xa7422d81f8725b29775f8eaafb874b2d997bb6e6d2a60eda98f5fb2eaeb58f3e.
//
// Solidity: event PenaltyPercentChanged(uint256 newPercentage)
func (_Hub *HubFilterer) WatchPenaltyPercentChanged(opts *bind.WatchOpts, sink chan<- *HubPenaltyPercentChanged) (event.Subscription, error) {

	logs, sub, err := _Hub.contract.WatchLogs(opts, "PenaltyPercentChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HubPenaltyPercentChanged)
				if err := _Hub.contract.UnpackLog(event, "PenaltyPercentChanged", log); err != nil {
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
func (_Hub *HubFilterer) ParsePenaltyPercentChanged(log types.Log) (*HubPenaltyPercentChanged, error) {
	event := new(HubPenaltyPercentChanged)
	if err := _Hub.contract.UnpackLog(event, "PenaltyPercentChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HubRelayerBlockedIterator is returned from FilterRelayerBlocked and is used to iterate over the raw logs and unpacked data for RelayerBlocked events raised by the Hub contract.
type HubRelayerBlockedIterator struct {
	Event *HubRelayerBlocked // Event containing the contract specifics and raw log

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
func (it *HubRelayerBlockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HubRelayerBlocked)
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
		it.Event = new(HubRelayerBlocked)
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
func (it *HubRelayerBlockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HubRelayerBlockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HubRelayerBlocked represents a RelayerBlocked event raised by the Hub contract.
type HubRelayerBlocked struct {
	Relayer common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRelayerBlocked is a free log retrieval operation binding the contract event 0xa5ea068b2ed8e05403d639cea236649975a9712487c612f9a6945b6bad00d81d.
//
// Solidity: event RelayerBlocked(address indexed relayer)
func (_Hub *HubFilterer) FilterRelayerBlocked(opts *bind.FilterOpts, relayer []common.Address) (*HubRelayerBlockedIterator, error) {

	var relayerRule []interface{}
	for _, relayerItem := range relayer {
		relayerRule = append(relayerRule, relayerItem)
	}

	logs, sub, err := _Hub.contract.FilterLogs(opts, "RelayerBlocked", relayerRule)
	if err != nil {
		return nil, err
	}
	return &HubRelayerBlockedIterator{contract: _Hub.contract, event: "RelayerBlocked", logs: logs, sub: sub}, nil
}

// WatchRelayerBlocked is a free log subscription operation binding the contract event 0xa5ea068b2ed8e05403d639cea236649975a9712487c612f9a6945b6bad00d81d.
//
// Solidity: event RelayerBlocked(address indexed relayer)
func (_Hub *HubFilterer) WatchRelayerBlocked(opts *bind.WatchOpts, sink chan<- *HubRelayerBlocked, relayer []common.Address) (event.Subscription, error) {

	var relayerRule []interface{}
	for _, relayerItem := range relayer {
		relayerRule = append(relayerRule, relayerItem)
	}

	logs, sub, err := _Hub.contract.WatchLogs(opts, "RelayerBlocked", relayerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HubRelayerBlocked)
				if err := _Hub.contract.UnpackLog(event, "RelayerBlocked", log); err != nil {
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
func (_Hub *HubFilterer) ParseRelayerBlocked(log types.Log) (*HubRelayerBlocked, error) {
	event := new(HubRelayerBlocked)
	if err := _Hub.contract.UnpackLog(event, "RelayerBlocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HubRelayerRegisterIterator is returned from FilterRelayerRegister and is used to iterate over the raw logs and unpacked data for RelayerRegister events raised by the Hub contract.
type HubRelayerRegisterIterator struct {
	Event *HubRelayerRegister // Event containing the contract specifics and raw log

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
func (it *HubRelayerRegisterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HubRelayerRegister)
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
		it.Event = new(HubRelayerRegister)
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
func (it *HubRelayerRegisterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HubRelayerRegisterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HubRelayerRegister represents a RelayerRegister event raised by the Hub contract.
type HubRelayerRegister struct {
	Relayer common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRelayerRegister is a free log retrieval operation binding the contract event 0x03f64c5eb6e7ea7290123c03dbf1d28d676a583815f5f6a8ebd7f153cfb45215.
//
// Solidity: event RelayerRegister(address indexed _relayer)
func (_Hub *HubFilterer) FilterRelayerRegister(opts *bind.FilterOpts, _relayer []common.Address) (*HubRelayerRegisterIterator, error) {

	var _relayerRule []interface{}
	for _, _relayerItem := range _relayer {
		_relayerRule = append(_relayerRule, _relayerItem)
	}

	logs, sub, err := _Hub.contract.FilterLogs(opts, "RelayerRegister", _relayerRule)
	if err != nil {
		return nil, err
	}
	return &HubRelayerRegisterIterator{contract: _Hub.contract, event: "RelayerRegister", logs: logs, sub: sub}, nil
}

// WatchRelayerRegister is a free log subscription operation binding the contract event 0x03f64c5eb6e7ea7290123c03dbf1d28d676a583815f5f6a8ebd7f153cfb45215.
//
// Solidity: event RelayerRegister(address indexed _relayer)
func (_Hub *HubFilterer) WatchRelayerRegister(opts *bind.WatchOpts, sink chan<- *HubRelayerRegister, _relayer []common.Address) (event.Subscription, error) {

	var _relayerRule []interface{}
	for _, _relayerItem := range _relayer {
		_relayerRule = append(_relayerRule, _relayerItem)
	}

	logs, sub, err := _Hub.contract.WatchLogs(opts, "RelayerRegister", _relayerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HubRelayerRegister)
				if err := _Hub.contract.UnpackLog(event, "RelayerRegister", log); err != nil {
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
func (_Hub *HubFilterer) ParseRelayerRegister(log types.Log) (*HubRelayerRegister, error) {
	event := new(HubRelayerRegister)
	if err := _Hub.contract.UnpackLog(event, "RelayerRegister", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HubRelayerUnRegisterIterator is returned from FilterRelayerUnRegister and is used to iterate over the raw logs and unpacked data for RelayerUnRegister events raised by the Hub contract.
type HubRelayerUnRegisterIterator struct {
	Event *HubRelayerUnRegister // Event containing the contract specifics and raw log

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
func (it *HubRelayerUnRegisterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HubRelayerUnRegister)
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
		it.Event = new(HubRelayerUnRegister)
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
func (it *HubRelayerUnRegisterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HubRelayerUnRegisterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HubRelayerUnRegister represents a RelayerUnRegister event raised by the Hub contract.
type HubRelayerUnRegister struct {
	Relayer common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRelayerUnRegister is a free log retrieval operation binding the contract event 0x013c40a2b0c93c7ba3e8b43e7529d1d6f020670e66817894c8f86708afb3e62c.
//
// Solidity: event RelayerUnRegister(address indexed _relayer)
func (_Hub *HubFilterer) FilterRelayerUnRegister(opts *bind.FilterOpts, _relayer []common.Address) (*HubRelayerUnRegisterIterator, error) {

	var _relayerRule []interface{}
	for _, _relayerItem := range _relayer {
		_relayerRule = append(_relayerRule, _relayerItem)
	}

	logs, sub, err := _Hub.contract.FilterLogs(opts, "RelayerUnRegister", _relayerRule)
	if err != nil {
		return nil, err
	}
	return &HubRelayerUnRegisterIterator{contract: _Hub.contract, event: "RelayerUnRegister", logs: logs, sub: sub}, nil
}

// WatchRelayerUnRegister is a free log subscription operation binding the contract event 0x013c40a2b0c93c7ba3e8b43e7529d1d6f020670e66817894c8f86708afb3e62c.
//
// Solidity: event RelayerUnRegister(address indexed _relayer)
func (_Hub *HubFilterer) WatchRelayerUnRegister(opts *bind.WatchOpts, sink chan<- *HubRelayerUnRegister, _relayer []common.Address) (event.Subscription, error) {

	var _relayerRule []interface{}
	for _, _relayerItem := range _relayer {
		_relayerRule = append(_relayerRule, _relayerItem)
	}

	logs, sub, err := _Hub.contract.WatchLogs(opts, "RelayerUnRegister", _relayerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HubRelayerUnRegister)
				if err := _Hub.contract.UnpackLog(event, "RelayerUnRegister", log); err != nil {
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
func (_Hub *HubFilterer) ParseRelayerUnRegister(log types.Log) (*HubRelayerUnRegister, error) {
	event := new(HubRelayerUnRegister)
	if err := _Hub.contract.UnpackLog(event, "RelayerUnRegister", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HubRewardForRelayerIterator is returned from FilterRewardForRelayer and is used to iterate over the raw logs and unpacked data for RewardForRelayer events raised by the Hub contract.
type HubRewardForRelayerIterator struct {
	Event *HubRewardForRelayer // Event containing the contract specifics and raw log

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
func (it *HubRewardForRelayerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HubRewardForRelayer)
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
		it.Event = new(HubRewardForRelayer)
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
func (it *HubRewardForRelayerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HubRewardForRelayerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HubRewardForRelayer represents a RewardForRelayer event raised by the Hub contract.
type HubRewardForRelayer struct {
	Relayer common.Address
	Reward  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRewardForRelayer is a free log retrieval operation binding the contract event 0x63dd553807ea27cb996173133b29a47f29c6b3bec34043271f5a5f7dd1f949c9.
//
// Solidity: event RewardForRelayer(address relayer, uint256 reward)
func (_Hub *HubFilterer) FilterRewardForRelayer(opts *bind.FilterOpts) (*HubRewardForRelayerIterator, error) {

	logs, sub, err := _Hub.contract.FilterLogs(opts, "RewardForRelayer")
	if err != nil {
		return nil, err
	}
	return &HubRewardForRelayerIterator{contract: _Hub.contract, event: "RewardForRelayer", logs: logs, sub: sub}, nil
}

// WatchRewardForRelayer is a free log subscription operation binding the contract event 0x63dd553807ea27cb996173133b29a47f29c6b3bec34043271f5a5f7dd1f949c9.
//
// Solidity: event RewardForRelayer(address relayer, uint256 reward)
func (_Hub *HubFilterer) WatchRewardForRelayer(opts *bind.WatchOpts, sink chan<- *HubRewardForRelayer) (event.Subscription, error) {

	logs, sub, err := _Hub.contract.WatchLogs(opts, "RewardForRelayer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HubRewardForRelayer)
				if err := _Hub.contract.UnpackLog(event, "RewardForRelayer", log); err != nil {
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
func (_Hub *HubFilterer) ParseRewardForRelayer(log types.Log) (*HubRewardForRelayer, error) {
	event := new(HubRewardForRelayer)
	if err := _Hub.contract.UnpackLog(event, "RewardForRelayer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HubRewardPercentageChangedIterator is returned from FilterRewardPercentageChanged and is used to iterate over the raw logs and unpacked data for RewardPercentageChanged events raised by the Hub contract.
type HubRewardPercentageChangedIterator struct {
	Event *HubRewardPercentageChanged // Event containing the contract specifics and raw log

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
func (it *HubRewardPercentageChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HubRewardPercentageChanged)
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
		it.Event = new(HubRewardPercentageChanged)
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
func (it *HubRewardPercentageChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HubRewardPercentageChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HubRewardPercentageChanged represents a RewardPercentageChanged event raised by the Hub contract.
type HubRewardPercentageChanged struct {
	NewPercentage *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRewardPercentageChanged is a free log retrieval operation binding the contract event 0x23ad5751be9629264e3725ea971a625cff4033b5a3df9f4526e8a32159c0fa89.
//
// Solidity: event RewardPercentageChanged(uint256 newPercentage)
func (_Hub *HubFilterer) FilterRewardPercentageChanged(opts *bind.FilterOpts) (*HubRewardPercentageChangedIterator, error) {

	logs, sub, err := _Hub.contract.FilterLogs(opts, "RewardPercentageChanged")
	if err != nil {
		return nil, err
	}
	return &HubRewardPercentageChangedIterator{contract: _Hub.contract, event: "RewardPercentageChanged", logs: logs, sub: sub}, nil
}

// WatchRewardPercentageChanged is a free log subscription operation binding the contract event 0x23ad5751be9629264e3725ea971a625cff4033b5a3df9f4526e8a32159c0fa89.
//
// Solidity: event RewardPercentageChanged(uint256 newPercentage)
func (_Hub *HubFilterer) WatchRewardPercentageChanged(opts *bind.WatchOpts, sink chan<- *HubRewardPercentageChanged) (event.Subscription, error) {

	logs, sub, err := _Hub.contract.WatchLogs(opts, "RewardPercentageChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HubRewardPercentageChanged)
				if err := _Hub.contract.UnpackLog(event, "RewardPercentageChanged", log); err != nil {
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
func (_Hub *HubFilterer) ParseRewardPercentageChanged(log types.Log) (*HubRewardPercentageChanged, error) {
	event := new(HubRewardPercentageChanged)
	if err := _Hub.contract.UnpackLog(event, "RewardPercentageChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
