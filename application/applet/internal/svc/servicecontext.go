package svc

import (
	"beyond-learn-go-zero/application/applet/internal/config"
	"beyond-learn-go-zero/application/user/rpc/user"
	"beyond-learn-go-zero/pkg/interceptors"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config   config.Config
	UserRPC  user.User
	BizRedis *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	// 自定义拦截器
	userRPC := zrpc.MustNewClient(c.UserRPC, zrpc.WithUnaryClientInterceptor(interceptors.ClientErrorInterceptor()))
	rds, err := redis.NewRedis(redis.RedisConf{
		Host: c.BizRedis.Host,
		Type: c.BizRedis.Type,
		Pass: c.BizRedis.Pass,
	})
	if err != nil {
		panic(err)
	}

	return &ServiceContext{
		Config:   c,
		UserRPC:  user.NewUser(userRPC),
		BizRedis: rds,
	}
}
