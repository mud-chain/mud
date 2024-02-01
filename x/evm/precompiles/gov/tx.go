package gov

import (
	"errors"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	govkeeper "github.com/cosmos/cosmos-sdk/x/gov/keeper"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	v1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	"github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/evmos/evmos/v12/x/evm/types"
)

const (
	LegacySubmitProposalGas = 60_000
	VoteGas                 = 60_000

	LegacySubmitProposalMethodName = "legacySubmitProposal"
	VoteMethodName                 = "vote"

	LegacySubmitProposalEventName = "LegacySubmitProposal"
	VoteEventName                 = "Vote"
)

func (c *Contract) LegacySubmitProposal(ctx sdk.Context, evm *vm.EVM, contract *vm.Contract, readonly bool) ([]byte, error) {
	if readonly {
		return nil, errors.New("legacySubmitProposal method not readonly")
	}

	method := MustMethod(LegacySubmitProposalMethodName)

	var args LegacySubmitProposalArgs
	err := types.ParseMethodArgs(method, &args, contract.Input[4:])
	if err != nil {
		return nil, err
	}

	var amount sdk.Coins
	if args.InitialDeposit.Amount.Sign() > 0 {
		amount = amount.Add(sdk.Coin{
			Denom:  args.InitialDeposit.Denom,
			Amount: sdk.NewIntFromBigInt(args.InitialDeposit.Amount),
		})
	}

	content, _ := v1beta1.ContentFromProposalType(args.Title, args.Description, v1beta1.ProposalTypeText)
	msg, err := v1beta1.NewMsgSubmitProposal(content, amount, contract.Caller().Bytes())
	if err != nil {
		return nil, fmt.Errorf("invalid message: %w", err)
	}

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	msgServer := govkeeper.NewMsgServerImpl(c.govKeeper)
	server := govkeeper.NewLegacyMsgServerImpl(c.accountKeeper.GetModuleAddress(govtypes.ModuleName).String(), msgServer)
	res, err := server.SubmitProposal(ctx, msg)
	if err != nil {
		return nil, err
	}

	if err := c.AddLog(
		evm,
		MustEvent(LegacySubmitProposalEventName),
		[]common.Hash{common.BytesToHash(contract.Caller().Bytes())},
		res.ProposalId,
	); err != nil {
		return nil, err
	}

	return method.Outputs.Pack(res.ProposalId)
}

func (c *Contract) Vote(ctx sdk.Context, evm *vm.EVM, contract *vm.Contract, readonly bool) ([]byte, error) {
	if readonly {
		return nil, errors.New("legacyVote method not readonly")
	}

	method := MustMethod(VoteMethodName)

	var args VoteArgs
	err := types.ParseMethodArgs(method, &args, contract.Input[4:])
	if err != nil {
		return nil, err
	}

	msg := &v1.MsgVote{
		ProposalId: args.ProposalId,
		Voter:      sdk.AccAddress(contract.Caller().Bytes()).String(),
		Option:     v1.VoteOption(args.Option),
		Metadata:   args.Metadata,
	}
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	server := govkeeper.NewMsgServerImpl(c.govKeeper)
	_, err = server.Vote(ctx, msg)
	if err != nil {
		return nil, err
	}

	if err := c.AddLog(
		evm,
		MustEvent(VoteEventName),
		[]common.Hash{common.BytesToHash(contract.Caller().Bytes())},
		args.ProposalId,
		args.Option,
	); err != nil {
		return nil, err
	}

	return method.Outputs.Pack(true)
}
