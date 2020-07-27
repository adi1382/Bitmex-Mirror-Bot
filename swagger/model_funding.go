package swagger

// Swap Funding History
type Funding struct {
	Timestamp        NullTime    `json:"timestamp"`
	Symbol           NullString  `json:"symbol"`
	FundingInterval  NullTime    `json:"fundingInterval,omitempty"`
	FundingRate      NullFloat64 `json:"fundingRate,omitempty"`
	FundingRateDaily NullFloat64 `json:"fundingRateDaily,omitempty"`
}
