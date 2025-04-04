// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
)

type GetCustomersRequest struct {
	// A case-sensitive filter on the list based on the customer's `email` field. The value must be a string.
	Email *string `queryParam:"style=form,explode=true,name=email"`
	// A case-sensitive filter on the list based on the customer's `externalId` field. The value must be a string.
	ExternalID *string `queryParam:"style=form,explode=true,name=externalId"`
	// Whether to include expanded fields on the customer (`link`, `partner`, `discount`).
	IncludeExpandedFields *bool `queryParam:"style=form,explode=true,name=includeExpandedFields"`
}

func (o *GetCustomersRequest) GetEmail() *string {
	if o == nil {
		return nil
	}
	return o.Email
}

func (o *GetCustomersRequest) GetExternalID() *string {
	if o == nil {
		return nil
	}
	return o.ExternalID
}

func (o *GetCustomersRequest) GetIncludeExpandedFields() *bool {
	if o == nil {
		return nil
	}
	return o.IncludeExpandedFields
}

type Link struct {
	// The unique ID of the short link.
	ID string `json:"id"`
	// The domain of the short link. If not provided, the primary domain for the workspace will be used (or `dub.sh` if the workspace has no domains).
	Domain string `json:"domain"`
	// The short link slug. If not provided, a random 7-character slug will be generated.
	Key string `json:"key"`
	// The full URL of the short link, including the https protocol (e.g. `https://dub.sh/try`).
	ShortLink string `json:"shortLink"`
	// The ID of the program the short link is associated with.
	ProgramID *string `json:"programId"`
}

func (o *Link) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *Link) GetDomain() string {
	if o == nil {
		return ""
	}
	return o.Domain
}

func (o *Link) GetKey() string {
	if o == nil {
		return ""
	}
	return o.Key
}

func (o *Link) GetShortLink() string {
	if o == nil {
		return ""
	}
	return o.ShortLink
}

func (o *Link) GetProgramID() *string {
	if o == nil {
		return nil
	}
	return o.ProgramID
}

type GetCustomersPartner struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Email *string `json:"email"`
	Image *string `json:"image"`
}

func (o *GetCustomersPartner) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *GetCustomersPartner) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *GetCustomersPartner) GetEmail() *string {
	if o == nil {
		return nil
	}
	return o.Email
}

func (o *GetCustomersPartner) GetImage() *string {
	if o == nil {
		return nil
	}
	return o.Image
}

type Type string

const (
	TypePercentage Type = "percentage"
	TypeFlat       Type = "flat"
)

func (e Type) ToPointer() *Type {
	return &e
}
func (e *Type) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "percentage":
		fallthrough
	case "flat":
		*e = Type(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Type: %v", v)
	}
}

type Discount struct {
	ID            string   `json:"id"`
	Amount        float64  `json:"amount"`
	Type          Type     `json:"type"`
	MaxDuration   *float64 `json:"maxDuration"`
	Description   *string  `json:"description,omitempty"`
	CouponID      *string  `json:"couponId"`
	CouponTestID  *string  `json:"couponTestId"`
	PartnersCount *float64 `json:"partnersCount,omitempty"`
}

func (o *Discount) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *Discount) GetAmount() float64 {
	if o == nil {
		return 0.0
	}
	return o.Amount
}

func (o *Discount) GetType() Type {
	if o == nil {
		return Type("")
	}
	return o.Type
}

func (o *Discount) GetMaxDuration() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxDuration
}

func (o *Discount) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *Discount) GetCouponID() *string {
	if o == nil {
		return nil
	}
	return o.CouponID
}

func (o *Discount) GetCouponTestID() *string {
	if o == nil {
		return nil
	}
	return o.CouponTestID
}

func (o *Discount) GetPartnersCount() *float64 {
	if o == nil {
		return nil
	}
	return o.PartnersCount
}

type GetCustomersResponseBody struct {
	// The unique ID of the customer. You may use either the customer's `id` on Dub (obtained via `/customers` endpoint) or their `externalId` (unique ID within your system, prefixed with `ext_`, e.g. `ext_123`).
	ID string `json:"id"`
	// Unique identifier for the customer in the client's app.
	ExternalID string `json:"externalId"`
	// Name of the customer.
	Name string `json:"name"`
	// Email of the customer.
	Email *string `json:"email,omitempty"`
	// Avatar URL of the customer.
	Avatar *string `json:"avatar,omitempty"`
	// Country of the customer.
	Country *string `json:"country,omitempty"`
	// The date the customer was created.
	CreatedAt string               `json:"createdAt"`
	Link      *Link                `json:"link,omitempty"`
	Partner   *GetCustomersPartner `json:"partner,omitempty"`
	Discount  *Discount            `json:"discount,omitempty"`
}

func (o *GetCustomersResponseBody) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *GetCustomersResponseBody) GetExternalID() string {
	if o == nil {
		return ""
	}
	return o.ExternalID
}

func (o *GetCustomersResponseBody) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *GetCustomersResponseBody) GetEmail() *string {
	if o == nil {
		return nil
	}
	return o.Email
}

func (o *GetCustomersResponseBody) GetAvatar() *string {
	if o == nil {
		return nil
	}
	return o.Avatar
}

func (o *GetCustomersResponseBody) GetCountry() *string {
	if o == nil {
		return nil
	}
	return o.Country
}

func (o *GetCustomersResponseBody) GetCreatedAt() string {
	if o == nil {
		return ""
	}
	return o.CreatedAt
}

func (o *GetCustomersResponseBody) GetLink() *Link {
	if o == nil {
		return nil
	}
	return o.Link
}

func (o *GetCustomersResponseBody) GetPartner() *GetCustomersPartner {
	if o == nil {
		return nil
	}
	return o.Partner
}

func (o *GetCustomersResponseBody) GetDiscount() *Discount {
	if o == nil {
		return nil
	}
	return o.Discount
}
