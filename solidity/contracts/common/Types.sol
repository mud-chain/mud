// SPDX-License-Identifier: Apache-2.0

pragma solidity ^0.8.0;

struct Coin {
    string denom;
    uint256 amount;
}

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