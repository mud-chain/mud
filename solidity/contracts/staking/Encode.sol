// SPDX-License-Identifier: Apache-2.0

pragma solidity ^0.8.0;

import "../common/Types.sol";

library Encode {
    function createValidator(
        Description calldata description,
        CommissionRates calldata commission,
        uint256 minSelfDelegation,
        string memory pubkey,
        uint256 value
    ) internal pure returns (bytes memory) {
        return abi.encodeWithSignature(
            "createValidator(Description,CommissionRates,uint256,uint256)",
            description,
            commission,
            minSelfDelegation,
            pubkey,
            value
        );
    }

    function delegate(
        address validatorAddress,
        uint256 amount
    ) internal pure returns (bytes memory) {
        return abi.encodeWithSignature(
            "delegate(address,uint256)",
            validatorAddress,
            amount
        );
    }

    function undelegate(
        address validatorAddress,
        uint256 amount
    ) internal pure returns (bytes memory) {
        return
            abi.encodeWithSignature(
            "undelegate(address,uint256)",
            validatorAddress,
            amount
        );
    }

    function redelegate(
        address validatorSrcAddress,
        address validatorDstAddress,
        uint256 amount
    ) internal pure returns (bytes memory) {
        return
            abi.encodeWithSignature(
            "redelegate(address,address,uint256)",
            validatorSrcAddress,
            validatorDstAddress,
            amount
        );
    }

    function cancelUnbondingDelegation(
        address validatorAddress,
        uint256 amount,
        uint256 creationHeight
    ) internal pure returns (bytes memory) {
        return
            abi.encodeWithSignature(
            "cancelUnbondingDelegation(address,uint256,uint256)",
            validatorAddress,
            amount,
            creationHeight
        );
    }

    function delegation(
        address delegatorAddress,
        address validatorAddress
    ) internal pure returns (bytes memory) {
        return
            abi.encodeWithSignature(
            "delegation(address,address)",
            delegatorAddress,
            validatorAddress
        );
    }
}
