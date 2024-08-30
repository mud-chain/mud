// SPDX-License-Identifier: Apache-2.0

pragma solidity ^0.8.0;

import "../common/Types.sol";

/// @dev The IAuth contract's address.
address constant AUTH_PRECOMPILE_ADDRESS = 0x0000000000000000000000000000000000001001;

/// @dev The IAuth contract's instance.
IAuth constant AUTH_CONTRACT = IAuth(AUTH_PRECOMPILE_ADDRESS);

/**
 * @dev Account defines the account for the auth module.
 */
struct Account {
    address addr;
    string pubKey;
    uint64 accountNumber;
    uint64 sequence;
}

/**
 * @dev ModuleAccount defines the module account for the auth module.
 */
struct ModuleAccount {
    address addr;
    string pubKey;
    uint64 accountNumber;
    uint64 sequence;
    string name;
    string[] permissions;
}

/**
 * @dev Params defines the parameters for the auth module.
 */
struct Params {
    uint64 maxMemoCharacters;
    uint64 txSigLimit;
    uint64 txSizeCostPerByte;
    uint64 sigVerifyCostEd25519;
    uint64 sigVerifyCostSecp256k1;
}

interface IAuth {
    /**
     * @dev account returns account details based on address.
     */
    function account(
        address addr
    ) external view returns (Account memory account);

    /**
     * @dev accounts returns all the existing accounts
     */
    function accounts(
        PageRequest memory pageRequest
    ) external view returns (Account[] memory accounts, PageResponse memory pageResponse);

    /**
     * @dev moduleAccount returns module account details based on address.
     */
    function moduleAccountByName(
        string memory name
    ) external view returns (ModuleAccount memory moduleAccount);

    /**
     * @dev moduleAccounts returns all the existing module accounts
     */
    function moduleAccounts( ) external view returns (ModuleAccount[] memory moduleAccounts);

    /**
     * @dev accountAddressByID returns account address based on account number.
     */
    function accountAddressByID(
        int64 id
    ) external view returns (address addr);

    /**
     * @dev params queries the parameters of x/auth module.
     */
    function params() external view returns (Params memory params);
}
