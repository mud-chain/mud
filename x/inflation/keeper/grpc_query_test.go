package keeper_test

import (
	"fmt"
	"github.com/evmos/evmos/v12/cmd/config"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	evmostypes "github.com/evmos/evmos/v12/types"
	"github.com/evmos/evmos/v12/x/inflation/types"
)

func (suite *KeeperTestSuite) TestPeriod() { //nolint:dupl
	var (
		req    *types.QueryPeriodRequest
		expRes *types.QueryPeriodResponse
	)

	testCases := []struct {
		name     string
		malleate func()
		expPass  bool
	}{
		{
			"default period",
			func() {
				req = &types.QueryPeriodRequest{}
				expRes = &types.QueryPeriodResponse{}
			},
			true,
		},
		{
			"set period",
			func() {
				period := uint64(9)
				suite.app.InflationKeeper.SetPeriod(suite.ctx, period)
				suite.Commit()

				req = &types.QueryPeriodRequest{}
				expRes = &types.QueryPeriodResponse{Period: period}
			},
			true,
		},
	}
	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.name), func() {
			suite.SetupTest() // reset

			ctx := sdk.WrapSDKContext(suite.ctx)
			tc.malleate()

			res, err := suite.queryClient.Period(ctx, req)
			if tc.expPass {
				suite.Require().NoError(err)
				suite.Require().Equal(expRes, res)
			} else {
				suite.Require().Error(err)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestEpochMintProvision() {
	var (
		req    *types.QueryEpochProvisionRequest
		expRes *types.QueryEpochProvisionResponse
	)

	testCases := []struct {
		name     string
		malleate func()
		expPass  bool
	}{
		{
			"default epochMintProvision",
			func() {
				req = &types.QueryEpochProvisionRequest{}
				expRes = &types.QueryEpochProvisionResponse{
					Mint: sdk.Coin{
						Denom:  types.DefaultInflationDenom,
						Amount: sdk.MustNewDecFromStr("102739726027397260273972").TruncateInt(),
					},
					Reward: sdk.Coin{
						Denom:  types.DefaultInflationDenom,
						Amount: sdk.MustNewDecFromStr("547945205479452").TruncateInt(),
					},
					Burn: sdk.Coin{
						Denom:  types.DefaultInflationDenom,
						Amount: sdk.MustNewDecFromStr("102739725479452054794520").TruncateInt(),
					},
				}
			},
			true,
		},
	}
	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.name), func() {
			suite.SetupTest() // reset

			ctx := sdk.WrapSDKContext(suite.ctx)
			tc.malleate()

			res, err := suite.queryClient.EpochProvision(ctx, req)
			if tc.expPass {
				suite.Require().NoError(err)
				suite.Require().Equal(expRes, res)
			} else {
				suite.Require().Error(err)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestSkippedEpochs() { //nolint:dupl
	var (
		req    *types.QuerySkippedEpochsRequest
		expRes *types.QuerySkippedEpochsResponse
	)

	testCases := []struct {
		name     string
		malleate func()
		expPass  bool
	}{
		{
			"default skipped epochs",
			func() {
				req = &types.QuerySkippedEpochsRequest{}
				expRes = &types.QuerySkippedEpochsResponse{}
			},
			true,
		},
		{
			"set skipped epochs",
			func() {
				skippedEpochs := uint64(9)
				suite.app.InflationKeeper.SetSkippedEpochs(suite.ctx, skippedEpochs)
				suite.Commit()

				req = &types.QuerySkippedEpochsRequest{}
				expRes = &types.QuerySkippedEpochsResponse{SkippedEpochs: skippedEpochs}
			},
			true,
		},
	}
	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.name), func() {
			suite.SetupTest() // reset

			ctx := sdk.WrapSDKContext(suite.ctx)
			tc.malleate()

			res, err := suite.queryClient.SkippedEpochs(ctx, req)
			if tc.expPass {
				suite.Require().NoError(err)
				suite.Require().Equal(expRes, res)
			} else {
				suite.Require().Error(err)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestQueryCirculatingSupply() {
	// Team allocation is only set on mainnet
	ctx := sdk.WrapSDKContext(suite.ctx)

	// Mint coins to increase supply
	mintDenom := config.BaseDenom
	mintCoin := sdk.NewCoin(mintDenom, sdk.TokensFromConsensusPower(int64(400_000_000), evmostypes.PowerReduction))
	err := suite.app.InflationKeeper.MintCoins(suite.ctx, mintCoin)
	suite.Require().NoError(err)

	// the total bonded tokens for the 2 accounts initialized on the setup
	bondedAmt := sdk.NewInt64DecCoin(evmostypes.AttoEvmos, 1000100000000000000)

	res, err := suite.queryClient.CirculatingSupply(ctx, &types.QueryCirculatingSupplyRequest{})
	suite.Require().NoError(err)
	suite.Require().Equal(bondedAmt.Amount.Add(sdk.NewDecFromInt(mintCoin.Amount)).Add(sdk.NewDecFromInt(types.DefaultInflationAmount)), res.CirculatingSupply.Amount)
}

func (suite *KeeperTestSuite) TestQueryInflationRate() {
	ctx := sdk.WrapSDKContext(suite.ctx)

	// the total bonded tokens for the 2 accounts initialized on the setup
	bondedAmt := math.NewInt(1000100000000000000)

	// Mint coins to increase supply
	mintDenom := config.BaseDenom
	mintCoin := sdk.NewCoin(mintDenom, sdk.TokensFromConsensusPower(int64(400_000_000), evmostypes.PowerReduction).Sub(bondedAmt))
	err := suite.app.InflationKeeper.MintCoins(suite.ctx, mintCoin)
	suite.Require().NoError(err)

	expInflationRate := sdk.MustNewDecFromStr("20.000000000000000000")
	res, err := suite.queryClient.InflationRate(ctx, &types.QueryInflationRateRequest{})
	suite.Require().NoError(err)
	suite.Require().Equal(expInflationRate, res.InflationRate)
}

func (suite *KeeperTestSuite) TestQueryParams() {
	ctx := sdk.WrapSDKContext(suite.ctx)
	expParams := types.DefaultParams()

	res, err := suite.queryClient.Params(ctx, &types.QueryParamsRequest{})
	suite.Require().NoError(err)
	suite.Require().Equal(expParams, res.Params)
}
