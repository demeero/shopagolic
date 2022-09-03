package cart

type Cart struct {
	UserID string
	Items  []Item
}

type Item struct {
	ProductID string
	Quantity  uint16
}
