package swagger

// Public Announcements
type Announcement struct {
	Id      NullInt    `json:"id"`
	Link    NullString `json:"link,omitempty"`
	Title   NullString `json:"title,omitempty"`
	Content NullString `json:"content,omitempty"`
	Date    NullTime   `json:"date,omitempty"`
}
