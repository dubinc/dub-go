// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/dubinc/dub-go/internal/utils"
)

type TrackLeadRequestBody struct {
	// The ID of the click in th Dub. You can read this value from `dub_id` cookie.
	ClickID string `json:"clickId"`
	// The name of the event to track.
	EventName string `json:"eventName"`
	// This is the unique identifier for the customer in the client's app. This is used to track the customer's journey.
	CustomerID string `json:"customerId"`
	// Name of the customer in the client's app.
	CustomerName *string `default:"null" json:"customerName"`
	// Email of the customer in the client's app.
	CustomerEmail *string `default:"null" json:"customerEmail"`
	// Avatar of the customer in the client's app.
	CustomerAvatar *string `default:"null" json:"customerAvatar"`
	// Additional metadata to be stored with the lead event
	Metadata map[string]any `json:"metadata,omitempty"`
}

func (t TrackLeadRequestBody) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *TrackLeadRequestBody) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *TrackLeadRequestBody) GetClickID() string {
	if o == nil {
		return ""
	}
	return o.ClickID
}

func (o *TrackLeadRequestBody) GetEventName() string {
	if o == nil {
		return ""
	}
	return o.EventName
}

func (o *TrackLeadRequestBody) GetCustomerID() string {
	if o == nil {
		return ""
	}
	return o.CustomerID
}

func (o *TrackLeadRequestBody) GetCustomerName() *string {
	if o == nil {
		return nil
	}
	return o.CustomerName
}

func (o *TrackLeadRequestBody) GetCustomerEmail() *string {
	if o == nil {
		return nil
	}
	return o.CustomerEmail
}

func (o *TrackLeadRequestBody) GetCustomerAvatar() *string {
	if o == nil {
		return nil
	}
	return o.CustomerAvatar
}

func (o *TrackLeadRequestBody) GetMetadata() map[string]any {
	if o == nil {
		return nil
	}
	return o.Metadata
}

type Click struct {
	ID string `json:"id"`
}

func (o *Click) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type Customer struct {
	ID     string  `json:"id"`
	Name   *string `json:"name"`
	Email  *string `json:"email"`
	Avatar *string `json:"avatar"`
}

func (o *Customer) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *Customer) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *Customer) GetEmail() *string {
	if o == nil {
		return nil
	}
	return o.Email
}

func (o *Customer) GetAvatar() *string {
	if o == nil {
		return nil
	}
	return o.Avatar
}

// TrackLeadResponseBody - A lead was tracked.
type TrackLeadResponseBody struct {
	Click    Click    `json:"click"`
	Customer Customer `json:"customer"`
}

func (o *TrackLeadResponseBody) GetClick() Click {
	if o == nil {
		return Click{}
	}
	return o.Click
}

func (o *TrackLeadResponseBody) GetCustomer() Customer {
	if o == nil {
		return Customer{}
	}
	return o.Customer
}
