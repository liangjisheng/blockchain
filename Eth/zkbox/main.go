package main

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"time"
	"zkbox.org/coachnft"
)

//immense once civil glass mad raccoon aspect sustain beef orbit proud math

func main() {
	client, err := ethclient.Dial("https://rinkeby.infura.io/v3/2fa35fe0ec5e42cebd9f4b37f9ba29dd")
	if err != nil {
		log.Println("Dial err:", err)
		return
	}
	defer client.Close()

	contractAddress := common.HexToAddress("0x2e85Dd915577fdd183ca1f0441316e97c3975D11")
	contract, err := coachnft.NewCoachnft(contractAddress, client)
	//_, err = coachnft.NewCoachnft(contractAddress, client)
	if err != nil {
		log.Println("NewCoachnft err:", err)
		return
	}

	privateKey, err := crypto.HexToECDSA("d5262805b5a55c6f90aa8abba83c82d45fec722d3f5c6023734fc52c0e060b45")
	if err != nil {
		log.Fatal(err)
	}
	//privateKeyBytes := crypto.FromECDSA(privateKey)
	//log.Println(hexutil.Encode(privateKeyBytes)[2:])

	//publicKey := privateKey.Public()
	//publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	//if !ok {
	//	log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	//}
	//
	//publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	//log.Println(hexutil.Encode(publicKeyBytes)[4:])
	//
	//address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	//log.Println("address:", address)

	fromAddress := common.HexToAddress("0xA8e9123885688AB63bF440bB9863A055652dC65E")
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Println("PendingNonceAt err:", err)
		return
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Println("SuggestGasPrice err:", err)
		return
	}

	chainID := big.NewInt(4)
	auth, _ := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(500000) // in units
	auth.GasPrice = gasPrice
	auth.From = fromAddress

	toAddress := common.HexToAddress("0x9f3ec47994Da2e6d1c378a02253365e60f1874AA")
	transaction, err := contract.MintToFans(auth, toAddress)
	if err != nil {
		log.Println("MintToFans err:", err)
		return
	}
	txHash := transaction.Hash()
	log.Println("txid:", txHash.String())

	tokenID := big.NewInt(-1)
	for {
		receipt, err := client.TransactionReceipt(context.Background(), txHash)
		if err != nil {
			time.Sleep(time.Second * 5)
		} else {
			tokenID = receipt.Logs[0].Topics[3].Big()
			log.Println("tokenID:", tokenID)
			break
		}
	}

	tokenURI, err := contract.TokenURI(&bind.CallOpts{}, tokenID)
	if err != nil {
		log.Println("TokenURI err:", err)
		return
	}
	log.Println("tokenURI:", tokenURI)
	//http://xj6-4.s.filfox.io:55002/ipfs/QmSqxz8vxeBxFypNNx2LxCHfZWA8E1KszxrzTLHTshREhK/0

	//baseURI, err := contract.BaseURI(&bind.CallOpts{})
	//if err != nil {
	//	log.Println("err:", err)
	//	return
	//}
	//log.Println("baseURI: ", baseURI)
	//baseURI: http://xj6-4.s.filfox.io:55002/ipfs/QmSqxz8vxeBxFypNNx2LxCHfZWA8E1KszxrzTLHTshREhK/

	//name, err := contract.Name(&bind.CallOpts{})
	//if err != nil {
	//	log.Println("err:", err)
	//	return
	//}
	//log.Println("name:", name)
	//name: CoachNFT

	//totalSupply, err := contract.TotalSupply(&bind.CallOpts{})
	//if err != nil {
	//	log.Println("err:", err)
	//	return
	//}
	//log.Println("total supply:", totalSupply.String())
	//total supply: 0

	//ok , err := contract.ClaimedFans(&bind.CallOpts{}, toAddress)
	//if err != nil {
	//	log.Println("ClaimedFans err:", err)
	//	return
	//}
	//log.Println("ok:", ok)
	//ok: false

	//tokenID := big.NewInt(0)
	//tokenIDAddress, err := contract.GetApproved(&bind.CallOpts{}, tokenID)
	//if err != nil {
	//	log.Println("GetApproved err:", err)
	//	return
	//}
	//log.Println("tokenIDAddress:", tokenIDAddress)
	//tokenIDAddress: 0x0000000000000000000000000000000000000000

	//index := big.NewInt(1)
	//tokenIndex, err := contract.TokenOfOwnerByIndex(&bind.CallOpts{}, toAddress, index)
	//if err != nil {
	//	log.Println("TokenOfOwnerByIndex err:", err)
	//	return
	//}
	//log.Println("tokenIndex:", tokenIndex)
	//tokenIndex: 0

	//zeroAddress := common.HexToAddress("0x0000000000000000000000000000000000000000")
	//index := big.NewInt(10000)
	//contract.FilterTransfer(&bind.FilterOpts{}, zeroAddress, toAddress, index)

	//symbol, err := contract.Symbol(&bind.CallOpts{})
	//if err != nil {
	//	log.Println("err:", err)
	//	return
	//}
	//log.Println("symbol:", symbol)
	//symbol: CoachNFT

	//owner, err := contract.Owner(&bind.CallOpts{})
	//if err != nil {
	//	log.Println("err:", err)
	//	return
	//}
	//log.Println("owner:", owner.String())
	//owner: 0xA8e9123885688AB63bF440bB9863A055652dC65E

	//contractURI, err := contract.ContractURI(&bind.CallOpts{})
	//if err != nil {
	//	log.Println("err:", err)
	//	return
	//}
	//log.Println("contractURI: ", contractURI)
	//contractURI: ipfs://QmV1SZzgaCWGvExViNRaRYg396NgEknp4cmYihcrJSPhKm
}
