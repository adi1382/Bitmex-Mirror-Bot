package swagger

// Active Liquidations
type Liquidation struct {
	OrderID   NullString  `json:"orderID"`
	Symbol    NullString  `json:"symbol,omitempty"`
	Side      NullString  `json:"side,omitempty"`
	Price     NullFloat64 `json:"price,omitempty"`
	LeavesQty NullInt     `json:"leavesQty,omitempty"`
}
