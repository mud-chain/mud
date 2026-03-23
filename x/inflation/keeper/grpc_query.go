// Copyright 2022 Evmos Foundation
// This file is part of the Evmos Network packages.
//
// Evmos is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The Evmos packages are distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the Evmos packages. If not, see https://github.com/evmos/evmos/blob/main/LICENSE

package keeper

import (
	"context"
	"github.com/evmos/evmos/v12/cmd/config"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/evmos/evmos/v12/x/inflation/types"
)

var _ types.QueryServer = Keeper{}

// Period returns the current period of the inflation module.
func (k Keeper) Period(
	c context.Context,
	_ *types.QueryPeriodRequest,
) (*types.QueryPeriodResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	period := k.GetPeriod(ctx)
	return &types.QueryPeriodResponse{Period: period}, nil
}

// EpochProvision returns the EpochProvision of the inflation module.
func (k Keeper) EpochProvision(
	c context.Context,
	_ *types.QueryEpochProvisionRequest,
) (*types.QueryEpochProvisionResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	params := k.GetParams(ctx)
	bondedTokens := k.stakingKeeper.TotalBondedTokens(ctx)
	epochsPerPeriod := k.GetEpochsPerPeriod(ctx)
	inflationMaxAmount := sdk.NewDecFromInt(bondedTokens).Mul(params.InflationMax).QuoInt64(epochsPerPeriod).TruncateInt()
	mint := k.GetEpochMintProvision(ctx)

	// Due to excessive delegate, the annualized rate has exceeded the annual coin - minting limit.
	if inflationMaxAmount.GT(mint.Amount) {
		inflationMaxAmount = mint.Amount
	}

	reward := sdk.Coin{
		Denom:  mint.Denom,
		Amount: inflationMaxAmount,
	}
	burn := sdk.Coin{
		Denom:  mint.Denom,
		Amount: mint.Amount.Sub(reward.Amount),
	}

	return &types.QueryEpochProvisionResponse{Mint: mint, Reward: reward, Burn: burn}, nil
}

// SkippedEpochs returns the number of skipped Epochs of the inflation module.
func (k Keeper) SkippedEpochs(
	c context.Context,
	_ *types.QuerySkippedEpochsRequest,
) (*types.QuerySkippedEpochsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	skippedEpochs := k.GetSkippedEpochs(ctx)
	return &types.QuerySkippedEpochsResponse{SkippedEpochs: skippedEpochs}, nil
}

// InflationRate returns the inflation rate for the current period.
func (k Keeper) InflationRate(
	c context.Context,
	_ *types.QueryInflationRateRequest,
) (*types.QueryInflationRateResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	inflationAmount := k.GetInflationAmount(ctx)
	bondedTokens := k.stakingKeeper.TotalBondedTokens(ctx)

	inflationRate := types.InflationRate(k.GetParams(ctx), inflationAmount.Amount, bondedTokens)
	inflationRate = inflationRate.MulInt64(100)

	return &types.QueryInflationRateResponse{InflationRate: inflationRate}, nil
}

// CirculatingSupply returns the total supply in circulation excluding the team
// allocation in the first year
func (k Keeper) CirculatingSupply(
	c context.Context,
	_ *types.QueryCirculatingSupplyRequest,
) (*types.QueryCirculatingSupplyResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	circulatingSupply := k.GetCirculatingSupply(ctx, config.BaseDenom)
	coin := sdk.NewDecCoinFromDec(config.BaseDenom, circulatingSupply)

	return &types.QueryCirculatingSupplyResponse{CirculatingSupply: coin}, nil
}

// Params returns params of the mint module.
func (k Keeper) Params(
	c context.Context,
	_ *types.QueryParamsRequest,
) (*types.QueryParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	params := k.GetParams(ctx)
	return &types.QueryParamsResponse{Params: params}, nil
}
