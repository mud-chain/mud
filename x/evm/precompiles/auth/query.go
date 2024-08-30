package auth

import (
	"bytes"
	"encoding/base64"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
	evmostypes "github.com/evmos/evmos/v12/types"
	"github.com/evmos/evmos/v12/utils"
	"github.com/evmos/evmos/v12/x/evm/types"
)

const (
	AccountGas             = 30_000
	AccountsGas            = 50_000
	ModuleAccountByNameGas = 30_000
	ModuleAccountsGas      = 50_000
	AccountAddressByIDGas  = 30_000
	ParamsGas              = 50_000

	AccountMethodName             = "account"
	AccountsMethodName            = "accounts"
	ModuleAccountByNameMethodName = "moduleAccountByName"
	ModuleAccountsMethodName      = "moduleAccounts"
	AccountAddressByIDMethodName  = "accountAddressByID"
	ParamsMethodName              = "params"
)

// Account returns account details based on address.
func (c *Contract) Account(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := MustMethod(AccountMethodName)

	var args AccountArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}
	msg := &authtypes.QueryAccountRequest{
		Address: sdk.AccAddress(args.Addr.Bytes()).String(),
	}

	res, err := c.accountKeeper.Account(ctx, msg)
	if err != nil {
		return nil, err
	}

	return method.Outputs.Pack(OutputsAccount(res.Account))
}

// Accounts returns all the existing accounts
func (c *Contract) Accounts(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := MustMethod(AccountsMethodName)

	var args AccountsArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}
	if bytes.Equal(args.PageRequest.Key, []byte{0}) {
		args.PageRequest.Key = nil
	}

	msg := &authtypes.QueryAccountsRequest{
		Pagination: &query.PageRequest{
			Key:        args.PageRequest.Key,
			Offset:     args.PageRequest.Offset,
			Limit:      args.PageRequest.Limit,
			CountTotal: args.PageRequest.CountTotal,
			Reverse:    args.PageRequest.Reverse,
		},
	}

	res, err := c.accountKeeper.Accounts(ctx, msg)
	if err != nil {
		return nil, err
	}

	var accounts []Account
	for _, account := range res.Accounts {
		accounts = append(accounts, OutputsAccount(account))
	}

	var pageResponse PageResponse
	pageResponse.NextKey = res.Pagination.NextKey
	pageResponse.Total = res.Pagination.Total

	return method.Outputs.Pack(accounts, pageResponse)
}

// ModuleAccountByName returns the module account info by module name
func (c *Contract) ModuleAccountByName(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := MustMethod(ModuleAccountByNameMethodName)

	var args ModuleAccountByNameArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}
	msg := &authtypes.QueryModuleAccountByNameRequest{
		Name: args.Name,
	}

	res, err := c.accountKeeper.ModuleAccountByName(ctx, msg)
	if err != nil {
		return nil, err
	}

	return method.Outputs.Pack(OutputsModuleAccount(res.Account))
}

// ModuleAccounts returns all the existing module accounts.
func (c *Contract) ModuleAccounts(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := MustMethod(ModuleAccountsMethodName)

	msg := &authtypes.QueryModuleAccountsRequest{}

	res, err := c.accountKeeper.ModuleAccounts(ctx, msg)
	if err != nil {
		return nil, err
	}

	var accounts []ModuleAccount
	for _, account := range res.Accounts {
		accounts = append(accounts, OutputsModuleAccount(account))
	}

	return method.Outputs.Pack(accounts)
}

// AccountAddressByID returns account address based on account number.
func (c *Contract) AccountAddressByID(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := MustMethod(AccountAddressByIDMethodName)

	var args AccountAddressByIDArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}

	msg := &authtypes.QueryAccountAddressByIDRequest{
		Id: args.Id,
	}

	res, err := c.accountKeeper.AccountAddressByID(ctx, msg)
	if err != nil {
		return nil, err
	}

	return method.Outputs.Pack(utils.AccAddressMustToHexAddress(res.AccountAddress))
}

// Params queries the parameters of x/auth module.
func (c *Contract) Params(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := MustMethod(ParamsMethodName)

	msg := &authtypes.QueryParamsRequest{}

	res, err := c.accountKeeper.Params(ctx, msg)
	if err != nil {
		return nil, err
	}

	params := Params{
		MaxMemoCharacters:      res.Params.MaxMemoCharacters,
		TxSigLimit:             res.Params.TxSigLimit,
		TxSizeCostPerByte:      res.Params.TxSizeCostPerByte,
		SigVerifyCostEd25519:   res.Params.SigVerifyCostED25519,
		SigVerifyCostSecp256k1: res.Params.SigVerifyCostSecp256k1,
	}

	return method.Outputs.Pack(params)
}

func OutputsAccount(account *codectypes.Any) Account {
	baseAccount, ok := account.GetCachedValue().(*authtypes.BaseAccount)
	if ok {
		return Account{
			Addr:          common.BytesToAddress(baseAccount.GetAddress()),
			PubKey:        FormatPubkey(baseAccount.PubKey),
			AccountNumber: baseAccount.AccountNumber,
			Sequence:      baseAccount.Sequence,
		}
	}

	moduleAccount, ok := account.GetCachedValue().(*authtypes.ModuleAccount)
	if ok {
		return Account{
			Addr:          common.BytesToAddress(moduleAccount.GetAddress()),
			PubKey:        FormatPubkey(moduleAccount.PubKey),
			AccountNumber: moduleAccount.AccountNumber,
			Sequence:      moduleAccount.Sequence,
		}
	}

	ethAccount, ok := account.GetCachedValue().(*evmostypes.EthAccount)
	if ok {
		return Account{
			Addr:          common.BytesToAddress(ethAccount.GetAddress()),
			PubKey:        FormatPubkey(ethAccount.PubKey),
			AccountNumber: ethAccount.AccountNumber,
			Sequence:      ethAccount.Sequence,
		}
	}

	return Account{
		Addr:          common.Address{},
		PubKey:        "",
		AccountNumber: 0,
		Sequence:      0,
	}
}

func OutputsModuleAccount(account *codectypes.Any) ModuleAccount {
	moduleAccount, ok := account.GetCachedValue().(*authtypes.ModuleAccount)
	if ok {
		return ModuleAccount{
			Addr:          common.BytesToAddress(moduleAccount.GetAddress()),
			PubKey:        FormatPubkey(moduleAccount.PubKey),
			AccountNumber: moduleAccount.AccountNumber,
			Sequence:      moduleAccount.Sequence,
			Name:          moduleAccount.Name,
			Permissions:   moduleAccount.Permissions,
		}
	}

	return ModuleAccount{
		Addr:          common.Address{},
		PubKey:        "",
		AccountNumber: 0,
		Sequence:      0,
		Name:          "",
		Permissions:   nil,
	}
}

// FormatPubkey format Pubkey into a base64 string
func FormatPubkey(pubkey *codectypes.Any) string {
	if pubkey == nil {
		return ""
	}
	pubKey, ok := pubkey.GetCachedValue().(cryptotypes.PubKey)
	if ok {
		return base64.StdEncoding.EncodeToString(pubKey.Bytes())
	}
	return pubkey.String()
}
