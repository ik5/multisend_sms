package multisendsms

//ResponseErrors is a map for holding error messages arrives from the Response
type ResponseErrors map[string]string

// Response holds the structure for response messages
type Response struct {
	Success bool           `json:"success"`
	Message string         `json:"message,omitempty"`
	Error   ResponseErrors `json:"error,omitempty"`
}

// FromXMLResponse implements the interface of Response
func (r *Response) FromXMLResponse(status []byte) error {
	return nil
}

// FromJSONResponse implements the interface of Response
func (r *Response) FromJSONResponse(status []byte) error {
	return nil
}

// ToError implements the interface of Response
func (r *Response) ToError() error {
	return nil
}

// IsOK implements the interface of Response
func (r *Response) IsOK() bool {
	return false
}
