package datagen

import (
	"context"
	"fmt"

	"github.com/brianvoe/gofakeit/v6"
	"go.uber.org/zap"

	"github.com/demeero/shopagolic/productcatalog/catalog"
)

var predefinedCategories = []string{
	"accessories",
	"clothing",
	"footwear",
	"hair",
	"beauty",
	"decor",
	"home",
	"kitchen",
}

type Generator struct {
	creator *catalog.Creator
	loader  *catalog.Loader
}

func New(creator *catalog.Creator, loader *catalog.Loader) *Generator {
	return &Generator{creator: creator, loader: loader}
}

func (g *Generator) Products(n int) ([]catalog.Product, error) {
	result := make([]catalog.Product, 0, n)
	for i := 0; i < n; i++ {
		productName := gofakeit.Word()
		if len([]rune(productName)) < 3 {
			productName += gofakeit.Word()
		}
		gofakeit.ShuffleStrings(predefinedCategories)
		id, err := g.creator.Create(context.Background(), catalog.CreateParams{
			Name:        productName,
			Description: gofakeit.HipsterSentence(10),
			Picture:     gofakeit.URL(),
			Price: catalog.Money{
				CurrencyCode: gofakeit.CurrencyShort(),
				Units:        int64(gofakeit.Number(0, 1000)),
				Nanos:        int32(gofakeit.Number(0, 1000)),
			},
			Categories: predefinedCategories[:gofakeit.Number(1, 4)],
		})
		if err != nil {
			return nil, fmt.Errorf("failed create product %s: %w", productName, err)
		}
		p, err := g.loader.LoadByID(context.Background(), id)
		if err != nil {
			return nil, fmt.Errorf("failed load product %s : %w", id, err)
		}
		zap.L().Info("product generated", zap.String("id", p.ID), zap.String("name", p.Name))
		result = append(result, p)
	}
	return result, nil
}
