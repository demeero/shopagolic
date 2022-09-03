package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"

	"github.com/demeero/shopagolic/productcatalog/catalog"
	"github.com/demeero/shopagolic/productcatalog/internal/repository"
	"github.com/demeero/shopagolic/productcatalog/internal/rpc"
	"github.com/demeero/shopagolic/services/go/bricks/zaplogger"
)

var (
	dbConnectTimeout  = 30 * time.Second
	dbShutdownTimeout = 20 * time.Second
)

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
	zlog = zlog.With(zap.String("service_name", "productcatalog"))

	cfg := config{}
	if err := envconfig.Process("", &cfg); err != nil {
		zlog.Fatal("failed init config", zap.Error(err))
	}

	zlog.Info("configuration initialized")

	mDB, closeDBFunc, err := mongoDBClient(cfg.Mongo)
	if err != nil {
		zlog.Fatal("failed connect to DB", zap.Error(err))
	}
	productColl := mDB.Database("shopagolic-catalog").Collection("products")
	zlog.Info("database connection established")

	components := rpc.ProductComponents{
		CatalogLoader:   catalog.NewLoader(repository.NewLoader(productColl)),
		CatalogSearcher: catalog.NewSearcher(repository.NewSearcher(productColl, 10)),
		CatalogCreator:  catalog.NewCreator(repository.NewCreator(productColl)),
	}
	grpcStopFunc := grpcServ(cfg.GRPC, components, mDB, zlog)
	zlog.Info("GRPC server established")

	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-interruptChan

	zlog.Info("application stopping")

	ctx, cancel := context.WithTimeout(context.Background(), dbShutdownTimeout)
	defer cancel()

	var appStoppedWithErr bool

	if err := closeDBFunc(ctx); err != nil {
		zlog.Error("failed disconnect MongoDB", zap.Error(err))
		appStoppedWithErr = true
	}

	if appStoppedWithErr {
		zlog.Fatal("failed stop app gracefully")
	}

	grpcStopFunc()

	zlog.Info("application stopped gracefully")
}
