package repository

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"

	"github.com/demeero/shopagolic/cartservice/cart"
	"github.com/demeero/shopagolic/services/go/bricks/zaplogger"
)

type Loader struct {
	rds       redis.Cmdable
	keyPrefix string
	ttl       time.Duration
}

func NewLoader(keyPrefix string, ttl time.Duration, rds redis.Cmdable) *Loader {
	return &Loader{rds: rds, ttl: ttl, keyPrefix: keyPrefix}
}

func (l *Loader) LoadByUserID(ctx context.Context, userID string) (cart.Cart, error) {
	k := key(l.keyPrefix, userID)
	result, err := l.rds.HGetAll(ctx, k).Result()
	if err != nil {
		return cart.Cart{}, fmt.Errorf("failed exec HGetAll: %w", err)
	}
	if len(result) == 0 {
		return cart.Cart{}, fmt.Errorf("%w: %s", cart.ErrNotFound, userID)
	}
	if err := l.rds.Expire(ctx, k, l.ttl).Err(); err != nil {
		zaplogger.FromCtx(ctx).Error("failed exec expire", zap.Error(err), zap.String("user_id", userID))
	}
	return cart.Cart{
		UserID: userID,
		Items:  convertRedisCartItems(result, zaplogger.FromCtx(ctx)),
	}, nil
}

func convertRedisCartItems(items map[string]string, zlog *zap.Logger) []cart.Item {
	cartItems := make([]cart.Item, 0, len(items))
	for productID, quantity := range items {
		cartItems = append(cartItems, convertRedisCartItem(productID, quantity, zlog))
	}
	return cartItems
}

func convertRedisCartItem(productID string, quantity string, zlog *zap.Logger) cart.Item {
	n, err := strconv.ParseInt(quantity, 10, 64)
	if err != nil {
		n = 0
		zlog.Error("failed parse quantity: %w", zap.String("val", quantity), zap.Error(err))
	}
	return cart.Item{
		ProductID: productID,
		Quantity:  uint16(n),
	}
}
