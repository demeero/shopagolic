package rpc

import (
	"context"
	"errors"
	"fmt"
	"math"

	"github.com/demeero/shopagolic/recommendationservice/recommendation"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/demeero/shopagolic/services/go/bricks/zaplogger"
	recommendationpb "github.com/demeero/shopagolic/services/proto/gen/go/shopagolic/recommendation/v1beta1"
)

type Recommendation struct {
	recommendationpb.UnimplementedRecommendationServiceServer
	loader *recommendation.Loader
}

func NewRecommendation(components Components) *Recommendation {
	return &Recommendation{loader: components.Loader}
}

func (c *Recommendation) GetRecommendation(ctx context.Context, req *recommendationpb.GetRecommendationRequest) (*recommendationpb.GetRecommendationResponse, error) {
	if req.GetLimit() < 0 || req.GetLimit() > math.MaxUint8 {
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("limit must be in range: 0 - %d", math.MaxUint8))
	}
	rec, err := c.loader.Load(ctx, req.GetProductId(), uint8(req.GetLimit()))
	if err := errHandler(ctx, err); err != nil {
		return nil, err
	}
	return &recommendationpb.GetRecommendationResponse{ProductIds: rec.ProductIDs}, nil
}

func errHandler(ctx context.Context, err error) error {
	if err == nil {
		return nil
	}
	if errors.Is(err, recommendation.ErrNotFound) {
		return status.Error(codes.NotFound, err.Error())
	}
	if errors.Is(err, recommendation.ErrInvalidData) {
		return status.Error(codes.InvalidArgument, err.Error())
	}
	zaplogger.FromCtx(ctx).Error("failed handle request", zap.Error(err))
	return err
}
