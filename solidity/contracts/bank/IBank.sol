// SPDX-License-Identifier: Apache-2.0

pragma solidity ^0.8.0;

import "../common/Types.sol";

interface IBank {
    function send(
        address toAddress,
        Coin[] memory amount
    ) external returns (bool success);

    event Send(
        address indexed fromAddress,
        address indexed toAddress,
        string amount
    );
}
