package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"

	"github.com/demeero/shopagolic/cartservice/cart"
	"github.com/demeero/shopagolic/services/go/bricks/zaplogger"
)

type Adder struct {
	rds       redis.Cmdable
	keyPrefix string
	ttl       time.Duration
}

func NewAdder(keyPrefix string, ttl time.Duration, rds redis.Cmdable) *Adder {
	return &Adder{rds: rds, ttl: ttl, keyPrefix: keyPrefix}
}

func (a *Adder) AddItem(ctx context.Context, userID string, item cart.Item) error {
	k := key(a.keyPrefix, userID)
	if err := a.rds.HSet(ctx, k, item.ProductID, item.Quantity).Err(); err != nil {
		return fmt.Errorf("failed exec hset: %w", err)
	}
	if err := a.rds.Expire(ctx, k, a.ttl).Err(); err != nil {
		zaplogger.FromCtx(ctx).Error("failed exec expire", zap.Error(err), zap.String("user_id", userID))
	}
	return nil
}
