package repository

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/demeero/shopagolic/productcatalog/catalog"
)

type moneyDoc struct {
	CurrencyCode string `bson:"currency_code"`
	Units        int64  `bson:"units"`
	Nanos        int32  `bson:"nanos"`
}

type productDoc struct {
	ID          primitive.ObjectID `bson:"_id"`
	Name        string             `bson:"name"`
	Description string             `bson:"description"`
	Picture     string             `bson:"picture"`
	Price       moneyDoc
	Categories  []string  `bson:"categories"`
	CreatedAt   time.Time `bson:"created_at"`
}

func newProductDoc(in catalog.CreateParams) productDoc {
	return productDoc{
		ID:          primitive.NewObjectID(),
		Name:        in.Name,
		Description: in.Description,
		Picture:     in.Picture,
		Price: moneyDoc{
			CurrencyCode: in.Price.CurrencyCode,
			Units:        in.Price.Units,
			Nanos:        in.Price.Nanos,
		},
		Categories: in.Categories,
		CreatedAt:  time.Now().UTC(),
	}
}

func (c productDoc) Domain() catalog.Product {
	return catalog.Product{
		ID:          c.ID.Hex(),
		Name:        c.Name,
		Description: c.Description,
		Picture:     c.Picture,
		Price: catalog.Money{
			CurrencyCode: c.Price.CurrencyCode,
			Units:        c.Price.Units,
			Nanos:        c.Price.Nanos,
		},
		Categories: c.Categories,
		CreatedAt:  c.CreatedAt,
	}
}
