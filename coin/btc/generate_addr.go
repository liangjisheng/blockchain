package main

import (
	"encoding/hex"
	"fmt"

	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
)

// 生成比特币1开头的公私钥地址
func main() {
	curve := btcec.S256()
	privkey, err := btcec.NewPrivateKey(curve)
	if err != nil {
		panic(err)
	}
	fmt.Println("prikey:", hex.EncodeToString(privkey.Serialize()))

	pubkey := privkey.PubKey()
	fmt.Println("pubkey:", hex.EncodeToString(pubkey.SerializeUncompressed()))

	hash160 := btcutil.Hash160(pubkey.SerializeUncompressed())
	fmt.Println("hash160:", hex.EncodeToString(hash160))

	pubKeyHashAddress, err := btcutil.NewAddressPubKeyHash(hash160, &chaincfg.MainNetParams)
	address := pubKeyHashAddress.EncodeAddress()
	fmt.Println("address:", address)
}
