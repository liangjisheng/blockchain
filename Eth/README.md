# README

[eth-hd-wallet](https://github.com/miguelmota/go-ethereum-hdwallet)
[ethereum-mac-env](http://blog.hubwiz.com/2019/06/01/ethereum-mac-env/)
[ethereum-hot-wallet](http://blog.hubwiz.com/2019/06/25/ethereum-hot-wallet/)

可以使用 [remix](https://remix.ethereum.org/)来在线编译合约, 得到合约ABI和Bytecode
然后可以使用 abigen 生成包装好的 go 文件

安装 abigen

```shell
go install github.com/ethereum/go-ethereum/cmd/abigen
```

生成包装好的 go 文件

```shell
abigen --abi token.abi --pkg main --type Token --out token.go
```

[eip1559](https://metamask.io/1559/)

eth api

[api](https://eth.wiki/json-rpc/API)
[api](http://cw.hubwiz.com/card/c/ethereum-json-rpc-api/)
[api](https://learnblockchain.cn/docs/etherscan/index.html)

etherscan api

[docs](https://docs.etherscan.io/)
