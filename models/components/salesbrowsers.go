// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type SalesBrowsers struct {
	// The name of the browser
	Browser string `json:"browser"`
	// The number of sales from this browser
	Sales float64 `json:"sales"`
	// The total amount of sales from this browser
	Amount float64 `json:"amount"`
}

func (o *SalesBrowsers) GetBrowser() string {
	if o == nil {
		return ""
	}
	return o.Browser
}

func (o *SalesBrowsers) GetSales() float64 {
	if o == nil {
		return 0.0
	}
	return o.Sales
}

func (o *SalesBrowsers) GetAmount() float64 {
	if o == nil {
		return 0.0
	}
	return o.Amount
}
