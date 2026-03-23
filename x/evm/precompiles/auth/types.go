package auth

import (
	"bytes"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/evmos/evmos/v12/types"
)

var (
	authAddress = common.HexToAddress(types.AuthAddress)
	authABI     = types.MustABIJson(IAuthMetaData.ABI)
)

func GetAddress() common.Address {
	return authAddress
}

func GetMethod(name string) (abi.Method, error) {
	method := authABI.Methods[name]
	if method.ID == nil {
		return abi.Method{}, fmt.Errorf("method %s is not exist", name)
	}
	return method, nil
}

func GetMethodByID(input []byte) (abi.Method, error) {
	if len(input) < 4 {
		return abi.Method{}, fmt.Errorf("input length %d is too short", len(input))
	}
	for _, method := range authABI.Methods {
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
	event := authABI.Events[name]
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

type AccountArgs struct {
	Addr common.Address `abi:"addr"`
}

// Validate Account args
func (args *AccountArgs) Validate() error {
	return nil
}

type PageRequestJson = PageRequest

type AccountsArgs struct {
	PageRequest PageRequestJson `abi:"pageRequest"`
}

// Validate Accounts args
func (args *AccountsArgs) Validate() error {
	return nil
}

type ModuleAccountByNameArgs struct {
	Name string `abi:"name"`
}

func (args *ModuleAccountByNameArgs) Validate() error {
	if args.Name == "" {
		return fmt.Errorf("name is empty")
	}
	return nil
}

type AccountAddressByIDArgs struct {
	Id int64 `abi:"id"`
}

func (args *AccountAddressByIDArgs) Validate() error {
	if args.Id < 0 {
		return fmt.Errorf("id must be great than 0")
	}
	return nil
}
