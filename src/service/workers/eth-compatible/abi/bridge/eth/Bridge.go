// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ethBr

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

// EthBrMetaData contains all meta data concerning the EthBr contract.
var EthBrMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes8\",\"name\":\"chainID\",\"type\":\"bytes8\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"initBackendSrvAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes8\",\"name\":\"originChainID\",\"type\":\"bytes8\"},{\"indexed\":true,\"internalType\":\"bytes8\",\"name\":\"destinationChainID\",\"type\":\"bytes8\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"depositNonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipientAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes8\",\"name\":\"originChainID\",\"type\":\"bytes8\"},{\"indexed\":true,\"internalType\":\"bytes8\",\"name\":\"destinationChainID\",\"type\":\"bytes8\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipientAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"depositNonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"enumBridge.ProposalStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"name\":\"ProposalEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_chainID\",\"outputs\":[{\"internalType\":\"bytes8\",\"name\":\"\",\"type\":\"bytes8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes8\",\"name\":\"\",\"type\":\"bytes8\"}],\"name\":\"_depositCounts\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"bytes8\",\"name\":\"\",\"type\":\"bytes8\"}],\"name\":\"_depositRecords\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"_executedProposals\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_fee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"_resourceIDToHandlerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newFee\",\"type\":\"uint256\"}],\"name\":\"adminChangeFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"adminPauseTransfers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newBackendSrv\",\"type\":\"address\"}],\"name\":\"adminSetBackendSrv\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"handlerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"adminSetBurnable\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"handlerAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"adminSetResource\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"adminUnpauseTransfers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"handlerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountOrTokenID\",\"type\":\"uint256\"}],\"name\":\"adminWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes8\",\"name\":\"destinationChainID\",\"type\":\"bytes8\"},{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipientAddress\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes8\",\"name\":\"originChainID\",\"type\":\"bytes8\"},{\"internalType\":\"bytes8\",\"name\":\"destinationChainID\",\"type\":\"bytes8\"},{\"internalType\":\"uint64\",\"name\":\"depositNonce\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"recipientAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"executeProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"renounceAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"b\",\"type\":\"uint256\"}],\"name\":\"sub\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
}

// EthBrABI is the input ABI used to generate the binding from.
// Deprecated: Use EthBrMetaData.ABI instead.
var EthBrABI = EthBrMetaData.ABI

// EthBr is an auto generated Go binding around an Ethereum contract.
type EthBr struct {
	EthBrCaller     // Read-only binding to the contract
	EthBrTransactor // Write-only binding to the contract
	EthBrFilterer   // Log filterer for contract events
}

// EthBrCaller is an auto generated read-only Go binding around an Ethereum contract.
type EthBrCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthBrTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EthBrTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthBrFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EthBrFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthBrSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EthBrSession struct {
	Contract     *EthBr            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EthBrCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EthBrCallerSession struct {
	Contract *EthBrCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// EthBrTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EthBrTransactorSession struct {
	Contract     *EthBrTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EthBrRaw is an auto generated low-level Go binding around an Ethereum contract.
type EthBrRaw struct {
	Contract *EthBr // Generic contract binding to access the raw methods on
}

// EthBrCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EthBrCallerRaw struct {
	Contract *EthBrCaller // Generic read-only contract binding to access the raw methods on
}

// EthBrTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EthBrTransactorRaw struct {
	Contract *EthBrTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEthBr creates a new instance of EthBr, bound to a specific deployed contract.
func NewEthBr(address common.Address, backend bind.ContractBackend) (*EthBr, error) {
	contract, err := bindEthBr(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EthBr{EthBrCaller: EthBrCaller{contract: contract}, EthBrTransactor: EthBrTransactor{contract: contract}, EthBrFilterer: EthBrFilterer{contract: contract}}, nil
}

// NewEthBrCaller creates a new read-only instance of EthBr, bound to a specific deployed contract.
func NewEthBrCaller(address common.Address, caller bind.ContractCaller) (*EthBrCaller, error) {
	contract, err := bindEthBr(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EthBrCaller{contract: contract}, nil
}

// NewEthBrTransactor creates a new write-only instance of EthBr, bound to a specific deployed contract.
func NewEthBrTransactor(address common.Address, transactor bind.ContractTransactor) (*EthBrTransactor, error) {
	contract, err := bindEthBr(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EthBrTransactor{contract: contract}, nil
}

// NewEthBrFilterer creates a new log filterer instance of EthBr, bound to a specific deployed contract.
func NewEthBrFilterer(address common.Address, filterer bind.ContractFilterer) (*EthBrFilterer, error) {
	contract, err := bindEthBr(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EthBrFilterer{contract: contract}, nil
}

// bindEthBr binds a generic wrapper to an already deployed contract.
func bindEthBr(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EthBrABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EthBr *EthBrRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EthBr.Contract.EthBrCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EthBr *EthBrRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthBr.Contract.EthBrTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EthBr *EthBrRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EthBr.Contract.EthBrTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EthBr *EthBrCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EthBr.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EthBr *EthBrTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthBr.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EthBr *EthBrTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EthBr.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_EthBr *EthBrCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _EthBr.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_EthBr *EthBrSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _EthBr.Contract.DEFAULTADMINROLE(&_EthBr.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_EthBr *EthBrCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _EthBr.Contract.DEFAULTADMINROLE(&_EthBr.CallOpts)
}

// ChainID is a free data retrieval call binding the contract method 0xbeab7131.
//
// Solidity: function _chainID() view returns(bytes8)
func (_EthBr *EthBrCaller) ChainID(opts *bind.CallOpts) ([8]byte, error) {
	var out []interface{}
	err := _EthBr.contract.Call(opts, &out, "_chainID")

	if err != nil {
		return *new([8]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([8]byte)).(*[8]byte)

	return out0, err

}

// ChainID is a free data retrieval call binding the contract method 0xbeab7131.
//
// Solidity: function _chainID() view returns(bytes8)
func (_EthBr *EthBrSession) ChainID() ([8]byte, error) {
	return _EthBr.Contract.ChainID(&_EthBr.CallOpts)
}

// ChainID is a free data retrieval call binding the contract method 0xbeab7131.
//
// Solidity: function _chainID() view returns(bytes8)
func (_EthBr *EthBrCallerSession) ChainID() ([8]byte, error) {
	return _EthBr.Contract.ChainID(&_EthBr.CallOpts)
}

// DepositCounts is a free data retrieval call binding the contract method 0x174e300b.
//
// Solidity: function _depositCounts(bytes8 ) view returns(uint64)
func (_EthBr *EthBrCaller) DepositCounts(opts *bind.CallOpts, arg0 [8]byte) (uint64, error) {
	var out []interface{}
	err := _EthBr.contract.Call(opts, &out, "_depositCounts", arg0)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// DepositCounts is a free data retrieval call binding the contract method 0x174e300b.
//
// Solidity: function _depositCounts(bytes8 ) view returns(uint64)
func (_EthBr *EthBrSession) DepositCounts(arg0 [8]byte) (uint64, error) {
	return _EthBr.Contract.DepositCounts(&_EthBr.CallOpts, arg0)
}

// DepositCounts is a free data retrieval call binding the contract method 0x174e300b.
//
// Solidity: function _depositCounts(bytes8 ) view returns(uint64)
func (_EthBr *EthBrCallerSession) DepositCounts(arg0 [8]byte) (uint64, error) {
	return _EthBr.Contract.DepositCounts(&_EthBr.CallOpts, arg0)
}

// DepositRecords is a free data retrieval call binding the contract method 0xa48ace38.
//
// Solidity: function _depositRecords(uint64 , bytes8 ) view returns(bytes)
func (_EthBr *EthBrCaller) DepositRecords(opts *bind.CallOpts, arg0 uint64, arg1 [8]byte) ([]byte, error) {
	var out []interface{}
	err := _EthBr.contract.Call(opts, &out, "_depositRecords", arg0, arg1)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// DepositRecords is a free data retrieval call binding the contract method 0xa48ace38.
//
// Solidity: function _depositRecords(uint64 , bytes8 ) view returns(bytes)
func (_EthBr *EthBrSession) DepositRecords(arg0 uint64, arg1 [8]byte) ([]byte, error) {
	return _EthBr.Contract.DepositRecords(&_EthBr.CallOpts, arg0, arg1)
}

// DepositRecords is a free data retrieval call binding the contract method 0xa48ace38.
//
// Solidity: function _depositRecords(uint64 , bytes8 ) view returns(bytes)
func (_EthBr *EthBrCallerSession) DepositRecords(arg0 uint64, arg1 [8]byte) ([]byte, error) {
	return _EthBr.Contract.DepositRecords(&_EthBr.CallOpts, arg0, arg1)
}

// ExecutedProposals is a free data retrieval call binding the contract method 0x9c1650e0.
//
// Solidity: function _executedProposals(bytes32 , bytes32 ) view returns(bool)
func (_EthBr *EthBrCaller) ExecutedProposals(opts *bind.CallOpts, arg0 [32]byte, arg1 [32]byte) (bool, error) {
	var out []interface{}
	err := _EthBr.contract.Call(opts, &out, "_executedProposals", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ExecutedProposals is a free data retrieval call binding the contract method 0x9c1650e0.
//
// Solidity: function _executedProposals(bytes32 , bytes32 ) view returns(bool)
func (_EthBr *EthBrSession) ExecutedProposals(arg0 [32]byte, arg1 [32]byte) (bool, error) {
	return _EthBr.Contract.ExecutedProposals(&_EthBr.CallOpts, arg0, arg1)
}

// ExecutedProposals is a free data retrieval call binding the contract method 0x9c1650e0.
//
// Solidity: function _executedProposals(bytes32 , bytes32 ) view returns(bool)
func (_EthBr *EthBrCallerSession) ExecutedProposals(arg0 [32]byte, arg1 [32]byte) (bool, error) {
	return _EthBr.Contract.ExecutedProposals(&_EthBr.CallOpts, arg0, arg1)
}

// Fee is a free data retrieval call binding the contract method 0xc5b37c22.
//
// Solidity: function _fee() view returns(uint256)
func (_EthBr *EthBrCaller) Fee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EthBr.contract.Call(opts, &out, "_fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fee is a free data retrieval call binding the contract method 0xc5b37c22.
//
// Solidity: function _fee() view returns(uint256)
func (_EthBr *EthBrSession) Fee() (*big.Int, error) {
	return _EthBr.Contract.Fee(&_EthBr.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xc5b37c22.
//
// Solidity: function _fee() view returns(uint256)
func (_EthBr *EthBrCallerSession) Fee() (*big.Int, error) {
	return _EthBr.Contract.Fee(&_EthBr.CallOpts)
}

// ResourceIDToHandlerAddress is a free data retrieval call binding the contract method 0x84db809f.
//
// Solidity: function _resourceIDToHandlerAddress(bytes32 ) view returns(address)
func (_EthBr *EthBrCaller) ResourceIDToHandlerAddress(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _EthBr.contract.Call(opts, &out, "_resourceIDToHandlerAddress", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ResourceIDToHandlerAddress is a free data retrieval call binding the contract method 0x84db809f.
//
// Solidity: function _resourceIDToHandlerAddress(bytes32 ) view returns(address)
func (_EthBr *EthBrSession) ResourceIDToHandlerAddress(arg0 [32]byte) (common.Address, error) {
	return _EthBr.Contract.ResourceIDToHandlerAddress(&_EthBr.CallOpts, arg0)
}

// ResourceIDToHandlerAddress is a free data retrieval call binding the contract method 0x84db809f.
//
// Solidity: function _resourceIDToHandlerAddress(bytes32 ) view returns(address)
func (_EthBr *EthBrCallerSession) ResourceIDToHandlerAddress(arg0 [32]byte) (common.Address, error) {
	return _EthBr.Contract.ResourceIDToHandlerAddress(&_EthBr.CallOpts, arg0)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_EthBr *EthBrCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _EthBr.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_EthBr *EthBrSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _EthBr.Contract.GetRoleAdmin(&_EthBr.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_EthBr *EthBrCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _EthBr.Contract.GetRoleAdmin(&_EthBr.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_EthBr *EthBrCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _EthBr.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_EthBr *EthBrSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _EthBr.Contract.GetRoleMember(&_EthBr.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_EthBr *EthBrCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _EthBr.Contract.GetRoleMember(&_EthBr.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_EthBr *EthBrCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _EthBr.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_EthBr *EthBrSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _EthBr.Contract.GetRoleMemberCount(&_EthBr.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_EthBr *EthBrCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _EthBr.Contract.GetRoleMemberCount(&_EthBr.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_EthBr *EthBrCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _EthBr.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_EthBr *EthBrSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _EthBr.Contract.HasRole(&_EthBr.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_EthBr *EthBrCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _EthBr.Contract.HasRole(&_EthBr.CallOpts, role, account)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_EthBr *EthBrCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _EthBr.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_EthBr *EthBrSession) Paused() (bool, error) {
	return _EthBr.Contract.Paused(&_EthBr.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_EthBr *EthBrCallerSession) Paused() (bool, error) {
	return _EthBr.Contract.Paused(&_EthBr.CallOpts)
}

// Sub is a free data retrieval call binding the contract method 0xb67d77c5.
//
// Solidity: function sub(uint256 a, uint256 b) pure returns(uint256)
func (_EthBr *EthBrCaller) Sub(opts *bind.CallOpts, a *big.Int, b *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _EthBr.contract.Call(opts, &out, "sub", a, b)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Sub is a free data retrieval call binding the contract method 0xb67d77c5.
//
// Solidity: function sub(uint256 a, uint256 b) pure returns(uint256)
func (_EthBr *EthBrSession) Sub(a *big.Int, b *big.Int) (*big.Int, error) {
	return _EthBr.Contract.Sub(&_EthBr.CallOpts, a, b)
}

// Sub is a free data retrieval call binding the contract method 0xb67d77c5.
//
// Solidity: function sub(uint256 a, uint256 b) pure returns(uint256)
func (_EthBr *EthBrCallerSession) Sub(a *big.Int, b *big.Int) (*big.Int, error) {
	return _EthBr.Contract.Sub(&_EthBr.CallOpts, a, b)
}

// AdminChangeFee is a paid mutator transaction binding the contract method 0x91c404ac.
//
// Solidity: function adminChangeFee(uint256 newFee) returns()
func (_EthBr *EthBrTransactor) AdminChangeFee(opts *bind.TransactOpts, newFee *big.Int) (*types.Transaction, error) {
	return _EthBr.contract.Transact(opts, "adminChangeFee", newFee)
}

// AdminChangeFee is a paid mutator transaction binding the contract method 0x91c404ac.
//
// Solidity: function adminChangeFee(uint256 newFee) returns()
func (_EthBr *EthBrSession) AdminChangeFee(newFee *big.Int) (*types.Transaction, error) {
	return _EthBr.Contract.AdminChangeFee(&_EthBr.TransactOpts, newFee)
}

// AdminChangeFee is a paid mutator transaction binding the contract method 0x91c404ac.
//
// Solidity: function adminChangeFee(uint256 newFee) returns()
func (_EthBr *EthBrTransactorSession) AdminChangeFee(newFee *big.Int) (*types.Transaction, error) {
	return _EthBr.Contract.AdminChangeFee(&_EthBr.TransactOpts, newFee)
}

// AdminPauseTransfers is a paid mutator transaction binding the contract method 0x80ae1c28.
//
// Solidity: function adminPauseTransfers() returns()
func (_EthBr *EthBrTransactor) AdminPauseTransfers(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthBr.contract.Transact(opts, "adminPauseTransfers")
}

// AdminPauseTransfers is a paid mutator transaction binding the contract method 0x80ae1c28.
//
// Solidity: function adminPauseTransfers() returns()
func (_EthBr *EthBrSession) AdminPauseTransfers() (*types.Transaction, error) {
	return _EthBr.Contract.AdminPauseTransfers(&_EthBr.TransactOpts)
}

// AdminPauseTransfers is a paid mutator transaction binding the contract method 0x80ae1c28.
//
// Solidity: function adminPauseTransfers() returns()
func (_EthBr *EthBrTransactorSession) AdminPauseTransfers() (*types.Transaction, error) {
	return _EthBr.Contract.AdminPauseTransfers(&_EthBr.TransactOpts)
}

// AdminSetBackendSrv is a paid mutator transaction binding the contract method 0x06f3640b.
//
// Solidity: function adminSetBackendSrv(address newBackendSrv) returns()
func (_EthBr *EthBrTransactor) AdminSetBackendSrv(opts *bind.TransactOpts, newBackendSrv common.Address) (*types.Transaction, error) {
	return _EthBr.contract.Transact(opts, "adminSetBackendSrv", newBackendSrv)
}

// AdminSetBackendSrv is a paid mutator transaction binding the contract method 0x06f3640b.
//
// Solidity: function adminSetBackendSrv(address newBackendSrv) returns()
func (_EthBr *EthBrSession) AdminSetBackendSrv(newBackendSrv common.Address) (*types.Transaction, error) {
	return _EthBr.Contract.AdminSetBackendSrv(&_EthBr.TransactOpts, newBackendSrv)
}

// AdminSetBackendSrv is a paid mutator transaction binding the contract method 0x06f3640b.
//
// Solidity: function adminSetBackendSrv(address newBackendSrv) returns()
func (_EthBr *EthBrTransactorSession) AdminSetBackendSrv(newBackendSrv common.Address) (*types.Transaction, error) {
	return _EthBr.Contract.AdminSetBackendSrv(&_EthBr.TransactOpts, newBackendSrv)
}

// AdminSetBurnable is a paid mutator transaction binding the contract method 0x8c0c2631.
//
// Solidity: function adminSetBurnable(address handlerAddress, address tokenAddress) returns()
func (_EthBr *EthBrTransactor) AdminSetBurnable(opts *bind.TransactOpts, handlerAddress common.Address, tokenAddress common.Address) (*types.Transaction, error) {
	return _EthBr.contract.Transact(opts, "adminSetBurnable", handlerAddress, tokenAddress)
}

// AdminSetBurnable is a paid mutator transaction binding the contract method 0x8c0c2631.
//
// Solidity: function adminSetBurnable(address handlerAddress, address tokenAddress) returns()
func (_EthBr *EthBrSession) AdminSetBurnable(handlerAddress common.Address, tokenAddress common.Address) (*types.Transaction, error) {
	return _EthBr.Contract.AdminSetBurnable(&_EthBr.TransactOpts, handlerAddress, tokenAddress)
}

// AdminSetBurnable is a paid mutator transaction binding the contract method 0x8c0c2631.
//
// Solidity: function adminSetBurnable(address handlerAddress, address tokenAddress) returns()
func (_EthBr *EthBrTransactorSession) AdminSetBurnable(handlerAddress common.Address, tokenAddress common.Address) (*types.Transaction, error) {
	return _EthBr.Contract.AdminSetBurnable(&_EthBr.TransactOpts, handlerAddress, tokenAddress)
}

// AdminSetResource is a paid mutator transaction binding the contract method 0xcb10f215.
//
// Solidity: function adminSetResource(address handlerAddress, bytes32 resourceID, address tokenAddress) returns()
func (_EthBr *EthBrTransactor) AdminSetResource(opts *bind.TransactOpts, handlerAddress common.Address, resourceID [32]byte, tokenAddress common.Address) (*types.Transaction, error) {
	return _EthBr.contract.Transact(opts, "adminSetResource", handlerAddress, resourceID, tokenAddress)
}

// AdminSetResource is a paid mutator transaction binding the contract method 0xcb10f215.
//
// Solidity: function adminSetResource(address handlerAddress, bytes32 resourceID, address tokenAddress) returns()
func (_EthBr *EthBrSession) AdminSetResource(handlerAddress common.Address, resourceID [32]byte, tokenAddress common.Address) (*types.Transaction, error) {
	return _EthBr.Contract.AdminSetResource(&_EthBr.TransactOpts, handlerAddress, resourceID, tokenAddress)
}

// AdminSetResource is a paid mutator transaction binding the contract method 0xcb10f215.
//
// Solidity: function adminSetResource(address handlerAddress, bytes32 resourceID, address tokenAddress) returns()
func (_EthBr *EthBrTransactorSession) AdminSetResource(handlerAddress common.Address, resourceID [32]byte, tokenAddress common.Address) (*types.Transaction, error) {
	return _EthBr.Contract.AdminSetResource(&_EthBr.TransactOpts, handlerAddress, resourceID, tokenAddress)
}

// AdminUnpauseTransfers is a paid mutator transaction binding the contract method 0xffaac0eb.
//
// Solidity: function adminUnpauseTransfers() returns()
func (_EthBr *EthBrTransactor) AdminUnpauseTransfers(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthBr.contract.Transact(opts, "adminUnpauseTransfers")
}

// AdminUnpauseTransfers is a paid mutator transaction binding the contract method 0xffaac0eb.
//
// Solidity: function adminUnpauseTransfers() returns()
func (_EthBr *EthBrSession) AdminUnpauseTransfers() (*types.Transaction, error) {
	return _EthBr.Contract.AdminUnpauseTransfers(&_EthBr.TransactOpts)
}

// AdminUnpauseTransfers is a paid mutator transaction binding the contract method 0xffaac0eb.
//
// Solidity: function adminUnpauseTransfers() returns()
func (_EthBr *EthBrTransactorSession) AdminUnpauseTransfers() (*types.Transaction, error) {
	return _EthBr.Contract.AdminUnpauseTransfers(&_EthBr.TransactOpts)
}

// AdminWithdraw is a paid mutator transaction binding the contract method 0x780cf004.
//
// Solidity: function adminWithdraw(address handlerAddress, address tokenAddress, address recipient, uint256 amountOrTokenID) returns()
func (_EthBr *EthBrTransactor) AdminWithdraw(opts *bind.TransactOpts, handlerAddress common.Address, tokenAddress common.Address, recipient common.Address, amountOrTokenID *big.Int) (*types.Transaction, error) {
	return _EthBr.contract.Transact(opts, "adminWithdraw", handlerAddress, tokenAddress, recipient, amountOrTokenID)
}

// AdminWithdraw is a paid mutator transaction binding the contract method 0x780cf004.
//
// Solidity: function adminWithdraw(address handlerAddress, address tokenAddress, address recipient, uint256 amountOrTokenID) returns()
func (_EthBr *EthBrSession) AdminWithdraw(handlerAddress common.Address, tokenAddress common.Address, recipient common.Address, amountOrTokenID *big.Int) (*types.Transaction, error) {
	return _EthBr.Contract.AdminWithdraw(&_EthBr.TransactOpts, handlerAddress, tokenAddress, recipient, amountOrTokenID)
}

// AdminWithdraw is a paid mutator transaction binding the contract method 0x780cf004.
//
// Solidity: function adminWithdraw(address handlerAddress, address tokenAddress, address recipient, uint256 amountOrTokenID) returns()
func (_EthBr *EthBrTransactorSession) AdminWithdraw(handlerAddress common.Address, tokenAddress common.Address, recipient common.Address, amountOrTokenID *big.Int) (*types.Transaction, error) {
	return _EthBr.Contract.AdminWithdraw(&_EthBr.TransactOpts, handlerAddress, tokenAddress, recipient, amountOrTokenID)
}

// Deposit is a paid mutator transaction binding the contract method 0x92b7064c.
//
// Solidity: function deposit(bytes8 destinationChainID, bytes32 resourceID, uint256 amount, address recipientAddress) payable returns()
func (_EthBr *EthBrTransactor) Deposit(opts *bind.TransactOpts, destinationChainID [8]byte, resourceID [32]byte, amount *big.Int, recipientAddress common.Address) (*types.Transaction, error) {
	return _EthBr.contract.Transact(opts, "deposit", destinationChainID, resourceID, amount, recipientAddress)
}

// Deposit is a paid mutator transaction binding the contract method 0x92b7064c.
//
// Solidity: function deposit(bytes8 destinationChainID, bytes32 resourceID, uint256 amount, address recipientAddress) payable returns()
func (_EthBr *EthBrSession) Deposit(destinationChainID [8]byte, resourceID [32]byte, amount *big.Int, recipientAddress common.Address) (*types.Transaction, error) {
	return _EthBr.Contract.Deposit(&_EthBr.TransactOpts, destinationChainID, resourceID, amount, recipientAddress)
}

// Deposit is a paid mutator transaction binding the contract method 0x92b7064c.
//
// Solidity: function deposit(bytes8 destinationChainID, bytes32 resourceID, uint256 amount, address recipientAddress) payable returns()
func (_EthBr *EthBrTransactorSession) Deposit(destinationChainID [8]byte, resourceID [32]byte, amount *big.Int, recipientAddress common.Address) (*types.Transaction, error) {
	return _EthBr.Contract.Deposit(&_EthBr.TransactOpts, destinationChainID, resourceID, amount, recipientAddress)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0x0da380e5.
//
// Solidity: function executeProposal(bytes8 originChainID, bytes8 destinationChainID, uint64 depositNonce, bytes32 resourceID, address recipientAddress, uint256 amount) returns()
func (_EthBr *EthBrTransactor) ExecuteProposal(opts *bind.TransactOpts, originChainID [8]byte, destinationChainID [8]byte, depositNonce uint64, resourceID [32]byte, recipientAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EthBr.contract.Transact(opts, "executeProposal", originChainID, destinationChainID, depositNonce, resourceID, recipientAddress, amount)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0x0da380e5.
//
// Solidity: function executeProposal(bytes8 originChainID, bytes8 destinationChainID, uint64 depositNonce, bytes32 resourceID, address recipientAddress, uint256 amount) returns()
func (_EthBr *EthBrSession) ExecuteProposal(originChainID [8]byte, destinationChainID [8]byte, depositNonce uint64, resourceID [32]byte, recipientAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EthBr.Contract.ExecuteProposal(&_EthBr.TransactOpts, originChainID, destinationChainID, depositNonce, resourceID, recipientAddress, amount)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0x0da380e5.
//
// Solidity: function executeProposal(bytes8 originChainID, bytes8 destinationChainID, uint64 depositNonce, bytes32 resourceID, address recipientAddress, uint256 amount) returns()
func (_EthBr *EthBrTransactorSession) ExecuteProposal(originChainID [8]byte, destinationChainID [8]byte, depositNonce uint64, resourceID [32]byte, recipientAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EthBr.Contract.ExecuteProposal(&_EthBr.TransactOpts, originChainID, destinationChainID, depositNonce, resourceID, recipientAddress, amount)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_EthBr *EthBrTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _EthBr.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_EthBr *EthBrSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _EthBr.Contract.GrantRole(&_EthBr.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_EthBr *EthBrTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _EthBr.Contract.GrantRole(&_EthBr.TransactOpts, role, account)
}

// RenounceAdmin is a paid mutator transaction binding the contract method 0x5e1fab0f.
//
// Solidity: function renounceAdmin(address newAdmin) returns()
func (_EthBr *EthBrTransactor) RenounceAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _EthBr.contract.Transact(opts, "renounceAdmin", newAdmin)
}

// RenounceAdmin is a paid mutator transaction binding the contract method 0x5e1fab0f.
//
// Solidity: function renounceAdmin(address newAdmin) returns()
func (_EthBr *EthBrSession) RenounceAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _EthBr.Contract.RenounceAdmin(&_EthBr.TransactOpts, newAdmin)
}

// RenounceAdmin is a paid mutator transaction binding the contract method 0x5e1fab0f.
//
// Solidity: function renounceAdmin(address newAdmin) returns()
func (_EthBr *EthBrTransactorSession) RenounceAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _EthBr.Contract.RenounceAdmin(&_EthBr.TransactOpts, newAdmin)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_EthBr *EthBrTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _EthBr.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_EthBr *EthBrSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _EthBr.Contract.RenounceRole(&_EthBr.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_EthBr *EthBrTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _EthBr.Contract.RenounceRole(&_EthBr.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_EthBr *EthBrTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _EthBr.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_EthBr *EthBrSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _EthBr.Contract.RevokeRole(&_EthBr.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_EthBr *EthBrTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _EthBr.Contract.RevokeRole(&_EthBr.TransactOpts, role, account)
}

// EthBrDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the EthBr contract.
type EthBrDepositIterator struct {
	Event *EthBrDeposit // Event containing the contract specifics and raw log

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
func (it *EthBrDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthBrDeposit)
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
		it.Event = new(EthBrDeposit)
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
func (it *EthBrDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthBrDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthBrDeposit represents a Deposit event raised by the EthBr contract.
type EthBrDeposit struct {
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
// Solidity: event Deposit(bytes8 originChainID, bytes8 indexed destinationChainID, bytes32 indexed resourceID, uint64 indexed depositNonce, address depositor, address recipientAddress, address tokenAddress, uint256 amount, bytes32 dataHash)
func (_EthBr *EthBrFilterer) FilterDeposit(opts *bind.FilterOpts, destinationChainID [][8]byte, resourceID [][32]byte, depositNonce []uint64) (*EthBrDepositIterator, error) {

	var destinationChainIDRule []interface{}
	for _, destinationChainIDItem := range destinationChainID {
		destinationChainIDRule = append(destinationChainIDRule, destinationChainIDItem)
	}
	var resourceIDRule []interface{}
	for _, resourceIDItem := range resourceID {
		resourceIDRule = append(resourceIDRule, resourceIDItem)
	}
	var depositNonceRule []interface{}
	for _, depositNonceItem := range depositNonce {
		depositNonceRule = append(depositNonceRule, depositNonceItem)
	}

	logs, sub, err := _EthBr.contract.FilterLogs(opts, "Deposit", destinationChainIDRule, resourceIDRule, depositNonceRule)
	if err != nil {
		return nil, err
	}
	return &EthBrDepositIterator{contract: _EthBr.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x370525803ffa9a7c0e6adb3868e393dca45d8b42b2f62fd1f23ecfe99f6ce8fc.
//
// Solidity: event Deposit(bytes8 originChainID, bytes8 indexed destinationChainID, bytes32 indexed resourceID, uint64 indexed depositNonce, address depositor, address recipientAddress, address tokenAddress, uint256 amount, bytes32 dataHash)
func (_EthBr *EthBrFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *EthBrDeposit, destinationChainID [][8]byte, resourceID [][32]byte, depositNonce []uint64) (event.Subscription, error) {

	var destinationChainIDRule []interface{}
	for _, destinationChainIDItem := range destinationChainID {
		destinationChainIDRule = append(destinationChainIDRule, destinationChainIDItem)
	}
	var resourceIDRule []interface{}
	for _, resourceIDItem := range resourceID {
		resourceIDRule = append(resourceIDRule, resourceIDItem)
	}
	var depositNonceRule []interface{}
	for _, depositNonceItem := range depositNonce {
		depositNonceRule = append(depositNonceRule, depositNonceItem)
	}

	logs, sub, err := _EthBr.contract.WatchLogs(opts, "Deposit", destinationChainIDRule, resourceIDRule, depositNonceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthBrDeposit)
				if err := _EthBr.contract.UnpackLog(event, "Deposit", log); err != nil {
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
// Solidity: event Deposit(bytes8 originChainID, bytes8 indexed destinationChainID, bytes32 indexed resourceID, uint64 indexed depositNonce, address depositor, address recipientAddress, address tokenAddress, uint256 amount, bytes32 dataHash)
func (_EthBr *EthBrFilterer) ParseDeposit(log types.Log) (*EthBrDeposit, error) {
	event := new(EthBrDeposit)
	if err := _EthBr.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthBrPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the EthBr contract.
type EthBrPausedIterator struct {
	Event *EthBrPaused // Event containing the contract specifics and raw log

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
func (it *EthBrPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthBrPaused)
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
		it.Event = new(EthBrPaused)
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
func (it *EthBrPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthBrPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthBrPaused represents a Paused event raised by the EthBr contract.
type EthBrPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_EthBr *EthBrFilterer) FilterPaused(opts *bind.FilterOpts) (*EthBrPausedIterator, error) {

	logs, sub, err := _EthBr.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &EthBrPausedIterator{contract: _EthBr.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_EthBr *EthBrFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *EthBrPaused) (event.Subscription, error) {

	logs, sub, err := _EthBr.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthBrPaused)
				if err := _EthBr.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_EthBr *EthBrFilterer) ParsePaused(log types.Log) (*EthBrPaused, error) {
	event := new(EthBrPaused)
	if err := _EthBr.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthBrProposalEventIterator is returned from FilterProposalEvent and is used to iterate over the raw logs and unpacked data for ProposalEvent events raised by the EthBr contract.
type EthBrProposalEventIterator struct {
	Event *EthBrProposalEvent // Event containing the contract specifics and raw log

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
func (it *EthBrProposalEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthBrProposalEvent)
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
		it.Event = new(EthBrProposalEvent)
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
func (it *EthBrProposalEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthBrProposalEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthBrProposalEvent represents a ProposalEvent event raised by the EthBr contract.
type EthBrProposalEvent struct {
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
// Solidity: event ProposalEvent(bytes8 indexed originChainID, bytes8 indexed destinationChainID, address indexed recipientAddress, uint256 amount, uint64 depositNonce, uint8 status, bytes32 resourceID, bytes32 dataHash)
func (_EthBr *EthBrFilterer) FilterProposalEvent(opts *bind.FilterOpts, originChainID [][8]byte, destinationChainID [][8]byte, recipientAddress []common.Address) (*EthBrProposalEventIterator, error) {

	var originChainIDRule []interface{}
	for _, originChainIDItem := range originChainID {
		originChainIDRule = append(originChainIDRule, originChainIDItem)
	}
	var destinationChainIDRule []interface{}
	for _, destinationChainIDItem := range destinationChainID {
		destinationChainIDRule = append(destinationChainIDRule, destinationChainIDItem)
	}
	var recipientAddressRule []interface{}
	for _, recipientAddressItem := range recipientAddress {
		recipientAddressRule = append(recipientAddressRule, recipientAddressItem)
	}

	logs, sub, err := _EthBr.contract.FilterLogs(opts, "ProposalEvent", originChainIDRule, destinationChainIDRule, recipientAddressRule)
	if err != nil {
		return nil, err
	}
	return &EthBrProposalEventIterator{contract: _EthBr.contract, event: "ProposalEvent", logs: logs, sub: sub}, nil
}

// WatchProposalEvent is a free log subscription operation binding the contract event 0x9686dcabd0450cad86a88df15a9d35b08b35d1b08a19008df37cf8538c467516.
//
// Solidity: event ProposalEvent(bytes8 indexed originChainID, bytes8 indexed destinationChainID, address indexed recipientAddress, uint256 amount, uint64 depositNonce, uint8 status, bytes32 resourceID, bytes32 dataHash)
func (_EthBr *EthBrFilterer) WatchProposalEvent(opts *bind.WatchOpts, sink chan<- *EthBrProposalEvent, originChainID [][8]byte, destinationChainID [][8]byte, recipientAddress []common.Address) (event.Subscription, error) {

	var originChainIDRule []interface{}
	for _, originChainIDItem := range originChainID {
		originChainIDRule = append(originChainIDRule, originChainIDItem)
	}
	var destinationChainIDRule []interface{}
	for _, destinationChainIDItem := range destinationChainID {
		destinationChainIDRule = append(destinationChainIDRule, destinationChainIDItem)
	}
	var recipientAddressRule []interface{}
	for _, recipientAddressItem := range recipientAddress {
		recipientAddressRule = append(recipientAddressRule, recipientAddressItem)
	}

	logs, sub, err := _EthBr.contract.WatchLogs(opts, "ProposalEvent", originChainIDRule, destinationChainIDRule, recipientAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthBrProposalEvent)
				if err := _EthBr.contract.UnpackLog(event, "ProposalEvent", log); err != nil {
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
// Solidity: event ProposalEvent(bytes8 indexed originChainID, bytes8 indexed destinationChainID, address indexed recipientAddress, uint256 amount, uint64 depositNonce, uint8 status, bytes32 resourceID, bytes32 dataHash)
func (_EthBr *EthBrFilterer) ParseProposalEvent(log types.Log) (*EthBrProposalEvent, error) {
	event := new(EthBrProposalEvent)
	if err := _EthBr.contract.UnpackLog(event, "ProposalEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthBrRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the EthBr contract.
type EthBrRoleAdminChangedIterator struct {
	Event *EthBrRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *EthBrRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthBrRoleAdminChanged)
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
		it.Event = new(EthBrRoleAdminChanged)
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
func (it *EthBrRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthBrRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthBrRoleAdminChanged represents a RoleAdminChanged event raised by the EthBr contract.
type EthBrRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_EthBr *EthBrFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*EthBrRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _EthBr.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &EthBrRoleAdminChangedIterator{contract: _EthBr.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_EthBr *EthBrFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *EthBrRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _EthBr.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthBrRoleAdminChanged)
				if err := _EthBr.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_EthBr *EthBrFilterer) ParseRoleAdminChanged(log types.Log) (*EthBrRoleAdminChanged, error) {
	event := new(EthBrRoleAdminChanged)
	if err := _EthBr.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthBrRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the EthBr contract.
type EthBrRoleGrantedIterator struct {
	Event *EthBrRoleGranted // Event containing the contract specifics and raw log

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
func (it *EthBrRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthBrRoleGranted)
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
		it.Event = new(EthBrRoleGranted)
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
func (it *EthBrRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthBrRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthBrRoleGranted represents a RoleGranted event raised by the EthBr contract.
type EthBrRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_EthBr *EthBrFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*EthBrRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _EthBr.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &EthBrRoleGrantedIterator{contract: _EthBr.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_EthBr *EthBrFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *EthBrRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _EthBr.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthBrRoleGranted)
				if err := _EthBr.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_EthBr *EthBrFilterer) ParseRoleGranted(log types.Log) (*EthBrRoleGranted, error) {
	event := new(EthBrRoleGranted)
	if err := _EthBr.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthBrRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the EthBr contract.
type EthBrRoleRevokedIterator struct {
	Event *EthBrRoleRevoked // Event containing the contract specifics and raw log

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
func (it *EthBrRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthBrRoleRevoked)
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
		it.Event = new(EthBrRoleRevoked)
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
func (it *EthBrRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthBrRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthBrRoleRevoked represents a RoleRevoked event raised by the EthBr contract.
type EthBrRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_EthBr *EthBrFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*EthBrRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _EthBr.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &EthBrRoleRevokedIterator{contract: _EthBr.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_EthBr *EthBrFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *EthBrRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _EthBr.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthBrRoleRevoked)
				if err := _EthBr.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_EthBr *EthBrFilterer) ParseRoleRevoked(log types.Log) (*EthBrRoleRevoked, error) {
	event := new(EthBrRoleRevoked)
	if err := _EthBr.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthBrUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the EthBr contract.
type EthBrUnpausedIterator struct {
	Event *EthBrUnpaused // Event containing the contract specifics and raw log

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
func (it *EthBrUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthBrUnpaused)
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
		it.Event = new(EthBrUnpaused)
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
func (it *EthBrUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthBrUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthBrUnpaused represents a Unpaused event raised by the EthBr contract.
type EthBrUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_EthBr *EthBrFilterer) FilterUnpaused(opts *bind.FilterOpts) (*EthBrUnpausedIterator, error) {

	logs, sub, err := _EthBr.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &EthBrUnpausedIterator{contract: _EthBr.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_EthBr *EthBrFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *EthBrUnpaused) (event.Subscription, error) {

	logs, sub, err := _EthBr.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthBrUnpaused)
				if err := _EthBr.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_EthBr *EthBrFilterer) ParseUnpaused(log types.Log) (*EthBrUnpaused, error) {
	event := new(EthBrUnpaused)
	if err := _EthBr.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
