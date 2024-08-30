package types

import (
	fmt "fmt"
	"testing"

	"cosmossdk.io/math"

	"github.com/stretchr/testify/suite"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type InflationTestSuite struct {
	suite.Suite
}

func TestInflationSuite(t *testing.T) {
	suite.Run(t, new(InflationTestSuite))
}

func (suite *InflationTestSuite) TestNextInflationRate() {
	epochsPerPeriod := int64(365)

	testCases := []struct {
		name             string
		params           Params
		bondedRatio      sdk.Dec
		inflationRate    sdk.Dec
		expInflationRate string
		expPass          bool
	}{
		{
			"pass - goal bonded ratio",
			DefaultParams(),
			sdk.NewDecWithPrec(67, 2),
			sdk.NewDecWithPrec(13, 2),
			"0.130000000",
			true,
		},
		{
			"pass - below goal bonded ratio",
			DefaultParams(),
			sdk.NewDecWithPrec(50, 2),
			sdk.NewDecWithPrec(13, 2),
			"0.130090370",
			true,
		},
		{
			"pass - above goal bonded ratio",
			DefaultParams(),
			sdk.NewDecWithPrec(70, 2),
			sdk.NewDecWithPrec(13, 2),
			"0.129984052",
			true,
		},
	}
	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.name), func() {
			newInflationRate := NextInflationRate(
				tc.params,
				tc.bondedRatio,
				tc.inflationRate,
				epochsPerPeriod,
			)

			suite.Require().Equal(tc.expInflationRate, newInflationRate.String()[:11])
		})
	}
}

func (suite *InflationTestSuite) TestEpochProvision() {
	epochsPerPeriod := int64(365)

	testCases := []struct {
		name              string
		params            Params
		totalSupply       math.Int
		inflationRate     sdk.Dec
		expEpochProvision math.Int
		expPass           bool
	}{
		{
			"pass - initial perid",
			DefaultParams(),
			sdk.NewInt(100000000),
			sdk.NewDecWithPrec(13, 2),
			sdk.NewInt(35616),
			true,
		},
		{
			"pass - period 1",
			DefaultParams(),
			sdk.NewInt(100035616),
			sdk.NewDecWithPrec(13, 2),
			sdk.NewInt(35629),
			true,
		},
	}
	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.name), func() {
			epochMintProvisions := EpochProvision(
				tc.params,
				tc.totalSupply,
				epochsPerPeriod,
				tc.inflationRate,
			)

			suite.Require().Equal(tc.expEpochProvision, epochMintProvisions.Amount)
		})
	}
}
