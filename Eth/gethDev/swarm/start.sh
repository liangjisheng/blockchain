#!/bin/bash

# 安装geth和bzzd
go get -d github.com/ethereum/go-ethereum
go install github.com/ethereum/go-ethereum/cmd/geth
go install github.com/ethereum/go-ethereum/cmd/swarm

cd $GOBIN

# 生成一个新的geth帐户
./geth account new

# Public address of the key:   0x1b2D389DB8565ed8B9B3b73f1aAc95DAbDD5f2d8
# Path of the secret key file: /root/.ethereum/keystore/UTC--2019-05-14T08-27-18.747420850Z--1b2d389db8565ed8b9b3b73f1aac95dabdd5f2d8

# 将环境变量BZZKEY导出，并设定为我们刚刚生成的geth帐户地址
export BZZKEY=1b2D389DB8565ed8B9B3b73f1aAc95DAbDD5f2d8

# 然后使用设定的帐户运行swarm，并作为我们的swarm帐户。 默认情况下，Swarm将在端口“8500”上运行
./swarm --bzzaccount $BZZKEY
