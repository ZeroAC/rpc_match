package match_core

import (
	"fmt"
	"sort"

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
func Abs(x int32) int32 {
	if x < 0 {
		return -x
	}
	return x
}
func (f *MatchPool) MatchProcess() {
	//更改匹配策略v1.0 分数差在100以内的进行匹配
	//排序后必定是相邻的匹配
	sort.Slice(f.Users, func(i, j int) bool {
		return f.Users[i].Score < f.Users[j].Score
	})
	delIndex := []int{} //匹配成功的下标 待删除
	for i := 1; i < len(f.Users); i++ {
		a, b := f.Users[i-1], f.Users[i]
		if Abs(a.Score-b.Score) <= 100 {
			fmt.Printf("%v %v matched!\n", a, b)
			delIndex = append(delIndex, i-1, i)
			i++
		}
	}

	newUser := make([]*match_service.User, 0)
	for i, j := 0, 0; i < len(f.Users); i++ {
		if j == len(delIndex) || i != delIndex[j] {
			newUser = append(newUser, f.Users[i])
		} else {
			j++
		}
	}
	f.Users = newUser //剔除掉匹配成功的人
}
