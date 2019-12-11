package main

import "sync"

// 本地消息池(模拟持久化层) 只有确认提交成功后才会存入此池
var localMessagePool = []Message{}

type node struct {
	nodeID     string // 节点ID
	addr       string // 节点监听地址
	rsaPrivKey []byte // RSA私钥
	rsaPubKey  []byte // RSA公钥
}

type pbft struct {
	node                node                       // 节点信息
	sequenceID          int                        // 每笔请求自增序号
	lock                sync.Mutex                 // 锁
	messagePool         map[string]Request         // 临时消息池，消息摘要对应消息本体
	prePareConfirmCount map[string]map[string]bool // 存放收到的prepare数量(至少需要收到并确认2f个)，根据摘要来对应
	commitConfirmCount  map[string]map[string]bool // 存放收到的commit数量（至少需要收到并确认2f+1个），根据摘要来对应
	isCommitBordcast    map[string]bool            // 该笔消息是否已进行Commit广播
	isReply             map[string]bool            // 该笔消息是否已对客户端进行Reply
}
