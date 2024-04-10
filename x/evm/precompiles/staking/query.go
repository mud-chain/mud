package staking

import (
	"bytes"
	"encoding/base64"
	"github.com/evmos/evmos/v12/utils"

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
	DelegationGas                    = 30_000
	UnbondingDelegationGas           = 30_000
	ValidatorGas                     = 30_000
	ValidatorsGas                    = 60_000
	ValidatorDelegationsGas          = 90_000
	ValidatorUnbondingDelegationsGas = 90_000
	DelegatorDelegationsGas          = 90_000
	DelegatorUnbondingDelegationsGas = 90_000

	DelegationMethodName              = "delegation"
	UnbondingDelegationMethodName     = "unbondingDelegation"
	ValidatorMethodName               = "validator"
	ValidatorsMethodName              = "validators"
	ValidatorDelegationsName          = "validatorDelegations"
	ValidatorUnbondingDelegationsName = "validatorUnbondingDelegations"
	DelegatorDelegationsName          = "delegatorDelegations"
	DelegatorUnbondingDelegationsName = "delegatorUnbondingDelegations"
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

	return method.Outputs.Pack(OutputsDelegation(*res.DelegationResponse))
}

func (c *Contract) UnbondingDelegation(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := MustMethod(UnbondingDelegationMethodName)

	// parse args
	var args UnbondingDelegationArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}
	msg := &stakingtypes.QueryUnbondingDelegationRequest{
		DelegatorAddr: args.GetDelegator().String(),
		ValidatorAddr: args.GetValidator().String(),
	}

	querier := stakingkeeper.Querier{Keeper: c.stakingKeeper}

	res, err := querier.UnbondingDelegation(ctx, msg)
	if err != nil {
		return nil, err
	}

	return method.Outputs.Pack(OutputsUnbondingDelegation(res.Unbond))
}

func (c *Contract) Validators(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := MustMethod(ValidatorsMethodName)

	// parse args
	var args ValidatorsArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}
	if bytes.Equal(args.Pagination.Key, []byte{0}) {
		args.Pagination.Key = nil
	}

	msg := &stakingtypes.QueryValidatorsRequest{
		Status: args.GetStatus(),
		Pagination: &query.PageRequest{
			Key:        args.Pagination.Key,
			Offset:     args.Pagination.Offset,
			Limit:      args.Pagination.Limit,
			CountTotal: args.Pagination.CountTotal,
			Reverse:    args.Pagination.Reverse,
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

// ValidatorDelegations queries delegate info for given validator.
func (c *Contract) ValidatorDelegations(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := MustMethod(ValidatorDelegationsName)

	// parse args
	var args ValidatorDelegationsArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}
	if bytes.Equal(args.Pagination.Key, []byte{0}) {
		args.Pagination.Key = nil
	}
	msg := &stakingtypes.QueryValidatorDelegationsRequest{
		ValidatorAddr: args.GetValidator().String(),
		Pagination: &query.PageRequest{
			Key:        args.Pagination.Key,
			Offset:     args.Pagination.Offset,
			Limit:      args.Pagination.Limit,
			CountTotal: args.Pagination.CountTotal,
			Reverse:    args.Pagination.Reverse,
		},
	}

	querier := stakingkeeper.Querier{Keeper: c.stakingKeeper}

	res, err := querier.ValidatorDelegations(ctx, msg)
	if err != nil {
		return nil, err
	}

	var delegations []DelegationResponse
	for _, delegation := range res.DelegationResponses {
		delegations = append(delegations, OutputsDelegation(delegation))
	}

	var pageResponse PageResponse
	pageResponse.NextKey = res.Pagination.NextKey
	pageResponse.Total = res.Pagination.Total

	return method.Outputs.Pack(delegations, pageResponse)
}

// ValidatorUnbondingDelegations queries unbonding delegations of a validator.
func (c *Contract) ValidatorUnbondingDelegations(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := MustMethod(ValidatorUnbondingDelegationsName)

	// parse args
	var args ValidatorDelegationsArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}
	if bytes.Equal(args.Pagination.Key, []byte{0}) {
		args.Pagination.Key = nil
	}
	msg := &stakingtypes.QueryValidatorUnbondingDelegationsRequest{
		ValidatorAddr: args.GetValidator().String(),
		Pagination: &query.PageRequest{
			Key:        args.Pagination.Key,
			Offset:     args.Pagination.Offset,
			Limit:      args.Pagination.Limit,
			CountTotal: args.Pagination.CountTotal,
			Reverse:    args.Pagination.Reverse,
		},
	}

	querier := stakingkeeper.Querier{Keeper: c.stakingKeeper}

	res, err := querier.ValidatorUnbondingDelegations(ctx, msg)
	if err != nil {
		return nil, err
	}

	var unbondingDelegations []UnbondingDelegation
	for _, unbondingDelegation := range res.UnbondingResponses {
		unbondingDelegations = append(unbondingDelegations, OutputsUnbondingDelegation(unbondingDelegation))
	}

	var pageResponse PageResponse
	pageResponse.NextKey = res.Pagination.NextKey
	pageResponse.Total = res.Pagination.Total

	return method.Outputs.Pack(unbondingDelegations, pageResponse)
}

// DelegatorDelegations queries all delegations of a given delegator address.
func (c *Contract) DelegatorDelegations(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := MustMethod(DelegatorDelegationsName)

	// parse args
	var args DelegatorDelegationsArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}
	if bytes.Equal(args.Pagination.Key, []byte{0}) {
		args.Pagination.Key = nil
	}
	msg := &stakingtypes.QueryDelegatorDelegationsRequest{
		DelegatorAddr: args.GetDelegator().String(),
		Pagination: &query.PageRequest{
			Key:        args.Pagination.Key,
			Offset:     args.Pagination.Offset,
			Limit:      args.Pagination.Limit,
			CountTotal: args.Pagination.CountTotal,
			Reverse:    args.Pagination.Reverse,
		},
	}

	querier := stakingkeeper.Querier{Keeper: c.stakingKeeper}

	res, err := querier.DelegatorDelegations(ctx, msg)
	if err != nil {
		return nil, err
	}

	var delegations []DelegationResponse
	for _, delegation := range res.DelegationResponses {
		delegations = append(delegations, OutputsDelegation(delegation))
	}

	var pageResponse PageResponse
	pageResponse.NextKey = res.Pagination.NextKey
	pageResponse.Total = res.Pagination.Total

	return method.Outputs.Pack(delegations, pageResponse)
}

// DelegatorUnbondingDelegations queries all unbonding delegations of a given
// delegator address.
func (c *Contract) DelegatorUnbondingDelegations(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := MustMethod(DelegatorUnbondingDelegationsName)

	// parse args
	var args DelegatorDelegationsArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}
	if bytes.Equal(args.Pagination.Key, []byte{0}) {
		args.Pagination.Key = nil
	}
	msg := &stakingtypes.QueryDelegatorUnbondingDelegationsRequest{
		DelegatorAddr: args.GetDelegator().String(),
		Pagination: &query.PageRequest{
			Key:        args.Pagination.Key,
			Offset:     args.Pagination.Offset,
			Limit:      args.Pagination.Limit,
			CountTotal: args.Pagination.CountTotal,
			Reverse:    args.Pagination.Reverse,
		},
	}

	querier := stakingkeeper.Querier{Keeper: c.stakingKeeper}

	res, err := querier.DelegatorUnbondingDelegations(ctx, msg)
	if err != nil {
		return nil, err
	}

	var unbondingDelegations []UnbondingDelegation
	for _, unbondingDelegation := range res.UnbondingResponses {
		unbondingDelegations = append(unbondingDelegations, OutputsUnbondingDelegation(unbondingDelegation))
	}

	var pageResponse PageResponse
	pageResponse.NextKey = res.Pagination.NextKey
	pageResponse.Total = res.Pagination.Total

	return method.Outputs.Pack(unbondingDelegations, pageResponse)
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

func OutputsDelegation(delegationResponse stakingtypes.DelegationResponse) DelegationResponse {
	deletation := delegationResponse.Delegation
	balance := delegationResponse.Balance

	return DelegationResponse{
		Delegation: Delegation{
			DelegatorAddress: utils.AccAddressMustToHexAddress(deletation.DelegatorAddress),
			ValidatorAddress: utils.ValAddressMustToHexAddress(deletation.ValidatorAddress),
			Shares: Dec{
				Amount:    deletation.Shares.BigInt(),
				Precision: sdk.Precision,
			},
		},
		Balance: Coin{
			Denom:  balance.Denom,
			Amount: balance.Amount.BigInt(),
		},
	}
}

func OutputsUnbondingDelegation(unbondingDelegation stakingtypes.UnbondingDelegation) UnbondingDelegation {
	var entries []UnbondingDelegationEntry
	for _, entry := range unbondingDelegation.Entries {
		entries = append(entries, UnbondingDelegationEntry{
			CreationHeight: entry.CreationHeight,
			CompletionTime: entry.CompletionTime.Unix(),
			InitialBalance: entry.InitialBalance.BigInt(),
			Balance:        entry.Balance.BigInt(),
		})
	}

	return UnbondingDelegation{
		DelegatorAddress: utils.AccAddressMustToHexAddress(unbondingDelegation.DelegatorAddress),
		ValidatorAddress: utils.ValAddressMustToHexAddress(unbondingDelegation.ValidatorAddress),
		Entries:          entries,
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
