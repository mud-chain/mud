package evidence

import (
	"bytes"
	"encoding/hex"
	"fmt"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	evidencetypes "github.com/cosmos/cosmos-sdk/x/evidence/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/evmos/evmos/v12/x/evm/types"
)

const (
	EvidenceGas    = 30_000
	AllEvidenceGas = 50_000

	EvidenceMethodName    = "evidence"
	AllEvidenceMethodName = "allEvidence"
)

func (c *Contract) Evidence(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := MustMethod(EvidenceMethodName)

	var args EvidenceArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}

	hashBytes, err := hex.DecodeString(args.EvidenceHash)
	if err != nil {
		return nil, fmt.Errorf("invalid evidence hash: %s", err)
	}

	msg := &evidencetypes.QueryEvidenceRequest{EvidenceHash: hashBytes}

	res, err := c.evidenceKeeper.Evidence(ctx, msg)
	if err != nil {
		return nil, err
	}

	valSigningInfo := OutPutEvidence(*res.Evidence)

	return method.Outputs.Pack(valSigningInfo)
}

func (c *Contract) AllEvidence(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := MustMethod(AllEvidenceMethodName)

	var args AllEvidenceArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}

	if bytes.Equal(args.Pagination.Key, []byte{0}) {
		args.Pagination.Key = nil
	}
	msg := &evidencetypes.QueryAllEvidenceRequest{
		Pagination: &query.PageRequest{
			Key:        args.Pagination.Key,
			Offset:     args.Pagination.Offset,
			Limit:      args.Pagination.Limit,
			CountTotal: args.Pagination.CountTotal,
			Reverse:    args.Pagination.Reverse,
		},
	}

	res, err := c.evidenceKeeper.AllEvidence(ctx, msg)
	if err != nil {
		return nil, err
	}

	var evidences []Equivocation
	for _, evidence := range res.Evidence {
		evidences = append(evidences, OutPutEvidence(*evidence))
	}

	var pageResponse PageResponse
	pageResponse.NextKey = res.Pagination.NextKey
	pageResponse.Total = res.Pagination.Total

	return method.Outputs.Pack(evidences, pageResponse)
}

func OutPutEvidence(evidence codectypes.Any) Equivocation {
	e := evidence.GetCachedValue().(*evidencetypes.Equivocation)

	consAddress, err := sdk.ConsAddressFromBech32(e.ConsensusAddress)
	var hexAddress common.Address
	if err == nil {
		hexAddress = common.BytesToAddress(consAddress)
	}

	return Equivocation{
		Height:           e.Height,
		Time:             e.Time.UnixMilli(),
		Power:            e.Power,
		ConsensusAddress: hexAddress,
		EvidenceHash:     e.Hash().String(),
	}
}
