// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/dubinc/dub-go/internal/utils"
)

// PaymentProcessor - The payment processor via which the sale was made.
type PaymentProcessor string

const (
	PaymentProcessorStripe     PaymentProcessor = "stripe"
	PaymentProcessorShopify    PaymentProcessor = "shopify"
	PaymentProcessorPolar      PaymentProcessor = "polar"
	PaymentProcessorPaddle     PaymentProcessor = "paddle"
	PaymentProcessorRevenuecat PaymentProcessor = "revenuecat"
	PaymentProcessorCustom     PaymentProcessor = "custom"
)

func (e PaymentProcessor) ToPointer() *PaymentProcessor {
	return &e
}
func (e *PaymentProcessor) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "stripe":
		fallthrough
	case "shopify":
		fallthrough
	case "polar":
		fallthrough
	case "paddle":
		fallthrough
	case "revenuecat":
		fallthrough
	case "custom":
		*e = PaymentProcessor(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PaymentProcessor: %v", v)
	}
}

type TrackSaleRequestBody struct {
	// The unique ID of the customer in your system. Will be used to identify and attribute all future events to this customer.
	CustomerExternalID string `json:"customerExternalId"`
	// The amount of the sale in cents (for all two-decimal currencies). If the sale is in a zero-decimal currency, pass the full integer value (e.g. `1437` JPY). Learn more: https://d.to/currency
	Amount int64 `json:"amount"`
	// The currency of the sale. Accepts ISO 4217 currency codes. Sales will be automatically converted and stored as USD at the latest exchange rates. Learn more: https://d.to/currency
	Currency *string `default:"usd" json:"currency"`
	// The name of the sale event. Recommended format: `Invoice paid` or `Subscription created`.
	EventName *string `default:"Purchase" json:"eventName"`
	// The payment processor via which the sale was made.
	PaymentProcessor PaymentProcessor `json:"paymentProcessor"`
	// The invoice ID of the sale. Can be used as a idempotency key – only one sale event can be recorded for a given invoice ID.
	InvoiceID *string `default:"null" json:"invoiceId"`
	// The name of the lead event that occurred before the sale (case-sensitive). This is used to associate the sale event with a particular lead event (instead of the latest lead event for a link-customer combination, which is the default behavior).
	LeadEventName *string `default:"null" json:"leadEventName"`
	// Additional metadata to be stored with the sale event. Max 10,000 characters when stringified.
	Metadata map[string]any `json:"metadata,omitempty"`
}

func (t TrackSaleRequestBody) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *TrackSaleRequestBody) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *TrackSaleRequestBody) GetCustomerExternalID() string {
	if o == nil {
		return ""
	}
	return o.CustomerExternalID
}

func (o *TrackSaleRequestBody) GetAmount() int64 {
	if o == nil {
		return 0
	}
	return o.Amount
}

func (o *TrackSaleRequestBody) GetCurrency() *string {
	if o == nil {
		return nil
	}
	return o.Currency
}

func (o *TrackSaleRequestBody) GetEventName() *string {
	if o == nil {
		return nil
	}
	return o.EventName
}

func (o *TrackSaleRequestBody) GetPaymentProcessor() PaymentProcessor {
	if o == nil {
		return PaymentProcessor("")
	}
	return o.PaymentProcessor
}

func (o *TrackSaleRequestBody) GetInvoiceID() *string {
	if o == nil {
		return nil
	}
	return o.InvoiceID
}

func (o *TrackSaleRequestBody) GetLeadEventName() *string {
	if o == nil {
		return nil
	}
	return o.LeadEventName
}

func (o *TrackSaleRequestBody) GetMetadata() map[string]any {
	if o == nil {
		return nil
	}
	return o.Metadata
}

type TrackSaleCustomer struct {
	ID         string  `json:"id"`
	Name       *string `json:"name"`
	Email      *string `json:"email"`
	Avatar     *string `json:"avatar"`
	ExternalID *string `json:"externalId"`
}

func (o *TrackSaleCustomer) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *TrackSaleCustomer) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *TrackSaleCustomer) GetEmail() *string {
	if o == nil {
		return nil
	}
	return o.Email
}

func (o *TrackSaleCustomer) GetAvatar() *string {
	if o == nil {
		return nil
	}
	return o.Avatar
}

func (o *TrackSaleCustomer) GetExternalID() *string {
	if o == nil {
		return nil
	}
	return o.ExternalID
}

type Sale struct {
	Amount           float64        `json:"amount"`
	Currency         string         `json:"currency"`
	PaymentProcessor string         `json:"paymentProcessor"`
	InvoiceID        *string        `json:"invoiceId"`
	Metadata         map[string]any `json:"metadata"`
}

func (o *Sale) GetAmount() float64 {
	if o == nil {
		return 0.0
	}
	return o.Amount
}

func (o *Sale) GetCurrency() string {
	if o == nil {
		return ""
	}
	return o.Currency
}

func (o *Sale) GetPaymentProcessor() string {
	if o == nil {
		return ""
	}
	return o.PaymentProcessor
}

func (o *Sale) GetInvoiceID() *string {
	if o == nil {
		return nil
	}
	return o.InvoiceID
}

func (o *Sale) GetMetadata() map[string]any {
	if o == nil {
		return nil
	}
	return o.Metadata
}

// TrackSaleResponseBody - A sale was tracked.
type TrackSaleResponseBody struct {
	EventName string             `json:"eventName"`
	Customer  *TrackSaleCustomer `json:"customer"`
	Sale      *Sale              `json:"sale"`
}

func (o *TrackSaleResponseBody) GetEventName() string {
	if o == nil {
		return ""
	}
	return o.EventName
}

func (o *TrackSaleResponseBody) GetCustomer() *TrackSaleCustomer {
	if o == nil {
		return nil
	}
	return o.Customer
}

func (o *TrackSaleResponseBody) GetSale() *Sale {
	if o == nil {
		return nil
	}
	return o.Sale
}
