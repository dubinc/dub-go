// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
	"github.com/dubinc/dub-go/internal/utils"
)

// Trigger - The type of trigger method: link click or QR scan
type Trigger string

const (
	TriggerQr   Trigger = "qr"
	TriggerLink Trigger = "link"
)

func (e Trigger) ToPointer() *Trigger {
	return &e
}
func (e *Trigger) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "qr":
		fallthrough
	case "link":
		*e = Trigger(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Trigger: %v", v)
	}
}

type AnalyticsTriggers struct {
	// The type of trigger method: link click or QR scan
	Trigger Trigger `json:"trigger"`
	// The number of clicks from this trigger method
	Clicks *float64 `default:"0" json:"clicks"`
	// The number of leads from this trigger method
	Leads *float64 `default:"0" json:"leads"`
	// The number of sales from this trigger method
	Sales *float64 `default:"0" json:"sales"`
	// The total amount of sales from this trigger method, in cents
	SaleAmount *float64 `default:"0" json:"saleAmount"`
}

func (a AnalyticsTriggers) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AnalyticsTriggers) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *AnalyticsTriggers) GetTrigger() Trigger {
	if o == nil {
		return Trigger("")
	}
	return o.Trigger
}

func (o *AnalyticsTriggers) GetClicks() *float64 {
	if o == nil {
		return nil
	}
	return o.Clicks
}

func (o *AnalyticsTriggers) GetLeads() *float64 {
	if o == nil {
		return nil
	}
	return o.Leads
}

func (o *AnalyticsTriggers) GetSales() *float64 {
	if o == nil {
		return nil
	}
	return o.Sales
}

func (o *AnalyticsTriggers) GetSaleAmount() *float64 {
	if o == nil {
		return nil
	}
	return o.SaleAmount
}