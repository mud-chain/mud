// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package inflation

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
	_ = abi.ConvertType
)

// Dec is an auto generated low-level Go binding around an user-defined struct.
type Dec struct {
	Amount    *big.Int
	Precision uint8
}

// DecCoin is an auto generated low-level Go binding around an user-defined struct.
type DecCoin struct {
	Denom     string
	Amount    *big.Int
	Precision uint8
}

// Params is an auto generated low-level Go binding around an user-defined struct.
type Params struct {
	StakingRewards  Dec
	CommunityPool   Dec
	EnableInflation bool
	InflationMax    Dec
	InflationDecay  Dec
}

// IInflationMetaData contains all meta data concerning the IInflation contract.
var IInflationMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"circulatingSupply\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"precision\",\"type\":\"uint8\"}],\"internalType\":\"structDecCoin\",\"name\":\"circulatingSupply\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"epochMintProvision\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"precision\",\"type\":\"uint8\"}],\"internalType\":\"structDecCoin\",\"name\":\"epochMintProvision\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"inflationRate\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"precision\",\"type\":\"uint8\"}],\"internalType\":\"structDec\",\"name\":\"inflationRate\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"params\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"precision\",\"type\":\"uint8\"}],\"internalType\":\"structDec\",\"name\":\"stakingRewards\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"precision\",\"type\":\"uint8\"}],\"internalType\":\"structDec\",\"name\":\"communityPool\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"enableInflation\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"precision\",\"type\":\"uint8\"}],\"internalType\":\"structDec\",\"name\":\"inflationMax\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"precision\",\"type\":\"uint8\"}],\"internalType\":\"structDec\",\"name\":\"inflationDecay\",\"type\":\"tuple\"}],\"internalType\":\"structParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"period\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"period\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"skippedEpochs\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"skippedEpochs\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IInflationABI is the input ABI used to generate the binding from.
// Deprecated: Use IInflationMetaData.ABI instead.
var IInflationABI = IInflationMetaData.ABI

// IInflation is an auto generated Go binding around an Ethereum contract.
type IInflation struct {
	IInflationCaller     // Read-only binding to the contract
	IInflationTransactor // Write-only binding to the contract
	IInflationFilterer   // Log filterer for contract events
}

// IInflationCaller is an auto generated read-only Go binding around an Ethereum contract.
type IInflationCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IInflationTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IInflationTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IInflationFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IInflationFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IInflationSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IInflationSession struct {
	Contract     *IInflation       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IInflationCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IInflationCallerSession struct {
	Contract *IInflationCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// IInflationTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IInflationTransactorSession struct {
	Contract     *IInflationTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// IInflationRaw is an auto generated low-level Go binding around an Ethereum contract.
type IInflationRaw struct {
	Contract *IInflation // Generic contract binding to access the raw methods on
}

// IInflationCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IInflationCallerRaw struct {
	Contract *IInflationCaller // Generic read-only contract binding to access the raw methods on
}

// IInflationTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IInflationTransactorRaw struct {
	Contract *IInflationTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIInflation creates a new instance of IInflation, bound to a specific deployed contract.
func NewIInflation(address common.Address, backend bind.ContractBackend) (*IInflation, error) {
	contract, err := bindIInflation(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IInflation{IInflationCaller: IInflationCaller{contract: contract}, IInflationTransactor: IInflationTransactor{contract: contract}, IInflationFilterer: IInflationFilterer{contract: contract}}, nil
}

// NewIInflationCaller creates a new read-only instance of IInflation, bound to a specific deployed contract.
func NewIInflationCaller(address common.Address, caller bind.ContractCaller) (*IInflationCaller, error) {
	contract, err := bindIInflation(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IInflationCaller{contract: contract}, nil
}

// NewIInflationTransactor creates a new write-only instance of IInflation, bound to a specific deployed contract.
func NewIInflationTransactor(address common.Address, transactor bind.ContractTransactor) (*IInflationTransactor, error) {
	contract, err := bindIInflation(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IInflationTransactor{contract: contract}, nil
}

// NewIInflationFilterer creates a new log filterer instance of IInflation, bound to a specific deployed contract.
func NewIInflationFilterer(address common.Address, filterer bind.ContractFilterer) (*IInflationFilterer, error) {
	contract, err := bindIInflation(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IInflationFilterer{contract: contract}, nil
}

// bindIInflation binds a generic wrapper to an already deployed contract.
func bindIInflation(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IInflationMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IInflation *IInflationRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IInflation.Contract.IInflationCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IInflation *IInflationRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IInflation.Contract.IInflationTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IInflation *IInflationRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IInflation.Contract.IInflationTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IInflation *IInflationCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IInflation.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IInflation *IInflationTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IInflation.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IInflation *IInflationTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IInflation.Contract.contract.Transact(opts, method, params...)
}

// CirculatingSupply is a free data retrieval call binding the contract method 0x9358928b.
//
// Solidity: function circulatingSupply() view returns((string,uint256,uint8) circulatingSupply)
func (_IInflation *IInflationCaller) CirculatingSupply(opts *bind.CallOpts) (DecCoin, error) {
	var out []interface{}
	err := _IInflation.contract.Call(opts, &out, "circulatingSupply")

	if err != nil {
		return *new(DecCoin), err
	}

	out0 := *abi.ConvertType(out[0], new(DecCoin)).(*DecCoin)

	return out0, err

}

// CirculatingSupply is a free data retrieval call binding the contract method 0x9358928b.
//
// Solidity: function circulatingSupply() view returns((string,uint256,uint8) circulatingSupply)
func (_IInflation *IInflationSession) CirculatingSupply() (DecCoin, error) {
	return _IInflation.Contract.CirculatingSupply(&_IInflation.CallOpts)
}

// CirculatingSupply is a free data retrieval call binding the contract method 0x9358928b.
//
// Solidity: function circulatingSupply() view returns((string,uint256,uint8) circulatingSupply)
func (_IInflation *IInflationCallerSession) CirculatingSupply() (DecCoin, error) {
	return _IInflation.Contract.CirculatingSupply(&_IInflation.CallOpts)
}

// EpochMintProvision is a free data retrieval call binding the contract method 0x49627060.
//
// Solidity: function epochMintProvision() view returns((string,uint256,uint8) epochMintProvision)
func (_IInflation *IInflationCaller) EpochMintProvision(opts *bind.CallOpts) (DecCoin, error) {
	var out []interface{}
	err := _IInflation.contract.Call(opts, &out, "epochMintProvision")

	if err != nil {
		return *new(DecCoin), err
	}

	out0 := *abi.ConvertType(out[0], new(DecCoin)).(*DecCoin)

	return out0, err

}

// EpochMintProvision is a free data retrieval call binding the contract method 0x49627060.
//
// Solidity: function epochMintProvision() view returns((string,uint256,uint8) epochMintProvision)
func (_IInflation *IInflationSession) EpochMintProvision() (DecCoin, error) {
	return _IInflation.Contract.EpochMintProvision(&_IInflation.CallOpts)
}

// EpochMintProvision is a free data retrieval call binding the contract method 0x49627060.
//
// Solidity: function epochMintProvision() view returns((string,uint256,uint8) epochMintProvision)
func (_IInflation *IInflationCallerSession) EpochMintProvision() (DecCoin, error) {
	return _IInflation.Contract.EpochMintProvision(&_IInflation.CallOpts)
}

// InflationRate is a free data retrieval call binding the contract method 0x31f9e35b.
//
// Solidity: function inflationRate() view returns((uint256,uint8) inflationRate)
func (_IInflation *IInflationCaller) InflationRate(opts *bind.CallOpts) (Dec, error) {
	var out []interface{}
	err := _IInflation.contract.Call(opts, &out, "inflationRate")

	if err != nil {
		return *new(Dec), err
	}

	out0 := *abi.ConvertType(out[0], new(Dec)).(*Dec)

	return out0, err

}

// InflationRate is a free data retrieval call binding the contract method 0x31f9e35b.
//
// Solidity: function inflationRate() view returns((uint256,uint8) inflationRate)
func (_IInflation *IInflationSession) InflationRate() (Dec, error) {
	return _IInflation.Contract.InflationRate(&_IInflation.CallOpts)
}

// InflationRate is a free data retrieval call binding the contract method 0x31f9e35b.
//
// Solidity: function inflationRate() view returns((uint256,uint8) inflationRate)
func (_IInflation *IInflationCallerSession) InflationRate() (Dec, error) {
	return _IInflation.Contract.InflationRate(&_IInflation.CallOpts)
}

// Params is a free data retrieval call binding the contract method 0xcff0ab96.
//
// Solidity: function params() view returns(((uint256,uint8),(uint256,uint8),bool,(uint256,uint8),(uint256,uint8)) params)
func (_IInflation *IInflationCaller) Params(opts *bind.CallOpts) (Params, error) {
	var out []interface{}
	err := _IInflation.contract.Call(opts, &out, "params")

	if err != nil {
		return *new(Params), err
	}

	out0 := *abi.ConvertType(out[0], new(Params)).(*Params)

	return out0, err

}

// Params is a free data retrieval call binding the contract method 0xcff0ab96.
//
// Solidity: function params() view returns(((uint256,uint8),(uint256,uint8),bool,(uint256,uint8),(uint256,uint8)) params)
func (_IInflation *IInflationSession) Params() (Params, error) {
	return _IInflation.Contract.Params(&_IInflation.CallOpts)
}

// Params is a free data retrieval call binding the contract method 0xcff0ab96.
//
// Solidity: function params() view returns(((uint256,uint8),(uint256,uint8),bool,(uint256,uint8),(uint256,uint8)) params)
func (_IInflation *IInflationCallerSession) Params() (Params, error) {
	return _IInflation.Contract.Params(&_IInflation.CallOpts)
}

// Period is a free data retrieval call binding the contract method 0xef78d4fd.
//
// Solidity: function period() view returns(uint64 period)
func (_IInflation *IInflationCaller) Period(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _IInflation.contract.Call(opts, &out, "period")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// Period is a free data retrieval call binding the contract method 0xef78d4fd.
//
// Solidity: function period() view returns(uint64 period)
func (_IInflation *IInflationSession) Period() (uint64, error) {
	return _IInflation.Contract.Period(&_IInflation.CallOpts)
}

// Period is a free data retrieval call binding the contract method 0xef78d4fd.
//
// Solidity: function period() view returns(uint64 period)
func (_IInflation *IInflationCallerSession) Period() (uint64, error) {
	return _IInflation.Contract.Period(&_IInflation.CallOpts)
}

// SkippedEpochs is a free data retrieval call binding the contract method 0xa4c586d4.
//
// Solidity: function skippedEpochs() view returns(uint64 skippedEpochs)
func (_IInflation *IInflationCaller) SkippedEpochs(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _IInflation.contract.Call(opts, &out, "skippedEpochs")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// SkippedEpochs is a free data retrieval call binding the contract method 0xa4c586d4.
//
// Solidity: function skippedEpochs() view returns(uint64 skippedEpochs)
func (_IInflation *IInflationSession) SkippedEpochs() (uint64, error) {
	return _IInflation.Contract.SkippedEpochs(&_IInflation.CallOpts)
}

// SkippedEpochs is a free data retrieval call binding the contract method 0xa4c586d4.
//
// Solidity: function skippedEpochs() view returns(uint64 skippedEpochs)
func (_IInflation *IInflationCallerSession) SkippedEpochs() (uint64, error) {
	return _IInflation.Contract.SkippedEpochs(&_IInflation.CallOpts)
}
