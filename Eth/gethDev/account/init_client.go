package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

// 若您没有现有以太坊客户端，您可以连接到infura网关.Infura管理着一批安全，
// 可靠，可扩展的以太坊[geth和parity]节点，并且在接入以太坊网络时降低了
// 新人的入门门槛

// 若您运行了本地geth实例，您还可以将路径传递给IPC端点文件
// client, err := ethclient.Dial("/home/user/.ethereum/geth.ipc")

// 使用Ganache
// npm install -g ganache-cli
// ganache-cli
// 现在连到http://localhost:8584上的ganache RPC主机
// client, err := ethclient.Dial("http://localhost:8545")

func initClient() {
	client, err := ethclient.Dial("https://mainnet.infura.io")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("we have a connection")
	_ = client // we'll use this in the upcoming sections
}
