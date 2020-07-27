package swagger

type Transaction struct {
	TransactID     NullString `json:"transactID"`
	Account        NullInt    `json:"account,omitempty"`
	Currency       NullString `json:"currency,omitempty"`
	TransactType   NullString `json:"transactType,omitempty"`
	Amount         NullInt    `json:"amount,omitempty"`
	Fee            NullInt    `json:"fee,omitempty"`
	TransactStatus NullString `json:"transactStatus,omitempty"`
	Address        NullString `json:"address,omitempty"`
	Tx             NullString `json:"tx,omitempty"`
	Text           NullString `json:"text,omitempty"`
	TransactTime   NullTime   `json:"transactTime,omitempty"`
	Timestamp      NullTime   `json:"timestamp,omitempty"`
}
