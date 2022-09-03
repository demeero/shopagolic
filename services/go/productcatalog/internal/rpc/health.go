package rpc

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	catalogpb "github.com/demeero/shopagolic/services/proto/gen/go/shopagolic/productcatalog/v1beta1"
)

type Health struct {
	catalogpb.UnimplementedHealthServiceServer
	mclient *mongo.Client
}

func NewHealth(mclient *mongo.Client) *Health {
	return &Health{mclient: mclient}
}

func (c *Health) Health(ctx context.Context, _ *catalogpb.HealthRequest) (*catalogpb.HealthResponse, error) {
	if err := c.mclient.Ping(ctx, readpref.Primary()); err != nil {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("failed ping MongoDB: %s", err))
	}
	return &catalogpb.HealthResponse{}, nil
}
