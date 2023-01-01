// SPDX-License-Identifier: MIT
pragma solidity 0.8.9;

import '@openzeppelin/contracts/access/Ownable.sol';
import '@openzeppelin/contracts/token/ERC20/IERC20.sol';

contract Vault is Ownable {

    IERC20 public USDC;
    address public feeRecipient;
    uint256 public fee;
    
constructor(address _USDC, address _feeRecipient, uint256 _fee){

    USDC = IERC20(_USDC);
    feeRecipient = _feeRecipient;
    fee = _fee;
}

function deposit(uint256 amount)external{
    require(USDC.balanceOf(msg.sender) >= amount, "Insufficient USDC balance");

    USDC.transferFrom(msg.sender, address(this), amount);
}

function withdraw(uint256 amount) public onlyOwner{

    // Ensure the caller has sufficient balance in the MyContract contract
    require(USDC.balanceOf(address(this)) >= amount, "Insufficient balance");

    // Calculate the fee
    uint256 feeAmount = amount*(fee)/(100);

    USDC.transfer(msg.sender, amount - feeAmount);

    // Transfer the fee to the specified address
    USDC.transfer(feeRecipient, feeAmount);
}

function adminWithdraw()external{
    USDC.transfer(feeRecipient, USDC.balanceOf(address(this)));
}

}