// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package distribution

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

// DecCoin is an auto generated low-level Go binding around an user-defined struct.
type DecCoin struct {
	Denom     string
	Amount    *big.Int
	Precision uint8
}

// IDistributionMetaData contains all meta data concerning the IDistribution contract.
var IDistributionMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"amount\",\"type\":\"string\"}],\"name\":\"FundCommunityPool\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegatorAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"withdrawAddress\",\"type\":\"address\"}],\"name\":\"SetWithdrawAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegatorAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"withdrawAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"amount\",\"type\":\"string\"}],\"name\":\"WithdrawDelegatorReward\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"amount\",\"type\":\"string\"}],\"name\":\"WithdrawValidatorCommission\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegatorAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"}],\"name\":\"delegationRewards\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"precision\",\"type\":\"uint8\"}],\"internalType\":\"structDecCoin[]\",\"name\":\"rewards\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structCoin[]\",\"name\":\"amount\",\"type\":\"tuple[]\"}],\"name\":\"fundCommunityPool\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"withdrawAddress\",\"type\":\"address\"}],\"name\":\"setWithdrawAddress\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"}],\"name\":\"withdrawDelegatorReward\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structCoin[]\",\"name\":\"amount\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawValidatorCommission\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structCoin[]\",\"name\":\"amount\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IDistributionABI is the input ABI used to generate the binding from.
// Deprecated: Use IDistributionMetaData.ABI instead.
var IDistributionABI = IDistributionMetaData.ABI

// IDistribution is an auto generated Go binding around an Ethereum contract.
type IDistribution struct {
	IDistributionCaller     // Read-only binding to the contract
	IDistributionTransactor // Write-only binding to the contract
	IDistributionFilterer   // Log filterer for contract events
}

// IDistributionCaller is an auto generated read-only Go binding around an Ethereum contract.
type IDistributionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDistributionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IDistributionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDistributionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IDistributionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDistributionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IDistributionSession struct {
	Contract     *IDistribution    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IDistributionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IDistributionCallerSession struct {
	Contract *IDistributionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// IDistributionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IDistributionTransactorSession struct {
	Contract     *IDistributionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// IDistributionRaw is an auto generated low-level Go binding around an Ethereum contract.
type IDistributionRaw struct {
	Contract *IDistribution // Generic contract binding to access the raw methods on
}

// IDistributionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IDistributionCallerRaw struct {
	Contract *IDistributionCaller // Generic read-only contract binding to access the raw methods on
}

// IDistributionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IDistributionTransactorRaw struct {
	Contract *IDistributionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIDistribution creates a new instance of IDistribution, bound to a specific deployed contract.
func NewIDistribution(address common.Address, backend bind.ContractBackend) (*IDistribution, error) {
	contract, err := bindIDistribution(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IDistribution{IDistributionCaller: IDistributionCaller{contract: contract}, IDistributionTransactor: IDistributionTransactor{contract: contract}, IDistributionFilterer: IDistributionFilterer{contract: contract}}, nil
}

// NewIDistributionCaller creates a new read-only instance of IDistribution, bound to a specific deployed contract.
func NewIDistributionCaller(address common.Address, caller bind.ContractCaller) (*IDistributionCaller, error) {
	contract, err := bindIDistribution(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IDistributionCaller{contract: contract}, nil
}

// NewIDistributionTransactor creates a new write-only instance of IDistribution, bound to a specific deployed contract.
func NewIDistributionTransactor(address common.Address, transactor bind.ContractTransactor) (*IDistributionTransactor, error) {
	contract, err := bindIDistribution(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IDistributionTransactor{contract: contract}, nil
}

// NewIDistributionFilterer creates a new log filterer instance of IDistribution, bound to a specific deployed contract.
func NewIDistributionFilterer(address common.Address, filterer bind.ContractFilterer) (*IDistributionFilterer, error) {
	contract, err := bindIDistribution(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IDistributionFilterer{contract: contract}, nil
}

// bindIDistribution binds a generic wrapper to an already deployed contract.
func bindIDistribution(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IDistributionMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IDistribution *IDistributionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IDistribution.Contract.IDistributionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IDistribution *IDistributionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IDistribution.Contract.IDistributionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IDistribution *IDistributionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IDistribution.Contract.IDistributionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IDistribution *IDistributionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IDistribution.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IDistribution *IDistributionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IDistribution.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IDistribution *IDistributionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IDistribution.Contract.contract.Transact(opts, method, params...)
}

// DelegationRewards is a free data retrieval call binding the contract method 0xc9a21b7b.
//
// Solidity: function delegationRewards(address delegatorAddress, address validatorAddress) view returns((string,uint256,uint8)[] rewards)
func (_IDistribution *IDistributionCaller) DelegationRewards(opts *bind.CallOpts, delegatorAddress common.Address, validatorAddress common.Address) ([]DecCoin, error) {
	var out []interface{}
	err := _IDistribution.contract.Call(opts, &out, "delegationRewards", delegatorAddress, validatorAddress)

	if err != nil {
		return *new([]DecCoin), err
	}

	out0 := *abi.ConvertType(out[0], new([]DecCoin)).(*[]DecCoin)

	return out0, err

}

// DelegationRewards is a free data retrieval call binding the contract method 0xc9a21b7b.
//
// Solidity: function delegationRewards(address delegatorAddress, address validatorAddress) view returns((string,uint256,uint8)[] rewards)
func (_IDistribution *IDistributionSession) DelegationRewards(delegatorAddress common.Address, validatorAddress common.Address) ([]DecCoin, error) {
	return _IDistribution.Contract.DelegationRewards(&_IDistribution.CallOpts, delegatorAddress, validatorAddress)
}

// DelegationRewards is a free data retrieval call binding the contract method 0xc9a21b7b.
//
// Solidity: function delegationRewards(address delegatorAddress, address validatorAddress) view returns((string,uint256,uint8)[] rewards)
func (_IDistribution *IDistributionCallerSession) DelegationRewards(delegatorAddress common.Address, validatorAddress common.Address) ([]DecCoin, error) {
	return _IDistribution.Contract.DelegationRewards(&_IDistribution.CallOpts, delegatorAddress, validatorAddress)
}

// FundCommunityPool is a paid mutator transaction binding the contract method 0x515a092c.
//
// Solidity: function fundCommunityPool((string,uint256)[] amount) returns(bool success)
func (_IDistribution *IDistributionTransactor) FundCommunityPool(opts *bind.TransactOpts, amount []Coin) (*types.Transaction, error) {
	return _IDistribution.contract.Transact(opts, "fundCommunityPool", amount)
}

// FundCommunityPool is a paid mutator transaction binding the contract method 0x515a092c.
//
// Solidity: function fundCommunityPool((string,uint256)[] amount) returns(bool success)
func (_IDistribution *IDistributionSession) FundCommunityPool(amount []Coin) (*types.Transaction, error) {
	return _IDistribution.Contract.FundCommunityPool(&_IDistribution.TransactOpts, amount)
}

// FundCommunityPool is a paid mutator transaction binding the contract method 0x515a092c.
//
// Solidity: function fundCommunityPool((string,uint256)[] amount) returns(bool success)
func (_IDistribution *IDistributionTransactorSession) FundCommunityPool(amount []Coin) (*types.Transaction, error) {
	return _IDistribution.Contract.FundCommunityPool(&_IDistribution.TransactOpts, amount)
}

// SetWithdrawAddress is a paid mutator transaction binding the contract method 0x3ab1a494.
//
// Solidity: function setWithdrawAddress(address withdrawAddress) returns(bool success)
func (_IDistribution *IDistributionTransactor) SetWithdrawAddress(opts *bind.TransactOpts, withdrawAddress common.Address) (*types.Transaction, error) {
	return _IDistribution.contract.Transact(opts, "setWithdrawAddress", withdrawAddress)
}

// SetWithdrawAddress is a paid mutator transaction binding the contract method 0x3ab1a494.
//
// Solidity: function setWithdrawAddress(address withdrawAddress) returns(bool success)
func (_IDistribution *IDistributionSession) SetWithdrawAddress(withdrawAddress common.Address) (*types.Transaction, error) {
	return _IDistribution.Contract.SetWithdrawAddress(&_IDistribution.TransactOpts, withdrawAddress)
}

// SetWithdrawAddress is a paid mutator transaction binding the contract method 0x3ab1a494.
//
// Solidity: function setWithdrawAddress(address withdrawAddress) returns(bool success)
func (_IDistribution *IDistributionTransactorSession) SetWithdrawAddress(withdrawAddress common.Address) (*types.Transaction, error) {
	return _IDistribution.Contract.SetWithdrawAddress(&_IDistribution.TransactOpts, withdrawAddress)
}

// WithdrawDelegatorReward is a paid mutator transaction binding the contract method 0x346683f0.
//
// Solidity: function withdrawDelegatorReward(address validatorAddress) returns((string,uint256)[] amount)
func (_IDistribution *IDistributionTransactor) WithdrawDelegatorReward(opts *bind.TransactOpts, validatorAddress common.Address) (*types.Transaction, error) {
	return _IDistribution.contract.Transact(opts, "withdrawDelegatorReward", validatorAddress)
}

// WithdrawDelegatorReward is a paid mutator transaction binding the contract method 0x346683f0.
//
// Solidity: function withdrawDelegatorReward(address validatorAddress) returns((string,uint256)[] amount)
func (_IDistribution *IDistributionSession) WithdrawDelegatorReward(validatorAddress common.Address) (*types.Transaction, error) {
	return _IDistribution.Contract.WithdrawDelegatorReward(&_IDistribution.TransactOpts, validatorAddress)
}

// WithdrawDelegatorReward is a paid mutator transaction binding the contract method 0x346683f0.
//
// Solidity: function withdrawDelegatorReward(address validatorAddress) returns((string,uint256)[] amount)
func (_IDistribution *IDistributionTransactorSession) WithdrawDelegatorReward(validatorAddress common.Address) (*types.Transaction, error) {
	return _IDistribution.Contract.WithdrawDelegatorReward(&_IDistribution.TransactOpts, validatorAddress)
}

// WithdrawValidatorCommission is a paid mutator transaction binding the contract method 0x0bde076d.
//
// Solidity: function withdrawValidatorCommission() returns((string,uint256)[] amount)
func (_IDistribution *IDistributionTransactor) WithdrawValidatorCommission(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IDistribution.contract.Transact(opts, "withdrawValidatorCommission")
}

// WithdrawValidatorCommission is a paid mutator transaction binding the contract method 0x0bde076d.
//
// Solidity: function withdrawValidatorCommission() returns((string,uint256)[] amount)
func (_IDistribution *IDistributionSession) WithdrawValidatorCommission() (*types.Transaction, error) {
	return _IDistribution.Contract.WithdrawValidatorCommission(&_IDistribution.TransactOpts)
}

// WithdrawValidatorCommission is a paid mutator transaction binding the contract method 0x0bde076d.
//
// Solidity: function withdrawValidatorCommission() returns((string,uint256)[] amount)
func (_IDistribution *IDistributionTransactorSession) WithdrawValidatorCommission() (*types.Transaction, error) {
	return _IDistribution.Contract.WithdrawValidatorCommission(&_IDistribution.TransactOpts)
}

// IDistributionFundCommunityPoolIterator is returned from FilterFundCommunityPool and is used to iterate over the raw logs and unpacked data for FundCommunityPool events raised by the IDistribution contract.
type IDistributionFundCommunityPoolIterator struct {
	Event *IDistributionFundCommunityPool // Event containing the contract specifics and raw log

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
func (it *IDistributionFundCommunityPoolIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IDistributionFundCommunityPool)
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
		it.Event = new(IDistributionFundCommunityPool)
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
func (it *IDistributionFundCommunityPoolIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IDistributionFundCommunityPoolIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IDistributionFundCommunityPool represents a FundCommunityPool event raised by the IDistribution contract.
type IDistributionFundCommunityPool struct {
	Depositor common.Address
	Amount    string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterFundCommunityPool is a free log retrieval operation binding the contract event 0x270d117e6a45828d092cd049178c5c1b739117dffbe4a19865b08565e288d1ed.
//
// Solidity: event FundCommunityPool(address indexed depositor, string amount)
func (_IDistribution *IDistributionFilterer) FilterFundCommunityPool(opts *bind.FilterOpts, depositor []common.Address) (*IDistributionFundCommunityPoolIterator, error) {

	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}

	logs, sub, err := _IDistribution.contract.FilterLogs(opts, "FundCommunityPool", depositorRule)
	if err != nil {
		return nil, err
	}
	return &IDistributionFundCommunityPoolIterator{contract: _IDistribution.contract, event: "FundCommunityPool", logs: logs, sub: sub}, nil
}

// WatchFundCommunityPool is a free log subscription operation binding the contract event 0x270d117e6a45828d092cd049178c5c1b739117dffbe4a19865b08565e288d1ed.
//
// Solidity: event FundCommunityPool(address indexed depositor, string amount)
func (_IDistribution *IDistributionFilterer) WatchFundCommunityPool(opts *bind.WatchOpts, sink chan<- *IDistributionFundCommunityPool, depositor []common.Address) (event.Subscription, error) {

	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}

	logs, sub, err := _IDistribution.contract.WatchLogs(opts, "FundCommunityPool", depositorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IDistributionFundCommunityPool)
				if err := _IDistribution.contract.UnpackLog(event, "FundCommunityPool", log); err != nil {
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

// ParseFundCommunityPool is a log parse operation binding the contract event 0x270d117e6a45828d092cd049178c5c1b739117dffbe4a19865b08565e288d1ed.
//
// Solidity: event FundCommunityPool(address indexed depositor, string amount)
func (_IDistribution *IDistributionFilterer) ParseFundCommunityPool(log types.Log) (*IDistributionFundCommunityPool, error) {
	event := new(IDistributionFundCommunityPool)
	if err := _IDistribution.contract.UnpackLog(event, "FundCommunityPool", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IDistributionSetWithdrawAddressIterator is returned from FilterSetWithdrawAddress and is used to iterate over the raw logs and unpacked data for SetWithdrawAddress events raised by the IDistribution contract.
type IDistributionSetWithdrawAddressIterator struct {
	Event *IDistributionSetWithdrawAddress // Event containing the contract specifics and raw log

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
func (it *IDistributionSetWithdrawAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IDistributionSetWithdrawAddress)
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
		it.Event = new(IDistributionSetWithdrawAddress)
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
func (it *IDistributionSetWithdrawAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IDistributionSetWithdrawAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IDistributionSetWithdrawAddress represents a SetWithdrawAddress event raised by the IDistribution contract.
type IDistributionSetWithdrawAddress struct {
	DelegatorAddress common.Address
	WithdrawAddress  common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterSetWithdrawAddress is a free log retrieval operation binding the contract event 0xae416f064415339eb2fc98ef48a0fc06ee07e3b33d469e001c477eac6e68947c.
//
// Solidity: event SetWithdrawAddress(address indexed delegatorAddress, address indexed withdrawAddress)
func (_IDistribution *IDistributionFilterer) FilterSetWithdrawAddress(opts *bind.FilterOpts, delegatorAddress []common.Address, withdrawAddress []common.Address) (*IDistributionSetWithdrawAddressIterator, error) {

	var delegatorAddressRule []interface{}
	for _, delegatorAddressItem := range delegatorAddress {
		delegatorAddressRule = append(delegatorAddressRule, delegatorAddressItem)
	}
	var withdrawAddressRule []interface{}
	for _, withdrawAddressItem := range withdrawAddress {
		withdrawAddressRule = append(withdrawAddressRule, withdrawAddressItem)
	}

	logs, sub, err := _IDistribution.contract.FilterLogs(opts, "SetWithdrawAddress", delegatorAddressRule, withdrawAddressRule)
	if err != nil {
		return nil, err
	}
	return &IDistributionSetWithdrawAddressIterator{contract: _IDistribution.contract, event: "SetWithdrawAddress", logs: logs, sub: sub}, nil
}

// WatchSetWithdrawAddress is a free log subscription operation binding the contract event 0xae416f064415339eb2fc98ef48a0fc06ee07e3b33d469e001c477eac6e68947c.
//
// Solidity: event SetWithdrawAddress(address indexed delegatorAddress, address indexed withdrawAddress)
func (_IDistribution *IDistributionFilterer) WatchSetWithdrawAddress(opts *bind.WatchOpts, sink chan<- *IDistributionSetWithdrawAddress, delegatorAddress []common.Address, withdrawAddress []common.Address) (event.Subscription, error) {

	var delegatorAddressRule []interface{}
	for _, delegatorAddressItem := range delegatorAddress {
		delegatorAddressRule = append(delegatorAddressRule, delegatorAddressItem)
	}
	var withdrawAddressRule []interface{}
	for _, withdrawAddressItem := range withdrawAddress {
		withdrawAddressRule = append(withdrawAddressRule, withdrawAddressItem)
	}

	logs, sub, err := _IDistribution.contract.WatchLogs(opts, "SetWithdrawAddress", delegatorAddressRule, withdrawAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IDistributionSetWithdrawAddress)
				if err := _IDistribution.contract.UnpackLog(event, "SetWithdrawAddress", log); err != nil {
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

// ParseSetWithdrawAddress is a log parse operation binding the contract event 0xae416f064415339eb2fc98ef48a0fc06ee07e3b33d469e001c477eac6e68947c.
//
// Solidity: event SetWithdrawAddress(address indexed delegatorAddress, address indexed withdrawAddress)
func (_IDistribution *IDistributionFilterer) ParseSetWithdrawAddress(log types.Log) (*IDistributionSetWithdrawAddress, error) {
	event := new(IDistributionSetWithdrawAddress)
	if err := _IDistribution.contract.UnpackLog(event, "SetWithdrawAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IDistributionWithdrawDelegatorRewardIterator is returned from FilterWithdrawDelegatorReward and is used to iterate over the raw logs and unpacked data for WithdrawDelegatorReward events raised by the IDistribution contract.
type IDistributionWithdrawDelegatorRewardIterator struct {
	Event *IDistributionWithdrawDelegatorReward // Event containing the contract specifics and raw log

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
func (it *IDistributionWithdrawDelegatorRewardIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IDistributionWithdrawDelegatorReward)
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
		it.Event = new(IDistributionWithdrawDelegatorReward)
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
func (it *IDistributionWithdrawDelegatorRewardIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IDistributionWithdrawDelegatorRewardIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IDistributionWithdrawDelegatorReward represents a WithdrawDelegatorReward event raised by the IDistribution contract.
type IDistributionWithdrawDelegatorReward struct {
	DelegatorAddress common.Address
	WithdrawAddress  common.Address
	Amount           string
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterWithdrawDelegatorReward is a free log retrieval operation binding the contract event 0xcb466dd767c14616f58fbfdb7792bb89fe8d7b264ac68471f7e09d73705adfd0.
//
// Solidity: event WithdrawDelegatorReward(address indexed delegatorAddress, address indexed withdrawAddress, string amount)
func (_IDistribution *IDistributionFilterer) FilterWithdrawDelegatorReward(opts *bind.FilterOpts, delegatorAddress []common.Address, withdrawAddress []common.Address) (*IDistributionWithdrawDelegatorRewardIterator, error) {

	var delegatorAddressRule []interface{}
	for _, delegatorAddressItem := range delegatorAddress {
		delegatorAddressRule = append(delegatorAddressRule, delegatorAddressItem)
	}
	var withdrawAddressRule []interface{}
	for _, withdrawAddressItem := range withdrawAddress {
		withdrawAddressRule = append(withdrawAddressRule, withdrawAddressItem)
	}

	logs, sub, err := _IDistribution.contract.FilterLogs(opts, "WithdrawDelegatorReward", delegatorAddressRule, withdrawAddressRule)
	if err != nil {
		return nil, err
	}
	return &IDistributionWithdrawDelegatorRewardIterator{contract: _IDistribution.contract, event: "WithdrawDelegatorReward", logs: logs, sub: sub}, nil
}

// WatchWithdrawDelegatorReward is a free log subscription operation binding the contract event 0xcb466dd767c14616f58fbfdb7792bb89fe8d7b264ac68471f7e09d73705adfd0.
//
// Solidity: event WithdrawDelegatorReward(address indexed delegatorAddress, address indexed withdrawAddress, string amount)
func (_IDistribution *IDistributionFilterer) WatchWithdrawDelegatorReward(opts *bind.WatchOpts, sink chan<- *IDistributionWithdrawDelegatorReward, delegatorAddress []common.Address, withdrawAddress []common.Address) (event.Subscription, error) {

	var delegatorAddressRule []interface{}
	for _, delegatorAddressItem := range delegatorAddress {
		delegatorAddressRule = append(delegatorAddressRule, delegatorAddressItem)
	}
	var withdrawAddressRule []interface{}
	for _, withdrawAddressItem := range withdrawAddress {
		withdrawAddressRule = append(withdrawAddressRule, withdrawAddressItem)
	}

	logs, sub, err := _IDistribution.contract.WatchLogs(opts, "WithdrawDelegatorReward", delegatorAddressRule, withdrawAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IDistributionWithdrawDelegatorReward)
				if err := _IDistribution.contract.UnpackLog(event, "WithdrawDelegatorReward", log); err != nil {
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

// ParseWithdrawDelegatorReward is a log parse operation binding the contract event 0xcb466dd767c14616f58fbfdb7792bb89fe8d7b264ac68471f7e09d73705adfd0.
//
// Solidity: event WithdrawDelegatorReward(address indexed delegatorAddress, address indexed withdrawAddress, string amount)
func (_IDistribution *IDistributionFilterer) ParseWithdrawDelegatorReward(log types.Log) (*IDistributionWithdrawDelegatorReward, error) {
	event := new(IDistributionWithdrawDelegatorReward)
	if err := _IDistribution.contract.UnpackLog(event, "WithdrawDelegatorReward", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IDistributionWithdrawValidatorCommissionIterator is returned from FilterWithdrawValidatorCommission and is used to iterate over the raw logs and unpacked data for WithdrawValidatorCommission events raised by the IDistribution contract.
type IDistributionWithdrawValidatorCommissionIterator struct {
	Event *IDistributionWithdrawValidatorCommission // Event containing the contract specifics and raw log

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
func (it *IDistributionWithdrawValidatorCommissionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IDistributionWithdrawValidatorCommission)
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
		it.Event = new(IDistributionWithdrawValidatorCommission)
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
func (it *IDistributionWithdrawValidatorCommissionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IDistributionWithdrawValidatorCommissionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IDistributionWithdrawValidatorCommission represents a WithdrawValidatorCommission event raised by the IDistribution contract.
type IDistributionWithdrawValidatorCommission struct {
	ValidatorAddress common.Address
	Amount           string
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterWithdrawValidatorCommission is a free log retrieval operation binding the contract event 0x35d850023c3dbe4c2b9373eb44803767ed387ca0cdf596e14dc3f43331c7a042.
//
// Solidity: event WithdrawValidatorCommission(address indexed validatorAddress, string amount)
func (_IDistribution *IDistributionFilterer) FilterWithdrawValidatorCommission(opts *bind.FilterOpts, validatorAddress []common.Address) (*IDistributionWithdrawValidatorCommissionIterator, error) {

	var validatorAddressRule []interface{}
	for _, validatorAddressItem := range validatorAddress {
		validatorAddressRule = append(validatorAddressRule, validatorAddressItem)
	}

	logs, sub, err := _IDistribution.contract.FilterLogs(opts, "WithdrawValidatorCommission", validatorAddressRule)
	if err != nil {
		return nil, err
	}
	return &IDistributionWithdrawValidatorCommissionIterator{contract: _IDistribution.contract, event: "WithdrawValidatorCommission", logs: logs, sub: sub}, nil
}

// WatchWithdrawValidatorCommission is a free log subscription operation binding the contract event 0x35d850023c3dbe4c2b9373eb44803767ed387ca0cdf596e14dc3f43331c7a042.
//
// Solidity: event WithdrawValidatorCommission(address indexed validatorAddress, string amount)
func (_IDistribution *IDistributionFilterer) WatchWithdrawValidatorCommission(opts *bind.WatchOpts, sink chan<- *IDistributionWithdrawValidatorCommission, validatorAddress []common.Address) (event.Subscription, error) {

	var validatorAddressRule []interface{}
	for _, validatorAddressItem := range validatorAddress {
		validatorAddressRule = append(validatorAddressRule, validatorAddressItem)
	}

	logs, sub, err := _IDistribution.contract.WatchLogs(opts, "WithdrawValidatorCommission", validatorAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IDistributionWithdrawValidatorCommission)
				if err := _IDistribution.contract.UnpackLog(event, "WithdrawValidatorCommission", log); err != nil {
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

// ParseWithdrawValidatorCommission is a log parse operation binding the contract event 0x35d850023c3dbe4c2b9373eb44803767ed387ca0cdf596e14dc3f43331c7a042.
//
// Solidity: event WithdrawValidatorCommission(address indexed validatorAddress, string amount)
func (_IDistribution *IDistributionFilterer) ParseWithdrawValidatorCommission(log types.Log) (*IDistributionWithdrawValidatorCommission, error) {
	event := new(IDistributionWithdrawValidatorCommission)
	if err := _IDistribution.contract.UnpackLog(event, "WithdrawValidatorCommission", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
