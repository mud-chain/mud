package epochs

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"

	epochskeeper "github.com/evmos/evmos/v12/x/epochs/keeper"

	"github.com/evmos/evmos/v12/x/evm/types"
)

type Contract struct {
	ctx          sdk.Context
	epochsKeeper epochskeeper.Keeper
}

func NewPrecompiledContract(ctx sdk.Context, epochsKeeper epochskeeper.Keeper) *Contract {
	return &Contract{
		ctx:          ctx,
		epochsKeeper: epochsKeeper,
	}
}

func (c *Contract) Address() common.Address {
	return epochsAddress
}

func (c *Contract) RequiredGas(input []byte) uint64 {
	method, err := GetMethodByID(input)
	if err != nil {
		return 0
	}

	switch method.Name {
	case EpochInfosMethodName:
		return EpochInfosGas
	case CurrentEpochMethodName:
		return CurrentEpochGas
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
		case EpochInfosMethodName:
			ret, err = c.EpochInfos(ctx, evm, contract, readonly)
		case CurrentEpochMethodName:
			ret, err = c.CurrentEpoch(ctx, evm, contract, readonly)
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
