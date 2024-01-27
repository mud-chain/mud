package staking

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/evmos/evmos/v12/types"
	"github.com/evmos/evmos/v12/x/evm/precompiles/contracts"
)

const (
	DelegateGas          = 40_000 // 98000 - 160000 // 165000
	UndelegateGas        = 45_000 // 94000 - 163000 // 172000
	RedelegateGas        = 60_000 // undelegate_gas+delegate_gas+withdraw_gas*2
	WithdrawGas          = 30_000 // 94000 // 120000
	DelegationGas        = 30_000 // 98000
	DelegationRewardsGas = 30_000 // 94000

	DelegateMethodName          = "delegate"
	UndelegateMethodName        = "undelegate"
	RedelegateMethodName        = "redelegate"
	WithdrawMethodName          = "withdraw"
	DelegationMethodName        = "delegation"
	DelegationRewardsMethodName = "delegationRewards"

	DelegateEventName   = "Delegate"
	UndelegateEventName = "Undelegate"
	RedelegateEventName = "Redelegate"
	WithdrawEventName   = "Withdraw"
)

var (
	stakingAddress = common.HexToAddress(types.StakingAddress)
	stakingABI     = types.MustABIJson(contracts.IStakingMetaData.ABI)
)

func GetAddress() common.Address {
	return stakingAddress
}

func GetABI() abi.ABI {
	return stakingABI
}
