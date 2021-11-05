package main

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

func main() {
	//client, err := ethclient.Dial("https://rinkeby.infura.io/v3/2fa35fe0ec5e42cebd9f4b37f9ba29dd")
	client, err := ethclient.Dial("https://mainnet.eth.zks.app")
	if err != nil {
		log.Println("err:", err)
		return
	}
	defer client.Close()

	address := common.HexToAddress("0xb67cC7BFAf1945aa40ccECC3fA86d28b8cb3f5c0")
	nonce, err := client.PendingNonceAt(context.Background(), address)
	if err != nil {
		log.Println("PendingNonceAt err:", err)
		return
	}
	log.Println("nonce:", nonce)

	//txID := "0x3437efdff347e7d1ad7ec925c77ece351ba5b36572f3dcb82b9426cb3e67b8a2"
	//txHash := common.HexToHash(txID)
	//txInfo, isPending, err := client.TransactionByHash(context.Background(), txHash)
	//if err != nil {
	//	log.Println("TransactionByHash err:", err)
	//	return
	//}
	//log.Println("is pending:", isPending)
	//log.Println("type:", txInfo.Type())

	//height, err := client.BlockNumber(context.Background())
	//if err != nil {
	//	log.Println("err:", err)
	//	return
	//}
	//
	//log.Println("height:", height)

	//address := "0x0E78b8c07d6eeffADE621294673E34E2DB5C5553"
	//balance, err := client.BalanceAt(context.Background(), common.HexToAddress(address), nil)
	//if err != nil {
	//	log.Println("err:", err)
	//	return
	//}
	//
	//log.Println("balance:", balance.String())
}
