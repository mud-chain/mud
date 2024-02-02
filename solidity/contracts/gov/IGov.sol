// SPDX-License-Identifier: Apache-2.0

pragma solidity ^0.8.0;

import "../common/Types.sol";

interface IGov {
    function legacySubmitProposal(
        string memory title,
        string memory description,
        Coin[] memory initialDeposit
    ) external returns (uint64 proposalId);

    function submitProposal(
        string memory messages,
        Coin[] memory initialDeposit,
        string memory metadata
    ) external returns (uint64 proposalId);

    function vote(
        uint64 proposalId,
        int32 option,
        string memory metadata
    ) external returns (bool success);

    function voteWeighted(
        uint64 proposalId,
        WeightedVoteOption[] memory options,
        string memory metadata
    ) external returns (bool success);

    function deposit(
        uint64 proposalId,
        Coin[] memory amount
    ) external returns (bool success);

    event LegacySubmitProposal(
        address indexed proposer,
        uint64 proposalId
    );

    event SubmitProposal(
        address indexed proposer,
        uint64 proposalId
    );

    event Vote(
        address indexed voter,
        uint64 proposalId,
        int32 option
    );

    event VoteWeighted(
        address indexed voter,
        uint64 proposalId
    );

    event Deposit(
        address indexed depositor,
        uint64 proposalId
    );
}
