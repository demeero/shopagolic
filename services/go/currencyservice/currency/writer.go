package currency

import (
	"context"
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type WriterRepository interface {
	Put(ctx context.Context, currCode string, val float32) error
	Delete(ctx context.Context, currCode string) error
}

type Writer struct {
	repo WriterRepository
}

func NewWriter(repo WriterRepository) *Writer {
	return &Writer{repo: repo}
}

func (w *Writer) Put(ctx context.Context, currCode string, val float32) error {
	if err := validation.Validate(currCode, validation.Required, is.CurrencyCode); err != nil {
		return fmt.Errorf("%w: %s", ErrInvalidData, err)
	}
	return w.repo.Put(ctx, currCode, val)
}

func (w *Writer) Delete(ctx context.Context, currCode string) error {
	if err := validation.Validate(currCode, validation.Required, is.CurrencyCode); err != nil {
		return fmt.Errorf("%w: %s", ErrInvalidData, err)
	}
	return w.repo.Delete(ctx, currCode)
}
