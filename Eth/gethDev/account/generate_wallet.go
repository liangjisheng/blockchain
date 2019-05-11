package main

import (
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/InWeCrypto/sha3"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func generateWallet() {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	privateKeyBytes := crypto.FromECDSA(privateKey)
	fmt.Println(hexutil.Encode(privateKeyBytes)[2:]) // 0x7a2c1cf99fcfcf82acce6a5c3f7e984278972d7b1a32046520042fe88aeedd24

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	fmt.Println(hexutil.Encode(publicKeyBytes)[4:]) // 0x0458b6fed06f94889bc88a6c095216388537fc48dcba0d074e0d7fda76e02d17aef07220cbef873d57475ea8a22c8f3c24d53eebfd819464b2664bbd38c6897321

	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Println(address) // 0xe1bb5C4cfC1ad63e592d81e57778C7021c5E8918

	hash := sha3.NewKeccak256()
	hash.Write(publicKeyBytes[1:])
	fmt.Println(hexutil.Encode(hash.Sum(nil)[12:])) // 0xe1bb5C4cfC1ad63e592d81e57778C7021c5E8918
}
