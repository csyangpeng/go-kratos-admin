// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/csyangpeng/go-kratos-admin/app/center/interface/internal/biz"
	"github.com/csyangpeng/go-kratos-admin/app/center/interface/internal/conf"
	"github.com/csyangpeng/go-kratos-admin/app/center/interface/internal/data"
	"github.com/csyangpeng/go-kratos-admin/app/center/interface/internal/server"
	"github.com/csyangpeng/go-kratos-admin/app/center/interface/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, registry *conf.Registry, confData *conf.Data, auth *conf.Auth, logger log.Logger) (*kratos.App, func(), error) {
	discovery := data.NewDiscovery(registry)
	userClient := data.NewUserServiceClient(auth, discovery)
	dataData, err := data.NewData(confData, logger, userClient)
	if err != nil {
		return nil, nil, err
	}
	userRepo := data.NewUserRepo(dataData, logger)
	authUseCase := biz.NewAuthUseCase(auth, userRepo)
	userUseCase := biz.NewUserUseCase(userRepo, logger)
	centerInterface := service.NewCenterInterface(authUseCase, userUseCase, logger)
	httpServer := server.NewHTTPServer(confServer, auth, centerInterface, logger)
	registrar := data.NewRegistrar(registry)
	app := newApp(logger, httpServer, registrar)
	return app, func() {
	}, nil
}
