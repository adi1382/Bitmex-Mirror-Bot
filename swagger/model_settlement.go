package swagger

// Historical Settlement Data
type Settlement struct {
	Timestamp             NullTime    `json:"timestamp"`
	Symbol                NullString  `json:"symbol"`
	SettlementType        NullString  `json:"settlementType,omitempty"`
	SettledPrice          NullFloat64 `json:"settledPrice,omitempty"`
	OptionStrikePrice     NullFloat64 `json:"optionStrikePrice,omitempty"`
	OptionUnderlyingPrice NullFloat64 `json:"optionUnderlyingPrice,omitempty"`
	Bankrupt              NullInt     `json:"bankrupt,omitempty"`
	TaxBase               NullInt     `json:"taxBase,omitempty"`
	TaxRate               NullFloat64 `json:"taxRate,omitempty"`
}
