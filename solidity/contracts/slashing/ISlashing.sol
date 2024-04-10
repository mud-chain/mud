// SPDX-License-Identifier: Apache-2.0

pragma solidity ^0.8.0;

interface ISlashing {
    // tx
    function unjail() external returns (bool success);

    // events
    event Unjail(
        address indexed validator
    );
}
