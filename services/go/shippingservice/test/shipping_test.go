package test

import (
	"context"
	"strconv"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"

	shippingpb "github.com/demeero/shopagolic/services/proto/gen/go/shopagolic/shipping/v1beta1"
)

func TestIntegrationShipping(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	suite.Run(t, &shippingTestSuite{})
}

type shippingTestSuite struct {
	suite.Suite
	ctx        context.Context
	grpcClient shippingpb.ShippingServiceClient
}

func (ts *shippingTestSuite) SetupSuite() {
	ts.ctx = context.Background()
	c := getConfig(ts.T())
	ts.T().Logf("shipping test config: %+v", c)
	ts.grpcClient = currGRPCClient(ts.T())
}

func (ts *shippingTestSuite) TestGetQuote() {
	req := &shippingpb.GetQuoteRequest{
		Address: &shippingpb.Address{
			StreetAddress: gofakeit.Street(),
			City:          gofakeit.City(),
			State:         gofakeit.State(),
			Country:       gofakeit.Country(),
			ZipCode:       fakeZIP(ts.T()),
		},
		Items: []*shippingpb.Item{
			{
				ProductId: gofakeit.UUID(),
				Quantity:  2,
			},
		},
	}
	resp, err := ts.grpcClient.GetQuote(ts.ctx, req)
	ts.NoError(err)
	ts.NotEmpty(resp.GetCost().GetCurrencyCode())
}

func (ts *shippingTestSuite) TestGetQuote_InvalidData() {
	tests := []struct {
		name string
		req  *shippingpb.GetQuoteRequest
	}{
		{
			name: "street empty",
			req: &shippingpb.GetQuoteRequest{
				Address: &shippingpb.Address{
					City:    gofakeit.City(),
					State:   gofakeit.State(),
					Country: gofakeit.Country(),
					ZipCode: fakeZIP(ts.T()),
				},
				Items: []*shippingpb.Item{
					{
						ProductId: gofakeit.UUID(),
						Quantity:  2,
					},
				},
			},
		},
		{
			name: "street too short",
			req: &shippingpb.GetQuoteRequest{
				Address: &shippingpb.Address{
					StreetAddress: gofakeit.Letter(),
					City:          gofakeit.City(),
					State:         gofakeit.State(),
					Country:       gofakeit.Country(),
					ZipCode:       fakeZIP(ts.T()),
				},
				Items: []*shippingpb.Item{
					{
						ProductId: gofakeit.UUID(),
						Quantity:  2,
					},
				},
			},
		},
		{
			name: "street too long",
			req: &shippingpb.GetQuoteRequest{
				Address: &shippingpb.Address{
					StreetAddress: gofakeit.LetterN(1500),
					City:          gofakeit.City(),
					State:         gofakeit.State(),
					Country:       gofakeit.Country(),
					ZipCode:       fakeZIP(ts.T()),
				},
				Items: []*shippingpb.Item{
					{
						ProductId: gofakeit.UUID(),
						Quantity:  2,
					},
				},
			},
		},
		{
			name: "city empty",
			req: &shippingpb.GetQuoteRequest{
				Address: &shippingpb.Address{
					StreetAddress: gofakeit.Street(),
					State:         gofakeit.State(),
					Country:       gofakeit.Country(),
					ZipCode:       fakeZIP(ts.T()),
				},
				Items: []*shippingpb.Item{
					{
						ProductId: gofakeit.UUID(),
						Quantity:  2,
					},
				},
			},
		},
		{
			name: "city too short",
			req: &shippingpb.GetQuoteRequest{
				Address: &shippingpb.Address{
					StreetAddress: gofakeit.Street(),
					City:          gofakeit.Letter(),
					State:         gofakeit.State(),
					Country:       gofakeit.Country(),
					ZipCode:       fakeZIP(ts.T()),
				},
				Items: []*shippingpb.Item{
					{
						ProductId: gofakeit.UUID(),
						Quantity:  2,
					},
				},
			},
		},
		{
			name: "city too long",
			req: &shippingpb.GetQuoteRequest{
				Address: &shippingpb.Address{
					StreetAddress: gofakeit.Street(),
					City:          gofakeit.LetterN(500),
					State:         gofakeit.State(),
					Country:       gofakeit.Country(),
					ZipCode:       fakeZIP(ts.T()),
				},
				Items: []*shippingpb.Item{
					{
						ProductId: gofakeit.UUID(),
						Quantity:  2,
					},
				},
			},
		},
		{
			name: "state empty",
			req: &shippingpb.GetQuoteRequest{
				Address: &shippingpb.Address{
					StreetAddress: gofakeit.Street(),
					City:          gofakeit.City(),
					Country:       gofakeit.Country(),
					ZipCode:       fakeZIP(ts.T()),
				},
				Items: []*shippingpb.Item{
					{
						ProductId: gofakeit.UUID(),
						Quantity:  2,
					},
				},
			},
		},
		{
			name: "state too short",
			req: &shippingpb.GetQuoteRequest{
				Address: &shippingpb.Address{
					StreetAddress: gofakeit.Street(),
					City:          gofakeit.City(),
					State:         gofakeit.Letter(),
					Country:       gofakeit.Country(),
					ZipCode:       fakeZIP(ts.T()),
				},
				Items: []*shippingpb.Item{
					{
						ProductId: gofakeit.UUID(),
						Quantity:  2,
					},
				},
			},
		},
		{
			name: "state too long",
			req: &shippingpb.GetQuoteRequest{
				Address: &shippingpb.Address{
					StreetAddress: gofakeit.Street(),
					City:          gofakeit.City(),
					State:         gofakeit.LetterN(500),
					Country:       gofakeit.Country(),
					ZipCode:       fakeZIP(ts.T()),
				},
				Items: []*shippingpb.Item{
					{
						ProductId: gofakeit.UUID(),
						Quantity:  2,
					},
				},
			},
		},
		{
			name: "country empty",
			req: &shippingpb.GetQuoteRequest{
				Address: &shippingpb.Address{
					StreetAddress: gofakeit.Street(),
					City:          gofakeit.City(),
					State:         gofakeit.State(),
					ZipCode:       fakeZIP(ts.T()),
				},
				Items: []*shippingpb.Item{
					{
						ProductId: gofakeit.UUID(),
						Quantity:  2,
					},
				},
			},
		},
		{
			name: "country too short",
			req: &shippingpb.GetQuoteRequest{
				Address: &shippingpb.Address{
					StreetAddress: gofakeit.Street(),
					City:          gofakeit.City(),
					State:         gofakeit.State(),
					Country:       gofakeit.Letter(),
					ZipCode:       fakeZIP(ts.T()),
				},
				Items: []*shippingpb.Item{
					{
						ProductId: gofakeit.UUID(),
						Quantity:  2,
					},
				},
			},
		},
		{
			name: "country too long",
			req: &shippingpb.GetQuoteRequest{
				Address: &shippingpb.Address{
					StreetAddress: gofakeit.Street(),
					City:          gofakeit.City(),
					State:         gofakeit.State(),
					Country:       gofakeit.LetterN(500),
					ZipCode:       fakeZIP(ts.T()),
				},
				Items: []*shippingpb.Item{
					{
						ProductId: gofakeit.UUID(),
						Quantity:  2,
					},
				},
			},
		},
		{
			name: "zip empty",
			req: &shippingpb.GetQuoteRequest{
				Address: &shippingpb.Address{
					StreetAddress: gofakeit.Street(),
					City:          gofakeit.City(),
					State:         gofakeit.State(),
					Country:       gofakeit.Country(),
				},
				Items: []*shippingpb.Item{
					{
						ProductId: gofakeit.UUID(),
						Quantity:  2,
					},
				},
			},
		},
		{
			name: "item without product_id",
			req: &shippingpb.GetQuoteRequest{
				Address: &shippingpb.Address{
					StreetAddress: gofakeit.Street(),
					City:          gofakeit.City(),
					State:         gofakeit.State(),
					Country:       gofakeit.Country(),
					ZipCode:       fakeZIP(ts.T()),
				},
				Items: []*shippingpb.Item{{Quantity: 2}},
			},
		},
		{
			name: "item with zero quantity",
			req: &shippingpb.GetQuoteRequest{
				Address: &shippingpb.Address{
					StreetAddress: gofakeit.Street(),
					City:          gofakeit.City(),
					State:         gofakeit.State(),
					Country:       gofakeit.Country(),
					ZipCode:       fakeZIP(ts.T()),
				},
				Items: []*shippingpb.Item{{ProductId: gofakeit.UUID()}},
			},
		},
	}
	for _, tt := range tests {
		ts.Run(tt.name, func() {
			resp, err := ts.grpcClient.GetQuote(ts.ctx, tt.req)
			ts.Nil(resp)
			ts.Error(err)
			s, _ := status.FromError(err)
			ts.Equal(codes.InvalidArgument.String(), s.Code().String())
		})
	}
}

func (ts *shippingTestSuite) TestShipOrder() {
	req := &shippingpb.ShipOrderRequest{
		Address: &shippingpb.Address{
			StreetAddress: gofakeit.Street(),
			City:          gofakeit.City(),
			State:         gofakeit.State(),
			Country:       gofakeit.Country(),
			ZipCode:       fakeZIP(ts.T()),
		},
		Items: []*shippingpb.Item{
			{
				ProductId: gofakeit.UUID(),
				Quantity:  2,
			},
		},
	}
	resp, err := ts.grpcClient.ShipOrder(ts.ctx, req)
	ts.NoError(err)
	ts.NotEmpty(resp.GetTrackingId())
}

func (ts *shippingTestSuite) TestShipOrder_InvalidData() {
	tests := []struct {
		name string
		req  *shippingpb.ShipOrderRequest
	}{
		{
			name: "street empty",
			req: &shippingpb.ShipOrderRequest{
				Address: &shippingpb.Address{
					City:    gofakeit.City(),
					State:   gofakeit.State(),
					Country: gofakeit.Country(),
					ZipCode: fakeZIP(ts.T()),
				},
				Items: []*shippingpb.Item{
					{
						ProductId: gofakeit.UUID(),
						Quantity:  2,
					},
				},
			},
		},
		{
			name: "street too short",
			req: &shippingpb.ShipOrderRequest{
				Address: &shippingpb.Address{
					StreetAddress: gofakeit.Letter(),
					City:          gofakeit.City(),
					State:         gofakeit.State(),
					Country:       gofakeit.Country(),
					ZipCode:       fakeZIP(ts.T()),
				},
				Items: []*shippingpb.Item{
					{
						ProductId: gofakeit.UUID(),
						Quantity:  2,
					},
				},
			},
		},
		{
			name: "street too long",
			req: &shippingpb.ShipOrderRequest{
				Address: &shippingpb.Address{
					StreetAddress: gofakeit.LetterN(1500),
					City:          gofakeit.City(),
					State:         gofakeit.State(),
					Country:       gofakeit.Country(),
					ZipCode:       fakeZIP(ts.T()),
				},
				Items: []*shippingpb.Item{
					{
						ProductId: gofakeit.UUID(),
						Quantity:  2,
					},
				},
			},
		},
		{
			name: "city empty",
			req: &shippingpb.ShipOrderRequest{
				Address: &shippingpb.Address{
					StreetAddress: gofakeit.Street(),
					State:         gofakeit.State(),
					Country:       gofakeit.Country(),
					ZipCode:       fakeZIP(ts.T()),
				},
				Items: []*shippingpb.Item{
					{
						ProductId: gofakeit.UUID(),
						Quantity:  2,
					},
				},
			},
		},
		{
			name: "city too short",
			req: &shippingpb.ShipOrderRequest{
				Address: &shippingpb.Address{
					StreetAddress: gofakeit.Street(),
					City:          gofakeit.Letter(),
					State:         gofakeit.State(),
					Country:       gofakeit.Country(),
					ZipCode:       fakeZIP(ts.T()),
				},
				Items: []*shippingpb.Item{
					{
						ProductId: gofakeit.UUID(),
						Quantity:  2,
					},
				},
			},
		},
		{
			name: "city too long",
			req: &shippingpb.ShipOrderRequest{
				Address: &shippingpb.Address{
					StreetAddress: gofakeit.Street(),
					City:          gofakeit.LetterN(500),
					State:         gofakeit.State(),
					Country:       gofakeit.Country(),
					ZipCode:       fakeZIP(ts.T()),
				},
				Items: []*shippingpb.Item{
					{
						ProductId: gofakeit.UUID(),
						Quantity:  2,
					},
				},
			},
		},
		{
			name: "state empty",
			req: &shippingpb.ShipOrderRequest{
				Address: &shippingpb.Address{
					StreetAddress: gofakeit.Street(),
					City:          gofakeit.City(),
					Country:       gofakeit.Country(),
					ZipCode:       fakeZIP(ts.T()),
				},
				Items: []*shippingpb.Item{
					{
						ProductId: gofakeit.UUID(),
						Quantity:  2,
					},
				},
			},
		},
		{
			name: "state too short",
			req: &shippingpb.ShipOrderRequest{
				Address: &shippingpb.Address{
					StreetAddress: gofakeit.Street(),
					City:          gofakeit.City(),
					State:         gofakeit.Letter(),
					Country:       gofakeit.Country(),
					ZipCode:       fakeZIP(ts.T()),
				},
				Items: []*shippingpb.Item{
					{
						ProductId: gofakeit.UUID(),
						Quantity:  2,
					},
				},
			},
		},
		{
			name: "state too long",
			req: &shippingpb.ShipOrderRequest{
				Address: &shippingpb.Address{
					StreetAddress: gofakeit.Street(),
					City:          gofakeit.City(),
					State:         gofakeit.LetterN(500),
					Country:       gofakeit.Country(),
					ZipCode:       fakeZIP(ts.T()),
				},
				Items: []*shippingpb.Item{
					{
						ProductId: gofakeit.UUID(),
						Quantity:  2,
					},
				},
			},
		},
		{
			name: "country empty",
			req: &shippingpb.ShipOrderRequest{
				Address: &shippingpb.Address{
					StreetAddress: gofakeit.Street(),
					City:          gofakeit.City(),
					State:         gofakeit.State(),
					ZipCode:       fakeZIP(ts.T()),
				},
				Items: []*shippingpb.Item{
					{
						ProductId: gofakeit.UUID(),
						Quantity:  2,
					},
				},
			},
		},
		{
			name: "country too short",
			req: &shippingpb.ShipOrderRequest{
				Address: &shippingpb.Address{
					StreetAddress: gofakeit.Street(),
					City:          gofakeit.City(),
					State:         gofakeit.State(),
					Country:       gofakeit.Letter(),
					ZipCode:       fakeZIP(ts.T()),
				},
				Items: []*shippingpb.Item{
					{
						ProductId: gofakeit.UUID(),
						Quantity:  2,
					},
				},
			},
		},
		{
			name: "country too long",
			req: &shippingpb.ShipOrderRequest{
				Address: &shippingpb.Address{
					StreetAddress: gofakeit.Street(),
					City:          gofakeit.City(),
					State:         gofakeit.State(),
					Country:       gofakeit.LetterN(500),
					ZipCode:       fakeZIP(ts.T()),
				},
				Items: []*shippingpb.Item{
					{
						ProductId: gofakeit.UUID(),
						Quantity:  2,
					},
				},
			},
		},
		{
			name: "zip empty",
			req: &shippingpb.ShipOrderRequest{
				Address: &shippingpb.Address{
					StreetAddress: gofakeit.Street(),
					City:          gofakeit.City(),
					State:         gofakeit.State(),
					Country:       gofakeit.Country(),
				},
				Items: []*shippingpb.Item{
					{
						ProductId: gofakeit.UUID(),
						Quantity:  2,
					},
				},
			},
		},
		{
			name: "item without product_id",
			req: &shippingpb.ShipOrderRequest{
				Address: &shippingpb.Address{
					StreetAddress: gofakeit.Street(),
					City:          gofakeit.City(),
					State:         gofakeit.State(),
					Country:       gofakeit.Country(),
					ZipCode:       fakeZIP(ts.T()),
				},
				Items: []*shippingpb.Item{{Quantity: 2}},
			},
		},
		{
			name: "item with zero quantity",
			req: &shippingpb.ShipOrderRequest{
				Address: &shippingpb.Address{
					StreetAddress: gofakeit.Street(),
					City:          gofakeit.City(),
					State:         gofakeit.State(),
					Country:       gofakeit.Country(),
					ZipCode:       fakeZIP(ts.T()),
				},
				Items: []*shippingpb.Item{{ProductId: gofakeit.UUID()}},
			},
		},
	}
	for _, tt := range tests {
		ts.Run(tt.name, func() {
			resp, err := ts.grpcClient.ShipOrder(ts.ctx, tt.req)
			ts.Nil(resp)
			ts.Error(err)
			s, _ := status.FromError(err)
			ts.Equal(codes.InvalidArgument.String(), s.Code().String())
		})
	}
}

func currGRPCClient(t *testing.T) shippingpb.ShippingServiceClient {
	t.Helper()
	conn, err := grpc.Dial(getConfig(t).GRPCAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	require.NoError(t, err)
	return shippingpb.NewShippingServiceClient(conn)
}

func fakeZIP(t *testing.T) int32 {
	t.Helper()
	digitZIP, err := strconv.Atoi(gofakeit.Zip())
	require.NoError(t, err, "failed convert string zip to digit zip")
	return int32(digitZIP)
}
