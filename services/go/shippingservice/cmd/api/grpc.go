package main

import (
	"context"
	"fmt"
	"net"

	grpcmiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	grpcrecovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/demeero/shopagolic/services/go/bricks/zaplogger"
	pb "github.com/demeero/shopagolic/services/proto/gen/go/shopagolic/shipping/v1beta1"
	"github.com/demeero/shopagolic/shippingservice/internal/rpc"
	"github.com/demeero/shopagolic/shippingservice/shipping"
)

const healthGRPCMethodName = "/shopagolic.shipping.v1beta1.HealthService/Health"

func grpcServ(cfg grpcCfg, shipper *shipping.Shipper, zlog *zap.Logger) func() {
	interceptors := []grpc.UnaryServerInterceptor{
		grpcrecovery.UnaryServerInterceptor(),
		grpc_zap.UnaryServerInterceptor(zlog, grpc_zap.WithDecider(func(methodName string, err error) bool {
			if methodName == healthGRPCMethodName {
				return false
			}
			return methodName != healthGRPCMethodName
		})),
		grpc_zap.PayloadUnaryServerInterceptor(zlog, func(_ context.Context, methodName string, _ interface{}) bool {
			if !cfg.LogPayload {
				return false
			}
			return methodName != healthGRPCMethodName
		}),
		grpcZapLogCtxInterceptor(),
	}
	grpcServ := grpc.NewServer(grpcmiddleware.WithUnaryServerChain(interceptors...))
	reflection.Register(grpcServ)
	pb.RegisterShippingServiceServer(grpcServ, rpc.NewShipping(shipper))
	pb.RegisterHealthServiceServer(grpcServ, rpc.NewHealth())

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		zap.L().Fatal("failed listen GRPC port", zap.Error(err))
	}
	go func() {
		if err := grpcServ.Serve(lis); err != nil {
			zap.L().Fatal("failed serve GRPC", zap.Error(err))
		}
	}()
	return func() {
		grpcServ.GracefulStop()
	}
}

func grpcZapLogCtxInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		return handler(zaplogger.ToCtx(ctx, ctxzap.Extract(ctx)), req)
	}
}
