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
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/evmos/evmos/v12/cmd/config"
	"github.com/evmos/evmos/v12/x/inflation/types"
)

// MintAndAllocateInflation performs inflation minting and allocation
func (k Keeper) MintAndAllocateInflation(
	ctx sdk.Context,
	bondedTokens sdkmath.Int,
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

	distributionCoin := coin

	// check inflation module has enough coin for burn and reward
	moduleAddr := k.accountKeeper.GetModuleAddress(types.ModuleName)
	leftCoin := k.bankKeeper.GetBalance(ctx, moduleAddr, coin.Denom)
	if leftCoin.Amount.LT(coin.Amount) {
		if leftCoin.Amount.IsZero() {
			return nil, nil, nil
		} else {
			distributionCoin = leftCoin
		}
	}

	// If the amount of staking is excessive, we have an annualized upper limit.
	// Once it exceeds this annualized upper limit, we will directly destroy the surplus coins.
	inflationMaxAmount := sdk.NewDecFromInt(bondedTokens).Mul(params.InflationMax).TruncateInt()
	if distributionCoin.Amount.GT(inflationMaxAmount) {
		burnAmount := distributionCoin.Amount.Sub(inflationMaxAmount)
		burnCoins := sdk.Coins{
			sdk.Coin{
				Denom:  distributionCoin.Denom,
				Amount: burnAmount,
			},
		}
		if err := k.bankKeeper.BurnCoins(ctx, types.ModuleName, burnCoins); err != nil {
			return nil, nil, err
		}

		distributionCoin = sdk.Coin{
			Denom:  distributionCoin.Denom,
			Amount: inflationMaxAmount,
		}
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

// GetInflationAmount get the inflation amount coin
func (k Keeper) GetInflationAmount(ctx sdk.Context) (amount sdk.Coin) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.KeyPrefixInflationAmount)
	if len(bz) == 0 {
		return sdk.Coin{
			Denom:  config.BaseDenom,
			Amount: sdkmath.NewInt(0),
		}
	}
	err := amount.Unmarshal(bz)
	if err != nil {
		panic(err)
	}

	return amount
}

// SetInflationAmount set the inflation amount coin
func (k Keeper) SetInflationAmount(ctx sdk.Context, amount sdk.Coin) {
	store := ctx.KVStore(k.storeKey)
	bytes, err := amount.Marshal()
	if err != nil {
		panic(err)
	}
	store.Set(types.KeyPrefixInflationAmount, bytes)
}

// GetEpochMintProvision get the epoch mint provision
func (k Keeper) GetEpochMintProvision(ctx sdk.Context) (epochMintProvision sdk.Coin) {
	inflationAmount := k.GetInflationAmount(ctx)
	epochsPerPeriod := k.GetEpochsPerPeriod(ctx)

	epochMintProvision = types.EpochProvision(inflationAmount, epochsPerPeriod)

	return epochMintProvision
}
