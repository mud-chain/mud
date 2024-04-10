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

// Delegation represents the bond with tokens held by an account. It is
// owned by one delegator, and is associated with the voting power of one
// validator.
struct Delegation {
    address delegatorAddress;
    address validatorAddress;
    Dec shares;
}

// DelegationResponse is equivalent to Delegation except that it contains a
// balance in addition to shares which is more suitable for client responses.
struct DelegationResponse {
    Delegation delegation;
    Coin balance;
}

// UnbondingDelegationEntry defines an unbonding object with relevant metadata.
struct UnbondingDelegationEntry {
    int64 creationHeight;
    int64 completionTime;
    uint256 initialBalance;
    uint256 balance;
}

// UnbondingDelegation stores all of a single delegator's unbonding bonds
// for a single validator in an time-ordered list.
struct UnbondingDelegation {
    address delegatorAddress;
    address validatorAddress;
    UnbondingDelegationEntry[] entries;
}

interface IStaking {
    // tx
    function createValidator(
        Description calldata description,
        CommissionRates calldata commission,
        uint256 minSelfDelegation,
        string calldata pubkey,
        uint256 value
    ) external returns (bool success);

    function editValidator(
        Description calldata description,
        int256 commissionRate,
        int256 minSelfDelegation
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
    function validators(
        BondStatus status,
        PageRequest calldata pagination
    ) external view returns (Validator[] calldata validators, PageResponse calldata pageResponse);

    function validator(
        address validatorAddr
    ) external view returns (Validator calldata validator);

    // validatorDelegations queries delegate info for given validator.
    function validatorDelegations(
        address validatorAddr,
        PageRequest calldata pagination
    ) external view returns (DelegationResponse[] calldata response, PageResponse calldata pageResponse);

    // validatorUnbondingDelegations queries unbonding delegations of a validator.
    function validatorUnbondingDelegations(
        address validatorAddr,
        PageRequest calldata pagination
    ) external view returns (UnbondingDelegation[] calldata response, PageResponse calldata pageResponse);

    // delegation queries delegate info for given validator delegator pair.
    function delegation(
        address delegatorAddr,
        address validatorAddr
    ) external view returns (DelegationResponse calldata response);

    // unbondingDelegation queries unbonding info for given validator delegator pair.
    function unbondingDelegation(
        address delegatorAddr,
        address validatorAddr
    ) external view returns (UnbondingDelegation calldata response);

    // delegatorDelegations queries all delegations of a given delegator address.
    function delegatorDelegations(
        address delegatorAddr,
        PageRequest calldata pagination
    ) external view returns (DelegationResponse[] calldata response, PageResponse calldata pageResponse);

    // delegatorUnbondingDelegations queries all unbonding delegations of a given delegator address.
    function delegatorUnbondingDelegations(
        address delegatorAddr,
        PageRequest calldata pagination
    ) external view returns (UnbondingDelegation[] calldata response, PageResponse calldata pageResponse);

    // delegatorValidators queries all validators info for given delegator address.
    function delegatorValidators(
        address delegatorAddr,
        PageRequest calldata pagination
    ) external view returns (Validator[] calldata validators, PageResponse calldata pageResponse);

    // delegatorValidator queries validator info for given delegator validator pair.
    function delegatorValidator(
        address delegatorAddr,
        address validatorAddr
    ) external view returns (Validator calldata validator);

    // events
    event CreateValidator(
        address indexed validator,
        uint256 value
    );

    event EditValidator(
        address indexed validator,
        int256 commissionRate,
        int256 minSelfDelegation
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
