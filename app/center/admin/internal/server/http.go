package server

import (
	"context"
	v1 "github.com/csyangpeng/go-kratos-admin/api/center/admin/v1"
	"github.com/csyangpeng/go-kratos-admin/app/center/admin/internal/conf"
	"github.com/csyangpeng/go-kratos-admin/app/center/admin/internal/service"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-kratos/swagger-api/openapiv2"
	jwtv4 "github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/handlers"
)

func NewWhiteListMatcher() selector.MatchFunc {

	whiteList := make(map[string]struct{})
	whiteList["/center.interface.v1.CenterInterface/Login"] = struct{}{}
	whiteList["/center.interface.v1.CenterInterface/Register"] = struct{}{}
	return func(ctx context.Context, operation string) bool {
		if _, ok := whiteList[operation]; ok {
			return false
		}
		return true
	}
}

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, ac *conf.Auth, s *service.CenterAdmin, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			logging.Server(logger),
			tracing.Server(),
			validate.Validator(),
			selector.Server(
				jwt.Server(func(token *jwtv4.Token) (interface{}, error) {
					return []byte(ac.ApiKey), nil
				}, jwt.WithSigningMethod(jwtv4.SigningMethodHS256), jwt.WithClaims(func() jwtv4.Claims {
					return &jwtv4.MapClaims{}
				})),
			).
				Match(NewWhiteListMatcher()).
				Build(),
		),
		http.Filter(handlers.CORS(
			handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
			handlers.AllowedOrigins([]string{"*"}),
		)),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	h := openapiv2.NewHandler()
	srv.HandlePrefix("/q/", h)
	v1.RegisterCenterAdminHTTPServer(srv, s)
	return srv
}
