package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
)

func account() {
	address := common.HexToAddress("0xBC08c485Ee0daC03a30cA9bEbdd59CbB98E5A09c")

	fmt.Println(address.Hex())
	fmt.Println(address.Hash().Hex())
	fmt.Println(address.Bytes())
}
