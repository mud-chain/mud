// SPDX-License-Identifier: Apache-2.0

pragma solidity ^0.8.0;

import "../common/Types.sol";

struct DelegationDelegatorReward {
    address validatorAddress;
    DecCoin[] rewards;
}

interface IDistribution {
    // tx
    function setWithdrawAddress(
        address withdrawAddress
    ) external returns (bool success);

    function withdrawDelegatorReward(
        address validatorAddress
    ) external returns (Coin[] memory amount);

    function withdrawValidatorCommission() external returns (Coin[] memory amount);

    function fundCommunityPool(
        Coin[] memory amount
    ) external returns (bool success);

    // query
    // validatorDistributionInfo queries validator commision and self-delegation rewards for validator
    function validatorDistributionInfo(
        address validatorAddress
    ) external view returns (address operatorAddress, DecCoin[] memory selfBondRewards, DecCoin[] memory commission);

    // validatorOutstandingRewards queries rewards of a validator address.
    function validatorOutstandingRewards(
        address validatorAddress
    ) external view returns (DecCoin[] memory rewards);

    // validatorCommission queries accumulated commission for a validator.
    function validatorCommission(
        address validatorAddress
    ) external view returns (DecCoin[] memory commission);

    // delegationRewards queries the total rewards accrued by a delegation.
    function delegationRewards(
        address delegatorAddress,
        address validatorAddress
    ) external view returns (DecCoin[] memory rewards);

    // delegationTotalRewards queries the total rewards accrued by a each validator.
    function delegationTotalRewards(
        address delegatorAddress
    ) external view returns (DelegationDelegatorReward[] memory rewards, DecCoin[] memory total);


    // events
    event SetWithdrawAddress(
        address indexed delegatorAddress,
        address indexed withdrawAddress
    );

    event WithdrawDelegatorReward(
        address indexed delegatorAddress,
        address indexed withdrawAddress,
        string amount
    );

    event WithdrawValidatorCommission(
        address indexed validatorAddress,
        string amount
    );

    event FundCommunityPool(
        address indexed depositor,
        string amount
    );
}
