package service

import (
	"context"
	v1 "github.com/csyangpeng/go-kratos-admin/api/center/interface/v1"
)

func (c *CenterInterface) Register(ctx context.Context, req *v1.RegisterReq) (*v1.RegisterReply, error) {
	c.log.Infof("%v", req)
	return c.ac.Register(ctx, req)
}

func (c *CenterInterface) Login(ctx context.Context, req *v1.LoginReq) (*v1.LoginReply, error) {
	return c.ac.Login(ctx, req)
}

func (c *CenterInterface) Logout(ctx context.Context, req *v1.LogoutReq) (*v1.LogoutReply, error) {
	//TODO implement me
	panic("implement me")
}
