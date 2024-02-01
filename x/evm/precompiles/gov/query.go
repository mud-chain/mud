package gov

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/core/vm"
)

const (
	ProposalGas = 30_000

	ProposalMethodName = "proposal"
)

func (c *Contract) Proposal(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := MustMethod(ProposalMethodName)

	//var args DelegationRewardsArgs
	//if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
	//	return nil, err
	//}
	//msg := &govtypes.QueryDelegationRewardsRequest{
	//	DelegatorAddress: args.GetDelegator().String(),
	//	ValidatorAddress: args.GetValidator().String(),
	//}
	//
	//resp, err := c.govKeeper.DelegationRewards(ctx, msg)
	//if err != nil {
	//	return nil, err
	//}
	//
	//var rewards []DecCoin
	//for _, reward := range resp.Rewards {
	//	rewards = append(rewards, DecCoin{
	//		Denom:     reward.Denom,
	//		Amount:    reward.Amount.BigInt(),
	//		Precision: uint8(sdk.Precision),
	//	})
	//}

	return method.Outputs.Pack(true)
}
