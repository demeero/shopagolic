package shipping

import "fmt"

type Address struct {
	Street  string
	City    string
	State   string
	Country string
	ZIPCode int32
}

type Item struct {
	ProductID string
	Quantity  uint32
}

type Money struct {
	CurrencyCode string
	Units        int64
	Nanos        int32
}

type Quote struct {
	Units uint32
	Cents uint32
}

func (q Quote) String() string {
	return fmt.Sprintf("$%d.%d", q.Units, q.Cents)
}
