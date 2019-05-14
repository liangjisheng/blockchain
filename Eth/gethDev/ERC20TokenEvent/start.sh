#!/bin/bash

# 创建ERC-20智能合约的事件日志的interface文件 erc20.sol

# 然后在给定abi使用abigen创建Go包
# solc --abi erc20.sol | sed -n '4,$p' > erc20.abi

PKG_NAME="token"
mkdir "$PKG_NAME"
if [ -f $GOBIN/abigen ]
then
    $GOBIN/abigen --abi=erc20.abi --pkg="$PKG_NAME" --out="$PKG_NAME"/"$PKG_NAME".go
fi
