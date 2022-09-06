package rpc

import (
	"context"
	"errors"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/demeero/shopagolic/services/go/bricks/zaplogger"
	moneypb "github.com/demeero/shopagolic/services/proto/gen/go/shopagolic/money/v1"
	shippingpb "github.com/demeero/shopagolic/services/proto/gen/go/shopagolic/shipping/v1beta1"
	"github.com/demeero/shopagolic/shippingservice/shipping"
)

type Shipping struct {
	shippingpb.UnimplementedShippingServiceServer
	shipper *shipping.Shipper
}

func NewShipping(shipper *shipping.Shipper) *Shipping {
	return &Shipping{shipper: shipper}
}

func (s *Shipping) GetQuote(ctx context.Context, req *shippingpb.GetQuoteRequest) (*shippingpb.GetQuoteResponse, error) {
	result, err := s.shipper.Quote(ctx, convertProtoAddress(req.GetAddress()), convertProtoItems(req.GetItems()))
	if err := errHandler(ctx, err); err != nil {
		return nil, err
	}
	return &shippingpb.GetQuoteResponse{Cost: convertMoney(result)}, nil
}

func (s *Shipping) ShipOrder(ctx context.Context, req *shippingpb.ShipOrderRequest) (*shippingpb.ShipOrderResponse, error) {
	trackID, err := s.shipper.ShipOrder(ctx, convertProtoAddress(req.GetAddress()), convertProtoItems(req.GetItems()))
	if err := errHandler(ctx, err); err != nil {
		return nil, err
	}
	return &shippingpb.ShipOrderResponse{TrackingId: trackID}, nil
}

func convertProtoAddress(address *shippingpb.Address) shipping.Address {
	return shipping.Address{
		Street:  address.GetStreetAddress(),
		City:    address.GetCity(),
		State:   address.GetState(),
		Country: address.GetCountry(),
		ZIPCode: address.GetZipCode(),
	}
}

func convertMoney(m shipping.Money) *moneypb.Money {
	return &moneypb.Money{
		CurrencyCode: m.CurrencyCode,
		Units:        m.Units,
		Nanos:        m.Nanos,
	}
}

func convertProtoItems(items []*shippingpb.Item) []shipping.Item {
	result := make([]shipping.Item, 0, len(items))
	for _, i := range items {
		result = append(result, convertProtoItem(i))
	}
	return result
}

func convertProtoItem(item *shippingpb.Item) shipping.Item {
	return shipping.Item{
		ProductID: item.GetProductId(),
		Quantity:  uint32(item.GetQuantity()),
	}
}

func errHandler(ctx context.Context, err error) error {
	if err == nil {
		return nil
	}
	if errors.Is(err, shipping.ErrInvalidData) {
		return status.Error(codes.InvalidArgument, err.Error())
	}
	zaplogger.FromCtx(ctx).Error("failed handle request", zap.Error(err))
	return err
}
