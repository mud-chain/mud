// SPDX-License-Identifier: Apache-2.0

pragma solidity ^0.8.0;

import "../common/Types.sol";

interface IDistribution {
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

    function delegationRewards(
        address delegatorAddress,
        address validatorAddress
    ) external view returns (DecCoin[] memory rewards);


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
