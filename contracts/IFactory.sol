// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

// Interface contract for the Factory contract
interface IFactory {

    // Create a new Vault contract
    function createVault(address) external;

    // Get the fee for the Vault contract
    function getFee() external view returns (uint256);

    // Get the fee recipient for the Vault contract
    function getFeeRecipient() external view returns (address);

    // Set the fee for the Vault contract
    function setFee(uint256 newFee) external;

    // Set the fee recipient for the Vault contract
    function setFeeRecipient(address newFeeRecipient) external;
}