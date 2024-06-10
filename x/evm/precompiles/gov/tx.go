package gov

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	govkeeper "github.com/cosmos/cosmos-sdk/x/gov/keeper"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	govv1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	govv1beta1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
	"github.com/cosmos/cosmos-sdk/x/params/types/proposal"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"

	erc20types "github.com/evmos/evmos/v12/x/erc20/types"
	"github.com/evmos/evmos/v12/x/evm/types"
	feemarkettypes "github.com/evmos/evmos/v12/x/feemarket/types"
	inflationtypes "github.com/evmos/evmos/v12/x/inflation/types"
)

const (
	LegacySubmitProposalGas = 60_000
	SubmitProposalGas       = 60_000
	VoteGas                 = 60_000
	VoteWeightedGas         = 60_000
	DepositGas              = 60_000

	LegacySubmitProposalMethodName = "legacySubmitProposal"
	SubmitProposalMethodName       = "submitProposal"
	VoteMethodName                 = "vote"
	VoteWeightedMethodName         = "voteWeighted"
	DepositMethodName              = "deposit"

	LegacySubmitProposalEventName = "LegacySubmitProposal"
	SubmitProposalEventName       = "SubmitProposal"
	VoteEventName                 = "Vote"
	VoteWeightedEventName         = "VoteWeighted"
	DepositEventName              = "Deposit"
)

func (c *Contract) LegacySubmitProposal(ctx sdk.Context, evm *vm.EVM, contract *vm.Contract, readonly bool) ([]byte, error) {
	if readonly {
		return nil, errors.New("legacySubmitProposal method readonly")
	}
	if evm.Origin != contract.Caller() {
		return nil, errors.New("only allow EOA contract call this method")
	}

	method := MustMethod(LegacySubmitProposalMethodName)

	var args LegacySubmitProposalArgs
	err := types.ParseMethodArgs(method, &args, contract.Input[4:])
	if err != nil {
		return nil, err
	}

	var amount sdk.Coins
	for _, deposit := range args.InitialDeposit {
		if deposit.Amount.Sign() > 0 {
			amount = amount.Add(sdk.Coin{
				Denom:  deposit.Denom,
				Amount: sdk.NewIntFromBigInt(deposit.Amount),
			})
		}
	}

	content, _ := govv1beta1.ContentFromProposalType(args.Title, args.Description, govv1beta1.ProposalTypeText)
	msg, err := govv1beta1.NewMsgSubmitProposal(content, amount, contract.Caller().Bytes())
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

func (c *Contract) SubmitProposal(ctx sdk.Context, evm *vm.EVM, contract *vm.Contract, readonly bool) ([]byte, error) {
	if readonly {
		return nil, errors.New("submitProposal method readonly")
	}
	if evm.Origin != contract.Caller() {
		return nil, errors.New("only allow EOA contract call this method")
	}

	method := MustMethod(SubmitProposalMethodName)

	var args SubmitProposalArgs
	err := types.ParseMethodArgs(method, &args, contract.Input[4:])
	if err != nil {
		return nil, err
	}

	var messages []json.RawMessage
	err = json.Unmarshal([]byte(args.Messages), &messages)
	if err != nil {
		return nil, err
	}

	interfaceRegistry := codectypes.NewInterfaceRegistry()
	feemarkettypes.RegisterInterfaces(interfaceRegistry)
	proposal.RegisterInterfaces(interfaceRegistry)
	upgradetypes.RegisterInterfaces(interfaceRegistry)
	govv1.RegisterInterfaces(interfaceRegistry)
	govv1beta1.RegisterInterfaces(interfaceRegistry)
	types.RegisterInterfaces(interfaceRegistry)
	erc20types.RegisterInterfaces(interfaceRegistry)
	inflationtypes.RegisterInterfaces(interfaceRegistry)
	ethosCodec := codec.NewProtoCodec(interfaceRegistry)

	msgs := make([]sdk.Msg, len(messages))
	for i, message := range messages {
		var msg sdk.Msg
		err := ethosCodec.UnmarshalInterfaceJSON(message, &msg)
		if err != nil {
			return nil, err
		}

		msgs[i] = msg
	}

	var amount sdk.Coins
	for _, deposit := range args.InitialDeposit {
		if deposit.Amount.Sign() > 0 {
			amount = amount.Add(sdk.Coin{
				Denom:  deposit.Denom,
				Amount: sdk.NewIntFromBigInt(deposit.Amount),
			})
		}
	}

	msg, err := govv1.NewMsgSubmitProposal(msgs, amount, sdk.AccAddress(contract.Caller().Bytes()).String(), args.Metadata)
	if err != nil {
		return nil, fmt.Errorf("invalid message: %w", err)
	}

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	server := govkeeper.NewMsgServerImpl(c.govKeeper)
	res, err := server.SubmitProposal(ctx, msg)
	if err != nil {
		return nil, err
	}

	if err := c.AddLog(
		evm,
		MustEvent(SubmitProposalEventName),
		[]common.Hash{common.BytesToHash(contract.Caller().Bytes())},
		res.ProposalId,
	); err != nil {
		return nil, err
	}

	return method.Outputs.Pack(res.ProposalId)
}

func (c *Contract) Vote(ctx sdk.Context, evm *vm.EVM, contract *vm.Contract, readonly bool) ([]byte, error) {
	if readonly {
		return nil, errors.New("vote method readonly")
	}
	if evm.Origin != contract.Caller() {
		return nil, errors.New("only allow EOA contract call this method")
	}

	method := MustMethod(VoteMethodName)

	var args VoteArgs
	err := types.ParseMethodArgs(method, &args, contract.Input[4:])
	if err != nil {
		return nil, err
	}

	msg := &govv1.MsgVote{
		ProposalId: args.ProposalId,
		Voter:      sdk.AccAddress(contract.Caller().Bytes()).String(),
		Option:     govv1.VoteOption(args.Option),
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

func (c *Contract) VoteWeighted(ctx sdk.Context, evm *vm.EVM, contract *vm.Contract, readonly bool) ([]byte, error) {
	if readonly {
		return nil, errors.New("voteWeighted method readonly")
	}
	if evm.Origin != contract.Caller() {
		return nil, errors.New("only allow EOA contract call this method")
	}

	method := MustMethod(VoteWeightedMethodName)

	var args VoteWeightedArgs
	err := types.ParseMethodArgs(method, &args, contract.Input[4:])
	if err != nil {
		return nil, err
	}

	var options []*govv1.WeightedVoteOption
	for _, option := range args.Options {
		options = append(options, &govv1.WeightedVoteOption{
			Option: govv1.VoteOption(option.Option),
			Weight: option.Weight,
		})
	}

	msg := &govv1.MsgVoteWeighted{
		ProposalId: args.ProposalId,
		Voter:      sdk.AccAddress(contract.Caller().Bytes()).String(),
		Options:    options,
		Metadata:   args.Metadata,
	}
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	server := govkeeper.NewMsgServerImpl(c.govKeeper)
	_, err = server.VoteWeighted(ctx, msg)
	if err != nil {
		return nil, err
	}

	if err := c.AddLog(
		evm,
		MustEvent(VoteWeightedEventName),
		[]common.Hash{common.BytesToHash(contract.Caller().Bytes())},
		args.ProposalId,
	); err != nil {
		return nil, err
	}

	return method.Outputs.Pack(true)
}

func (c *Contract) Deposit(ctx sdk.Context, evm *vm.EVM, contract *vm.Contract, readonly bool) ([]byte, error) {
	if readonly {
		return nil, errors.New("deposit method readonly")
	}
	if evm.Origin != contract.Caller() {
		return nil, errors.New("only allow EOA contract call this method")
	}

	method := MustMethod(DepositMethodName)

	var args DepositArgs
	err := types.ParseMethodArgs(method, &args, contract.Input[4:])
	if err != nil {
		return nil, err
	}

	var amount sdk.Coins
	for _, deposit := range args.Amount {
		if deposit.Amount.Sign() > 0 {
			amount = amount.Add(sdk.Coin{
				Denom:  deposit.Denom,
				Amount: sdk.NewIntFromBigInt(deposit.Amount),
			})
		}
	}

	msg := &govv1.MsgDeposit{
		ProposalId: args.ProposalId,
		Depositor:  sdk.AccAddress(contract.Caller().Bytes()).String(),
		Amount:     amount,
	}
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	server := govkeeper.NewMsgServerImpl(c.govKeeper)
	_, err = server.Deposit(ctx, msg)
	if err != nil {
		return nil, err
	}

	if err := c.AddLog(
		evm,
		MustEvent(DepositEventName),
		[]common.Hash{common.BytesToHash(contract.Caller().Bytes())},
		args.ProposalId,
	); err != nil {
		return nil, err
	}

	return method.Outputs.Pack(true)
}
