package match_core

import (
	"sync"

	"github.com/ZeroAC/rpc_match/kitex_gen/match_service"
	"github.com/ZeroAC/rpc_match/src/myqueue"
)

//手写消息队列
//消息队列模型
//队列重点任务数据
type Task struct {
	User     *match_service.User
	FuncName string
}

type MessageQueue struct {
	Q  *myqueue.Queue[*Task]
	CH chan struct{} //控制并发量
	Mu sync.Mutex
}

func NewMessageQueue() *MessageQueue {
	ans := MessageQueue{}
	ans.Q = &myqueue.Queue[*Task]{}
	ans.CH = make(chan struct{}, 1000)
	return &ans
}

//向消息队列中推入一个添加用户的操作
func (f *MessageQueue) AddUser(user *match_service.User) {
	f.Q.Enqueue(&Task{user, "add"})
}

//向消息队列中推入一个删除用户的操作
func (f *MessageQueue) RemoveUser(user *match_service.User) {
	f.Q.Enqueue(&Task{user, "remove"})
}
