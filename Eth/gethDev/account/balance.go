package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func getBalance() {
	// 读取一个账户的余额相当简单。调用客户端的BalanceAt方法，给它传递账户
	// 地址和可选的区块号。将区块号设置为nil将返回最新的余额
	// client, err := ethclient.Dial("https://mainnet.infura.io")
	client, err := ethclient.Dial("https://ropsten.infura.io")
	if err != nil {
		log.Fatal(err)
	}

	// account := common.HexToAddress("0x71c7656ec7ab88b098defb751b7401b5f6d8976f")
	account := common.HexToAddress("0x8332fc0fA6ca0004Ed6D81a55599ad0dC43c5b92")
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(balance)

	// 传区块号能让您读取该区块时的账户余额
	blockNumber := big.NewInt(5552993)
	balanceAt, err := client.BalanceAt(context.Background(), account, blockNumber)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(balanceAt)

	fbalance := new(big.Float)
	fbalance.SetString(balanceAt.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	fmt.Println(ethValue)

	// 有时您想知道待处理的账户余额是多少，例如，在提交或等待交易确认后。客户端
	// 提供了类似BalanceAt的方法，名为PendingBalanceAt，它接收账户地址作为参数
	pendingBalance, err := client.PendingBalanceAt(context.Background(), account)
	fmt.Println(pendingBalance)
}
