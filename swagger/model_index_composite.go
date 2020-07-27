package swagger

type IndexComposite struct {
	Timestamp   NullTime    `json:"timestamp"`
	Symbol      NullString  `json:"symbol,omitempty"`
	IndexSymbol NullString  `json:"indexSymbol,omitempty"`
	Reference   NullString  `json:"reference,omitempty"`
	LastPrice   NullFloat64 `json:"lastPrice,omitempty"`
	Weight      NullFloat64 `json:"weight,omitempty"`
	Logged      NullTime    `json:"logged,omitempty"`
}
