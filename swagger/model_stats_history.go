package swagger

type StatsHistory struct {
	Date       NullTime   `json:"date"`
	RootSymbol NullString `json:"rootSymbol"`
	Currency   NullString `json:"currency,omitempty"`
	Volume     NullInt    `json:"volume,omitempty"`
	Turnover   NullInt    `json:"turnover,omitempty"`
}
