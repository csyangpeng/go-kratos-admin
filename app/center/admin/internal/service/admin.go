package service

import (
	"context"
	"github.com/csyangpeng/go-kratos-admin/api/center/admin/v1"
)

func (c *CenterAdmin) Login(ctx context.Context, req *v1.LoginReq) (*v1.LoginReply, error) {
	return c.ac.Login(ctx, req)
}

func (c *CenterAdmin) Logout(ctx context.Context, req *v1.LogoutReq) (*v1.LogoutReply, error) {
	//TODO implement me
	panic("implement me")
}
