// Code generated by Kitex v0.3.4. DO NOT EDIT.

package match

import (
	"context"
	"github.com/ZeroAC/rpc_match/kitex_gen/match_service"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	AddUser(ctx context.Context, user *match_service.User, info string, callOptions ...callopt.Option) (r int32, err error)
	RemoveUser(ctx context.Context, user *match_service.User, info string, callOptions ...callopt.Option) (r int32, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kMatchClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kMatchClient struct {
	*kClient
}

func (p *kMatchClient) AddUser(ctx context.Context, user *match_service.User, info string, callOptions ...callopt.Option) (r int32, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AddUser(ctx, user, info)
}

func (p *kMatchClient) RemoveUser(ctx context.Context, user *match_service.User, info string, callOptions ...callopt.Option) (r int32, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.RemoveUser(ctx, user, info)
}
