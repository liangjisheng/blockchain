package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func blockSubscribe() {
	// 在本节中，我们将讨论如何设置订阅以便在新区块被开采时获取事件。首先，
	// 我们需要一个支持websocket RPC的以太坊服务提供者。在示例中，我们
	// 将使用infura 的websocket端点
	client, err := ethclient.Dial("wss://ropsten.infura.io/ws")
	if err != nil {
		log.Fatal(err)
	}

	// 创建一个新的通道，用于接收最新的区块头
	headers := make(chan *types.Header)
	// 现在我们调用客户端的SubscribeNewHead方法，它接收我们刚创建的区块头通道
	// 该方法将返回一个订阅对象
	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatal(err)
	}

	// 订阅将推送新的区块头事件到我们的通道，因此我们可以使用一个select语句
	// 来监听新消息。订阅对象还包括一个error通道，该通道将在订阅失败时发送消息
	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case header := <-headers:
			fmt.Println(header.Hash().Hex())

			// 要获得该区块的完整内容，我们可以将区块头的摘要传递给客户端的BlockByHash函数
			block, err := client.BlockByHash(context.Background(), header.Hash())
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(block.Hash().Hex())
			fmt.Println(block.Number().Uint64())
			fmt.Println(block.Time())
			fmt.Println(block.Nonce())
			fmt.Println(len(block.Transactions()))
		}
	}
}
