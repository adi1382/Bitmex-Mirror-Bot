package swagger

type ErrorError struct {
	Message NullString `json:"message,omitempty"`
	Name    NullString `json:"name,omitempty"`
}
