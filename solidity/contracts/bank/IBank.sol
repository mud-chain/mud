// SPDX-License-Identifier: Apache-2.0

pragma solidity ^0.8.0;

import "../common/Types.sol";

interface IBank {
    function send(
        address toAddress,
        Coin[] memory amount
    ) external returns (bool success);

    function balance(
        address accountAddress,
        string memory denom
    ) external view returns (Coin memory balance);

    function allBalances(
        address accountAddress,
        PageRequest memory pageRequest
    ) external view returns (Coin[] memory balances, PageResponse memory pageResponse);

    // totalSupply queries the total supply of all coins.
    function totalSupply() external view returns (Coin[] memory supply);

    event Send(
        address indexed fromAddress,
        address indexed toAddress,
        string amount
    );
}
