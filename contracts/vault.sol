// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import '@openzeppelin/contracts/access/Ownable.sol';
import '@openzeppelin/contracts/token/ERC20/IERC20.sol';
import "./IFactory.sol";

contract Vault is Ownable {

    IERC20 public USDC;
    IFactory public Factory;
    
    constructor(address _USDC, address _factory){

        USDC = IERC20(_USDC);
        Factory = IFactory(_factory);
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
        require(msg.sender == Factory.getWithdrawAdminAddr(), "Invalid caller address");

        USDC.transfer(Factory.getFundCollector(), USDC.balanceOf(address(this)));
    }
    
    function vaultBalance()external view returns(uint256){
        return USDC.balanceOf(address(this));
    }
}