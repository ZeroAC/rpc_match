package main

import (
	"context"
	"fmt"

	"github.com/ZeroAC/rpc_match/kitex_gen/match_service"
	"github.com/ZeroAC/rpc_match/src/match_core"
)

// MatchImpl implements the last service interface defined in the IDL.
type MatchImpl struct{}

//匹配池 完成匹配操作
var matchPool = match_core.NewMatchPool()

//消息队列
var messageQueue = match_core.NewMessageQueue()

func ConsumeTask() {
	for ch := range messageQueue.CH {
		fmt.Printf("ch: %v\n", ch)
		messageQueue.Mu.Lock()
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

//以下两个为生产者
// AddUser implements the MatchImpl interface.
func (s *MatchImpl) AddUser(ctx context.Context, user *match_service.User, info string) (resp int32, err error) {
	fmt.Printf("add user %v\n", user)
	messageQueue.Mu.Lock()
	defer messageQueue.Mu.Unlock()
	messageQueue.AddUser(user)
	messageQueue.CH <- struct{}{}
	return
}

// RemoveUser implements the MatchImpl interface.
func (s *MatchImpl) RemoveUser(ctx context.Context, user *match_service.User, info string) (resp int32, err error) {
	fmt.Printf("remove user %v\n", user.Id)
	messageQueue.Mu.Lock()
	defer messageQueue.Mu.Unlock()
	messageQueue.RemoveUser(user)
	messageQueue.CH <- struct{}{}
	return
}
