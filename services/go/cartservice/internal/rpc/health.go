package rpc

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	cartpb "github.com/demeero/shopagolic/services/proto/gen/go/shopagolic/cart/v1beta1"
)

type Health struct {
	cartpb.UnimplementedHealthServiceServer
	rClient *redis.Client
}

func NewHealth(rClient *redis.Client) *Health {
	return &Health{rClient: rClient}
}

func (c *Health) Health(ctx context.Context, _ *cartpb.HealthRequest) (*cartpb.HealthResponse, error) {
	if err := c.rClient.Ping(ctx).Err(); err != nil {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("failed ping Redis: %s", err))
	}
	return &cartpb.HealthResponse{}, nil
}
