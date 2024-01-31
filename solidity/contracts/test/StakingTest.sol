// SPDX-License-Identifier: Apache-2.0

pragma solidity ^0.8.0;

/* solhint-disable no-global-import */
import "../staking/IStaking.sol";
import "../staking/StakingCall.sol";

/* solhint-enable no-global-import */

contract StakingTest is IStaking {
    mapping(string => uint256) public validatorShares;

    function createValidator(
        Description calldata description,
        CommissionRates calldata commissionRates,
        uint256 minSelfDelegation,
        string memory pubkey,
        uint256 value
    ) external override returns (bool) {
        bool success = StakingCall.createValidator(
            description,
            commissionRates,
            minSelfDelegation,
            pubkey,
            value
        );
        return success;
    }

    function delegate(
        address validatorAddress,
        uint256 amount
    ) external override returns (bool) {
        bool success = StakingCall.delegate(
            validatorAddress,
            amount
        );
        return success;
    }

    function undelegate(
        address validatorAddress,
        uint256 amount
    ) external override returns (uint256) {
        uint256 completionTime = StakingCall.undelegate(
            validatorAddress,
            amount
        );
        return completionTime;
    }

    function redelegate(
        address validatorSrcAddress,
        address validatorDstAddress,
        uint256 amount
    ) external override returns (uint256) {
        uint256 completionTime = StakingCall.redelegate(
            validatorSrcAddress,
            validatorDstAddress,
            amount
        );
        return completionTime;
    }

    function cancelUnbondingDelegation(
        address validatorAddress,
        uint256 amount,
        uint256 creationHeight
    ) external override returns (bool) {
        bool success = StakingCall.cancelUnbondingDelegation(
            validatorAddress,
            amount,
            creationHeight
        );
        return success;
    }

    function delegation(
        address delegatorAddress,
        address validatorAddress
    ) public view override returns (uint256, types.Coin memory) {
        return StakingCall.delegation(delegatorAddress, validatorAddress);
    }
}
