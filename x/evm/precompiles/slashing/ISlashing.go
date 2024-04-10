// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package slashing

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

// ISlashingMetaData contains all meta data concerning the ISlashing contract.
var ISlashingMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"Unjail\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"unjail\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ISlashingABI is the input ABI used to generate the binding from.
// Deprecated: Use ISlashingMetaData.ABI instead.
var ISlashingABI = ISlashingMetaData.ABI

// ISlashing is an auto generated Go binding around an Ethereum contract.
type ISlashing struct {
	ISlashingCaller     // Read-only binding to the contract
	ISlashingTransactor // Write-only binding to the contract
	ISlashingFilterer   // Log filterer for contract events
}

// ISlashingCaller is an auto generated read-only Go binding around an Ethereum contract.
type ISlashingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISlashingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ISlashingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISlashingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ISlashingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISlashingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ISlashingSession struct {
	Contract     *ISlashing        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ISlashingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ISlashingCallerSession struct {
	Contract *ISlashingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ISlashingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ISlashingTransactorSession struct {
	Contract     *ISlashingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ISlashingRaw is an auto generated low-level Go binding around an Ethereum contract.
type ISlashingRaw struct {
	Contract *ISlashing // Generic contract binding to access the raw methods on
}

// ISlashingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ISlashingCallerRaw struct {
	Contract *ISlashingCaller // Generic read-only contract binding to access the raw methods on
}

// ISlashingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ISlashingTransactorRaw struct {
	Contract *ISlashingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewISlashing creates a new instance of ISlashing, bound to a specific deployed contract.
func NewISlashing(address common.Address, backend bind.ContractBackend) (*ISlashing, error) {
	contract, err := bindISlashing(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ISlashing{ISlashingCaller: ISlashingCaller{contract: contract}, ISlashingTransactor: ISlashingTransactor{contract: contract}, ISlashingFilterer: ISlashingFilterer{contract: contract}}, nil
}

// NewISlashingCaller creates a new read-only instance of ISlashing, bound to a specific deployed contract.
func NewISlashingCaller(address common.Address, caller bind.ContractCaller) (*ISlashingCaller, error) {
	contract, err := bindISlashing(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ISlashingCaller{contract: contract}, nil
}

// NewISlashingTransactor creates a new write-only instance of ISlashing, bound to a specific deployed contract.
func NewISlashingTransactor(address common.Address, transactor bind.ContractTransactor) (*ISlashingTransactor, error) {
	contract, err := bindISlashing(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ISlashingTransactor{contract: contract}, nil
}

// NewISlashingFilterer creates a new log filterer instance of ISlashing, bound to a specific deployed contract.
func NewISlashingFilterer(address common.Address, filterer bind.ContractFilterer) (*ISlashingFilterer, error) {
	contract, err := bindISlashing(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ISlashingFilterer{contract: contract}, nil
}

// bindISlashing binds a generic wrapper to an already deployed contract.
func bindISlashing(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ISlashingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISlashing *ISlashingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISlashing.Contract.ISlashingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISlashing *ISlashingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISlashing.Contract.ISlashingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISlashing *ISlashingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISlashing.Contract.ISlashingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISlashing *ISlashingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISlashing.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISlashing *ISlashingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISlashing.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISlashing *ISlashingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISlashing.Contract.contract.Transact(opts, method, params...)
}

// Unjail is a paid mutator transaction binding the contract method 0xf679d305.
//
// Solidity: function unjail() returns(bool success)
func (_ISlashing *ISlashingTransactor) Unjail(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISlashing.contract.Transact(opts, "unjail")
}

// Unjail is a paid mutator transaction binding the contract method 0xf679d305.
//
// Solidity: function unjail() returns(bool success)
func (_ISlashing *ISlashingSession) Unjail() (*types.Transaction, error) {
	return _ISlashing.Contract.Unjail(&_ISlashing.TransactOpts)
}

// Unjail is a paid mutator transaction binding the contract method 0xf679d305.
//
// Solidity: function unjail() returns(bool success)
func (_ISlashing *ISlashingTransactorSession) Unjail() (*types.Transaction, error) {
	return _ISlashing.Contract.Unjail(&_ISlashing.TransactOpts)
}

// ISlashingUnjailIterator is returned from FilterUnjail and is used to iterate over the raw logs and unpacked data for Unjail events raised by the ISlashing contract.
type ISlashingUnjailIterator struct {
	Event *ISlashingUnjail // Event containing the contract specifics and raw log

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
func (it *ISlashingUnjailIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ISlashingUnjail)
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
		it.Event = new(ISlashingUnjail)
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
func (it *ISlashingUnjailIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ISlashingUnjailIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ISlashingUnjail represents a Unjail event raised by the ISlashing contract.
type ISlashingUnjail struct {
	Validator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUnjail is a free log retrieval operation binding the contract event 0xc3ef55ddda4bc9300706e15ab3aed03c762d8afd43a7d358a7b9503cb39f281b.
//
// Solidity: event Unjail(address indexed validator)
func (_ISlashing *ISlashingFilterer) FilterUnjail(opts *bind.FilterOpts, validator []common.Address) (*ISlashingUnjailIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _ISlashing.contract.FilterLogs(opts, "Unjail", validatorRule)
	if err != nil {
		return nil, err
	}
	return &ISlashingUnjailIterator{contract: _ISlashing.contract, event: "Unjail", logs: logs, sub: sub}, nil
}

// WatchUnjail is a free log subscription operation binding the contract event 0xc3ef55ddda4bc9300706e15ab3aed03c762d8afd43a7d358a7b9503cb39f281b.
//
// Solidity: event Unjail(address indexed validator)
func (_ISlashing *ISlashingFilterer) WatchUnjail(opts *bind.WatchOpts, sink chan<- *ISlashingUnjail, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _ISlashing.contract.WatchLogs(opts, "Unjail", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ISlashingUnjail)
				if err := _ISlashing.contract.UnpackLog(event, "Unjail", log); err != nil {
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

// ParseUnjail is a log parse operation binding the contract event 0xc3ef55ddda4bc9300706e15ab3aed03c762d8afd43a7d358a7b9503cb39f281b.
//
// Solidity: event Unjail(address indexed validator)
func (_ISlashing *ISlashingFilterer) ParseUnjail(log types.Log) (*ISlashingUnjail, error) {
	event := new(ISlashingUnjail)
	if err := _ISlashing.contract.UnpackLog(event, "Unjail", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
