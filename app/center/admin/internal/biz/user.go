package biz

import (
	"context"
	"errors"
	v1 "github.com/csyangpeng/go-kratos-admin/api/center/admin/v1"
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
	GetUser(ctx context.Context, id int64) (*User, error)
	FindByUsername(ctx context.Context, username string) (*User, error)
	VerifyPassword(ctx context.Context, u *User, password string) error
	ListUser(ctx context.Context, req *v1.ListUserReq) (*v1.ListUserReply, error)
}

type UserUseCase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUseCase(repo UserRepo, logger log.Logger) *UserUseCase {
	l := log.NewHelper(log.With(logger, "module", "usecase/admin"))
	return &UserUseCase{
		repo: repo,
		log:  l,
	}
}

func (uc *UserUseCase) Get(ctx context.Context, req *v1.GetUserReq) (*v1.GetUserReply, error) {
	user, err := uc.repo.GetUser(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &v1.GetUserReply{
		Id:       user.Id,
		Username: user.Username,
	}, nil
}

func (uc *UserUseCase) List(ctx context.Context, req *v1.ListUserReq) (*v1.ListUserReply, error) {
	return uc.repo.ListUser(ctx, req)
}
