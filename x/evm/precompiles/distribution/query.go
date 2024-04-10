package distribution

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	distributiontypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/evmos/evmos/v12/utils"
	"github.com/evmos/evmos/v12/x/evm/types"
)

const (
	ValidatorDistributionInfoGas   = 30_000
	ValidatorOutstandingRewardsGas = 30_000
	ValidatorCommissionGas         = 30_000
	DelegationRewardsGas           = 30_000
	DelegationTotalRewardsGas      = 30_000

	ValidatorDistributionInfoMethodName   = "validatorDistributionInfo"
	ValidatorOutstandingRewardsMethodName = "validatorOutstandingRewards"
	ValidatorCommissionMethodName         = "validatorCommission"
	DelegationRewardsMethodName           = "delegationRewards"
	DelegationTotalRewardsMethodName      = "delegationTotalRewards"
)

// ValidatorDistributionInfo queries validator commision and self-delegation rewards for validator
func (c *Contract) ValidatorDistributionInfo(ctx sdk.Context, evm *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := MustMethod(ValidatorDistributionInfoMethodName)

	var args ValidatorAddressArgs
	err := types.ParseMethodArgs(method, &args, contract.Input[4:])
	if err != nil {
		return nil, err
	}

	msg := &distributiontypes.QueryValidatorDistributionInfoRequest{
		ValidatorAddress: args.GetValidator().String(),
	}

	res, err := c.distributionKeeper.ValidatorDistributionInfo(ctx, msg)
	if err != nil {
		return nil, err
	}

	operatorAddress := utils.AccAddressMustToHexAddress(res.OperatorAddress)

	var selfBondRewards []DecCoin
	for _, reward := range res.SelfBondRewards {
		selfBondRewards = append(selfBondRewards, DecCoin{
			Denom:     reward.Denom,
			Amount:    reward.Amount.BigInt(),
			Precision: uint8(sdk.Precision),
		})
	}

	var commission []DecCoin
	for _, reward := range res.Commission {
		commission = append(commission, DecCoin{
			Denom:     reward.Denom,
			Amount:    reward.Amount.BigInt(),
			Precision: uint8(sdk.Precision),
		})
	}

	return method.Outputs.Pack(operatorAddress, selfBondRewards, commission)
}

// ValidatorOutstandingRewards queries rewards of a validator address.
func (c *Contract) ValidatorOutstandingRewards(ctx sdk.Context, evm *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := MustMethod(ValidatorOutstandingRewardsMethodName)

	var args ValidatorAddressArgs
	err := types.ParseMethodArgs(method, &args, contract.Input[4:])
	if err != nil {
		return nil, err
	}

	msg := &distributiontypes.QueryValidatorOutstandingRewardsRequest{
		ValidatorAddress: args.GetValidator().String(),
	}

	res, err := c.distributionKeeper.ValidatorOutstandingRewards(ctx, msg)
	if err != nil {
		return nil, err
	}

	var rewards []DecCoin
	for _, reward := range res.Rewards.Rewards {
		rewards = append(rewards, DecCoin{
			Denom:     reward.Denom,
			Amount:    reward.Amount.BigInt(),
			Precision: uint8(sdk.Precision),
		})
	}

	return method.Outputs.Pack(rewards)
}

// ValidatorCommission queries accumulated commission for a validator.
func (c *Contract) ValidatorCommission(ctx sdk.Context, evm *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := MustMethod(ValidatorCommissionMethodName)

	var args ValidatorAddressArgs
	err := types.ParseMethodArgs(method, &args, contract.Input[4:])
	if err != nil {
		return nil, err
	}

	msg := &distributiontypes.QueryValidatorCommissionRequest{
		ValidatorAddress: args.GetValidator().String(),
	}

	res, err := c.distributionKeeper.ValidatorCommission(ctx, msg)
	if err != nil {
		return nil, err
	}

	var rewards []DecCoin
	for _, reward := range res.Commission.Commission {
		rewards = append(rewards, DecCoin{
			Denom:     reward.Denom,
			Amount:    reward.Amount.BigInt(),
			Precision: uint8(sdk.Precision),
		})
	}

	return method.Outputs.Pack(rewards)
}

// DelegationRewards queries the total rewards accrued by a delegation.
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

	res, err := c.distributionKeeper.DelegationRewards(ctx, msg)
	if err != nil {
		return nil, err
	}

	var rewards []DecCoin
	for _, reward := range res.Rewards {
		rewards = append(rewards, DecCoin{
			Denom:     reward.Denom,
			Amount:    reward.Amount.BigInt(),
			Precision: uint8(sdk.Precision),
		})
	}

	return method.Outputs.Pack(rewards)
}

// DelegationTotalRewards queries the total rewards accrued by a each validator.
func (c *Contract) DelegationTotalRewards(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := MustMethod(DelegationTotalRewardsMethodName)

	var args DelegatorAddressArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}
	msg := &distributiontypes.QueryDelegationTotalRewardsRequest{
		DelegatorAddress: args.GetDelegator().String(),
	}

	res, err := c.distributionKeeper.DelegationTotalRewards(ctx, msg)
	if err != nil {
		return nil, err
	}

	var delegationDelegatorReward []DelegationDelegatorReward
	for _, reward := range res.Rewards {
		var rewards []DecCoin
		for _, r := range reward.Reward {
			rewards = append(rewards, DecCoin{
				Denom:     r.Denom,
				Amount:    r.Amount.BigInt(),
				Precision: uint8(sdk.Precision),
			})
		}
		delegationDelegatorReward = append(delegationDelegatorReward, DelegationDelegatorReward{
			ValidatorAddress: utils.ValAddressMustToHexAddress(reward.ValidatorAddress),
			Rewards:          rewards,
		})
	}

	var total []DecCoin
	for _, reward := range res.Total {
		total = append(total, DecCoin{
			Denom:     reward.Denom,
			Amount:    reward.Amount.BigInt(),
			Precision: uint8(sdk.Precision),
		})
	}

	return method.Outputs.Pack(delegationDelegatorReward, total)
}
