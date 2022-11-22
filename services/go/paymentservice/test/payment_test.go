package test

import (
	"context"
	"math"
	"testing"

	moneypb "github.com/demeero/shopagolic/services/proto/gen/go/shopagolic/money/v1"
	paymentpb "github.com/demeero/shopagolic/services/proto/gen/go/shopagolic/payment/v1beta1"
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
	grpcClient paymentpb.PaymentServiceClient
}

func (ts *recommendationTestSuite) SetupSuite() {
	ts.ctx = context.Background()
	c := getConfig(ts.T())
	ts.T().Logf("payment test config: %+v", c)
	ts.grpcClient = paymentGRPCClient(ts.T())
}

func (ts *recommendationTestSuite) TestCharge() {
	tests := []struct {
		name string
		req  *paymentpb.ChargeRequest
	}{
		{
			name: "card1",
			req: &paymentpb.ChargeRequest{
				Amount: &moneypb.Money{
					CurrencyCode: "USD",
					Units:        1,
					Nanos:        2,
				},
				CreditCard: &paymentpb.CreditCardInfo{
					Number:          "4595347396176636",
					Cvv:             313,
					ExpirationYear:  2023,
					ExpirationMonth: 7,
				},
			},
		},
		{
			name: "card2",
			req: &paymentpb.ChargeRequest{
				Amount: &moneypb.Money{
					CurrencyCode: "UAH",
					Units:        0,
					Nanos:        1,
				},
				CreditCard: &paymentpb.CreditCardInfo{
					Number:          "5127086932981018",
					Cvv:             676,
					ExpirationYear:  2026,
					ExpirationMonth: 11,
				},
			},
		},
		{
			name: "card3",
			req: &paymentpb.ChargeRequest{
				Amount: &moneypb.Money{
					CurrencyCode: "EUR",
					Units:        3,
					Nanos:        22,
				},
				CreditCard: &paymentpb.CreditCardInfo{
					Number:          "4117498921282644",
					Cvv:             756,
					ExpirationYear:  2025,
					ExpirationMonth: 11,
				},
			},
		},
	}
	for _, tt := range tests {
		ts.Run(tt.name, func() {
			actual, err := ts.grpcClient.Charge(ts.ctx, tt.req)
			ts.NoError(err)
			ts.NotEmpty(actual.GetTransactionId())
		})
	}
}

func (ts *recommendationTestSuite) TestCharge_InvalidArgErr() {
	tests := []struct {
		name string
		req  *paymentpb.ChargeRequest
	}{
		{
			name: "invalid-card-num-1",
			req: &paymentpb.ChargeRequest{
				Amount: &moneypb.Money{
					CurrencyCode: "EUR",
					Units:        3,
					Nanos:        22,
				},
				CreditCard: &paymentpb.CreditCardInfo{
					Number:          "",
					Cvv:             756,
					ExpirationYear:  2040,
					ExpirationMonth: 1,
				},
			},
		},
		{
			name: "invalid-card-num-2",
			req: &paymentpb.ChargeRequest{
				Amount: &moneypb.Money{
					CurrencyCode: "EUR",
					Units:        3,
					Nanos:        22,
				},
				CreditCard: &paymentpb.CreditCardInfo{
					Number:          "23524",
					Cvv:             756,
					ExpirationYear:  2040,
					ExpirationMonth: 1,
				},
			},
		},
		{
			name: "invalid-card-num-3",
			req: &paymentpb.ChargeRequest{
				Amount: &moneypb.Money{
					CurrencyCode: "EUR",
					Units:        3,
					Nanos:        22,
				},
				CreditCard: &paymentpb.CreditCardInfo{
					Number:          "4117498921f82644",
					Cvv:             756,
					ExpirationYear:  2040,
					ExpirationMonth: 1,
				},
			},
		},
		{
			name: "invalid-year-1",
			req: &paymentpb.ChargeRequest{
				Amount: &moneypb.Money{
					CurrencyCode: "EUR",
					Units:        3,
					Nanos:        22,
				},
				CreditCard: &paymentpb.CreditCardInfo{
					Number:          "4117498921282644",
					Cvv:             756,
					ExpirationYear:  math.MaxInt32,
					ExpirationMonth: 1,
				},
			},
		},
		{
			name: "invalid-year-2",
			req: &paymentpb.ChargeRequest{
				Amount: &moneypb.Money{
					CurrencyCode: "EUR",
					Units:        3,
					Nanos:        22,
				},
				CreditCard: &paymentpb.CreditCardInfo{
					Number:          "4117498921282644",
					Cvv:             756,
					ExpirationYear:  0,
					ExpirationMonth: 1,
				},
			},
		},
		{
			name: "invalid-year-3",
			req: &paymentpb.ChargeRequest{
				Amount: &moneypb.Money{
					CurrencyCode: "EUR",
					Units:        3,
					Nanos:        22,
				},
				CreditCard: &paymentpb.CreditCardInfo{
					Number:          "4117498921282644",
					Cvv:             756,
					ExpirationYear:  -1,
					ExpirationMonth: 1,
				},
			},
		},
		{
			name: "invalid-month-1",
			req: &paymentpb.ChargeRequest{
				Amount: &moneypb.Money{
					CurrencyCode: "EUR",
					Units:        3,
					Nanos:        22,
				},
				CreditCard: &paymentpb.CreditCardInfo{
					Number:          "4117498921282644",
					Cvv:             756,
					ExpirationYear:  2025,
					ExpirationMonth: -1,
				},
			},
		},
		{
			name: "invalid-month-2",
			req: &paymentpb.ChargeRequest{
				Amount: &moneypb.Money{
					CurrencyCode: "EUR",
					Units:        3,
					Nanos:        22,
				},
				CreditCard: &paymentpb.CreditCardInfo{
					Number:          "4117498921282644",
					Cvv:             756,
					ExpirationYear:  2025,
					ExpirationMonth: 0,
				},
			},
		},
		{
			name: "invalid-month-3",
			req: &paymentpb.ChargeRequest{
				Amount: &moneypb.Money{
					CurrencyCode: "EUR",
					Units:        3,
					Nanos:        22,
				},
				CreditCard: &paymentpb.CreditCardInfo{
					Number:          "4117498921282644",
					Cvv:             756,
					ExpirationYear:  2025,
					ExpirationMonth: 13,
				},
			},
		},
		{
			name: "invalid-cvv-1",
			req: &paymentpb.ChargeRequest{
				Amount: &moneypb.Money{
					CurrencyCode: "EUR",
					Units:        3,
					Nanos:        22,
				},
				CreditCard: &paymentpb.CreditCardInfo{
					Number:          "4117498921282644",
					Cvv:             -1,
					ExpirationYear:  2040,
					ExpirationMonth: 1,
				},
			},
		},
		{
			name: "invalid-cvv-2",
			req: &paymentpb.ChargeRequest{
				Amount: &moneypb.Money{
					CurrencyCode: "EUR",
					Units:        3,
					Nanos:        22,
				},
				CreditCard: &paymentpb.CreditCardInfo{
					Number:          "4117498921282644",
					Cvv:             0,
					ExpirationYear:  2040,
					ExpirationMonth: 1,
				},
			},
		},
		{
			name: "invalid-cvv-3",
			req: &paymentpb.ChargeRequest{
				Amount: &moneypb.Money{
					CurrencyCode: "EUR",
					Units:        3,
					Nanos:        22,
				},
				CreditCard: &paymentpb.CreditCardInfo{
					Number:          "4117498921282644",
					Cvv:             math.MaxInt32,
					ExpirationYear:  2040,
					ExpirationMonth: 1,
				},
			},
		},
	}
	for _, tt := range tests {
		ts.Run(tt.name, func() {
			actual, err := ts.grpcClient.Charge(ts.ctx, tt.req)
			ts.Nil(actual)
			ts.Error(err)
			s, _ := status.FromError(err)
			ts.Equal(codes.InvalidArgument, s.Code())
		})
	}
}

func paymentGRPCClient(t *testing.T) paymentpb.PaymentServiceClient {
	t.Helper()
	conn, err := grpc.Dial(getConfig(t).GRPCAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	require.NoError(t, err)
	return paymentpb.NewPaymentServiceClient(conn)
}
