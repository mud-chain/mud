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

package types

import (
	"errors"
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	evm "github.com/evmos/evmos/v12/x/evm/types"
)

var ParamsKey = []byte("Params")

var (
	DefaultInflationDenom        = evm.DefaultEVMDenom
	DefaultEnableInflation       = true
	DefaultInflationDistribution = InflationDistribution{
		StakingRewards: sdk.NewDecWithPrec(8, 1), // 0.8
		CommunityPool:  sdk.NewDecWithPrec(2, 1), // 0.2
	}
)

func NewParams(
	inflationDistribution InflationDistribution,
	enableInflation bool,
	inflationMax sdk.Dec,
	inflationDecay sdk.Dec,
) Params {
	return Params{
		InflationDistribution: inflationDistribution,
		EnableInflation:       enableInflation,
		InflationMax:          inflationMax,
		InflationDecay:        inflationDecay,
	}
}

// default minting module parameters
func DefaultParams() Params {
	return Params{
		InflationDistribution: DefaultInflationDistribution,
		EnableInflation:       DefaultEnableInflation,
		InflationMax:          sdk.NewDecWithPrec(20, 2),
		InflationDecay:        sdk.NewDecWithPrec(8, 1),
	}
}

func validateMintDenom(i interface{}) error {
	v, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if strings.TrimSpace(v) == "" {
		return errors.New("mint denom cannot be blank")
	}

	return sdk.ValidateDenom(v)
}

func validateInflationDistribution(i interface{}) error {
	v, ok := i.(InflationDistribution)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.StakingRewards.IsNegative() {
		return errors.New("staking distribution ratio must not be negative")
	}

	if v.CommunityPool.IsNegative() {
		return errors.New("community pool distribution ratio must not be negative")
	}

	totalProportions := v.StakingRewards.Add(v.CommunityPool)
	if !totalProportions.Equal(sdk.NewDec(1)) {
		return errors.New("total distributions ratio should be 1")
	}

	return nil
}

func validateBool(i interface{}) error {
	_, ok := i.(bool)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	return nil
}

func validateInflationMax(i interface{}) error {
	v, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNegative() {
		return fmt.Errorf("max inflation cannot be negative: %s", v)
	}
	if v.GT(sdk.OneDec()) {
		return fmt.Errorf("max inflation too large: %s", v)
	}

	return nil
}

func (p Params) Validate() error {
	if err := validateInflationDistribution(p.InflationDistribution); err != nil {
		return err
	}
	if err := validateInflationMax(p.InflationMax); err != nil {
		return err
	}
	return validateBool(p.EnableInflation)
}
