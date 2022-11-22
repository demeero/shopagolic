package main

type config struct {
	GRPC grpcCfg
}

type grpcCfg struct {
	Port       int  `default:"8080"`
	LogPayload bool `default:"true" split_words:"true"`
}
