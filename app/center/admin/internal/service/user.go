package service

import (
	"context"
	v1 "github.com/csyangpeng/go-kratos-admin/api/center/admin/v1"
)

func (c *CenterAdmin) GetUser(ctx context.Context, req *v1.GetUserReq) (*v1.GetUserReply, error) {
	return c.uc.Get(ctx, req)
}

func (c *CenterAdmin) ListUser(ctx context.Context, req *v1.ListUserReq) (*v1.ListUserReply, error) {
	return c.uc.List(ctx, req)
}
