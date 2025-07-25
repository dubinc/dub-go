// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package dubgo

// Generated from OpenAPI doc version 0.0.1 and generator version 2.662.0

import (
	"context"
	"fmt"
	"github.com/dubinc/dub-go/internal/config"
	"github.com/dubinc/dub-go/internal/hooks"
	"github.com/dubinc/dub-go/internal/utils"
	"github.com/dubinc/dub-go/models/components"
	"github.com/dubinc/dub-go/retry"
	"net/http"
	"time"
)

// ServerList contains the list of servers available to the SDK
var ServerList = []string{
	// Production API
	"https://api.dub.co",
}

// HTTPClient provides an interface for supplying the SDK with a custom HTTP client
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// String provides a helper function to return a pointer to a string
func String(s string) *string { return &s }

// Bool provides a helper function to return a pointer to a bool
func Bool(b bool) *bool { return &b }

// Int provides a helper function to return a pointer to an int
func Int(i int) *int { return &i }

// Int64 provides a helper function to return a pointer to an int64
func Int64(i int64) *int64 { return &i }

// Float32 provides a helper function to return a pointer to a float32
func Float32(f float32) *float32 { return &f }

// Float64 provides a helper function to return a pointer to a float64
func Float64(f float64) *float64 { return &f }

// Pointer provides a helper function to return a pointer to a type
func Pointer[T any](v T) *T { return &v }

// Dub API: Dub is link management infrastructure for companies to create marketing campaigns, link sharing features, and referral programs.
type Dub struct {
	SDKVersion string
	Links      *Links
	// Retrieve analytics for a partner
	// Retrieve analytics for a partner within a program. The response type vary based on the `groupBy` query parameter.
	Analytics   *Analytics
	Events      *Events
	Tags        *Tags
	Folders     *Folders
	Domains     *Domains
	Track       *Track
	Customers   *Customers
	Partners    *Partners
	Commissions *Commissions
	Workspaces  *Workspaces
	EmbedTokens *EmbedTokens
	QRCodes     *QRCodes

	sdkConfiguration config.SDKConfiguration
	hooks            *hooks.Hooks
}

type SDKOption func(*Dub)

// WithServerURL allows the overriding of the default server URL
func WithServerURL(serverURL string) SDKOption {
	return func(sdk *Dub) {
		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithTemplatedServerURL allows the overriding of the default server URL with a templated URL populated with the provided parameters
func WithTemplatedServerURL(serverURL string, params map[string]string) SDKOption {
	return func(sdk *Dub) {
		if params != nil {
			serverURL = utils.ReplaceParameters(serverURL, params)
		}

		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithServerIndex allows the overriding of the default server by index
func WithServerIndex(serverIndex int) SDKOption {
	return func(sdk *Dub) {
		if serverIndex < 0 || serverIndex >= len(ServerList) {
			panic(fmt.Errorf("server index %d out of range", serverIndex))
		}

		sdk.sdkConfiguration.ServerIndex = serverIndex
	}
}

// WithClient allows the overriding of the default HTTP client used by the SDK
func WithClient(client HTTPClient) SDKOption {
	return func(sdk *Dub) {
		sdk.sdkConfiguration.Client = client
	}
}

// WithSecurity configures the SDK to use the provided security details
func WithSecurity(token string) SDKOption {
	return func(sdk *Dub) {
		security := components.Security{Token: token}
		sdk.sdkConfiguration.Security = utils.AsSecuritySource(&security)
	}
}

// WithSecuritySource configures the SDK to invoke the Security Source function on each method call to determine authentication
func WithSecuritySource(security func(context.Context) (components.Security, error)) SDKOption {
	return func(sdk *Dub) {
		sdk.sdkConfiguration.Security = func(ctx context.Context) (interface{}, error) {
			return security(ctx)
		}
	}
}

func WithRetryConfig(retryConfig retry.Config) SDKOption {
	return func(sdk *Dub) {
		sdk.sdkConfiguration.RetryConfig = &retryConfig
	}
}

// WithTimeout Optional request timeout applied to each operation
func WithTimeout(timeout time.Duration) SDKOption {
	return func(sdk *Dub) {
		sdk.sdkConfiguration.Timeout = &timeout
	}
}

// New creates a new instance of the SDK with the provided options
func New(opts ...SDKOption) *Dub {
	sdk := &Dub{
		SDKVersion: "0.16.9",
		sdkConfiguration: config.SDKConfiguration{
			UserAgent:  "speakeasy-sdk/go 0.16.9 2.662.0 0.0.1 github.com/dubinc/dub-go",
			ServerList: ServerList,
		},
		hooks: hooks.New(),
	}
	for _, opt := range opts {
		opt(sdk)
	}

	// Use WithClient to override the default client if you would like to customize the timeout
	if sdk.sdkConfiguration.Client == nil {
		sdk.sdkConfiguration.Client = &http.Client{Timeout: 60 * time.Second}
	}

	currentServerURL, _ := sdk.sdkConfiguration.GetServerDetails()
	serverURL := currentServerURL
	serverURL, sdk.sdkConfiguration.Client = sdk.hooks.SDKInit(currentServerURL, sdk.sdkConfiguration.Client)
	if currentServerURL != serverURL {
		sdk.sdkConfiguration.ServerURL = serverURL
	}

	sdk.Links = newLinks(sdk, sdk.sdkConfiguration, sdk.hooks)
	sdk.Analytics = newAnalytics(sdk, sdk.sdkConfiguration, sdk.hooks)
	sdk.Events = newEvents(sdk, sdk.sdkConfiguration, sdk.hooks)
	sdk.Tags = newTags(sdk, sdk.sdkConfiguration, sdk.hooks)
	sdk.Folders = newFolders(sdk, sdk.sdkConfiguration, sdk.hooks)
	sdk.Domains = newDomains(sdk, sdk.sdkConfiguration, sdk.hooks)
	sdk.Track = newTrack(sdk, sdk.sdkConfiguration, sdk.hooks)
	sdk.Customers = newCustomers(sdk, sdk.sdkConfiguration, sdk.hooks)
	sdk.Partners = newPartners(sdk, sdk.sdkConfiguration, sdk.hooks)
	sdk.Commissions = newCommissions(sdk, sdk.sdkConfiguration, sdk.hooks)
	sdk.Workspaces = newWorkspaces(sdk, sdk.sdkConfiguration, sdk.hooks)
	sdk.EmbedTokens = newEmbedTokens(sdk, sdk.sdkConfiguration, sdk.hooks)
	sdk.QRCodes = newQRCodes(sdk, sdk.sdkConfiguration, sdk.hooks)

	return sdk
}
