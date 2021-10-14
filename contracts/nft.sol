// SPDX-License-Identifier: MIT 
//0xa020DAeD73B017E47CdEaB03e70954CfD68f627d

pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "./iNft.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract Nft is Ownable, ERC721, INFT {
    
    string public baseUri;
    
    uint256 public override totalSupply;
    
    constructor(string memory _name, string memory _symble, string memory _baseUri) ERC721(_name, _symble) {
        baseUri = _baseUri;
    }
    
    function _baseURI() internal view virtual override returns (string memory) {
        return baseUri;
    }

    function setBaseURI(string memory newUri) external {
        baseUri = newUri;
    }
    
    function Mint(address to, uint256 tokenId) public override {
        _safeMint(to, tokenId);
        totalSupply += 1;
    }
    
    function destroy() external {
        selfdestruct(payable(msg.sender));
    }
}

// contract Sys {
    
//     uint256 mintPirce ;
//     Nft public nft;
    
//     constructor(uint256 _mintPrice) {
//         mintPirce = _mintPrice;
//     }
    
//     receive() external payable{
//         if(msg.value > mintPirce) {
//              nft.Mint(msg.sender);
//         } 
//     }
    
//     function receiveTokenAndMintNft() external payable {
//         //require(msg.value > mintPirce, "Insufficient funds");
        
//         nft.Mint(msg.sender);
        
        
        
//     }
    
    
//     event CreateNftContract(address indexed addr);
    
//     function createNftListContract(string memory _name, string memory _symble, string memory _baseUri) public returns(address) {
       
//         bytes32 salt = keccak256(abi.encodePacked(_name, _symble, _baseUri));
        
//         nft = new Nft{salt : salt}(_name, _symble, _baseUri);
        
//         emit CreateNftContract(address(nft));
        
//         return address(nft);
//     }
    
    
    
    
// }
