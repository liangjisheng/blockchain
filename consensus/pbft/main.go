package main

import (
	"log"
	"os"
)

const nodeCount = 4

// 客户端的监听地址
var clientAddr = "127.0.0.1:8080"

// 节点池 主要用来存储监听地址
var nodeTable map[string]string

func main() {
	// 为四个节点生成公私钥
	genRsaKeys(1024)
	nodeTable = map[string]string{
		"N0": "127.0.0.1:8000",
		"N1": "127.0.0.1:8001",
		"N2": "127.0.0.1:8002",
		"N3": "127.0.0.1:8003",
	}
	if len(os.Args) != 2 {
		log.Panic("输入的参数有误！")
	}
}
