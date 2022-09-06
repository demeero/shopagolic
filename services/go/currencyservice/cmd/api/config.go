package main

type config struct {
	Redis                 rdsCfg
	GRPC                  grpcCfg
	InitCurrenciesIfEmpty bool `default:"true" split_words:"true"`
}

type grpcCfg struct {
	Port       int  `default:"8080"`
	LogPayload bool `default:"true" split_words:"true"`
}

type rdsCfg struct {
	Addr              string `default:"localhost:6379"`
	CurrencyKeyPrefix string `default:"currency:"`
}
