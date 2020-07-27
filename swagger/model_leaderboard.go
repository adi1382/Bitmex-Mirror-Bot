package swagger

// Information on Top Users
type Leaderboard struct {
	Name       NullString  `json:"name"`
	IsRealName NullBool    `json:"isRealName,omitempty"`
	Profit     NullFloat64 `json:"profit,omitempty"`
}
