package grpc

import (
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	"github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func NewGRPCServer(tracer opentracing.Tracer, logger *logrus.Entry, options ...grpc.ServerOption) *grpc.Server {

	var (
		opts = append([]grpc.ServerOption{
			grpc.UnaryInterceptor(defaultServerInterceptors(tracer, logger)),
		}, options...)
	)

	return grpc.NewServer(opts...)
}

func defaultServerInterceptors(tracer opentracing.Tracer, logger *logrus.Entry) grpc.UnaryServerInterceptor {
	return grpc_middleware.ChainUnaryServer(
		grpc_opentracing.UnaryServerInterceptor(
			grpc_opentracing.WithTracer(tracer),
		),
		grpc_logrus.UnaryServerInterceptor(logger),
		grpc_prometheus.UnaryServerInterceptor,
	)
}
