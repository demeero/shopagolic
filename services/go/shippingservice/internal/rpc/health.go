package rpc

import (
	"context"

	shippingpb "github.com/demeero/shopagolic/services/proto/gen/go/shopagolic/shipping/v1beta1"
)

type Health struct {
	shippingpb.UnimplementedHealthServiceServer
}

func NewHealth() *Health {
	return &Health{}
}

func (c *Health) Health(context.Context, *shippingpb.HealthRequest) (*shippingpb.HealthResponse, error) {
	return &shippingpb.HealthResponse{}, nil
}
