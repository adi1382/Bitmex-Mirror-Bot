package swagger

// Account Notifications
type Notification struct {
	Id                NullInt    `json:"id,omitempty"`
	Date              NullTime   `json:"date"`
	Title             NullString `json:"title"`
	Body              NullString `json:"body"`
	Ttl               NullInt    `json:"ttl"`
	Type_             NullString `json:"type,omitempty"`
	Closable          NullBool   `json:"closable,omitempty"`
	Persist           NullBool   `json:"persist,omitempty"`
	WaitForVisibility NullBool   `json:"waitForVisibility,omitempty"`
	Sound             NullString `json:"sound,omitempty"`
}
