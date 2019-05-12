package main

import (
	"blockchain/Eth/gethDev/contract/store"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func contractLoad() {
	// client, err := ethclient.Dial("https://rinkeby.infura.io")
	client, err := ethclient.Dial("https://ropsten.infura.io")
	if err != nil {
		log.Fatal(err)
	}

	// 一旦使用abigen工具将智能合约的ABI编译为Go包，下一步就是调用“New”方法，
	// 其格式为“New”，所以在我们的例子中如果你 回想一下它将是NewStore
	address := common.HexToAddress("0x062549Da5B962f82574DfC1380994De4Fc8F0f16")
	instance, err := store.NewStore(address, client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("contract is loaded")
	_ = instance
}
