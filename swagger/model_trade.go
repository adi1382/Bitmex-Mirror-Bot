package swagger

// Individual & Bucketed Trades
type Trade struct {
	Timestamp       NullTime    `json:"timestamp"`
	Symbol          NullString  `json:"symbol"`
	Side            NullString  `json:"side,omitempty"`
	Size            NullInt     `json:"size,omitempty"`
	Price           NullFloat64 `json:"price,omitempty"`
	TickDirection   NullString  `json:"tickDirection,omitempty"`
	TrdMatchID      NullString  `json:"trdMatchID,omitempty"`
	GrossValue      NullInt     `json:"grossValue,omitempty"`
	HomeNotional    NullFloat64 `json:"homeNotional,omitempty"`
	ForeignNotional NullFloat64 `json:"foreignNotional,omitempty"`
}
