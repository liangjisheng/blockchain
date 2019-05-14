package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/whisper/shhclient"
)

func whisperClient() {
	// 我们将导入在whisper/shhclient中找到的go-ethereum whisper客户端软件包
	// 并初始化客户端，使用默认的websocket端口“8546”通过websockets连接我们的
	// 本地geth节点
	client, err := shhclient.Dial("ws://127.0.0.1:8546")
	if err != nil {
		log.Fatal(err)
	}

	_ = client // we'll be using this in the section.
	fmt.Println("we have a whisper connection")
}
