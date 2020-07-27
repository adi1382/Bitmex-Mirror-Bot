package swagger

type Affiliate struct {
	Account         NullInt     `json:"account"`
	Currency        NullString  `json:"currency"`
	PrevPayout      NullInt     `json:"prevPayout,omitempty"`
	PrevTurnover    NullInt     `json:"prevTurnover,omitempty"`
	PrevComm        NullInt     `json:"prevComm,omitempty"`
	PrevTimestamp   NullTime    `json:"prevTimestamp,omitempty"`
	ExecTurnover    NullInt     `json:"execTurnover,omitempty"`
	ExecComm        NullInt     `json:"execComm,omitempty"`
	TotalReferrals  NullInt     `json:"totalReferrals,omitempty"`
	TotalTurnover   NullInt     `json:"totalTurnover,omitempty"`
	TotalComm       NullInt     `json:"totalComm,omitempty"`
	PayoutPcnt      NullFloat64 `json:"payoutPcnt,omitempty"`
	PendingPayout   NullInt     `json:"pendingPayout,omitempty"`
	Timestamp       NullTime    `json:"timestamp,omitempty"`
	ReferrerAccount NullInt     `json:"referrerAccount,omitempty"`
}
