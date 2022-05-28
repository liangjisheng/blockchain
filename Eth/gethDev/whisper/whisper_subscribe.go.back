package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/whisper/shhclient"
	"github.com/ethereum/go-ethereum/whisper/whisperv6"
)

func whisperSubscribe() {
	client, err := shhclient.Dial("ws://127.0.0.1:8546")
	if err != nil {
		log.Fatal(err)
	}

	keyID, err := client.NewKeyPair(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(keyID) // 06262769e352bdd1a99ff59f72b384d8774630f8f911f60c3139811b40fa2803

	// 我们将订阅websockets上的Whisper消息。 我们首先需要的是一个通道，
	// 它将从whisper/whisperv6包中的Message类型接收Whispe消息
	messages := make(chan *whisperv6.Message)
	// 在我们调用订阅之前，我们首先需要确定消息的过滤标准。 从whisperv6包中
	// 初始化一个新的Criteria对象。 由于我们只对定位到我们的消息感兴趣，
	// 因此我们将条件对象上的PrivateKeyID属性设置为我们用于加密消息的相同密钥ID
	criteria := whisperv6.Criteria{
		PrivateKeyID: keyID,
	}

	// 接下来，我们调用客户端的SubscribeMessages方法，该方法订阅符合给定条件
	// 的消息.HTTP不支持此方法; 仅支持双向连接，例如websockets和IPC
	// 最后一个参数是我们之前创建的消息通道
	sub, err := client.SubscribeMessages(context.Background(), criteria, messages)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for {
			select {
			case err := <-sub.Err():
				log.Fatal(err)
			case message := <-messages:
				// 消息内容在Payload属性中作为字节切片，我们可以将其转换回
				// 人类可读的字符串
				fmt.Println(string(message.Payload)) // hello
				os.Exit(0)
			}
		}
	}()

	publicKey, err := client.PublicKey(context.Background(), keyID)
	if err != nil {
		log.Print(err)
	}
	fmt.Println(hexutil.Encode(publicKey)) // 0x0460dfe8dd9bec1dc2457d0243c5e2aecc0d9e6b4e43ff05a27d85a86206a52dfa81446d90e41b370d3059d0c5753f9db378b659d93bef73b4162deace90ff6844

	message := whisperv6.NewMessage{
		Payload:   []byte("hello"),
		PublicKey: publicKey,
		TTL:       60,
		PowTime:   2,
		PowTarget: 2.5,
	}

	messageHash, err := client.Post(context.Background(), message)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messageHash) // 0x2bc551058032ad1741599534996111dde3e5333b1553bb137c165a9607a020f0

	runtime.Goexit() // wait for goroutines to finish
}
