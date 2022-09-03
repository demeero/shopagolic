package rpc

import (
	"context"
	"errors"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/demeero/shopagolic/productcatalog/catalog"
	"github.com/demeero/shopagolic/services/go/bricks/zaplogger"
	moneypb "github.com/demeero/shopagolic/services/proto/gen/go/shopagolic/money/v1"
	catalogpb "github.com/demeero/shopagolic/services/proto/gen/go/shopagolic/productcatalog/v1beta1"
)

type ProductComponents struct {
	CatalogLoader   *catalog.Loader
	CatalogSearcher *catalog.Searcher
	CatalogCreator  *catalog.Creator
}

var sortKeyMapping = map[catalogpb.SortKey]catalog.SortKey{
	catalogpb.SortKey_SORT_KEY_CREATED_AT:  catalog.SortKeyCreatedAt,
	catalogpb.SortKey_SORT_KEY_NAME:        catalog.SortKeyName,
	catalogpb.SortKey_SORT_KEY_UNSPECIFIED: catalog.SortKeyUnknown,
}

type Product struct {
	catalogpb.UnimplementedProductCatalogServiceServer
	loader   *catalog.Loader
	searcher *catalog.Searcher
	creator  *catalog.Creator
}

func NewProduct(components ProductComponents) *Product {
	return &Product{
		loader:   components.CatalogLoader,
		searcher: components.CatalogSearcher,
		creator:  components.CatalogCreator,
	}
}

func (c *Product) CreateProduct(ctx context.Context, req *catalogpb.CreateProductRequest) (*catalogpb.CreateProductResponse, error) {
	id, err := c.creator.Create(ctx, catalog.CreateParams{
		Name:        req.GetName(),
		Description: req.GetDescription(),
		Picture:     req.GetPicture(),
		Price:       convertProtoMoney(req.GetPrice()),
		Categories:  req.GetCategories(),
	})
	if err := errHandler(ctx, err); err != nil {
		return nil, err
	}
	return &catalogpb.CreateProductResponse{Id: id}, nil
}

func (c *Product) ListProducts(ctx context.Context, req *catalogpb.ListProductsRequest) (*catalogpb.ListProductsResponse, error) {
	pagination := catalog.Pagination{
		PageToken: req.GetPageToken(),
		PageSize:  uint16(req.GetPageSize()),
	}
	pList, err := c.loader.Load(ctx, pagination, sortKeyMapping[req.GetSortKey()], req.GetAsc())
	if err := errHandler(ctx, err); err != nil {
		return nil, err
	}
	return convertToListProductsResponse(pList), nil
}

func (c *Product) GetProduct(ctx context.Context, req *catalogpb.GetProductRequest) (*catalogpb.GetProductResponse, error) {
	p, err := c.loader.LoadByID(ctx, req.GetId())
	if err := errHandler(ctx, err); err != nil {
		return nil, err
	}
	return &catalogpb.GetProductResponse{Product: convertProduct(p)}, nil
}

func (c *Product) SearchProducts(ctx context.Context, req *catalogpb.SearchProductsRequest) (*catalogpb.SearchProductsResponse, error) {
	products, err := c.searcher.Search(ctx, req.GetQuery())
	if err := errHandler(ctx, err); err != nil {
		return nil, err
	}
	return &catalogpb.SearchProductsResponse{Products: convertProducts(products)}, nil
}

func convertToListProductsResponse(in catalog.ProductList) *catalogpb.ListProductsResponse {
	products := make([]*catalogpb.Product, 0, len(in.Page))
	for _, item := range in.Page {
		products = append(products, convertProduct(item))
	}
	return &catalogpb.ListProductsResponse{
		Products:      products,
		NextPageToken: in.NextTokenPage,
	}
}

func convertProduct(in catalog.Product) *catalogpb.Product {
	return &catalogpb.Product{
		Id:          in.ID,
		Name:        in.Name,
		Description: in.Description,
		Picture:     in.Picture,
		Price: &moneypb.Money{
			CurrencyCode: in.Price.CurrencyCode,
			Units:        in.Price.Units,
			Nanos:        in.Price.Nanos,
		},
		Categories: in.Categories,
		CreatedAt:  timestamppb.New(in.CreatedAt),
	}
}

func convertProducts(in []catalog.Product) []*catalogpb.Product {
	res := make([]*catalogpb.Product, 0, len(in))
	for _, p := range in {
		res = append(res, convertProduct(p))
	}
	return res
}

func convertProtoMoney(in *moneypb.Money) catalog.Money {
	return catalog.Money{
		CurrencyCode: in.GetCurrencyCode(),
		Units:        in.GetUnits(),
		Nanos:        in.GetNanos(),
	}
}

func errHandler(ctx context.Context, err error) error {
	if err == nil {
		return nil
	}
	if errors.Is(err, catalog.ErrNotFound) {
		return status.Error(codes.NotFound, err.Error())
	}
	if errors.Is(err, catalog.ErrInvalidData) {
		return status.Error(codes.InvalidArgument, err.Error())
	}
	zaplogger.FromCtx(ctx).Error("failed handle request", zap.Error(err))
	return err
}
