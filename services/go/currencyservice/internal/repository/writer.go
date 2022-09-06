package repository

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"

	"github.com/demeero/shopagolic/currencyservice/currency"
)

type Writer struct {
	rds       redis.Cmdable
	keyPrefix string
}

func NewWriter(keyPrefix string, rds redis.Cmdable) *Writer {
	return &Writer{keyPrefix: keyPrefix, rds: rds}
}

func (w *Writer) Put(ctx context.Context, currCode string, val float32) error {
	k := key(w.keyPrefix, currCode)
	v := fmt.Sprintf("%f", val)
	if err := w.rds.Set(ctx, k, v, 0).Err(); err != nil {
		return fmt.Errorf("failed exec set: %w", err)
	}
	return nil
}

func (w *Writer) Delete(ctx context.Context, currCode string) error {
	result, err := w.rds.Del(ctx, key(w.keyPrefix, currCode)).Result()
	if err != nil {
		return fmt.Errorf("failed exec del: %w", err)
	}
	if result == 0 {
		return fmt.Errorf("%w: %s", currency.ErrNotFound, currCode)
	}
	return nil
}
