package gov

import (
	"bytes"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	govv1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/evmos/evmos/v12/utils"

	"github.com/evmos/evmos/v12/x/evm/types"
)

const (
	ProposalGas     = 30_000
	ProposalsGas    = 30_000
	VoteQueryGas    = 30_000
	VotesGas        = 30_000
	DepositQueryGas = 30_000
	DepositsGas     = 30_000
	TallyResultGas  = 30_000

	ProposalMethodName     = "proposal"
	ProposalsMethodName    = "proposals"
	VoteQueryMethodName    = "vote0"
	VotesMethodName        = "votes"
	DepositQueryMethodName = "depositQuery"
	DepositsMethodName     = "deposits"
	TallyResultMethodName  = "tallyResult"
)

// Proposal returns proposal details based on ProposalID
func (c *Contract) Proposal(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := MustMethod(ProposalMethodName)

	var args ProposalArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}
	msg := &govv1.QueryProposalRequest{
		ProposalId: args.ProposalId,
	}

	res, err := c.govKeeper.Proposal(ctx, msg)
	if err != nil {
		return nil, err
	}

	return method.Outputs.Pack(OutputsProposal(*res.Proposal))
}

// Proposals queries all proposals based on given status.
func (c *Contract) Proposals(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := MustMethod(ProposalsMethodName)

	var args ProposalsArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}

	if bytes.Equal(args.Pagination.Key, []byte{0}) {
		args.Pagination.Key = nil
	}

	voter := ""
	if args.Voter != (common.Address{}) {
		voter = sdk.AccAddress(args.Voter.Bytes()).String()
	}

	depositor := ""
	if args.Depositor != (common.Address{}) {
		depositor = sdk.AccAddress(args.Depositor.Bytes()).String()
	}

	msg := &govv1.QueryProposalsRequest{
		ProposalStatus: govv1.ProposalStatus(args.Status),
		Voter:          voter,
		Depositor:      depositor,
		Pagination: &query.PageRequest{
			Key:        args.Pagination.Key,
			Offset:     args.Pagination.Offset,
			Limit:      args.Pagination.Limit,
			CountTotal: args.Pagination.CountTotal,
			Reverse:    args.Pagination.Reverse,
		},
	}

	res, err := c.govKeeper.Proposals(ctx, msg)
	if err != nil {
		return nil, err
	}

	var proposals []Proposal
	for _, proposal := range res.Proposals {
		proposals = append(proposals, OutputsProposal(*proposal))
	}

	var pageResponse PageResponse
	pageResponse.NextKey = res.Pagination.NextKey
	pageResponse.Total = res.Pagination.Total

	return method.Outputs.Pack(proposals, pageResponse)
}

// VoteQuery returns Voted information based on proposalID, voterAddr
func (c *Contract) VoteQuery(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := MustMethod(VoteQueryMethodName)

	var args VoteQueryArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}
	msg := &govv1.QueryVoteRequest{
		ProposalId: args.ProposalId,
		Voter:      sdk.AccAddress(args.Voter.Bytes()).String(),
	}

	res, err := c.govKeeper.Vote(ctx, msg)
	if err != nil {
		return nil, err
	}

	return method.Outputs.Pack(OutputsVote(*res.Vote))
}

// Votes returns single proposal's votes
func (c *Contract) Votes(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := MustMethod(VotesMethodName)

	var args VotesArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}

	if bytes.Equal(args.Pagination.Key, []byte{0}) {
		args.Pagination.Key = nil
	}

	msg := &govv1.QueryVotesRequest{
		ProposalId: args.ProposalId,
		Pagination: &query.PageRequest{
			Key:        args.Pagination.Key,
			Offset:     args.Pagination.Offset,
			Limit:      args.Pagination.Limit,
			CountTotal: args.Pagination.CountTotal,
			Reverse:    args.Pagination.Reverse,
		},
	}

	res, err := c.govKeeper.Votes(ctx, msg)
	if err != nil {
		return nil, err
	}

	var votes []VoteData
	for _, vote := range res.Votes {
		votes = append(votes, OutputsVote(*vote))
	}

	var pageResponse PageResponse
	pageResponse.NextKey = res.Pagination.NextKey
	pageResponse.Total = res.Pagination.Total

	return method.Outputs.Pack(votes, pageResponse)
}

// DepositQuery queries single deposit information based on proposalID, depositAddr.
func (c *Contract) DepositQuery(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := MustMethod(DepositQueryMethodName)

	var args DepositQueryArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}
	msg := &govv1.QueryDepositRequest{
		ProposalId: args.ProposalId,
		Depositor:  sdk.AccAddress(args.Depositor.Bytes()).String(),
	}

	res, err := c.govKeeper.Deposit(ctx, msg)
	if err != nil {
		return nil, err
	}

	return method.Outputs.Pack(OutputsDeposit(*res.Deposit))
}

// Deposits returns single proposal's all deposits
func (c *Contract) Deposits(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := MustMethod(DepositsMethodName)

	var args DepositsArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}

	if bytes.Equal(args.Pagination.Key, []byte{0}) {
		args.Pagination.Key = nil
	}

	msg := &govv1.QueryDepositsRequest{
		ProposalId: args.ProposalId,
		Pagination: &query.PageRequest{
			Key:        args.Pagination.Key,
			Offset:     args.Pagination.Offset,
			Limit:      args.Pagination.Limit,
			CountTotal: args.Pagination.CountTotal,
			Reverse:    args.Pagination.Reverse,
		},
	}

	res, err := c.govKeeper.Deposits(ctx, msg)
	if err != nil {
		return nil, err
	}

	var deposits []DepositData
	for _, vote := range res.Deposits {
		deposits = append(deposits, OutputsDeposit(*vote))
	}

	var pageResponse PageResponse
	pageResponse.NextKey = res.Pagination.NextKey
	pageResponse.Total = res.Pagination.Total

	return method.Outputs.Pack(deposits, pageResponse)
}

// TallyResult queries the tally of a proposal vote.
func (c *Contract) TallyResult(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := MustMethod(TallyResultMethodName)

	var args ProposalArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}
	msg := &govv1.QueryTallyResultRequest{
		ProposalId: args.ProposalId,
	}

	res, err := c.govKeeper.TallyResult(ctx, msg)
	if err != nil {
		return nil, err
	}

	return method.Outputs.Pack(TallyResult(*res.Tally))
}

func OutputsProposal(proposal govv1.Proposal) Proposal {
	var messages []string
	for _, message := range proposal.Messages {
		// TODO must format to readable string
		msg, ok := message.GetCachedValue().(govv1.MsgExecLegacyContent)
		if ok {
			messages = append(messages, msg.String())
			continue
		}
		messages = append(messages, message.String())
	}
	var totalDeposit []Coin
	for _, coin := range proposal.TotalDeposit {
		totalDeposit = append(totalDeposit, Coin{
			Denom:  coin.Denom,
			Amount: coin.Amount.BigInt(),
		})
	}

	return Proposal{
		Id:               proposal.Id,
		Messages:         messages,
		Status:           uint8(proposal.Status),
		FinalTallyResult: TallyResult(*proposal.FinalTallyResult),
		SubmitTime:       proposal.SubmitTime.Unix(),
		DepositEndTime:   proposal.DepositEndTime.Unix(),
		TotalDeposit:     totalDeposit,
		VotingStartTime:  proposal.VotingStartTime.Unix(),
		VotingEndTime:    proposal.VotingEndTime.Unix(),
		Metadata:         proposal.Metadata,
	}
}

func OutputsVote(vote govv1.Vote) VoteData {
	var options []WeightedVoteOption
	for _, option := range vote.Options {
		options = append(options, WeightedVoteOption{
			Option: uint8(option.Option),
			Weight: option.Weight,
		})
	}

	return VoteData{
		ProposalId: vote.ProposalId,
		Voter:      utils.AccAddressMustToHexAddress(vote.Voter),
		Options:    options,
		Metadata:   vote.Metadata,
	}
}

func OutputsDeposit(deposit govv1.Deposit) DepositData {
	var amount []Coin
	for _, coin := range deposit.Amount {
		amount = append(amount, Coin{
			Denom:  coin.Denom,
			Amount: coin.Amount.BigInt(),
		})
	}

	return DepositData{
		ProposalId: deposit.ProposalId,
		Depositor:  utils.AccAddressMustToHexAddress(deposit.Depositor),
		Amount:     amount,
	}
}
