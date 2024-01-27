// SPDX-License-Identifier: Apache-2.0

pragma solidity ^0.8.0;

interface IStaking {
    function delegate(
        string memory _val
    ) external payable returns (uint256 _shares, uint256 _reward);

    function undelegate(
        string memory _val,
        uint256 _shares
    )
        external
        returns (uint256 _amount, uint256 _reward, uint256 _completionTime);

    function redelegate(
        string memory _valSrc,
        string memory _valDst,
        uint256 _shares
    )
        external
        returns (uint256 _amount, uint256 _reward, uint256 _completionTime);

    function withdraw(string memory _val) external returns (uint256 _reward);

    function delegation(
        string memory _val,
        address _del
    ) external view returns (uint256 _shares, uint256 _delegateAmount);

    function delegationRewards(
        string memory _val,
        address _del
    ) external view returns (uint256 _reward);

    event Delegate(
        address indexed delegator,
        string validator,
        uint256 amount,
        uint256 shares
    );

    event Undelegate(
        address indexed sender,
        string validator,
        uint256 shares,
        uint256 amount,
        uint256 completionTime
    );

    event Redelegate(
        address indexed sender,
        string valSrc,
        string valDst,
        uint256 shares,
        uint256 amount,
        uint256 completionTime
    );

    event Withdraw(address indexed sender, string validator, uint256 reward);
}
