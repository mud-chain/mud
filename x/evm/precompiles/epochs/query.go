package epochs

import (
	"bytes"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/ethereum/go-ethereum/core/vm"
	epochstypes "github.com/evmos/evmos/v12/x/epochs/types"
	"github.com/evmos/evmos/v12/x/evm/types"
)

const (
	EpochInfosGas   = 50_000
	CurrentEpochGas = 30_000

	EpochInfosMethodName   = "epochInfos"
	CurrentEpochMethodName = "currentEpoch"
)

// EpochInfos provide running epochInfos
func (c *Contract) EpochInfos(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := MustMethod(EpochInfosMethodName)

	var args EpochInfosArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}

	if bytes.Equal(args.Pagination.Key, []byte{0}) {
		args.Pagination.Key = nil
	}
	msg := &epochstypes.QueryEpochsInfoRequest{
		Pagination: &query.PageRequest{
			Key:        args.Pagination.Key,
			Offset:     args.Pagination.Offset,
			Limit:      args.Pagination.Limit,
			CountTotal: args.Pagination.CountTotal,
			Reverse:    args.Pagination.Reverse,
		},
	}

	res, err := c.epochsKeeper.EpochInfos(ctx, msg)
	if err != nil {
		return nil, err
	}

	var epochs []EpochInfo
	for _, epoch := range res.Epochs {
		epochs = append(epochs, EpochInfo{
			Identifier:              epoch.Identifier,
			StartTime:               epoch.StartTime.UnixMilli(),
			Duration:                epoch.Duration.Milliseconds(),
			CurrentEpoch:            epoch.CurrentEpoch,
			CurrentEpochStartTime:   epoch.CurrentEpochStartTime.UnixMilli(),
			EpochCountingStarted:    epoch.EpochCountingStarted,
			CurrentEpochStartHeight: epoch.CurrentEpochStartHeight,
		})
	}

	var pageResponse PageResponse
	pageResponse.NextKey = res.Pagination.NextKey
	pageResponse.Total = res.Pagination.Total

	return method.Outputs.Pack(epochs, pageResponse)
}

// CurrentEpoch provides current epoch of specified identifier
func (c *Contract) CurrentEpoch(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := MustMethod(CurrentEpochMethodName)

	var args CurrentEpochArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}

	msg := &epochstypes.QueryCurrentEpochRequest{Identifier: args.Identifier}

	res, err := c.epochsKeeper.CurrentEpoch(ctx, msg)
	if err != nil {
		return nil, err
	}

	return method.Outputs.Pack(res.CurrentEpoch)
}
