package multisendsms

import "encoding/json"

//ResponseErrors is a map for holding error messages arrives from the Response
type ResponseErrors map[string]string

// Response holds the structure for response messages
type Response struct {
	Success  bool           `json:"success"`
	Message  string         `json:"message,omitempty"`
	SMSCount int            `json:"smsCount,omitempty"`
	Error    ResponseErrors `json:"error,omitempty"`
}

// FromResponse implements the interface of Response
func (r *Response) FromResponse(status []byte) error {
	return json.Unmarshal(status, r)
}

// ToError implements the interface of Response
func (r *Response) ToError() error {
	return nil
}

// IsOK implements the interface of Response
func (r Response) IsOK() bool {
	return r.Success
}
