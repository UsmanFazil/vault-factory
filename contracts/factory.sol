// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import "@openzeppelin/contracts/access/Ownable.sol";
import "./vault.sol";
import "./IFactory.sol";

contract Factory is Ownable, IFactory{

    uint256 private fee;
    address private feeRecipient;
    address private usdcAddress;
    uint256 public totalVaults;
    address private adminWallet;

    // Mapping to store the deployed contract addresses for each user
    mapping(address => address) public userContract;
    mapping(uint256 => address) public contractAddresses;

    constructor(address _USDC, address _feeRecipient, uint256 _fee, address adminWallet_){
        usdcAddress = _USDC;
        feeRecipient = _feeRecipient;
        fee = _fee;
        totalVaults = 0;
        adminWallet = adminWallet_;
    }

    // Create a new Vault contract
    function createVault(address _fundCollector) external override{
        // Check if the caller has already deployed a contract
        require(userContract[msg.sender] == address(0), "You have already deployed a contract");

        // Deploy a new instance of the Vault contract
        Vault vault = new Vault(usdcAddress, address(this), _fundCollector, adminWallet);

        // Transfer vault ownership to msg.sender 
        vault.transferOwnership(msg.sender);
        totalVaults = totalVaults +1;

        // Store the address of the deployed contract in the mapping
        userContract[msg.sender] = address(vault);
        contractAddresses[totalVaults] = address(vault);
    }
    function setFee(uint256 newFee) public override onlyOwner {
    // Set the value of the fee variable
        fee = newFee;
    }

    function setFeeRecipient(address newFeeRecipient) public override onlyOwner {
    // Set the value of the feeRecipient address
        feeRecipient = newFeeRecipient;
    }

    function getFee()external view override returns(uint256){
        return fee;
    }

    function getFeeRecipient()external view returns(address){
        return feeRecipient;
    }

    function setAdminWallet(address newAdminWallet)  external onlyOwner {
        // set the value of admin wallet address for withdrawal
        adminWallet = newAdminWallet;
    }

    function getAdminWallet()  external view returns(address) {
        // set the value of admin wallet address for withdrawal
        return adminWallet;
    }
    
}