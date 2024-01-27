package staking

import (
	sdkmath "cosmossdk.io/math"
	"errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/evmos/evmos/v12/x/evm/types"
)

func (c *Contract) Delegate(ctx sdk.Context, evm *vm.EVM, contract *vm.Contract, readonly bool) ([]byte, error) {
	if readonly {
		return nil, errors.New("delegate method not readonly")
	}
	// parse args
	var args DelegateArgs
	err := types.ParseMethodArgs(DelegateMethod, &args, contract.Input[4:])
	if err != nil {
		return nil, err
	}

	msg := &stakingtypes.MsgDelegate{
		DelegatorAddress: sdk.AccAddress(contract.Caller().Bytes()).String(),
		ValidatorAddress: args.GetValidator().String(),
		Amount: sdk.Coin{
			Denom:  c.stakingKeeper.GetParams(ctx).BondDenom,
			Amount: sdkmath.NewIntFromBigInt(args.Amount),
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
		DelegateEvent,
		[]common.Hash{common.BytesToHash(contract.Caller().Bytes()), common.BytesToHash(args.GetValidator().Bytes())},
		args.Amount,
	); err != nil {
		return nil, err
	}

	return DelegateMethod.Outputs.Pack(true)
}
