// SPDX-License-Identifier: MIT
pragma solidity 0.8.9;

import "@openzeppelin/contracts/access/Ownable.sol";
import "./vault.sol";

contract Factory is Ownable {

    address public feeRecipient;
    uint256 public fee;
    address public usdcAddress;
    // Mapping to store the deployed contract addresses for each user
    mapping(address => address) public userContracts;

    constructor(address _USDC, address _feeRecipient, uint256 _fee){
        usdcAddress = _USDC;
        feeRecipient = _feeRecipient;
        fee = _fee;
    }

    // Create a new Vault contract
    function createVault() public {
        // Check if the caller has already deployed a contract
        require(userContracts[msg.sender] == address(0), "You have already deployed a contract");

        // Deploy a new instance of the Vault contract
        Vault vault = new Vault(usdcAddress, feeRecipient, fee);
        vault.transferOwnership(msg.sender);

        // Store the address of the deployed contract in the mapping
        userContracts[msg.sender] = address(vault);
    }
}