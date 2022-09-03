package catalog

import (
	"context"
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type SearchRepository interface {
	Search(ctx context.Context, query string) ([]Product, error)
}

type Searcher struct {
	repo SearchRepository
}

func NewSearcher(repo SearchRepository) *Searcher {
	return &Searcher{repo: repo}
}

func (s *Searcher) Search(ctx context.Context, query string) ([]Product, error) {
	if err := validation.Validate(query, validation.Required, validation.Length(3, 128)); err != nil {
		return nil, fmt.Errorf("%w: %s", ErrInvalidData, err)
	}
	records, err := s.repo.Search(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed search products: %w", err)
	}
	return records, nil
}
