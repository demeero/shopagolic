package test

import (
	"context"
	"math"
	"testing"

	recommendationpb "github.com/demeero/shopagolic/services/proto/gen/go/shopagolic/recommendation/v1beta1"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

func TestIntegrationRecommendation(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	suite.Run(t, &recommendationTestSuite{})
}

type recommendationTestSuite struct {
	suite.Suite
	ctx        context.Context
	grpcClient recommendationpb.RecommendationServiceClient
}

func (ts *recommendationTestSuite) SetupSuite() {
	ts.ctx = context.Background()
	c := getConfig(ts.T())
	ts.T().Logf("recommendation test config: %+v", c)
	ts.grpcClient = recommendationGRPCClient(ts.T())
}

func (ts *recommendationTestSuite) TestLoad() {
	tests := []struct {
		name  string
		limit int32
	}{
		{name: "default-limit"},
		{name: "small-limit", limit: 2},
		{name: "big-limit", limit: 100},
	}
	for _, tt := range tests {
		ts.Run(tt.name, func() {
			actual, err := ts.grpcClient.GetRecommendation(ts.ctx, &recommendationpb.GetRecommendationRequest{ProductId: "1", Limit: tt.limit})
			ts.NoError(err)
			ts.NotEmpty(actual.GetProductIds())
			if tt.limit > 0 {
				ts.LessOrEqual(len(actual.GetProductIds()), int(tt.limit))
			}
		})
	}
}

func (ts *recommendationTestSuite) TestLoad_Err() {
	tests := []struct {
		name            string
		productID       string
		limit           int32
		expectedErrCode codes.Code
	}{
		{
			name:            "negative-limit",
			productID:       "1",
			limit:           -1,
			expectedErrCode: codes.InvalidArgument,
		},
		{
			name:            "over-limit",
			productID:       "1",
			limit:           math.MaxInt32,
			expectedErrCode: codes.InvalidArgument,
		},
		{
			name:            "undefined-product-id",
			productID:       "100500",
			limit:           100,
			expectedErrCode: codes.NotFound,
		},
	}
	for _, tt := range tests {
		ts.Run(tt.name, func() {
			req := &recommendationpb.GetRecommendationRequest{ProductId: tt.productID, Limit: tt.limit}
			actual, err := ts.grpcClient.GetRecommendation(ts.ctx, req)
			ts.Nil(actual)
			ts.Error(err)
			s, _ := status.FromError(err)
			ts.Equal(tt.expectedErrCode.String(), s.Code().String())
		})
	}
}

func recommendationGRPCClient(t *testing.T) recommendationpb.RecommendationServiceClient {
	t.Helper()
	conn, err := grpc.Dial(getConfig(t).GRPCAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	require.NoError(t, err)
	return recommendationpb.NewRecommendationServiceClient(conn)
}
