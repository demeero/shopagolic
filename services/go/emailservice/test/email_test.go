package test

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	emailpb "github.com/demeero/shopagolic/services/proto/gen/go/shopagolic/email/v1beta1"
	moneypb "github.com/demeero/shopagolic/services/proto/gen/go/shopagolic/money/v1"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

func TestIntegrationEmailConfirmation(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	suite.Run(t, &emailConfirmationTestSuite{})
}

type emailConfirmationTestSuite struct {
	suite.Suite
	ctx        context.Context
	grpcClient emailpb.EmailServiceClient
}

func (ts *emailConfirmationTestSuite) SetupSuite() {
	ts.ctx = context.Background()
	c := getConfig(ts.T())
	ts.T().Logf("email test config: %+v", c)
	ts.grpcClient = emailGRPCClient(ts.T())
}

func (ts *emailConfirmationTestSuite) TestEmailConfirmation() {
	actual, err := ts.grpcClient.SendOrderConfirmation(context.Background(), &emailpb.SendOrderConfirmationRequest{
		Email: gofakeit.Email(),
		Order: &emailpb.OrderResult{
			OrderId:            gofakeit.UUID(),
			ShippingTrackingId: gofakeit.UUID(),
			ShippingCost: &moneypb.Money{
				CurrencyCode: gofakeit.CurrencyShort(),
				Units:        1,
				Nanos:        2,
			},
			ShippingAddress: &emailpb.Address{
				StreetAddress: gofakeit.Street(),
				City:          gofakeit.City(),
				State:         gofakeit.State(),
				Country:       gofakeit.Country(),
				ZipCode:       int32(gofakeit.Number(1000, 9999)),
			},
			Items: []*emailpb.OrderItem{
				{
					Item: &emailpb.CartItem{
						ProductId: gofakeit.UUID(),
						Quantity:  2,
					},
					Cost: &moneypb.Money{
						CurrencyCode: gofakeit.CurrencyShort(),
						Units:        2,
						Nanos:        1,
					},
				},
			},
		},
	})
	ts.NoError(err)
	ts.NotNil(actual)
}

func (ts *emailConfirmationTestSuite) TestEmailConfirmation_InvalidArgErr() {
	invalidEmails := []string{"", "Abc.example.com", "A@b@c@example.com", `a"b(c)d,e:f;g<h>i[j\k]l@example.com`,
		`just"not"right@example.com`, `his is"not\allowed@example.com`, `this\ still\"notallowed@example.com`}
	for _, email := range invalidEmails {
		ts.Run(email, func() {
			actual, err := ts.grpcClient.SendOrderConfirmation(context.Background(), &emailpb.SendOrderConfirmationRequest{
				Email: email,
				Order: &emailpb.OrderResult{
					OrderId:            gofakeit.UUID(),
					ShippingTrackingId: gofakeit.UUID(),
					ShippingCost: &moneypb.Money{
						CurrencyCode: gofakeit.CurrencyShort(),
						Units:        1,
						Nanos:        2,
					},
					ShippingAddress: &emailpb.Address{
						StreetAddress: gofakeit.Street(),
						City:          gofakeit.City(),
						State:         gofakeit.State(),
						Country:       gofakeit.Country(),
						ZipCode:       int32(gofakeit.Number(1000, 9999)),
					},
					Items: []*emailpb.OrderItem{
						{
							Item: &emailpb.CartItem{
								ProductId: gofakeit.UUID(),
								Quantity:  2,
							},
							Cost: &moneypb.Money{
								CurrencyCode: gofakeit.CurrencyShort(),
								Units:        2,
								Nanos:        1,
							},
						},
					},
				},
			})
			ts.Nil(actual)
			ts.Error(err)
			s, _ := status.FromError(err)
			ts.Equal(codes.InvalidArgument, s.Code())
		})
	}
}

func emailGRPCClient(t *testing.T) emailpb.EmailServiceClient {
	t.Helper()
	conn, err := grpc.Dial(getConfig(t).GRPCAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	require.NoError(t, err)
	return emailpb.NewEmailServiceClient(conn)
}
