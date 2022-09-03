package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.uber.org/zap"

	"github.com/demeero/shopagolic/productcatalog/catalog"
	"github.com/demeero/shopagolic/productcatalog/datagen"
	"github.com/demeero/shopagolic/productcatalog/internal/repository"
	"github.com/demeero/shopagolic/services/go/bricks/zaplogger"
)

type config struct {
	Mongo         mongoCfg
	ProductAmount int `default:"1000"`
}

type mongoCfg struct {
	DBName string `default:"shopagolic"`
	URI    string `default:"mongodb://localhost:27017"`
}

func main() {
	// Load environment variables from a `.env` file if one exists
	if err := godotenv.Load(); err != nil && !os.IsNotExist(err) {
		log.Fatalln("failed to load .env file", err)
	}

	zlog, err := zaplogger.NewZapLogger(zaplogger.ZapLoggingConfig{Level: os.Getenv("LOG_LEVEL")})
	if err != nil {
		log.Fatalln("failed to init zap logger", err)
	}
	defer func() {
		if err := zlog.Sync(); err != nil {
			log.Printf("failed to sync zap logger: %v", err)
		}
	}()

	cfg := config{}
	if err := envconfig.Process("", &cfg); err != nil {
		zlog.Fatal("failed init config", zap.Error(err))
	}

	mDB, closeDBFunc, err := mongoDBClient(cfg.Mongo)
	if err != nil {
		zlog.Fatal("failed connect to DB", zap.Error(err))
	}
	defer func() {
		if err := closeDBFunc(context.Background()); err != nil {
			zlog.Error("failed close mongo DB", zap.Error(err))
		}
	}()

	productColl := mDB.Database("shopagolic-catalog").Collection("products")

	creator := catalog.NewCreator(repository.NewCreator(productColl))
	loader := catalog.NewLoader(repository.NewLoader(productColl))
	gen := datagen.New(creator, loader)
	products, err := gen.Products(cfg.ProductAmount)
	if err != nil {
		zlog.Fatal("failed generate products", zap.Error(err))
	}
	zlog.Info("products generated", zap.Int("n", len(products)))
}

func mongoDBClient(cfg mongoCfg) (*mongo.Client, func(ctx context.Context) error, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
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
