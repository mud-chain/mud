// SPDX-License-Identifier: Apache-2.0

pragma solidity ^0.8.0;

struct Coin {
    string denom;
    uint256 amount;
}

struct DecCoin {
    string denom;
    uint256 amount;
    uint8 precision;
}

struct Dec {
    uint256 amount;
    uint8 precision;
}

struct PageRequest {
    bytes key;
    uint64 offset;
    uint64 limit;
    bool countTotal;
    bool reverse;
}

struct PageResponse {
    bytes nextKey;
    uint64 total;
}