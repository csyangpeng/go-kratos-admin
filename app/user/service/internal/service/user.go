package service

import (
	"context"
	v1 "github.com/csyangpeng/go-kratos-admin/api/user/service/v1"
	"github.com/csyangpeng/go-kratos-admin/app/user/service/internal/biz"
)

func (s *UserService) CreateUser(ctx context.Context, req *v1.CreateUserReq) (*v1.CreateUserReply, error) {
	user, err := s.uc.Create(ctx, &biz.User{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}

	return &v1.CreateUserReply{
		Id:       user.Id,
		Username: user.Username,
	}, nil
}

func (s *UserService) Save(ctx context.Context, req *v1.SaveUserReq) (*v1.SaveUserReply, error) {
	return s.uc.Save(ctx, req)
}

func (s *UserService) GetUser(ctx context.Context, req *v1.GetUserReq) (*v1.GetUserReply, error) {
	user, err := s.uc.Get(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &v1.GetUserReply{
		Id:       user.Id,
		Username: user.Username,
	}, nil
}

func (s *UserService) GetUserByUsername(ctx context.Context, req *v1.GetUserByUsernameReq) (*v1.GetUserByUsernameReply, error) {
	return s.uc.GetUserByUsername(ctx, req)
}

func (s *UserService) VerifyPassword(ctx context.Context, req *v1.VerifyPasswordReq) (*v1.VerifyPasswordReply, error) {
	rv, err := s.uc.VerifyPassword(ctx, &biz.User{Username: req.Username, Password: req.Password})
	if err != nil {
		return nil, err
	}

	return &v1.VerifyPasswordReply{
		Ok: rv,
	}, nil
}

func (s *UserService) ListUser(ctx context.Context, req *v1.ListUserReq) (*v1.ListUserReply, error) {
	list, total, err := s.uc.List(ctx, int(req.PageIndex), int(req.PageSize))
	if err != nil {
		return nil, err
	}
	res := &v1.ListUserReply{}
	res.Total = int64(total)
	for _, user := range list {
		res.Results = append(res.Results, &v1.ListUserReply_User{
			Id:       user.Id,
			Username: user.Username,
		})
	}

	return res, nil
}
