package data

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/csyangpeng/go-kratos-admin/app/user/service/internal/biz"
	"github.com/csyangpeng/go-kratos-admin/app/user/service/internal/data/ent"
	"github.com/csyangpeng/go-kratos-admin/app/user/service/internal/data/ent/user"
	"github.com/csyangpeng/go-kratos-admin/app/user/service/internal/pkg/util"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

var _ biz.UserRepo = (*userRepo)(nil)

var cacheKey = func(suffix string) string {
	return "user_cache_key_" + suffix
}

type userRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/server-service")),
	}
}

func (r *userRepo) GetUser(ctx context.Context, id int64) (*biz.User, error) {
	cacheKey := cacheKey(fmt.Sprintf("%d", id))
	target, err := r.getUserFromCache(ctx, cacheKey)
	if err != nil {
		target, err = r.data.db.User.Get(ctx, id)
		if err != nil {
			return nil, biz.ErrUserNotFound
		}
		r.setUserCache(ctx, target, cacheKey)
	}
	return &biz.User{
		Id:       target.ID,
		Username: target.Username,
	}, nil
}

func (r *userRepo) FindByUsername(ctx context.Context, username string) (*biz.User, error) {
	var target *ent.User
	cacheKey := cacheKey(username)
	target, err := r.getUserFromCache(ctx, cacheKey)
	if err != nil {
		target, err = r.data.db.User.
			Query().
			Where(user.UsernameEQ(username)).
			Only(ctx)
		if err != nil {
			return nil, biz.ErrUserNotFound
		}
		r.setUserCache(ctx, target, cacheKey)
	}

	return &biz.User{
		Id:       target.ID,
		Username: target.Username,
	}, nil
}

func (r *userRepo) CreateUser(ctx context.Context, u *biz.User) (*biz.User, error) {
	ph, err := util.HashPassword(u.Password)
	if err != nil {
		return nil, err
	}
	res, err := r.data.db.User.
		Create().
		SetUsername(u.Username).
		SetPasswordHash(ph).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &biz.User{
		Id:       res.ID,
		Username: res.Username,
	}, nil
}

func (r *userRepo) getUserFromCache(ctx context.Context, key string) (*ent.User, error) {
	result, err := r.data.redis.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	var cacheUser = &ent.User{}
	err = json.Unmarshal([]byte(result), cacheUser)
	if err != nil {
		return nil, err
	}
	return cacheUser, nil
}

func (r *userRepo) setUserCache(ctx context.Context, user *ent.User, key string) {
	marshal, err := json.Marshal(user)
	if err != nil {
		r.log.Errorf("fail to set user cache:json.Marshal(%v) error(%v)", user, err)
	}
	err = r.data.redis.Set(ctx, key, marshal, time.Minute*30).Err()
	if err != nil {
		r.log.Errorf("fail to set user cache:redis.Set(%v) error(%v)", user, err)
	}
}

func (r *userRepo) VerifyPassword(ctx context.Context, u *biz.User) (bool, error) {
	po, err := r.data.db.User.Query().Where(user.UsernameEQ(u.Username)).Only(ctx)
	if err != nil {
		return false, err
	}

	return util.CheckPassword(u.Password, po.PasswordHash), nil
}

func (r *userRepo) ListUser(ctx context.Context, pageIndex, pageSize int) ([]*biz.User, int, error) {
	users, err := r.paginate(pageIndex, pageSize).All(ctx)
	if err != nil {
		return nil, 0, err
	}
	total, err := r.data.db.User.Query().Count(ctx)

	if err != nil {
		return nil, 0, err
	}

	rv := make([]*biz.User, 0)
	for _, u := range users {
		rv = append(rv, &biz.User{
			Id:       u.ID,
			Username: u.Username,
			Password: u.PasswordHash,
		})
	}
	return rv, total, nil
}

func (r *userRepo) paginate(page, pageSize int) *ent.UserQuery {
	if page == 0 {
		page = 1
	}
	switch {
	case pageSize > 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	offset := (page - 1) * pageSize
	return r.data.db.User.Query().Offset(offset).Limit(pageSize)

}
