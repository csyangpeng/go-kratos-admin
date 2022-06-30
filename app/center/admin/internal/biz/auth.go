package biz

import (
	"context"
	"errors"
	v1 "github.com/csyangpeng/go-kratos-admin/api/center/admin/v1"
	"github.com/csyangpeng/go-kratos-admin/app/center/admin/internal/conf"
	"github.com/golang-jwt/jwt/v4"
)

var (
	ErrLoginFailed    = errors.New("login failed")
	ErrPasswordFailed = errors.New("password failed")
	ErrUserDeactivate = errors.New("user is deactivated")
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

	// check isActive
	if !user.IsActive {
		return nil, v1.ErrorLoginFailed(ErrUserDeactivate.Error())
	}

	// check password
	err = ac.userRepo.VerifyPassword(ctx, user, req.Password)
	if err != nil {
		return nil, err
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
