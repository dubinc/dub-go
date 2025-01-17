// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

// Code - The error code.
type Code string

const (
	CodeBadRequest          Code = "bad_request"
	CodeNotFound            Code = "not_found"
	CodeInternalServerError Code = "internal_server_error"
	CodeUnauthorized        Code = "unauthorized"
	CodeForbidden           Code = "forbidden"
	CodeRateLimitExceeded   Code = "rate_limit_exceeded"
	CodeInviteExpired       Code = "invite_expired"
	CodeInvitePending       Code = "invite_pending"
	CodeExceededLimit       Code = "exceeded_limit"
	CodeConflict            Code = "conflict"
	CodeUnprocessableEntity Code = "unprocessable_entity"
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
		fallthrough
	case "not_found":
		fallthrough
	case "internal_server_error":
		fallthrough
	case "unauthorized":
		fallthrough
	case "forbidden":
		fallthrough
	case "rate_limit_exceeded":
		fallthrough
	case "invite_expired":
		fallthrough
	case "invite_pending":
		fallthrough
	case "exceeded_limit":
		fallthrough
	case "conflict":
		fallthrough
	case "unprocessable_entity":
		*e = Code(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Code: %v", v)
	}
}

type LinkErrorSchema struct {
	// The link that caused the error.
	Link any `json:"link,omitempty"`
	// The error message.
	Error string `json:"error"`
	// The error code.
	Code Code `json:"code"`
}

func (o *LinkErrorSchema) GetLink() any {
	if o == nil {
		return nil
	}
	return o.Link
}

func (o *LinkErrorSchema) GetError() string {
	if o == nil {
		return ""
	}
	return o.Error
}

func (o *LinkErrorSchema) GetCode() Code {
	if o == nil {
		return Code("")
	}
	return o.Code
}
