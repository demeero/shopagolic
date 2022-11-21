package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/demeero/shopagolic/recommendationservice/internal/rpc"
	"github.com/demeero/shopagolic/recommendationservice/recommendation"
	"github.com/demeero/shopagolic/services/go/bricks/zaplogger"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"
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
	zlog = zlog.With(zap.String("service_name", "recommendation"))

	cfg := config{}
	if err := envconfig.Process("", &cfg); err != nil {
		zlog.Fatal("failed init config", zap.Error(err))
	}

	zlog.Info("configuration initialized")

	zlog.Info("initializing product catalog client")
	catalogLoader, closeCataloFunc, err := createProductCatalogLoader(cfg.ProductCatalogLoaderClient)
	if err != nil {
		zlog.Fatal("failed create product catalog loader", zap.Error(err))
	}
	catalogLoaderHealth, closeCatalogHealthFunc, err := createProductCatalogHealthClient(cfg.ProductCatalogLoaderClient)
	if err != nil {
		zlog.Fatal("failed create catalog loader health client", zap.Error(err))
	}
	zlog.Info("product catalog client initialized")

	components := rpc.Components{
		Loader:              recommendation.NewLoader(catalogLoader),
		CatalogHealthClient: catalogLoaderHealth,
	}
	grpcStopFunc := grpcServ(cfg.GRPC, components, zlog)
	zlog.Info("GRPC server established")

	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-interruptChan

	zlog.Info("application stopping")

	var appStoppedWithErr bool

	if err := closeCataloFunc(); err != nil {
		zlog.Error("failed close catalog", zap.Error(err))
		appStoppedWithErr = true
	}

	if err := closeCatalogHealthFunc(); err != nil {
		zlog.Error("failed close catalog health", zap.Error(err))
		appStoppedWithErr = true
	}

	if appStoppedWithErr {
		zlog.Fatal("failed stop app gracefully")
	}

	grpcStopFunc()

	zlog.Info("application stopped gracefully")
}
