// SPDX-License-Identifier: Apache-2.0

pragma solidity ^0.8.0;

import "../common/Types.sol";

enum ProposalStatus {
    Unspecified,
    DepositPeriod,
    VotingPeriod,
    Passed,
    Rejected,
    Failed
}

enum VoteOption {
    Unspecified,
    Yes,
    Abstain,
    No,
    NoWithWeto
}

struct TallyResult {
    string yesCount;
    string abstainCount;
    string noCount;
    string noWithVetoCount;
}

struct Proposal {
    uint64 id;
    string[] messages;
    ProposalStatus status;
    TallyResult finalTallyResult;
    int64 submitTime;
    int64 depositEndTime;
    Coin[] totalDeposit;
    int64 votingStartTime;
    int64 votingEndTime;
    string metadata;
}

struct WeightedVoteOption {
    VoteOption option;
    string weight;
}

struct VoteData {
    uint64 proposalId;
    address voter;
    WeightedVoteOption[] options;
    string metadata;
}

struct DepositData {
    uint64 proposalId;
    address depositor;
    Coin[] amount;
}

interface IGov {
    // tx
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

    // query

    // proposal queries proposal details based on ProposalID.
    function proposal(
        uint64 proposalId
    ) external view returns (Proposal calldata proposal);

    // proposals queries all proposals based on given status.
    function proposals(
        ProposalStatus status,
        address voter,
        address depositor,
        PageRequest calldata pagination
    ) external view returns (Proposal[] calldata proposals, PageResponse calldata pageResponse);

    // vote queries voted information based on proposalID, voterAddr.
    function vote(
        uint64 proposalId,
        address voter
    ) external view returns (VoteData calldata vote);

    // votes queries votes of a given proposal.
    function votes(
        uint64 proposalId,
        PageRequest calldata pagination
    ) external view returns (VoteData[] calldata votes, PageResponse calldata pageResponse);

    // deposit queries single deposit information based proposalID, depositAddr.
    function depositQuery(
        uint64 proposalId,
        address depositor
    ) external view returns (DepositData calldata deposit);

    // deposits queries all deposits of a single proposal.
    function deposits(
        uint64 proposalId,
        PageRequest calldata pagination
    ) external view returns (DepositData[] calldata deposits, PageResponse calldata pageResponse);

    // tallyResult queries the tally of a proposal vote.
    function tallyResult(
        uint64 proposalId
    ) external view returns (TallyResult calldata tallyResult);

    // event
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
