package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/go-redis/redis/v8"
)

type Loader struct {
	rds       redis.Cmdable
	keyPrefix string
}

func NewLoader(keyPrefix string, rds redis.Cmdable) *Loader {
	return &Loader{
		rds:       rds,
		keyPrefix: keyPrefix,
	}
}

func (l *Loader) LoadCurrencyCodes(ctx context.Context) ([]string, error) {
	result, err := l.rds.Keys(ctx, l.keyPrefix+"*").Result()
	if err != nil {
		return nil, fmt.Errorf("failed exec keys: %w", err)
	}
	for i, k := range result {
		result[i] = strings.TrimPrefix(k, fmt.Sprintf(l.keyPrefix+internalPrefix))
	}
	return result, nil
}
