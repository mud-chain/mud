package staking

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"

	"github.com/evmos/evmos/v12/types"
)

var (
	stakingAddress = common.HexToAddress(types.StakingAddress)
	stakingABI     = types.MustABIJson(IStakingMetaData.ABI)
)

func GetAddress() common.Address {
	return stakingAddress
}

func GetMethod(name string) (abi.Method, error) {
	method := stakingABI.Methods[name]
	if method.ID == nil {
		return abi.Method{}, fmt.Errorf("method %s is not exist", name)
	}
	return method, nil
}

func GetMethodByID(input []byte) (abi.Method, error) {
	if len(input) <= 4 {
		return abi.Method{}, fmt.Errorf("input length %d is too short", len(input))
	}
	for _, method := range stakingABI.Methods {
		if bytes.Equal(input[:4], method.ID) {
			return method, nil
		}
	}
	return abi.Method{}, fmt.Errorf("method id %s is not exist", string(input[:4]))
}

func MustMethod(name string) abi.Method {
	method, err := GetMethod(name)
	if err != nil {
		panic(err)
	}
	return method
}

func GetEvent(name string) (abi.Event, error) {
	event := stakingABI.Events[name]
	if event.ID == (common.Hash{}) {
		return abi.Event{}, fmt.Errorf("event %s is not exist", name)
	}
	return event, nil
}

func MustEvent(name string) abi.Event {
	event, err := GetEvent(name)
	if err != nil {
		panic(err)
	}
	return event
}

type DescriptionJson = Description
type CommissionRatesJson = CommissionRates

type CreateValidatorArgs struct {
	Description       DescriptionJson     `abi:"description"`
	Commission        CommissionRatesJson `abi:"commission"`
	MinSelfDelegation *big.Int            `abi:"minSelfDelegation"`
	Pubkey            string              `abi:"pubkey"`
	Value             *big.Int            `abi:"value"`
}

// Validate validates the args
func (args *CreateValidatorArgs) Validate() error {
	pubkeyBytes, err := base64.StdEncoding.DecodeString(args.Pubkey)
	if err != nil {
		return err
	}

	if len(pubkeyBytes) != ed25519.PubKeySize {
		return fmt.Errorf("pubkey %s is invalid, len is: %d, expected len: %d", hex.EncodeToString(pubkeyBytes), len(pubkeyBytes), ed25519.PubKeySize)
	}

	if args.Value == nil || args.Value.Sign() <= 0 {
		return errors.New("invalid value")
	}
	return nil
}

// GetPubkey returns the validator pubkey
func (args *CreateValidatorArgs) GetPubkey() *codectypes.Any {
	pubkeyBytes, err := base64.StdEncoding.DecodeString(args.Pubkey)
	if err != nil {
		return nil
	}

	var ed25519pk cryptotypes.PubKey = &ed25519.PubKey{Key: pubkeyBytes}
	pubkey, err := codectypes.NewAnyWithValue(ed25519pk)
	if err != nil {
		return nil
	}

	return pubkey
}

type DelegateArgs struct {
	ValidatorAddress common.Address `abi:"validatorAddress"`
	Amount           *big.Int       `abi:"amount"`
}

// Validate validates the args
func (args *DelegateArgs) Validate() error {
	if args.ValidatorAddress == (common.Address{}) {
		return fmt.Errorf("invalid validator address: %s", args.ValidatorAddress)
	}
	if args.Amount == nil || args.Amount.Sign() <= 0 {
		return errors.New("invalid amount")
	}
	return nil
}

// GetValidator returns the validator address, caller must ensure the validator address is valid
func (args *DelegateArgs) GetValidator() sdk.ValAddress {
	valAddr := sdk.ValAddress(args.ValidatorAddress.Bytes())
	return valAddr
}

type DelegationArgs struct {
	DelegatorAddress common.Address `abi:"delegatorAddress"`
	ValidatorAddress common.Address `abi:"validatorAddress"`
}

// Validate validates the args
func (args *DelegationArgs) Validate() error {
	if args.DelegatorAddress == (common.Address{}) {
		return fmt.Errorf("invalid delegator address: %s", args.DelegatorAddress)
	}
	if args.ValidatorAddress == (common.Address{}) {
		return fmt.Errorf("invalid validator address: %s", args.ValidatorAddress)
	}

	return nil
}

// GetValidator returns the validator address, caller must ensure the validator address is valid
func (args *DelegationArgs) GetValidator() sdk.ValAddress {
	valAddr := sdk.ValAddress(args.ValidatorAddress.Bytes())
	return valAddr
}

// GetDelegator returns the Delegator address, caller must ensure the delegator address is valid
func (args *DelegationArgs) GetDelegator() sdk.AccAddress {
	accAddr := sdk.AccAddress(args.DelegatorAddress.Bytes())
	return accAddr
}

type UndelegateArgs struct {
	ValidatorAddress common.Address `abi:"validatorAddress"`
	Amount           *big.Int       `abi:"amount"`
}

// Validate validates the args
func (args *UndelegateArgs) Validate() error {
	if args.ValidatorAddress == (common.Address{}) {
		return fmt.Errorf("invalid validator address: %s", args.ValidatorAddress)
	}
	if args.Amount == nil || args.Amount.Sign() <= 0 {
		return errors.New("invalid amount")
	}
	return nil
}

// GetValidator returns the validator address, caller must ensure the validator address is valid
func (args *UndelegateArgs) GetValidator() sdk.ValAddress {
	valAddr := sdk.ValAddress(args.ValidatorAddress.Bytes())
	return valAddr
}

type RedelegateArgs struct {
	ValidatorSrcAddress common.Address `abi:"validatorSrcAddress"`
	ValidatorDstAddress common.Address `abi:"validatorDstAddress"`
	Amount              *big.Int       `abi:"amount"`
}

// Validate validates the args
func (args *RedelegateArgs) Validate() error {
	if args.ValidatorSrcAddress == (common.Address{}) {
		return fmt.Errorf("invalid src validator address: %s", args.ValidatorSrcAddress)
	}
	if args.ValidatorDstAddress == (common.Address{}) {
		return fmt.Errorf("invalid dst validator address: %s", args.ValidatorDstAddress)
	}
	if args.Amount == nil || args.Amount.Sign() <= 0 {
		return errors.New("invalid amount")
	}
	return nil
}

// GetSrcValidator returns the validator src address, caller must ensure the validator address is valid
func (args *RedelegateArgs) GetSrcValidator() sdk.ValAddress {
	valAddr := sdk.ValAddress(args.ValidatorSrcAddress.Bytes())
	return valAddr
}

// GetDstValidator returns the validator dest address, caller must ensure the validator address is valid
func (args *RedelegateArgs) GetDstValidator() sdk.ValAddress {
	valAddr := sdk.ValAddress(args.ValidatorDstAddress.Bytes())
	return valAddr
}

type CancelUnbondingDelegationArgs struct {
	ValidatorAddress common.Address `abi:"validatorAddress"`
	Amount           *big.Int       `abi:"amount"`
	CreationHeight   *big.Int       `abi:"creationHeight"`
}

// Validate validates the args
func (args *CancelUnbondingDelegationArgs) Validate() error {
	if args.ValidatorAddress == (common.Address{}) {
		return fmt.Errorf("invalid validator address: %s", args.ValidatorAddress)
	}

	if args.Amount == nil || args.Amount.Sign() <= 0 {
		return errors.New("invalid amount")
	}

	if args.CreationHeight == nil || args.CreationHeight.Sign() <= 0 || !args.CreationHeight.IsInt64() {
		return errors.New("invalid creation height")
	}

	return nil
}

// GetValidator returns the validator address
func (args *CancelUnbondingDelegationArgs) GetValidator() sdk.ValAddress {
	valAddr := sdk.ValAddress(args.ValidatorAddress.Bytes())
	return valAddr
}

// GetCreationHeight returns the creation height
func (args *CancelUnbondingDelegationArgs) GetCreationHeight() int64 {
	return args.CreationHeight.Int64()
}
