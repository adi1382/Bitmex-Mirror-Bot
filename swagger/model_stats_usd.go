package swagger

type StatsUsd struct {
	RootSymbol   NullString `json:"rootSymbol"`
	Currency     NullString `json:"currency,omitempty"`
	Turnover24h  NullInt    `json:"turnover24h,omitempty"`
	Turnover30d  NullInt    `json:"turnover30d,omitempty"`
	Turnover365d NullInt    `json:"turnover365d,omitempty"`
	Turnover     NullInt    `json:"turnover,omitempty"`
}
