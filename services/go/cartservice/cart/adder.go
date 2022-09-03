package cart

import (
	"context"
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type AdderRepository interface {
	AddItem(ctx context.Context, userID string, item Item) error
}

type Adder struct {
	repo AdderRepository
}

func NewAdder(repo AdderRepository) *Adder {
	return &Adder{repo: repo}
}

func (a *Adder) AddItem(ctx context.Context, userID string, item Item) error {
	err := validation.Errors{
		"user_id":    validation.Validate(userID, validation.Required),
		"product_id": validation.Validate(item.ProductID, validation.Required),
		"quantity":   validation.Validate(int64(item.Quantity), validation.Required, validation.Min(1)),
	}.Filter()
	if err != nil {
		return fmt.Errorf("%w: %s", ErrInvalidData, err)
	}
	return a.repo.AddItem(ctx, userID, item)
}
