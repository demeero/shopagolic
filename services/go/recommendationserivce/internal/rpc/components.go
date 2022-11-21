package rpc

import (
	"github.com/demeero/shopagolic/recommendationservice/recommendation"
	catalogpb "github.com/demeero/shopagolic/services/proto/gen/go/shopagolic/productcatalog/v1beta1"
)

type Components struct {
	Loader              *recommendation.Loader
	CatalogHealthClient catalogpb.HealthServiceClient
}
