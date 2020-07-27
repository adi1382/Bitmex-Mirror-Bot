package swagger

type ConnectedUsers struct {
	Users NullInt `json:"users,omitempty"`
	Bots  NullInt `json:"bots,omitempty"`
}
