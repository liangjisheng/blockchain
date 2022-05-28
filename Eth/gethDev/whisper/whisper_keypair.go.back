package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/whisper/shhclient"
)

func whisperKeypair() {
	client, err := shhclient.Dial("ws://127.0.0.1:8546")
	if err != nil {
		log.Fatal(err)
	}

	// 生成 Whisper 密匙对
	// 在Whisper中，消息必须使用对称或非对称密钥加密，以防止除预期接收者以外
	// 的任何人读取消息
	// 在连接到Whisper客户端后，您需要调用客户端的NewKeyPair方法来生成该节点
	// 将管理的新公共和私有对。 此函数的结果将是一个唯一的ID，它引用我们将在接
	// 下来的几节中用于加密和解密消息的密钥对
	keyID, err := client.NewKeyPair(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(keyID) // ed385add19349e260b36b257f64977c666276aa4598d041fda4b79bce81d5460
}
