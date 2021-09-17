// SPDX-License-Identifier: MIT 

pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract Nft is Ownable, ERC721 {
    
    string public baseUri;
    
    uint256 public totalSupply;
    
    constructor(string memory _name, string memory _symble, string memory _baseUri) ERC721(_name, _symble) {
        baseUri = _baseUri;
    }
    
    function _baseURI() internal view virtual override returns (string memory) {
        return baseUri;
    }
    
    function Mint(address to) public {
        _safeMint(to, totalSupply);
        totalSupply += 1;
    }
    
    function destroy() external {
        selfdestruct(payable(msg.sender));
    }
}

contract Sys {
    
    uint256 mintPirce ;
    Nft public nft;
    
    constructor(uint256 _mintPrice) {
        mintPirce = _mintPrice;
    }
    
    receive() external payable{
        if(msg.value > mintPirce) {
             nft.Mint(msg.sender);
        } 
    }
    
    function receiveTokenAndMintNft() external payable {
        //require(msg.value > mintPirce, "Insufficient funds");
        
        nft.Mint(msg.sender);
        
        
        
    }
    
    
    event CreateNftContract(address indexed addr);
    
    function createNftListContract(string memory _name, string memory _symble, string memory _baseUri) public returns(address) {
       
        bytes32 salt = keccak256(abi.encodePacked(_name, _symble, _baseUri));
        
        nft = new Nft{salt : salt}(_name, _symble, _baseUri);
        
        emit CreateNftContract(address(nft));
        
        return address(nft);
    }
    
    
    
    
}
