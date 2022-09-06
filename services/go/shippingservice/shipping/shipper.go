package shipping

import (
	"context"
	"errors"
	"fmt"
	"math"
	"math/rand"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

var ErrInvalidData = errors.New("invalid data")

// Shipper is just a mock that imitates real work.
type Shipper struct{}

func NewShipper() *Shipper {
	rand.Seed(time.Now().UnixNano())
	return &Shipper{}
}

// ShipOrder mocks that the requested items will be shipped.
// It supplies a tracking ID for notional lookup of shipment delivery status.
func (s *Shipper) ShipOrder(_ context.Context, address Address, items []Item) (string, error) {
	if err := validateAddress(address); err != nil {
		return "", fmt.Errorf("%w: invalid address: %s", ErrInvalidData, err)
	}
	if err := validateItems(items); err != nil {
		return "", fmt.Errorf("%w: invalid address: %s", ErrInvalidData, err)
	}
	return createTrackingID(fmt.Sprintf("%s, %s, %s", address.Street, address.City, address.State)), nil
}

// Quote mocks a shipping quote (cost).
func (s *Shipper) Quote(_ context.Context, address Address, items []Item) (Money, error) {
	if err := validateAddress(address); err != nil {
		return Money{}, fmt.Errorf("%w: invalid address: %s", ErrInvalidData, err)
	}
	if err := validateItems(items); err != nil {
		return Money{}, fmt.Errorf("%w: invalid address: %s", ErrInvalidData, err)
	}
	quote := createQuoteFromCount(0)
	return Money{
		CurrencyCode: "USD",
		Units:        int64(quote.Units),
		Nanos:        int32(quote.Cents * 10000000),
	}, nil
}

func validateAddress(addr Address) error {
	return validation.Errors{
		"street":   validation.Validate(addr.Street, validation.Required, validation.Length(2, 1000)),
		"city":     validation.Validate(addr.City, validation.Required, validation.Length(2, 300)),
		"state":    validation.Validate(addr.State, validation.Required, validation.Length(2, 300)),
		"country":  validation.Validate(addr.Country, validation.Required, validation.Length(2, 300)),
		"zip_code": validation.Validate(addr.ZIPCode, validation.Required),
	}.Filter()
}

func validateItems(items []Item) error {
	for _, i := range items {
		if err := validateItem(i); err != nil {
			return err
		}
	}
	return nil
}

func validateItem(item Item) error {
	return validation.Errors{
		"product_id": validation.Validate(item.ProductID, validation.Required),
		"quantity":   validation.Validate(int(item.Quantity), validation.Required, validation.Min(1)),
	}.Filter()
}

func createTrackingID(salt string) string {
	return fmt.Sprintf("%c%c-%d%s-%d%s",
		getRandomLetterCode(),
		getRandomLetterCode(),
		len(salt),
		getRandomNumber(3),
		len(salt)/2,
		getRandomNumber(7),
	)
}

func getRandomLetterCode() uint32 {
	return 65 + uint32(rand.Intn(25))
}

func getRandomNumber(digits int) string {
	str := ""
	for i := 0; i < digits; i++ {
		str = fmt.Sprintf("%s%d", str, rand.Intn(10))
	}
	return str
}

func createQuoteFromCount(count int) Quote {
	return createQuoteFromFloat(8.99)
}

func createQuoteFromFloat(value float64) Quote {
	units, fraction := math.Modf(value)
	return Quote{
		uint32(units),
		uint32(math.Trunc(fraction * 100)),
	}
}
