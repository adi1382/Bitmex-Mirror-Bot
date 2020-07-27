package swagger

// Exchange Statistics
type Stats struct {
	RootSymbol       NullString `json:"rootSymbol"`
	Currency         NullString `json:"currency,omitempty"`
	Volume24h        NullInt    `json:"volume24h,omitempty"`
	Turnover24h      NullInt    `json:"turnover24h,omitempty"`
	OpenNullInterest NullInt    `json:"openNullInterest,omitempty"`
	OpenValue        NullInt    `json:"openValue,omitempty"`
}
