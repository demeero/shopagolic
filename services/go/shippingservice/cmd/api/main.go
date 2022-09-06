package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"

	"github.com/demeero/shopagolic/services/go/bricks/zaplogger"
	"github.com/demeero/shopagolic/shippingservice/shipping"
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
	zlog = zlog.With(zap.String("service_name", "shipping"))

	cfg := config{}
	if err := envconfig.Process("", &cfg); err != nil {
		zlog.Fatal("failed init config", zap.Error(err))
	}

	zlog.Info("configuration initialized")

	grpcStopFunc := grpcServ(cfg.GRPC, shipping.NewShipper(), zlog)
	zlog.Info("GRPC server established")

	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-interruptChan

	zlog.Info("application stopping")

	var appStoppedWithErr bool

	if appStoppedWithErr {
		zlog.Fatal("failed stop app gracefully")
	}

	grpcStopFunc()

	zlog.Info("application stopped gracefully")
}
