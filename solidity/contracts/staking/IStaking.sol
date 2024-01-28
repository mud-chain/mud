// SPDX-License-Identifier: Apache-2.0

pragma solidity ^0.8.0;

import "../common/Types.sol";

interface IStaking {
    function delegate(
        address validatorAddress,
        uint256 amount
    ) external returns (bool success);

    function undelegate(
        address validatorAddress,
        uint256 amount
    ) external returns (uint256 completionTime);

    function redelegate(
        address validatorSrcAddress,
        address validatorDstAddress,
        uint256 amount
    ) external returns (uint256 completionTime);

    function cancelUnbondingDelegation(
        address validatorAddress,
        uint256 amount,
        uint256 creationHeight
    ) external returns (bool success);

    function delegation(
        address delegatorAddress,
        address validatorAddress
    ) external view returns (uint256 shares, Coin memory balance);

    event Delegate(
        address indexed delegator,
        address indexed validator,
        uint256 amount
    );

    event Undelegate(
        address indexed delegatorAddress,
        address indexed validatorAddress,
        uint256 amount,
        uint256 completionTime
    );

    event Redelegate(
        address indexed delegatorAddress,
        address indexed validatorSrcAddress,
        address indexed validatorDstAddress,
        uint256 amount,
        uint256 completionTime
    );

    event CancelUnbondingDelegation(
        address indexed delegatorAddress,
        address indexed validatorAddress,
        uint256 amount,
        uint256 creationHeight
    );
}
