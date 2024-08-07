// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package sdkerrors

import (
	"encoding/json"
	"fmt"
)

// InviteExpiredCode - A short code indicating the error code returned.
type InviteExpiredCode string

const (
	InviteExpiredCodeInviteExpired InviteExpiredCode = "invite_expired"
)

func (e InviteExpiredCode) ToPointer() *InviteExpiredCode {
	return &e
}
func (e *InviteExpiredCode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "invite_expired":
		*e = InviteExpiredCode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InviteExpiredCode: %v", v)
	}
}

type InviteExpiredError struct {
	// A short code indicating the error code returned.
	Code InviteExpiredCode `json:"code"`
	// A human readable explanation of what went wrong.
	Message string `json:"message"`
	// A link to our documentation with more details about this error code
	DocURL *string `json:"doc_url,omitempty"`
}

func (o *InviteExpiredError) GetCode() InviteExpiredCode {
	if o == nil {
		return InviteExpiredCode("")
	}
	return o.Code
}

func (o *InviteExpiredError) GetMessage() string {
	if o == nil {
		return ""
	}
	return o.Message
}

func (o *InviteExpiredError) GetDocURL() *string {
	if o == nil {
		return nil
	}
	return o.DocURL
}

// InviteExpired - This response is sent when the requested content has been permanently deleted from server, with no forwarding address.
type InviteExpired struct {
	Error_ InviteExpiredError `json:"error"`
}

var _ error = &InviteExpired{}

func (e *InviteExpired) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
