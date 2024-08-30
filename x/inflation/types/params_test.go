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
	inflationRateChange := sdk.NewDecWithPrec(13, 2)
	inflationMax := sdk.NewDecWithPrec(20, 2)
	inflationMin := sdk.NewDecWithPrec(7, 2)
	goalBonded := sdk.NewDecWithPrec(67, 2)

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
				"acbo",
				validInflationDistribution,
				true,
				inflationRateChange,
				inflationMax,
				inflationMin,
				goalBonded,
			),
			false,
		},
		{
			"valid param literal",
			Params{
				MintDenom:             "acbo",
				InflationDistribution: validInflationDistribution,
				EnableInflation:       true,
				InflationRateChange:   inflationRateChange,
				InflationMax:          inflationMax,
				InflationMin:          inflationMin,
				GoalBonded:            goalBonded,
			},
			false,
		},
		{
			"invalid - denom",
			NewParams(
				"/acbo",
				validInflationDistribution,
				true,
				inflationRateChange,
				inflationMax,
				inflationMin,
				goalBonded,
			),
			true,
		},
		{
			"invalid - denom",
			Params{
				MintDenom:             "",
				InflationDistribution: validInflationDistribution,
				EnableInflation:       true,
				InflationRateChange:   inflationRateChange,
				InflationMax:          inflationMax,
				InflationMin:          inflationMin,
				GoalBonded:            goalBonded,
			},
			true,
		},

		{
			"invalid - inflation distribution - negative staking rewards",
			Params{
				MintDenom: "acbo",
				InflationDistribution: InflationDistribution{
					StakingRewards: sdk.OneDec().Neg(),
					CommunityPool:  sdk.NewDecWithPrec(20, 2),
				},
				EnableInflation:     true,
				InflationRateChange: inflationRateChange,
				InflationMax:        inflationMax,
				InflationMin:        inflationMin,
				GoalBonded:          goalBonded,
			},
			true,
		},
		{
			"invalid - inflation distribution - negative community pool rewards",
			Params{
				MintDenom: "acbo",
				InflationDistribution: InflationDistribution{
					StakingRewards: sdk.NewDecWithPrec(80, 2),
					CommunityPool:  sdk.OneDec().Neg(),
				},
				EnableInflation:     true,
				InflationRateChange: inflationRateChange,
				InflationMax:        inflationMax,
				InflationMin:        inflationMin,
				GoalBonded:          goalBonded,
			},
			true,
		},
		{
			"invalid - inflation distribution - total distribution ratio unequal 1",
			Params{
				MintDenom: "acbo",
				InflationDistribution: InflationDistribution{
					StakingRewards: sdk.NewDecWithPrec(80, 2),
					CommunityPool:  sdk.NewDecWithPrec(30, 2),
				},
				EnableInflation:     true,
				InflationRateChange: inflationRateChange,
				InflationMax:        inflationMax,
				InflationMin:        inflationMin,
				GoalBonded:          goalBonded,
			},
			true,
		},
		{
			"invalid - inflation rate change - negative inflation rate change",
			Params{
				MintDenom:             "acbo",
				InflationDistribution: validInflationDistribution,
				EnableInflation:       true,
				InflationRateChange:   inflationRateChange.Neg(),
				InflationMax:          inflationMax,
				InflationMin:          inflationMin,
				GoalBonded:            goalBonded,
			},
			true,
		},
		{
			"invalid - inflation rate change - inflation rate change is greater than 1",
			Params{
				MintDenom:             "acbo",
				InflationDistribution: validInflationDistribution,
				EnableInflation:       true,
				InflationRateChange:   sdk.NewDecWithPrec(110, 2),
				InflationMax:          inflationMax,
				InflationMin:          inflationMin,
				GoalBonded:            goalBonded,
			},
			true,
		},
		{
			"invalid - inflation max - negative inflation max",
			Params{
				MintDenom:             "acbo",
				InflationDistribution: validInflationDistribution,
				EnableInflation:       true,
				InflationRateChange:   inflationRateChange,
				InflationMax:          inflationMax.Neg(),
				InflationMin:          inflationMin,
				GoalBonded:            goalBonded,
			},
			true,
		},
		{
			"invalid - inflation max - inflation max is greater than 1",
			Params{
				MintDenom:             "acbo",
				InflationDistribution: validInflationDistribution,
				EnableInflation:       true,
				InflationRateChange:   inflationRateChange,
				InflationMax:          sdk.NewDecWithPrec(110, 2),
				InflationMin:          inflationMin,
				GoalBonded:            goalBonded,
			},
			true,
		},
		{
			"invalid - inflation min - negative inflation min",
			Params{
				MintDenom:             "acbo",
				InflationDistribution: validInflationDistribution,
				EnableInflation:       true,
				InflationRateChange:   inflationRateChange,
				InflationMax:          inflationMax,
				InflationMin:          inflationMin.Neg(),
				GoalBonded:            goalBonded,
			},
			true,
		},
		{
			"invalid - inflation min - inflation min is greater than 1",
			Params{
				MintDenom:             "acbo",
				InflationDistribution: validInflationDistribution,
				EnableInflation:       true,
				InflationRateChange:   inflationRateChange,
				InflationMax:          inflationMax,
				InflationMin:          sdk.NewDecWithPrec(110, 2),
				GoalBonded:            goalBonded,
			},
			true,
		},

		{
			"invalid - goal bonded - negative goal bonded",
			Params{
				MintDenom:             "acbo",
				InflationDistribution: validInflationDistribution,
				EnableInflation:       true,
				InflationRateChange:   inflationRateChange,
				InflationMax:          inflationMax,
				InflationMin:          inflationMin,
				GoalBonded:            goalBonded.Neg(),
			},
			true,
		},
		{
			"invalid - goal bonded - goal bonded is greater than 1",
			Params{
				MintDenom:             "acbo",
				InflationDistribution: validInflationDistribution,
				EnableInflation:       true,
				InflationRateChange:   inflationRateChange,
				InflationMax:          inflationMax,
				InflationMin:          inflationMin,
				GoalBonded:            sdk.NewDecWithPrec(110, 2),
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
