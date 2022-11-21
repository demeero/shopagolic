package cart

type Cart struct {
	Items  []Item
	UserID string
}

type Item struct {
	Quantity  uint16
	ProductID string
}
