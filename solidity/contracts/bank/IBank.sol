// SPDX-License-Identifier: Apache-2.0

pragma solidity ^0.8.0;

import "../common/Types.sol";

interface IBank {
    /**
     * @dev send defines a method for sending coins from one account to another account.
     */
    function send(
        address toAddress,
        Coin[] memory amount
    ) external returns (bool success);

    /**
     * @dev balance queries the balance of a single coin for a single account.
     */
    function balance(
        address accountAddress,
        string memory denom
    ) external view returns (Coin memory balance);

    /**
     * @dev allBalances queries the balance of all coins for a single account.
     */
    function allBalances(
        address accountAddress,
        PageRequest memory pageRequest
    ) external view returns (Coin[] memory balances, PageResponse memory pageResponse);

    /**
     * @dev totalSupply queries the total supply of all coins.
     */
    function totalSupply() external view returns (Coin[] memory supply);

    /**
     * @dev Send defines an Event emitted when a given amount of tokens send fromAddress to toAddress
     */
    event Send(
        address indexed fromAddress,
        address indexed toAddress,
        string amount
    );
}
