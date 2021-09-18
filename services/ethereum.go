package services

import (
	"fmt"
	"log"
	"context"
	"crypto/ecdsa"
	"math/big"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
    "github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	_ "github.com/ethereum/go-ethereum/common/hexutil"
    beego "github.com/beego/beego/v2/server/web"
    nftContract "MermaidNFT/contracts" 
    _ "github.com/beego/beego/v2/core/logs"
)

const myAddr = `0xD32e5c150b9Ca49506D6f04C5498B71e6fC9d027`

func MintNFT(addr string) (interface{}) {

    client, privateResource, _ := getEnv()

    method := "Mint(address)"

    var data []byte
    addrHex := common.HexToAddress(addr)

    data = append(data, common.LeftPadBytes(addrHex.Bytes(), 32)...) 

    return sendTransaction(client, privateResource, method, data)



}

/**
 * 获取当前tokenId
 * 
 */ 
func GetTokenId() (interface{}) {
    
    client, _, _ := getEnv()

    nftAddressStr, _ := beego.AppConfig.String("userNFT")
    nftAddress := common.HexToAddress(nftAddressStr)


    instance, _ := nftContract.NewContracts(nftAddress, client)


    totalSupply, _ := instance.TotalSupply(&bind.CallOpts{})

    
    return totalSupply
}


func LoopSearchTx(hash string) (bool) {
    
    client, _, _ := getEnv()

    txHash := common.HexToAddress(hash) 
    _, isPending, err := client.TransactionByHash(context.Background(), txHash.Hash())
    
    if err != nil {
        log.Fatal(err)
    }

    return isPending;
    // _, err = client.TransactionReceipt(context.Background(), tx.Hash())
    // if err != nil {
    //     log.Fatal(err)
    // }
}


/**
 * @dev 获取私钥对应的公钥地址
 * 
 * @params
 * - `privateKey` *ecdsa.PrivateKey
 * 
 * @return 
 * 
 * - `address` common.Address
 */ 
func getAddressFromPrivateKey(privateKey *ecdsa.PrivateKey) (address common.Address) {
	publicKey := privateKey.Public()
    publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
    if !ok {
        log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
    }

    address = crypto.PubkeyToAddress(*publicKeyECDSA)

    return
}

/**
 * @dev 发送ETH给地址，这个eth是泛指，公链的gas代币统一
 * 
 * @params 
 * - *ethclient.Client 客户端
 * - *ecdsa.PrivateKey 私钥
 * 
 */
func sendEthToAddress(client *ethclient.Client, privateKey *ecdsa.PrivateKey) {

	fromAddress := getAddressFromPrivateKey(privateKey);

    nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
    if err != nil {
        log.Fatal(err)
    }

    value := big.NewInt(100000) // in wei (1 eth)
    gasLimit := uint64(901716)                // in units
    gasPrice, err := client.SuggestGasPrice(context.Background())
    if err != nil {
        log.Fatal(err)
    }

    toAddress := common.HexToAddress("0xD32e5c150b9Ca49506D6f04C5498B71e6fC9d027")
    var data []byte
    tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

    chainID, err := client.NetworkID(context.Background())
    if err != nil {
        log.Fatal(err)
    }

    signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
    if err != nil {
        log.Fatal(err)
    }

    err = client.SendTransaction(context.Background(), signedTx)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
	return
}




/**
 * @dev 发送交易 
 * 
 * @step
 * - 获取接口名称和参数,不包含空格,uint转为全称 uint2456  
 * - `approve(address,uint256)`
 * 
 * - 设置目的合约
 * - 设置参数，参数需要转为[]byte,并补0，满足32byte
 * - 接口名称和参数追加,作为整体
 * 
 * - 设置gas 和gas limit
 * - 注意：一般接口不支持同时调用方法和转移gas,特殊除外
 * 
 */ 
func sendTransaction(
    client *ethclient.Client, 
    privateKey *ecdsa.PrivateKey, 
    method string,
    data []byte) (string) {
    fromAddress := getAddressFromPrivateKey(privateKey);

    nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
    if err != nil {
        log.Fatal(err)
    }

    value := big.NewInt(0) 
    gasLimit := uint64(91716)                // in units
    gasPrice, err := client.SuggestGasPrice(context.Background())
    gasBiger := big.NewInt(10)
    gasPrice.Mul(gasPrice, gasBiger)
    if err != nil {
        log.Fatal(err)
    }

    methodId := getMethodSignature(method) 

    var sendData []byte
    sendData = append(sendData, methodId...)
    sendData = append(sendData, data...)

    nftAddress, _ := beego.AppConfig.String("userNFT")
    toAddress := common.HexToAddress(nftAddress)
    
    tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, sendData)

    chainID, err := client.NetworkID(context.Background())
    if err != nil {
        log.Fatal(err)
    }

    signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
    if err != nil {
        log.Fatal(err)
    }

    err = client.SendTransaction(context.Background(), signedTx)
    if err != nil {
        log.Fatal(err)
    }

    return signedTx.Hash().Hex()
}

/**
 * @dev 获取接口ID,调用合约的方法，需要获取方法的id
 * 
 * - 方法是无空格完整方法结构: approve(address,uint256)
 * 
 * @return []byte 4个byte大小的[]byte
 */ 
func getMethodSignature(method string) (methodId []byte ){
	methodBytes := []byte(method);
    keccak := crypto.Keccak256(methodBytes);
    methodId = keccak[:4]

    return 
}



func getEnv() (client *ethclient.Client, privateResource *ecdsa.PrivateKey, err error) {

    network, _ := beego.AppConfig.String("network::polygon")
    client, err = ethclient.Dial(network)
    
    if err != nil {
        log.Fatalf("Failed to connect to the Ethereum client: %v", err)
    }
    defer client.Close()

    
    // //账户
    // account := common.HexToAddress(address)
    
    // //合约
    // balance, err := client.BalanceAt(context.Background(), account, nil)
    // if err != nil {
    //   log.Fatal(err)
    // }

    // fmt.Printf("\nbalance is %v\n", balance)

    // 获取私钥
    privateKey, _ := beego.AppConfig.String("privateKey")
    privateResource, err = crypto.HexToECDSA(privateKey)
    if err != nil {
      log.Fatal(err)
    }   
   
    return

}





