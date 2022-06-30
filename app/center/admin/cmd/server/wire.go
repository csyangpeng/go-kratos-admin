//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/csyangpeng/go-kratos-admin/app/center/admin/internal/biz"
	"github.com/csyangpeng/go-kratos-admin/app/center/admin/internal/conf"
	"github.com/csyangpeng/go-kratos-admin/app/center/admin/internal/data"
	"github.com/csyangpeng/go-kratos-admin/app/center/admin/internal/server"
	"github.com/csyangpeng/go-kratos-admin/app/center/admin/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, *conf.Auth, *conf.Registry, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
