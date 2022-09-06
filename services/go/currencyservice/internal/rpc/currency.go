package rpc

import (
	"context"
	"errors"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/demeero/shopagolic/currencyservice/currency"
	"github.com/demeero/shopagolic/services/go/bricks/zaplogger"
	currpb "github.com/demeero/shopagolic/services/proto/gen/go/shopagolic/currency/v1beta1"
	moneypb "github.com/demeero/shopagolic/services/proto/gen/go/shopagolic/money/v1"
)

type CurrencyComponents struct {
	Loader    *currency.Loader
	Converter *currency.Converter
	Writer    *currency.Writer
}

type Currency struct {
	currpb.UnimplementedCurrencyServiceServer
	loader    *currency.Loader
	converter *currency.Converter
	writer    *currency.Writer
}

func NewCurrency(components CurrencyComponents) *Currency {
	return &Currency{
		loader:    components.Loader,
		converter: components.Converter,
		writer:    components.Writer,
	}
}

func (c *Currency) GetSupportedCurrencies(ctx context.Context, _ *currpb.GetSupportedCurrenciesRequest) (*currpb.GetSupportedCurrenciesResponse, error) {
	currCodes, err := c.loader.LoadCurrencyCodes(ctx)
	if err := errHandler(ctx, err); err != nil {
		return nil, err
	}
	return &currpb.GetSupportedCurrenciesResponse{CurrencyCodes: currCodes}, nil
}

func (c *Currency) Convert(ctx context.Context, req *currpb.ConvertRequest) (*currpb.ConvertResponse, error) {
	conversion, err := c.converter.Convert(ctx, convertProtoMoney(req.GetFrom()), req.GetToCode())
	if err := errHandler(ctx, err); err != nil {
		return nil, err
	}
	return &currpb.ConvertResponse{ConversionResult: convertMoney(conversion)}, nil
}

func (c *Currency) PutCurrency(ctx context.Context, req *currpb.PutCurrencyRequest) (*currpb.PutCurrencyResponse, error) {
	if err := errHandler(ctx, c.writer.Put(ctx, req.GetCurrencyCode(), req.GetValue())); err != nil {
		return nil, err
	}
	return &currpb.PutCurrencyResponse{}, nil
}

func (c *Currency) DeleteCurrency(ctx context.Context, req *currpb.DeleteCurrencyRequest) (*currpb.DeleteCurrencyResponse, error) {
	if err := errHandler(ctx, c.writer.Delete(ctx, req.GetCurrencyCode())); err != nil {
		return nil, err
	}
	return &currpb.DeleteCurrencyResponse{}, nil
}

func convertMoney(m currency.Money) *moneypb.Money {
	return &moneypb.Money{
		CurrencyCode: m.CurrencyCode,
		Units:        m.Units,
		Nanos:        m.Nanos,
	}
}

func convertProtoMoney(m *moneypb.Money) currency.Money {
	return currency.Money{
		CurrencyCode: m.GetCurrencyCode(),
		Units:        m.GetUnits(),
		Nanos:        m.GetNanos(),
	}
}
func errHandler(ctx context.Context, err error) error {
	if err == nil {
		return nil
	}
	if errors.Is(err, currency.ErrNotFound) {
		return status.Error(codes.NotFound, err.Error())
	}
	if errors.Is(err, currency.ErrInvalidData) {
		return status.Error(codes.InvalidArgument, err.Error())
	}
	if errors.Is(err, currency.ErrConflictData) {
		return status.Error(codes.AlreadyExists, err.Error())
	}
	zaplogger.FromCtx(ctx).Error("failed handle request", zap.Error(err))
	return err
}
