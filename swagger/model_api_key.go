package swagger

// Persistent API keys for Developers
type ApiKey struct {
	Id          NullString   `json:"id"`
	Secret      NullString   `json:"secret"`
	Name        NullString   `json:"name"`
	Nonce       NullInt      `json:"nonce"`
	Cidr        NullString   `json:"cidr,omitempty"`
	Permissions []NullString `json:"permissions,omitempty"`
	Enabled     NullBool     `json:"enabled,omitempty"`
	UserId      NullInt      `json:"userId"`
	Created     NullTime     `json:"created,omitempty"`
}
