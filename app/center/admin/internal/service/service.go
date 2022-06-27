package service

import (
	v1 "github.com/csyangpeng/go-kratos-admin/api/center/admin/v1"
	"github.com/csyangpeng/go-kratos-admin/app/center/admin/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewCenterAdmin)

type CenterAdmin struct {
	v1.UnimplementedCenterAdminServer

	log *log.Helper
	uc  *biz.UserUseCase
	ac  *biz.AuthUseCase
}

func NewCenterAdmin(uc *biz.UserUseCase, ac *biz.AuthUseCase, logger log.Logger) *CenterAdmin {
	return &CenterAdmin{
		log: log.NewHelper(log.With(logger, "module", "service/admin")),
		uc:  uc,
		ac:  ac,
	}
}
