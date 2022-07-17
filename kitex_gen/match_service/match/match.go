// Code generated by Kitex v0.3.4. DO NOT EDIT.

package match

import (
	"context"
	"github.com/ZeroAC/rpc_match/kitex_gen/match_service"
	"github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

func serviceInfo() *kitex.ServiceInfo {
	return matchServiceInfo
}

var matchServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "Match"
	handlerType := (*match_service.Match)(nil)
	methods := map[string]kitex.MethodInfo{
		"add_user":    kitex.NewMethodInfo(addUserHandler, newMatchAddUserArgs, newMatchAddUserResult, false),
		"remove_user": kitex.NewMethodInfo(removeUserHandler, newMatchRemoveUserArgs, newMatchRemoveUserResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "match_service",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.3.4",
		Extra:           extra,
	}
	return svcInfo
}

func addUserHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*match_service.MatchAddUserArgs)
	realResult := result.(*match_service.MatchAddUserResult)
	success, err := handler.(match_service.Match).AddUser(ctx, realArg.User, realArg.Info)
	if err != nil {
		return err
	}
	realResult.Success = &success
	return nil
}
func newMatchAddUserArgs() interface{} {
	return match_service.NewMatchAddUserArgs()
}

func newMatchAddUserResult() interface{} {
	return match_service.NewMatchAddUserResult()
}

func removeUserHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*match_service.MatchRemoveUserArgs)
	realResult := result.(*match_service.MatchRemoveUserResult)
	success, err := handler.(match_service.Match).RemoveUser(ctx, realArg.User, realArg.Info)
	if err != nil {
		return err
	}
	realResult.Success = &success
	return nil
}
func newMatchRemoveUserArgs() interface{} {
	return match_service.NewMatchRemoveUserArgs()
}

func newMatchRemoveUserResult() interface{} {
	return match_service.NewMatchRemoveUserResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) AddUser(ctx context.Context, user *match_service.User, info string) (r int32, err error) {
	var _args match_service.MatchAddUserArgs
	_args.User = user
	_args.Info = info
	var _result match_service.MatchAddUserResult
	if err = p.c.Call(ctx, "add_user", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) RemoveUser(ctx context.Context, user *match_service.User, info string) (r int32, err error) {
	var _args match_service.MatchRemoveUserArgs
	_args.User = user
	_args.Info = info
	var _result match_service.MatchRemoveUserResult
	if err = p.c.Call(ctx, "remove_user", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
