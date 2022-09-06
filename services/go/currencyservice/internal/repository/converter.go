package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-redis/redis/v8"

	"github.com/demeero/shopagolic/currencyservice/currency"
)

type Converter struct {
	rds       redis.Cmdable
	keyPrefix string
}

func NewConverter(keyPrefix string, rds redis.Cmdable) *Converter {
	return &Converter{
		rds:       rds,
		keyPrefix: keyPrefix,
	}
}

func (c *Converter) LoadCurrencyConversion(ctx context.Context, currCode string) (float64, error) {
	result, err := c.rds.Get(ctx, key(c.keyPrefix, currCode)).Float64()
	if errors.Is(err, redis.Nil) {
		return 0, currency.ErrNotFound
	}
	if err != nil {
		return 0, fmt.Errorf("failed exec get: %w", err)
	}
	return result, nil
}
