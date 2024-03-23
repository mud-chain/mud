package staking

import (
	"bytes"
	"encoding/base64"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/evmos/evmos/v12/x/evm/types"
)

const (
	DelegationGas = 30_000 // 98000
	ValidatorGas  = 30_000
	ValidatorsGas = 60_000

	DelegationMethodName = "delegation"
	ValidatorMethodName  = "validator"
	ValidatorsMethodName = "validators"
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

func (c *Contract) Validators(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := MustMethod(ValidatorsMethodName)

	// parse args
	var args ValidatorsArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}
	if bytes.Equal(args.PageRequest.Key, []byte{0}) {
		args.PageRequest.Key = nil
	}

	msg := &stakingtypes.QueryValidatorsRequest{
		Status: args.GetStatus(),
		Pagination: &query.PageRequest{
			Key:        args.PageRequest.Key,
			Offset:     args.PageRequest.Offset,
			Limit:      args.PageRequest.Limit,
			CountTotal: args.PageRequest.CountTotal,
			Reverse:    args.PageRequest.Reverse,
		},
	}

	querier := stakingkeeper.Querier{Keeper: c.stakingKeeper}

	res, err := querier.Validators(ctx, msg)
	if err != nil {
		return nil, err
	}

	var validators []Validator
	for _, validator := range res.Validators {
		validators = append(validators, OutputsValidator(validator))
	}

	var pageResponse PageResponse
	pageResponse.NextKey = res.Pagination.NextKey
	pageResponse.Total = res.Pagination.Total

	return method.Outputs.Pack(validators, pageResponse)
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

	return method.Outputs.Pack(OutputsValidator(res.Validator))
}

func OutputsValidator(validator stakingtypes.Validator) Validator {
	operatorValAddress, err := sdk.ValAddressFromBech32(validator.OperatorAddress)
	var operatorHexAddress common.Address
	if err == nil {
		operatorHexAddress = common.BytesToAddress(operatorValAddress)
	}

	return Validator{
		OperatorAddress: operatorHexAddress,
		ConsensusPubkey: FormatConsensusPubkey(validator.ConsensusPubkey),
		Jailed:          validator.Jailed,
		Status:          uint8(validator.Status),
		Tokens:          validator.Tokens.BigInt(),
		DelegatorShares: validator.DelegatorShares.BigInt(),
		Description:     Description(validator.Description),
		UnbondingHeight: validator.UnbondingHeight,
		UnbondingTime:   validator.UnbondingTime.Unix(),
		Commission: Commission{
			CommissionRates: CommissionRates{
				Rate:          validator.Commission.Rate.BigInt(),
				MaxRate:       validator.Commission.MaxRate.BigInt(),
				MaxChangeRate: validator.Commission.MaxChangeRate.BigInt(),
			},
			UpdateTime: validator.Commission.UpdateTime.Unix(),
		},
		MinSelfDelegation: validator.MinSelfDelegation.BigInt(),
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
