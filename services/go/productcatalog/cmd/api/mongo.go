package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var dbConnectTimeout = 30 * time.Second

func mongoDBClient(cfg mongoCfg) (*mongo.Client, func(ctx context.Context) error, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbConnectTimeout)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.URI))
	if err != nil {
		return nil, nil, fmt.Errorf("failed connect MongoDB: %w", err)
	}
	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, nil, fmt.Errorf("failed ping MongoDB: %w", err)
	}
	return client, client.Disconnect, nil
}
