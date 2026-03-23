package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (suite *KeeperTestSuite) TestInitGenesis() {
	// check calculated epochMintProvision at genesis
	epochMintProvision := suite.app.InflationKeeper.GetEpochMintProvision(suite.ctx)
	expMintProvision := sdk.MustNewDecFromStr("102739726027397260273972.000000000000000000")
	suite.Require().Equal(expMintProvision, sdk.NewDecFromInt(epochMintProvision.Amount))
}
