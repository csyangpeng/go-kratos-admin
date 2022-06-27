package data

import (
	"context"
	userv1 "github.com/csyangpeng/go-kratos-admin/api/user/service/v1"
	"github.com/csyangpeng/go-kratos-admin/app/center/admin/internal/conf"
	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	jwtv4 "github.com/golang-jwt/jwt/v4"
	"github.com/google/wire"
	consulAPI "github.com/hashicorp/consul/api"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	NewDiscovery,
	NewRegistrar,
	NewUserServiceClient,
	NewUserRepo,
)

// Data .
type Data struct {
	log *log.Helper
	uc  userv1.UserClient
}

// NewData .
func NewData(c *conf.Data, logger log.Logger, uc userv1.UserClient) (*Data, error) {
	l := log.NewHelper(log.With(logger, "module", "data"))
	return &Data{
		log: l,
		uc:  uc,
	}, nil
}

func NewDiscovery(conf *conf.Registry) registry.Discovery {
	c := consulAPI.DefaultConfig()
	c.Address = conf.Consul.Address
	c.Scheme = conf.Consul.Scheme
	cli, err := consulAPI.NewClient(c)
	if err != nil {
		panic(err)
	}
	r := consul.New(cli, consul.WithHealthCheck(false))
	return r
}

func NewRegistrar(conf *conf.Registry) registry.Registrar {
	c := consulAPI.DefaultConfig()
	c.Address = conf.Consul.Address
	c.Scheme = conf.Consul.Scheme
	cli, err := consulAPI.NewClient(c)
	if err != nil {
		panic(err)
	}
	r := consul.New(cli, consul.WithHealthCheck(false))
	return r
}

func NewUserServiceClient(r registry.Discovery, ac *conf.Auth) userv1.UserClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///gka.user.service"),
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(
			tracing.Client(),
			recovery.Recovery(),
			jwt.Client(func(token *jwtv4.Token) (interface{}, error) {
				return []byte(ac.ServiceKey), nil
			}, jwt.WithSigningMethod(jwtv4.SigningMethodHS256)),
		),
	)
	if err != nil {
		panic(err)
	}
	c := userv1.NewUserClient(conn)
	return c
}
