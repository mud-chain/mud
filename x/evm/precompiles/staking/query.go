package staking

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/evmos/evmos/v12/x/evm/types"
)

const (
	DelegationGas = 30_000 // 98000

	DelegationMethodName = "delegation"
)

func (c *Contract) Delegation(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := MustMethod(DelegationMethodName)

	// parse args
	var args DelegationArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}
	msg := &stakingtypes.QueryDelegationRequest{
		DelegatorAddr: args.GetDelegator().String(),
		ValidatorAddr: args.GetValidator().String(),
	}

	querier := stakingkeeper.Querier{Keeper: c.stakingKeeper}

	resp, err := querier.Delegation(ctx, msg)
	if err != nil {
		return nil, err
	}

	shares := resp.DelegationResponse.Delegation.GetShares().TruncateInt().BigInt()
	balance := Coin{
		Denom:  resp.DelegationResponse.Balance.Denom,
		Amount: resp.DelegationResponse.Balance.Amount.BigInt(),
	}

	return method.Outputs.Pack(shares, balance)
}
