package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func getTransactions() {
	client, err := ethclient.Dial("https://mainnet.infura.io")
	if err != nil {
		log.Fatal(err)
	}

	blockNumber := big.NewInt(5671744)
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	// 调用Transactions方法来读取块中的事务，该方法返回一个Transaction类型的列表
	// 然后，重复遍历集合并获取有关事务的任何信息就变得简单了
	for _, tx := range block.Transactions() {
		fmt.Println(tx.Hash().Hex())
		fmt.Println(tx.Value().String())
		fmt.Println(tx.Gas())
		fmt.Println(tx.GasPrice().Uint64())
		fmt.Println(tx.Nonce())
		fmt.Println(tx.Data())
		fmt.Println(tx.To().Hex())

		chainID, err := client.NetworkID(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		// 为了读取发送方的地址，我们需要在事务上调用AsMessage，它返回一个
		// Message类型，其中包含一个返回sender（from）地址的函数
		// AsMessage方法需要EIP155签名者，这个我们从客户端拿到链ID
		if msg, err := tx.AsMessage(types.NewEIP155Signer(chainID)); err == nil {
			fmt.Println(msg.From().Hex())
		}

		// 每个事务都有一个收据，其中包含执行事务的结果，例如任何返回值和日志
		// 以及为“1”（成功）或“0”（失败）的事件结果状态
		receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(receipt.Status)
	}

	// 在不获取块的情况下遍历事务的另一种方法是调用客户端的TransactionInBlock
	// 方法。 此方法仅接受块哈希和块内事务的索引值。 您可以调用TransactionCount
	// 来了解块中有多少个事务
	blockHash := common.HexToHash("0x9e8751ebb5069389b855bba72d94902cc385042661498a415979b7b6ee9ba4b9")
	count, err := client.TransactionCount(context.Background(), blockHash)
	if err != nil {
		log.Fatal(err)
	}

	for idx := uint(0); idx < count; idx++ {
		tx, err := client.TransactionInBlock(context.Background(), blockHash, idx)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(tx.Hash().Hex())
	}

	// 您还可以使用TransactionByHash在给定具体事务哈希值的情况下直接查询单个事务
	txHash := common.HexToHash("0x5d49fcaa394c97ec8a9c3e7bd9e8388d420fb050a52083ca52ff24b3b65bc9c2")
	tx, isPending, err := client.TransactionByHash(context.Background(), txHash)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(tx.Hash().Hex())
	fmt.Println(isPending)
}
