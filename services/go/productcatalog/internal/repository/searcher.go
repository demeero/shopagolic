package repository

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"

	"github.com/demeero/shopagolic/productcatalog/catalog"
	"github.com/demeero/shopagolic/services/go/bricks/zaplogger"
)

type Searcher struct {
	coll *mongo.Collection
	size int
}

func NewSearcher(coll *mongo.Collection, size uint8) *Searcher {
	return &Searcher{coll: coll, size: int(size)}
}

func (r *Searcher) Search(ctx context.Context, keywords string) ([]catalog.Product, error) {
	cur, err := r.coll.Aggregate(ctx, r.buildTextSearchQuery(keywords, r.size))
	if err != nil {
		return nil, fmt.Errorf("failed search products by keywords %s: %w", keywords, err)
	}
	return r.processListCursor(ctx, cur)
}

func (r *Searcher) buildTextSearchQuery(keywords string, size int) mongo.Pipeline {
	textSearchFilter := bson.D{{"$match", bson.M{"$text": bson.M{"$search": keywords}}}}
	addScoreField := bson.D{{"$addFields", bson.M{"score": bson.M{"$meta": "textScore"}}}}
	sorting := bson.D{{"$sort", bson.D{{"score", -1}, {"_id", -1}}}}
	pipeline := mongo.Pipeline{
		textSearchFilter,
		addScoreField,
		sorting,
	}
	limit := bson.D{{"$limit", size}}
	return append(pipeline, limit)
}

func (r *Searcher) processListCursor(ctx context.Context, cur *mongo.Cursor) ([]catalog.Product, error) {
	defer func() {
		if err := cur.Close(ctx); err != nil {
			zaplogger.FromCtx(ctx).Error("failed to close mongo cursor", zap.Error(err))
		}
	}()
	var products []catalog.Product
	for cur.Next(ctx) {
		var doc productDoc
		if err := cur.Decode(&doc); err != nil {
			id, _ := cur.Current.Lookup("_id").ObjectIDOK()
			return nil, fmt.Errorf("failed decode document %s: %w", id, err)
		}
		products = append(products, doc.Domain())
	}
	return products, nil
}
