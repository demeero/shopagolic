package stub

import (
	"context"
	"fmt"

	catalogpb "github.com/demeero/shopagolic/services/proto/gen/go/shopagolic/productcatalog/v1beta1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ProductCatalogLoader struct {
	Products []*catalogpb.Product
}

func (l ProductCatalogLoader) GetProduct(_ context.Context, in *catalogpb.GetProductRequest, _ ...grpc.CallOption) (*catalogpb.GetProductResponse, error) {
	for _, p := range l.Products {
		if p.GetId() == in.GetId() {
			return &catalogpb.GetProductResponse{Product: p}, nil
		}
	}
	return nil, status.Error(codes.NotFound, fmt.Sprintf("not found: %s", in.GetId()))
}

func (l ProductCatalogLoader) SearchProducts(context.Context, *catalogpb.SearchProductsRequest, ...grpc.CallOption) (*catalogpb.SearchProductsResponse, error) {
	return &catalogpb.SearchProductsResponse{Products: l.Products}, nil
}
