package repository

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"

	"github.com/demeero/shopagolic/cartservice/cart"
)

type Deleter struct {
	rds       redis.Cmdable
	keyPrefix string
}

func NewDeleter(keyPrefix string, rds redis.Cmdable) *Deleter {
	return &Deleter{rds: rds, keyPrefix: keyPrefix}
}

func (d *Deleter) DeleteAll(ctx context.Context, userID string) error {
	result, err := d.rds.Del(ctx, key(d.keyPrefix, userID)).Result()
	if err != nil {
		return fmt.Errorf("failed exec del: %w", err)
	}
	if result == 0 {
		return fmt.Errorf("%w: %s", cart.ErrNotFound, userID)
	}
	return nil
}
