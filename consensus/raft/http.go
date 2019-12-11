package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"net/rpc"
)

func (rf *Raft) getRequest(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	// curl http://localhost:8080/req?message=ohmygod
	if len(r.Form["message"]) > 0 && rf.currentLeader != "-1" {
		message := r.Form["message"][0]
		m := new(Message)
		m.MsgID = getRandom()
		m.Msg = message
		// 接收到消息后，直接转发到领导者
		fmt.Println("http监听到了消息 准备发送给领导者，消息id:", m.MsgID)
		port := nodeTable[rf.currentLeader]
		rp, err := rpc.DialHTTP("tcp", "127.0.0.1"+port)
		if err != nil {
			log.Panic(err)
		}
		b := false
		err = rp.Call("Raft.LeaderReceiveMessage", m, &b)
		if err != nil {
			log.Panic(err)
		}
		fmt.Println("消息是否已发送到领导者:", b)
		w.Write([]byte("ok!!!"))
	}
}

func (rf *Raft) httpListen() {
	http.HandleFunc("/req", rf.getRequest)
	fmt.Println("监听8080")
	if err := http.ListenAndServe("127.0.0.1:8080", nil); err != nil {
		fmt.Println(err)
		return
	}
}

// 返回一个十位数的随机数 作为消息 id
func getRandom() int {
	x := big.NewInt(10000000000)
	for {
		result, err := rand.Int(rand.Reader, x)
		if err != nil {
			log.Panic(err)
		}
		if result.Int64() > 1000000000 {
			return int(result.Int64())
		}
	}
}
