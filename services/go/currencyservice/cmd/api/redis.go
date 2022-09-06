package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

var rdsPingTimeout = 30 * time.Second

func rdsClient(cfg rdsCfg) (*redis.Client, func() error, error) {
	ctx, cancel := context.WithTimeout(context.Background(), rdsPingTimeout)
	defer cancel()
	rdb := redis.NewClient(&redis.Options{Addr: cfg.Addr})
	if err := rdb.Ping(ctx).Err(); err != nil {
		return nil, nil, fmt.Errorf("failed ping redis db: %w", err)
	}

	return rdb, func() error {
		return rdb.Close()
	}, nil
}
