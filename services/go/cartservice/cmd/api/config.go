package main

import "time"

type config struct {
	GRPC    grpcCfg
	Redis   rdsCfg
	CartTTL time.Duration `default:"48h"`
}

type grpcCfg struct {
	Port       int  `default:"8080"`
	LogPayload bool `default:"true" split_words:"true"`
}

type rdsCfg struct {
	Addr          string `default:"localhost:6379"`
	CartKeyPrefix string `default:"cart:"`
}
