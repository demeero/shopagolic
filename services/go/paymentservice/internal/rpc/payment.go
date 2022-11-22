package rpc

import (
	"context"
	"errors"
	"math"

	"github.com/demeero/shopagolic/paymentservice/payment"
	"github.com/demeero/shopagolic/services/go/bricks/zaplogger"
	paymentpb "github.com/demeero/shopagolic/services/proto/gen/go/shopagolic/payment/v1beta1"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Payment struct {
	paymentpb.UnimplementedPaymentServiceServer
	charger *payment.Charger
}

func NewPayment(charger *payment.Charger) *Payment {
	return &Payment{charger: charger}
}

func (p *Payment) Charge(ctx context.Context, req *paymentpb.ChargeRequest) (*paymentpb.ChargeResponse, error) {
	if err := p.validateCR(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	txID, err := p.charger.Charge(ctx, p.convertCRFromProto(req))
	if err := errHandler(ctx, err); err != nil {
		return nil, err
	}
	return &paymentpb.ChargeResponse{TransactionId: txID}, nil
}

func (p *Payment) validateCR(req *paymentpb.ChargeRequest) error {
	return validation.Errors{
		"cvv":   validation.Validate(req.GetCreditCard().GetCvv(), validation.Min(0), validation.Max(math.MaxUint16)),
		"year":  validation.Validate(req.GetCreditCard().GetExpirationYear(), validation.Min(0), validation.Max(math.MaxUint16)),
		"month": validation.Validate(req.GetCreditCard().GetExpirationMonth(), validation.Min(0), validation.Max(math.MaxUint8)),
	}.Filter()
}

func (p *Payment) convertCRFromProto(req *paymentpb.ChargeRequest) payment.ChargeInput {
	card := req.GetCreditCard()
	amount := req.GetAmount()
	return payment.ChargeInput{
		CreditCard: payment.CreditCard{
			Number:   card.GetNumber(),
			CVV:      uint16(card.GetCvv()),
			ExpYear:  uint16(card.GetExpirationYear()),
			ExpMonth: uint8(card.GetExpirationMonth()),
		},
		Amount: payment.Money{
			CurrencyCode: amount.GetCurrencyCode(),
			Units:        amount.GetUnits(),
			Nanos:        amount.GetNanos(),
		},
	}
}

func errHandler(ctx context.Context, err error) error {
	if err == nil {
		return nil
	}
	if errors.Is(err, payment.ErrInvalidData) {
		return status.Error(codes.InvalidArgument, err.Error())
	}
	zaplogger.FromCtx(ctx).Error("failed handle request", zap.Error(err))
	return err
}
