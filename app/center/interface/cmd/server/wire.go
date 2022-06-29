//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/csyangpeng/go-kratos-admin/app/center/interface/internal/biz"
	"github.com/csyangpeng/go-kratos-admin/app/center/interface/internal/conf"
	"github.com/csyangpeng/go-kratos-admin/app/center/interface/internal/data"
	"github.com/csyangpeng/go-kratos-admin/app/center/interface/internal/server"
	"github.com/csyangpeng/go-kratos-admin/app/center/interface/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Registry, *conf.Data, *conf.Auth, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
