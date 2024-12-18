// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

type ResponseBody struct {
	// The unique identifier of the customer in Dub.
	ID string `json:"id"`
	// Unique identifier for the customer in the client's app.
	ExternalID string `json:"externalId"`
	// Name of the customer.
	Name string `json:"name"`
	// Email of the customer.
	Email *string `json:"email,omitempty"`
	// Avatar URL of the customer.
	Avatar *string `json:"avatar,omitempty"`
	// The date the customer was created.
	CreatedAt string `json:"createdAt"`
}

func (o *ResponseBody) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *ResponseBody) GetExternalID() string {
	if o == nil {
		return ""
	}
	return o.ExternalID
}

func (o *ResponseBody) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *ResponseBody) GetEmail() *string {
	if o == nil {
		return nil
	}
	return o.Email
}

func (o *ResponseBody) GetAvatar() *string {
	if o == nil {
		return nil
	}
	return o.Avatar
}

func (o *ResponseBody) GetCreatedAt() string {
	if o == nil {
		return ""
	}
	return o.CreatedAt
}
