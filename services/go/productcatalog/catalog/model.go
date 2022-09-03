package catalog

import "time"

type Money struct {
	CurrencyCode string
	Units        int64
	Nanos        int32
}

type Product struct {
	CreatedAt   time.Time
	Name        string
	Description string
	Picture     string
	ID          string
	Categories  []string
	Price       Money
}

type ProductList struct {
	NextTokenPage string
	Page          []Product
	Total         uint
}

type CreateParams struct {
	Name        string
	Description string
	Picture     string
	Categories  []string
	Price       Money
}
