// SPDX-License-Identifier: MIT 

pragma solidity ^0.8.0;

interface INFT {

	function Mint(address to) external;

	function totalSupply() external returns(uint256) ;
}