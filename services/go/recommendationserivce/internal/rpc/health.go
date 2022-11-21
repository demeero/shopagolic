package rpc

import (
	"context"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	catalogpb "github.com/demeero/shopagolic/services/proto/gen/go/shopagolic/productcatalog/v1beta1"
	recommendationpb "github.com/demeero/shopagolic/services/proto/gen/go/shopagolic/recommendation/v1beta1"
)

type Health struct {
	recommendationpb.UnimplementedHealthServiceServer
	catalogHealthClient catalogpb.HealthServiceClient
}

func NewHealth(components Components) *Health {
	return &Health{catalogHealthClient: components.CatalogHealthClient}
}

func (c *Health) Health(ctx context.Context, _ *recommendationpb.HealthRequest) (*recommendationpb.HealthResponse, error) {
	if _, err := c.catalogHealthClient.Health(ctx, &catalogpb.HealthRequest{}); err != nil {
		return nil, status.Error(codes.Unavailable, fmt.Sprintf("failed health catalog service: %s", err))
	}
	return &recommendationpb.HealthResponse{}, nil
}
