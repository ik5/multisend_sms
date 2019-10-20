package multisendsms

//ResponseErrors is a map for holding error messages arrives from the Response
type ResponseErrors map[string]string

// Response holds the structure for response messages
type Response struct {
	Success bool           `json:"success"`
	Message string         `json:"message,omitempty"`
	Error   ResponseErrors `json:"error,omitempty"`
}
