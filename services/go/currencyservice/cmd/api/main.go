package main

import (
	"context"
	_ "embed"
	"encoding/json"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"

	"github.com/demeero/shopagolic/currencyservice/currency"
	"github.com/demeero/shopagolic/currencyservice/internal/repository"
	"github.com/demeero/shopagolic/currencyservice/internal/rpc"
	"github.com/demeero/shopagolic/services/go/bricks/zaplogger"
)

//go:embed currency_conversion_fixture.json
var currencyConversionFixture string

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
	zlog = zlog.With(zap.String("service_name", "currency"))

	cfg := config{}
	if err := envconfig.Process("", &cfg); err != nil {
		zlog.Fatal("failed init config", zap.Error(err))
	}

	zlog.Info("configuration initialized")

	rds, closeRedisFunc, err := rdsClient(cfg.Redis)
	if err != nil {
		zlog.Fatal("failed init redis client", zap.Error(err))
	}
	zlog.Info("Redis connection established")

	components := rpc.CurrencyComponents{
		Converter: currency.NewConverter(repository.NewConverter(cfg.Redis.CurrencyKeyPrefix, rds)),
		Loader:    currency.NewLoader(repository.NewLoader(cfg.Redis.CurrencyKeyPrefix, rds)),
		Writer:    currency.NewWriter(repository.NewWriter(cfg.Redis.CurrencyKeyPrefix, rds)),
	}

	if cfg.InitCurrenciesIfEmpty {
		go initCurrencies(components.Writer, components.Loader, zlog)
	}

	grpcStopFunc := grpcServ(cfg.GRPC, components, rds, zlog)
	zlog.Info("GRPC server established")

	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-interruptChan

	zlog.Info("application stopping")

	var appStoppedWithErr bool

	if err := closeRedisFunc(); err != nil {
		zlog.Error("failed disconnect Redis", zap.Error(err))
		appStoppedWithErr = true
	}

	if appStoppedWithErr {
		zlog.Fatal("failed stop app gracefully")
	}

	grpcStopFunc()

	zlog.Info("application stopped gracefully")
}

func initCurrencies(writer *currency.Writer, loader *currency.Loader, zlog *zap.Logger) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*50)
	defer cancel()
	currCodes, err := loader.LoadCurrencyCodes(ctx)
	if err != nil {
		zlog.Error("failed load currency codes for initialization", zap.Error(err))
		return
	}
	if len(currCodes) > 0 {
		zlog.Info("currency storage not empty - skip currency initialization")
		return
	}
	currencies := map[string]float32{}
	if err := json.Unmarshal([]byte(currencyConversionFixture), &currencies); err != nil {
		zlog.Error("failed decode currency_conversion_fixture", zap.Error(err))
		return
	}
	for currCode, val := range currencies {
		if err := writer.Put(ctx, currCode, val); err != nil {
			zlog.Error("failed put currency during initialization", zap.Error(err), zap.String("currency", currCode))
			return
		}
	}
	zlog.Info("currency_conversion_fixture initialized")
}
