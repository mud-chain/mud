package distribution

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	distributiontypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/evmos/evmos/v12/x/evm/types"
)

const (
	DelegationRewardsGas = 30_000

	DelegationRewardsMethodName = "delegationRewards"
)

func (c *Contract) DelegationRewards(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := MustMethod(DelegationRewardsMethodName)

	var args DelegationRewardsArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}
	msg := &distributiontypes.QueryDelegationRewardsRequest{
		DelegatorAddress: args.GetDelegator().String(),
		ValidatorAddress: args.GetValidator().String(),
	}

	resp, err := c.distributionKeeper.DelegationRewards(ctx, msg)
	if err != nil {
		return nil, err
	}

	var rewards []DecCoin
	for _, reward := range resp.Rewards {
		rewards = append(rewards, DecCoin{
			Denom:     reward.Denom,
			Amount:    reward.Amount.BigInt(),
			Precision: uint8(sdk.Precision),
		})
	}

	return method.Outputs.Pack(rewards)
}
