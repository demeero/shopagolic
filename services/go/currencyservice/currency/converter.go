package currency

import (
	"context"
	"fmt"
	"math"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type ConverterRepository interface {
	LoadCurrencyConversion(ctx context.Context, currCode string) (float64, error)
}

type Money struct {
	CurrencyCode string
	Units        int64
	Nanos        int32
}

type Converter struct {
	repo ConverterRepository
}

func NewConverter(repo ConverterRepository) *Converter {
	return &Converter{repo: repo}
}

func (c *Converter) Convert(ctx context.Context, from Money, to string) (Money, error) {
	err := validation.Errors{
		"from_currency_code": validation.Validate(from.CurrencyCode, validation.Required, is.CurrencyCode),
		"to":                 validation.Validate(to, validation.Required, is.CurrencyCode),
	}.Filter()
	if err != nil {
		return Money{}, fmt.Errorf("%w: %s", ErrInvalidData, err)
	}
	fromVal, err := c.repo.LoadCurrencyConversion(ctx, from.CurrencyCode)
	if err != nil {
		return Money{}, fmt.Errorf("failed load currency conversion (%s): %w", from.CurrencyCode, err)
	}
	toVal, err := c.repo.LoadCurrencyConversion(ctx, to)
	if err != nil {
		return Money{}, fmt.Errorf("failed load currency conversion (%s): %w", to, err)
	}
	total := int64(math.Floor(float64(from.Units*10^9+int64(from.Nanos)) / fromVal * toVal))
	return Money{
		CurrencyCode: to,
		Units:        total / 1e9,
		Nanos:        int32(total % 1e9),
	}, nil
}
