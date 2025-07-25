// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dubinc/dub-go/internal/utils"
	"github.com/dubinc/dub-go/models/components"
)

// Event - The type of event to retrieve analytics for. Defaults to `clicks`.
type Event string

const (
	EventClicks    Event = "clicks"
	EventLeads     Event = "leads"
	EventSales     Event = "sales"
	EventComposite Event = "composite"
)

func (e Event) ToPointer() *Event {
	return &e
}
func (e *Event) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "clicks":
		fallthrough
	case "leads":
		fallthrough
	case "sales":
		fallthrough
	case "composite":
		*e = Event(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Event: %v", v)
	}
}

// QueryParamGroupBy - The parameter to group the analytics data points by. Defaults to `count` if undefined.
type QueryParamGroupBy string

const (
	QueryParamGroupByCount        QueryParamGroupBy = "count"
	QueryParamGroupByTimeseries   QueryParamGroupBy = "timeseries"
	QueryParamGroupByContinents   QueryParamGroupBy = "continents"
	QueryParamGroupByRegions      QueryParamGroupBy = "regions"
	QueryParamGroupByCountries    QueryParamGroupBy = "countries"
	QueryParamGroupByCities       QueryParamGroupBy = "cities"
	QueryParamGroupByDevices      QueryParamGroupBy = "devices"
	QueryParamGroupByBrowsers     QueryParamGroupBy = "browsers"
	QueryParamGroupByOs           QueryParamGroupBy = "os"
	QueryParamGroupByTrigger      QueryParamGroupBy = "trigger"
	QueryParamGroupByTriggers     QueryParamGroupBy = "triggers"
	QueryParamGroupByReferers     QueryParamGroupBy = "referers"
	QueryParamGroupByRefererUrls  QueryParamGroupBy = "referer_urls"
	QueryParamGroupByTopPartners  QueryParamGroupBy = "top_partners"
	QueryParamGroupByTopLinks     QueryParamGroupBy = "top_links"
	QueryParamGroupByTopUrls      QueryParamGroupBy = "top_urls"
	QueryParamGroupByUtmSources   QueryParamGroupBy = "utm_sources"
	QueryParamGroupByUtmMediums   QueryParamGroupBy = "utm_mediums"
	QueryParamGroupByUtmCampaigns QueryParamGroupBy = "utm_campaigns"
	QueryParamGroupByUtmTerms     QueryParamGroupBy = "utm_terms"
	QueryParamGroupByUtmContents  QueryParamGroupBy = "utm_contents"
)

func (e QueryParamGroupBy) ToPointer() *QueryParamGroupBy {
	return &e
}
func (e *QueryParamGroupBy) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "count":
		fallthrough
	case "timeseries":
		fallthrough
	case "continents":
		fallthrough
	case "regions":
		fallthrough
	case "countries":
		fallthrough
	case "cities":
		fallthrough
	case "devices":
		fallthrough
	case "browsers":
		fallthrough
	case "os":
		fallthrough
	case "trigger":
		fallthrough
	case "triggers":
		fallthrough
	case "referers":
		fallthrough
	case "referer_urls":
		fallthrough
	case "top_partners":
		fallthrough
	case "top_links":
		fallthrough
	case "top_urls":
		fallthrough
	case "utm_sources":
		fallthrough
	case "utm_mediums":
		fallthrough
	case "utm_campaigns":
		fallthrough
	case "utm_terms":
		fallthrough
	case "utm_contents":
		*e = QueryParamGroupBy(v)
		return nil
	default:
		return fmt.Errorf("invalid value for QueryParamGroupBy: %v", v)
	}
}

// Interval - The interval to retrieve analytics for. If undefined, defaults to 24h.
type Interval string

const (
	IntervalTwentyFourh Interval = "24h"
	IntervalSevend      Interval = "7d"
	IntervalThirtyd     Interval = "30d"
	IntervalNinetyd     Interval = "90d"
	IntervalOney        Interval = "1y"
	IntervalMtd         Interval = "mtd"
	IntervalQtd         Interval = "qtd"
	IntervalYtd         Interval = "ytd"
	IntervalAll         Interval = "all"
)

func (e Interval) ToPointer() *Interval {
	return &e
}
func (e *Interval) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "24h":
		fallthrough
	case "7d":
		fallthrough
	case "30d":
		fallthrough
	case "90d":
		fallthrough
	case "1y":
		fallthrough
	case "mtd":
		fallthrough
	case "qtd":
		fallthrough
	case "ytd":
		fallthrough
	case "all":
		*e = Interval(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Interval: %v", v)
	}
}

// Trigger - The trigger to retrieve analytics for. If undefined, return both QR and link clicks.
type Trigger string

const (
	TriggerQr       Trigger = "qr"
	TriggerLink     Trigger = "link"
	TriggerPageview Trigger = "pageview"
	TriggerDeeplink Trigger = "deeplink"
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
		fallthrough
	case "pageview":
		fallthrough
	case "deeplink":
		*e = Trigger(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Trigger: %v", v)
	}
}

type RetrieveAnalyticsQueryParamTagIdsType string

const (
	RetrieveAnalyticsQueryParamTagIdsTypeStr        RetrieveAnalyticsQueryParamTagIdsType = "str"
	RetrieveAnalyticsQueryParamTagIdsTypeArrayOfStr RetrieveAnalyticsQueryParamTagIdsType = "arrayOfStr"
)

// RetrieveAnalyticsQueryParamTagIds - The tag IDs to retrieve analytics for.
type RetrieveAnalyticsQueryParamTagIds struct {
	Str        *string  `queryParam:"inline"`
	ArrayOfStr []string `queryParam:"inline"`

	Type RetrieveAnalyticsQueryParamTagIdsType
}

func CreateRetrieveAnalyticsQueryParamTagIdsStr(str string) RetrieveAnalyticsQueryParamTagIds {
	typ := RetrieveAnalyticsQueryParamTagIdsTypeStr

	return RetrieveAnalyticsQueryParamTagIds{
		Str:  &str,
		Type: typ,
	}
}

func CreateRetrieveAnalyticsQueryParamTagIdsArrayOfStr(arrayOfStr []string) RetrieveAnalyticsQueryParamTagIds {
	typ := RetrieveAnalyticsQueryParamTagIdsTypeArrayOfStr

	return RetrieveAnalyticsQueryParamTagIds{
		ArrayOfStr: arrayOfStr,
		Type:       typ,
	}
}

func (u *RetrieveAnalyticsQueryParamTagIds) UnmarshalJSON(data []byte) error {

	var str string = ""
	if err := utils.UnmarshalJSON(data, &str, "", true, false); err == nil {
		u.Str = &str
		u.Type = RetrieveAnalyticsQueryParamTagIdsTypeStr
		return nil
	}

	var arrayOfStr []string = []string{}
	if err := utils.UnmarshalJSON(data, &arrayOfStr, "", true, false); err == nil {
		u.ArrayOfStr = arrayOfStr
		u.Type = RetrieveAnalyticsQueryParamTagIdsTypeArrayOfStr
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for RetrieveAnalyticsQueryParamTagIds", string(data))
}

func (u RetrieveAnalyticsQueryParamTagIds) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.ArrayOfStr != nil {
		return utils.MarshalJSON(u.ArrayOfStr, "", true)
	}

	return nil, errors.New("could not marshal union type RetrieveAnalyticsQueryParamTagIds: all fields are null")
}

// SaleType - Filter sales by type: 'new' for first-time purchases, 'recurring' for repeat purchases. If undefined, returns both.
type SaleType string

const (
	SaleTypeNew       SaleType = "new"
	SaleTypeRecurring SaleType = "recurring"
)

func (e SaleType) ToPointer() *SaleType {
	return &e
}
func (e *SaleType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "new":
		fallthrough
	case "recurring":
		*e = SaleType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SaleType: %v", v)
	}
}

type RetrieveAnalyticsRequest struct {
	// The type of event to retrieve analytics for. Defaults to `clicks`.
	Event *Event `default:"clicks" queryParam:"style=form,explode=true,name=event"`
	// The parameter to group the analytics data points by. Defaults to `count` if undefined.
	GroupBy *QueryParamGroupBy `default:"count" queryParam:"style=form,explode=true,name=groupBy"`
	// The domain to filter analytics for.
	Domain *string `queryParam:"style=form,explode=true,name=domain"`
	// The slug of the short link to retrieve analytics for. Must be used along with the corresponding `domain` of the short link to fetch analytics for a specific short link.
	Key *string `queryParam:"style=form,explode=true,name=key"`
	// The unique ID of the short link on Dub to retrieve analytics for.
	LinkID *string `queryParam:"style=form,explode=true,name=linkId"`
	// The ID of the link in the your database. Must be prefixed with 'ext_' when passed as a query parameter.
	ExternalID *string `queryParam:"style=form,explode=true,name=externalId"`
	// The ID of the tenant that created the link inside your system.
	TenantID *string `queryParam:"style=form,explode=true,name=tenantId"`
	// The ID of the program to retrieve analytics for.
	ProgramID *string `queryParam:"style=form,explode=true,name=programId"`
	// The ID of the partner to retrieve analytics for.
	PartnerID *string `queryParam:"style=form,explode=true,name=partnerId"`
	// The ID of the customer to retrieve analytics for.
	CustomerID *string `queryParam:"style=form,explode=true,name=customerId"`
	// The interval to retrieve analytics for. If undefined, defaults to 24h.
	Interval *Interval `queryParam:"style=form,explode=true,name=interval"`
	// The start date and time when to retrieve analytics from. If set, takes precedence over `interval`.
	Start *string `queryParam:"style=form,explode=true,name=start"`
	// The end date and time when to retrieve analytics from. If not provided, defaults to the current date. If set along with `start`, takes precedence over `interval`.
	End *string `queryParam:"style=form,explode=true,name=end"`
	// The IANA time zone code for aligning timeseries granularity (e.g. America/New_York). Defaults to UTC.
	Timezone *string `default:"UTC" queryParam:"style=form,explode=true,name=timezone"`
	// The country to retrieve analytics for. Must be passed as a 2-letter ISO 3166-1 country code. Learn more: https://d.to/geo
	Country *components.CountryCode `queryParam:"style=form,explode=true,name=country"`
	// The city to retrieve analytics for.
	City *string `queryParam:"style=form,explode=true,name=city"`
	// The ISO 3166-2 region code to retrieve analytics for.
	Region *string `queryParam:"style=form,explode=true,name=region"`
	// The continent to retrieve analytics for.
	Continent *components.ContinentCode `queryParam:"style=form,explode=true,name=continent"`
	// The device to retrieve analytics for.
	Device *string `queryParam:"style=form,explode=true,name=device"`
	// The browser to retrieve analytics for.
	Browser *string `queryParam:"style=form,explode=true,name=browser"`
	// The OS to retrieve analytics for.
	Os *string `queryParam:"style=form,explode=true,name=os"`
	// The trigger to retrieve analytics for. If undefined, return both QR and link clicks.
	Trigger *Trigger `queryParam:"style=form,explode=true,name=trigger"`
	// The referer to retrieve analytics for.
	Referer *string `queryParam:"style=form,explode=true,name=referer"`
	// The full referer URL to retrieve analytics for.
	RefererURL *string `queryParam:"style=form,explode=true,name=refererUrl"`
	// The URL to retrieve analytics for.
	URL *string `queryParam:"style=form,explode=true,name=url"`
	// Deprecated. Use `tagIds` instead. The tag ID to retrieve analytics for.
	TagID *string `queryParam:"style=form,explode=true,name=tagId"`
	// The tag IDs to retrieve analytics for.
	TagIds *RetrieveAnalyticsQueryParamTagIds `queryParam:"style=form,explode=true,name=tagIds"`
	// The folder ID to retrieve analytics for. If not provided, return analytics for unsorted links.
	FolderID *string `queryParam:"style=form,explode=true,name=folderId"`
	// Deprecated. Use the `trigger` field instead. Filter for QR code scans. If true, filter for QR codes only. If false, filter for links only. If undefined, return both.
	Qr *bool `queryParam:"style=form,explode=true,name=qr"`
	// Filter for root domains. If true, filter for domains only. If false, filter for links only. If undefined, return both.
	Root *bool `queryParam:"style=form,explode=true,name=root"`
	// Filter sales by type: 'new' for first-time purchases, 'recurring' for repeat purchases. If undefined, returns both.
	SaleType *SaleType `queryParam:"style=form,explode=true,name=saleType"`
	// The UTM source of the short link.
	UtmSource *string `queryParam:"style=form,explode=true,name=utm_source"`
	// The UTM medium of the short link.
	UtmMedium *string `queryParam:"style=form,explode=true,name=utm_medium"`
	// The UTM campaign of the short link.
	UtmCampaign *string `queryParam:"style=form,explode=true,name=utm_campaign"`
	// The UTM term of the short link.
	UtmTerm *string `queryParam:"style=form,explode=true,name=utm_term"`
	// The UTM content of the short link.
	UtmContent *string `queryParam:"style=form,explode=true,name=utm_content"`
}

func (r RetrieveAnalyticsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *RetrieveAnalyticsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *RetrieveAnalyticsRequest) GetEvent() *Event {
	if o == nil {
		return nil
	}
	return o.Event
}

func (o *RetrieveAnalyticsRequest) GetGroupBy() *QueryParamGroupBy {
	if o == nil {
		return nil
	}
	return o.GroupBy
}

func (o *RetrieveAnalyticsRequest) GetDomain() *string {
	if o == nil {
		return nil
	}
	return o.Domain
}

func (o *RetrieveAnalyticsRequest) GetKey() *string {
	if o == nil {
		return nil
	}
	return o.Key
}

func (o *RetrieveAnalyticsRequest) GetLinkID() *string {
	if o == nil {
		return nil
	}
	return o.LinkID
}

func (o *RetrieveAnalyticsRequest) GetExternalID() *string {
	if o == nil {
		return nil
	}
	return o.ExternalID
}

func (o *RetrieveAnalyticsRequest) GetTenantID() *string {
	if o == nil {
		return nil
	}
	return o.TenantID
}

func (o *RetrieveAnalyticsRequest) GetProgramID() *string {
	if o == nil {
		return nil
	}
	return o.ProgramID
}

func (o *RetrieveAnalyticsRequest) GetPartnerID() *string {
	if o == nil {
		return nil
	}
	return o.PartnerID
}

func (o *RetrieveAnalyticsRequest) GetCustomerID() *string {
	if o == nil {
		return nil
	}
	return o.CustomerID
}

func (o *RetrieveAnalyticsRequest) GetInterval() *Interval {
	if o == nil {
		return nil
	}
	return o.Interval
}

func (o *RetrieveAnalyticsRequest) GetStart() *string {
	if o == nil {
		return nil
	}
	return o.Start
}

func (o *RetrieveAnalyticsRequest) GetEnd() *string {
	if o == nil {
		return nil
	}
	return o.End
}

func (o *RetrieveAnalyticsRequest) GetTimezone() *string {
	if o == nil {
		return nil
	}
	return o.Timezone
}

func (o *RetrieveAnalyticsRequest) GetCountry() *components.CountryCode {
	if o == nil {
		return nil
	}
	return o.Country
}

func (o *RetrieveAnalyticsRequest) GetCity() *string {
	if o == nil {
		return nil
	}
	return o.City
}

func (o *RetrieveAnalyticsRequest) GetRegion() *string {
	if o == nil {
		return nil
	}
	return o.Region
}

func (o *RetrieveAnalyticsRequest) GetContinent() *components.ContinentCode {
	if o == nil {
		return nil
	}
	return o.Continent
}

func (o *RetrieveAnalyticsRequest) GetDevice() *string {
	if o == nil {
		return nil
	}
	return o.Device
}

func (o *RetrieveAnalyticsRequest) GetBrowser() *string {
	if o == nil {
		return nil
	}
	return o.Browser
}

func (o *RetrieveAnalyticsRequest) GetOs() *string {
	if o == nil {
		return nil
	}
	return o.Os
}

func (o *RetrieveAnalyticsRequest) GetTrigger() *Trigger {
	if o == nil {
		return nil
	}
	return o.Trigger
}

func (o *RetrieveAnalyticsRequest) GetReferer() *string {
	if o == nil {
		return nil
	}
	return o.Referer
}

func (o *RetrieveAnalyticsRequest) GetRefererURL() *string {
	if o == nil {
		return nil
	}
	return o.RefererURL
}

func (o *RetrieveAnalyticsRequest) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}

func (o *RetrieveAnalyticsRequest) GetTagID() *string {
	if o == nil {
		return nil
	}
	return o.TagID
}

func (o *RetrieveAnalyticsRequest) GetTagIds() *RetrieveAnalyticsQueryParamTagIds {
	if o == nil {
		return nil
	}
	return o.TagIds
}

func (o *RetrieveAnalyticsRequest) GetFolderID() *string {
	if o == nil {
		return nil
	}
	return o.FolderID
}

func (o *RetrieveAnalyticsRequest) GetQr() *bool {
	if o == nil {
		return nil
	}
	return o.Qr
}

func (o *RetrieveAnalyticsRequest) GetRoot() *bool {
	if o == nil {
		return nil
	}
	return o.Root
}

func (o *RetrieveAnalyticsRequest) GetSaleType() *SaleType {
	if o == nil {
		return nil
	}
	return o.SaleType
}

func (o *RetrieveAnalyticsRequest) GetUtmSource() *string {
	if o == nil {
		return nil
	}
	return o.UtmSource
}

func (o *RetrieveAnalyticsRequest) GetUtmMedium() *string {
	if o == nil {
		return nil
	}
	return o.UtmMedium
}

func (o *RetrieveAnalyticsRequest) GetUtmCampaign() *string {
	if o == nil {
		return nil
	}
	return o.UtmCampaign
}

func (o *RetrieveAnalyticsRequest) GetUtmTerm() *string {
	if o == nil {
		return nil
	}
	return o.UtmTerm
}

func (o *RetrieveAnalyticsRequest) GetUtmContent() *string {
	if o == nil {
		return nil
	}
	return o.UtmContent
}

type RetrieveAnalyticsResponseBodyType string

const (
	RetrieveAnalyticsResponseBodyTypeAnalyticsCount              RetrieveAnalyticsResponseBodyType = "AnalyticsCount"
	RetrieveAnalyticsResponseBodyTypeArrayOfAnalyticsTimeseries  RetrieveAnalyticsResponseBodyType = "arrayOfAnalyticsTimeseries"
	RetrieveAnalyticsResponseBodyTypeArrayOfAnalyticsContinents  RetrieveAnalyticsResponseBodyType = "arrayOfAnalyticsContinents"
	RetrieveAnalyticsResponseBodyTypeArrayOfAnalyticsCountries   RetrieveAnalyticsResponseBodyType = "arrayOfAnalyticsCountries"
	RetrieveAnalyticsResponseBodyTypeArrayOfAnalyticsRegions     RetrieveAnalyticsResponseBodyType = "arrayOfAnalyticsRegions"
	RetrieveAnalyticsResponseBodyTypeArrayOfAnalyticsCities      RetrieveAnalyticsResponseBodyType = "arrayOfAnalyticsCities"
	RetrieveAnalyticsResponseBodyTypeArrayOfAnalyticsDevices     RetrieveAnalyticsResponseBodyType = "arrayOfAnalyticsDevices"
	RetrieveAnalyticsResponseBodyTypeArrayOfAnalyticsBrowsers    RetrieveAnalyticsResponseBodyType = "arrayOfAnalyticsBrowsers"
	RetrieveAnalyticsResponseBodyTypeArrayOfAnalyticsOS          RetrieveAnalyticsResponseBodyType = "arrayOfAnalyticsOS"
	RetrieveAnalyticsResponseBodyTypeArrayOfAnalyticsTriggers    RetrieveAnalyticsResponseBodyType = "arrayOfAnalyticsTriggers"
	RetrieveAnalyticsResponseBodyTypeArrayOfAnalyticsReferers    RetrieveAnalyticsResponseBodyType = "arrayOfAnalyticsReferers"
	RetrieveAnalyticsResponseBodyTypeArrayOfAnalyticsRefererUrls RetrieveAnalyticsResponseBodyType = "arrayOfAnalyticsRefererUrls"
	RetrieveAnalyticsResponseBodyTypeArrayOfAnalyticsTopLinks    RetrieveAnalyticsResponseBodyType = "arrayOfAnalyticsTopLinks"
	RetrieveAnalyticsResponseBodyTypeArrayOfAnalyticsTopUrls     RetrieveAnalyticsResponseBodyType = "arrayOfAnalyticsTopUrls"
)

// RetrieveAnalyticsResponseBody - Analytics data
type RetrieveAnalyticsResponseBody struct {
	AnalyticsCount              *components.AnalyticsCount        `queryParam:"inline"`
	ArrayOfAnalyticsTimeseries  []components.AnalyticsTimeseries  `queryParam:"inline"`
	ArrayOfAnalyticsContinents  []components.AnalyticsContinents  `queryParam:"inline"`
	ArrayOfAnalyticsCountries   []components.AnalyticsCountries   `queryParam:"inline"`
	ArrayOfAnalyticsRegions     []components.AnalyticsRegions     `queryParam:"inline"`
	ArrayOfAnalyticsCities      []components.AnalyticsCities      `queryParam:"inline"`
	ArrayOfAnalyticsDevices     []components.AnalyticsDevices     `queryParam:"inline"`
	ArrayOfAnalyticsBrowsers    []components.AnalyticsBrowsers    `queryParam:"inline"`
	ArrayOfAnalyticsOS          []components.AnalyticsOS          `queryParam:"inline"`
	ArrayOfAnalyticsTriggers    []components.AnalyticsTriggers    `queryParam:"inline"`
	ArrayOfAnalyticsReferers    []components.AnalyticsReferers    `queryParam:"inline"`
	ArrayOfAnalyticsRefererUrls []components.AnalyticsRefererUrls `queryParam:"inline"`
	ArrayOfAnalyticsTopLinks    []components.AnalyticsTopLinks    `queryParam:"inline"`
	ArrayOfAnalyticsTopUrls     []components.AnalyticsTopUrls     `queryParam:"inline"`

	Type RetrieveAnalyticsResponseBodyType
}

func CreateRetrieveAnalyticsResponseBodyAnalyticsCount(analyticsCount components.AnalyticsCount) RetrieveAnalyticsResponseBody {
	typ := RetrieveAnalyticsResponseBodyTypeAnalyticsCount

	return RetrieveAnalyticsResponseBody{
		AnalyticsCount: &analyticsCount,
		Type:           typ,
	}
}

func CreateRetrieveAnalyticsResponseBodyArrayOfAnalyticsTimeseries(arrayOfAnalyticsTimeseries []components.AnalyticsTimeseries) RetrieveAnalyticsResponseBody {
	typ := RetrieveAnalyticsResponseBodyTypeArrayOfAnalyticsTimeseries

	return RetrieveAnalyticsResponseBody{
		ArrayOfAnalyticsTimeseries: arrayOfAnalyticsTimeseries,
		Type:                       typ,
	}
}

func CreateRetrieveAnalyticsResponseBodyArrayOfAnalyticsContinents(arrayOfAnalyticsContinents []components.AnalyticsContinents) RetrieveAnalyticsResponseBody {
	typ := RetrieveAnalyticsResponseBodyTypeArrayOfAnalyticsContinents

	return RetrieveAnalyticsResponseBody{
		ArrayOfAnalyticsContinents: arrayOfAnalyticsContinents,
		Type:                       typ,
	}
}

func CreateRetrieveAnalyticsResponseBodyArrayOfAnalyticsCountries(arrayOfAnalyticsCountries []components.AnalyticsCountries) RetrieveAnalyticsResponseBody {
	typ := RetrieveAnalyticsResponseBodyTypeArrayOfAnalyticsCountries

	return RetrieveAnalyticsResponseBody{
		ArrayOfAnalyticsCountries: arrayOfAnalyticsCountries,
		Type:                      typ,
	}
}

func CreateRetrieveAnalyticsResponseBodyArrayOfAnalyticsRegions(arrayOfAnalyticsRegions []components.AnalyticsRegions) RetrieveAnalyticsResponseBody {
	typ := RetrieveAnalyticsResponseBodyTypeArrayOfAnalyticsRegions

	return RetrieveAnalyticsResponseBody{
		ArrayOfAnalyticsRegions: arrayOfAnalyticsRegions,
		Type:                    typ,
	}
}

func CreateRetrieveAnalyticsResponseBodyArrayOfAnalyticsCities(arrayOfAnalyticsCities []components.AnalyticsCities) RetrieveAnalyticsResponseBody {
	typ := RetrieveAnalyticsResponseBodyTypeArrayOfAnalyticsCities

	return RetrieveAnalyticsResponseBody{
		ArrayOfAnalyticsCities: arrayOfAnalyticsCities,
		Type:                   typ,
	}
}

func CreateRetrieveAnalyticsResponseBodyArrayOfAnalyticsDevices(arrayOfAnalyticsDevices []components.AnalyticsDevices) RetrieveAnalyticsResponseBody {
	typ := RetrieveAnalyticsResponseBodyTypeArrayOfAnalyticsDevices

	return RetrieveAnalyticsResponseBody{
		ArrayOfAnalyticsDevices: arrayOfAnalyticsDevices,
		Type:                    typ,
	}
}

func CreateRetrieveAnalyticsResponseBodyArrayOfAnalyticsBrowsers(arrayOfAnalyticsBrowsers []components.AnalyticsBrowsers) RetrieveAnalyticsResponseBody {
	typ := RetrieveAnalyticsResponseBodyTypeArrayOfAnalyticsBrowsers

	return RetrieveAnalyticsResponseBody{
		ArrayOfAnalyticsBrowsers: arrayOfAnalyticsBrowsers,
		Type:                     typ,
	}
}

func CreateRetrieveAnalyticsResponseBodyArrayOfAnalyticsOS(arrayOfAnalyticsOS []components.AnalyticsOS) RetrieveAnalyticsResponseBody {
	typ := RetrieveAnalyticsResponseBodyTypeArrayOfAnalyticsOS

	return RetrieveAnalyticsResponseBody{
		ArrayOfAnalyticsOS: arrayOfAnalyticsOS,
		Type:               typ,
	}
}

func CreateRetrieveAnalyticsResponseBodyArrayOfAnalyticsTriggers(arrayOfAnalyticsTriggers []components.AnalyticsTriggers) RetrieveAnalyticsResponseBody {
	typ := RetrieveAnalyticsResponseBodyTypeArrayOfAnalyticsTriggers

	return RetrieveAnalyticsResponseBody{
		ArrayOfAnalyticsTriggers: arrayOfAnalyticsTriggers,
		Type:                     typ,
	}
}

func CreateRetrieveAnalyticsResponseBodyArrayOfAnalyticsReferers(arrayOfAnalyticsReferers []components.AnalyticsReferers) RetrieveAnalyticsResponseBody {
	typ := RetrieveAnalyticsResponseBodyTypeArrayOfAnalyticsReferers

	return RetrieveAnalyticsResponseBody{
		ArrayOfAnalyticsReferers: arrayOfAnalyticsReferers,
		Type:                     typ,
	}
}

func CreateRetrieveAnalyticsResponseBodyArrayOfAnalyticsRefererUrls(arrayOfAnalyticsRefererUrls []components.AnalyticsRefererUrls) RetrieveAnalyticsResponseBody {
	typ := RetrieveAnalyticsResponseBodyTypeArrayOfAnalyticsRefererUrls

	return RetrieveAnalyticsResponseBody{
		ArrayOfAnalyticsRefererUrls: arrayOfAnalyticsRefererUrls,
		Type:                        typ,
	}
}

func CreateRetrieveAnalyticsResponseBodyArrayOfAnalyticsTopLinks(arrayOfAnalyticsTopLinks []components.AnalyticsTopLinks) RetrieveAnalyticsResponseBody {
	typ := RetrieveAnalyticsResponseBodyTypeArrayOfAnalyticsTopLinks

	return RetrieveAnalyticsResponseBody{
		ArrayOfAnalyticsTopLinks: arrayOfAnalyticsTopLinks,
		Type:                     typ,
	}
}

func CreateRetrieveAnalyticsResponseBodyArrayOfAnalyticsTopUrls(arrayOfAnalyticsTopUrls []components.AnalyticsTopUrls) RetrieveAnalyticsResponseBody {
	typ := RetrieveAnalyticsResponseBodyTypeArrayOfAnalyticsTopUrls

	return RetrieveAnalyticsResponseBody{
		ArrayOfAnalyticsTopUrls: arrayOfAnalyticsTopUrls,
		Type:                    typ,
	}
}

func (u *RetrieveAnalyticsResponseBody) UnmarshalJSON(data []byte) error {

	var analyticsCount components.AnalyticsCount = components.AnalyticsCount{}
	if err := utils.UnmarshalJSON(data, &analyticsCount, "", true, false); err == nil {
		u.AnalyticsCount = &analyticsCount
		u.Type = RetrieveAnalyticsResponseBodyTypeAnalyticsCount
		return nil
	}

	var arrayOfAnalyticsTimeseries []components.AnalyticsTimeseries = []components.AnalyticsTimeseries{}
	if err := utils.UnmarshalJSON(data, &arrayOfAnalyticsTimeseries, "", true, false); err == nil {
		u.ArrayOfAnalyticsTimeseries = arrayOfAnalyticsTimeseries
		u.Type = RetrieveAnalyticsResponseBodyTypeArrayOfAnalyticsTimeseries
		return nil
	}

	var arrayOfAnalyticsContinents []components.AnalyticsContinents = []components.AnalyticsContinents{}
	if err := utils.UnmarshalJSON(data, &arrayOfAnalyticsContinents, "", true, false); err == nil {
		u.ArrayOfAnalyticsContinents = arrayOfAnalyticsContinents
		u.Type = RetrieveAnalyticsResponseBodyTypeArrayOfAnalyticsContinents
		return nil
	}

	var arrayOfAnalyticsCountries []components.AnalyticsCountries = []components.AnalyticsCountries{}
	if err := utils.UnmarshalJSON(data, &arrayOfAnalyticsCountries, "", true, false); err == nil {
		u.ArrayOfAnalyticsCountries = arrayOfAnalyticsCountries
		u.Type = RetrieveAnalyticsResponseBodyTypeArrayOfAnalyticsCountries
		return nil
	}

	var arrayOfAnalyticsRegions []components.AnalyticsRegions = []components.AnalyticsRegions{}
	if err := utils.UnmarshalJSON(data, &arrayOfAnalyticsRegions, "", true, false); err == nil {
		u.ArrayOfAnalyticsRegions = arrayOfAnalyticsRegions
		u.Type = RetrieveAnalyticsResponseBodyTypeArrayOfAnalyticsRegions
		return nil
	}

	var arrayOfAnalyticsCities []components.AnalyticsCities = []components.AnalyticsCities{}
	if err := utils.UnmarshalJSON(data, &arrayOfAnalyticsCities, "", true, false); err == nil {
		u.ArrayOfAnalyticsCities = arrayOfAnalyticsCities
		u.Type = RetrieveAnalyticsResponseBodyTypeArrayOfAnalyticsCities
		return nil
	}

	var arrayOfAnalyticsDevices []components.AnalyticsDevices = []components.AnalyticsDevices{}
	if err := utils.UnmarshalJSON(data, &arrayOfAnalyticsDevices, "", true, false); err == nil {
		u.ArrayOfAnalyticsDevices = arrayOfAnalyticsDevices
		u.Type = RetrieveAnalyticsResponseBodyTypeArrayOfAnalyticsDevices
		return nil
	}

	var arrayOfAnalyticsBrowsers []components.AnalyticsBrowsers = []components.AnalyticsBrowsers{}
	if err := utils.UnmarshalJSON(data, &arrayOfAnalyticsBrowsers, "", true, false); err == nil {
		u.ArrayOfAnalyticsBrowsers = arrayOfAnalyticsBrowsers
		u.Type = RetrieveAnalyticsResponseBodyTypeArrayOfAnalyticsBrowsers
		return nil
	}

	var arrayOfAnalyticsOS []components.AnalyticsOS = []components.AnalyticsOS{}
	if err := utils.UnmarshalJSON(data, &arrayOfAnalyticsOS, "", true, false); err == nil {
		u.ArrayOfAnalyticsOS = arrayOfAnalyticsOS
		u.Type = RetrieveAnalyticsResponseBodyTypeArrayOfAnalyticsOS
		return nil
	}

	var arrayOfAnalyticsTriggers []components.AnalyticsTriggers = []components.AnalyticsTriggers{}
	if err := utils.UnmarshalJSON(data, &arrayOfAnalyticsTriggers, "", true, false); err == nil {
		u.ArrayOfAnalyticsTriggers = arrayOfAnalyticsTriggers
		u.Type = RetrieveAnalyticsResponseBodyTypeArrayOfAnalyticsTriggers
		return nil
	}

	var arrayOfAnalyticsReferers []components.AnalyticsReferers = []components.AnalyticsReferers{}
	if err := utils.UnmarshalJSON(data, &arrayOfAnalyticsReferers, "", true, false); err == nil {
		u.ArrayOfAnalyticsReferers = arrayOfAnalyticsReferers
		u.Type = RetrieveAnalyticsResponseBodyTypeArrayOfAnalyticsReferers
		return nil
	}

	var arrayOfAnalyticsRefererUrls []components.AnalyticsRefererUrls = []components.AnalyticsRefererUrls{}
	if err := utils.UnmarshalJSON(data, &arrayOfAnalyticsRefererUrls, "", true, false); err == nil {
		u.ArrayOfAnalyticsRefererUrls = arrayOfAnalyticsRefererUrls
		u.Type = RetrieveAnalyticsResponseBodyTypeArrayOfAnalyticsRefererUrls
		return nil
	}

	var arrayOfAnalyticsTopLinks []components.AnalyticsTopLinks = []components.AnalyticsTopLinks{}
	if err := utils.UnmarshalJSON(data, &arrayOfAnalyticsTopLinks, "", true, false); err == nil {
		u.ArrayOfAnalyticsTopLinks = arrayOfAnalyticsTopLinks
		u.Type = RetrieveAnalyticsResponseBodyTypeArrayOfAnalyticsTopLinks
		return nil
	}

	var arrayOfAnalyticsTopUrls []components.AnalyticsTopUrls = []components.AnalyticsTopUrls{}
	if err := utils.UnmarshalJSON(data, &arrayOfAnalyticsTopUrls, "", true, false); err == nil {
		u.ArrayOfAnalyticsTopUrls = arrayOfAnalyticsTopUrls
		u.Type = RetrieveAnalyticsResponseBodyTypeArrayOfAnalyticsTopUrls
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for RetrieveAnalyticsResponseBody", string(data))
}

func (u RetrieveAnalyticsResponseBody) MarshalJSON() ([]byte, error) {
	if u.AnalyticsCount != nil {
		return utils.MarshalJSON(u.AnalyticsCount, "", true)
	}

	if u.ArrayOfAnalyticsTimeseries != nil {
		return utils.MarshalJSON(u.ArrayOfAnalyticsTimeseries, "", true)
	}

	if u.ArrayOfAnalyticsContinents != nil {
		return utils.MarshalJSON(u.ArrayOfAnalyticsContinents, "", true)
	}

	if u.ArrayOfAnalyticsCountries != nil {
		return utils.MarshalJSON(u.ArrayOfAnalyticsCountries, "", true)
	}

	if u.ArrayOfAnalyticsRegions != nil {
		return utils.MarshalJSON(u.ArrayOfAnalyticsRegions, "", true)
	}

	if u.ArrayOfAnalyticsCities != nil {
		return utils.MarshalJSON(u.ArrayOfAnalyticsCities, "", true)
	}

	if u.ArrayOfAnalyticsDevices != nil {
		return utils.MarshalJSON(u.ArrayOfAnalyticsDevices, "", true)
	}

	if u.ArrayOfAnalyticsBrowsers != nil {
		return utils.MarshalJSON(u.ArrayOfAnalyticsBrowsers, "", true)
	}

	if u.ArrayOfAnalyticsOS != nil {
		return utils.MarshalJSON(u.ArrayOfAnalyticsOS, "", true)
	}

	if u.ArrayOfAnalyticsTriggers != nil {
		return utils.MarshalJSON(u.ArrayOfAnalyticsTriggers, "", true)
	}

	if u.ArrayOfAnalyticsReferers != nil {
		return utils.MarshalJSON(u.ArrayOfAnalyticsReferers, "", true)
	}

	if u.ArrayOfAnalyticsRefererUrls != nil {
		return utils.MarshalJSON(u.ArrayOfAnalyticsRefererUrls, "", true)
	}

	if u.ArrayOfAnalyticsTopLinks != nil {
		return utils.MarshalJSON(u.ArrayOfAnalyticsTopLinks, "", true)
	}

	if u.ArrayOfAnalyticsTopUrls != nil {
		return utils.MarshalJSON(u.ArrayOfAnalyticsTopUrls, "", true)
	}

	return nil, errors.New("could not marshal union type RetrieveAnalyticsResponseBody: all fields are null")
}
