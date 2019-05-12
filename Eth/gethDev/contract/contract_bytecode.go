package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func contractBytecode() {
	// client, err := ethclient.Dial("https://rinkeby.infura.io")
	client, err := ethclient.Dial("https://ropsten.infura.io")
	if err != nil {
		log.Fatal(err)
	}

	// 有时您需要读取已部署的智能合约的字节码。 由于所有智能合约字节码都存在于
	// 区块链中，因此我们可以轻松获取它
	contractAddress := common.HexToAddress("0x062549Da5B962f82574DfC1380994De4Fc8F0f16")
	bytecode, err := client.CodeAt(context.Background(), contractAddress, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(hex.EncodeToString(bytecode))
}
