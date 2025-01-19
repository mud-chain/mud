package inflation

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/core/vm"
	inflationtypes "github.com/evmos/evmos/v12/x/inflation/types"
)

const (
	PeriodGas            = 30_000
	EpochProvisionGas    = 50_000
	SkippedEpochsGas     = 30_000
	CirculatingSupplyGas = 50_000
	InflationRateGas     = 30_000
	ParamsGas            = 50_000

	PeriodMethodName            = "period"
	EpochProvisionMethodName    = "epochProvision"
	SkippedEpochsMethodName     = "skippedEpochs"
	CirculatingSupplyMethodName = "circulatingSupply"
	InflationRateMethodName     = "inflationRate"
	ParamsMethodName            = "params"
)

// Period returns the current period
func (c *Contract) Period(ctx sdk.Context, _ *vm.EVM, _ *vm.Contract, _ bool) ([]byte, error) {
	method := MustMethod(PeriodMethodName)

	msg := &inflationtypes.QueryPeriodRequest{}

	res, err := c.inflationKeeper.Period(ctx, msg)
	if err != nil {
		return nil, err
	}

	return method.Outputs.Pack(res.Period)
}

// EpochProvision retrieves current minting epoch provision value.
func (c *Contract) EpochProvision(ctx sdk.Context, _ *vm.EVM, _ *vm.Contract, _ bool) ([]byte, error) {
	method := MustMethod(EpochProvisionMethodName)

	msg := &inflationtypes.QueryEpochProvisionRequest{}

	res, err := c.inflationKeeper.EpochProvision(ctx, msg)
	if err != nil {
		return nil, err
	}

	provision := Provision{
		Mint: Coin{
			Denom:  res.Mint.Denom,
			Amount: res.Mint.Amount.BigInt(),
		},
		Reward: Coin{
			Denom:  res.Mint.Denom,
			Amount: res.Reward.Amount.BigInt(),
		},
		Burn: Coin{
			Denom:  res.Mint.Denom,
			Amount: res.Burn.Amount.BigInt(),
		},
	}

	return method.Outputs.Pack(provision)
}

// SkippedEpochs retrieves the total number of skipped epochs.
func (c *Contract) SkippedEpochs(ctx sdk.Context, _ *vm.EVM, _ *vm.Contract, _ bool) ([]byte, error) {
	method := MustMethod(SkippedEpochsMethodName)

	msg := &inflationtypes.QuerySkippedEpochsRequest{}

	res, err := c.inflationKeeper.SkippedEpochs(ctx, msg)
	if err != nil {
		return nil, err
	}

	return method.Outputs.Pack(res.SkippedEpochs)
}

// CirculatingSupply retrieves the total number of tokens that are in circulation (i.e. excluding unvested tokens).
func (c *Contract) CirculatingSupply(ctx sdk.Context, _ *vm.EVM, _ *vm.Contract, _ bool) ([]byte, error) {
	method := MustMethod(CirculatingSupplyMethodName)

	msg := &inflationtypes.QueryCirculatingSupplyRequest{}

	res, err := c.inflationKeeper.CirculatingSupply(ctx, msg)
	if err != nil {
		return nil, err
	}

	return method.Outputs.Pack(DecCoin{
		Denom:     res.CirculatingSupply.Denom,
		Amount:    res.CirculatingSupply.Amount.BigInt(),
		Precision: uint8(sdk.Precision),
	})
}

// InflationRate retrieves the inflation rate of the current period.
func (c *Contract) InflationRate(ctx sdk.Context, _ *vm.EVM, _ *vm.Contract, _ bool) ([]byte, error) {
	method := MustMethod(InflationRateMethodName)

	msg := &inflationtypes.QueryInflationRateRequest{}

	res, err := c.inflationKeeper.InflationRate(ctx, msg)
	if err != nil {
		return nil, err
	}

	return method.Outputs.Pack(Dec{
		Amount:    res.InflationRate.BigInt(),
		Precision: uint8(sdk.Precision),
	})
}

// Params queries the parameters of x/inflation module.
func (c *Contract) Params(ctx sdk.Context, _ *vm.EVM, _ *vm.Contract, _ bool) ([]byte, error) {
	method := MustMethod(ParamsMethodName)

	msg := &inflationtypes.QueryParamsRequest{}

	res, err := c.inflationKeeper.Params(ctx, msg)
	if err != nil {
		return nil, err
	}

	params := Params{
		StakingRewards: Dec{
			Amount:    res.Params.InflationDistribution.StakingRewards.BigInt(),
			Precision: uint8(sdk.Precision),
		},
		CommunityPool: Dec{
			Amount:    res.Params.InflationDistribution.CommunityPool.BigInt(),
			Precision: uint8(sdk.Precision),
		},
		EnableInflation: res.Params.EnableInflation,
		InflationMax: Dec{
			Amount:    res.Params.InflationMax.BigInt(),
			Precision: uint8(sdk.Precision),
		},
		InflationDecay: Dec{
			Amount:    res.Params.InflationDecay.BigInt(),
			Precision: uint8(sdk.Precision),
		},
	}

	return method.Outputs.Pack(params)
}
