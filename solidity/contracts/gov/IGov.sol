// SPDX-License-Identifier: Apache-2.0

pragma solidity ^0.8.0;

import "../common/Types.sol";

interface IGov {
    function legacySubmitProposal(
        string memory title,
        string memory description,
        Coin memory initialDeposit
    ) external returns (uint64 proposalId);

    function vote(
        uint64 proposalId,
        int32 option,
        string memory metadata
    ) external returns (bool success);

    event LegacySubmitProposal(
        address indexed proposer,
        uint64 proposalId
    );

    event Vote(
        address indexed voter,
        uint64 proposalId,
        int32 option
    );
}
