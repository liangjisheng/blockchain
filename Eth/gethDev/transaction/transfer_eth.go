package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func transferEth() {
	// client, err := ethclient.Dial("https://mainnet.infura.io")
	// client, err := ethclient.Dial("https://rinkeby.infura.io")
	client, err := ethclient.Dial("https://ropsten.infura.io")
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA("be2390e9eede3f4112c4ada66b6ef96ee9e63f511659f4f93e53624a30a24602")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// 之后我们需要获得帐户的随机数(nonce)。 每笔交易都需要一个nonce
	// 根据定义，nonce是仅使用一次的数字。 如果是发送交易的新帐户，
	// 则该随机数将为“0”。 来自帐户的每个新事务都必须具有前一个nonce增加1的
	// nonce。很难对所有nonce进行手动跟踪，于是ethereum客户端提供一个帮助
	// 方法PendingNonceAt，它将返回你应该使用的下一个nonce
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	value := big.NewInt(10000000000000000) // in wei (0.01 eth)
	// ETH转账的燃气应设上限为“21000”单位
	gasLimit := uint64(21000) // in units
	// 燃气价格总是根据市场需求和用户愿意支付的价格而波动的，因此对燃气价格
	// 进行硬编码有时并不理想。 go-ethereum客户端提供SuggestGasPrice函数，
	// 用于根据'x'个先前块来获得平均燃气价格
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	toAddress := common.HexToAddress("0xd65E7bBF1122753F023FB8760008Dd371763Ac49")
	var data []byte
	// 现在我们最终可以通过导入go-ethereumcore/types包并调用NewTransaction
	// 来生成我们的未签名以太坊事务，这个函数需要接收nonce，地址，值，燃气上限值，
	// 燃气价格和可选发的数据。 发送ETH的数据字段为“nil”。 在与智能合约进行交互时，
	// 我们将使用数据字段，仅仅转账以太币是不需要数据字段的
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

	// 下一步是使用发件人的私钥对事务进行签名。 为此，我们调用SignTx方法，
	// 该方法接受一个未签名的事务和我们之前构造的私钥。 SignTx方法需要EIP155
	// 签名者，这个也需要我们先从客户端拿到链ID
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s\n", signedTx.Hash().Hex())
}
