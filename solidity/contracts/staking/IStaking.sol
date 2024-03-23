// SPDX-License-Identifier: Apache-2.0

pragma solidity ^0.8.0;

import "../common/Types.sol";

struct Description {
    string moniker;
    string identity;
    string website;
    string securityContact;
    string details;
}

struct CommissionRates {
    uint256 rate;
    uint256 maxRate;
    uint256 maxChangeRate;
}

struct Commission {
    CommissionRates commissionRates;
    int64 updateTime;
}

enum BondStatus {
    Unspecified,
    Unbonded,
    Unbonding,
    Bonded
}

struct Validator {
    address operatorAddress;
    string consensusPubkey;
    bool jailed;
    BondStatus status;
    uint256 tokens;
    uint256 delegatorShares;
    Description description;
    int64 unbondingHeight;
    int64 unbondingTime;
    Commission commission;
    uint256 minSelfDelegation;
}

interface IStaking {
    // tx
    function createValidator(
        Description calldata description,
        CommissionRates calldata commission,
        uint256 minSelfDelegation,
        string memory pubkey,
        uint256 value
    ) external returns (bool success);

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

    // query
    function delegation(
        address delegatorAddress,
        address validatorAddress
    ) external view returns (uint256 shares, Coin memory balance);

    function validators(
        BondStatus status,
        PageRequest memory pagination
    ) external view returns (Validator[] memory validators, PageResponse memory pageResponse);

    function validator(
        address validatorAddr
    ) external view returns (Validator memory validator);

    function validatorDelegations(
        address validatorAddr,
        PageRequest memory pagination
    ) external view returns (Validator memory validator);

    // events
    event CreateValidator(
        address indexed validator,
        uint256 value
    );

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
