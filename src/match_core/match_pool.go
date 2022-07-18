package match_core

import (
	"fmt"

	"github.com/ZeroAC/rpc_match/kitex_gen/match_service"
)

//匹配池 保存发出匹配请求的所有用户
type MatchPool struct {
	Users []*match_service.User
}

func NewMatchPool() *MatchPool {
	ans := MatchPool{}
	ans.Users = make([]*match_service.User, 0)
	return &ans
}
func (f *MatchPool) AddUser(user *match_service.User) {
	f.Users = append(f.Users, user)

}
func (f *MatchPool) RemoveByUserIndex(i int32) error {

	f.Users = append(f.Users[:i], f.Users[i+1:]...)
	return nil
}
func (f *MatchPool) RemoveUser(id int32) error {
	for i, v := range f.Users {
		if v.Id == id {
			f.RemoveByUserIndex(int32(i))
			return nil
		}
	}
	return fmt.Errorf("not found id : %d", id)
}

//Matching algorithm v1.0
//Directly let the first two match directly
func (f *MatchPool) MatchProcess() {
	for len(f.Users) > 1 {
		a, b := f.Users[0], f.Users[1]
		f.RemoveUser(a.Id)
		f.RemoveUser(b.Id)
		fmt.Printf("%v %v\n matched!\n", a, b)
	}
}
