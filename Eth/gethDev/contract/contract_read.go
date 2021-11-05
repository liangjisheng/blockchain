package main

import (
	"contract/store"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func contractRead() {
	// client, err := ethclient.Dial("https://rinkeby.infura.io")
	client, err := ethclient.Dial("https://ropsten.infura.io")
	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress("0x062549Da5B962f82574DfC1380994De4Fc8F0f16")
	instance, err := store.NewStore(address, client)
	if err != nil {
		log.Fatal(err)
	}

	// 如果你还记得我们在部署过程中设置的合约中有一个名为version的全局变量
	// 因为它是公开的，这意味着它们将成为我们自动创建的getter函数。常量和view
	// 函数也接受bind.CallOpts作为第一个参数。了解可用的具体选项要看相应类
	// 的文档 一般情况下我们可以用 nil
	version, err := instance.Version(nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(version)
}
