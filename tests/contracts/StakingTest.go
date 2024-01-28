// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// Coin is an auto generated low-level Go binding around an user-defined struct.
type Coin struct {
	Denom  string
	Amount *big.Int
}

// StakingTestMetaData contains all meta data concerning the StakingTest contract.
var StakingTestMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegatorAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"creationHeight\",\"type\":\"uint256\"}],\"name\":\"CancelUnbondingDelegation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Delegate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegatorAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validatorSrcAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validatorDstAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"completionTime\",\"type\":\"uint256\"}],\"name\":\"Redelegate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegatorAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"completionTime\",\"type\":\"uint256\"}],\"name\":\"Undelegate\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"creationHeight\",\"type\":\"uint256\"}],\"name\":\"cancelUnbondingDelegation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"delegate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegatorAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"}],\"name\":\"delegation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structCoin\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validatorSrcAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"validatorDstAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"redelegate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"undelegate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"validatorShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610bae806100206000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c8063026e402b14610067578063046d33071461008f5780634d99dd16146100b057806350826aef146100d15780636bd8f804146100e4578063bf98d772146100f7575b600080fd5b61007a610075366004610847565b610122565b60405190151581526020015b60405180910390f35b6100a261009d3660046107da565b610137565b604051610086929190610aa1565b6100c36100be366004610847565b61016b565b604051908152602001610086565b61007a6100df366004610870565b610178565b6100c36100f236600461080c565b61018f565b6100c36101053660046108c9565b805160208183018101805160008252928201919093012091525481565b60008061012f848461019d565b949350505050565b6000610156604051806040016040528060608152602001600081525090565b6101608484610289565b915091509250929050565b60008061012f84846103a1565b60008061018685858561048f565b95945050505050565b60008061018685858561058f565b600080806110036101fa86866040516001600160a01b03831660248201526044810182905260609060640160408051601f198184030181529190526020810180516001600160e01b031663026e402b60e01b179052905092915050565b6040516102079190610a35565b6000604051808303816000865af19150503d8060008114610244576040519150601f19603f3d011682016040523d82523d6000602084013e610249565b606091505b509150915061028082826040518060400160405280600f81526020016e19195b1959d85d194819985a5b1959608a1b81525061067a565b61018681610702565b60006102a8604051806040016040528060608152602001600081525090565b60008061100361030587876040516001600160a01b0380841660248301528216604482015260609060640160408051601f198184030181529190526020810180516001600160e01b031663046d330760e01b179052905092915050565b6040516103129190610a35565b600060405180830381855afa9150503d806000811461034d576040519150601f19603f3d011682016040523d82523d6000602084013e610352565b606091505b509150915061038b82826040518060400160405280601181526020017019195b1959d85d1a5bdb8819985a5b1959607a1b81525061067a565b61039481610722565b9350935050509250929050565b600080806110036103fe86866040516001600160a01b03831660248201526044810182905260609060640160408051601f198184030181529190526020810180516001600160e01b03166326ccee8b60e11b179052905092915050565b60405161040b9190610a35565b6000604051808303816000865af19150503d8060008114610448576040519150601f19603f3d011682016040523d82523d6000602084013e61044d565b606091505b50915091506104868282604051806040016040528060118152602001701d5b99195b1959d85d194819985a5b1959607a1b81525061067a565b61018681610764565b604080516001600160a01b03851660248201526044810184905260648082018490528251808303909101815260849091019091526020810180516001600160e01b03166350826aef60e01b17905260009081908190611003906040516104f59190610a35565b6000604051808303816000865af19150503d8060008114610532576040519150601f19603f3d011682016040523d82523d6000602084013e610537565b606091505b509150915061057c82826040518060400160405280602081526020017f63616e63656c556e626f6e64696e6744656c65676174696f6e206661696c656481525061067a565b61058581610702565b9695505050505050565b604080516001600160a01b0385811660248301528416604482015260648082018490528251808303909101815260849091019091526020810180516001600160e01b0316631af63e0160e21b17905260009081908190611003906040516105f69190610a35565b6000604051808303816000865af19150503d8060008114610633576040519150601f19603f3d011682016040523d82523d6000602084013e610638565b606091505b50915091506106718282604051806040016040528060118152602001701c9959195b1959d85d194819985a5b1959607a1b81525061067a565b61058581610764565b826106fd57600082806020019051810190610695919061093d565b90506001825110156106c4578060405162461bcd60e51b81526004016106bb9190610a8e565b60405180910390fd5b81816040516020016106d7929190610a51565b60408051601f198184030181529082905262461bcd60e51b82526106bb91600401610a8e565b505050565b6000808280602001905181019061071991906108a2565b9150505b919050565b6000610741604051806040016040528060608152602001600081525090565b600080848060200190518101906107589190610988565b90945092505050915091565b600080828060200190518101906107199190610970565b80356001600160a01b038116811461071d57600080fd5b600082601f8301126107a2578081fd5b81516107b56107b082610b0a565b610ad9565b8181528460208386010111156107c9578283fd5b61012f826020830160208701610b32565b600080604083850312156107ec578182fd5b6107f58361077b565b91506108036020840161077b565b90509250929050565b600080600060608486031215610820578081fd5b6108298461077b565b92506108376020850161077b565b9150604084013590509250925092565b60008060408385031215610859578182fd5b6108628361077b565b946020939093013593505050565b600080600060608486031215610884578283fd5b61088d8461077b565b95602085013595506040909401359392505050565b6000602082840312156108b3578081fd5b815180151581146108c2578182fd5b9392505050565b6000602082840312156108da578081fd5b813567ffffffffffffffff8111156108f0578182fd5b8201601f81018413610900578182fd5b803561090e6107b082610b0a565b818152856020838501011115610922578384fd5b81602084016020830137908101602001929092525092915050565b60006020828403121561094e578081fd5b815167ffffffffffffffff811115610964578182fd5b61012f84828501610792565b600060208284031215610981578081fd5b5051919050565b6000806040838503121561099a578182fd5b82519150602083015167ffffffffffffffff808211156109b8578283fd5b90840190604082870312156109cb578283fd5b6109d56040610ad9565b8251828111156109e3578485fd5b6109ef88828601610792565b825250602083015160208201528093505050509250929050565b60008151808452610a21816020860160208601610b32565b601f01601f19169290920160200192915050565b60008251610a47818460208701610b32565b9190910192915050565b60008351610a63818460208801610b32565b6101d160f51b9083019081528351610a82816002840160208801610b32565b01600201949350505050565b6000602082526108c26020830184610a09565b6000838252604060208301528251604080840152610ac26080840182610a09565b905060208401516060840152809150509392505050565b604051601f8201601f1916810167ffffffffffffffff81118282101715610b0257610b02610b62565b604052919050565b600067ffffffffffffffff821115610b2457610b24610b62565b50601f01601f191660200190565b60005b83811015610b4d578181015183820152602001610b35565b83811115610b5c576000848401525b50505050565b634e487b7160e01b600052604160045260246000fdfea264697066735822122079465eb42f8293768ce14ee08a3d0454422e857e11d53e923110a68c4b39acec64736f6c63430008020033",
}

// StakingTestABI is the input ABI used to generate the binding from.
// Deprecated: Use StakingTestMetaData.ABI instead.
var StakingTestABI = StakingTestMetaData.ABI

// StakingTestBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StakingTestMetaData.Bin instead.
var StakingTestBin = StakingTestMetaData.Bin

// DeployStakingTest deploys a new Ethereum contract, binding an instance of StakingTest to it.
func DeployStakingTest(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *StakingTest, error) {
	parsed, err := StakingTestMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StakingTestBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StakingTest{StakingTestCaller: StakingTestCaller{contract: contract}, StakingTestTransactor: StakingTestTransactor{contract: contract}, StakingTestFilterer: StakingTestFilterer{contract: contract}}, nil
}

// StakingTest is an auto generated Go binding around an Ethereum contract.
type StakingTest struct {
	StakingTestCaller     // Read-only binding to the contract
	StakingTestTransactor // Write-only binding to the contract
	StakingTestFilterer   // Log filterer for contract events
}

// StakingTestCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakingTestCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingTestTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakingTestTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingTestFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakingTestFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingTestSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakingTestSession struct {
	Contract     *StakingTest      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakingTestCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakingTestCallerSession struct {
	Contract *StakingTestCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// StakingTestTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakingTestTransactorSession struct {
	Contract     *StakingTestTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// StakingTestRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakingTestRaw struct {
	Contract *StakingTest // Generic contract binding to access the raw methods on
}

// StakingTestCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakingTestCallerRaw struct {
	Contract *StakingTestCaller // Generic read-only contract binding to access the raw methods on
}

// StakingTestTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakingTestTransactorRaw struct {
	Contract *StakingTestTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStakingTest creates a new instance of StakingTest, bound to a specific deployed contract.
func NewStakingTest(address common.Address, backend bind.ContractBackend) (*StakingTest, error) {
	contract, err := bindStakingTest(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StakingTest{StakingTestCaller: StakingTestCaller{contract: contract}, StakingTestTransactor: StakingTestTransactor{contract: contract}, StakingTestFilterer: StakingTestFilterer{contract: contract}}, nil
}

// NewStakingTestCaller creates a new read-only instance of StakingTest, bound to a specific deployed contract.
func NewStakingTestCaller(address common.Address, caller bind.ContractCaller) (*StakingTestCaller, error) {
	contract, err := bindStakingTest(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakingTestCaller{contract: contract}, nil
}

// NewStakingTestTransactor creates a new write-only instance of StakingTest, bound to a specific deployed contract.
func NewStakingTestTransactor(address common.Address, transactor bind.ContractTransactor) (*StakingTestTransactor, error) {
	contract, err := bindStakingTest(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakingTestTransactor{contract: contract}, nil
}

// NewStakingTestFilterer creates a new log filterer instance of StakingTest, bound to a specific deployed contract.
func NewStakingTestFilterer(address common.Address, filterer bind.ContractFilterer) (*StakingTestFilterer, error) {
	contract, err := bindStakingTest(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakingTestFilterer{contract: contract}, nil
}

// bindStakingTest binds a generic wrapper to an already deployed contract.
func bindStakingTest(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StakingTestMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakingTest *StakingTestRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakingTest.Contract.StakingTestCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakingTest *StakingTestRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingTest.Contract.StakingTestTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakingTest *StakingTestRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakingTest.Contract.StakingTestTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakingTest *StakingTestCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakingTest.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakingTest *StakingTestTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingTest.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakingTest *StakingTestTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakingTest.Contract.contract.Transact(opts, method, params...)
}

// Delegation is a free data retrieval call binding the contract method 0x046d3307.
//
// Solidity: function delegation(address delegatorAddress, address validatorAddress) view returns(uint256, (string,uint256))
func (_StakingTest *StakingTestCaller) Delegation(opts *bind.CallOpts, delegatorAddress common.Address, validatorAddress common.Address) (*big.Int, Coin, error) {
	var out []interface{}
	err := _StakingTest.contract.Call(opts, &out, "delegation", delegatorAddress, validatorAddress)
	if err != nil {
		return *new(*big.Int), *new(Coin), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(Coin)).(*Coin)

	return out0, out1, err
}

// Delegation is a free data retrieval call binding the contract method 0x046d3307.
//
// Solidity: function delegation(address delegatorAddress, address validatorAddress) view returns(uint256, (string,uint256))
func (_StakingTest *StakingTestSession) Delegation(delegatorAddress common.Address, validatorAddress common.Address) (*big.Int, Coin, error) {
	return _StakingTest.Contract.Delegation(&_StakingTest.CallOpts, delegatorAddress, validatorAddress)
}

// Delegation is a free data retrieval call binding the contract method 0x046d3307.
//
// Solidity: function delegation(address delegatorAddress, address validatorAddress) view returns(uint256, (string,uint256))
func (_StakingTest *StakingTestCallerSession) Delegation(delegatorAddress common.Address, validatorAddress common.Address) (*big.Int, Coin, error) {
	return _StakingTest.Contract.Delegation(&_StakingTest.CallOpts, delegatorAddress, validatorAddress)
}

// ValidatorShares is a free data retrieval call binding the contract method 0xbf98d772.
//
// Solidity: function validatorShares(string ) view returns(uint256)
func (_StakingTest *StakingTestCaller) ValidatorShares(opts *bind.CallOpts, arg0 string) (*big.Int, error) {
	var out []interface{}
	err := _StakingTest.contract.Call(opts, &out, "validatorShares", arg0)
	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err
}

// ValidatorShares is a free data retrieval call binding the contract method 0xbf98d772.
//
// Solidity: function validatorShares(string ) view returns(uint256)
func (_StakingTest *StakingTestSession) ValidatorShares(arg0 string) (*big.Int, error) {
	return _StakingTest.Contract.ValidatorShares(&_StakingTest.CallOpts, arg0)
}

// ValidatorShares is a free data retrieval call binding the contract method 0xbf98d772.
//
// Solidity: function validatorShares(string ) view returns(uint256)
func (_StakingTest *StakingTestCallerSession) ValidatorShares(arg0 string) (*big.Int, error) {
	return _StakingTest.Contract.ValidatorShares(&_StakingTest.CallOpts, arg0)
}

// CancelUnbondingDelegation is a paid mutator transaction binding the contract method 0x50826aef.
//
// Solidity: function cancelUnbondingDelegation(address validatorAddress, uint256 amount, uint256 creationHeight) returns(bool)
func (_StakingTest *StakingTestTransactor) CancelUnbondingDelegation(opts *bind.TransactOpts, validatorAddress common.Address, amount *big.Int, creationHeight *big.Int) (*types.Transaction, error) {
	return _StakingTest.contract.Transact(opts, "cancelUnbondingDelegation", validatorAddress, amount, creationHeight)
}

// CancelUnbondingDelegation is a paid mutator transaction binding the contract method 0x50826aef.
//
// Solidity: function cancelUnbondingDelegation(address validatorAddress, uint256 amount, uint256 creationHeight) returns(bool)
func (_StakingTest *StakingTestSession) CancelUnbondingDelegation(validatorAddress common.Address, amount *big.Int, creationHeight *big.Int) (*types.Transaction, error) {
	return _StakingTest.Contract.CancelUnbondingDelegation(&_StakingTest.TransactOpts, validatorAddress, amount, creationHeight)
}

// CancelUnbondingDelegation is a paid mutator transaction binding the contract method 0x50826aef.
//
// Solidity: function cancelUnbondingDelegation(address validatorAddress, uint256 amount, uint256 creationHeight) returns(bool)
func (_StakingTest *StakingTestTransactorSession) CancelUnbondingDelegation(validatorAddress common.Address, amount *big.Int, creationHeight *big.Int) (*types.Transaction, error) {
	return _StakingTest.Contract.CancelUnbondingDelegation(&_StakingTest.TransactOpts, validatorAddress, amount, creationHeight)
}

// Delegate is a paid mutator transaction binding the contract method 0x026e402b.
//
// Solidity: function delegate(address validatorAddress, uint256 amount) returns(bool)
func (_StakingTest *StakingTestTransactor) Delegate(opts *bind.TransactOpts, validatorAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakingTest.contract.Transact(opts, "delegate", validatorAddress, amount)
}

// Delegate is a paid mutator transaction binding the contract method 0x026e402b.
//
// Solidity: function delegate(address validatorAddress, uint256 amount) returns(bool)
func (_StakingTest *StakingTestSession) Delegate(validatorAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakingTest.Contract.Delegate(&_StakingTest.TransactOpts, validatorAddress, amount)
}

// Delegate is a paid mutator transaction binding the contract method 0x026e402b.
//
// Solidity: function delegate(address validatorAddress, uint256 amount) returns(bool)
func (_StakingTest *StakingTestTransactorSession) Delegate(validatorAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakingTest.Contract.Delegate(&_StakingTest.TransactOpts, validatorAddress, amount)
}

// Redelegate is a paid mutator transaction binding the contract method 0x6bd8f804.
//
// Solidity: function redelegate(address validatorSrcAddress, address validatorDstAddress, uint256 amount) returns(uint256)
func (_StakingTest *StakingTestTransactor) Redelegate(opts *bind.TransactOpts, validatorSrcAddress common.Address, validatorDstAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakingTest.contract.Transact(opts, "redelegate", validatorSrcAddress, validatorDstAddress, amount)
}

// Redelegate is a paid mutator transaction binding the contract method 0x6bd8f804.
//
// Solidity: function redelegate(address validatorSrcAddress, address validatorDstAddress, uint256 amount) returns(uint256)
func (_StakingTest *StakingTestSession) Redelegate(validatorSrcAddress common.Address, validatorDstAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakingTest.Contract.Redelegate(&_StakingTest.TransactOpts, validatorSrcAddress, validatorDstAddress, amount)
}

// Redelegate is a paid mutator transaction binding the contract method 0x6bd8f804.
//
// Solidity: function redelegate(address validatorSrcAddress, address validatorDstAddress, uint256 amount) returns(uint256)
func (_StakingTest *StakingTestTransactorSession) Redelegate(validatorSrcAddress common.Address, validatorDstAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakingTest.Contract.Redelegate(&_StakingTest.TransactOpts, validatorSrcAddress, validatorDstAddress, amount)
}

// Undelegate is a paid mutator transaction binding the contract method 0x4d99dd16.
//
// Solidity: function undelegate(address validatorAddress, uint256 amount) returns(uint256)
func (_StakingTest *StakingTestTransactor) Undelegate(opts *bind.TransactOpts, validatorAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakingTest.contract.Transact(opts, "undelegate", validatorAddress, amount)
}

// Undelegate is a paid mutator transaction binding the contract method 0x4d99dd16.
//
// Solidity: function undelegate(address validatorAddress, uint256 amount) returns(uint256)
func (_StakingTest *StakingTestSession) Undelegate(validatorAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakingTest.Contract.Undelegate(&_StakingTest.TransactOpts, validatorAddress, amount)
}

// Undelegate is a paid mutator transaction binding the contract method 0x4d99dd16.
//
// Solidity: function undelegate(address validatorAddress, uint256 amount) returns(uint256)
func (_StakingTest *StakingTestTransactorSession) Undelegate(validatorAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakingTest.Contract.Undelegate(&_StakingTest.TransactOpts, validatorAddress, amount)
}

// StakingTestCancelUnbondingDelegationIterator is returned from FilterCancelUnbondingDelegation and is used to iterate over the raw logs and unpacked data for CancelUnbondingDelegation events raised by the StakingTest contract.
type StakingTestCancelUnbondingDelegationIterator struct {
	Event *StakingTestCancelUnbondingDelegation // Event containing the contract specifics and raw log

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
func (it *StakingTestCancelUnbondingDelegationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingTestCancelUnbondingDelegation)
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
		it.Event = new(StakingTestCancelUnbondingDelegation)
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
func (it *StakingTestCancelUnbondingDelegationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingTestCancelUnbondingDelegationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingTestCancelUnbondingDelegation represents a CancelUnbondingDelegation event raised by the StakingTest contract.
type StakingTestCancelUnbondingDelegation struct {
	DelegatorAddress common.Address
	ValidatorAddress common.Address
	Amount           *big.Int
	CreationHeight   *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterCancelUnbondingDelegation is a free log retrieval operation binding the contract event 0x6dbe2fb6b2613bdd8e3d284a6111592e06c3ab0af846ff89b6688d48f408dbb5.
//
// Solidity: event CancelUnbondingDelegation(address indexed delegatorAddress, address indexed validatorAddress, uint256 amount, uint256 creationHeight)
func (_StakingTest *StakingTestFilterer) FilterCancelUnbondingDelegation(opts *bind.FilterOpts, delegatorAddress []common.Address, validatorAddress []common.Address) (*StakingTestCancelUnbondingDelegationIterator, error) {
	var delegatorAddressRule []interface{}
	for _, delegatorAddressItem := range delegatorAddress {
		delegatorAddressRule = append(delegatorAddressRule, delegatorAddressItem)
	}
	var validatorAddressRule []interface{}
	for _, validatorAddressItem := range validatorAddress {
		validatorAddressRule = append(validatorAddressRule, validatorAddressItem)
	}

	logs, sub, err := _StakingTest.contract.FilterLogs(opts, "CancelUnbondingDelegation", delegatorAddressRule, validatorAddressRule)
	if err != nil {
		return nil, err
	}
	return &StakingTestCancelUnbondingDelegationIterator{contract: _StakingTest.contract, event: "CancelUnbondingDelegation", logs: logs, sub: sub}, nil
}

// WatchCancelUnbondingDelegation is a free log subscription operation binding the contract event 0x6dbe2fb6b2613bdd8e3d284a6111592e06c3ab0af846ff89b6688d48f408dbb5.
//
// Solidity: event CancelUnbondingDelegation(address indexed delegatorAddress, address indexed validatorAddress, uint256 amount, uint256 creationHeight)
func (_StakingTest *StakingTestFilterer) WatchCancelUnbondingDelegation(opts *bind.WatchOpts, sink chan<- *StakingTestCancelUnbondingDelegation, delegatorAddress []common.Address, validatorAddress []common.Address) (event.Subscription, error) {
	var delegatorAddressRule []interface{}
	for _, delegatorAddressItem := range delegatorAddress {
		delegatorAddressRule = append(delegatorAddressRule, delegatorAddressItem)
	}
	var validatorAddressRule []interface{}
	for _, validatorAddressItem := range validatorAddress {
		validatorAddressRule = append(validatorAddressRule, validatorAddressItem)
	}

	logs, sub, err := _StakingTest.contract.WatchLogs(opts, "CancelUnbondingDelegation", delegatorAddressRule, validatorAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingTestCancelUnbondingDelegation)
				if err := _StakingTest.contract.UnpackLog(event, "CancelUnbondingDelegation", log); err != nil {
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

// ParseCancelUnbondingDelegation is a log parse operation binding the contract event 0x6dbe2fb6b2613bdd8e3d284a6111592e06c3ab0af846ff89b6688d48f408dbb5.
//
// Solidity: event CancelUnbondingDelegation(address indexed delegatorAddress, address indexed validatorAddress, uint256 amount, uint256 creationHeight)
func (_StakingTest *StakingTestFilterer) ParseCancelUnbondingDelegation(log types.Log) (*StakingTestCancelUnbondingDelegation, error) {
	event := new(StakingTestCancelUnbondingDelegation)
	if err := _StakingTest.contract.UnpackLog(event, "CancelUnbondingDelegation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingTestDelegateIterator is returned from FilterDelegate and is used to iterate over the raw logs and unpacked data for Delegate events raised by the StakingTest contract.
type StakingTestDelegateIterator struct {
	Event *StakingTestDelegate // Event containing the contract specifics and raw log

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
func (it *StakingTestDelegateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingTestDelegate)
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
		it.Event = new(StakingTestDelegate)
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
func (it *StakingTestDelegateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingTestDelegateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingTestDelegate represents a Delegate event raised by the StakingTest contract.
type StakingTestDelegate struct {
	Delegator common.Address
	Validator common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDelegate is a free log retrieval operation binding the contract event 0x510b11bb3f3c799b11307c01ab7db0d335683ef5b2da98f7697de744f465eacc.
//
// Solidity: event Delegate(address indexed delegator, address indexed validator, uint256 amount)
func (_StakingTest *StakingTestFilterer) FilterDelegate(opts *bind.FilterOpts, delegator []common.Address, validator []common.Address) (*StakingTestDelegateIterator, error) {
	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _StakingTest.contract.FilterLogs(opts, "Delegate", delegatorRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return &StakingTestDelegateIterator{contract: _StakingTest.contract, event: "Delegate", logs: logs, sub: sub}, nil
}

// WatchDelegate is a free log subscription operation binding the contract event 0x510b11bb3f3c799b11307c01ab7db0d335683ef5b2da98f7697de744f465eacc.
//
// Solidity: event Delegate(address indexed delegator, address indexed validator, uint256 amount)
func (_StakingTest *StakingTestFilterer) WatchDelegate(opts *bind.WatchOpts, sink chan<- *StakingTestDelegate, delegator []common.Address, validator []common.Address) (event.Subscription, error) {
	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _StakingTest.contract.WatchLogs(opts, "Delegate", delegatorRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingTestDelegate)
				if err := _StakingTest.contract.UnpackLog(event, "Delegate", log); err != nil {
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

// ParseDelegate is a log parse operation binding the contract event 0x510b11bb3f3c799b11307c01ab7db0d335683ef5b2da98f7697de744f465eacc.
//
// Solidity: event Delegate(address indexed delegator, address indexed validator, uint256 amount)
func (_StakingTest *StakingTestFilterer) ParseDelegate(log types.Log) (*StakingTestDelegate, error) {
	event := new(StakingTestDelegate)
	if err := _StakingTest.contract.UnpackLog(event, "Delegate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingTestRedelegateIterator is returned from FilterRedelegate and is used to iterate over the raw logs and unpacked data for Redelegate events raised by the StakingTest contract.
type StakingTestRedelegateIterator struct {
	Event *StakingTestRedelegate // Event containing the contract specifics and raw log

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
func (it *StakingTestRedelegateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingTestRedelegate)
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
		it.Event = new(StakingTestRedelegate)
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
func (it *StakingTestRedelegateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingTestRedelegateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingTestRedelegate represents a Redelegate event raised by the StakingTest contract.
type StakingTestRedelegate struct {
	DelegatorAddress    common.Address
	ValidatorSrcAddress common.Address
	ValidatorDstAddress common.Address
	Amount              *big.Int
	CompletionTime      *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterRedelegate is a free log retrieval operation binding the contract event 0x82b07f2421474f1e3f1e0b34738cb5ffb925273f408e7591d9c803dcae8da657.
//
// Solidity: event Redelegate(address indexed delegatorAddress, address indexed validatorSrcAddress, address indexed validatorDstAddress, uint256 amount, uint256 completionTime)
func (_StakingTest *StakingTestFilterer) FilterRedelegate(opts *bind.FilterOpts, delegatorAddress []common.Address, validatorSrcAddress []common.Address, validatorDstAddress []common.Address) (*StakingTestRedelegateIterator, error) {
	var delegatorAddressRule []interface{}
	for _, delegatorAddressItem := range delegatorAddress {
		delegatorAddressRule = append(delegatorAddressRule, delegatorAddressItem)
	}
	var validatorSrcAddressRule []interface{}
	for _, validatorSrcAddressItem := range validatorSrcAddress {
		validatorSrcAddressRule = append(validatorSrcAddressRule, validatorSrcAddressItem)
	}
	var validatorDstAddressRule []interface{}
	for _, validatorDstAddressItem := range validatorDstAddress {
		validatorDstAddressRule = append(validatorDstAddressRule, validatorDstAddressItem)
	}

	logs, sub, err := _StakingTest.contract.FilterLogs(opts, "Redelegate", delegatorAddressRule, validatorSrcAddressRule, validatorDstAddressRule)
	if err != nil {
		return nil, err
	}
	return &StakingTestRedelegateIterator{contract: _StakingTest.contract, event: "Redelegate", logs: logs, sub: sub}, nil
}

// WatchRedelegate is a free log subscription operation binding the contract event 0x82b07f2421474f1e3f1e0b34738cb5ffb925273f408e7591d9c803dcae8da657.
//
// Solidity: event Redelegate(address indexed delegatorAddress, address indexed validatorSrcAddress, address indexed validatorDstAddress, uint256 amount, uint256 completionTime)
func (_StakingTest *StakingTestFilterer) WatchRedelegate(opts *bind.WatchOpts, sink chan<- *StakingTestRedelegate, delegatorAddress []common.Address, validatorSrcAddress []common.Address, validatorDstAddress []common.Address) (event.Subscription, error) {
	var delegatorAddressRule []interface{}
	for _, delegatorAddressItem := range delegatorAddress {
		delegatorAddressRule = append(delegatorAddressRule, delegatorAddressItem)
	}
	var validatorSrcAddressRule []interface{}
	for _, validatorSrcAddressItem := range validatorSrcAddress {
		validatorSrcAddressRule = append(validatorSrcAddressRule, validatorSrcAddressItem)
	}
	var validatorDstAddressRule []interface{}
	for _, validatorDstAddressItem := range validatorDstAddress {
		validatorDstAddressRule = append(validatorDstAddressRule, validatorDstAddressItem)
	}

	logs, sub, err := _StakingTest.contract.WatchLogs(opts, "Redelegate", delegatorAddressRule, validatorSrcAddressRule, validatorDstAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingTestRedelegate)
				if err := _StakingTest.contract.UnpackLog(event, "Redelegate", log); err != nil {
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

// ParseRedelegate is a log parse operation binding the contract event 0x82b07f2421474f1e3f1e0b34738cb5ffb925273f408e7591d9c803dcae8da657.
//
// Solidity: event Redelegate(address indexed delegatorAddress, address indexed validatorSrcAddress, address indexed validatorDstAddress, uint256 amount, uint256 completionTime)
func (_StakingTest *StakingTestFilterer) ParseRedelegate(log types.Log) (*StakingTestRedelegate, error) {
	event := new(StakingTestRedelegate)
	if err := _StakingTest.contract.UnpackLog(event, "Redelegate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingTestUndelegateIterator is returned from FilterUndelegate and is used to iterate over the raw logs and unpacked data for Undelegate events raised by the StakingTest contract.
type StakingTestUndelegateIterator struct {
	Event *StakingTestUndelegate // Event containing the contract specifics and raw log

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
func (it *StakingTestUndelegateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingTestUndelegate)
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
		it.Event = new(StakingTestUndelegate)
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
func (it *StakingTestUndelegateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingTestUndelegateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingTestUndelegate represents a Undelegate event raised by the StakingTest contract.
type StakingTestUndelegate struct {
	DelegatorAddress common.Address
	ValidatorAddress common.Address
	Amount           *big.Int
	CompletionTime   *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterUndelegate is a free log retrieval operation binding the contract event 0x54feacf7d18f42300ff65ba8829f17724df4915a8b1984c5fee2cc42d4027f07.
//
// Solidity: event Undelegate(address indexed delegatorAddress, address indexed validatorAddress, uint256 amount, uint256 completionTime)
func (_StakingTest *StakingTestFilterer) FilterUndelegate(opts *bind.FilterOpts, delegatorAddress []common.Address, validatorAddress []common.Address) (*StakingTestUndelegateIterator, error) {
	var delegatorAddressRule []interface{}
	for _, delegatorAddressItem := range delegatorAddress {
		delegatorAddressRule = append(delegatorAddressRule, delegatorAddressItem)
	}
	var validatorAddressRule []interface{}
	for _, validatorAddressItem := range validatorAddress {
		validatorAddressRule = append(validatorAddressRule, validatorAddressItem)
	}

	logs, sub, err := _StakingTest.contract.FilterLogs(opts, "Undelegate", delegatorAddressRule, validatorAddressRule)
	if err != nil {
		return nil, err
	}
	return &StakingTestUndelegateIterator{contract: _StakingTest.contract, event: "Undelegate", logs: logs, sub: sub}, nil
}

// WatchUndelegate is a free log subscription operation binding the contract event 0x54feacf7d18f42300ff65ba8829f17724df4915a8b1984c5fee2cc42d4027f07.
//
// Solidity: event Undelegate(address indexed delegatorAddress, address indexed validatorAddress, uint256 amount, uint256 completionTime)
func (_StakingTest *StakingTestFilterer) WatchUndelegate(opts *bind.WatchOpts, sink chan<- *StakingTestUndelegate, delegatorAddress []common.Address, validatorAddress []common.Address) (event.Subscription, error) {
	var delegatorAddressRule []interface{}
	for _, delegatorAddressItem := range delegatorAddress {
		delegatorAddressRule = append(delegatorAddressRule, delegatorAddressItem)
	}
	var validatorAddressRule []interface{}
	for _, validatorAddressItem := range validatorAddress {
		validatorAddressRule = append(validatorAddressRule, validatorAddressItem)
	}

	logs, sub, err := _StakingTest.contract.WatchLogs(opts, "Undelegate", delegatorAddressRule, validatorAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingTestUndelegate)
				if err := _StakingTest.contract.UnpackLog(event, "Undelegate", log); err != nil {
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

// ParseUndelegate is a log parse operation binding the contract event 0x54feacf7d18f42300ff65ba8829f17724df4915a8b1984c5fee2cc42d4027f07.
//
// Solidity: event Undelegate(address indexed delegatorAddress, address indexed validatorAddress, uint256 amount, uint256 completionTime)
func (_StakingTest *StakingTestFilterer) ParseUndelegate(log types.Log) (*StakingTestUndelegate, error) {
	event := new(StakingTestUndelegate)
	if err := _StakingTest.contract.UnpackLog(event, "Undelegate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
