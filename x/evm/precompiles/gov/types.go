package gov

import (
	"bytes"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/evmos/evmos/v12/types"
)

var (
	govAddress = common.HexToAddress(types.GovAddress)
	govABI     = types.MustABIJson(IGovMetaData.ABI)
)

func GetAddress() common.Address {
	return govAddress
}

func GetMethod(name string) (abi.Method, error) {
	method := govABI.Methods[name]
	if method.ID == nil {
		return abi.Method{}, fmt.Errorf("method %s is not exist", name)
	}
	return method, nil
}

func GetMethodByID(input []byte) (abi.Method, error) {
	if len(input) < 4 {
		return abi.Method{}, fmt.Errorf("input length %d is too short", len(input))
	}
	for _, method := range govABI.Methods {
		if bytes.Equal(input[:4], method.ID) {
			return method, nil
		}
	}
	return abi.Method{}, fmt.Errorf("method id %s is not exist", string(input[:4]))
}

func MustMethod(name string) abi.Method {
	method, err := GetMethod(name)
	if err != nil {
		panic(err)
	}
	return method
}

func GetEvent(name string) (abi.Event, error) {
	event := govABI.Events[name]
	if event.ID == (common.Hash{}) {
		return abi.Event{}, fmt.Errorf("event %s is not exist", name)
	}
	return event, nil
}

func MustEvent(name string) abi.Event {
	event, err := GetEvent(name)
	if err != nil {
		panic(err)
	}
	return event
}

type CoinJson = Coin

type LegacySubmitProposalArgs struct {
	Title          string   `abi:"title"`
	Description    string   `abi:"description"`
	InitialDeposit CoinJson `abi:"initialDeposit"`
}

// Validate LegacySubmitProposal the args
func (args *LegacySubmitProposalArgs) Validate() error {
	return nil
}

type VoteArgs struct {
	ProposalId uint64 `abi:"proposalId"`
	Option     int32  `abi:"option"`
	Metadata   string `abi:"metadata"`
}

// Validate Vote the args
func (args *VoteArgs) Validate() error {
	if args.ProposalId == 0 {
		return fmt.Errorf("proposal id must greater than 0")
	}

	return nil
}
