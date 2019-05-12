package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/rlp"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func txRawBytesSend() {
	// client, err := ethclient.Dial("https://mainnet.infura.io")
	// client, err := ethclient.Dial("https://rinkeby.infura.io")
	client, err := ethclient.Dial("https://ropsten.infura.io")
	if err != nil {
		log.Fatal(err)
	}

	rawTxHex := "f86a07843b9aca0082520894d65e7bbf1122753f023fb8760008dd371763ac49872386f26fc100008029a0f7c2a2b70b13dfe6e726fd1bd050077c0f23ba7a1ed68cf2bfd1821a34fd83d1a029c2c24c81fe30e7432c94441d7fc299a340bfd6806d8440791b428cfe4e1b9c"
	rawTxBytes, err := hex.DecodeString(rawTxHex)

	tx := new(types.Transaction)
	rlp.DecodeBytes(rawTxBytes, &tx)

	err = client.SendTransaction(context.Background(), tx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", tx.Hash().Hex())
}
