package swagger

type UserCommission struct {
	MakerFee      NullFloat64 `json:"makerFee,omitempty"`
	TakerFee      NullFloat64 `json:"takerFee,omitempty"`
	SettlementFee NullFloat64 `json:"settlementFee,omitempty"`
	MaxFee        NullFloat64 `json:"maxFee,omitempty"`
}
