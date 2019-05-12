#!/bin/bash

# 可以从一个solidity文件生成ABI
solc --abi store.sol | sed -n '4,$p' > store.abi

# 现在让我们用abigen将ABI转换为我们可以导入的Go文件。 这个新文件将包含
# 我们可以用来与Go应用程序中的智能合约进行交互的所有可用方法
mkdir store
if [ -f $GOBIN/abigen ]
then
    $GOBIN/abigen --abi=store.abi --pkg=store --out=store/store.go
fi

# 为了从Go部署智能合约，我们还需要将solidity智能合约编译为EVM字节码
# EVM字节码将在事务的数据字段中发送。 在Go文件上生成部署方法需要bin文件
solc --bin store.sol | sed -n '4,$p' > store.bin

# 现在我们编译Go合约文件，其中包括deploy方法，因为我们包含了bin文件
if [ -f $GOBIN/abigen ]
then
    $GOBIN/abigen --bin=store.bin --abi=store.abi --pkg=store --out=store/store.go
fi
