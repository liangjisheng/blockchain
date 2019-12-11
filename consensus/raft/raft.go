package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Raft ...
type Raft struct {
	node               *NodeInfo  // 节点信息
	vote               int        // 本节点获得的投票数
	lock               sync.Mutex // 线程锁
	me                 string     // 节点编号
	currentTerm        int        // 当前任期
	votedFor           string     // 为哪个节点投票
	state              int        // 当前节点状态 0 follower  1 candidate  2 leader
	lastMessageTime    int64      // 发送最后一条消息的时间
	lastHeartBeartTime int64      // 发送最后一条心跳消息的时间
	currentLeader      string     // 当前节点的领导
	timeout            int        // 心跳超时时间(单位：秒)
	voteCh             chan bool  // 接收投票成功通道
	heartBeat          chan bool  // 心跳信号
}

// NodeInfo ...
type NodeInfo struct {
	ID   string
	Port string
}

// Message ...
type Message struct {
	Msg   string
	MsgID int
}

// NewRaft ...
func NewRaft(id, port string) *Raft {
	node := new(NodeInfo)
	node.ID = id
	node.Port = port

	rf := new(Raft)
	rf.node = node
	rf.setVote(0)             // 当前节点获得票数
	rf.me = id                // 编号
	rf.setVoteFor("-1")       // 给0 1 2三个节点投票 给谁都不投
	rf.setStatus(0)           // 0 follower
	rf.lastHeartBeartTime = 0 // 最后一次心跳检测时间
	rf.timeout = heartBeatTimeout
	rf.setCurrentLeader("-1")      // 最初没有领导
	rf.setTerm(0)                  // 设置任期
	rf.voteCh = make(chan bool)    // 投票通道
	rf.heartBeat = make(chan bool) // 心跳通道

	return rf
}

// 修改节点为候选人状态
func (rf *Raft) becomeCandidate() bool {
	r := randRange(1500, 5000)
	// 休眠随机时间后，再开始成为候选人
	time.Sleep(time.Duration(r) * time.Millisecond)
	// 如果发现本节点已经投过票，或者已经存在领导者，则不用变身候选人状态
	if rf.state == 0 && rf.currentLeader == "-1" && rf.votedFor == "-1" {
		rf.setStatus(1)                // 将节点状态变为1
		rf.setVoteFor(rf.me)           // 设置为哪个节点投票
		rf.setTerm(rf.currentTerm + 1) // 节点任期加1
		rf.setCurrentLeader("-1")      // 当前没有领导
		rf.voteAdd()                   // 为自己投票
		fmt.Println("本节点已变更为候选人状态")
		fmt.Printf("当前得票数：%d\n", rf.vote)
		// 开启选举通道
		return true
	}
	return false
}

// 进行选举
func (rf *Raft) election() bool {
	fmt.Println("开始进行领导者选举，向其他节点进行广播")
	go rf.broadcast("Raft.Vote", rf.node, func(ok bool) {
		rf.voteCh <- ok
	})
	for {
		select {
		case <-time.After(time.Second * time.Duration(timeout)):
			fmt.Println("领导者选举超时，节点变更为追随者状态")
			rf.reDefault()
			return false
		case ok := <-rf.voteCh:
			if ok {
				rf.voteAdd()
				fmt.Printf("获得来自其他节点的投票，当前得票数：%d\n", rf.vote)
			}
			if rf.vote > raftCount/2 && rf.currentLeader == "-1" {
				fmt.Println("获得超过网络节点二分之一的得票数，本节点被选举成为了leader")
				// 节点状态变为2，代表leader
				rf.setStatus(2)
				// 当前领导者为自己
				rf.setCurrentLeader(rf.me)
				fmt.Println("向其他节点进行广播...")
				go rf.broadcast("Raft.ConfirmationLeader", rf.node, func(ok bool) {
					fmt.Println(ok)
				})
				// 开启心跳检测通道
				rf.heartBeat <- true
				return true
			}
		}
	}
}

// 心跳检测方法
func (rf *Raft) heartbeat() {
	// 如果收到通道开启的信息,将会向其他节点进行固定频率的心跳检测
	if <-rf.heartBeat {
		for {
			fmt.Println("本节点开始发送心跳检测...")
			rf.broadcast("Raft.HeartbeatRe", rf.node, func(ok bool) {
				fmt.Println("收到回复:", ok)
			})
			rf.lastHeartBeartTime = millisecond()
			time.Sleep(time.Second * time.Duration(heartBeatTimes))
		}
	}
}

func randRange(min, max int64) int64 {
	rand.Seed(time.Now().UnixNano())
	return rand.Int63n(max-min) + min
}

// 获取当前时间的毫秒数
func millisecond() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

// 设置任期
func (rf *Raft) setTerm(term int) {
	rf.lock.Lock()
	defer rf.lock.Unlock()
	rf.currentTerm = term
}

// 设置为谁投票
func (rf *Raft) setVoteFor(id string) {
	rf.lock.Lock()
	defer rf.lock.Unlock()
	rf.votedFor = id
}

// 设置当前领导者
func (rf *Raft) setCurrentLeader(leader string) {
	rf.lock.Lock()
	defer rf.lock.Unlock()
	rf.currentLeader = leader
}

// 设置当前节点状态
func (rf *Raft) setStatus(state int) {
	rf.lock.Lock()
	defer rf.lock.Unlock()
	rf.state = state
}

// 投票累加
func (rf *Raft) voteAdd() {
	rf.lock.Lock()
	defer rf.lock.Unlock()
	rf.vote++
}

// 设置投票数量
func (rf *Raft) setVote(num int) {
	rf.lock.Lock()
	defer rf.lock.Unlock()
	rf.vote = num
}

// 恢复默认设置
func (rf *Raft) reDefault() {
	rf.setVote(0)
	rf.setVoteFor("-1")
	rf.setStatus(0)
}
