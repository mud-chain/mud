package evidence

import (
	"bytes"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/evmos/evmos/v12/types"
)

var (
	evidenceAddress = common.HexToAddress(types.EvidenceAddress)
	evidenceABI     = types.MustABIJson(IEvidenceMetaData.ABI)
)

func GetAddress() common.Address {
	return evidenceAddress
}

func GetMethod(name string) (abi.Method, error) {
	method := evidenceABI.Methods[name]
	if method.ID == nil {
		return abi.Method{}, fmt.Errorf("method %s is not exist", name)
	}
	return method, nil
}

func GetMethodByID(input []byte) (abi.Method, error) {
	if len(input) < 4 {
		return abi.Method{}, fmt.Errorf("input length %d is too short", len(input))
	}
	for _, method := range evidenceABI.Methods {
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
	event := evidenceABI.Events[name]
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

type PageRequestJson = PageRequest

type SubmitEvidenceArgs struct {
	Height           int64          `abi:"height"`
	Time             int64          `abi:"time"`
	Power            int64          `abi:"power"`
	ConsensusAddress common.Address `abi:"consensusAddress"`
}

// Validate submit evidence args
func (args *SubmitEvidenceArgs) Validate() error {
	if args.Height < 0 {
		return fmt.Errorf("invalid height: %v", args.Height)
	}

	if args.ConsensusAddress == (common.Address{}) {
		return fmt.Errorf("invalid consensus address: %s", args.ConsensusAddress)
	}

	return nil
}

// GetConsAddress returns the consensus address, caller must ensure the consensus address is valid
func (args *SubmitEvidenceArgs) GetConsAddress() sdk.ConsAddress {
	consAddress := sdk.ConsAddress(args.ConsensusAddress.Bytes())
	return consAddress
}

type EvidenceArgs struct {
	EvidenceHash string `abi:"evidenceHash"`
}

// Validate evidence args
func (args *EvidenceArgs) Validate() error {
	if args.EvidenceHash == "" {
		return fmt.Errorf("invalid evidence hash: %s", args.EvidenceHash)
	}

	return nil
}

type AllEvidenceArgs struct {
	Pagination PageRequestJson `abi:"pagination"`
}

// Validate all evidence args
func (args *AllEvidenceArgs) Validate() error {
	return nil
}
