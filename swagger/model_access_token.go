package swagger

type AccessToken struct {
	Id NullString `json:"id"`
	// time to live in seconds (2 weeks by default)
	Ttl     NullFloat64 `json:"ttl,omitempty"`
	Created NullTime    `json:"created,omitempty"`
	UserId  NullFloat64 `json:"userId,omitempty"`
}
