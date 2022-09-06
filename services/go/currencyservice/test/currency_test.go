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

	currpb "github.com/demeero/shopagolic/services/proto/gen/go/shopagolic/currency/v1beta1"
	moneypb "github.com/demeero/shopagolic/services/proto/gen/go/shopagolic/money/v1"
)

func TestIntegrationCurrency(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	suite.Run(t, &currencyTestSuite{})
}

type currencyTestSuite struct {
	suite.Suite
	ctx        context.Context
	grpcClient currpb.CurrencyServiceClient
	rds        *redis.Client
	keyPrefix  string
}

func (ts *currencyTestSuite) SetupSuite() {
	ts.ctx = context.Background()
	c := getConfig(ts.T())
	ts.keyPrefix = c.CurrencyKeyPrefix
	ts.T().Logf("currency test config: %+v", c)
	ts.rds = redis.NewClient(&redis.Options{Addr: c.RedisAddr})
	ts.grpcClient = currGRPCClient(ts.T())
	ts.Require().NoError(ts.rds.FlushAll(ts.ctx).Err())
}

func (ts *currencyTestSuite) TearDownSuite() {
	ts.Require().NoError(ts.rds.Close())
}

func (ts *currencyTestSuite) TearDownTest() {
	ts.Require().NoError(ts.rds.FlushAll(ts.ctx).Err())
}

func (ts *currencyTestSuite) TestGetSupportedCurrencies_Empty() {
	resp, err := ts.grpcClient.GetSupportedCurrencies(ts.ctx, &currpb.GetSupportedCurrenciesRequest{})
	ts.NoError(err)
	ts.NotNil(resp)
	ts.Empty(resp.GetCurrencyCodes())
}

func (ts *currencyTestSuite) TestGetSupportedCurrencies() {
	var (
		curr1    = gofakeit.CurrencyShort()
		curr2    = gofakeit.CurrencyShort()
		expected = []string{curr1, curr2}
	)

	err := ts.rds.Set(ts.ctx, key(ts.keyPrefix, curr1), gofakeit.Float32(), 0).Err()
	ts.Require().NoError(err)
	ts.rds.Set(ts.ctx, key(ts.keyPrefix, curr2), gofakeit.Float32(), 0)
	ts.Require().NoError(err)

	resp, err := ts.grpcClient.GetSupportedCurrencies(ts.ctx, &currpb.GetSupportedCurrenciesRequest{})
	ts.NoError(err)
	ts.ElementsMatch(expected, resp.GetCurrencyCodes())
}

func (ts *currencyTestSuite) TestPutCurrency() {
	var (
		curr1 = gofakeit.CurrencyShort()
		val   = gofakeit.Float32()
	)

	resp, err := ts.grpcClient.PutCurrency(ts.ctx, &currpb.PutCurrencyRequest{
		CurrencyCode: curr1,
		Value:        val,
	})
	ts.NoError(err)
	ts.NotNil(resp)

	result, err := ts.rds.Get(ts.ctx, key(ts.keyPrefix, curr1)).Float32()
	ts.NoError(err)
	ts.Equal(val, result)
}

func (ts *currencyTestSuite) TestPutCurrency_InvalidInput() {
	resp, err := ts.grpcClient.PutCurrency(ts.ctx, &currpb.PutCurrencyRequest{CurrencyCode: "", Value: 3})
	ts.Error(err)
	ts.Equal(codes.InvalidArgument.String(), status.Code(err).String())
	ts.Nil(resp)
}

func (ts *currencyTestSuite) TestDeleteCurrency() {
	var (
		curr1 = gofakeit.CurrencyShort()
		curr2 = gofakeit.CurrencyShort()
		val2  = gofakeit.Float32()
	)

	err := ts.rds.Set(ts.ctx, key(ts.keyPrefix, curr1), gofakeit.Float32(), 0).Err()
	ts.Require().NoError(err)
	ts.rds.Set(ts.ctx, key(ts.keyPrefix, curr2), val2, 0)
	ts.Require().NoError(err)

	resp, err := ts.grpcClient.DeleteCurrency(ts.ctx, &currpb.DeleteCurrencyRequest{CurrencyCode: curr1})
	ts.NoError(err)
	ts.NotNil(resp)

	_, err = ts.rds.Get(ts.ctx, key(ts.keyPrefix, curr1)).Float32()
	ts.ErrorIs(err, redis.Nil)

	result, err := ts.rds.Get(ts.ctx, key(ts.keyPrefix, curr2)).Float32()
	ts.NoError(err)
	ts.Equal(val2, result)
}

func (ts *currencyTestSuite) TestDeleteCurrency_Notfound() {
	resp, err := ts.grpcClient.DeleteCurrency(ts.ctx, &currpb.DeleteCurrencyRequest{CurrencyCode: gofakeit.CurrencyShort()})
	ts.Error(err)
	ts.Nil(resp)
	ts.Equal(codes.NotFound.String(), status.Code(err).String())
}

func (ts *currencyTestSuite) TestDeleteCurrency_InvalidInput() {
	tests := []struct {
		req  *currpb.DeleteCurrencyRequest
		name string
	}{
		{
			name: "empty currency",
			req:  &currpb.DeleteCurrencyRequest{},
		},
		{
			name: "invalid currency",
			req:  &currpb.DeleteCurrencyRequest{CurrencyCode: gofakeit.Username()},
		},
	}
	for _, tt := range tests {
		ts.Run(tt.name, func() {
			resp, err := ts.grpcClient.DeleteCurrency(ts.ctx, tt.req)
			ts.Error(err)
			ts.Nil(resp)
			ts.Equal(codes.InvalidArgument.String(), status.Code(err).String())
		})
	}
}

func (ts *currencyTestSuite) TestConvert_NotFound_From() {
	curr1 := gofakeit.CurrencyShort()
	err := ts.rds.Set(ts.ctx, key(ts.keyPrefix, curr1), gofakeit.Float32(), 0).Err()
	ts.Require().NoError(err)

	resp, err := ts.grpcClient.Convert(ts.ctx, &currpb.ConvertRequest{
		From: &moneypb.Money{
			CurrencyCode: gofakeit.CurrencyShort(),
			Units:        1,
			Nanos:        2,
		},
		ToCode: curr1,
	})
	ts.Error(err)
	ts.Nil(resp)
	ts.Equal(codes.NotFound.String(), status.Code(err).String())
}

func (ts *currencyTestSuite) TestConvert_NotFound_To() {
	curr1 := gofakeit.CurrencyShort()
	err := ts.rds.Set(ts.ctx, key(ts.keyPrefix, curr1), gofakeit.Float32(), 0).Err()
	ts.Require().NoError(err)

	resp, err := ts.grpcClient.Convert(ts.ctx, &currpb.ConvertRequest{
		From: &moneypb.Money{
			CurrencyCode: curr1,
			Units:        1,
			Nanos:        2,
		},
		ToCode: gofakeit.CurrencyShort(),
	})
	ts.Error(err)
	ts.Nil(resp)
	ts.Equal(codes.NotFound.String(), status.Code(err).String())
}

func (ts *currencyTestSuite) TestConvert_NotFound_InvalidInput() {
	tests := []struct {
		req  *currpb.ConvertRequest
		name string
	}{
		{
			name: "empty to_code",
			req: &currpb.ConvertRequest{From: &moneypb.Money{
				CurrencyCode: "",
				Units:        1,
				Nanos:        2,
			}},
		},
		{
			name: "invalid to_code",
			req: &currpb.ConvertRequest{From: &moneypb.Money{
				CurrencyCode: "",
				Units:        1,
				Nanos:        2,
			}, ToCode: gofakeit.Username()},
		},
		{
			name: "empty from_code",
			req: &currpb.ConvertRequest{From: &moneypb.Money{
				Units: 1,
				Nanos: 2,
			}, ToCode: gofakeit.CurrencyShort()},
		},
		{
			name: "invalid from_code",
			req: &currpb.ConvertRequest{From: &moneypb.Money{
				Units: 1,
				Nanos: 2,
			}, ToCode: gofakeit.Username()},
		},
	}
	for _, tt := range tests {
		ts.Run(tt.name, func() {
			resp, err := ts.grpcClient.Convert(ts.ctx, tt.req)
			ts.Error(err)
			ts.Nil(resp)
			ts.Equal(codes.InvalidArgument.String(), status.Code(err).String())
		})
	}
}

func currGRPCClient(t *testing.T) currpb.CurrencyServiceClient {
	t.Helper()
	conn, err := grpc.Dial(getConfig(t).GRPCAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	require.NoError(t, err)
	return currpb.NewCurrencyServiceClient(conn)
}

func key(prefix, val string) string {
	return fmt.Sprintf("%s:currency_conversion:%s", prefix, val)
}
