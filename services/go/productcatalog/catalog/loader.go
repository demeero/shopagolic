package catalog

import (
	"context"
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"go.uber.org/zap"

	"github.com/demeero/shopagolic/services/go/bricks/zaplogger"
)

type Pagination struct {
	PageToken string `json:"page_token"`
	PageSize  uint16 `json:"page_size"`
}

type SortKey uint8

const (
	SortKeyUnknown SortKey = iota
	SortKeyName
	SortKeyCreatedAt
)

type LoaderRepository interface {
	Load(ctx context.Context, p Pagination, sortKey SortKey, asc bool) ([]Product, string, error)
	LoadByID(context.Context, string) (Product, error)
	Count(ctx context.Context) (uint, error)
}

type Loader struct {
	repo LoaderRepository
}

func NewLoader(repo LoaderRepository) *Loader {
	return &Loader{repo: repo}
}

func (l *Loader) LoadByID(ctx context.Context, id string) (Product, error) {
	if err := validation.Validate(id, validation.Required); err != nil {
		return Product{}, fmt.Errorf("%w: %s", ErrInvalidData, err)
	}
	p, err := l.repo.LoadByID(ctx, id)
	if err != nil {
		return Product{}, fmt.Errorf("failed load product by id: %w", err)
	}
	return p, nil
}

func (l *Loader) Load(ctx context.Context, p Pagination, sortKey SortKey, asc bool) (ProductList, error) {
	records, nextTokenPage, err := l.repo.Load(ctx, p, sortKey, asc)
	if err != nil {
		return ProductList{}, fmt.Errorf("failed load products: %w", err)
	}
	total, err := l.repo.Count(ctx)
	if err != nil {
		zaplogger.FromCtx(ctx).Error("failed get total amount of products", zap.Error(err))
	}
	return ProductList{Page: records, Total: total, NextTokenPage: nextTokenPage}, nil
}
