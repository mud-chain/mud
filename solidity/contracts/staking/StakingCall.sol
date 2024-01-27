// SPDX-License-Identifier: Apache-2.0

pragma solidity ^0.8.0;

/* solhint-disable no-global-import */
import "./Encode.sol";
import "./Decode.sol";

/* solhint-enable no-global-import */

library StakingCall {
    address public constant STAKING_ADDRESS =
        address(0x0000000000000000000000000000000000001003);

    function delegate(
        string memory _val,
        uint256 _value
    ) internal returns (uint256, uint256) {
        (bool result, bytes memory data) = STAKING_ADDRESS.call{value: _value}(
            Encode.delegate(_val)
        );
        Decode.ok(result, data, "delegate failed");
        return Decode.delegate(data);
    }

    function undelegate(
        string memory _val,
        uint256 _shares
    ) internal returns (uint256, uint256, uint256) {
        // solhint-disable-next-line avoid-low-level-calls
        (bool result, bytes memory data) = STAKING_ADDRESS.call(
            Encode.undelegate(_val, _shares)
        );
        Decode.ok(result, data, "undelegate failed");
        return Decode.undelegate(data);
    }

    function redelegate(
        string memory _valSrc,
        string memory _valDst,
        uint256 _shares
    ) internal returns (uint256, uint256, uint256) {
        // solhint-disable-next-line avoid-low-level-calls
        (bool result, bytes memory data) = STAKING_ADDRESS.call(
            Encode.redelegate(_valSrc, _valDst, _shares)
        );
        Decode.ok(result, data, "redelegate failed");
        return Decode.redelegate(data);
    }

    function withdraw(string memory _val) internal returns (uint256) {
        // solhint-disable-next-line avoid-low-level-calls
        (bool result, bytes memory data) = STAKING_ADDRESS.call(
            Encode.withdraw(_val)
        );
        Decode.ok(result, data, "withdraw failed");
        return Decode.withdraw(data);
    }

    function delegation(
        string memory _val,
        address _del
    ) internal view returns (uint256, uint256) {
        (bool result, bytes memory data) = STAKING_ADDRESS.staticcall(
            Encode.delegation(_val, _del)
        );
        Decode.ok(result, data, "delegation failed");
        return Decode.delegation(data);
    }

    function delegationRewards(
        string memory _val,
        address _del
    ) internal view returns (uint256) {
        (bool result, bytes memory data) = STAKING_ADDRESS.staticcall(
            Encode.delegationRewards(_val, _del)
        );
        Decode.ok(result, data, "delegationRewards failed");
        return Decode.delegationRewards(data);
    }
}
