package server

import (
	v1 "github.com/csyangpeng/go-kratos-admin/api/user/service/v1"
	"github.com/csyangpeng/go-kratos-admin/app/user/service/internal/conf"
	"github.com/csyangpeng/go-kratos-admin/app/user/service/internal/service"
	prom "github.com/go-kratos/kratos/contrib/metrics/prometheus/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/metrics"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	jwtv4 "github.com/golang-jwt/jwt/v4"
	"github.com/prometheus/client_golang/prometheus"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Server, ac *conf.Auth, u *service.UserService, logger log.Logger) *grpc.Server {

	// https://github.com/go-kratos/examples/tree/main/metrics
	_metricSeconds := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "server",
		Subsystem: "requests",
		Name:      "duration_ms",
		Help:      "server requests duration(ms).",
		Buckets:   []float64{5, 10, 25, 50, 100, 250, 500, 1000},
	}, []string{"kind", "operation"})

	_metricRequests := prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "client",
		Subsystem: "requests",
		Name:      "code_total",
		Help:      "The total number of processed requests",
	}, []string{"kind", "operation", "code", "reason"})

	prometheus.MustRegister(_metricSeconds, _metricRequests)

	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
			tracing.Server(), // 链路追踪
			metrics.Server(
				metrics.WithSeconds(prom.NewHistogram(_metricSeconds)),
				metrics.WithRequests(prom.NewCounter(_metricRequests)),
			),
			jwt.Server(func(token *jwtv4.Token) (interface{}, error) {
				return []byte(ac.Key), nil
			}, jwt.WithSigningMethod(jwtv4.SigningMethodHS256)),
		),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	v1.RegisterUserServer(srv, u)
	return srv
}
