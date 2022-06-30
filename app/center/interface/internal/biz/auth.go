package biz

import (
	"context"
	"errors"
	v1 "github.com/csyangpeng/go-kratos-admin/api/center/interface/v1"
	"github.com/csyangpeng/go-kratos-admin/app/center/interface/internal/conf"
	"github.com/golang-jwt/jwt/v4"
)

var (
	ErrLoginFailed    = errors.New("login failed")
	ErrPasswordFailed = errors.New("password failed")
)

type AuthUseCase struct {
	key      string
	userRepo UserRepo
}

func NewAuthUseCase(conf *conf.Auth, userRepo UserRepo) *AuthUseCase {
	return &AuthUseCase{
		key:      conf.ApiKey,
		userRepo: userRepo,
	}
}

func (ac *AuthUseCase) Login(ctx context.Context, req *v1.LoginReq) (*v1.LoginReply, error) {
	// get user
	user, err := ac.userRepo.FindByUsername(ctx, req.Username)
	if err != nil {
		return nil, v1.ErrorLoginFailed("user not found: %s", err.Error())
	}
	// check password
	err = ac.userRepo.VerifyPassword(ctx, user, req.Password)
	if err != nil {
		return nil, v1.ErrorLoginFailed("password not match")
	}
	// generate token
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.Id,
	})
	signedString, err := claims.SignedString([]byte(ac.key))
	if err != nil {
		return nil, v1.ErrorLoginFailed("generate token failed: %s", err.Error())
	}

	return &v1.LoginReply{Token: signedString}, nil
}

func (ac *AuthUseCase) Register(ctx context.Context, req *v1.RegisterReq) (*v1.RegisterReply, error) {
	_, err := ac.userRepo.FindByUsername(ctx, req.Username)
	if !errors.Is(err, ErrUserNotFound) {
		return nil, v1.ErrorRegisterFailed("username already exists")
	}
	user, err := NewUser(req.Username, req.Password)
	if err != nil {
		return nil, v1.ErrorRegisterFailed("create user failed: %s", err.Error())
	}
	err = ac.userRepo.Save(ctx, &user)
	if err != nil {
		return nil, v1.ErrorRegisterFailed("save user failed: %s", err.Error())
	}
	return &v1.RegisterReply{Id: user.Id}, nil
}
