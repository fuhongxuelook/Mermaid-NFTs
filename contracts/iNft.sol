// SPDX-License-Identifier: MIT 

pragma solidity ^0.8.0;

interface INFT {

	function Mint(address to, uint256 tokenId) external;

	function totalSupply() external returns(uint256) ;
}