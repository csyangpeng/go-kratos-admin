package data

import (
	"context"
	"fmt"

	v1 "github.com/csyangpeng/go-kratos-admin/api/center/admin/v1"
	userv1 "github.com/csyangpeng/go-kratos-admin/api/user/service/v1"
	"github.com/csyangpeng/go-kratos-admin/app/center/admin/internal/biz"
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

func (r *userRepo) GetUser(ctx context.Context, id int64) (*biz.User, error) {
	do, err, _ := r.sg.Do(fmt.Sprintf("get_user_id_%d", id), func() (interface{}, error) {
		u, err := r.data.uc.GetUser(ctx, &userv1.GetUserReq{Id: id})

		if err != nil {
			return nil, biz.ErrUserNotFound
		}
		return &biz.User{
			Id:       u.Id,
			Username: u.Username,
		}, nil
	})
	if err != nil {
		return nil, err
	}

	return do.(*biz.User), nil
}

func (r *userRepo) FindByUsername(ctx context.Context, username string) (*biz.User, error) {
	r.log.Info(username)
	do, err, _ := r.sg.Do(fmt.Sprintf("find_user_by_username_%s", username), func() (interface{}, error) {
		user, err := r.data.uc.GetUserByUsername(ctx, &userv1.GetUserByUsernameReq{Username: username})
		if err != nil {
			return nil, err
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

func (r *userRepo) ListUser(ctx context.Context, req *v1.ListUserReq) (*v1.ListUserReply, error) {
	lst, err := r.data.uc.ListUser(ctx, &userv1.ListUserReq{
		PageIndex: req.PageIndex,
		PageSize:  req.PageSize,
	})
	if err != nil {
		return nil, err
	}
	v := &v1.ListUserReply{}
	v.Total = lst.Total
	for _, u := range lst.Results {
		v.Results = append(v.Results, &v1.ListUserReply_User{
			Id:       u.Id,
			Username: u.Username,
		})
	}

	return v, nil
}
