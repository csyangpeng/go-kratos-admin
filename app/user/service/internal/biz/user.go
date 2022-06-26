package biz

import (
	"context"
	"errors"
	v1 "github.com/csyangpeng/go-kratos-admin/api/user/service/v1"
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
	CreateUser(ctx context.Context, u *User) (*User, error)
	VerifyPassword(ctx context.Context, u *User) (bool, error)
}

type UserUseCase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUseCase(repo UserRepo, logger log.Logger) *UserUseCase {
	return &UserUseCase{
		repo: repo,
		log:  log.NewHelper(log.With(logger, "module", "usecase/user")),
	}
}

func (uc *UserUseCase) Get(ctx context.Context, id int64) (*User, error) {
	return uc.repo.GetUser(ctx, id)
}

func (uc *UserUseCase) Create(ctx context.Context, u *User) (*User, error) {
	out, err := uc.repo.CreateUser(ctx, u)
	if err != nil {
		return nil, err
	}

	return out, nil
}

func (uc *UserUseCase) Save(ctx context.Context, req *v1.SaveUserReq) (*v1.SaveUserReply, error) {
	user := &User{
		Username: req.Username,
		Password: req.Password,
	}
	_, err := uc.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	return &v1.SaveUserReply{Id: user.Id}, nil
}

func (uc *UserUseCase) VerifyPassword(ctx context.Context, u *User) (bool, error) {
	return uc.repo.VerifyPassword(ctx, u)
}

func (uc *UserUseCase) GetUserByUsername(ctx context.Context, req *v1.GetUserByUsernameReq) (*v1.GetUserByUsernameReply, error) {
	user, err := uc.repo.FindByUsername(ctx, req.Username)
	if err != nil {
		return nil, err
	}

	return &v1.GetUserByUsernameReply{
		Id:       user.Id,
		Username: user.Username,
	}, nil
}
