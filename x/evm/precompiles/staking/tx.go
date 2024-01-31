package staking

import (
	"errors"
	"math/big"

	"cosmossdk.io/math"

	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/evmos/evmos/v12/x/evm/types"
)

const (
	CreateValidatorGas           = 60_000
	DelegateGas                  = 40_000 // 98000 - 160000 // 165000
	UndelegateGas                = 45_000 // 94000 - 163000 // 172000
	RedelegateGas                = 60_000 // undelegate_gas+delegate_gas+withdraw_gas*2
	CancelUnbondingDelegationGas = 30_000 // 98000

	CreateValidatorMethodName           = "createValidator"
	DelegateMethodName                  = "delegate"
	UndelegateMethodName                = "undelegate"
	RedelegateMethodName                = "redelegate"
	CancelUnbondingDelegationMethodName = "cancelUnbondingDelegation"

	CreateValidatorEventName           = "CreateValidator"
	DelegateEventName                  = "Delegate"
	UndelegateEventName                = "Undelegate"
	RedelegateEventName                = "Redelegate"
	CancelUnbondingDelegationEventName = "CancelUnbondingDelegation"
)

func (c *Contract) CreateValidator(ctx sdk.Context, evm *vm.EVM, contract *vm.Contract, readonly bool) ([]byte, error) {
	if readonly {
		return nil, errors.New("createValidator method not readonly")
	}

	method := MustMethod(CreateValidatorMethodName)

	// parse args
	var args CreateValidatorArgs
	err := types.ParseMethodArgs(method, &args, contract.Input[4:])
	if err != nil {
		return nil, err
	}

	msg := &stakingtypes.MsgCreateValidator{
		Description: stakingtypes.Description(args.Description),
		Commission: stakingtypes.CommissionRates{
			Rate:          sdk.NewDecFromBigIntWithPrec(args.Commission.Rate, sdk.Precision),
			MaxRate:       sdk.NewDecFromBigIntWithPrec(args.Commission.MaxRate, sdk.Precision),
			MaxChangeRate: sdk.NewDecFromBigIntWithPrec(args.Commission.MaxChangeRate, sdk.Precision),
		},
		MinSelfDelegation: math.NewIntFromBigInt(args.MinSelfDelegation),
		DelegatorAddress:  sdk.AccAddress(contract.Caller().Bytes()).String(),
		ValidatorAddress:  sdk.ValAddress(contract.Caller().Bytes()).String(),
		Pubkey:            args.GetPubkey(),
		Value: sdk.Coin{
			Denom:  c.stakingKeeper.GetParams(ctx).BondDenom,
			Amount: math.NewIntFromBigInt(args.Value),
		},
	}

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	server := stakingkeeper.NewMsgServerImpl(c.stakingKeeper)

	_, err = server.CreateValidator(ctx, msg)
	if err != nil {
		return nil, err
	}

	// add delegate log
	if err := c.AddLog(
		evm,
		MustEvent(CreateValidatorEventName),
		[]common.Hash{common.BytesToHash(contract.Caller().Bytes())},
		args.Value,
	); err != nil {
		return nil, err
	}

	return method.Outputs.Pack(true)
}

func (c *Contract) Delegate(ctx sdk.Context, evm *vm.EVM, contract *vm.Contract, readonly bool) ([]byte, error) {
	if readonly {
		return nil, errors.New("delegate method not readonly")
	}

	method := MustMethod(DelegateMethodName)

	// parse args
	var args DelegateArgs
	err := types.ParseMethodArgs(method, &args, contract.Input[4:])
	if err != nil {
		return nil, err
	}

	msg := &stakingtypes.MsgDelegate{
		DelegatorAddress: sdk.AccAddress(contract.Caller().Bytes()).String(),
		ValidatorAddress: args.GetValidator().String(),
		Amount: sdk.Coin{
			Denom:  c.stakingKeeper.GetParams(ctx).BondDenom,
			Amount: math.NewIntFromBigInt(args.Amount),
		},
	}

	server := stakingkeeper.NewMsgServerImpl(c.stakingKeeper)

	_, err = server.Delegate(ctx, msg)
	if err != nil {
		return nil, err
	}

	// add delegate log
	if err := c.AddLog(
		evm,
		MustEvent(DelegateEventName),
		[]common.Hash{common.BytesToHash(contract.Caller().Bytes()), common.BytesToHash(args.GetValidator().Bytes())},
		args.Amount,
	); err != nil {
		return nil, err
	}

	return method.Outputs.Pack(true)
}

func (c *Contract) Undelegate(ctx sdk.Context, evm *vm.EVM, contract *vm.Contract, readonly bool) ([]byte, error) {
	if readonly {
		return nil, errors.New("undelegate method not readonly")
	}

	method := MustMethod(UndelegateMethodName)
	// parse args
	var args UndelegateArgs
	err := types.ParseMethodArgs(method, &args, contract.Input[4:])
	if err != nil {
		return nil, err
	}

	msg := &stakingtypes.MsgUndelegate{
		DelegatorAddress: sdk.AccAddress(contract.Caller().Bytes()).String(),
		ValidatorAddress: args.GetValidator().String(),
		Amount: sdk.Coin{
			Denom:  c.stakingKeeper.GetParams(ctx).BondDenom,
			Amount: math.NewIntFromBigInt(args.Amount),
		},
	}

	server := stakingkeeper.NewMsgServerImpl(c.stakingKeeper)

	res, err := server.Undelegate(ctx, msg)
	if err != nil {
		return nil, err
	}
	completionTime := big.NewInt(res.CompletionTime.Unix())

	// add undelegate log
	if err := c.AddLog(
		evm,
		MustEvent(UndelegateEventName),
		[]common.Hash{common.BytesToHash(contract.Caller().Bytes()), common.BytesToHash(args.GetValidator().Bytes())},
		args.Amount,
		completionTime,
	); err != nil {
		return nil, err
	}

	return method.Outputs.Pack(completionTime)
}

func (c *Contract) Redelegatge(ctx sdk.Context, evm *vm.EVM, contract *vm.Contract, readonly bool) ([]byte, error) {
	if readonly {
		return nil, errors.New("redelegate method not readonly")
	}
	method := MustMethod(RedelegateMethodName)

	// parse args
	var args RedelegateArgs
	err := types.ParseMethodArgs(method, &args, contract.Input[4:])
	if err != nil {
		return nil, err
	}

	msg := &stakingtypes.MsgBeginRedelegate{
		DelegatorAddress:    sdk.AccAddress(contract.Caller().Bytes()).String(),
		ValidatorSrcAddress: args.GetSrcValidator().String(),
		ValidatorDstAddress: args.GetDstValidator().String(),
		Amount: sdk.Coin{
			Denom:  c.stakingKeeper.GetParams(ctx).BondDenom,
			Amount: math.NewIntFromBigInt(args.Amount),
		},
	}

	server := stakingkeeper.NewMsgServerImpl(c.stakingKeeper)

	res, err := server.BeginRedelegate(ctx, msg)
	if err != nil {
		return nil, err
	}
	completionTime := big.NewInt(res.CompletionTime.Unix())

	// add redelegate log
	if err := c.AddLog(
		evm,
		MustEvent(RedelegateEventName),
		[]common.Hash{common.BytesToHash(contract.Caller().Bytes()), common.BytesToHash(args.GetSrcValidator().Bytes()), common.BytesToHash(args.GetDstValidator().Bytes())},
		args.Amount,
		completionTime,
	); err != nil {
		return nil, err
	}

	return method.Outputs.Pack(completionTime)
}

func (c *Contract) CancelUnbondingDelegation(ctx sdk.Context, evm *vm.EVM, contract *vm.Contract, readonly bool) ([]byte, error) {
	if readonly {
		return nil, errors.New("cancelUnbondingDelegation method not readonly")
	}

	method := MustMethod(CancelUnbondingDelegationMethodName)

	// parse args
	var args CancelUnbondingDelegationArgs
	err := types.ParseMethodArgs(method, &args, contract.Input[4:])
	if err != nil {
		return nil, err
	}

	msg := &stakingtypes.MsgCancelUnbondingDelegation{
		DelegatorAddress: sdk.AccAddress(contract.Caller().Bytes()).String(),
		ValidatorAddress: args.GetValidator().String(),
		Amount: sdk.Coin{
			Denom:  c.stakingKeeper.GetParams(ctx).BondDenom,
			Amount: math.NewIntFromBigInt(args.Amount),
		},
		CreationHeight: args.GetCreationHeight(),
	}

	server := stakingkeeper.NewMsgServerImpl(c.stakingKeeper)

	_, err = server.CancelUnbondingDelegation(ctx, msg)
	if err != nil {
		return nil, err
	}

	// add CancelUnbondingDelegationMethod log
	if err := c.AddLog(
		evm,
		MustEvent(CancelUnbondingDelegationEventName),
		[]common.Hash{common.BytesToHash(contract.Caller().Bytes()), common.BytesToHash(args.GetValidator().Bytes())},
		args.Amount,
		args.CreationHeight,
	); err != nil {
		return nil, err
	}

	return method.Outputs.Pack(true)
}
