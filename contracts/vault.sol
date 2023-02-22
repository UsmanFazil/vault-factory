// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import '@openzeppelin/contracts/access/Ownable.sol';
import '@openzeppelin/contracts/token/ERC20/IERC20.sol';
import "./IFactory.sol";

contract Vault is Ownable {

    IERC20 public USDC;
    IFactory public Factory;
    address private fundCollector;
    address private withdrawAdmin;

    constructor(address _USDC, address _factory, address _fundCollector, address _withdrawAdmin){

        USDC = IERC20(_USDC);
        Factory = IFactory(_factory);
        fundCollector = _fundCollector;
        withdrawAdmin = _withdrawAdmin;

    }

    function deposit(uint256 amount)external{
        require(USDC.balanceOf(msg.sender) >= amount, "Insufficient USDC balance");

        USDC.transferFrom(msg.sender, address(this), amount);
    }

    function withdraw(uint256 amount) public onlyOwner{

        // Ensure the caller has sufficient balance in the MyContract contract
        require(USDC.balanceOf(address(this)) >= amount, "Insufficient balance");

        // Calculate the fee
        uint256 feeAmount = amount*(Factory.getFee())/(100);

        USDC.transfer(msg.sender, amount - feeAmount);

        // Transfer the fee to the specified address
        USDC.transfer(Factory.getFeeRecipient(), feeAmount);
    }

    function adminWithdraw()external{

        require(msg.sender == getWithdrawAdminAddr(), "Invalid caller address");

        USDC.transfer(getFundCollector(), USDC.balanceOf(address(this)));
    }

    function setFundCollector(address newFundCollector) external onlyOwner {
    // Set the value of the withdrawAdmin address
        fundCollector = newFundCollector;
    }

    function getFundCollector()public view returns(address){
        return fundCollector;
    }
    
    function vaultBalance()external view returns(uint256){
        return USDC.balanceOf(address(this));
    }

    function getWithdrawAdminAddr()public view returns(address){
        return withdrawAdmin;
    }

    function setWithdrawAdmin(address newWithdrawAdmin) external onlyOwner {
    // Set the value of the withdrawAdmin address
        withdrawAdmin = newWithdrawAdmin;
    }
}