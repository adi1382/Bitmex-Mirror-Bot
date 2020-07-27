package swagger

type TradeBin struct {
	Timestamp       NullTime    `json:"timestamp"`
	Symbol          NullString  `json:"symbol"`
	Open            NullFloat64 `json:"open,omitempty"`
	High            NullFloat64 `json:"high,omitempty"`
	Low             NullFloat64 `json:"low,omitempty"`
	Close           NullFloat64 `json:"close,omitempty"`
	Trades          NullInt     `json:"trades,omitempty"`
	Volume          NullInt     `json:"volume,omitempty"`
	Vwap            NullFloat64 `json:"vwap,omitempty"`
	LastSize        NullInt     `json:"lastSize,omitempty"`
	Turnover        NullInt     `json:"turnover,omitempty"`
	HomeNotional    NullFloat64 `json:"homeNotional,omitempty"`
	ForeignNotional NullFloat64 `json:"foreignNotional,omitempty"`
}
