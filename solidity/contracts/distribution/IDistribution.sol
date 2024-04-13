// SPDX-License-Identifier: Apache-2.0

pragma solidity ^0.8.0;

import "../common/Types.sol";

/**
 * @dev DelegationDelegatorReward represents the properties of a delegator's delegation reward
 */
struct DelegationDelegatorReward {
    address validatorAddress;
    DecCoin[] rewards;
}

interface IDistribution {
    /**
     * @dev setWithdrawAddress SetWithdrawAddress defines a method to change the withdraw address
     * for a delegator (or validator self-delegation).
     */
    function setWithdrawAddress(
        address withdrawAddress
    ) external returns (bool success);

    /**
     * @dev withdrawDelegatorReward defines a method to withdraw rewards of delegator
     * from a single validator.
     */
    function withdrawDelegatorReward(
        address validatorAddress
    ) external returns (Coin[] memory amount);

    /**
     * @dev withdrawValidatorCommission defines a method to withdraw the
     * full commission to the validator address.
     */
    function withdrawValidatorCommission() external returns (Coin[] memory amount);

    /**
     * @dev fundCommunityPool defines a method to allow an account to directly
     * fund the community pool.
     */
    function fundCommunityPool(
        Coin[] memory amount
    ) external returns (bool success);

    /**
     * @dev validatorDistributionInfo queries validator commision and self-delegation rewards for validator
     */
    function validatorDistributionInfo(
        address validatorAddress
    ) external view returns (address operatorAddress, DecCoin[] memory selfBondRewards, DecCoin[] memory commission);

    /**
     * @dev validatorOutstandingRewards queries rewards of a validator address.
     */
    function validatorOutstandingRewards(
        address validatorAddress
    ) external view returns (DecCoin[] memory rewards);

    /**
     * @dev validatorCommission queries accumulated commission for a validator.
     */
    function validatorCommission(
        address validatorAddress
    ) external view returns (DecCoin[] memory commission);

    /**
     * @dev delegationRewards queries the total rewards accrued by a delegation.
     */
    function delegationRewards(
        address delegatorAddress,
        address validatorAddress
    ) external view returns (DecCoin[] memory rewards);

    /**
     * @dev delegationTotalRewards queries the total rewards accrued by a each validator.
     */
    function delegationTotalRewards(
        address delegatorAddress
    ) external view returns (DelegationDelegatorReward[] memory rewards, DecCoin[] memory total);

    /**
     * @dev communityPool queries the community pool coins.
     */
    function communityPool() external view returns (DecCoin[] memory pool);


    /**
     * @dev SetWithdrawAddress defines an Event emitted when a user change the withdraw address
     */
    event SetWithdrawAddress(
        address indexed delegatorAddress,
        address indexed withdrawAddress
    );

    /**
     * @dev WithdrawDelegatorReward defines an Event emitted when withdraw rewards by delegator
     */
    event WithdrawDelegatorReward(
        address indexed delegatorAddress,
        address indexed withdrawAddress,
        string amount
    );

    /**
     * @dev WithdrawValidatorCommission defines an Event emitted when withdraw commission by validator
     */
    event WithdrawValidatorCommission(
        address indexed validatorAddress,
        string amount
    );

    /**
     * @dev FundCommunityPool defines an Event emitted when a user fund community pool
     */
    event FundCommunityPool(
        address indexed depositor,
        string amount
    );
}
