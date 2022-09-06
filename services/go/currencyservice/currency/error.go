package currency

import "errors"

var (
	ErrInvalidData  = errors.New("invalid data")
	ErrNotFound     = errors.New("not found")
	ErrConflictData = errors.New("conflict")
)
