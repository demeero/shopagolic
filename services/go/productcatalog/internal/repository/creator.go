package repository

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/demeero/shopagolic/productcatalog/catalog"
)

type Creator struct {
	coll *mongo.Collection
}

func NewCreator(coll *mongo.Collection) *Creator {
	return &Creator{coll: coll}
}

func (r *Creator) Create(ctx context.Context, params catalog.CreateParams) (string, error) {
	res, err := r.coll.InsertOne(ctx, newProductDoc(params))
	if err != nil {
		return "", fmt.Errorf("failed insert product: %w", err)
	}
	if docID, ok := res.InsertedID.(primitive.ObjectID); ok {
		return docID.Hex(), nil
	}
	return "", fmt.Errorf("failed cast %T to primitive.ObjectID", res.InsertedID)
}
