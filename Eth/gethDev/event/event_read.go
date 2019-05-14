package main

import (
	"blockchain/Eth/gethDev/contract/store"
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func eventRead() {
	client, err := ethclient.Dial("wss://ropsten.infura.io/ws")
	if err != nil {
		log.Fatal(err)
	}

	// 智能合约可以可选地释放“事件”，其作为交易收据的一部分存储日志。
	// 读取这些事件相当简单。首先我们需要构造一个过滤查询
	contractAddress := common.HexToAddress("0x062549Da5B962f82574DfC1380994De4Fc8F0f16")
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	// FilterLogs，它接收我们的查询并将返回所有的匹配事件日志
	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	// 返回的所有日志将是ABI编码，因此它们本身不会非常易读。为了解码日志，
	// 我们需要导入我们智能合约的ABI
	// abi.JSON函数返回一个我们可以在Go应用程序中使用的解析过的ABI接口
	contractAbi, err := abi.JSON(strings.NewReader(string(store.StoreABI)))
	if err != nil {
		log.Fatal(err)
	}

	for _, vLog := range logs {
		fmt.Println(vLog.BlockHash.Hex())
		fmt.Println(vLog.BlockNumber)
		fmt.Println(vLog.TxHash.Hex())

		// 创建一个匿名结构体，并将指针作为第一个参数传递给解析后的ABI接口的
		// Unpack函数，以解码原始的日志数据。第二个参数是我们尝试解码的事件名称
		// 最后一个参数是编码的日志数据
		event := struct {
			Key   [32]byte
			Value [32]byte
		}{}
		err := contractAbi.Unpack(&event, "ItemSet", vLog.Data)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(event.Key[:]))
		fmt.Println(string(event.Value[:]))

		// 若您的solidity事件包含indexed事件类型，那么它们将成为主题而不是日志
		// 的数据属性的一部分。在solidity中您最多只能有4个主题，但只有3个可索引
		// 的事件类型。第一个主题总是事件的签名
		var topics [4]string
		for i := range vLog.Topics {
			topics[i] = vLog.Topics[i].Hex()
		}

		fmt.Println(topics[0])
	}

	// 首个主题只是被哈希过的事件签名
	eventSignature := []byte("ItemSet(bytes32,bytes32)")
	hash := crypto.Keccak256Hash(eventSignature)
	fmt.Println(hash.Hex())
}
