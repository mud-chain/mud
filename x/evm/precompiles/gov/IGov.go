// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package gov

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

// IGovMetaData contains all meta data concerning the IGov contract.
var IGovMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"proposalId\",\"type\":\"uint64\"}],\"name\":\"LegacySubmitProposal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"proposalId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"int32\",\"name\":\"option\",\"type\":\"int32\"}],\"name\":\"Vote\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structCoin\",\"name\":\"initialDeposit\",\"type\":\"tuple\"}],\"name\":\"legacySubmitProposal\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"proposalId\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"proposalId\",\"type\":\"uint64\"},{\"internalType\":\"int32\",\"name\":\"option\",\"type\":\"int32\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"}],\"name\":\"vote\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IGovABI is the input ABI used to generate the binding from.
// Deprecated: Use IGovMetaData.ABI instead.
var IGovABI = IGovMetaData.ABI

// IGov is an auto generated Go binding around an Ethereum contract.
type IGov struct {
	IGovCaller     // Read-only binding to the contract
	IGovTransactor // Write-only binding to the contract
	IGovFilterer   // Log filterer for contract events
}

// IGovCaller is an auto generated read-only Go binding around an Ethereum contract.
type IGovCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGovTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IGovTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGovFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IGovFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGovSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IGovSession struct {
	Contract     *IGov             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IGovCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IGovCallerSession struct {
	Contract *IGovCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IGovTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IGovTransactorSession struct {
	Contract     *IGovTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IGovRaw is an auto generated low-level Go binding around an Ethereum contract.
type IGovRaw struct {
	Contract *IGov // Generic contract binding to access the raw methods on
}

// IGovCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IGovCallerRaw struct {
	Contract *IGovCaller // Generic read-only contract binding to access the raw methods on
}

// IGovTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IGovTransactorRaw struct {
	Contract *IGovTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIGov creates a new instance of IGov, bound to a specific deployed contract.
func NewIGov(address common.Address, backend bind.ContractBackend) (*IGov, error) {
	contract, err := bindIGov(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IGov{IGovCaller: IGovCaller{contract: contract}, IGovTransactor: IGovTransactor{contract: contract}, IGovFilterer: IGovFilterer{contract: contract}}, nil
}

// NewIGovCaller creates a new read-only instance of IGov, bound to a specific deployed contract.
func NewIGovCaller(address common.Address, caller bind.ContractCaller) (*IGovCaller, error) {
	contract, err := bindIGov(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IGovCaller{contract: contract}, nil
}

// NewIGovTransactor creates a new write-only instance of IGov, bound to a specific deployed contract.
func NewIGovTransactor(address common.Address, transactor bind.ContractTransactor) (*IGovTransactor, error) {
	contract, err := bindIGov(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IGovTransactor{contract: contract}, nil
}

// NewIGovFilterer creates a new log filterer instance of IGov, bound to a specific deployed contract.
func NewIGovFilterer(address common.Address, filterer bind.ContractFilterer) (*IGovFilterer, error) {
	contract, err := bindIGov(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IGovFilterer{contract: contract}, nil
}

// bindIGov binds a generic wrapper to an already deployed contract.
func bindIGov(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IGovMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IGov *IGovRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IGov.Contract.IGovCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IGov *IGovRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IGov.Contract.IGovTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IGov *IGovRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IGov.Contract.IGovTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IGov *IGovCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IGov.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IGov *IGovTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IGov.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IGov *IGovTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IGov.Contract.contract.Transact(opts, method, params...)
}

// LegacySubmitProposal is a paid mutator transaction binding the contract method 0xbc4b9c1f.
//
// Solidity: function legacySubmitProposal(string title, string description, (string,uint256) initialDeposit) returns(uint64 proposalId)
func (_IGov *IGovTransactor) LegacySubmitProposal(opts *bind.TransactOpts, title string, description string, initialDeposit Coin) (*types.Transaction, error) {
	return _IGov.contract.Transact(opts, "legacySubmitProposal", title, description, initialDeposit)
}

// LegacySubmitProposal is a paid mutator transaction binding the contract method 0xbc4b9c1f.
//
// Solidity: function legacySubmitProposal(string title, string description, (string,uint256) initialDeposit) returns(uint64 proposalId)
func (_IGov *IGovSession) LegacySubmitProposal(title string, description string, initialDeposit Coin) (*types.Transaction, error) {
	return _IGov.Contract.LegacySubmitProposal(&_IGov.TransactOpts, title, description, initialDeposit)
}

// LegacySubmitProposal is a paid mutator transaction binding the contract method 0xbc4b9c1f.
//
// Solidity: function legacySubmitProposal(string title, string description, (string,uint256) initialDeposit) returns(uint64 proposalId)
func (_IGov *IGovTransactorSession) LegacySubmitProposal(title string, description string, initialDeposit Coin) (*types.Transaction, error) {
	return _IGov.Contract.LegacySubmitProposal(&_IGov.TransactOpts, title, description, initialDeposit)
}

// Vote is a paid mutator transaction binding the contract method 0x19f7a0fb.
//
// Solidity: function vote(uint64 proposalId, int32 option, string metadata) returns(bool success)
func (_IGov *IGovTransactor) Vote(opts *bind.TransactOpts, proposalId uint64, option int32, metadata string) (*types.Transaction, error) {
	return _IGov.contract.Transact(opts, "vote", proposalId, option, metadata)
}

// Vote is a paid mutator transaction binding the contract method 0x19f7a0fb.
//
// Solidity: function vote(uint64 proposalId, int32 option, string metadata) returns(bool success)
func (_IGov *IGovSession) Vote(proposalId uint64, option int32, metadata string) (*types.Transaction, error) {
	return _IGov.Contract.Vote(&_IGov.TransactOpts, proposalId, option, metadata)
}

// Vote is a paid mutator transaction binding the contract method 0x19f7a0fb.
//
// Solidity: function vote(uint64 proposalId, int32 option, string metadata) returns(bool success)
func (_IGov *IGovTransactorSession) Vote(proposalId uint64, option int32, metadata string) (*types.Transaction, error) {
	return _IGov.Contract.Vote(&_IGov.TransactOpts, proposalId, option, metadata)
}

// IGovLegacySubmitProposalIterator is returned from FilterLegacySubmitProposal and is used to iterate over the raw logs and unpacked data for LegacySubmitProposal events raised by the IGov contract.
type IGovLegacySubmitProposalIterator struct {
	Event *IGovLegacySubmitProposal // Event containing the contract specifics and raw log

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
func (it *IGovLegacySubmitProposalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IGovLegacySubmitProposal)
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
		it.Event = new(IGovLegacySubmitProposal)
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
func (it *IGovLegacySubmitProposalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IGovLegacySubmitProposalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IGovLegacySubmitProposal represents a LegacySubmitProposal event raised by the IGov contract.
type IGovLegacySubmitProposal struct {
	Proposer   common.Address
	ProposalId uint64
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLegacySubmitProposal is a free log retrieval operation binding the contract event 0x749a006802f9b6dde9b6ee9f8f3a317d8a98cb9017684a7b59bfcf21bf93f477.
//
// Solidity: event LegacySubmitProposal(address indexed proposer, uint64 proposalId)
func (_IGov *IGovFilterer) FilterLegacySubmitProposal(opts *bind.FilterOpts, proposer []common.Address) (*IGovLegacySubmitProposalIterator, error) {

	var proposerRule []interface{}
	for _, proposerItem := range proposer {
		proposerRule = append(proposerRule, proposerItem)
	}

	logs, sub, err := _IGov.contract.FilterLogs(opts, "LegacySubmitProposal", proposerRule)
	if err != nil {
		return nil, err
	}
	return &IGovLegacySubmitProposalIterator{contract: _IGov.contract, event: "LegacySubmitProposal", logs: logs, sub: sub}, nil
}

// WatchLegacySubmitProposal is a free log subscription operation binding the contract event 0x749a006802f9b6dde9b6ee9f8f3a317d8a98cb9017684a7b59bfcf21bf93f477.
//
// Solidity: event LegacySubmitProposal(address indexed proposer, uint64 proposalId)
func (_IGov *IGovFilterer) WatchLegacySubmitProposal(opts *bind.WatchOpts, sink chan<- *IGovLegacySubmitProposal, proposer []common.Address) (event.Subscription, error) {

	var proposerRule []interface{}
	for _, proposerItem := range proposer {
		proposerRule = append(proposerRule, proposerItem)
	}

	logs, sub, err := _IGov.contract.WatchLogs(opts, "LegacySubmitProposal", proposerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IGovLegacySubmitProposal)
				if err := _IGov.contract.UnpackLog(event, "LegacySubmitProposal", log); err != nil {
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

// ParseLegacySubmitProposal is a log parse operation binding the contract event 0x749a006802f9b6dde9b6ee9f8f3a317d8a98cb9017684a7b59bfcf21bf93f477.
//
// Solidity: event LegacySubmitProposal(address indexed proposer, uint64 proposalId)
func (_IGov *IGovFilterer) ParseLegacySubmitProposal(log types.Log) (*IGovLegacySubmitProposal, error) {
	event := new(IGovLegacySubmitProposal)
	if err := _IGov.contract.UnpackLog(event, "LegacySubmitProposal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IGovVoteIterator is returned from FilterVote and is used to iterate over the raw logs and unpacked data for Vote events raised by the IGov contract.
type IGovVoteIterator struct {
	Event *IGovVote // Event containing the contract specifics and raw log

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
func (it *IGovVoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IGovVote)
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
		it.Event = new(IGovVote)
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
func (it *IGovVoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IGovVoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IGovVote represents a Vote event raised by the IGov contract.
type IGovVote struct {
	Voter      common.Address
	ProposalId uint64
	Option     int32
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterVote is a free log retrieval operation binding the contract event 0x136803a3fea2cd3a4136956fbdf8bc0a7d375ee66ae3ca58caba112be1343d4c.
//
// Solidity: event Vote(address indexed voter, uint64 proposalId, int32 option)
func (_IGov *IGovFilterer) FilterVote(opts *bind.FilterOpts, voter []common.Address) (*IGovVoteIterator, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _IGov.contract.FilterLogs(opts, "Vote", voterRule)
	if err != nil {
		return nil, err
	}
	return &IGovVoteIterator{contract: _IGov.contract, event: "Vote", logs: logs, sub: sub}, nil
}

// WatchVote is a free log subscription operation binding the contract event 0x136803a3fea2cd3a4136956fbdf8bc0a7d375ee66ae3ca58caba112be1343d4c.
//
// Solidity: event Vote(address indexed voter, uint64 proposalId, int32 option)
func (_IGov *IGovFilterer) WatchVote(opts *bind.WatchOpts, sink chan<- *IGovVote, voter []common.Address) (event.Subscription, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _IGov.contract.WatchLogs(opts, "Vote", voterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IGovVote)
				if err := _IGov.contract.UnpackLog(event, "Vote", log); err != nil {
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

// ParseVote is a log parse operation binding the contract event 0x136803a3fea2cd3a4136956fbdf8bc0a7d375ee66ae3ca58caba112be1343d4c.
//
// Solidity: event Vote(address indexed voter, uint64 proposalId, int32 option)
func (_IGov *IGovFilterer) ParseVote(log types.Log) (*IGovVote, error) {
	event := new(IGovVote)
	if err := _IGov.contract.UnpackLog(event, "Vote", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
