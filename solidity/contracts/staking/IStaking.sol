// SPDX-License-Identifier: Apache-2.0

pragma solidity ^0.8.0;

import "../common/Types.sol";

/**
 * @dev Description defines a validator description.
 */
struct Description {
    // moniker defines a human-readable name for the validator.
    string moniker;
    // identity defines an optional identity signature (ex. UPort or Keybase).
    string identity;
    // website defines an optional website link.
    string website;
    // securityContact defines an optional email for security contact.
    string securityContact;
    // details define other optional details.
    string details;
}

/**
 * @dev CommissionRates defines the initial commission rates to be used for creating a validator.
 */
struct CommissionRates {
    // rate defines the maximum commission rate which validator can ever charge, as a fraction.
    uint256 rate;
    // maxRate defines the maximum commission rate which validator can ever charge, as a fraction.
    uint256 maxRate;
    // maxChangeRate defines the maximum daily increase of the validator commission, as a fraction.
    uint256 maxChangeRate;
}

/**
 * @dev Commission defines commission parameters for a given validator.
 */
struct Commission {
    // commissionRates defines the initial commission rates to be used for creating a validator
    CommissionRates commissionRates;
    // updateTime defines the validator update commissionRates time
    int64 updateTime;
}

/**
 * @dev BondStatus is the status of a validator.
 */
enum BondStatus {
    // Unspecified defines an invalid validator status.
    Unspecified,
    // Unbonded defines a validator that is not bonded.
    Unbonded,
    // Unbonding defines a validator that is unbonding.
    Unbonding,
    // Bonded defines a validator that is bonded.
    Bonded
}

/**
 * @dev Validator defines a validator, together with the total amount of the
 * Validator's bond shares and their exchange rate to coins. Slashing results in
 * a decrease in the exchange rate, allowing correct calculation of future
 * undelegations without iterating over delegators. When coins are delegated to
 * this validator, the validator is credited with a delegation whose number of
 * bond shares is based on the amount of coins delegated divided by the current
 * exchange rate. Voting power can be calculated as total bonded shares
 * multiplied by exchange rate.
 */
struct Validator {
    // operatorAddress defines the address of the validator's operator
    address operatorAddress;
    // consensusPubkey is the consensus public key of the validator
    string consensusPubkey;
    // jailed defined whether the validator has been jailed from bonded status or not.
    bool jailed;
    // status is the validator status (bonded/unbonding/unbonded).
    BondStatus status;
    // tokens define the delegated tokens (incl. self-delegation).
    uint256 tokens;
    // delegatorShares defines total shares issued to a validator's delegators.
    uint256 delegatorShares;
    // description defines the description terms for the validator.
    Description description;
    // unbonding_height defines, if unbonding, the height at which this validator has begun unbonding.
    int64 unbondingHeight;
    // unbonding_time defines, if unbonding, the min time for the validator to complete unbonding.
    int64 unbondingTime;
    // commission defines the commission parameters.
    Commission commission;
    // minSelfDelegation is the validator's self declared minimum self delegation.
    uint256 minSelfDelegation;
}

/**
 * @dev Delegation represents the bond with tokens held by an account. It is
 * owned by one delegator, and is associated with the voting power of one
 * validator.
 */
struct Delegation {
    // delegatorAddress is the address of the delegator.
    address delegatorAddress;
    // validatorAddress is the address of the validator.
    address validatorAddress;
    // shares define the delegation shares received.
    Dec shares;
}

/**
 * @dev DelegationResponse is equivalent to Delegation except that it contains a
 * balance in addition to shares which is more suitable for client responses.
 */
struct DelegationResponse {
    Delegation delegation;
    Coin balance;
}

/**
 * @dev UnbondingDelegationEntry defines an unbonding object with relevant metadata.
 */
struct UnbondingDelegationEntry {
    // creationHeight is the height which the unbonding took place.
    int64 creationHeight;
    // completionTime is the unix time for unbonding completion.
    int64 completionTime;
    // initialBalance defines the tokens initially scheduled to receive at completion.
    uint256 initialBalance;
    // balance defines the tokens to receive at completion.
    uint256 balance;
}

/**
 * @dev UnbondingDelegation stores all of a single delegator's unbonding bonds
 * for a single validator in an time-ordered list.
 */
struct UnbondingDelegation {
    address delegatorAddress;
    address validatorAddress;
    UnbondingDelegationEntry[] entries;
}

interface IStaking {
    /**
     * @dev createValidator defines a method for creating a new validator.
     */
    function createValidator(
        Description calldata description,
        CommissionRates calldata commission,
        uint256 minSelfDelegation,
        string calldata pubkey,
        uint256 value
    ) external returns (bool success);

    /**
     * @dev editValidator defines a method for editing an existing validator.
     */
    function editValidator(
        Description calldata description,
        int256 commissionRate,
        int256 minSelfDelegation
    ) external returns (bool success);

    /**
     * @dev delegate defines a method for performing a delegation of coins
     * from a delegator to a validator.
     */
    function delegate(
        address validatorAddress,
        uint256 amount
    ) external returns (bool success);

    /**
     * @dev undelegate defines a method for performing an undelegation from a
     * delegate and a validator.
     */
    function undelegate(
        address validatorAddress,
        uint256 amount
    ) external returns (uint256 completionTime);

    /**
     * @dev redelegate defines a method for performing a redelegation
     * of coins from a delegator and source validator to a destination validator.
     */
    function redelegate(
        address validatorSrcAddress,
        address validatorDstAddress,
        uint256 amount
    ) external returns (uint256 completionTime);

    /**
     * @dev cancelUnbondingDelegation defines a method for performing canceling the unbonding delegation
     * and delegate back to previous validator.
     */
    function cancelUnbondingDelegation(
        address validatorAddress,
        uint256 amount,
        uint256 creationHeight
    ) external returns (bool success);

    /**
     * @dev validators queries all validators that match the given status.
     */
    function validators(
        BondStatus status,
        PageRequest calldata pagination
    ) external view returns (Validator[] calldata validators, PageResponse calldata pageResponse);

    /**
     * @dev validator queries validator info for given validator address.
     */
    function validator(
        address validatorAddr
    ) external view returns (Validator calldata validator);

    /**
     * @dev validatorDelegations queries delegate info for given validator.
     */
    function validatorDelegations(
        address validatorAddr,
        PageRequest calldata pagination
    ) external view returns (DelegationResponse[] calldata response, PageResponse calldata pageResponse);

    /**
     * @dev validatorUnbondingDelegations queries unbonding delegations of a validator.
     */
    function validatorUnbondingDelegations(
        address validatorAddr,
        PageRequest calldata pagination
    ) external view returns (UnbondingDelegation[] calldata response, PageResponse calldata pageResponse);

    /**
     * @dev delegation queries delegate info for given validator delegator pair.
     */
    function delegation(
        address delegatorAddr,
        address validatorAddr
    ) external view returns (DelegationResponse calldata response);

    /**
     * @dev unbondingDelegation queries unbonding info for given validator delegator pair.
     */
    function unbondingDelegation(
        address delegatorAddr,
        address validatorAddr
    ) external view returns (UnbondingDelegation calldata response);

    /**
     * @dev delegatorDelegations queries all delegations of a given delegator address.
     */
    function delegatorDelegations(
        address delegatorAddr,
        PageRequest calldata pagination
    ) external view returns (DelegationResponse[] calldata response, PageResponse calldata pageResponse);

    /**
     * @dev delegatorUnbondingDelegations queries all unbonding delegations of a given delegator address.
     */
    function delegatorUnbondingDelegations(
        address delegatorAddr,
        PageRequest calldata pagination
    ) external view returns (UnbondingDelegation[] calldata response, PageResponse calldata pageResponse);

    /**
     * @dev CreateValidator defines an Event emitted when a new validator created.
     */
    event CreateValidator(
        address indexed validator,
        uint256 value
    );

    /**
     * @dev EditValidator defines an Event emitted when a validator edited.
     */
    event EditValidator(
        address indexed validator,
        int256 commissionRate,
        int256 minSelfDelegation
    );

    /**
     * @dev Delegate defines an Event emitted when a given amount of tokens are delegated from the
     * delegator address to the validator address.
     */
    event Delegate(
        address indexed delegator,
        address indexed validator,
        uint256 amount
    );

    /**
     * @dev Undelegate defines an Event emitted when a given amount of tokens are undelegate by delegator
     */
    event Undelegate(
        address indexed delegatorAddress,
        address indexed validatorAddress,
        uint256 amount,
        uint256 completionTime
    );

    /**
     * @dev Redelegate defines an Event emitted when a given amount of tokens are redelegated from
     * the source validator address to the destination validator address.
     */
    event Redelegate(
        address indexed delegatorAddress,
        address indexed validatorSrcAddress,
        address indexed validatorDstAddress,
        uint256 amount,
        uint256 completionTime
    );

    /**
     * @dev CancelUnbondingDelegation defines an Event emitted when a given amount of tokens are cancel
     * bond delegate from the validator by the delegator
     */
    event CancelUnbondingDelegation(
        address indexed delegatorAddress,
        address indexed validatorAddress,
        uint256 amount,
        uint256 creationHeight
    );
}
