package types

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"
)

type ParamsTestSuite struct {
	suite.Suite
}

func TestParamsTestSuite(t *testing.T) {
	suite.Run(t, new(ParamsTestSuite))
}

func (suite *ParamsTestSuite) TestParamsValidate() {
	validInflationDistribution := InflationDistribution{
		StakingRewards: sdk.NewDecWithPrec(80, 2),
		CommunityPool:  sdk.NewDecWithPrec(20, 2),
	}

	inflationMax := sdk.NewDecWithPrec(20, 2)
	inflationDecay := sdk.NewDecWithPrec(8, 1)

	testCases := []struct {
		name     string
		params   Params
		expError bool
	}{
		{
			"default",
			DefaultParams(),
			false,
		},
		{
			"valid",
			NewParams(
				validInflationDistribution,
				true,
				inflationMax,
				inflationDecay,
			),
			false,
		},
		{
			"valid param literal",
			Params{
				InflationDistribution: validInflationDistribution,
				EnableInflation:       true,
				InflationMax:          inflationMax,
				InflationDecay:        inflationDecay,
			},
			false,
		},
		{
			"invalid - inflation distribution - negative staking rewards",
			Params{
				InflationDistribution: InflationDistribution{
					StakingRewards: sdk.OneDec().Neg(),
					CommunityPool:  sdk.NewDecWithPrec(20, 2),
				},
				EnableInflation: true,
				InflationMax:    inflationMax,
				InflationDecay:  inflationDecay,
			},
			true,
		},
		{
			"invalid - inflation distribution - negative community pool rewards",
			Params{
				InflationDistribution: InflationDistribution{
					StakingRewards: sdk.NewDecWithPrec(80, 2),
					CommunityPool:  sdk.OneDec().Neg(),
				},
				EnableInflation: true,
				InflationMax:    inflationMax,
				InflationDecay:  inflationDecay,
			},
			true,
		},
		{
			"invalid - inflation distribution - total distribution ratio unequal 1",
			Params{
				InflationDistribution: InflationDistribution{
					StakingRewards: sdk.NewDecWithPrec(80, 2),
					CommunityPool:  sdk.NewDecWithPrec(30, 2),
				},
				EnableInflation: true,
				InflationMax:    inflationMax,
				InflationDecay:  inflationDecay,
			},
			true,
		},
		{
			"invalid - inflation max - negative inflation max",
			Params{
				InflationDistribution: validInflationDistribution,
				EnableInflation:       true,
				InflationMax:          inflationMax.Neg(),
				InflationDecay:        inflationDecay,
			},
			true,
		},
		{
			"invalid - inflation max - inflation max is greater than 1",
			Params{
				InflationDistribution: validInflationDistribution,
				EnableInflation:       true,
				InflationMax:          sdk.NewDecWithPrec(110, 2),
				InflationDecay:        inflationDecay,
			},
			true,
		},
	}

	for _, tc := range testCases {
		err := tc.params.Validate()

		if tc.expError {
			suite.Require().Error(err, tc.name)
		} else {
			suite.Require().NoError(err, tc.name)
		}
	}
}
