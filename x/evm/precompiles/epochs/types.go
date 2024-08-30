package epochs

import (
	"bytes"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/evmos/evmos/v12/types"
)

var (
	epochsAddress = common.HexToAddress(types.EpochsAddress)
	evidenceABI   = types.MustABIJson(IEpochsMetaData.ABI)
)

func GetAddress() common.Address {
	return epochsAddress
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

type PageRequestJson = PageRequest

type EpochInfosArgs struct {
	Pagination PageRequestJson `abi:"pagination"`
}

// Validate epoch infos args
func (args *EpochInfosArgs) Validate() error {
	return nil
}

type CurrentEpochArgs struct {
	Identifier string `abi:"identifier"`
}

// Validate current epoch args
func (args *CurrentEpochArgs) Validate() error {
	if args.Identifier == "" {
		return fmt.Errorf("invalid identifier: %s", args.Identifier)
	}

	return nil
}
