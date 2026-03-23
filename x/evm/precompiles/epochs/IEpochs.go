// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package epochs

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

// EpochInfo is an auto generated low-level Go binding around an user-defined struct.
type EpochInfo struct {
	Identifier              string
	StartTime               int64
	Duration                int64
	CurrentEpoch            int64
	CurrentEpochStartTime   int64
	EpochCountingStarted    bool
	CurrentEpochStartHeight int64
}

// PageRequest is an auto generated low-level Go binding around an user-defined struct.
type PageRequest struct {
	Key        []byte
	Offset     uint64
	Limit      uint64
	CountTotal bool
	Reverse    bool
}

// PageResponse is an auto generated low-level Go binding around an user-defined struct.
type PageResponse struct {
	NextKey []byte
	Total   uint64
}

// IEpochsMetaData contains all meta data concerning the IEpochs contract.
var IEpochsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"identifier\",\"type\":\"string\"}],\"name\":\"currentEpoch\",\"outputs\":[{\"internalType\":\"int64\",\"name\":\"currentEpoch\",\"type\":\"int64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"offset\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"limit\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"countTotal\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"reverse\",\"type\":\"bool\"}],\"internalType\":\"structPageRequest\",\"name\":\"pagination\",\"type\":\"tuple\"}],\"name\":\"epochInfos\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"identifier\",\"type\":\"string\"},{\"internalType\":\"int64\",\"name\":\"startTime\",\"type\":\"int64\"},{\"internalType\":\"int64\",\"name\":\"duration\",\"type\":\"int64\"},{\"internalType\":\"int64\",\"name\":\"currentEpoch\",\"type\":\"int64\"},{\"internalType\":\"int64\",\"name\":\"currentEpochStartTime\",\"type\":\"int64\"},{\"internalType\":\"bool\",\"name\":\"epochCountingStarted\",\"type\":\"bool\"},{\"internalType\":\"int64\",\"name\":\"currentEpochStartHeight\",\"type\":\"int64\"}],\"internalType\":\"structEpochInfo[]\",\"name\":\"epochs\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"nextKey\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"total\",\"type\":\"uint64\"}],\"internalType\":\"structPageResponse\",\"name\":\"pageResponse\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IEpochsABI is the input ABI used to generate the binding from.
// Deprecated: Use IEpochsMetaData.ABI instead.
var IEpochsABI = IEpochsMetaData.ABI

// IEpochs is an auto generated Go binding around an Ethereum contract.
type IEpochs struct {
	IEpochsCaller     // Read-only binding to the contract
	IEpochsTransactor // Write-only binding to the contract
	IEpochsFilterer   // Log filterer for contract events
}

// IEpochsCaller is an auto generated read-only Go binding around an Ethereum contract.
type IEpochsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IEpochsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IEpochsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IEpochsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IEpochsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IEpochsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IEpochsSession struct {
	Contract     *IEpochs          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IEpochsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IEpochsCallerSession struct {
	Contract *IEpochsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// IEpochsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IEpochsTransactorSession struct {
	Contract     *IEpochsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IEpochsRaw is an auto generated low-level Go binding around an Ethereum contract.
type IEpochsRaw struct {
	Contract *IEpochs // Generic contract binding to access the raw methods on
}

// IEpochsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IEpochsCallerRaw struct {
	Contract *IEpochsCaller // Generic read-only contract binding to access the raw methods on
}

// IEpochsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IEpochsTransactorRaw struct {
	Contract *IEpochsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIEpochs creates a new instance of IEpochs, bound to a specific deployed contract.
func NewIEpochs(address common.Address, backend bind.ContractBackend) (*IEpochs, error) {
	contract, err := bindIEpochs(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IEpochs{IEpochsCaller: IEpochsCaller{contract: contract}, IEpochsTransactor: IEpochsTransactor{contract: contract}, IEpochsFilterer: IEpochsFilterer{contract: contract}}, nil
}

// NewIEpochsCaller creates a new read-only instance of IEpochs, bound to a specific deployed contract.
func NewIEpochsCaller(address common.Address, caller bind.ContractCaller) (*IEpochsCaller, error) {
	contract, err := bindIEpochs(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IEpochsCaller{contract: contract}, nil
}

// NewIEpochsTransactor creates a new write-only instance of IEpochs, bound to a specific deployed contract.
func NewIEpochsTransactor(address common.Address, transactor bind.ContractTransactor) (*IEpochsTransactor, error) {
	contract, err := bindIEpochs(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IEpochsTransactor{contract: contract}, nil
}

// NewIEpochsFilterer creates a new log filterer instance of IEpochs, bound to a specific deployed contract.
func NewIEpochsFilterer(address common.Address, filterer bind.ContractFilterer) (*IEpochsFilterer, error) {
	contract, err := bindIEpochs(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IEpochsFilterer{contract: contract}, nil
}

// bindIEpochs binds a generic wrapper to an already deployed contract.
func bindIEpochs(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IEpochsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IEpochs *IEpochsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IEpochs.Contract.IEpochsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IEpochs *IEpochsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IEpochs.Contract.IEpochsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IEpochs *IEpochsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IEpochs.Contract.IEpochsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IEpochs *IEpochsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IEpochs.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IEpochs *IEpochsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IEpochs.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IEpochs *IEpochsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IEpochs.Contract.contract.Transact(opts, method, params...)
}

// CurrentEpoch is a free data retrieval call binding the contract method 0x75d11314.
//
// Solidity: function currentEpoch(string identifier) view returns(int64 currentEpoch)
func (_IEpochs *IEpochsCaller) CurrentEpoch(opts *bind.CallOpts, identifier string) (int64, error) {
	var out []interface{}
	err := _IEpochs.contract.Call(opts, &out, "currentEpoch", identifier)

	if err != nil {
		return *new(int64), err
	}

	out0 := *abi.ConvertType(out[0], new(int64)).(*int64)

	return out0, err

}

// CurrentEpoch is a free data retrieval call binding the contract method 0x75d11314.
//
// Solidity: function currentEpoch(string identifier) view returns(int64 currentEpoch)
func (_IEpochs *IEpochsSession) CurrentEpoch(identifier string) (int64, error) {
	return _IEpochs.Contract.CurrentEpoch(&_IEpochs.CallOpts, identifier)
}

// CurrentEpoch is a free data retrieval call binding the contract method 0x75d11314.
//
// Solidity: function currentEpoch(string identifier) view returns(int64 currentEpoch)
func (_IEpochs *IEpochsCallerSession) CurrentEpoch(identifier string) (int64, error) {
	return _IEpochs.Contract.CurrentEpoch(&_IEpochs.CallOpts, identifier)
}

// EpochInfos is a free data retrieval call binding the contract method 0x5fa28a21.
//
// Solidity: function epochInfos((bytes,uint64,uint64,bool,bool) pagination) view returns((string,int64,int64,int64,int64,bool,int64)[] epochs, (bytes,uint64) pageResponse)
func (_IEpochs *IEpochsCaller) EpochInfos(opts *bind.CallOpts, pagination PageRequest) (struct {
	Epochs       []EpochInfo
	PageResponse PageResponse
}, error) {
	var out []interface{}
	err := _IEpochs.contract.Call(opts, &out, "epochInfos", pagination)

	outstruct := new(struct {
		Epochs       []EpochInfo
		PageResponse PageResponse
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Epochs = *abi.ConvertType(out[0], new([]EpochInfo)).(*[]EpochInfo)
	outstruct.PageResponse = *abi.ConvertType(out[1], new(PageResponse)).(*PageResponse)

	return *outstruct, err

}

// EpochInfos is a free data retrieval call binding the contract method 0x5fa28a21.
//
// Solidity: function epochInfos((bytes,uint64,uint64,bool,bool) pagination) view returns((string,int64,int64,int64,int64,bool,int64)[] epochs, (bytes,uint64) pageResponse)
func (_IEpochs *IEpochsSession) EpochInfos(pagination PageRequest) (struct {
	Epochs       []EpochInfo
	PageResponse PageResponse
}, error) {
	return _IEpochs.Contract.EpochInfos(&_IEpochs.CallOpts, pagination)
}

// EpochInfos is a free data retrieval call binding the contract method 0x5fa28a21.
//
// Solidity: function epochInfos((bytes,uint64,uint64,bool,bool) pagination) view returns((string,int64,int64,int64,int64,bool,int64)[] epochs, (bytes,uint64) pageResponse)
func (_IEpochs *IEpochsCallerSession) EpochInfos(pagination PageRequest) (struct {
	Epochs       []EpochInfo
	PageResponse PageResponse
}, error) {
	return _IEpochs.Contract.EpochInfos(&_IEpochs.CallOpts, pagination)
}
