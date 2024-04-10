package slashing

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/core/vm"
)

const (
	SigningInfoGas = 30_000

	SigningInfoMethodName = "signingInfo"
)

func (c *Contract) SigningInfo(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	return nil, nil
}

func (c *Contract) SigningInfos(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	return nil, nil
}
