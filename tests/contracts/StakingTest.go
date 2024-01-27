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
)

// StakingTestMetaData contains all meta data concerning the StakingTest contract.
var StakingTestMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"validator\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"Delegate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"valSrc\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"valDst\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"completionTime\",\"type\":\"uint256\"}],\"name\":\"Redelegate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"validator\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"completionTime\",\"type\":\"uint256\"}],\"name\":\"Undelegate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"validator\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_val\",\"type\":\"string\"}],\"name\":\"delegate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_val\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_del\",\"type\":\"address\"}],\"name\":\"delegation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_val\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_del\",\"type\":\"address\"}],\"name\":\"delegationRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_valSrc\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_valDst\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_shares\",\"type\":\"uint256\"}],\"name\":\"redelegate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_val\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_shares\",\"type\":\"uint256\"}],\"name\":\"undelegate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"validatorShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_val\",\"type\":\"string\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610e4a806100206000396000f3fe6080604052600436106100705760003560e01c80638dfc88971161004e5780638dfc8897146101035780639ddb511a14610123578063bf98d7721461014b578063d5c498eb1461018357610070565b806331fb67c21461007557806351af513a146100a85780637dd0209d146100c8575b600080fd5b34801561008157600080fd5b50610095610090366004610a09565b6101a3565b6040519081526020015b60405180910390f35b3480156100b457600080fd5b506100956100c3366004610aa6565b6101b6565b3480156100d457600080fd5b506100e86100e3366004610b01565b6101c2565b6040805193845260208401929092529082015260600161009f565b34801561010f57600080fd5b506100e861011e366004610b6b565b610252565b610136610131366004610a09565b6102ab565b6040805192835260208301919091520161009f565b34801561015757600080fd5b50610095610166366004610a09565b805160208183018101805160008252928201919093012091525481565b34801561018f57600080fd5b5061013661019e366004610aa6565b6102fc565b6000806101af83610314565b9392505050565b60006101af83836103bb565b6000806000806000806101d6898989610470565b9250925092508660008a6040516101ed9190610c42565b9081526020016040518091039020600082825461020a9190610da1565b92505081905550866000896040516102229190610c42565b9081526020016040518091039020600082825461023f9190610d89565b9091555092999198509650945050505050565b6000806000806000806102658888610526565b9250925092508660008960405161027c9190610c42565b908152602001604051809103902060008282546102999190610da1565b90915550929891975095509350505050565b6000806000806102bb85346105d9565b91509150816000866040516102d09190610c42565b908152602001604051809103902060008282546102ed9190610d89565b90915550919350915050915091565b6000806103098484610687565b915091509250929050565b600080806110036103248561071f565b6040516103319190610c42565b6000604051808303816000865af19150503d806000811461036e576040519150601f19603f3d011682016040523d82523d6000602084013e610373565b606091505b50915091506103aa82826040518060400160405280600f81526020016e1dda5d1a191c985dc819985a5b1959608a1b815250610763565b6103b3816107eb565b949350505050565b600080806110036103cc8686610802565b6040516103d99190610c42565b600060405180830381855afa9150503d8060008114610414576040519150601f19603f3d011682016040523d82523d6000602084013e610419565b606091505b509150915061045e82826040518060400160405280601881526020017f64656c65676174696f6e52657761726473206661696c65640000000000000000815250610763565b610467816107eb565b95945050505050565b600080808080611003610484898989610849565b6040516104919190610c42565b6000604051808303816000865af19150503d80600081146104ce576040519150601f19603f3d011682016040523d82523d6000602084013e6104d3565b606091505b509150915061050c8282604051806040016040528060118152602001701c9959195b1959d85d194819985a5b1959607a1b815250610763565b61051581610893565b945094509450505093509350939050565b60008080808061100361053988886108c0565b6040516105469190610c42565b6000604051808303816000865af19150503d8060008114610583576040519150601f19603f3d011682016040523d82523d6000602084013e610588565b606091505b50915091506105c18282604051806040016040528060118152602001701d5b99195b1959d85d194819985a5b1959607a1b815250610763565b6105ca81610893565b94509450945050509250925092565b6000808080611003856105eb88610907565b6040516105f89190610c42565b60006040518083038185875af1925050503d8060008114610635576040519150601f19603f3d011682016040523d82523d6000602084013e61063a565b606091505b509150915061067182826040518060400160405280600f81526020016e19195b1959d85d194819985a5b1959608a1b815250610763565b61067a8161094b565b9350935050509250929050565b60008080806110036106998787610971565b6040516106a69190610c42565b600060405180830381855afa9150503d80600081146106e1576040519150601f19603f3d011682016040523d82523d6000602084013e6106e6565b606091505b509150915061067182826040518060400160405280601181526020017019195b1959d85d1a5bdb8819985a5b1959607a1b815250610763565b6060816040516024016107329190610c9b565b60408051601f198184030181529190526020810180516001600160e01b03166318fdb3e160e11b1790529050919050565b826107e65760008280602001905181019061077e9190610a3c565b90506001825110156107ad578060405162461bcd60e51b81526004016107a49190610c9b565b60405180910390fd5b81816040516020016107c0929190610c5e565b60408051601f198184030181529082905262461bcd60e51b82526107a491600401610c9b565b505050565b600080828060200190518101906101af9190610bae565b60608282604051602401610817929190610cae565b60408051601f198184030181529190526020810180516001600160e01b03166328d7a89d60e11b179052905092915050565b606083838360405160240161086093929190610cd8565b60408051601f198184030181529190526020810180516001600160e01b0316637dd0209d60e01b17905290509392505050565b600080600080600080868060200190518101906108b09190610be9565b9199909850909650945050505050565b606082826040516024016108d5929190610d0e565b60408051601f198184030181529190526020810180516001600160e01b0316638dfc889760e01b179052905092915050565b60608160405160240161091a9190610c9b565b60408051601f198184030181529190526020810180516001600160e01b0316634eeda88d60e11b1790529050919050565b600080600080848060200190518101906109659190610bc6565b90945092505050915091565b60608282604051602401610986929190610cae565b60408051601f198184030181529190526020810180516001600160e01b031663d5c498eb60e01b179052905092915050565b600082601f8301126109c8578081fd5b81356109db6109d682610d61565b610d30565b8181528460208386010111156109ef578283fd5b816020850160208301379081016020019190915292915050565b600060208284031215610a1a578081fd5b813567ffffffffffffffff811115610a30578182fd5b6103b3848285016109b8565b600060208284031215610a4d578081fd5b815167ffffffffffffffff811115610a63578182fd5b8201601f81018413610a73578182fd5b8051610a816109d682610d61565b818152856020838501011115610a95578384fd5b610467826020830160208601610db8565b60008060408385031215610ab8578081fd5b823567ffffffffffffffff811115610ace578182fd5b610ada858286016109b8565b92505060208301356001600160a01b0381168114610af6578182fd5b809150509250929050565b600080600060608486031215610b15578081fd5b833567ffffffffffffffff80821115610b2c578283fd5b610b38878388016109b8565b94506020860135915080821115610b4d578283fd5b50610b5a868287016109b8565b925050604084013590509250925092565b60008060408385031215610b7d578182fd5b823567ffffffffffffffff811115610b93578283fd5b610b9f858286016109b8565b95602094909401359450505050565b600060208284031215610bbf578081fd5b5051919050565b60008060408385031215610bd8578182fd5b505080516020909101519092909150565b600080600060608486031215610bfd578283fd5b8351925060208401519150604084015190509250925092565b60008151808452610c2e816020860160208601610db8565b601f01601f19169290920160200192915050565b60008251610c54818460208701610db8565b9190910192915050565b60008351610c70818460208801610db8565b6101d160f51b9083019081528351610c8f816002840160208801610db8565b01600201949350505050565b6000602082526101af6020830184610c16565b600060408252610cc16040830185610c16565b905060018060a01b03831660208301529392505050565b600060608252610ceb6060830186610c16565b8281036020840152610cfd8186610c16565b915050826040830152949350505050565b600060408252610d216040830185610c16565b90508260208301529392505050565b604051601f8201601f1916810167ffffffffffffffff81118282101715610d5957610d59610dfe565b604052919050565b600067ffffffffffffffff821115610d7b57610d7b610dfe565b50601f01601f191660200190565b60008219821115610d9c57610d9c610de8565b500190565b600082821015610db357610db3610de8565b500390565b60005b83811015610dd3578181015183820152602001610dbb565b83811115610de2576000848401525b50505050565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052604160045260246000fdfea264697066735822122005ed147156607a39b11a9f1ef93eff5ebc1da30180dd2b5c0b754d85b405635c64736f6c63430008020033",
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
	parsed, err := abi.JSON(strings.NewReader(StakingTestABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
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

// Delegation is a free data retrieval call binding the contract method 0xd5c498eb.
//
// Solidity: function delegation(string _val, address _del) view returns(uint256, uint256)
func (_StakingTest *StakingTestCaller) Delegation(opts *bind.CallOpts, _val string, _del common.Address) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _StakingTest.contract.Call(opts, &out, "delegation", _val, _del)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// Delegation is a free data retrieval call binding the contract method 0xd5c498eb.
//
// Solidity: function delegation(string _val, address _del) view returns(uint256, uint256)
func (_StakingTest *StakingTestSession) Delegation(_val string, _del common.Address) (*big.Int, *big.Int, error) {
	return _StakingTest.Contract.Delegation(&_StakingTest.CallOpts, _val, _del)
}

// Delegation is a free data retrieval call binding the contract method 0xd5c498eb.
//
// Solidity: function delegation(string _val, address _del) view returns(uint256, uint256)
func (_StakingTest *StakingTestCallerSession) Delegation(_val string, _del common.Address) (*big.Int, *big.Int, error) {
	return _StakingTest.Contract.Delegation(&_StakingTest.CallOpts, _val, _del)
}

// DelegationRewards is a free data retrieval call binding the contract method 0x51af513a.
//
// Solidity: function delegationRewards(string _val, address _del) view returns(uint256)
func (_StakingTest *StakingTestCaller) DelegationRewards(opts *bind.CallOpts, _val string, _del common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StakingTest.contract.Call(opts, &out, "delegationRewards", _val, _del)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DelegationRewards is a free data retrieval call binding the contract method 0x51af513a.
//
// Solidity: function delegationRewards(string _val, address _del) view returns(uint256)
func (_StakingTest *StakingTestSession) DelegationRewards(_val string, _del common.Address) (*big.Int, error) {
	return _StakingTest.Contract.DelegationRewards(&_StakingTest.CallOpts, _val, _del)
}

// DelegationRewards is a free data retrieval call binding the contract method 0x51af513a.
//
// Solidity: function delegationRewards(string _val, address _del) view returns(uint256)
func (_StakingTest *StakingTestCallerSession) DelegationRewards(_val string, _del common.Address) (*big.Int, error) {
	return _StakingTest.Contract.DelegationRewards(&_StakingTest.CallOpts, _val, _del)
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

// Delegate is a paid mutator transaction binding the contract method 0x9ddb511a.
//
// Solidity: function delegate(string _val) payable returns(uint256, uint256)
func (_StakingTest *StakingTestTransactor) Delegate(opts *bind.TransactOpts, _val string) (*types.Transaction, error) {
	return _StakingTest.contract.Transact(opts, "delegate", _val)
}

// Delegate is a paid mutator transaction binding the contract method 0x9ddb511a.
//
// Solidity: function delegate(string _val) payable returns(uint256, uint256)
func (_StakingTest *StakingTestSession) Delegate(_val string) (*types.Transaction, error) {
	return _StakingTest.Contract.Delegate(&_StakingTest.TransactOpts, _val)
}

// Delegate is a paid mutator transaction binding the contract method 0x9ddb511a.
//
// Solidity: function delegate(string _val) payable returns(uint256, uint256)
func (_StakingTest *StakingTestTransactorSession) Delegate(_val string) (*types.Transaction, error) {
	return _StakingTest.Contract.Delegate(&_StakingTest.TransactOpts, _val)
}

// Redelegate is a paid mutator transaction binding the contract method 0x7dd0209d.
//
// Solidity: function redelegate(string _valSrc, string _valDst, uint256 _shares) returns(uint256, uint256, uint256)
func (_StakingTest *StakingTestTransactor) Redelegate(opts *bind.TransactOpts, _valSrc string, _valDst string, _shares *big.Int) (*types.Transaction, error) {
	return _StakingTest.contract.Transact(opts, "redelegate", _valSrc, _valDst, _shares)
}

// Redelegate is a paid mutator transaction binding the contract method 0x7dd0209d.
//
// Solidity: function redelegate(string _valSrc, string _valDst, uint256 _shares) returns(uint256, uint256, uint256)
func (_StakingTest *StakingTestSession) Redelegate(_valSrc string, _valDst string, _shares *big.Int) (*types.Transaction, error) {
	return _StakingTest.Contract.Redelegate(&_StakingTest.TransactOpts, _valSrc, _valDst, _shares)
}

// Redelegate is a paid mutator transaction binding the contract method 0x7dd0209d.
//
// Solidity: function redelegate(string _valSrc, string _valDst, uint256 _shares) returns(uint256, uint256, uint256)
func (_StakingTest *StakingTestTransactorSession) Redelegate(_valSrc string, _valDst string, _shares *big.Int) (*types.Transaction, error) {
	return _StakingTest.Contract.Redelegate(&_StakingTest.TransactOpts, _valSrc, _valDst, _shares)
}

// Undelegate is a paid mutator transaction binding the contract method 0x8dfc8897.
//
// Solidity: function undelegate(string _val, uint256 _shares) returns(uint256, uint256, uint256)
func (_StakingTest *StakingTestTransactor) Undelegate(opts *bind.TransactOpts, _val string, _shares *big.Int) (*types.Transaction, error) {
	return _StakingTest.contract.Transact(opts, "undelegate", _val, _shares)
}

// Undelegate is a paid mutator transaction binding the contract method 0x8dfc8897.
//
// Solidity: function undelegate(string _val, uint256 _shares) returns(uint256, uint256, uint256)
func (_StakingTest *StakingTestSession) Undelegate(_val string, _shares *big.Int) (*types.Transaction, error) {
	return _StakingTest.Contract.Undelegate(&_StakingTest.TransactOpts, _val, _shares)
}

// Undelegate is a paid mutator transaction binding the contract method 0x8dfc8897.
//
// Solidity: function undelegate(string _val, uint256 _shares) returns(uint256, uint256, uint256)
func (_StakingTest *StakingTestTransactorSession) Undelegate(_val string, _shares *big.Int) (*types.Transaction, error) {
	return _StakingTest.Contract.Undelegate(&_StakingTest.TransactOpts, _val, _shares)
}

// Withdraw is a paid mutator transaction binding the contract method 0x31fb67c2.
//
// Solidity: function withdraw(string _val) returns(uint256)
func (_StakingTest *StakingTestTransactor) Withdraw(opts *bind.TransactOpts, _val string) (*types.Transaction, error) {
	return _StakingTest.contract.Transact(opts, "withdraw", _val)
}

// Withdraw is a paid mutator transaction binding the contract method 0x31fb67c2.
//
// Solidity: function withdraw(string _val) returns(uint256)
func (_StakingTest *StakingTestSession) Withdraw(_val string) (*types.Transaction, error) {
	return _StakingTest.Contract.Withdraw(&_StakingTest.TransactOpts, _val)
}

// Withdraw is a paid mutator transaction binding the contract method 0x31fb67c2.
//
// Solidity: function withdraw(string _val) returns(uint256)
func (_StakingTest *StakingTestTransactorSession) Withdraw(_val string) (*types.Transaction, error) {
	return _StakingTest.Contract.Withdraw(&_StakingTest.TransactOpts, _val)
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
	Validator string
	Amount    *big.Int
	Shares    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDelegate is a free log retrieval operation binding the contract event 0x5a5adf903ba232ef17ed8be4ef872e1f60d17c5ba26a1ecbf44e388a672b118a.
//
// Solidity: event Delegate(address indexed delegator, string validator, uint256 amount, uint256 shares)
func (_StakingTest *StakingTestFilterer) FilterDelegate(opts *bind.FilterOpts, delegator []common.Address) (*StakingTestDelegateIterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}

	logs, sub, err := _StakingTest.contract.FilterLogs(opts, "Delegate", delegatorRule)
	if err != nil {
		return nil, err
	}
	return &StakingTestDelegateIterator{contract: _StakingTest.contract, event: "Delegate", logs: logs, sub: sub}, nil
}

// WatchDelegate is a free log subscription operation binding the contract event 0x5a5adf903ba232ef17ed8be4ef872e1f60d17c5ba26a1ecbf44e388a672b118a.
//
// Solidity: event Delegate(address indexed delegator, string validator, uint256 amount, uint256 shares)
func (_StakingTest *StakingTestFilterer) WatchDelegate(opts *bind.WatchOpts, sink chan<- *StakingTestDelegate, delegator []common.Address) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}

	logs, sub, err := _StakingTest.contract.WatchLogs(opts, "Delegate", delegatorRule)
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

// ParseDelegate is a log parse operation binding the contract event 0x5a5adf903ba232ef17ed8be4ef872e1f60d17c5ba26a1ecbf44e388a672b118a.
//
// Solidity: event Delegate(address indexed delegator, string validator, uint256 amount, uint256 shares)
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
	Sender         common.Address
	ValSrc         string
	ValDst         string
	Shares         *big.Int
	Amount         *big.Int
	CompletionTime *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterRedelegate is a free log retrieval operation binding the contract event 0x14e0e9558f524ca41364e4e284ebe7aabee65559c8ea32a6fca4d812e0a1d9e6.
//
// Solidity: event Redelegate(address indexed sender, string valSrc, string valDst, uint256 shares, uint256 amount, uint256 completionTime)
func (_StakingTest *StakingTestFilterer) FilterRedelegate(opts *bind.FilterOpts, sender []common.Address) (*StakingTestRedelegateIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _StakingTest.contract.FilterLogs(opts, "Redelegate", senderRule)
	if err != nil {
		return nil, err
	}
	return &StakingTestRedelegateIterator{contract: _StakingTest.contract, event: "Redelegate", logs: logs, sub: sub}, nil
}

// WatchRedelegate is a free log subscription operation binding the contract event 0x14e0e9558f524ca41364e4e284ebe7aabee65559c8ea32a6fca4d812e0a1d9e6.
//
// Solidity: event Redelegate(address indexed sender, string valSrc, string valDst, uint256 shares, uint256 amount, uint256 completionTime)
func (_StakingTest *StakingTestFilterer) WatchRedelegate(opts *bind.WatchOpts, sink chan<- *StakingTestRedelegate, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _StakingTest.contract.WatchLogs(opts, "Redelegate", senderRule)
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

// ParseRedelegate is a log parse operation binding the contract event 0x14e0e9558f524ca41364e4e284ebe7aabee65559c8ea32a6fca4d812e0a1d9e6.
//
// Solidity: event Redelegate(address indexed sender, string valSrc, string valDst, uint256 shares, uint256 amount, uint256 completionTime)
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
	Sender         common.Address
	Validator      string
	Shares         *big.Int
	Amount         *big.Int
	CompletionTime *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUndelegate is a free log retrieval operation binding the contract event 0xadff14cd34035a6bbb90fbe80979f36398f244f1885f7612e6e33a05a0b90d0f.
//
// Solidity: event Undelegate(address indexed sender, string validator, uint256 shares, uint256 amount, uint256 completionTime)
func (_StakingTest *StakingTestFilterer) FilterUndelegate(opts *bind.FilterOpts, sender []common.Address) (*StakingTestUndelegateIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _StakingTest.contract.FilterLogs(opts, "Undelegate", senderRule)
	if err != nil {
		return nil, err
	}
	return &StakingTestUndelegateIterator{contract: _StakingTest.contract, event: "Undelegate", logs: logs, sub: sub}, nil
}

// WatchUndelegate is a free log subscription operation binding the contract event 0xadff14cd34035a6bbb90fbe80979f36398f244f1885f7612e6e33a05a0b90d0f.
//
// Solidity: event Undelegate(address indexed sender, string validator, uint256 shares, uint256 amount, uint256 completionTime)
func (_StakingTest *StakingTestFilterer) WatchUndelegate(opts *bind.WatchOpts, sink chan<- *StakingTestUndelegate, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _StakingTest.contract.WatchLogs(opts, "Undelegate", senderRule)
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

// ParseUndelegate is a log parse operation binding the contract event 0xadff14cd34035a6bbb90fbe80979f36398f244f1885f7612e6e33a05a0b90d0f.
//
// Solidity: event Undelegate(address indexed sender, string validator, uint256 shares, uint256 amount, uint256 completionTime)
func (_StakingTest *StakingTestFilterer) ParseUndelegate(log types.Log) (*StakingTestUndelegate, error) {
	event := new(StakingTestUndelegate)
	if err := _StakingTest.contract.UnpackLog(event, "Undelegate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingTestWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the StakingTest contract.
type StakingTestWithdrawIterator struct {
	Event *StakingTestWithdraw // Event containing the contract specifics and raw log

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
func (it *StakingTestWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingTestWithdraw)
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
		it.Event = new(StakingTestWithdraw)
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
func (it *StakingTestWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingTestWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingTestWithdraw represents a Withdraw event raised by the StakingTest contract.
type StakingTestWithdraw struct {
	Sender    common.Address
	Validator string
	Reward    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x901c03da5d88eb3d62ab4617e7b7d17d86db16356823a7971127d5181a842fef.
//
// Solidity: event Withdraw(address indexed sender, string validator, uint256 reward)
func (_StakingTest *StakingTestFilterer) FilterWithdraw(opts *bind.FilterOpts, sender []common.Address) (*StakingTestWithdrawIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _StakingTest.contract.FilterLogs(opts, "Withdraw", senderRule)
	if err != nil {
		return nil, err
	}
	return &StakingTestWithdrawIterator{contract: _StakingTest.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x901c03da5d88eb3d62ab4617e7b7d17d86db16356823a7971127d5181a842fef.
//
// Solidity: event Withdraw(address indexed sender, string validator, uint256 reward)
func (_StakingTest *StakingTestFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *StakingTestWithdraw, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _StakingTest.contract.WatchLogs(opts, "Withdraw", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingTestWithdraw)
				if err := _StakingTest.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x901c03da5d88eb3d62ab4617e7b7d17d86db16356823a7971127d5181a842fef.
//
// Solidity: event Withdraw(address indexed sender, string validator, uint256 reward)
func (_StakingTest *StakingTestFilterer) ParseWithdraw(log types.Log) (*StakingTestWithdraw, error) {
	event := new(StakingTestWithdraw)
	if err := _StakingTest.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
