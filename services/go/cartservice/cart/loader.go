package cart

import (
	"context"
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type LoaderRepository interface {
	LoadByUserID(ctx context.Context, userID string) (Cart, error)
}

type Loader struct {
	repo LoaderRepository
}

func NewLoader(repo LoaderRepository) *Loader {
	return &Loader{repo: repo}
}

func (l *Loader) LoadByUserID(ctx context.Context, userID string) (Cart, error) {
	if err := validation.Validate(userID, validation.Required); err != nil {
		return Cart{}, fmt.Errorf("%w: %s", ErrInvalidData, err)
	}
	return l.repo.LoadByUserID(ctx, userID)
}
