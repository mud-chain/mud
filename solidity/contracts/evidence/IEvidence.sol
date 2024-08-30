// SPDX-License-Identifier: Apache-2.0

pragma solidity ^0.8.0;

import "../common/Types.sol";

/**
 * @dev Equivocation implements the Evidence interface and defines evidence of double
 * signing misbehavior.
 */
struct Equivocation {
    int64 height;
    int64 time;
    int64 power;
    address consensusAddress;
    string evidenceHash;
}

interface IEvidence {
    /**
     * @dev SubmitEvidence submits an arbitrary Evidence of misbehavior such as equivocation or
     * counterfactual signing.
     */
    function submitEvidence(int64 height, int64 time, int64 power, address consensusAddress) external returns (string memory hash);

    /**
     * @dev evidence queries evidence based on evidence hash.
     */
    function evidence(string memory evidenceHash) external view returns (Equivocation memory evidence);

    /**
     * @dev allEvidence queries all evidence.
     */
    function allEvidence(PageRequest calldata pagination) external view returns (Equivocation[] memory evidence, PageResponse memory pageResponse);

    /**
     * @dev SubmitEvidence defines an event emitted when a submit evidence is submitted.
     */
    event SubmitEvidence(
        address indexed submitter,
        string  evidenceHash
    );
}
