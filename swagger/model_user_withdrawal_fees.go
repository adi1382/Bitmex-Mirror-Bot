package swagger

type UserWithdrawalFees struct {
	Currency string  `json:"currency,omitempty"`
	Fee      NullInt `json:"fee,omitempty"`
	MinFee   NullInt `json:"minFee,omitempty"`
	MaxFee   NullInt `json:"maxFee,omitempty"`
}
