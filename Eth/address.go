package main

import (
	"encoding/hex"
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	genEthAddr()
}

func genEthAddr() {
	key, _ := crypto.GenerateKey()
	strPrikey := "b38acf64ad3b4d0a53b2a513b73170330af3d09cacc3d01a39d0a852abbd6313"
	key.D.SetString(strPrikey, 16)

	privateKey := hex.EncodeToString(key.D.Bytes())
	strx := "9f3a050c7ec81b4450596f68d8251fab3accf3e8f2b6b07ea2a5b4a6bee02607"
	stry := "06148ad1f756d64bd66ec2ce576302cc9b91e88a088271300bd7b9994f9670aa"
	key.PublicKey.X.SetString(strx, 16)
	key.PublicKey.Y.SetString(stry, 16)
	address := crypto.PubkeyToAddress(key.PublicKey).Hex()

	fmt.Printf("privateKey: 0x%s\n", privateKey)
	fmt.Printf("addr: %s\n", address)
}
