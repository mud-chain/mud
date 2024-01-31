// SPDX-License-Identifier: Apache-2.0

pragma solidity ^0.8.0;

/* solhint-disable no-global-import */
import "./Encode.sol";
import "./Decode.sol";

/* solhint-enable no-global-import */

library StakingCall {
    address public constant STAKING_ADDRESS = address(0x0000000000000000000000000000000000001003);

    function createValidator(
        Description calldata description,
        CommissionRates calldata commission,
        uint256 minSelfDelegation,
        string memory pubkey,
        uint256 value
    ) internal returns (bool) {
        (bool result, bytes memory data) = STAKING_ADDRESS.call(
            Encode.createValidator(
                description,
                commission,
                minSelfDelegation,
                pubkey,
                value
            )
        );
        Decode.ok(result, data, "createValidator failed");
        return Decode.createValidator(data);
    }

    function delegate(
        address validatorAddress,
        uint256 amount
    ) internal returns (bool) {
        (bool result, bytes memory data) = STAKING_ADDRESS.call(
            Encode.delegate(validatorAddress, amount)
        );
        Decode.ok(result, data, "delegate failed");
        return Decode.delegate(data);
    }

    function undelegate(
        address validatorAddress,
        uint256 amount
    ) internal returns (uint256) {
        (bool result, bytes memory data) = STAKING_ADDRESS.call(
            Encode.undelegate(validatorAddress, amount)
        );
        Decode.ok(result, data, "undelegate failed");
        return Decode.undelegate(data);
    }

    function redelegate(
        address validatorSrcAddress,
        address validatorDstAddress,
        uint256 amount
    ) internal returns (uint256) {
        // solhint-disable-next-line avoid-low-level-calls
        (bool result, bytes memory data) = STAKING_ADDRESS.call(
            Encode.redelegate(validatorSrcAddress, validatorDstAddress, amount)
        );
        Decode.ok(result, data, "redelegate failed");
        return Decode.redelegate(data);
    }

    function cancelUnbondingDelegation(
        address validatorAddress,
        uint256 amount,
        uint256 creationHeight
    ) internal returns (bool) {
        // solhint-disable-next-line avoid-low-level-calls
        (bool result, bytes memory data) = STAKING_ADDRESS.call(
            Encode.cancelUnbondingDelegation(validatorAddress, amount, creationHeight)
        );
        Decode.ok(result, data, "cancelUnbondingDelegation failed");
        return Decode.cancelUnbondingDelegation(data);
    }

    function delegation(
        address delegatorAddress,
        address validatorAddress
    ) internal view returns (uint256, types.Coin memory) {
        (bool result, bytes memory data) = STAKING_ADDRESS.staticcall(
            Encode.delegation(delegatorAddress, validatorAddress)
        );
        Decode.ok(result, data, "delegation failed");
        return Decode.delegation(data);
    }
}
