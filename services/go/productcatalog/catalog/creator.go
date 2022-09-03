package catalog

import (
	"context"
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type CreatorRepository interface {
	Create(ctx context.Context, params CreateParams) (string, error)
}

type Creator struct {
	repo CreatorRepository
}

func NewCreator(repo CreatorRepository) *Creator {
	return &Creator{repo: repo}
}

func (c *Creator) Create(ctx context.Context, params CreateParams) (string, error) {
	err := validation.Errors{
		"picture":     validation.Validate(params.Picture, is.URL),
		"name":        validation.Validate(params.Name, validation.Required, validation.Length(3, 256)),
		"description": validation.Validate(params.Description, validation.Length(3, 4096)),
	}.Filter()
	if err != nil {
		return "", fmt.Errorf("%w: %s", ErrInvalidData, err)
	}
	id, err := c.repo.Create(ctx, params)
	if err != nil {
		return "", fmt.Errorf("failed create product: %w", err)
	}
	return id, nil
}
