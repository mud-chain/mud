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

func (suite *InflationTestSuite) TestInflationRate() {
	testCases := []struct {
		name             string
		params           Params
		inflationAmount  math.Int
		bondedTokens     math.Int
		expInflationRate string
		expPass          bool
	}{
		{
			"pass - goal bonded ratio",
			DefaultParams(),
			sdk.NewInt(20),
			sdk.NewInt(100),
			"0.200000000",
			true,
		},
		{
			"pass - below goal bonded ratio",
			DefaultParams(),
			sdk.NewInt(30),
			sdk.NewInt(100),
			"0.200000000",
			true,
		},
		{
			"pass - above goal bonded ratio",
			DefaultParams(),
			sdk.NewInt(10),
			sdk.NewInt(100),
			"0.100000000",
			true,
		},
	}
	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.name), func() {
			newInflationRate := InflationRate(
				tc.params,
				tc.inflationAmount,
				tc.bondedTokens,
			)

			suite.Require().Equal(tc.expInflationRate, newInflationRate.String()[:11])
		})
	}
}
