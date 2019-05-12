package main

import (
	"blockchain/Eth/gethDev/contract/store"
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func contractDeploy() {
	// client, err := ethclient.Dial("https://rinkeby.infura.io")
	client, err := ethclient.Dial("https://ropsten.infura.io")
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA("be2390e9eede3f4112c4ada66b6ef96ee9e63f511659f4f93e53624a30a24602")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// 首先从go-ethereum导入accounts/abi/bind包，然后调用传入私钥的
	// NewKeyedTransactor。 然后设置通常的属性，如nonce，燃气价格，
	// 燃气上线限制和ETH值
	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	// 如果你还记得上个章节的内容, 我们创建了一个非常简单的“Store”合约，
	// 用于设置和存储键/值对。 生成的Go合约文件提供了部署方法。 部署方法
	// 名称始终以单词Deploy开头，后跟合约名称，在本例中为Store
	// deploy函数接受有密匙的事务处理器，ethclient，以及智能合约构造函数
	// 可能接受的任何输入参数。我们测试的智能合约接受一个版本号的字符串参数
	// 此函数将返回新部署的合约地址，事务对象，我们可以交互的合约实例，
	// 还有错误（如果有）
	input := "1.0"
	address, tx, instance, err := store.DeployStore(auth, client, input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(address.Hex())   // 0x062549Da5B962f82574DfC1380994De4Fc8F0f16
	fmt.Println(tx.Hash().Hex()) // 0xeb28ab4272a09d99135932154e7aa3429b0d35a15fa92fa6240eb143b1e6568e

	_ = instance
}
