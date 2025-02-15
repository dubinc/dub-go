// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"github.com/dubinc/dub-go/internal/utils"
)

type PartnerAnalyticsTimeseries struct {
	// The starting timestamp of the interval
	Start string `json:"start"`
	// The number of clicks in the interval
	Clicks *float64 `default:"0" json:"clicks"`
	// The number of leads in the interval
	Leads *float64 `default:"0" json:"leads"`
	// The number of sales in the interval
	Sales *float64 `default:"0" json:"sales"`
	// The total amount of sales in the interval, in cents
	SaleAmount *float64 `default:"0" json:"saleAmount"`
	Earnings   *float64 `default:"0" json:"earnings"`
}

func (p PartnerAnalyticsTimeseries) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *PartnerAnalyticsTimeseries) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *PartnerAnalyticsTimeseries) GetStart() string {
	if o == nil {
		return ""
	}
	return o.Start
}

func (o *PartnerAnalyticsTimeseries) GetClicks() *float64 {
	if o == nil {
		return nil
	}
	return o.Clicks
}

func (o *PartnerAnalyticsTimeseries) GetLeads() *float64 {
	if o == nil {
		return nil
	}
	return o.Leads
}

func (o *PartnerAnalyticsTimeseries) GetSales() *float64 {
	if o == nil {
		return nil
	}
	return o.Sales
}

func (o *PartnerAnalyticsTimeseries) GetSaleAmount() *float64 {
	if o == nil {
		return nil
	}
	return o.SaleAmount
}

func (o *PartnerAnalyticsTimeseries) GetEarnings() *float64 {
	if o == nil {
		return nil
	}
	return o.Earnings
}
