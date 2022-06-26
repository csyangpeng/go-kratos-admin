package service

import (
	v1 "github.com/csyangpeng/go-kratos-admin/api/center/interface/v1"
	"github.com/csyangpeng/go-kratos-admin/app/center/interface/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewCenterInterface)

type CenterInterface struct {
	v1.UnimplementedCenterInterfaceServer

	ac  *biz.AuthUseCase
	uc  *biz.UserUseCase
	log *log.Helper
}

func NewCenterInterface(ac *biz.AuthUseCase, uc *biz.UserUseCase, logger log.Logger) *CenterInterface {
	return &CenterInterface{
		ac:  ac,
		uc:  uc,
		log: log.NewHelper(log.With(logger, "module", "service/interface")),
	}
}
