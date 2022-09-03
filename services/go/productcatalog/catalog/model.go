package catalog

import "time"

type Money struct {
	CurrencyCode string
	Units        int64
	Nanos        int32
}

type Product struct {
	ID          string
	Name        string
	Description string
	Picture     string
	Price       Money
	Categories  []string
	CreatedAt   time.Time
}

type ProductList struct {
	Page          []Product
	Total         uint
	NextTokenPage string
}

type CreateParams struct {
	Name        string
	Description string
	Picture     string
	Price       Money
	Categories  []string
}
