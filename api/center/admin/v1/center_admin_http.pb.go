// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.3.1

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationCenterAdminGetUser = "/center.admin.v1.CenterAdmin/GetUser"
const OperationCenterAdminListUser = "/center.admin.v1.CenterAdmin/ListUser"
const OperationCenterAdminLogin = "/center.admin.v1.CenterAdmin/Login"
const OperationCenterAdminLogout = "/center.admin.v1.CenterAdmin/Logout"

type CenterAdminHTTPServer interface {
	GetUser(context.Context, *GetUserReq) (*GetUserReply, error)
	ListUser(context.Context, *ListUserReq) (*ListUserReply, error)
	Login(context.Context, *LoginReq) (*LoginReply, error)
	Logout(context.Context, *LogoutReq) (*LogoutReply, error)
}

func RegisterCenterAdminHTTPServer(s *http.Server, srv CenterAdminHTTPServer) {
	r := s.Route("/")
	r.POST("/admin/v1/login", _CenterAdmin_Login0_HTTP_Handler(srv))
	r.POST("/admin/v1/logout", _CenterAdmin_Logout0_HTTP_Handler(srv))
	r.GET("/admin/v1/users/{id}", _CenterAdmin_GetUser0_HTTP_Handler(srv))
	r.GET("/v1/users", _CenterAdmin_ListUser0_HTTP_Handler(srv))
}

func _CenterAdmin_Login0_HTTP_Handler(srv CenterAdminHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LoginReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCenterAdminLogin)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Login(ctx, req.(*LoginReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*LoginReply)
		return ctx.Result(200, reply)
	}
}

func _CenterAdmin_Logout0_HTTP_Handler(srv CenterAdminHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LogoutReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCenterAdminLogout)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Logout(ctx, req.(*LogoutReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*LogoutReply)
		return ctx.Result(200, reply)
	}
}

func _CenterAdmin_GetUser0_HTTP_Handler(srv CenterAdminHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetUserReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCenterAdminGetUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetUser(ctx, req.(*GetUserReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetUserReply)
		return ctx.Result(200, reply)
	}
}

func _CenterAdmin_ListUser0_HTTP_Handler(srv CenterAdminHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListUserReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCenterAdminListUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListUser(ctx, req.(*ListUserReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListUserReply)
		return ctx.Result(200, reply)
	}
}

type CenterAdminHTTPClient interface {
	GetUser(ctx context.Context, req *GetUserReq, opts ...http.CallOption) (rsp *GetUserReply, err error)
	ListUser(ctx context.Context, req *ListUserReq, opts ...http.CallOption) (rsp *ListUserReply, err error)
	Login(ctx context.Context, req *LoginReq, opts ...http.CallOption) (rsp *LoginReply, err error)
	Logout(ctx context.Context, req *LogoutReq, opts ...http.CallOption) (rsp *LogoutReply, err error)
}

type CenterAdminHTTPClientImpl struct {
	cc *http.Client
}

func NewCenterAdminHTTPClient(client *http.Client) CenterAdminHTTPClient {
	return &CenterAdminHTTPClientImpl{client}
}

func (c *CenterAdminHTTPClientImpl) GetUser(ctx context.Context, in *GetUserReq, opts ...http.CallOption) (*GetUserReply, error) {
	var out GetUserReply
	pattern := "/admin/v1/users/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationCenterAdminGetUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *CenterAdminHTTPClientImpl) ListUser(ctx context.Context, in *ListUserReq, opts ...http.CallOption) (*ListUserReply, error) {
	var out ListUserReply
	pattern := "/v1/users"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationCenterAdminListUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *CenterAdminHTTPClientImpl) Login(ctx context.Context, in *LoginReq, opts ...http.CallOption) (*LoginReply, error) {
	var out LoginReply
	pattern := "/admin/v1/login"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationCenterAdminLogin))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *CenterAdminHTTPClientImpl) Logout(ctx context.Context, in *LogoutReq, opts ...http.CallOption) (*LogoutReply, error) {
	var out LogoutReply
	pattern := "/admin/v1/logout"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationCenterAdminLogout))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
