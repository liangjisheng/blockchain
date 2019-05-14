package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func eventSubscribe() {
	// 为了订阅事件日志，我们需要做的第一件事就是拨打启用websocket的以太坊
	// 客户端。 幸运的是，Infura支持websockets
	// client, err := ethclient.Dial("wss://rinkeby.infura.io/ws")
	client, err := ethclient.Dial("wss://ropsten.infura.io/ws")
	if err != nil {
		log.Fatal(err)
	}

	// 创建筛选查询
	contractAddress := common.HexToAddress("0x062549Da5B962f82574DfC1380994De4Fc8F0f16")
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	// 使用select语句设置一个连续循环来读入新的日志事件或订阅错误
	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			fmt.Println(vLog) // pointer to event log
		}
	}
}
