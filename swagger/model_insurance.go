package swagger

// Insurance Fund Data
type Insurance struct {
	Currency      NullString `json:"currency"`
	Timestamp     NullTime   `json:"timestamp"`
	WalletBalance NullInt    `json:"walletBalance,omitempty"`
}
