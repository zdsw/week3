package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"zg4-1/src/rk/week3/week3api/internal/config"
	"zg4-1/src/rk/week3/week3rpc/week3rpcclient"
)

type ServiceContext struct {
	Config      config.Config
	Week3Server week3rpcclient.Week3rpc
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		Week3Server: week3rpcclient.NewWeek3rpc(zrpc.MustNewClient(c.Week3Server)),
	}
}
