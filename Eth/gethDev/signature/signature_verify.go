package main

import (
	"bytes"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func signatureVerity() {
	privateKey, err := crypto.HexToECDSA("fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	fmt.Println(hex.EncodeToString(publicKeyBytes))

	data := []byte("hello")
	hash := crypto.Keccak256Hash(data) // 0x1c8aff950685c2ed4bc3174f3472287b56d9517b9c948127319a09a7a36deac8
	fmt.Println(hash.Hex())

	signature, err := crypto.Sign(hash.Bytes(), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(hexutil.Encode(signature)) // 0x789a80053e4927d0a898db8e065e948f5cf086e32f9ccaa54c1908e22ac430c62621578113ddbb62d509bf6049b8fb544ab06d36f916685a2eb8e57ffadde02301

	sigPublicKey, err := crypto.Ecrecover(hash.Bytes(), signature)
	if err != nil {
		log.Fatal(err)
	}

	matches := bytes.Equal(sigPublicKey, publicKeyBytes)
	fmt.Println(matches)

	sigPublicKeyECDSA, err := crypto.SigToPub(hash.Bytes(), signature)
	if err != nil {
		log.Fatal(err)
	}

	sigPublicKeyBytes := crypto.FromECDSAPub(sigPublicKeyECDSA)
	matches = bytes.Equal(sigPublicKeyBytes, publicKeyBytes)
	fmt.Println(matches)

	signatureNoRecoverID := signature[:len(signature)-1] // remove recovery id v
	verified := crypto.VerifySignature(publicKeyBytes, hash.Bytes(), signatureNoRecoverID)
	fmt.Println(verified)
}
