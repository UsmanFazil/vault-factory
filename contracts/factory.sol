// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import "@openzeppelin/contracts/access/Ownable.sol";
import "./vault.sol";
import "./IFactory.sol";

contract Factory is Ownable, IFactory{

    uint256 private fee;
    address private feeRecipient;
    address private withdrawAdmin;
    address private usdcAddress;
    address private fundCollector;
    uint256 public totalVaults;

    // Mapping to store the deployed contract addresses for each user
    mapping(address => address) public userContract;
    mapping(uint256 => address) public contractAddresses;

    constructor(address _USDC, address _feeRecipient, uint256 _fee, address _withdrawAdmin, address _fundCollector){
        usdcAddress = _USDC;
        feeRecipient = _feeRecipient;
        withdrawAdmin = _withdrawAdmin;
        fundCollector = _fundCollector;
        fee = _fee;
        totalVaults = 0;
    }

    // Create a new Vault contract
    function createVault() public override{
        // Check if the caller has already deployed a contract
        require(userContract[msg.sender] == address(0), "You have already deployed a contract");

        // Deploy a new instance of the Vault contract
        Vault vault = new Vault(usdcAddress, address(this));

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

    function setWithdrawAdmin(address newWithdrawAdmin) public override onlyOwner {
    // Set the value of the withdrawAdmin address
        withdrawAdmin = newWithdrawAdmin;
    }

    function setFundCollector(address newFundCollector) public override onlyOwner {
    // Set the value of the withdrawAdmin address
        fundCollector = newFundCollector;
    }

    function getFee()external view override returns(uint256){
        return fee;
    }

    function getFeeRecipient()external view returns(address){
        return feeRecipient;
    }

    function getWithdrawAdminAddr()external view returns(address){
        return withdrawAdmin;
    }

    function getFundCollector()external view returns(address){
        return fundCollector;
    }

}