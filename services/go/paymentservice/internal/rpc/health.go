package rpc

import (
	"context"

	pb "github.com/demeero/shopagolic/services/proto/gen/go/shopagolic/payment/v1beta1"
)

type Health struct {
	pb.UnimplementedHealthServiceServer
}

func NewHealth() *Health {
	return &Health{}
}

func (h *Health) Health(context.Context, *pb.HealthRequest) (*pb.HealthResponse, error) {
	return &pb.HealthResponse{}, nil
}
