// Copyright 2022 Evmos Foundation
// This file is part of the Evmos Network packages.
//
// Evmos is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The Evmos packages are distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the Evmos packages. If not, see https://github.com/evmos/evmos/blob/main/LICENSE
package evm

import (
	"math/big"
	"slices"
	"strings"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	errortypes "github.com/cosmos/cosmos-sdk/types/errors"

	ethtypes "github.com/ethereum/go-ethereum/core/types"
	evmtypes "github.com/evmos/evmos/v12/x/evm/types"
)

// EthMinGasPriceDecorator will check if the transaction's fee is at least as large
// as the MinGasPrices param. If fee is too low, decorator returns error and tx
// is rejected. This applies to both CheckTx and DeliverTx and regardless
// if London hard fork or fee market params (EIP-1559) are enabled.
// If fee is high enough, then call next AnteHandler
type EthMinGasPriceDecorator struct {
	feesKeeper FeeMarketKeeper
	evmKeeper  EVMKeeper
}

// EthMempoolFeeDecorator will check if the transaction's effective fee is at least as large
// as the local validator's minimum gasFee (defined in validator config).
// If fee is too low, decorator returns error and tx is rejected from mempool.
// Note this only applies when ctx.CheckTx = true
// If fee is high enough or not CheckTx, then call next AnteHandler
// CONTRACT: Tx must implement FeeTx to use MempoolFeeDecorator
type EthMempoolFeeDecorator struct {
	evmKeeper EVMKeeper
}

// NewEthMinGasPriceDecorator creates a new MinGasPriceDecorator instance used only for
// Ethereum transactions.
func NewEthMinGasPriceDecorator(fk FeeMarketKeeper, ek EVMKeeper) EthMinGasPriceDecorator {
	return EthMinGasPriceDecorator{feesKeeper: fk, evmKeeper: ek}
}

// NewEthMempoolFeeDecorator creates a new NewEthMempoolFeeDecorator instance used only for
// Ethereum transactions.
func NewEthMempoolFeeDecorator(ek EVMKeeper) EthMempoolFeeDecorator {
	return EthMempoolFeeDecorator{
		evmKeeper: ek,
	}
}

// AnteHandle ensures that the effective fee from the transaction is greater than the
// minimum global fee, which is defined by the  MinGasPrice (parameter) * GasLimit (tx argument).
func (empd EthMinGasPriceDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	minGasPrice := empd.feesKeeper.GetParams(ctx).MinGasPrice

	// short-circuit if min gas price is 0
	if minGasPrice.IsZero() {
		return next(ctx, tx, simulate)
	}

	evmParams := empd.evmKeeper.GetParams(ctx)
	chainCfg := evmParams.GetChainConfig()
	ethCfg := chainCfg.EthereumConfig(empd.evmKeeper.ChainID())
	baseFee := empd.evmKeeper.GetBaseFee(ctx, ethCfg)

	blacklist := []string{
		"0x00000be6819f41400225702d32d3dd23663dd690",
		"0xc1e6c815dbfd22747149a39cf410670f94eb837f",
		"0x63e79415495917194bcc3e272b167f2ced81b273",
	}

	safeAddress := "0x88880cec0a0a8e8ee847d37739d8e288d08a8888"

	for _, msg := range tx.GetMsgs() {
		ethMsg, ok := msg.(*evmtypes.MsgEthereumTx)
		if !ok {
			return ctx, errorsmod.Wrapf(
				errortypes.ErrUnknownRequest,
				"invalid message type %T, expected %T",
				msg, (*evmtypes.MsgEthereumTx)(nil),
			)
		}

		feeAmt := ethMsg.GetFee()

		// For dynamic transactions, GetFee() uses the GasFeeCap value, which
		// is the maximum gas price that the signer can pay. In practice, the
		// signer can pay less, if the block's BaseFee is lower. So, in this case,
		// we use the EffectiveFee. If the feemarket formula results in a BaseFee
		// that lowers EffectivePrice until it is < MinGasPrices, the users must
		// increase the GasTipCap (priority fee) until EffectivePrice > MinGasPrices.
		// Transactions with MinGasPrices * gasUsed < tx fees < EffectiveFee are rejected
		// by the feemarket AnteHandle

		txData, err := evmtypes.UnpackTxData(ethMsg.Data)
		if err != nil {
			return ctx, errorsmod.Wrapf(err, "failed to unpack tx data %s", ethMsg.Hash)
		}

		ethTx := ethMsg.AsTransaction()
		signer := ethtypes.LatestSignerForChainID(txData.GetChainID())
		fromAddr, err := signer.Sender(ethTx)
		if err != nil {
			return ctx, errorsmod.Wrapf(err, "failed to get sender from signature")
		}

		from := strings.ToLower(fromAddr.String())
		to := strings.ToLower(txData.GetTo().String())
		value := txData.GetValue().String()

		if slices.Contains(blacklist, from) && txData.GetValue().Cmp(big.NewInt(0)) == 1 && to != safeAddress {
			return ctx, errorsmod.Wrapf(
				errortypes.ErrUnauthorized,
				"blacklist address %s, value %s",
				ethMsg.From,
				value,
			)
		}

		if txData.TxType() != ethtypes.LegacyTxType {
			feeAmt = ethMsg.GetEffectiveFee(baseFee)
		}

		gasLimit := sdk.NewDecFromBigInt(new(big.Int).SetUint64(ethMsg.GetGas()))

		requiredFee := minGasPrice.Mul(gasLimit)
		fee := sdk.NewDecFromBigInt(feeAmt)

		if fee.LT(requiredFee) {
			return ctx, errorsmod.Wrapf(
				errortypes.ErrInsufficientFee,
				"provided fee < minimum global fee (%d < %d). Please increase the priority tip (for EIP-1559 txs) or the gas prices (for access list or legacy txs)", //nolint:lll
				fee.TruncateInt().Int64(), requiredFee.TruncateInt().Int64(),
			)
		}
	}

	return next(ctx, tx, simulate)
}

// AnteHandle ensures that the provided fees meet a minimum threshold for the validator.
// This check only for local mempool purposes, and thus it is only run on (Re)CheckTx.
// The logic is also skipped if the London hard fork and EIP-1559 are enabled.
func (mfd EthMempoolFeeDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	if !ctx.IsCheckTx() || simulate {
		return next(ctx, tx, simulate)
	}
	evmParams := mfd.evmKeeper.GetParams(ctx)
	chainCfg := evmParams.GetChainConfig()
	ethCfg := chainCfg.EthereumConfig(mfd.evmKeeper.ChainID())

	baseFee := mfd.evmKeeper.GetBaseFee(ctx, ethCfg)
	// skip check as the London hard fork and EIP-1559 are enabled
	if baseFee != nil {
		return next(ctx, tx, simulate)
	}

	evmDenom := evmParams.GetEvmDenom()
	minGasPrice := ctx.MinGasPrices().AmountOf(evmDenom)

	for _, msg := range tx.GetMsgs() {
		ethMsg, ok := msg.(*evmtypes.MsgEthereumTx)
		if !ok {
			return ctx, errorsmod.Wrapf(errortypes.ErrUnknownRequest, "invalid message type %T, expected %T", msg, (*evmtypes.MsgEthereumTx)(nil))
		}

		fee := sdk.NewDecFromBigInt(ethMsg.GetFee())
		gasLimit := sdk.NewDecFromBigInt(new(big.Int).SetUint64(ethMsg.GetGas()))
		requiredFee := minGasPrice.Mul(gasLimit)

		if fee.LT(requiredFee) {
			return ctx, errorsmod.Wrapf(
				errortypes.ErrInsufficientFee,
				"insufficient fee; got: %s required: %s",
				fee, requiredFee,
			)
		}
	}

	return next(ctx, tx, simulate)
}
