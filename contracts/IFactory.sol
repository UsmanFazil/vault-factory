// SPDX-License-Identifier: MIT
pragma solidity 0.8.9;

// Interface contract for the Factory contract
interface IFactory {

    // Create a new Vault contract
    function createVault() external;

    // Get the fee for the Vault contract
    function getFee() external view returns (uint256);

    // Get the fee recipient for the Vault contract
    function getFeeRecipient() external view returns (address);

    // Get the address of the withdraw admin
    function getWithdrawAdminAddr() external view returns (address);

    // Get the address of the fund collector 
    function getFundCollector() external view returns (address);

    // Set the fee for the Vault contract
    function setFee(uint256 newFee) external;

    // Set the fee recipient for the Vault contract
    function setFeeRecipient(address newFeeRecipient) external;

    // Set the address of the withdraw admin
    function setWithdrawAdmin(address newWithdrawAdmin) external;

    // Set the address of the fund collector
    function setFundCollector(address newFundCollector) external;
}