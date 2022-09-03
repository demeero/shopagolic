package repository

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"

	"github.com/demeero/shopagolic/productcatalog/catalog"
	"github.com/demeero/shopagolic/services/go/bricks/zaplogger"
)

var productSortKeyMapping = map[catalog.SortKey]string{
	catalog.SortKeyName:      "name",
	catalog.SortKeyCreatedAt: "created_at",
}

type Loader struct {
	coll *mongo.Collection
}

func NewLoader(coll *mongo.Collection) *Loader {
	return &Loader{coll: coll}
}

func (r *Loader) Load(ctx context.Context, p catalog.Pagination, sortKey catalog.SortKey, asc bool) ([]catalog.Product, string, error) {
	pt, err := decodePageToken(p.PageToken)
	if err != nil {
		return nil, "", errInvalidPageToken
	}
	mSortKey, sortType := r.resolveListSort(sortKey, asc)
	opts := options.Find().
		SetSort(r.buildSortListQuery(mSortKey, sortType)).
		SetLimit(int64(p.PageSize))
	cur, err := r.coll.Find(ctx, r.addPagination(bson.M{}, mSortKey, sortType, pt), opts)
	if err != nil {
		return nil, "", err
	}
	return r.processListCursor(ctx, int(p.PageSize), mSortKey, cur)
}

func (r *Loader) LoadByID(ctx context.Context, id string) (catalog.Product, error) {
	docID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return catalog.Product{}, errInvalidID
	}
	var doc productDoc
	err = r.coll.FindOne(ctx, bson.M{"_id": docID}).Decode(&doc)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return catalog.Product{}, catalog.ErrNotFound
	}
	if err != nil {
		return catalog.Product{}, err
	}
	return doc.Domain(), nil
}

func (r *Loader) Count(ctx context.Context) (uint, error) {
	n, err := r.coll.CountDocuments(ctx, bson.D{})
	if err != nil {
		return 0, fmt.Errorf("failed count documents: %w", err)
	}
	return uint(n), nil
}

func (r *Loader) buildSortListQuery(sortKey string, sortType int) bson.D {
	var sort bson.D
	if sortKey != "" {
		sort = append(sort, bson.E{Key: sortKey, Value: sortType})
	}
	return append(sort, bson.E{Key: "_id", Value: sortType})
}

func (r *Loader) resolveListSort(sortKey catalog.SortKey, asc bool) (mSortKey string, order int) {
	sortType := -1
	if asc {
		sortType = 1
	}
	return productSortKeyMapping[sortKey], sortType
}

func (r *Loader) addPagination(query bson.M, sortKey string, sortType int, pt *pageToken) bson.M {
	if pt == nil {
		return query
	}
	cond := "$lt"
	if sortType == 1 {
		cond = "$gt"
	}
	if sortKey == "" {
		query["_id"] = bson.D{{cond, pt.ID}}
		return query
	}
	query["$or"] = bson.A{
		bson.D{{sortKey, bson.D{{cond, pt.SortVal}}}},
		bson.D{
			{sortKey, pt.SortVal},
			{"_id", bson.D{{cond, pt.ID}}},
		},
	}
	return query
}

func (r *Loader) processListCursor(ctx context.Context, size int, sortKey string, cur *mongo.Cursor) ([]catalog.Product, string, error) {
	defer func() {
		if err := cur.Close(ctx); err != nil {
			zaplogger.FromCtx(ctx).Error("failed close mongo cursor", zap.Error(err))
		}
	}()
	var (
		products []catalog.Product
		last     bson.Raw
	)
	for cur.Next(ctx) {
		var pDoc productDoc
		if err := cur.Decode(&pDoc); err != nil {
			id, _ := cur.Current.Lookup("_id").ObjectIDOK()
			return nil, "", fmt.Errorf("failed decode document %s: %w", id, err)
		}
		products = append(products, pDoc.Domain())
		last = cur.Current
	}
	if len(products) < size {
		return products, "", nil
	}
	nextPageToken, err := encodePageToken(last, sortKey)
	if err != nil {
		return nil, "", err
	}
	return products, nextPageToken, nil
}
