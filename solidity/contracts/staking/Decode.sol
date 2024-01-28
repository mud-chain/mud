// SPDX-License-Identifier: Apache-2.0

pragma solidity ^0.8.0;

import "../common/Types.sol" as types;

library Decode {
    function delegate(
        bytes memory data
    ) internal pure returns (bool) {
        (bool success) = abi.decode(data, (bool));
        return success;
    }

    function undelegate(
        bytes memory data
    ) internal pure returns (uint256) {
        uint256 completionTime = abi.decode(data, (uint256));
        return completionTime;
    }

    function redelegate(
        bytes memory data
    ) internal pure returns (uint256) {
        uint256 completionTime = abi.decode(data, (uint256));
        return completionTime;
    }

    function cancelUnbondingDelegation(
        bytes memory data
    ) internal pure returns (bool) {
        (bool success) = abi.decode(data, (bool));
        return success;
    }

    function delegation(
        bytes memory data
    ) internal pure returns (uint256, types.Coin memory) {
        (uint256 shares, types.Coin memory balance) = abi.decode(data, (uint256, types.Coin));
        return (shares, balance);
    }

    function ok(
        bool _result,
        bytes memory _data,
        string memory _msg
    ) internal pure {
        if (!_result) {
            string memory errMsg = abi.decode(_data, (string));
            if (bytes(_msg).length < 1) {
                revert(errMsg);
            }
            revert(string(abi.encodePacked(_msg, ": ", errMsg)));
        }
    }
}
