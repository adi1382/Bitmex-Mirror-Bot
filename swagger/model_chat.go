package swagger

// Trollbox Data
type Chat struct {
	Id        NullInt    `json:"id,omitempty"`
	Date      NullTime   `json:"date"`
	User      NullString `json:"user"`
	Message   NullString `json:"message"`
	Html      NullString `json:"html"`
	FromBot   NullBool   `json:"fromBot,omitempty"`
	ChannelID NullInt    `json:"channelID,omitempty"`
}
