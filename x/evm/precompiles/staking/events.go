package staking

import (
	"github.com/ethereum/go-ethereum/accounts/abi"

	"github.com/evmos/evmos/v12/x/evm/types"
)

var (
	DelegateEvent = abi.NewEvent(
		DelegateEventName,
		DelegateEventName,
		false,
		abi.Arguments{
			abi.Argument{Name: "delegator", Type: types.TypeAddress, Indexed: true},
			abi.Argument{Name: "validator", Type: types.TypeString, Indexed: false},
			abi.Argument{Name: "amount", Type: types.TypeUint256, Indexed: false},
			abi.Argument{Name: "shares", Type: types.TypeUint256, Indexed: false},
		},
	)

	UndelegateEvent = abi.NewEvent(
		UndelegateEventName,
		UndelegateEventName,
		false,
		abi.Arguments{
			abi.Argument{Name: "sender", Type: types.TypeAddress, Indexed: true},
			abi.Argument{Name: "validator", Type: types.TypeString, Indexed: false},
			abi.Argument{Name: "shares", Type: types.TypeUint256, Indexed: false},
			abi.Argument{Name: "amount", Type: types.TypeUint256, Indexed: false},
			abi.Argument{Name: "completionTime", Type: types.TypeUint256, Indexed: false},
		},
	)

	WithdrawEvent = abi.NewEvent(
		WithdrawEventName,
		WithdrawEventName,
		false,
		abi.Arguments{
			abi.Argument{Name: "sender", Type: types.TypeAddress, Indexed: true},
			abi.Argument{Name: "validator", Type: types.TypeString, Indexed: false},
			abi.Argument{Name: "reward", Type: types.TypeUint256, Indexed: false},
		},
	)

	RedelegateEvent = abi.NewEvent(
		RedelegateEventName,
		RedelegateEventName,
		false,
		abi.Arguments{
			abi.Argument{Name: "sender", Type: types.TypeAddress, Indexed: true},
			abi.Argument{Name: "valSrc", Type: types.TypeString, Indexed: false},
			abi.Argument{Name: "valDst", Type: types.TypeString, Indexed: false},
			abi.Argument{Name: "shares", Type: types.TypeUint256, Indexed: false},
			abi.Argument{Name: "amount", Type: types.TypeUint256, Indexed: false},
			abi.Argument{Name: "completionTime", Type: types.TypeUint256, Indexed: false},
		},
	)
)
