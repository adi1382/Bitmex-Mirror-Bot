package swagger

// Best Bid/Offer Snapshots & Historical Bins
type Quote struct {
	Timestamp NullTime    `json:"timestamp"`
	Symbol    NullString  `json:"symbol"`
	BidSize   NullInt     `json:"bidSize,omitempty"`
	BidPrice  NullFloat64 `json:"bidPrice,omitempty"`
	AskPrice  NullFloat64 `json:"askPrice,omitempty"`
	AskSize   NullInt     `json:"askSize,omitempty"`
}
