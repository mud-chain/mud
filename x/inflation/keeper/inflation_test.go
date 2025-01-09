package keeper_test

import (
	sdkmath "cosmossdk.io/math"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	evmostypes "github.com/evmos/evmos/v12/types"
	"github.com/evmos/evmos/v12/x/inflation/types"
)

func (suite *KeeperTestSuite) TestMintAndAllocateInflation() {
	testCases := []struct {
		name                string
		mintCoin            sdk.Coin
		malleate            func()
		expStakingRewardAmt sdk.Coin
		expCommunityPoolAmt sdk.DecCoins
		expPass             bool
	}{
		{
			"pass",
			sdk.NewCoin(denomMint, sdk.NewInt(1_000_000)),
			func() {},
			sdk.NewCoin(denomMint, sdk.NewInt(800_000)),
			sdk.NewDecCoins(sdk.NewDecCoin(denomMint, sdk.NewInt(200_000))),
			true,
		},
		{
			"pass - no coins minted ",
			sdk.NewCoin(denomMint, sdk.ZeroInt()),
			func() {},
			sdk.NewCoin(denomMint, sdk.ZeroInt()),
			sdk.DecCoins(nil),
			true,
		},
	}
	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.name), func() {
			suite.SetupTest() // reset

			err := suite.app.InflationKeeper.MintCoins(suite.ctx, tc.mintCoin)
			suite.Require().NoError(err)

			tc.malleate()

			bondedTokens := suite.app.StakingKeeper.TotalBondedTokens(suite.ctx)
			_, _, err = suite.app.InflationKeeper.MintAndAllocateInflation(suite.ctx, bondedTokens, tc.mintCoin, types.DefaultParams())

			// Get balances
			balanceModule := suite.app.BankKeeper.GetBalance(
				suite.ctx,
				suite.app.AccountKeeper.GetModuleAddress(types.ModuleName),
				denomMint,
			)

			feeCollector := suite.app.AccountKeeper.GetModuleAddress(authtypes.FeeCollectorName)
			balanceStakingRewards := suite.app.BankKeeper.GetBalance(
				suite.ctx,
				feeCollector,
				denomMint,
			)

			balanceCommunityPool := suite.app.DistrKeeper.GetFeePoolCommunityCoins(suite.ctx)

			if tc.expPass {
				suite.Require().NoError(err, tc.name)
				suite.Require().True(balanceModule.IsZero())
				suite.Require().Equal(tc.expStakingRewardAmt, balanceStakingRewards)
				suite.Require().Equal(tc.expCommunityPoolAmt, balanceCommunityPool)
			} else {
				suite.Require().Error(err)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestGetCirculatingSupplyAndInflationRate() {
	// the total bonded tokens for the 2 accounts initialized on the setup
	bondedAmt := sdkmath.NewInt(1000100000000000000)
	bondedCoins := sdk.NewDecCoin(evmostypes.AttoEvmos, bondedAmt)

	testCases := []struct {
		name             string
		bankSupply       sdkmath.Int
		malleate         func()
		expInflationRate sdk.Dec
	}{
		{
			"no epochs per period",
			sdk.TokensFromConsensusPower(400_000_000, evmostypes.PowerReduction).Sub(bondedAmt),
			func() {
				suite.app.InflationKeeper.SetEpochsPerPeriod(suite.ctx, 0)
			},
			sdk.MustNewDecFromStr("0.200000000000000000"),
		},
		{
			"high supply",
			sdk.TokensFromConsensusPower(800_000_000, evmostypes.PowerReduction).Sub(bondedAmt),
			func() {},
			sdk.MustNewDecFromStr("0.200000000000000000"),
		},
		{
			"low supply",
			sdk.TokensFromConsensusPower(400_000_000, evmostypes.PowerReduction).Sub(bondedAmt),
			func() {},
			sdk.MustNewDecFromStr("0.200000000000000000"),
		},
		{
			"zero circulating supply",
			sdk.TokensFromConsensusPower(200_000_000, evmostypes.PowerReduction).Sub(bondedAmt),
			func() {},
			sdk.MustNewDecFromStr("0.200000000000000000"),
		},
	}
	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.name), func() {
			suite.SetupTest() // reset

			// Team allocation is only set on mainnet
			suite.ctx = suite.ctx.WithChainID("evmos_9001-1")
			tc.malleate()

			// Mint coins to increase supply
			coin := sdk.NewCoin(
				types.DefaultInflationDenom,
				tc.bankSupply,
			)
			decCoin := sdk.NewDecCoinFromCoin(coin)
			err := suite.app.InflationKeeper.MintCoins(suite.ctx, coin)
			suite.Require().NoError(err)

			circulatingSupply := s.app.InflationKeeper.GetCirculatingSupply(suite.ctx, types.DefaultInflationDenom)
			suite.Require().Equal(decCoin.Add(bondedCoins).Amount, circulatingSupply)

			params := suite.app.InflationKeeper.GetParams(suite.ctx)
			inflationAmount := suite.app.InflationKeeper.GetInflationAmount(suite.ctx).Amount
			bondedTokens := suite.app.StakingKeeper.TotalBondedTokens(suite.ctx)

			inflationRate := types.InflationRate(params, inflationAmount, bondedTokens)
			suite.Require().Equal(tc.expInflationRate, inflationRate)
		})
	}
}

func (suite *KeeperTestSuite) TestBondedRatio() {
	testCases := []struct {
		name         string
		isMainnet    bool
		malleate     func()
		expBondRatio sdk.Dec
	}{
		{
			"is mainnet",
			true,
			func() {},
			sdk.MustNewDecFromStr("0.999900009999000099"),
		},
		{
			"not mainnet",
			false,
			func() {},
			sdk.MustNewDecFromStr("0.999900009999000099"),
		},
	}
	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.name), func() {
			suite.SetupTest() // reset

			// Team allocation is only set on mainnet
			if tc.isMainnet {
				suite.ctx = suite.ctx.WithChainID("evmos_9001-1")
			} else {
				suite.ctx = suite.ctx.WithChainID("evmos_9999-666")
			}
			tc.malleate()

			bondRatio := suite.app.InflationKeeper.BondedRatio(suite.ctx)
			suite.Require().Equal(tc.expBondRatio, bondRatio)
		})
	}
}
