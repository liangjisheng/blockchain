package address

import (
	"encoding/hex"
	"fmt"
	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
)

func Address() {
	curve := btcec.S256()
	priKey, err := btcec.NewPrivateKey(curve)
	if err != nil {
		panic(err)
	}
	fmt.Println("priKey:", hex.EncodeToString(priKey.Serialize()))

	pubKey := priKey.PubKey()
	fmt.Println("pubKey:", hex.EncodeToString(pubKey.SerializeUncompressed()))

	hash160 := btcutil.Hash160(pubKey.SerializeUncompressed())
	fmt.Println("hash160:", hex.EncodeToString(hash160))

	pubKeyHashAddress, err := btcutil.NewAddressPubKeyHash(hash160, &chaincfg.MainNetParams)
	address := pubKeyHashAddress.EncodeAddress()
	fmt.Println("address:", address)
}
