package bank

import (
	"bytes"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/evmos/evmos/v12/x/evm/types"
)

const (
	BalanceGas     = 30_000
	AllBalancesGas = 50_000
	TotalSupplyGas = 50_000

	BalanceMethodName     = "balance"
	AllBalancesMethodName = "allBalances"
	TotalSupplyMethodName = "totalSupply"
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

func (c *Contract) AllBalances(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := MustMethod(AllBalancesMethodName)

	var args AllBalancesArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}
	if bytes.Equal(args.PageRequest.Key, []byte{0}) {
		args.PageRequest.Key = nil
	}

	msg := &banktypes.QueryAllBalancesRequest{
		Address: sdk.AccAddress(args.AccountAddress.Bytes()).String(),
		Pagination: &query.PageRequest{
			Key:        args.PageRequest.Key,
			Offset:     args.PageRequest.Offset,
			Limit:      args.PageRequest.Limit,
			CountTotal: args.PageRequest.CountTotal,
			Reverse:    args.PageRequest.Reverse,
		},
	}

	res, err := c.bankKeeper.AllBalances(ctx, msg)
	if err != nil {
		return nil, err
	}

	var balances []Coin
	for _, balance := range res.Balances {
		balances = append(balances, Coin{
			Denom:  balance.Denom,
			Amount: balance.Amount.BigInt(),
		})
	}

	var pageResponse PageResponse
	pageResponse.NextKey = res.Pagination.NextKey
	pageResponse.Total = res.Pagination.Total

	return method.Outputs.Pack(balances, pageResponse)
}

// TotalSupply queries the total supply of all coins.
func (c *Contract) TotalSupply(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := MustMethod(TotalSupplyMethodName)

	msg := &banktypes.QueryTotalSupplyRequest{}

	res, err := c.bankKeeper.TotalSupply(ctx, msg)
	if err != nil {
		return nil, err
	}

	var balances []Coin
	for _, balance := range res.Supply {
		balances = append(balances, Coin{
			Denom:  balance.Denom,
			Amount: balance.Amount.BigInt(),
		})
	}

	return method.Outputs.Pack(balances)
}
