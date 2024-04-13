// SPDX-License-Identifier: Apache-2.0

pragma solidity ^0.8.0;

interface ISlashing {
    /**
     * @dev unjail defines a method for unjailing a jailed validator, thus returning
     * them into the bonded validator set, so they can begin receiving provisions
     * and rewards again.
     */
    function unjail() external returns (bool success);

    /**
     * @dev Unjail defines an Event emitted when a validtor unjail
     */
    event Unjail(
        address indexed validator
    );
}
