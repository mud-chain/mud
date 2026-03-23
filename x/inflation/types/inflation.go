package types

import (
	"cosmossdk.io/math"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InflationRate returns the new inflation rate for the next epoch.
func InflationRate(params Params, inflationAmount, bondedTokens math.Int) sdk.Dec {
	if bondedTokens.IsZero() {
		return sdk.Dec{}
	}

	bondedTokensDec := sdk.NewDecFromBigInt(bondedTokens.BigInt())
	inflationAmountDec := sdk.NewDecFromBigInt(inflationAmount.BigInt())
	newInflation := inflationAmountDec.Quo(bondedTokensDec)
	if newInflation.GT(params.InflationMax) {
		newInflation = params.InflationMax
	}
	return newInflation
}

// EpochProvision returns the provisions for an epoch based on the current inflation.
func EpochProvision(inflationAmount sdk.Coin, epochsPerPeriod int64) sdk.Coin {
	provisionAmt := sdk.NewDecFromBigInt(inflationAmount.Amount.BigInt()).Quo(sdk.NewDec(epochsPerPeriod))
	return sdk.NewCoin(inflationAmount.Denom, provisionAmt.TruncateInt())
}
