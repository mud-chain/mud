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
	sdk "github.com/cosmos/cosmos-sdk/types"

	evmostypes "github.com/evmos/evmos/v12/types"

	"github.com/evmos/evmos/v12/x/inflation/types"
)

// 200M token at year 4 allocated to the team
var teamAlloc = sdk.NewInt(200_000_000).Mul(evmostypes.PowerReduction)

// MintAndAllocateInflation performs inflation minting and allocation
func (k Keeper) MintAndAllocateInflation(
	ctx sdk.Context,
	coin sdk.Coin,
	params types.Params,
) (
	staking, communityPool sdk.Coins,
	err error,
) {
	// skip as no coins need to be minted
	if coin.Amount.IsNil() || !coin.Amount.IsPositive() {
		return nil, nil, nil
	}

	// check inflation module has enough coin for burn and reward
	moduleAddr := k.accountKeeper.GetModuleAddress(types.ModuleName)
	if k.bankKeeper.GetBalance(ctx, moduleAddr, coin.Denom).Amount.LT(coin.Amount) {
		return nil, nil, nil
	}

	// 1/3 will be directly burn, and 2/3 will be used to reward validators.
	burnAmount := coin.Amount.QuoRaw(3)
	burnCoins := sdk.Coins{
		sdk.Coin{
			Denom:  coin.Denom,
			Amount: burnAmount,
		},
	}
	if err := k.bankKeeper.BurnCoins(ctx, types.ModuleName, burnCoins); err != nil {
		return nil, nil, err
	}

	distributionCoin := sdk.Coin{
		Denom:  coin.Denom,
		Amount: coin.Amount.Sub(burnAmount),
	}

	// Allocate minted coins according to allocation proportions (staking, community pool)
	return k.AllocateExponentialInflation(ctx, distributionCoin, params)
}

// MintCoins implements an alias call to the underlying supply keeper's
// MintCoins to be used in BeginBlocker.
func (k Keeper) MintCoins(ctx sdk.Context, coin sdk.Coin) error {
	coins := sdk.Coins{coin}
	return k.bankKeeper.MintCoins(ctx, types.ModuleName, coins)
}

// AllocateExponentialInflation allocates coins from the inflation to external
// modules according to allocation proportions:
//   - staking rewards -> sdk `auth` module fee collector
//   - community pool -> `sdk `distr` module community pool
func (k Keeper) AllocateExponentialInflation(
	ctx sdk.Context,
	mintedCoin sdk.Coin,
	params types.Params,
) (
	staking, communityPool sdk.Coins,
	err error,
) {
	distribution := params.InflationDistribution

	// Allocate staking rewards into fee collector account
	staking = sdk.Coins{k.GetProportions(ctx, mintedCoin, distribution.StakingRewards)}

	if err := k.bankKeeper.SendCoinsFromModuleToModule(
		ctx,
		types.ModuleName,
		k.feeCollectorName,
		staking,
	); err != nil {
		return nil, nil, err
	}

	communityPool = sdk.Coins{
		sdk.Coin{
			Denom:  mintedCoin.Denom,
			Amount: mintedCoin.Amount.Sub(staking.AmountOf(mintedCoin.Denom)),
		},
	}

	moduleAddr := k.accountKeeper.GetModuleAddress(types.ModuleName)
	err = k.distrKeeper.FundCommunityPool(
		ctx,
		communityPool,
		moduleAddr,
	)
	if err != nil {
		return nil, nil, err
	}

	return staking, communityPool, nil
}

// GetAllocationProportion calculates the proportion of coins that is to be
// allocated during inflation for a given distribution.
func (k Keeper) GetProportions(
	_ sdk.Context,
	coin sdk.Coin,
	distribution sdk.Dec,
) sdk.Coin {
	return sdk.Coin{
		Denom:  coin.Denom,
		Amount: sdk.NewDecFromInt(coin.Amount).Mul(distribution).TruncateInt(),
	}
}

// BondedRatio the fraction of the staking tokens which are currently bonded
func (k Keeper) BondedRatio(ctx sdk.Context) sdk.Dec {
	bondedRatio := k.stakingKeeper.BondedRatio(ctx)
	return bondedRatio
}

// GetCirculatingSupply returns the bank supply of the mintDenom
func (k Keeper) GetCirculatingSupply(ctx sdk.Context, mintDenom string) sdk.Dec {
	circulatingSupply := sdk.NewDecFromInt(k.bankKeeper.GetSupply(ctx, mintDenom).Amount)

	return circulatingSupply
}

// GetInflation get the epoch mint inflation
func (k Keeper) GetInflation(ctx sdk.Context) (inflation sdk.Dec) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.KeyPrefixInflation)
	if len(bz) == 0 {
		return sdk.ZeroDec()
	}
	err := inflation.Unmarshal(bz)
	if err != nil {
		panic(err)
	}

	return inflation
}

// SetInflation set the epoch mint inflation
func (k Keeper) SetInflation(ctx sdk.Context, inflation sdk.Dec) {
	store := ctx.KVStore(k.storeKey)
	bytes, err := inflation.Marshal()
	if err != nil {
		panic(err)
	}
	store.Set(types.KeyPrefixInflation, bytes)
}

// GetEpochMintProvision get the epoch mint provision
func (k Keeper) GetEpochMintProvision(ctx sdk.Context) (epochMintProvision sdk.Coin) {
	inflation := types.NextInflationRate(k.GetParams(ctx), k.BondedRatio(ctx), k.GetInflation(ctx), k.GetEpochsPerPeriod(ctx))
	epochMintProvision = types.EpochProvision(k.GetParams(ctx), k.stakingKeeper.StakingTokenSupply(ctx), k.GetEpochsPerPeriod(ctx), inflation)

	return epochMintProvision
}
