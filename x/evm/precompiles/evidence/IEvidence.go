// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package evidence

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

// Equivocation is an auto generated low-level Go binding around an user-defined struct.
type Equivocation struct {
	Height           int64
	Time             int64
	Power            int64
	ConsensusAddress common.Address
	EvidenceHash     string
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

// IEvidenceMetaData contains all meta data concerning the IEvidence contract.
var IEvidenceMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"submitter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"evidenceHash\",\"type\":\"string\"}],\"name\":\"SubmitEvidence\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"offset\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"limit\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"countTotal\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"reverse\",\"type\":\"bool\"}],\"internalType\":\"structPageRequest\",\"name\":\"pagination\",\"type\":\"tuple\"}],\"name\":\"allEvidence\",\"outputs\":[{\"components\":[{\"internalType\":\"int64\",\"name\":\"height\",\"type\":\"int64\"},{\"internalType\":\"int64\",\"name\":\"time\",\"type\":\"int64\"},{\"internalType\":\"int64\",\"name\":\"power\",\"type\":\"int64\"},{\"internalType\":\"address\",\"name\":\"consensusAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"evidenceHash\",\"type\":\"string\"}],\"internalType\":\"structEquivocation[]\",\"name\":\"evidence\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"nextKey\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"total\",\"type\":\"uint64\"}],\"internalType\":\"structPageResponse\",\"name\":\"pageResponse\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"evidenceHash\",\"type\":\"string\"}],\"name\":\"evidence\",\"outputs\":[{\"components\":[{\"internalType\":\"int64\",\"name\":\"height\",\"type\":\"int64\"},{\"internalType\":\"int64\",\"name\":\"time\",\"type\":\"int64\"},{\"internalType\":\"int64\",\"name\":\"power\",\"type\":\"int64\"},{\"internalType\":\"address\",\"name\":\"consensusAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"evidenceHash\",\"type\":\"string\"}],\"internalType\":\"structEquivocation\",\"name\":\"evidence\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int64\",\"name\":\"height\",\"type\":\"int64\"},{\"internalType\":\"int64\",\"name\":\"time\",\"type\":\"int64\"},{\"internalType\":\"int64\",\"name\":\"power\",\"type\":\"int64\"},{\"internalType\":\"address\",\"name\":\"consensusAddress\",\"type\":\"address\"}],\"name\":\"submitEvidence\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"hash\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IEvidenceABI is the input ABI used to generate the binding from.
// Deprecated: Use IEvidenceMetaData.ABI instead.
var IEvidenceABI = IEvidenceMetaData.ABI

// IEvidence is an auto generated Go binding around an Ethereum contract.
type IEvidence struct {
	IEvidenceCaller     // Read-only binding to the contract
	IEvidenceTransactor // Write-only binding to the contract
	IEvidenceFilterer   // Log filterer for contract events
}

// IEvidenceCaller is an auto generated read-only Go binding around an Ethereum contract.
type IEvidenceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IEvidenceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IEvidenceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IEvidenceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IEvidenceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IEvidenceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IEvidenceSession struct {
	Contract     *IEvidence        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IEvidenceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IEvidenceCallerSession struct {
	Contract *IEvidenceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// IEvidenceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IEvidenceTransactorSession struct {
	Contract     *IEvidenceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IEvidenceRaw is an auto generated low-level Go binding around an Ethereum contract.
type IEvidenceRaw struct {
	Contract *IEvidence // Generic contract binding to access the raw methods on
}

// IEvidenceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IEvidenceCallerRaw struct {
	Contract *IEvidenceCaller // Generic read-only contract binding to access the raw methods on
}

// IEvidenceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IEvidenceTransactorRaw struct {
	Contract *IEvidenceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIEvidence creates a new instance of IEvidence, bound to a specific deployed contract.
func NewIEvidence(address common.Address, backend bind.ContractBackend) (*IEvidence, error) {
	contract, err := bindIEvidence(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IEvidence{IEvidenceCaller: IEvidenceCaller{contract: contract}, IEvidenceTransactor: IEvidenceTransactor{contract: contract}, IEvidenceFilterer: IEvidenceFilterer{contract: contract}}, nil
}

// NewIEvidenceCaller creates a new read-only instance of IEvidence, bound to a specific deployed contract.
func NewIEvidenceCaller(address common.Address, caller bind.ContractCaller) (*IEvidenceCaller, error) {
	contract, err := bindIEvidence(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IEvidenceCaller{contract: contract}, nil
}

// NewIEvidenceTransactor creates a new write-only instance of IEvidence, bound to a specific deployed contract.
func NewIEvidenceTransactor(address common.Address, transactor bind.ContractTransactor) (*IEvidenceTransactor, error) {
	contract, err := bindIEvidence(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IEvidenceTransactor{contract: contract}, nil
}

// NewIEvidenceFilterer creates a new log filterer instance of IEvidence, bound to a specific deployed contract.
func NewIEvidenceFilterer(address common.Address, filterer bind.ContractFilterer) (*IEvidenceFilterer, error) {
	contract, err := bindIEvidence(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IEvidenceFilterer{contract: contract}, nil
}

// bindIEvidence binds a generic wrapper to an already deployed contract.
func bindIEvidence(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IEvidenceMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IEvidence *IEvidenceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IEvidence.Contract.IEvidenceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IEvidence *IEvidenceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IEvidence.Contract.IEvidenceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IEvidence *IEvidenceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IEvidence.Contract.IEvidenceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IEvidence *IEvidenceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IEvidence.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IEvidence *IEvidenceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IEvidence.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IEvidence *IEvidenceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IEvidence.Contract.contract.Transact(opts, method, params...)
}

// AllEvidence is a free data retrieval call binding the contract method 0x9e85e2c7.
//
// Solidity: function allEvidence((bytes,uint64,uint64,bool,bool) pagination) view returns((int64,int64,int64,address,string)[] evidence, (bytes,uint64) pageResponse)
func (_IEvidence *IEvidenceCaller) AllEvidence(opts *bind.CallOpts, pagination PageRequest) (struct {
	Evidence     []Equivocation
	PageResponse PageResponse
}, error) {
	var out []interface{}
	err := _IEvidence.contract.Call(opts, &out, "allEvidence", pagination)

	outstruct := new(struct {
		Evidence     []Equivocation
		PageResponse PageResponse
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Evidence = *abi.ConvertType(out[0], new([]Equivocation)).(*[]Equivocation)
	outstruct.PageResponse = *abi.ConvertType(out[1], new(PageResponse)).(*PageResponse)

	return *outstruct, err

}

// AllEvidence is a free data retrieval call binding the contract method 0x9e85e2c7.
//
// Solidity: function allEvidence((bytes,uint64,uint64,bool,bool) pagination) view returns((int64,int64,int64,address,string)[] evidence, (bytes,uint64) pageResponse)
func (_IEvidence *IEvidenceSession) AllEvidence(pagination PageRequest) (struct {
	Evidence     []Equivocation
	PageResponse PageResponse
}, error) {
	return _IEvidence.Contract.AllEvidence(&_IEvidence.CallOpts, pagination)
}

// AllEvidence is a free data retrieval call binding the contract method 0x9e85e2c7.
//
// Solidity: function allEvidence((bytes,uint64,uint64,bool,bool) pagination) view returns((int64,int64,int64,address,string)[] evidence, (bytes,uint64) pageResponse)
func (_IEvidence *IEvidenceCallerSession) AllEvidence(pagination PageRequest) (struct {
	Evidence     []Equivocation
	PageResponse PageResponse
}, error) {
	return _IEvidence.Contract.AllEvidence(&_IEvidence.CallOpts, pagination)
}

// Evidence is a free data retrieval call binding the contract method 0xbd283577.
//
// Solidity: function evidence(string evidenceHash) view returns((int64,int64,int64,address,string) evidence)
func (_IEvidence *IEvidenceCaller) Evidence(opts *bind.CallOpts, evidenceHash string) (Equivocation, error) {
	var out []interface{}
	err := _IEvidence.contract.Call(opts, &out, "evidence", evidenceHash)

	if err != nil {
		return *new(Equivocation), err
	}

	out0 := *abi.ConvertType(out[0], new(Equivocation)).(*Equivocation)

	return out0, err

}

// Evidence is a free data retrieval call binding the contract method 0xbd283577.
//
// Solidity: function evidence(string evidenceHash) view returns((int64,int64,int64,address,string) evidence)
func (_IEvidence *IEvidenceSession) Evidence(evidenceHash string) (Equivocation, error) {
	return _IEvidence.Contract.Evidence(&_IEvidence.CallOpts, evidenceHash)
}

// Evidence is a free data retrieval call binding the contract method 0xbd283577.
//
// Solidity: function evidence(string evidenceHash) view returns((int64,int64,int64,address,string) evidence)
func (_IEvidence *IEvidenceCallerSession) Evidence(evidenceHash string) (Equivocation, error) {
	return _IEvidence.Contract.Evidence(&_IEvidence.CallOpts, evidenceHash)
}

// SubmitEvidence is a paid mutator transaction binding the contract method 0x4330f90f.
//
// Solidity: function submitEvidence(int64 height, int64 time, int64 power, address consensusAddress) returns(string hash)
func (_IEvidence *IEvidenceTransactor) SubmitEvidence(opts *bind.TransactOpts, height int64, time int64, power int64, consensusAddress common.Address) (*types.Transaction, error) {
	return _IEvidence.contract.Transact(opts, "submitEvidence", height, time, power, consensusAddress)
}

// SubmitEvidence is a paid mutator transaction binding the contract method 0x4330f90f.
//
// Solidity: function submitEvidence(int64 height, int64 time, int64 power, address consensusAddress) returns(string hash)
func (_IEvidence *IEvidenceSession) SubmitEvidence(height int64, time int64, power int64, consensusAddress common.Address) (*types.Transaction, error) {
	return _IEvidence.Contract.SubmitEvidence(&_IEvidence.TransactOpts, height, time, power, consensusAddress)
}

// SubmitEvidence is a paid mutator transaction binding the contract method 0x4330f90f.
//
// Solidity: function submitEvidence(int64 height, int64 time, int64 power, address consensusAddress) returns(string hash)
func (_IEvidence *IEvidenceTransactorSession) SubmitEvidence(height int64, time int64, power int64, consensusAddress common.Address) (*types.Transaction, error) {
	return _IEvidence.Contract.SubmitEvidence(&_IEvidence.TransactOpts, height, time, power, consensusAddress)
}

// IEvidenceSubmitEvidenceIterator is returned from FilterSubmitEvidence and is used to iterate over the raw logs and unpacked data for SubmitEvidence events raised by the IEvidence contract.
type IEvidenceSubmitEvidenceIterator struct {
	Event *IEvidenceSubmitEvidence // Event containing the contract specifics and raw log

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
func (it *IEvidenceSubmitEvidenceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IEvidenceSubmitEvidence)
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
		it.Event = new(IEvidenceSubmitEvidence)
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
func (it *IEvidenceSubmitEvidenceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IEvidenceSubmitEvidenceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IEvidenceSubmitEvidence represents a SubmitEvidence event raised by the IEvidence contract.
type IEvidenceSubmitEvidence struct {
	Submitter    common.Address
	EvidenceHash string
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSubmitEvidence is a free log retrieval operation binding the contract event 0xd5358adc3d4486a89e80dc734cad1741beecd28933fa0a1691491f2b643804e1.
//
// Solidity: event SubmitEvidence(address indexed submitter, string evidenceHash)
func (_IEvidence *IEvidenceFilterer) FilterSubmitEvidence(opts *bind.FilterOpts, submitter []common.Address) (*IEvidenceSubmitEvidenceIterator, error) {

	var submitterRule []interface{}
	for _, submitterItem := range submitter {
		submitterRule = append(submitterRule, submitterItem)
	}

	logs, sub, err := _IEvidence.contract.FilterLogs(opts, "SubmitEvidence", submitterRule)
	if err != nil {
		return nil, err
	}
	return &IEvidenceSubmitEvidenceIterator{contract: _IEvidence.contract, event: "SubmitEvidence", logs: logs, sub: sub}, nil
}

// WatchSubmitEvidence is a free log subscription operation binding the contract event 0xd5358adc3d4486a89e80dc734cad1741beecd28933fa0a1691491f2b643804e1.
//
// Solidity: event SubmitEvidence(address indexed submitter, string evidenceHash)
func (_IEvidence *IEvidenceFilterer) WatchSubmitEvidence(opts *bind.WatchOpts, sink chan<- *IEvidenceSubmitEvidence, submitter []common.Address) (event.Subscription, error) {

	var submitterRule []interface{}
	for _, submitterItem := range submitter {
		submitterRule = append(submitterRule, submitterItem)
	}

	logs, sub, err := _IEvidence.contract.WatchLogs(opts, "SubmitEvidence", submitterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IEvidenceSubmitEvidence)
				if err := _IEvidence.contract.UnpackLog(event, "SubmitEvidence", log); err != nil {
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

// ParseSubmitEvidence is a log parse operation binding the contract event 0xd5358adc3d4486a89e80dc734cad1741beecd28933fa0a1691491f2b643804e1.
//
// Solidity: event SubmitEvidence(address indexed submitter, string evidenceHash)
func (_IEvidence *IEvidenceFilterer) ParseSubmitEvidence(log types.Log) (*IEvidenceSubmitEvidence, error) {
	event := new(IEvidenceSubmitEvidence)
	if err := _IEvidence.contract.UnpackLog(event, "SubmitEvidence", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
