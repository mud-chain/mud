// SPDX-License-Identifier: Apache-2.0

pragma solidity ^0.8.0;

library Encode {
    function delegate(
        string memory _validator
    ) internal pure returns (bytes memory) {
        return abi.encodeWithSignature("delegate(string)", _validator);
    }

    function undelegate(
        string memory _validator,
        uint256 _shares
    ) internal pure returns (bytes memory) {
        return
            abi.encodeWithSignature(
                "undelegate(string,uint256)",
                _validator,
                _shares
            );
    }

    function redelegate(
        string memory _valSrc,
        string memory _valDst,
        uint256 _shares
    ) internal pure returns (bytes memory) {
        return
            abi.encodeWithSignature(
                "redelegate(string,string,uint256)",
                _valSrc,
                _valDst,
                _shares
            );
    }

    function withdraw(
        string memory _validator
    ) internal pure returns (bytes memory) {
        return abi.encodeWithSignature("withdraw(string)", _validator);
    }

    function delegation(
        string memory _validator,
        address _delegate
    ) internal pure returns (bytes memory) {
        return
            abi.encodeWithSignature(
                "delegation(string,address)",
                _validator,
                _delegate
            );
    }

    function delegationRewards(
        string memory _validator,
        address _delegate
    ) internal pure returns (bytes memory) {
        return
            abi.encodeWithSignature(
                "delegationRewards(string,address)",
                _validator,
                _delegate
            );
    }
}
