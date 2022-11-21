package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/demeero/shopagolic/recommendationservice/recommendation"
	"github.com/demeero/shopagolic/recommendationservice/stub"
	catalogpb "github.com/demeero/shopagolic/services/proto/gen/go/shopagolic/productcatalog/v1beta1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

//go:embed stub-products.json
var defaultCatalogStub string

func createProductCatalogLoader(cfg productCatalogLoaderClientCfg) (recommendation.ProductCatalogLoader, func() error, error) {
	if cfg.UseStub {
		catalogStub, err := createProductCatalogLoaderStub(cfg)
		return catalogStub, func() error { return nil }, err
	}
	conn, err := grpc.Dial(cfg.ServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, nil, fmt.Errorf("failed dial product catalog service %s: %w", cfg.ServiceAddr, err)
	}
	return catalogpb.NewProductCatalogServiceClient(conn), func() error { return conn.Close() }, nil
}

func createProductCatalogLoaderStub(cfg productCatalogLoaderClientCfg) (recommendation.ProductCatalogLoader, error) {
	r, err := createProductCatalogStubReader(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed create catalog reader: %w", err)
	}
	var products []*catalogpb.Product
	if err := json.NewDecoder(r).Decode(&products); err != nil {
		return nil, fmt.Errorf("failed decode file with catalog stub data %s: %w", cfg.StubDataPath, err)
	}
	return stub.ProductCatalogLoader{Products: products}, nil
}

func createProductCatalogStubReader(cfg productCatalogLoaderClientCfg) (io.Reader, error) {
	if cfg.StubDataPath == "" {
		return strings.NewReader(defaultCatalogStub), nil
	}
	f, err := os.Open(cfg.StubDataPath)
	if err != nil {
		return nil, fmt.Errorf("failed open file with catalog stub data %s: %w", cfg.StubDataPath, err)
	}
	return f, nil
}

func createProductCatalogHealthClient(cfg productCatalogLoaderClientCfg) (catalogpb.HealthServiceClient, func() error, error) {
	conn, err := grpc.Dial(cfg.HealthAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, nil, fmt.Errorf("failed dial product catalog health service %s: %w", cfg.HealthAddr, err)
	}
	return catalogpb.NewHealthServiceClient(conn), func() error { return conn.Close() }, nil
}
