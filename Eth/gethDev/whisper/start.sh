#!/bin/bash

# 要使用连接Whisper客户端，我们必须首先连接到运行whisper的以太坊节点
# 不幸的是，诸如infura之类的公共网关不支持whisper，因为没有金钱动力免费
# 处理这些消息.Infura可能会在不久的将来支持whisper，但现在我们必须运行
# 我们自己的geth节点。一旦你安装 geth, 运行geth的时候加 --shh flag
# 来支持whisper协议, 并且加 --wsflag和 --rpc，来支持websocket来接收实时信息

cd $GOBIN

./geth --rpc --shh --ws
