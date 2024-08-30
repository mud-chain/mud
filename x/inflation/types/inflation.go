package types

import (
	"cosmossdk.io/math"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NextInflationRate returns the new inflation rate for the next epoch.
func NextInflationRate(params Params, bondedRatio, inflation sdk.Dec, epochsPerPeriod int64) sdk.Dec {
	// The target annual inflation rate is recalculated for each previsions cycle. The
	// inflation is also subject to a rate change (positive or negative) depending on
	// the distance from the desired ratio (67%). The maximum rate change possible is
	// defined to be 13% per year, however the annual inflation is capped as between
	// 7% and 20%.

	// (1 - bondedRatio/GoalBonded) * InflationRateChange
	inflationRateChangePerYear := sdk.OneDec().
		Sub(bondedRatio.Quo(params.GoalBonded)).
		Mul(params.InflationRateChange)
	inflationRateChange := inflationRateChangePerYear.Quo(sdk.NewDec(epochsPerPeriod))

	// adjust the new annual inflation for this next cycle
	newInflation := inflation.Add(inflationRateChange) // note inflationRateChange may be negative
	if newInflation.GT(params.InflationMax) {
		newInflation = params.InflationMax
	}
	if newInflation.LT(params.InflationMin) {
		newInflation = params.InflationMin
	}

	return newInflation
}

// EpochProvision returns the provisions for an epoch based on the current inflation.
func EpochProvision(params Params, totalStakingSupply math.Int, epochsPerPeriod int64, inflation sdk.Dec) sdk.Coin {
	annualProvisions := inflation.MulInt(totalStakingSupply)
	provisionAmt := annualProvisions.Quo(sdk.NewDec(epochsPerPeriod))
	return sdk.NewCoin(params.MintDenom, provisionAmt.TruncateInt())
}
