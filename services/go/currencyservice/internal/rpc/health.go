package rpc

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	currpb "github.com/demeero/shopagolic/services/proto/gen/go/shopagolic/currency/v1beta1"
)

type Health struct {
	currpb.UnimplementedHealthServiceServer
	rClient *redis.Client
}

func NewHealth(rClient *redis.Client) *Health {
	return &Health{rClient: rClient}
}

func (c *Health) Health(ctx context.Context, _ *currpb.HealthRequest) (*currpb.HealthResponse, error) {
	if err := c.rClient.Ping(ctx).Err(); err != nil {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("failed ping Redis: %s", err))
	}
	return &currpb.HealthResponse{}, nil
}
