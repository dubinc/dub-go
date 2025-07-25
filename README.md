# github.com/dubinc/dub-go

<div align="left">
    <a href="https://speakeasyapi.dev/"><img src="https://custom-icon-badges.demolab.com/badge/-Built%20By%20Speakeasy-212015?style=for-the-badge&logoColor=FBE331&logo=speakeasy&labelColor=545454" /></a>
    <a href="https://opensource.org/licenses/MIT">
        <img src="https://img.shields.io/badge/License-MIT-blue.svg" style="width: 100px; height: 28px;" />
    </a>
</div>

<!-- Start Summary [summary] -->
## Summary

Dub API: Dub is link management infrastructure for companies to create marketing campaigns, link sharing features, and referral programs.
<!-- End Summary [summary] -->

<!-- Start Table of Contents [toc] -->
## Table of Contents
<!-- $toc-max-depth=2 -->
* [github.com/dubinc/dub-go](#githubcomdubincdub-go)
  * [SDK Installation](#sdk-installation)
  * [SDK Example Usage](#sdk-example-usage)
  * [Available Resources and Operations](#available-resources-and-operations)
  * [Error Handling](#error-handling)
  * [Server Selection](#server-selection)
  * [Custom HTTP Client](#custom-http-client)
  * [Authentication](#authentication)
  * [Retries](#retries)
  * [Pagination](#pagination)
* [Development](#development)
  * [Contributions](#contributions)

<!-- End Table of Contents [toc] -->

<!-- Start SDK Installation [installation] -->
## SDK Installation

To add the SDK as a dependency to your project:
```bash
go get github.com/dubinc/dub-go
```
<!-- End SDK Installation [installation] -->

<!-- Start SDK Example Usage [usage] -->
## SDK Example Usage

### Example 1

```go
package main

import (
	"context"
	dubgo "github.com/dubinc/dub-go"
	"github.com/dubinc/dub-go/models/operations"
	"log"
)

func main() {
	ctx := context.Background()

	s := dubgo.New(
		dubgo.WithSecurity("DUB_API_KEY"),
	)

	res, err := s.Links.Create(ctx, &operations.CreateLinkRequestBody{
		URL:        "https://google.com",
		ExternalID: dubgo.String("123456"),
		TagIds: dubgo.Pointer(operations.CreateTagIdsArrayOfStr(
			[]string{
				"clux0rgak00011...",
			},
		)),
		TestVariants: []operations.TestVariants{
			operations.TestVariants{
				URL:        "https://example.com/variant-1",
				Percentage: 50,
			},
			operations.TestVariants{
				URL:        "https://example.com/variant-2",
				Percentage: 50,
			},
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	if res != nil {
		// handle response
	}
}

```

### Example 2

```go
package main

import (
	"context"
	dubgo "github.com/dubinc/dub-go"
	"github.com/dubinc/dub-go/models/operations"
	"log"
)

func main() {
	ctx := context.Background()

	s := dubgo.New(
		dubgo.WithSecurity("DUB_API_KEY"),
	)

	res, err := s.Links.Upsert(ctx, &operations.UpsertLinkRequestBody{
		URL:        "https://google.com",
		ExternalID: dubgo.String("123456"),
		TagIds: dubgo.Pointer(operations.CreateUpsertLinkTagIdsArrayOfStr(
			[]string{
				"clux0rgak00011...",
			},
		)),
		TestVariants: []operations.UpsertLinkTestVariants{
			operations.UpsertLinkTestVariants{
				URL:        "https://example.com/variant-1",
				Percentage: 50,
			},
			operations.UpsertLinkTestVariants{
				URL:        "https://example.com/variant-2",
				Percentage: 50,
			},
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	if res != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->

<!-- Start Available Resources and Operations [operations] -->
## Available Resources and Operations

<details open>
<summary>Available methods</summary>

### [Analytics](docs/sdks/analytics/README.md)

* [Retrieve](docs/sdks/analytics/README.md#retrieve) - Retrieve analytics for a link, a domain, or the authenticated workspace.

### [Commissions](docs/sdks/commissions/README.md)

* [List](docs/sdks/commissions/README.md#list) - Get commissions for a program.
* [Update](docs/sdks/commissions/README.md#update) - Update a commission.

### [Customers](docs/sdks/customers/README.md)

* [List](docs/sdks/customers/README.md#list) - Retrieve a list of customers
* [~~Create~~](docs/sdks/customers/README.md#create) - Create a customer :warning: **Deprecated**
* [Get](docs/sdks/customers/README.md#get) - Retrieve a customer
* [Update](docs/sdks/customers/README.md#update) - Update a customer
* [Delete](docs/sdks/customers/README.md#delete) - Delete a customer

### [Domains](docs/sdks/domains/README.md)

* [Create](docs/sdks/domains/README.md#create) - Create a domain
* [List](docs/sdks/domains/README.md#list) - Retrieve a list of domains
* [Update](docs/sdks/domains/README.md#update) - Update a domain
* [Delete](docs/sdks/domains/README.md#delete) - Delete a domain
* [Register](docs/sdks/domains/README.md#register) - Register a domain
* [CheckStatus](docs/sdks/domains/README.md#checkstatus) - Check the availability of one or more domains


### [EmbedTokens](docs/sdks/embedtokens/README.md)

* [Referrals](docs/sdks/embedtokens/README.md#referrals) - Create a referrals embed token

### [Events](docs/sdks/events/README.md)

* [List](docs/sdks/events/README.md#list) - Retrieve a list of events

### [Folders](docs/sdks/folders/README.md)

* [Create](docs/sdks/folders/README.md#create) - Create a folder
* [List](docs/sdks/folders/README.md#list) - Retrieve a list of folders
* [Update](docs/sdks/folders/README.md#update) - Update a folder
* [Delete](docs/sdks/folders/README.md#delete) - Delete a folder

### [Links](docs/sdks/links/README.md)

* [Create](docs/sdks/links/README.md#create) - Create a link
* [List](docs/sdks/links/README.md#list) - Retrieve a list of links
* [Count](docs/sdks/links/README.md#count) - Retrieve links count
* [Get](docs/sdks/links/README.md#get) - Retrieve a link
* [Update](docs/sdks/links/README.md#update) - Update a link
* [Delete](docs/sdks/links/README.md#delete) - Delete a link
* [CreateMany](docs/sdks/links/README.md#createmany) - Bulk create links
* [UpdateMany](docs/sdks/links/README.md#updatemany) - Bulk update links
* [DeleteMany](docs/sdks/links/README.md#deletemany) - Bulk delete links
* [Upsert](docs/sdks/links/README.md#upsert) - Upsert a link

### [Partners](docs/sdks/partners/README.md)

* [Create](docs/sdks/partners/README.md#create) - Create a partner
* [List](docs/sdks/partners/README.md#list) - List all partners
* [CreateLink](docs/sdks/partners/README.md#createlink) - Create a link for a partner
* [RetrieveLinks](docs/sdks/partners/README.md#retrievelinks) - Retrieve a partner's links.
* [UpsertLink](docs/sdks/partners/README.md#upsertlink) - Upsert a link for a partner
* [Analytics](docs/sdks/partners/README.md#analytics) - Retrieve analytics for a partner

### [QRCodes](docs/sdks/qrcodes/README.md)

* [Get](docs/sdks/qrcodes/README.md#get) - Retrieve a QR code

### [Tags](docs/sdks/tags/README.md)

* [Create](docs/sdks/tags/README.md#create) - Create a tag
* [List](docs/sdks/tags/README.md#list) - Retrieve a list of tags
* [Update](docs/sdks/tags/README.md#update) - Update a tag
* [Delete](docs/sdks/tags/README.md#delete) - Delete a tag

### [Track](docs/sdks/track/README.md)

* [Lead](docs/sdks/track/README.md#lead) - Track a lead
* [Sale](docs/sdks/track/README.md#sale) - Track a sale

### [Workspaces](docs/sdks/workspaces/README.md)

* [Get](docs/sdks/workspaces/README.md#get) - Retrieve a workspace
* [Update](docs/sdks/workspaces/README.md#update) - Update a workspace

</details>
<!-- End Available Resources and Operations [operations] -->

<!-- Start Error Handling [errors] -->
## Error Handling

Handling errors in this SDK should largely match your expectations. All operations return a response object or an error, they will never return both.

By Default, an API error will return `sdkerrors.SDKError`. When custom error responses are specified for an operation, the SDK may also return their associated error. You can refer to respective *Errors* tables in SDK docs for more details on possible error types for each operation.

For example, the `Create` function may return the following errors:

| Error Type                    | Status Code | Content Type     |
| ----------------------------- | ----------- | ---------------- |
| sdkerrors.BadRequest          | 400         | application/json |
| sdkerrors.Unauthorized        | 401         | application/json |
| sdkerrors.Forbidden           | 403         | application/json |
| sdkerrors.NotFound            | 404         | application/json |
| sdkerrors.Conflict            | 409         | application/json |
| sdkerrors.InviteExpired       | 410         | application/json |
| sdkerrors.UnprocessableEntity | 422         | application/json |
| sdkerrors.RateLimitExceeded   | 429         | application/json |
| sdkerrors.InternalServerError | 500         | application/json |
| sdkerrors.SDKError            | 4XX, 5XX    | \*/\*            |

### Example

```go
package main

import (
	"context"
	"errors"
	dubgo "github.com/dubinc/dub-go"
	"github.com/dubinc/dub-go/models/operations"
	"github.com/dubinc/dub-go/models/sdkerrors"
	"log"
)

func main() {
	ctx := context.Background()

	s := dubgo.New(
		dubgo.WithSecurity("DUB_API_KEY"),
	)

	res, err := s.Links.Create(ctx, &operations.CreateLinkRequestBody{
		URL:        "https://google.com",
		ExternalID: dubgo.String("123456"),
		TagIds: dubgo.Pointer(operations.CreateTagIdsArrayOfStr(
			[]string{
				"clux0rgak00011...",
			},
		)),
		TestVariants: []operations.TestVariants{
			operations.TestVariants{
				URL:        "https://example.com/variant-1",
				Percentage: 50,
			},
			operations.TestVariants{
				URL:        "https://example.com/variant-2",
				Percentage: 50,
			},
		},
	})
	if err != nil {

		var e *sdkerrors.BadRequest
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.Unauthorized
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.Forbidden
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.NotFound
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.Conflict
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.InviteExpired
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.UnprocessableEntity
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.RateLimitExceeded
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.InternalServerError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.SDKError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}
	}
}

```
<!-- End Error Handling [errors] -->

<!-- Start Server Selection [server] -->
## Server Selection

### Override Server URL Per-Client

The default server can be overridden globally using the `WithServerURL(serverURL string)` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	dubgo "github.com/dubinc/dub-go"
	"github.com/dubinc/dub-go/models/operations"
	"log"
)

func main() {
	ctx := context.Background()

	s := dubgo.New(
		dubgo.WithServerURL("https://api.dub.co"),
		dubgo.WithSecurity("DUB_API_KEY"),
	)

	res, err := s.Links.Create(ctx, &operations.CreateLinkRequestBody{
		URL:        "https://google.com",
		ExternalID: dubgo.String("123456"),
		TagIds: dubgo.Pointer(operations.CreateTagIdsArrayOfStr(
			[]string{
				"clux0rgak00011...",
			},
		)),
		TestVariants: []operations.TestVariants{
			operations.TestVariants{
				URL:        "https://example.com/variant-1",
				Percentage: 50,
			},
			operations.TestVariants{
				URL:        "https://example.com/variant-2",
				Percentage: 50,
			},
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	if res != nil {
		// handle response
	}
}

```
<!-- End Server Selection [server] -->

<!-- Start Custom HTTP Client [http-client] -->
## Custom HTTP Client

The Go SDK makes API calls that wrap an internal HTTP client. The requirements for the HTTP client are very simple. It must match this interface:

```go
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
```

The built-in `net/http` client satisfies this interface and a default client based on the built-in is provided by default. To replace this default with a client of your own, you can implement this interface yourself or provide your own client configured as desired. Here's a simple example, which adds a client with a 30 second timeout.

```go
import (
	"net/http"
	"time"
	"github.com/myorg/your-go-sdk"
)

var (
	httpClient = &http.Client{Timeout: 30 * time.Second}
	sdkClient  = sdk.New(sdk.WithClient(httpClient))
)
```

This can be a convenient way to configure timeouts, cookies, proxies, custom headers, and other low-level configuration.
<!-- End Custom HTTP Client [http-client] -->

<!-- Start Authentication [security] -->
## Authentication

### Per-Client Security Schemes

This SDK supports the following security scheme globally:

| Name    | Type | Scheme      |
| ------- | ---- | ----------- |
| `Token` | http | HTTP Bearer |

You can configure it using the `WithSecurity` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	dubgo "github.com/dubinc/dub-go"
	"github.com/dubinc/dub-go/models/operations"
	"log"
)

func main() {
	ctx := context.Background()

	s := dubgo.New(
		dubgo.WithSecurity("DUB_API_KEY"),
	)

	res, err := s.Links.Create(ctx, &operations.CreateLinkRequestBody{
		URL:        "https://google.com",
		ExternalID: dubgo.String("123456"),
		TagIds: dubgo.Pointer(operations.CreateTagIdsArrayOfStr(
			[]string{
				"clux0rgak00011...",
			},
		)),
		TestVariants: []operations.TestVariants{
			operations.TestVariants{
				URL:        "https://example.com/variant-1",
				Percentage: 50,
			},
			operations.TestVariants{
				URL:        "https://example.com/variant-2",
				Percentage: 50,
			},
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	if res != nil {
		// handle response
	}
}

```
<!-- End Authentication [security] -->

<!-- Start Retries [retries] -->
## Retries

Some of the endpoints in this SDK support retries. If you use the SDK without any configuration, it will fall back to the default retry strategy provided by the API. However, the default retry strategy can be overridden on a per-operation basis, or across the entire SDK.

To change the default retry strategy for a single API call, simply provide a `retry.Config` object to the call by using the `WithRetries` option:
```go
package main

import (
	"context"
	dubgo "github.com/dubinc/dub-go"
	"github.com/dubinc/dub-go/models/operations"
	"github.com/dubinc/dub-go/retry"
	"log"
	"models/operations"
)

func main() {
	ctx := context.Background()

	s := dubgo.New(
		dubgo.WithSecurity("DUB_API_KEY"),
	)

	res, err := s.Links.Create(ctx, &operations.CreateLinkRequestBody{
		URL:        "https://google.com",
		ExternalID: dubgo.String("123456"),
		TagIds: dubgo.Pointer(operations.CreateTagIdsArrayOfStr(
			[]string{
				"clux0rgak00011...",
			},
		)),
		TestVariants: []operations.TestVariants{
			operations.TestVariants{
				URL:        "https://example.com/variant-1",
				Percentage: 50,
			},
			operations.TestVariants{
				URL:        "https://example.com/variant-2",
				Percentage: 50,
			},
		},
	}, operations.WithRetries(
		retry.Config{
			Strategy: "backoff",
			Backoff: &retry.BackoffStrategy{
				InitialInterval: 1,
				MaxInterval:     50,
				Exponent:        1.1,
				MaxElapsedTime:  100,
			},
			RetryConnectionErrors: false,
		}))
	if err != nil {
		log.Fatal(err)
	}
	if res != nil {
		// handle response
	}
}

```

If you'd like to override the default retry strategy for all operations that support retries, you can use the `WithRetryConfig` option at SDK initialization:
```go
package main

import (
	"context"
	dubgo "github.com/dubinc/dub-go"
	"github.com/dubinc/dub-go/models/operations"
	"github.com/dubinc/dub-go/retry"
	"log"
)

func main() {
	ctx := context.Background()

	s := dubgo.New(
		dubgo.WithRetryConfig(
			retry.Config{
				Strategy: "backoff",
				Backoff: &retry.BackoffStrategy{
					InitialInterval: 1,
					MaxInterval:     50,
					Exponent:        1.1,
					MaxElapsedTime:  100,
				},
				RetryConnectionErrors: false,
			}),
		dubgo.WithSecurity("DUB_API_KEY"),
	)

	res, err := s.Links.Create(ctx, &operations.CreateLinkRequestBody{
		URL:        "https://google.com",
		ExternalID: dubgo.String("123456"),
		TagIds: dubgo.Pointer(operations.CreateTagIdsArrayOfStr(
			[]string{
				"clux0rgak00011...",
			},
		)),
		TestVariants: []operations.TestVariants{
			operations.TestVariants{
				URL:        "https://example.com/variant-1",
				Percentage: 50,
			},
			operations.TestVariants{
				URL:        "https://example.com/variant-2",
				Percentage: 50,
			},
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	if res != nil {
		// handle response
	}
}

```
<!-- End Retries [retries] -->

<!-- Start Pagination [pagination] -->
## Pagination

Some of the endpoints in this SDK support pagination. To use pagination, you make your SDK calls as usual, but the
returned response object will have a `Next` method that can be called to pull down the next group of results. If the
return value of `Next` is `nil`, then there are no more pages to be fetched.

Here's an example of one such pagination call:
```go
package main

import (
	"context"
	dubgo "github.com/dubinc/dub-go"
	"github.com/dubinc/dub-go/models/operations"
	"log"
)

func main() {
	ctx := context.Background()

	s := dubgo.New(
		dubgo.WithSecurity("DUB_API_KEY"),
	)

	res, err := s.Links.List(ctx, operations.GetLinksRequest{
		PageSize: dubgo.Float64(50),
	})
	if err != nil {
		log.Fatal(err)
	}
	if res != nil {
		for {
			// handle items

			res, err = res.Next()

			if err != nil {
				// handle error
			}

			if res == nil {
				break
			}
		}
	}
}

```
<!-- End Pagination [pagination] -->

<!-- Placeholder for Future Speakeasy SDK Sections -->

# Development

## Contributions

While we value open-source contributions to this SDK, this library is generated programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release!

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
