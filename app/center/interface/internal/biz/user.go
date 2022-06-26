package biz

import (
	"context"
	"errors"
	v1 "github.com/csyangpeng/go-kratos-admin/api/center/interface/v1"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

type User struct {
	Id       int64
	Username string
	Password string
}

type UserRepo interface {
	FindByUsername(ctx context.Context, username string) (*User, error)
	VerifyPassword(ctx context.Context, u *User, password string) error
	ListUser(ctx context.Context, req *v1.ListUserReq) (*v1.ListUserReply, error)
}

type UserUseCase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUseCase(repo UserRepo, logger log.Logger) *UserUseCase {
	return &UserUseCase{
		repo: repo,
		log:  log.NewHelper(log.With(logger, "module", "usecase/interface")),
	}
}

func (uc *UserUseCase) List(ctx context.Context, req *v1.ListUserReq) (*v1.ListUserReply, error) {
	return uc.repo.ListUser(ctx, req)
}
