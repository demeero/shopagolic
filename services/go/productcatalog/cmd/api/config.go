package main

type config struct {
	GRPC  grpcCfg
	Mongo mongoCfg
}

type grpcCfg struct {
	Port       int  `default:"8080"`
	LogPayload bool `default:"true" split_words:"true"`
}

type mongoCfg struct {
	DBName string `default:"shopagolic" split_words:"true"`
	URI    string `default:"mongodb://localhost:27017"`
}
