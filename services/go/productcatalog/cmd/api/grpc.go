package main

import (
	"context"
	"fmt"
	"net"

	grpcmiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	grpcrecovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/demeero/shopagolic/productcatalog/internal/rpc"
	"github.com/demeero/shopagolic/services/go/bricks/zaplogger"
	pb "github.com/demeero/shopagolic/services/proto/gen/go/shopagolic/productcatalog/v1beta1"
)

const healthGRPCMethodName = "/shopagolic.productcatalog.v1beta1.HealthService/Health"

func grpcServ(cfg grpcCfg, prodComponents rpc.ProductComponents, mclient *mongo.Client, zlog *zap.Logger) func() {
	interceptors := []grpc.UnaryServerInterceptor{
		grpcrecovery.UnaryServerInterceptor(),
		grpc_zap.UnaryServerInterceptor(zlog, grpc_zap.WithDecider(func(methodName string, err error) bool {
			if methodName == healthGRPCMethodName {
				return false
			}
			return true
		})),
		grpc_zap.PayloadUnaryServerInterceptor(zlog, func(_ context.Context, methodName string, _ interface{}) bool {
			if !cfg.LogPayload {
				return false
			}
			if methodName == healthGRPCMethodName {
				return false
			}
			return true
		}),
		grpcZapLogCtxInterceptor(),
	}
	grpcServ := grpc.NewServer(grpcmiddleware.WithUnaryServerChain(interceptors...))
	reflection.Register(grpcServ)
	pb.RegisterProductCatalogServiceServer(grpcServ, rpc.NewProduct(prodComponents))
	pb.RegisterHealthServiceServer(grpcServ, rpc.NewHealth(mclient))

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
