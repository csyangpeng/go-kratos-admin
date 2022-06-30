package biz

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"math/rand"
)

var (
	ErrUserNotFound    = errors.New("user not found")
	ErrUsernameInvalid = errors.New("username invalid")
	ErrPasswordInvalid = errors.New("password invalid")
)

type User struct {
	Id       int64
	Username string
	Password string
	IsActive bool
}

func NewUser(username, password string) (User, error) {
	if len(username) <= 0 {
		return User{}, ErrUsernameInvalid
	}

	if len(password) <= 0 {
		return User{}, ErrPasswordInvalid
	}

	return User{
		Id:       rand.Int63(),
		Username: username,
		Password: password,
	}, nil
}

type UserRepo interface {
	FindByUsername(ctx context.Context, username string) (*User, error)
	VerifyPassword(ctx context.Context, u *User, password string) error
	Save(ctx context.Context, u *User) error
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
