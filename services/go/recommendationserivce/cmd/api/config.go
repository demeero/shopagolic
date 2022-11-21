package main

type config struct {
	ProductCatalogLoaderClient productCatalogLoaderClientCfg `split_words:"true"`
	GRPC                       grpcCfg
}

type grpcCfg struct {
	Port       int  `default:"8080"`
	LogPayload bool `default:"true" split_words:"true"`
}

type productCatalogLoaderClientCfg struct {
	StubDataPath string `split_words:"true"`
	ServiceAddr  string `split_words:"true"`
	HealthAddr   string `split_words:"true"`
	UseStub      bool   `split_words:"true" default:"true"`
}
