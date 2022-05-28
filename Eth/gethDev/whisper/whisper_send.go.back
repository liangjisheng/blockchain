package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/whisper/whisperv6"

	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/ethereum/go-ethereum/whisper/shhclient"
)

func whisperSend() {
	client, err := shhclient.Dial("ws://127.0.0.1:8546")
	if err != nil {
		log.Fatal(err)
	}

	keyID, err := client.NewKeyPair(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(keyID) // faf494a1eec137019e210fddedacde7fa3647d4d5effb3388effe1daddaec699

	// 调用PublicKey函数以字节格式读取密钥对的公钥，我们将使用它来加密消息
	publicKey, err := client.PublicKey(context.Background(), keyID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(hexutil.Encode(publicKey)) // 0x04aa164ea6afd635d2d4d386be9f1f40a214713b2abbb3262d5fff8b5a4058077890a64a91fde3aec3b3f605592b817bc510651451fe3255be790acf79254680d4

	// 现在我们将通过从go-ethereumwhisper/whisperv6包中初始化NewMessage结构来
	// 构造我们的私语消息，这需要以下属性：
	// Payload 字节格式的消息内容
	// PublicKey 加密的公钥
	// TTL 消息的活跃时间
	// PowTime 做工证明的时间上限
	// PowTarget 做工证明的时间下限
	message := whisperv6.NewMessage{
		Payload:   []byte("hello"),
		PublicKey: publicKey,
		TTL:       60,
		PowTime:   2,
		PowTarget: 2.5,
	}

	// 返回消息的哈希值
	messageHash, err := client.Post(context.Background(), message)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messageHash) // 0xff46f778640150cd8f053cafdec4c9e2b076861f975df967b11b11b8a108a9c0
}
