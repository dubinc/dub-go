// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type SalesTimeseries struct {
	// The starting timestamp of the interval
	Start string `json:"start"`
	// The number of sales in the interval
	Sales float64 `json:"sales"`
	// The total amount of sales in the interval
	Amount float64 `json:"amount"`
}

func (o *SalesTimeseries) GetStart() string {
	if o == nil {
		return ""
	}
	return o.Start
}

func (o *SalesTimeseries) GetSales() float64 {
	if o == nil {
		return 0.0
	}
	return o.Sales
}

func (o *SalesTimeseries) GetAmount() float64 {
	if o == nil {
		return 0.0
	}
	return o.Amount
}
