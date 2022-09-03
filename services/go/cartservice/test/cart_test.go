package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"

	"github.com/demeero/shopagolic/cartservice/cart"
	cartpb "github.com/demeero/shopagolic/services/proto/gen/go/shopagolic/cart/v1beta1"
)

func TestIntegrationCart(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	suite.Run(t, &cartTestSuite{})
}

type cartTestSuite struct {
	suite.Suite
	ctx        context.Context
	keyPrefix  string
	rds        *redis.Client
	grpcClient cartpb.CartServiceClient
}

func (ts *cartTestSuite) SetupSuite() {
	ts.ctx = context.Background()
	c := getConfig(ts.T())
	ts.keyPrefix = c.CartKeyPrefix
	ts.T().Logf("cart test config: %+v", c)
	ts.rds = redis.NewClient(&redis.Options{Addr: c.RedisAddr})
	ts.grpcClient = cartGRPCClient(ts.T())
}

func (ts *cartTestSuite) TearDownSuite() {
	ts.Require().NoError(ts.rds.Close())
}

func (ts *cartTestSuite) TearDownTest() {
	ts.Require().NoError(ts.rds.FlushAll(context.Background()).Err())
}

func (ts *cartTestSuite) TestAddItem() {
	userID := gofakeit.UUID()
	resp, err := ts.grpcClient.AddItem(ts.ctx, &cartpb.AddItemRequest{
		UserId: userID,
		Item: &cartpb.CartItem{
			ProductId: gofakeit.UUID(),
			Quantity:  3,
		},
	})
	ts.NoError(err)
	ts.NotNil(resp)

	result, err := ts.rds.HGetAll(ts.ctx, key(ts.keyPrefix, userID)).Result()
	ts.NoError(err)
	ts.Equal(1, len(result))
}

func (ts *cartTestSuite) TestAddItem_InvalidInput() {
	tests := []struct {
		name string
		req  *cartpb.AddItemRequest
	}{
		{
			name: "empty user_id",
			req: &cartpb.AddItemRequest{
				Item: &cartpb.CartItem{
					ProductId: gofakeit.UUID(),
					Quantity:  3,
				},
			},
		},
		{
			name: "zero quantity",
			req: &cartpb.AddItemRequest{
				UserId: gofakeit.UUID(),
				Item:   &cartpb.CartItem{ProductId: gofakeit.UUID()},
			},
		},
		{
			name: "empty product_id",
			req: &cartpb.AddItemRequest{
				UserId: gofakeit.UUID(),
				Item:   &cartpb.CartItem{Quantity: 3},
			},
		},
	}
	for _, tt := range tests {
		ts.Run(tt.name, func() {
			resp, err := ts.grpcClient.AddItem(ts.ctx, tt.req)
			ts.Nil(resp)
			ts.Error(err)
			s, _ := status.FromError(err)
			ts.Equal("InvalidArgument", s.Code().String())
		})
	}
}

func (ts *cartTestSuite) TestGetCart() {
	var (
		cart1 = cart.Cart{
			UserID: gofakeit.UUID(),
			Items: []cart.Item{
				{ProductID: gofakeit.UUID(), Quantity: 1},
				{ProductID: gofakeit.UUID(), Quantity: 2},
			},
		}
		cart2 = cart.Cart{
			UserID: gofakeit.UUID(),
			Items: []cart.Item{
				{ProductID: gofakeit.UUID(), Quantity: 3},
				{ProductID: gofakeit.UUID(), Quantity: 4},
				{ProductID: gofakeit.UUID(), Quantity: 5},
			},
		}
		carts = []cart.Cart{cart1, cart2}
	)

	for _, c := range carts {
		for _, item := range c.Items {
			ts.Require().NoError(ts.rds.HSet(ts.ctx, key(ts.keyPrefix, c.UserID), item.ProductID, item.Quantity).Err())
		}
	}

	for _, c := range carts {
		expected := &cartpb.GetCartResponse{UserId: c.UserID}
		for _, item := range c.Items {
			expected.Items = append(expected.Items, &cartpb.CartItem{
				ProductId: item.ProductID,
				Quantity:  int32(item.Quantity),
			})
		}
		actual, err := ts.grpcClient.GetCart(ts.ctx, &cartpb.GetCartRequest{UserId: c.UserID})
		ts.NoError(err)
		ts.Equal(expected.UserId, actual.UserId)
		ts.ElementsMatch(expected.Items, actual.Items)
	}
}

func (ts *cartTestSuite) TestGetCart_InvalidInput() {
	actual, err := ts.grpcClient.GetCart(ts.ctx, &cartpb.GetCartRequest{})
	ts.Nil(actual)
	ts.Error(err)
	s, _ := status.FromError(err)
	ts.Equal("InvalidArgument", s.Code().String())
}

func (ts *cartTestSuite) TestEmptyCart() {
	var (
		cart1 = cart.Cart{
			UserID: gofakeit.UUID(),
			Items: []cart.Item{
				{ProductID: gofakeit.UUID(), Quantity: 1},
				{ProductID: gofakeit.UUID(), Quantity: 2},
			},
		}
		cart2 = cart.Cart{
			UserID: gofakeit.UUID(),
			Items: []cart.Item{
				{ProductID: gofakeit.UUID(), Quantity: 3},
				{ProductID: gofakeit.UUID(), Quantity: 4},
				{ProductID: gofakeit.UUID(), Quantity: 5},
			},
		}
		carts = []cart.Cart{cart1, cart2}
	)

	for _, c := range carts {
		for _, item := range c.Items {
			ts.Require().NoError(ts.rds.HSet(ts.ctx, key(ts.keyPrefix, c.UserID), item.ProductID, item.Quantity).Err())
		}
	}

	actual, err := ts.grpcClient.EmptyCart(ts.ctx, &cartpb.EmptyCartRequest{UserId: cart1.UserID})
	ts.NoError(err)
	ts.NotNil(actual)

	result, err := ts.rds.HGetAll(ts.ctx, key(ts.keyPrefix, cart1.UserID)).Result()
	ts.NoError(err)
	ts.Empty(result)

	result, err = ts.rds.HGetAll(ts.ctx, key(ts.keyPrefix, cart2.UserID)).Result()
	ts.NoError(err)
	ts.NotEmpty(result)
}

func (ts *cartTestSuite) TestEmptyCart_Error() {
	cart1 := cart.Cart{
		UserID: gofakeit.UUID(),
		Items:  []cart.Item{{ProductID: gofakeit.UUID(), Quantity: 1}},
	}
	err := ts.rds.HSet(ts.ctx, key(ts.keyPrefix, cart1.UserID), cart1.Items[0].ProductID, cart1.Items[0].Quantity).Err()
	ts.Require().NoError(err)

	tests := []struct {
		name            string
		req             *cartpb.EmptyCartRequest
		expectedErrCode codes.Code
	}{
		{
			name:            "empty user_id",
			req:             &cartpb.EmptyCartRequest{},
			expectedErrCode: codes.InvalidArgument,
		},
		{
			name:            "cart not found",
			req:             &cartpb.EmptyCartRequest{UserId: gofakeit.UUID()},
			expectedErrCode: codes.NotFound,
		},
	}
	for _, tt := range tests {
		ts.Run(tt.name, func() {
			actual, err := ts.grpcClient.EmptyCart(ts.ctx, tt.req)
			ts.Nil(actual)
			ts.Error(err)
			s, _ := status.FromError(err)
			ts.Equal(tt.expectedErrCode.String(), s.Code().String())
		})
	}
}

func cartGRPCClient(t *testing.T) cartpb.CartServiceClient {
	t.Helper()
	conn, err := grpc.Dial(getConfig(t).GRPCAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	require.NoError(t, err)
	return cartpb.NewCartServiceClient(conn)
}

func key(prefix, val string) string {
	return fmt.Sprintf("%s:%s", prefix, val)
}
