package repository

import (
	"fmt"

	"github.com/demeero/shopagolic/productcatalog/catalog"
)

var (
	errInvalidID        = fmt.Errorf("%w: invalid format of ID", catalog.ErrInvalidData)
	errInvalidPageToken = fmt.Errorf("%w: page token", catalog.ErrInvalidData)
)
