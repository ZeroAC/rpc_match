package match_core

import (
	"fmt"

	"github.com/ZeroAC/rpc_match/kitex_gen/match_service"
)

//匹配池 保存发出匹配请求的所有用户
type PoolUser struct {
	User      *match_service.User
	waitTimes int32
}
type MatchPool struct {
	Users []*PoolUser
}

func NewMatchPool() *MatchPool {
	ans := MatchPool{}
	ans.Users = make([]*PoolUser, 0)
	return &ans
}
func (f *MatchPool) AddUser(user *match_service.User) {
	f.Users = append(f.Users, &PoolUser{user, 0})
	fmt.Printf("add user %v\n", user)
}
func (f *MatchPool) RemoveByUserIndex(i int32) error {
	f.Users = append(f.Users[:i], f.Users[i+1:]...)
	return nil
}
func (f *MatchPool) RemoveUser(id int32) error {
	for i, v := range f.Users {
		if v.User.Id == id {
			f.RemoveByUserIndex(int32(i))
			fmt.Printf("remove user %v\n", id)
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
func Min(a, b int32) int32 {
	if a < b {
		return a
	}
	return b
}

//核心用户匹配算法
func (f *MatchPool) MatchProcess() {
	//更改匹配策略v2.0 每个用户随着匹配轮数的增加 能接受的分数阈值也变大
	//排序后必定是相邻的匹配
	delIndex := []int{} //匹配成功的下标 待删除
	mp := make(map[int]struct{})
	initScope := 100 //初始为100
	addScope := 50   //每轮一次加30
	for i := 0; i < len(f.Users); i++ {
		for j := i + 1; j < len(f.Users); j++ {
			if _, ok := mp[j]; ok { //该下标若已删除 则跳过
				continue
			}
			a, b := f.Users[i], f.Users[j]
			//若a和b都能相互接受 即二者的差值小于等于二者阈值的最小值
			if Abs(a.User.Score-b.User.Score) <= Min(a.waitTimes, b.waitTimes)*int32(addScope)+int32(initScope) {
				fmt.Printf("%v %v matched!!\n", a, b)
				delIndex = append(delIndex, i, j)
				mp[j] = struct{}{}
				break
			}
		}
	}

	newUser := make([]*PoolUser, 0)
	for i, j := 0, 0; i < len(f.Users); i++ {
		if j == len(delIndex) || i != delIndex[j] {
			newUser = append(newUser, f.Users[i])
		} else {
			j++
		}
	}
	f.Users = newUser //剔除掉匹配成功的人
	f.AddWaitTimes()  //为匹配成功的人 等待次数+1
}

//讲匹配池内的所有用户的等待次数加1
func (f *MatchPool) AddWaitTimes() {
	for i := 0; i < len(f.Users); i++ {
		f.Users[i].waitTimes++
	}
}
