package main

import (
	"context"
	"fmt"
	"time"

	"github.com/ZeroAC/rpc_match/kitex_gen/match_service"
	"github.com/ZeroAC/rpc_match/src/match_core"
	"github.com/ZeroAC/rpc_match/src/myqueue"
)

// MatchImpl implements the last service interface defined in the IDL.
type MatchImpl struct{}

//匹配池 完成匹配操作
var matchPool match_core.MatchPool

//消息队列
var messageQueue match_core.MessageQueue

//消费者不断处理消息队列中的消息
func ConsumeTask() {
	matchPool.Users = make([]*match_service.User, 0)
	messageQueue.Q = myqueue.NewQueue[*match_core.Task]()
	for {
		messageQueue.Mu.Lock()
		if messageQueue.Q.IsEmpty() {
			fmt.Println("Pending...")
			messageQueue.Mu.Unlock()
			time.Sleep(time.Second)
		} else {
			a := messageQueue.Q.Front()
			messageQueue.Q.Dequeue()
			messageQueue.Mu.Unlock()
			if a.FuncName == "add" {
				matchPool.AddUser(a.User)
			} else if a.FuncName == "remove" {
				matchPool.RemoveUser(a.User.Id)
			}
			matchPool.MatchProcess()
		}
	}
}

//以下两个为生产者
// AddUser implements the MatchImpl interface.
func (s *MatchImpl) AddUser(ctx context.Context, user *match_service.User, info string) (resp int32, err error) {
	fmt.Printf("add user %v\n", user)
	messageQueue.Mu.Lock()
	messageQueue.AddUser(user)
	messageQueue.Mu.Unlock()
	return
}

// RemoveUser implements the MatchImpl interface.
func (s *MatchImpl) RemoveUser(ctx context.Context, user *match_service.User, info string) (resp int32, err error) {
	fmt.Printf("remove user %v\n", user.Id)
	messageQueue.Mu.Lock()
	defer messageQueue.Mu.Unlock()
	messageQueue.RemoveUser(user)
	return
}
