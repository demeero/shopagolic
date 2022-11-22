package payment

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/demeero/shopagolic/services/go/bricks/zaplogger"
	creditcard "github.com/durango/go-credit-card"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

var ErrInvalidData = errors.New("invalid data")

type Money struct {
	CurrencyCode string
	Units        int64
	Nanos        int32
}

type ChargeInput struct {
	CreditCard CreditCard
	Amount     Money
}

type CreditCard struct {
	Number   string
	CVV      uint16
	ExpYear  uint16
	ExpMonth uint8
}

type Charger struct {
}

func NewCharger() *Charger {
	return &Charger{}
}

func (c *Charger) Charge(ctx context.Context, in ChargeInput) (string, error) {
	card := creditcard.Card{
		Number: in.CreditCard.Number,
		Cvv:    strconv.Itoa(int(in.CreditCard.CVV)),
		Year:   strconv.Itoa(int(in.CreditCard.ExpYear)),
		Month:  strconv.Itoa(int(in.CreditCard.ExpMonth)),
	}
	if err := card.Validate(); err != nil {
		return "", fmt.Errorf("%w: %s", ErrInvalidData, err)
	}

	lastFour, err := card.LastFourDigits()
	if err != nil {
		return "", fmt.Errorf("failed get last four card digits: %w", err)
	}

	transactionID := uuid.NewString()
	zaplogger.FromCtx(ctx).Info("transaction processed",
		zap.Any("amount", in.Amount),
		zap.Any("card_num", "***"+lastFour),
		zap.Any("id", transactionID))
	return transactionID, nil
}
