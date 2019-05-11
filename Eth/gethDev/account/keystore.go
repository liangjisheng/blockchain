package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/accounts/keystore"
)

// keystore是一个包含经过加密了的钱包私钥。go-ethereum中的keystore，每个文件
// 只能包含一个钱包密钥对。要生成keystore，首先您必须调用NewKeyStore，给它提供
// 保存keystore的目录路径。然后，您可调用NewAccount方法创建新的钱包，并给它传入
// 一个用于加密的口令。您每次调用NewAccount，它将在磁盘上生成新的keystore文件
func createKs() {
	ks := keystore.NewKeyStore("./wallets", keystore.StandardScryptN, keystore.StandardScryptP)
	password := "secret"
	account, err := ks.NewAccount(password)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(account.Address.Hex()) // 0x3f3043211f3051cDf61834a510c2C7c4b3b7A24d
}

// 现在要导入您的keystore，您基本上像往常一样再次调用NewKeyStore，然后调用
// Import方法，该方法接收keystore的JSON数据作为字节。第二个参数是用于加密私钥
// 的口令。第三个参数是指定一个新的加密口令，但我们在示例中使用一样的口令。导入
// 账户将允许您按期访问该账户，但它将生成新keystore文件！有两个相同的事物是没有
// 意义的，所以我们将删除旧的
func importKs() {
	file := "./wallets/UTC--2019-05-11T12-40-09.805691250Z--3f3043211f3051cdf61834a510c2c7c4b3b7a24d"
	ks := keystore.NewKeyStore("./wallets", keystore.StandardScryptN, keystore.StandardScryptP)
	jsonBytes, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	password := "secret"
	account, err := ks.Import(jsonBytes, password, password)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(account.Address.Hex()) // 0x3f3043211f3051cDf61834a510c2C7c4b3b7A24d

	if err := os.Remove(file); err != nil {
		log.Fatal(err)
	}
}
