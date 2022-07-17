package main

import (
	"context"
	"fmt"

	"github.com/ZeroAC/rpc_match/kitex_gen/match_service"
)

// MatchImpl implements the last service interface defined in the IDL.
type MatchImpl struct{}

// AddUser implements the MatchImpl interface.
func (s *MatchImpl) AddUser(ctx context.Context, user *match_service.User, info string) (resp int32, err error) {
	fmt.Printf("add user %v\n", user)
	return
}

// RemoveUser implements the MatchImpl interface.
func (s *MatchImpl) RemoveUser(ctx context.Context, user *match_service.User, info string) (resp int32, err error) {
	fmt.Printf("remove user %v\n", user.Id)
	return
}
