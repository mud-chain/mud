package staking

import (
	"encoding/base64"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/evmos/evmos/v12/x/evm/types"
)

const (
	DelegationGas = 30_000 // 98000
	ValidatorGas  = 30_000

	DelegationMethodName = "delegation"
	ValidatorMethodName  = "validator"
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

	res, err := querier.Delegation(ctx, msg)
	if err != nil {
		return nil, err
	}

	shares := res.DelegationResponse.Delegation.GetShares().TruncateInt().BigInt()
	balance := Coin{
		Denom:  res.DelegationResponse.Balance.Denom,
		Amount: res.DelegationResponse.Balance.Amount.BigInt(),
	}

	return method.Outputs.Pack(shares, balance)
}

func (c *Contract) Validator(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := MustMethod(ValidatorMethodName)

	// parse args
	var args ValidatorArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}
	msg := &stakingtypes.QueryValidatorRequest{
		ValidatorAddr: args.GetValidator().String(),
	}

	querier := stakingkeeper.Querier{Keeper: c.stakingKeeper}

	res, err := querier.Validator(ctx, msg)
	if err != nil {
		return nil, err
	}

	return method.Outputs.Pack(OutputsValidator(res))
}

func OutputsValidator(res *stakingtypes.QueryValidatorResponse) Validator {
	return Validator{
		OperatorAddress: res.Validator.OperatorAddress,
		ConsensusPubkey: FormatConsensusPubkey(res.Validator.ConsensusPubkey),
		Jailed:          res.Validator.Jailed,
		Status:          uint8(res.Validator.Status),
		Tokens:          res.Validator.Tokens.BigInt(),
		DelegatorShares: res.Validator.DelegatorShares.BigInt(),
		Description:     Description(res.Validator.Description),
		UnbondingHeight: res.Validator.UnbondingHeight,
		UnbondingTime:   res.Validator.UnbondingTime.Unix(),
		Commission: Commission{
			CommissionRates: CommissionRates{
				Rate:          res.Validator.Commission.Rate.BigInt(),
				MaxRate:       res.Validator.Commission.MaxRate.BigInt(),
				MaxChangeRate: res.Validator.Commission.MaxChangeRate.BigInt(),
			},
			UpdateTime: res.Validator.Commission.UpdateTime.Unix(),
		},
		MinSelfDelegation: res.Validator.MinSelfDelegation.BigInt(),
	}
}

// FormatConsensusPubkey format ConsensusPubkey into a base64 string
func FormatConsensusPubkey(consensusPubkey *codectypes.Any) string {
	ed25519pk, ok := consensusPubkey.GetCachedValue().(cryptotypes.PubKey)
	if ok {
		return base64.StdEncoding.EncodeToString(ed25519pk.Bytes())
	}
	return consensusPubkey.String()
}
