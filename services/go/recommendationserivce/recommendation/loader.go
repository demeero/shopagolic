package recommendation

import (
	"context"
	"fmt"
	"strings"

	catalogpb "github.com/demeero/shopagolic/services/proto/gen/go/shopagolic/productcatalog/v1beta1"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const defaultRecLimit uint8 = 5

type ProductCatalogLoader interface {
	GetProduct(ctx context.Context, in *catalogpb.GetProductRequest, opts ...grpc.CallOption) (*catalogpb.GetProductResponse, error)
	SearchProducts(ctx context.Context, in *catalogpb.SearchProductsRequest, opts ...grpc.CallOption) (*catalogpb.SearchProductsResponse, error)
}

type Loader struct {
	catalogLoader ProductCatalogLoader
}

func NewLoader(catalogLoader ProductCatalogLoader) *Loader {
	return &Loader{catalogLoader: catalogLoader}
}

func (l *Loader) Load(ctx context.Context, productID string, limit uint8) (Recommendation, error) {
	if err := validation.Validate(productID, validation.Required); err != nil {
		return Recommendation{}, fmt.Errorf("%w: %s", ErrInvalidData, err)
	}
	if limit == 0 {
		limit = defaultRecLimit
	}
	foundProduct, err := l.catalogLoader.GetProduct(ctx, &catalogpb.GetProductRequest{Id: productID})
	if err := handleProductCatalogLoaderErr(err); err != nil {
		return Recommendation{}, fmt.Errorf("failed get product %s from catalog: %w", productID, err)
	}
	recs, err := l.makeRecommendations(ctx, foundProduct.GetProduct(), int(limit))
	if err != nil {
		return Recommendation{}, fmt.Errorf("failed make recommendations for product %s: %w", productID, err)
	}
	return Recommendation{ProductIDs: recs}, nil
}

func (l *Loader) makeRecommendations(ctx context.Context, product *catalogpb.Product, limit int) ([]string, error) {
	categoriesQuery := strings.Join(product.GetCategories(), " ")
	foundProducts, err := l.catalogLoader.SearchProducts(ctx, &catalogpb.SearchProductsRequest{Query: categoriesQuery})
	if err := handleProductCatalogLoaderErr(err); err != nil {
		return nil, fmt.Errorf("failed search products in catalog: %w", err)
	}
	recommendations := make([]string, 0, limit)
	for i := 0; i < len(foundProducts.GetProducts()) && i < limit; i++ {
		if foundProductID := foundProducts.GetProducts()[i].GetId(); foundProductID != product.GetId() {
			recommendations = append(recommendations, foundProductID)
		}
	}
	return recommendations, nil
}

func handleProductCatalogLoaderErr(err error) error {
	if err == nil {
		return nil
	}
	s, ok := status.FromError(err)
	if !ok {
		return err
	}
	if s.Code() == codes.InvalidArgument {
		return fmt.Errorf("%w: %s", ErrInvalidData, s.Message())
	}
	if s.Code() == codes.NotFound {
		return fmt.Errorf("%w: %s", ErrNotFound, s.Message())
	}
	return err
}
