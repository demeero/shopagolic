package cart

import (
	"context"
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type DeleterRepository interface {
	DeleteAll(ctx context.Context, userID string) error
}

type Deleter struct {
	repo DeleterRepository
}

func NewDeleter(repo DeleterRepository) *Deleter {
	return &Deleter{repo: repo}
}

func (d *Deleter) DeleteAll(ctx context.Context, userID string) error {
	if err := validation.Validate(userID, validation.Required); err != nil {
		return fmt.Errorf("%w: %s", ErrInvalidData, err)
	}
	return d.repo.DeleteAll(ctx, userID)
}
