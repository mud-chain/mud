package bank

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/evmos/evmos/v12/x/evm/types"
)

const (
	BalanceGas = 30_000

	BalanceMethodName = "balance"
)

func (c *Contract) Balance(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := MustMethod(BalanceMethodName)

	var args BalanceArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}
	msg := &banktypes.QueryBalanceRequest{
		Address: sdk.AccAddress(args.AccountAddress.Bytes()).String(),
		Denom:   args.Denom,
	}

	res, err := c.bankKeeper.Balance(ctx, msg)
	if err != nil {
		return nil, err
	}

	return method.Outputs.Pack(Coin{
		Denom:  res.Balance.Denom,
		Amount: res.Balance.Amount.BigInt(),
	})
}
