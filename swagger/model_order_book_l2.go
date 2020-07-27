package swagger

type OrderBookL2 struct {
	Symbol NullString  `json:"symbol"`
	Id     NullInt     `json:"id"`
	Side   NullString  `json:"side"`
	Size   NullInt     `json:"size,omitempty"`
	Price  NullFloat64 `json:"price,omitempty"`
}
