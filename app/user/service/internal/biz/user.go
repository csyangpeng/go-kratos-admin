package biz

import (
	"context"
	"errors"
	"math/rand"

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
	IsActive bool
}

type UserRepo interface {
	GetUser(ctx context.Context, id int64) (*User, error)
	FindByUsername(ctx context.Context, username string) (*User, error)
	CreateUser(ctx context.Context, u *User) (*User, error)
	VerifyPassword(ctx context.Context, u *User) (bool, error)
	ListUser(ctx context.Context, pageIndex, pageSize int) ([]*User, int, error)
	ChangeActive(ctx context.Context, u *User, isActive bool) (bool, error)
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
	uc.repo.CreateUser(ctx, u)
	if err != nil {
		return nil, err
	}

	return out, nil
}

func (uc *UserUseCase) Save(ctx context.Context, req *v1.SaveUserReq) (*v1.SaveUserReply, error) {
	user := &User{
		Id:       rand.Int63(),
		Username: req.Username,
		Password: req.Password,
	}
	res, err := uc.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	return &v1.SaveUserReply{Id: res.Id}, nil
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
		IsActive: user.IsActive,
	}, nil
}

func (uc *UserUseCase) List(ctx context.Context, pageIndex, pageSize int) ([]*User, int, error) {
	return uc.repo.ListUser(ctx, pageIndex, pageSize)
}

func (uc *UserUseCase) ChangeActive(ctx context.Context, req *v1.ChangeActiveReq) (*v1.ChangeActiveReply, error) {
	u, err := uc.repo.GetUser(ctx, req.Id)
	if err != nil || u == nil {
		return nil, ErrUserNotFound
	}

	ok, err := uc.repo.ChangeActive(ctx, &User{Id: u.Id, Username: u.Username}, req.IsActive)
	if err != nil {
		return nil, err
	}

	return &v1.ChangeActiveReply{
		Ok: ok,
		Id: req.Id,
	}, nil
}
