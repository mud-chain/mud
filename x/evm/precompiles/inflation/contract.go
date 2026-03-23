package inflation

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"

	"github.com/evmos/evmos/v12/x/evm/types"
	inflationkeeper "github.com/evmos/evmos/v12/x/inflation/keeper"
)

type Contract struct {
	ctx             sdk.Context
	inflationKeeper inflationkeeper.Keeper
}

func NewPrecompiledContract(ctx sdk.Context, inflationKeeper inflationkeeper.Keeper) *Contract {
	return &Contract{
		ctx:             ctx,
		inflationKeeper: inflationKeeper,
	}
}

func (c *Contract) Address() common.Address {
	return inflationAddress
}

func (c *Contract) RequiredGas(input []byte) uint64 {
	method, err := GetMethodByID(input)
	if err != nil {
		return 0
	}

	switch method.Name {
	case PeriodMethodName:
		return PeriodGas
	case EpochProvisionMethodName:
		return EpochProvisionGas
	case SkippedEpochsMethodName:
		return SkippedEpochsGas
	case CirculatingSupplyMethodName:
		return CirculatingSupplyGas
	case InflationRateMethodName:
		return InflationRateGas
	case ParamsMethodName:
		return ParamsGas
	default:
		return 0
	}
}

func (c *Contract) Run(evm *vm.EVM, contract *vm.Contract, readonly bool) (ret []byte, err error) {
	if len(contract.Input) < 4 {
		return types.PackRetError("invalid input")
	}

	ctx, commit := c.ctx.CacheContext()
	snapshot := evm.StateDB.Snapshot()

	method, err := GetMethodByID(contract.Input)
	if err == nil {
		// parse input
		switch method.Name {
		case PeriodMethodName:
			ret, err = c.Period(ctx, evm, contract, readonly)
		case EpochProvisionMethodName:
			ret, err = c.EpochProvision(ctx, evm, contract, readonly)
		case SkippedEpochsMethodName:
			ret, err = c.SkippedEpochs(ctx, evm, contract, readonly)
		case CirculatingSupplyMethodName:
			ret, err = c.CirculatingSupply(ctx, evm, contract, readonly)
		case InflationRateMethodName:
			ret, err = c.InflationRate(ctx, evm, contract, readonly)
		case ParamsMethodName:
			ret, err = c.Params(ctx, evm, contract, readonly)
		default:
			err = fmt.Errorf("method %s is not handle", method.Name)
		}
	}

	if err != nil {
		// revert evm state
		evm.StateDB.RevertToSnapshot(snapshot)
		return types.PackRetError(err.Error())
	}

	// commit and append events
	commit()
	return ret, nil
}
