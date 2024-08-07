// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package sdkerrors

import (
	"encoding/json"
	"fmt"
)

// Code - A short code indicating the error code returned.
type Code string

const (
	CodeBadRequest Code = "bad_request"
)

func (e Code) ToPointer() *Code {
	return &e
}
func (e *Code) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "bad_request":
		*e = Code(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Code: %v", v)
	}
}

type Error struct {
	// A short code indicating the error code returned.
	Code Code `json:"code"`
	// A human readable explanation of what went wrong.
	Message string `json:"message"`
	// A link to our documentation with more details about this error code
	DocURL *string `json:"doc_url,omitempty"`
}

func (o *Error) GetCode() Code {
	if o == nil {
		return Code("")
	}
	return o.Code
}

func (o *Error) GetMessage() string {
	if o == nil {
		return ""
	}
	return o.Message
}

func (o *Error) GetDocURL() *string {
	if o == nil {
		return nil
	}
	return o.DocURL
}

// BadRequest - The server cannot or will not process the request due to something that is perceived to be a client error (e.g., malformed request syntax, invalid request message framing, or deceptive request routing).
type BadRequest struct {
	Error_ Error `json:"error"`
}

var _ error = &BadRequest{}

func (e *BadRequest) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
