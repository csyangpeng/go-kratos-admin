package data

import (
	"context"
	"fmt"
	userv1 "github.com/csyangpeng/go-kratos-admin/api/user/service/v1"
	"github.com/csyangpeng/go-kratos-admin/app/center/interface/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
	"golang.org/x/sync/singleflight"
)

type userRepo struct {
	data *Data
	log  *log.Helper
	sg   *singleflight.Group
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "repo/user")),
		sg:   &singleflight.Group{},
	}
}

func (r *userRepo) FindByUsername(ctx context.Context, username string) (*biz.User, error) {
	do, err, _ := r.sg.Do(fmt.Sprintf("find_user_by_username_%s", username), func() (interface{}, error) {
		user, err := r.data.uc.GetUserByUsername(ctx, &userv1.GetUserByUsernameReq{Username: username})
		if err != nil {
			return nil, biz.ErrUserNotFound
		}
		return &biz.User{
			Id:       user.Id,
			Username: user.Username,
		}, nil
	})
	if err != nil {
		return nil, err
	}
	return do.(*biz.User), nil
}

func (r *userRepo) VerifyPassword(ctx context.Context, u *biz.User, password string) error {
	rv, err := r.data.uc.VerifyPassword(ctx, &userv1.VerifyPasswordReq{
		Username: u.Username,
		Password: password,
	})
	if err != nil {
		return err
	}

	if !rv.Ok {
		return biz.ErrPasswordFailed
	}

	return nil
}
