package pagetoken

import (
	"testing"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/stretchr/testify/assert"
)

type pageTokenKey struct {
	SortVal interface{}
	ID      string
}

func TestPageToken(t *testing.T) {
	key := pageTokenKey{
		SortVal: 3100,
		ID:      "123",
	}
	actual, err := Encode(key)
	assert.NoError(t, err)
	assert.NotEmpty(t, actual)
	assert.NoError(t, validation.Validate(actual, is.Base64))

	actualKey := pageTokenKey{}
	err = Decode(actual, &actualKey)
	assert.NoError(t, err)
	assert.Equal(t, key, actualKey)
}
