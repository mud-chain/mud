// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package auth

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

// Account is an auto generated low-level Go binding around an user-defined struct.
type Account struct {
	Addr          common.Address
	PubKey        string
	AccountNumber uint64
	Sequence      uint64
}

// ModuleAccount is an auto generated low-level Go binding around an user-defined struct.
type ModuleAccount struct {
	Addr          common.Address
	PubKey        string
	AccountNumber uint64
	Sequence      uint64
	Name          string
	Permissions   []string
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

// Params is an auto generated low-level Go binding around an user-defined struct.
type Params struct {
	MaxMemoCharacters      uint64
	TxSigLimit             uint64
	TxSizeCostPerByte      uint64
	SigVerifyCostEd25519   uint64
	SigVerifyCostSecp256k1 uint64
}

// IAuthMetaData contains all meta data concerning the IAuth contract.
var IAuthMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"account\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"pubKey\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"accountNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"}],\"internalType\":\"structAccount\",\"name\":\"account\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int64\",\"name\":\"id\",\"type\":\"int64\"}],\"name\":\"accountAddressByID\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"offset\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"limit\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"countTotal\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"reverse\",\"type\":\"bool\"}],\"internalType\":\"structPageRequest\",\"name\":\"pageRequest\",\"type\":\"tuple\"}],\"name\":\"accounts\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"pubKey\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"accountNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"}],\"internalType\":\"structAccount[]\",\"name\":\"accounts\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"nextKey\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"total\",\"type\":\"uint64\"}],\"internalType\":\"structPageResponse\",\"name\":\"pageResponse\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"moduleAccountByName\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"pubKey\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"accountNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"permissions\",\"type\":\"string[]\"}],\"internalType\":\"structModuleAccount\",\"name\":\"moduleAccount\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"moduleAccounts\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"pubKey\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"accountNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"permissions\",\"type\":\"string[]\"}],\"internalType\":\"structModuleAccount[]\",\"name\":\"moduleAccounts\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"params\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"maxMemoCharacters\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"txSigLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"txSizeCostPerByte\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"sigVerifyCostEd25519\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"sigVerifyCostSecp256k1\",\"type\":\"uint64\"}],\"internalType\":\"structParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IAuthABI is the input ABI used to generate the binding from.
// Deprecated: Use IAuthMetaData.ABI instead.
var IAuthABI = IAuthMetaData.ABI

// IAuth is an auto generated Go binding around an Ethereum contract.
type IAuth struct {
	IAuthCaller     // Read-only binding to the contract
	IAuthTransactor // Write-only binding to the contract
	IAuthFilterer   // Log filterer for contract events
}

// IAuthCaller is an auto generated read-only Go binding around an Ethereum contract.
type IAuthCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAuthTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IAuthTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAuthFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IAuthFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAuthSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IAuthSession struct {
	Contract     *IAuth            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IAuthCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IAuthCallerSession struct {
	Contract *IAuthCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IAuthTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IAuthTransactorSession struct {
	Contract     *IAuthTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IAuthRaw is an auto generated low-level Go binding around an Ethereum contract.
type IAuthRaw struct {
	Contract *IAuth // Generic contract binding to access the raw methods on
}

// IAuthCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IAuthCallerRaw struct {
	Contract *IAuthCaller // Generic read-only contract binding to access the raw methods on
}

// IAuthTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IAuthTransactorRaw struct {
	Contract *IAuthTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIAuth creates a new instance of IAuth, bound to a specific deployed contract.
func NewIAuth(address common.Address, backend bind.ContractBackend) (*IAuth, error) {
	contract, err := bindIAuth(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IAuth{IAuthCaller: IAuthCaller{contract: contract}, IAuthTransactor: IAuthTransactor{contract: contract}, IAuthFilterer: IAuthFilterer{contract: contract}}, nil
}

// NewIAuthCaller creates a new read-only instance of IAuth, bound to a specific deployed contract.
func NewIAuthCaller(address common.Address, caller bind.ContractCaller) (*IAuthCaller, error) {
	contract, err := bindIAuth(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IAuthCaller{contract: contract}, nil
}

// NewIAuthTransactor creates a new write-only instance of IAuth, bound to a specific deployed contract.
func NewIAuthTransactor(address common.Address, transactor bind.ContractTransactor) (*IAuthTransactor, error) {
	contract, err := bindIAuth(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IAuthTransactor{contract: contract}, nil
}

// NewIAuthFilterer creates a new log filterer instance of IAuth, bound to a specific deployed contract.
func NewIAuthFilterer(address common.Address, filterer bind.ContractFilterer) (*IAuthFilterer, error) {
	contract, err := bindIAuth(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IAuthFilterer{contract: contract}, nil
}

// bindIAuth binds a generic wrapper to an already deployed contract.
func bindIAuth(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IAuthMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAuth *IAuthRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAuth.Contract.IAuthCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAuth *IAuthRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAuth.Contract.IAuthTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAuth *IAuthRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAuth.Contract.IAuthTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAuth *IAuthCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAuth.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAuth *IAuthTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAuth.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAuth *IAuthTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAuth.Contract.contract.Transact(opts, method, params...)
}

// Account is a free data retrieval call binding the contract method 0x73b9aa91.
//
// Solidity: function account(address addr) view returns((address,string,uint64,uint64) account)
func (_IAuth *IAuthCaller) Account(opts *bind.CallOpts, addr common.Address) (Account, error) {
	var out []interface{}
	err := _IAuth.contract.Call(opts, &out, "account", addr)

	if err != nil {
		return *new(Account), err
	}

	out0 := *abi.ConvertType(out[0], new(Account)).(*Account)

	return out0, err

}

// Account is a free data retrieval call binding the contract method 0x73b9aa91.
//
// Solidity: function account(address addr) view returns((address,string,uint64,uint64) account)
func (_IAuth *IAuthSession) Account(addr common.Address) (Account, error) {
	return _IAuth.Contract.Account(&_IAuth.CallOpts, addr)
}

// Account is a free data retrieval call binding the contract method 0x73b9aa91.
//
// Solidity: function account(address addr) view returns((address,string,uint64,uint64) account)
func (_IAuth *IAuthCallerSession) Account(addr common.Address) (Account, error) {
	return _IAuth.Contract.Account(&_IAuth.CallOpts, addr)
}

// AccountAddressByID is a free data retrieval call binding the contract method 0x2b51cd72.
//
// Solidity: function accountAddressByID(int64 id) view returns(address addr)
func (_IAuth *IAuthCaller) AccountAddressByID(opts *bind.CallOpts, id int64) (common.Address, error) {
	var out []interface{}
	err := _IAuth.contract.Call(opts, &out, "accountAddressByID", id)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AccountAddressByID is a free data retrieval call binding the contract method 0x2b51cd72.
//
// Solidity: function accountAddressByID(int64 id) view returns(address addr)
func (_IAuth *IAuthSession) AccountAddressByID(id int64) (common.Address, error) {
	return _IAuth.Contract.AccountAddressByID(&_IAuth.CallOpts, id)
}

// AccountAddressByID is a free data retrieval call binding the contract method 0x2b51cd72.
//
// Solidity: function accountAddressByID(int64 id) view returns(address addr)
func (_IAuth *IAuthCallerSession) AccountAddressByID(id int64) (common.Address, error) {
	return _IAuth.Contract.AccountAddressByID(&_IAuth.CallOpts, id)
}

// Accounts is a free data retrieval call binding the contract method 0x3f701b99.
//
// Solidity: function accounts((bytes,uint64,uint64,bool,bool) pageRequest) view returns((address,string,uint64,uint64)[] accounts, (bytes,uint64) pageResponse)
func (_IAuth *IAuthCaller) Accounts(opts *bind.CallOpts, pageRequest PageRequest) (struct {
	Accounts     []Account
	PageResponse PageResponse
}, error) {
	var out []interface{}
	err := _IAuth.contract.Call(opts, &out, "accounts", pageRequest)

	outstruct := new(struct {
		Accounts     []Account
		PageResponse PageResponse
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Accounts = *abi.ConvertType(out[0], new([]Account)).(*[]Account)
	outstruct.PageResponse = *abi.ConvertType(out[1], new(PageResponse)).(*PageResponse)

	return *outstruct, err

}

// Accounts is a free data retrieval call binding the contract method 0x3f701b99.
//
// Solidity: function accounts((bytes,uint64,uint64,bool,bool) pageRequest) view returns((address,string,uint64,uint64)[] accounts, (bytes,uint64) pageResponse)
func (_IAuth *IAuthSession) Accounts(pageRequest PageRequest) (struct {
	Accounts     []Account
	PageResponse PageResponse
}, error) {
	return _IAuth.Contract.Accounts(&_IAuth.CallOpts, pageRequest)
}

// Accounts is a free data retrieval call binding the contract method 0x3f701b99.
//
// Solidity: function accounts((bytes,uint64,uint64,bool,bool) pageRequest) view returns((address,string,uint64,uint64)[] accounts, (bytes,uint64) pageResponse)
func (_IAuth *IAuthCallerSession) Accounts(pageRequest PageRequest) (struct {
	Accounts     []Account
	PageResponse PageResponse
}, error) {
	return _IAuth.Contract.Accounts(&_IAuth.CallOpts, pageRequest)
}

// ModuleAccountByName is a free data retrieval call binding the contract method 0x61dd1dd4.
//
// Solidity: function moduleAccountByName(string name) view returns((address,string,uint64,uint64,string,string[]) moduleAccount)
func (_IAuth *IAuthCaller) ModuleAccountByName(opts *bind.CallOpts, name string) (ModuleAccount, error) {
	var out []interface{}
	err := _IAuth.contract.Call(opts, &out, "moduleAccountByName", name)

	if err != nil {
		return *new(ModuleAccount), err
	}

	out0 := *abi.ConvertType(out[0], new(ModuleAccount)).(*ModuleAccount)

	return out0, err

}

// ModuleAccountByName is a free data retrieval call binding the contract method 0x61dd1dd4.
//
// Solidity: function moduleAccountByName(string name) view returns((address,string,uint64,uint64,string,string[]) moduleAccount)
func (_IAuth *IAuthSession) ModuleAccountByName(name string) (ModuleAccount, error) {
	return _IAuth.Contract.ModuleAccountByName(&_IAuth.CallOpts, name)
}

// ModuleAccountByName is a free data retrieval call binding the contract method 0x61dd1dd4.
//
// Solidity: function moduleAccountByName(string name) view returns((address,string,uint64,uint64,string,string[]) moduleAccount)
func (_IAuth *IAuthCallerSession) ModuleAccountByName(name string) (ModuleAccount, error) {
	return _IAuth.Contract.ModuleAccountByName(&_IAuth.CallOpts, name)
}

// ModuleAccounts is a free data retrieval call binding the contract method 0x26d6d3aa.
//
// Solidity: function moduleAccounts() view returns((address,string,uint64,uint64,string,string[])[] moduleAccounts)
func (_IAuth *IAuthCaller) ModuleAccounts(opts *bind.CallOpts) ([]ModuleAccount, error) {
	var out []interface{}
	err := _IAuth.contract.Call(opts, &out, "moduleAccounts")

	if err != nil {
		return *new([]ModuleAccount), err
	}

	out0 := *abi.ConvertType(out[0], new([]ModuleAccount)).(*[]ModuleAccount)

	return out0, err

}

// ModuleAccounts is a free data retrieval call binding the contract method 0x26d6d3aa.
//
// Solidity: function moduleAccounts() view returns((address,string,uint64,uint64,string,string[])[] moduleAccounts)
func (_IAuth *IAuthSession) ModuleAccounts() ([]ModuleAccount, error) {
	return _IAuth.Contract.ModuleAccounts(&_IAuth.CallOpts)
}

// ModuleAccounts is a free data retrieval call binding the contract method 0x26d6d3aa.
//
// Solidity: function moduleAccounts() view returns((address,string,uint64,uint64,string,string[])[] moduleAccounts)
func (_IAuth *IAuthCallerSession) ModuleAccounts() ([]ModuleAccount, error) {
	return _IAuth.Contract.ModuleAccounts(&_IAuth.CallOpts)
}

// Params is a free data retrieval call binding the contract method 0xcff0ab96.
//
// Solidity: function params() view returns((uint64,uint64,uint64,uint64,uint64) params)
func (_IAuth *IAuthCaller) Params(opts *bind.CallOpts) (Params, error) {
	var out []interface{}
	err := _IAuth.contract.Call(opts, &out, "params")

	if err != nil {
		return *new(Params), err
	}

	out0 := *abi.ConvertType(out[0], new(Params)).(*Params)

	return out0, err

}

// Params is a free data retrieval call binding the contract method 0xcff0ab96.
//
// Solidity: function params() view returns((uint64,uint64,uint64,uint64,uint64) params)
func (_IAuth *IAuthSession) Params() (Params, error) {
	return _IAuth.Contract.Params(&_IAuth.CallOpts)
}

// Params is a free data retrieval call binding the contract method 0xcff0ab96.
//
// Solidity: function params() view returns((uint64,uint64,uint64,uint64,uint64) params)
func (_IAuth *IAuthCallerSession) Params() (Params, error) {
	return _IAuth.Contract.Params(&_IAuth.CallOpts)
}
