package repository

import (
	"encoding/gob"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/demeero/shopagolic/services/go/bricks/pagetoken"
)

func init() {
	gob.RegisterName("go.mongodb.org/mongo-driver/bson", bson.RawValue{})
}

type pageToken struct {
	ID      primitive.ObjectID
	SortVal bson.RawValue
}

func encodePageToken(doc bson.Raw, sortKey string) (string, error) {
	if doc == nil {
		return "", nil
	}
	id, ok := doc.Lookup("_id").ObjectIDOK()
	if !ok {
		return "", errors.New("failed to lookup _id on raw document")
	}
	pt := pageToken{
		ID:      id,
		SortVal: doc.Lookup(sortKey),
	}
	return pagetoken.Encode(pt)
}

func decodePageToken(token string) (*pageToken, error) {
	if token == "" {
		return nil, nil
	}
	pt := pageToken{}
	if err := pagetoken.Decode(token, &pt); err != nil {
		return nil, err
	}
	return &pt, nil
}
