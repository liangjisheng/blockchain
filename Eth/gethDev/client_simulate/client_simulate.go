package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

func clientSimulate() {
	// 我们需要一个带有初始ETH的账户
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	// accounts/abi/bind包创建一个NewKeyedTransactor，并为其传递私钥
	auth := bind.NewKeyedTransactor(privateKey)

	// 创建一个创世账户并为其分配初始余额。我们将使用core包的GenesisAccount类型
	balance := new(big.Int)
	balance.SetString("10000000000000000000", 10) // 10 eth in wei

	address := auth.From
	genesisAlloc := map[common.Address]core.GenesisAccount{
		address: {
			Balance: balance,
		},
	}

	// 现在我们将创世分配结构体和配置好的汽油上限传给account/abi/bind/backends
	// 包的NewSimulatedBackend方法，该方法将返回一个新的模拟以太坊客户端
	blockGasLimit := uint64(4712388)
	client := backends.NewSimulatedBackend(genesisAlloc, blockGasLimit)

	// 您可以像往常一样使用此客户端。作为一个示例，我们将构造一个新的交易并进行广播
	fromAddress := auth.From
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	value := big.NewInt(1000000000000000000) // in wei (1 eth)
	gasLimit := uint64(21000)                // in units
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	toAddress := common.HexToAddress("0xd65E7bBF1122753F023FB8760008Dd371763Ac49")
	var data []byte
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

	chainID := big.NewInt(1)
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s\n", signedTx.Hash().Hex())

	// 到现在为止，您可能想知道交易何时才会被开采。为了“开采”它，您还必须做
	// 一件额外的事情，在客户端调用Commit提交新开采的区块
	client.Commit()

	// 现在您可以获取交易收据并看见其已被处理
	receipt, err := client.TransactionReceipt(context.Background(), signedTx.Hash())
	if err != nil {
		log.Fatal(err)
	}
	if receipt == nil {
		log.Fatal("receipt is nil. Forgot to commit?")
	}

	fmt.Printf("status: %v\n", receipt.Status) // status: 1

	// 因此，请记住：模拟客户端允许您使用模拟客户端的Commit方法手动开采区块
}
