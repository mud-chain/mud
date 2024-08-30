// SPDX-License-Identifier: Apache-2.0

pragma solidity ^0.8.0;

import "../common/Types.sol";

/// @dev The IInflation contract's address.
address constant INFLATION_PRECOMPILE_ADDRESS = 0x0000000000000000000000000000000000001008;

/// @dev The IInflation contract's instance.
IInflation constant INFLATION_CONTRACT = IInflation(INFLATION_PRECOMPILE_ADDRESS);

/**
 * @dev Params defines the parameters for the auth module.
 */
struct Params {
    string mintDenom;
    Dec stakingRewards;
    Dec communityPool;
    bool enableInflation;
    Dec inflationRateChange;
    Dec inflationMax;
    Dec inflationMin;
    Dec goalBonded;
}

interface IInflation {
    /**
     * @dev period retrieves current period.
     */
    function period() external view returns (uint64 period);

    /**
     * @dev epochMintProvision retrieves current minting epoch provision value.
     */
    function epochMintProvision() external view returns (DecCoin memory epochMintProvision);

    /**
     * @dev skippedEpochs retrieves the total number of skipped epochs.
     */
    function skippedEpochs() external view returns (uint64 skippedEpochs);

    /**
     * @dev circulatingSupply retrieves the total number of tokens that are in circulation (i.e. excluding unvested tokens).
     */
    function circulatingSupply() external view returns (DecCoin memory circulatingSupply);

    /**
     * @dev inflationRate retrieves the inflation rate of the current period.
     */
    function inflationRate() external view returns (Dec memory inflationRate);

    /**
     * @dev params queries the parameters of x/inflation module.
     */
    function params() external view returns (Params memory params);
}
