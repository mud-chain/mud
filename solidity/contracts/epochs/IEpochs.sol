// SPDX-License-Identifier: Apache-2.0

pragma solidity ^0.8.0;

import "../common/Types.sol";

/**
 * @dev EpochInfo defines the message interface containing the relevant informations about
 * an epoch.
 */
struct EpochInfo {
    // identifier of the epoch
    string identifier;
    // startTime of the epoch
    int64 startTime;
    // duration of the epoch
    int64 duration;
    // currentEpoch is the integer identifier of the epoch
    int64 currentEpoch;
    // currentEpochStartTime defines the timestamp of the start of the epoch
    int64 currentEpochStartTime;
    // epochCountingStarted reflects if the counting for the epoch has started
    bool epochCountingStarted;
    // currentEpochStartHeight of the epoch
    int64 currentEpochStartHeight;
}

interface IEpochs {
    /**
     * @dev epochInfos provide running epochInfos
     */
    function epochInfos(PageRequest memory pagination) external view returns (EpochInfo[] memory epochs, PageResponse memory pageResponse);

    /**
     * @dev currentEpoch provide current epoch of specified identifier
     */
    function currentEpoch(string memory identifier) external view returns (int64 currentEpoch);
}
