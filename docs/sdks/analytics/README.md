# Analytics
(*Analytics*)

### Available Operations

* [~~Timeseries~~](#timeseries) - Retrieve timeseries click analytics :warning: **Deprecated** Use `Timeseries` instead.
* [~~Country~~](#country) - Retrieve top countries by clicks :warning: **Deprecated** Use `Countries` instead.
* [~~City~~](#city) - Retrieve top cities by clicks :warning: **Deprecated** Use `Cities` instead.
* [~~Device~~](#device) - Retrieve top devices by clicks :warning: **Deprecated** Use `Devices` instead.
* [~~Browser~~](#browser) - Retrieve top browsers by clicks :warning: **Deprecated** Use `Browsers` instead.
* [~~Os~~](#os) - Retrieve top OS by clicks :warning: **Deprecated** Use `Os` instead.
* [~~Referer~~](#referer) - Retrieve top referers by clicks :warning: **Deprecated** Use `Referers` instead.
* [~~TopLinks~~](#toplinks) - Retrieve top links by clicks :warning: **Deprecated** Use `TopLinks` instead.
* [~~TopUrls~~](#topurls) - Retrieve top URLs by clicks :warning: **Deprecated** Use `TopUrls` instead.

## ~~Timeseries~~

Retrieve timeseries click analytics for a link, a domain, or the authenticated workspace over a period of time.

> :warning: **DEPRECATED**: This method is deprecated. Use dub.analytics.clicks.timeseries instead.. Use `Timeseries` instead.

### Example Usage

```go
package main

import(
	"github.com/dubinc/dub-go/models/components"
	dubgo "github.com/dubinc/dub-go"
	"context"
	"github.com/dubinc/dub-go/models/operations"
	"log"
)

func main() {
    s := dubgo.New(
        dubgo.WithSecurity("DUB_API_KEY"),
        dubgo.WithWorkspaceID("<value>"),
    )

    ctx := context.Background()
    res, err := s.Analytics.Timeseries(ctx, operations.GetTimeseriesByClicksDeprecatedRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.ResponseBodies != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                              | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                  | :heavy_check_mark:                                                                                                     | The context to use for the request.                                                                                    |
| `request`                                                                                                              | [operations.GetTimeseriesByClicksDeprecatedRequest](../../models/operations/gettimeseriesbyclicksdeprecatedrequest.md) | :heavy_check_mark:                                                                                                     | The request object to use for the request.                                                                             |


### Response

**[*operations.GetTimeseriesByClicksDeprecatedResponse](../../models/operations/gettimeseriesbyclicksdeprecatedresponse.md), error**
| Error Object                  | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.BadRequest          | 400                           | application/json              |
| sdkerrors.Unauthorized        | 401                           | application/json              |
| sdkerrors.Forbidden           | 403                           | application/json              |
| sdkerrors.NotFound            | 404                           | application/json              |
| sdkerrors.Conflict            | 409                           | application/json              |
| sdkerrors.InviteExpired       | 410                           | application/json              |
| sdkerrors.UnprocessableEntity | 422                           | application/json              |
| sdkerrors.RateLimitExceeded   | 429                           | application/json              |
| sdkerrors.InternalServerError | 500                           | application/json              |
| sdkerrors.SDKError            | 4xx-5xx                       | */*                           |

## ~~Country~~

Retrieve the top countries by number of clicks for a link, a domain, or the authenticated workspace.

> :warning: **DEPRECATED**: This method is deprecated. Use dub.analytics.clicks.countries instead.. Use `Countries` instead.

### Example Usage

```go
package main

import(
	"github.com/dubinc/dub-go/models/components"
	dubgo "github.com/dubinc/dub-go"
	"context"
	"github.com/dubinc/dub-go/models/operations"
	"log"
)

func main() {
    s := dubgo.New(
        dubgo.WithSecurity("DUB_API_KEY"),
        dubgo.WithWorkspaceID("<value>"),
    )

    ctx := context.Background()
    res, err := s.Analytics.Country(ctx, operations.GetCountriesByClicksDeprecatedRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.ClicksByCountries != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [operations.GetCountriesByClicksDeprecatedRequest](../../models/operations/getcountriesbyclicksdeprecatedrequest.md) | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |


### Response

**[*operations.GetCountriesByClicksDeprecatedResponse](../../models/operations/getcountriesbyclicksdeprecatedresponse.md), error**
| Error Object                  | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.BadRequest          | 400                           | application/json              |
| sdkerrors.Unauthorized        | 401                           | application/json              |
| sdkerrors.Forbidden           | 403                           | application/json              |
| sdkerrors.NotFound            | 404                           | application/json              |
| sdkerrors.Conflict            | 409                           | application/json              |
| sdkerrors.InviteExpired       | 410                           | application/json              |
| sdkerrors.UnprocessableEntity | 422                           | application/json              |
| sdkerrors.RateLimitExceeded   | 429                           | application/json              |
| sdkerrors.InternalServerError | 500                           | application/json              |
| sdkerrors.SDKError            | 4xx-5xx                       | */*                           |

## ~~City~~

Retrieve the top countries by number of clicks for a link, a domain, or the authenticated workspace.

> :warning: **DEPRECATED**: This method is deprecated. Use dub.analytics.clicks.cities instead.. Use `Cities` instead.

### Example Usage

```go
package main

import(
	"github.com/dubinc/dub-go/models/components"
	dubgo "github.com/dubinc/dub-go"
	"context"
	"github.com/dubinc/dub-go/models/operations"
	"log"
)

func main() {
    s := dubgo.New(
        dubgo.WithSecurity("DUB_API_KEY"),
        dubgo.WithWorkspaceID("<value>"),
    )

    ctx := context.Background()
    res, err := s.Analytics.City(ctx, operations.GetCitiesByClicksDeprecatedRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.ClicksByCities != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.GetCitiesByClicksDeprecatedRequest](../../models/operations/getcitiesbyclicksdeprecatedrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |


### Response

**[*operations.GetCitiesByClicksDeprecatedResponse](../../models/operations/getcitiesbyclicksdeprecatedresponse.md), error**
| Error Object                  | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.BadRequest          | 400                           | application/json              |
| sdkerrors.Unauthorized        | 401                           | application/json              |
| sdkerrors.Forbidden           | 403                           | application/json              |
| sdkerrors.NotFound            | 404                           | application/json              |
| sdkerrors.Conflict            | 409                           | application/json              |
| sdkerrors.InviteExpired       | 410                           | application/json              |
| sdkerrors.UnprocessableEntity | 422                           | application/json              |
| sdkerrors.RateLimitExceeded   | 429                           | application/json              |
| sdkerrors.InternalServerError | 500                           | application/json              |
| sdkerrors.SDKError            | 4xx-5xx                       | */*                           |

## ~~Device~~

Retrieve the top devices by number of clicks for a link, a domain, or the authenticated workspace.

> :warning: **DEPRECATED**: This method is deprecated. Use dub.analytics.clicks.devices instead.. Use `Devices` instead.

### Example Usage

```go
package main

import(
	"github.com/dubinc/dub-go/models/components"
	dubgo "github.com/dubinc/dub-go"
	"context"
	"github.com/dubinc/dub-go/models/operations"
	"log"
)

func main() {
    s := dubgo.New(
        dubgo.WithSecurity("DUB_API_KEY"),
        dubgo.WithWorkspaceID("<value>"),
    )

    ctx := context.Background()
    res, err := s.Analytics.Device(ctx, operations.GetDevicesByClicksDeprecatedRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.ResponseBodies != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.GetDevicesByClicksDeprecatedRequest](../../models/operations/getdevicesbyclicksdeprecatedrequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |


### Response

**[*operations.GetDevicesByClicksDeprecatedResponse](../../models/operations/getdevicesbyclicksdeprecatedresponse.md), error**
| Error Object                  | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.BadRequest          | 400                           | application/json              |
| sdkerrors.Unauthorized        | 401                           | application/json              |
| sdkerrors.Forbidden           | 403                           | application/json              |
| sdkerrors.NotFound            | 404                           | application/json              |
| sdkerrors.Conflict            | 409                           | application/json              |
| sdkerrors.InviteExpired       | 410                           | application/json              |
| sdkerrors.UnprocessableEntity | 422                           | application/json              |
| sdkerrors.RateLimitExceeded   | 429                           | application/json              |
| sdkerrors.InternalServerError | 500                           | application/json              |
| sdkerrors.SDKError            | 4xx-5xx                       | */*                           |

## ~~Browser~~

Retrieve the top browsers by number of clicks for a link, a domain, or the authenticated workspace.

> :warning: **DEPRECATED**: This method is deprecated. Use dub.analytics.clicks.browsers instead.. Use `Browsers` instead.

### Example Usage

```go
package main

import(
	"github.com/dubinc/dub-go/models/components"
	dubgo "github.com/dubinc/dub-go"
	"context"
	"github.com/dubinc/dub-go/models/operations"
	"log"
)

func main() {
    s := dubgo.New(
        dubgo.WithSecurity("DUB_API_KEY"),
        dubgo.WithWorkspaceID("<value>"),
    )

    ctx := context.Background()
    res, err := s.Analytics.Browser(ctx, operations.GetBrowsersByClicksDeprecatedRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.ResponseBodies != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.GetBrowsersByClicksDeprecatedRequest](../../models/operations/getbrowsersbyclicksdeprecatedrequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |


### Response

**[*operations.GetBrowsersByClicksDeprecatedResponse](../../models/operations/getbrowsersbyclicksdeprecatedresponse.md), error**
| Error Object                  | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.BadRequest          | 400                           | application/json              |
| sdkerrors.Unauthorized        | 401                           | application/json              |
| sdkerrors.Forbidden           | 403                           | application/json              |
| sdkerrors.NotFound            | 404                           | application/json              |
| sdkerrors.Conflict            | 409                           | application/json              |
| sdkerrors.InviteExpired       | 410                           | application/json              |
| sdkerrors.UnprocessableEntity | 422                           | application/json              |
| sdkerrors.RateLimitExceeded   | 429                           | application/json              |
| sdkerrors.InternalServerError | 500                           | application/json              |
| sdkerrors.SDKError            | 4xx-5xx                       | */*                           |

## ~~Os~~

Retrieve the top OS by number of clicks for a link, a domain, or the authenticated workspace.

> :warning: **DEPRECATED**: This method is deprecated. Use dub.analytics.clicks.os instead.. Use `Os` instead.

### Example Usage

```go
package main

import(
	"github.com/dubinc/dub-go/models/components"
	dubgo "github.com/dubinc/dub-go"
	"context"
	"github.com/dubinc/dub-go/models/operations"
	"log"
)

func main() {
    s := dubgo.New(
        dubgo.WithSecurity("DUB_API_KEY"),
        dubgo.WithWorkspaceID("<value>"),
    )

    ctx := context.Background()
    res, err := s.Analytics.Os(ctx, operations.GetOSByClicksDeprecatedRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.ResponseBodies != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.GetOSByClicksDeprecatedRequest](../../models/operations/getosbyclicksdeprecatedrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |


### Response

**[*operations.GetOSByClicksDeprecatedResponse](../../models/operations/getosbyclicksdeprecatedresponse.md), error**
| Error Object                  | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.BadRequest          | 400                           | application/json              |
| sdkerrors.Unauthorized        | 401                           | application/json              |
| sdkerrors.Forbidden           | 403                           | application/json              |
| sdkerrors.NotFound            | 404                           | application/json              |
| sdkerrors.Conflict            | 409                           | application/json              |
| sdkerrors.InviteExpired       | 410                           | application/json              |
| sdkerrors.UnprocessableEntity | 422                           | application/json              |
| sdkerrors.RateLimitExceeded   | 429                           | application/json              |
| sdkerrors.InternalServerError | 500                           | application/json              |
| sdkerrors.SDKError            | 4xx-5xx                       | */*                           |

## ~~Referer~~

Retrieve the top referers by number of clicks for a link, a domain, or the authenticated workspace.

> :warning: **DEPRECATED**: This method is deprecated. Use dub.analytics.clicks.referers instead.. Use `Referers` instead.

### Example Usage

```go
package main

import(
	"github.com/dubinc/dub-go/models/components"
	dubgo "github.com/dubinc/dub-go"
	"context"
	"github.com/dubinc/dub-go/models/operations"
	"log"
)

func main() {
    s := dubgo.New(
        dubgo.WithSecurity("DUB_API_KEY"),
        dubgo.WithWorkspaceID("<value>"),
    )

    ctx := context.Background()
    res, err := s.Analytics.Referer(ctx, operations.GetReferersByClicksDeprecatedRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.ResponseBodies != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.GetReferersByClicksDeprecatedRequest](../../models/operations/getreferersbyclicksdeprecatedrequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |


### Response

**[*operations.GetReferersByClicksDeprecatedResponse](../../models/operations/getreferersbyclicksdeprecatedresponse.md), error**
| Error Object                  | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.BadRequest          | 400                           | application/json              |
| sdkerrors.Unauthorized        | 401                           | application/json              |
| sdkerrors.Forbidden           | 403                           | application/json              |
| sdkerrors.NotFound            | 404                           | application/json              |
| sdkerrors.Conflict            | 409                           | application/json              |
| sdkerrors.InviteExpired       | 410                           | application/json              |
| sdkerrors.UnprocessableEntity | 422                           | application/json              |
| sdkerrors.RateLimitExceeded   | 429                           | application/json              |
| sdkerrors.InternalServerError | 500                           | application/json              |
| sdkerrors.SDKError            | 4xx-5xx                       | */*                           |

## ~~TopLinks~~

Retrieve the top links by number of clicks for a domain or the authenticated workspace.

> :warning: **DEPRECATED**: This method is deprecated. Use dub.analytics.clicks.topLinks instead.. Use `TopLinks` instead.

### Example Usage

```go
package main

import(
	"github.com/dubinc/dub-go/models/components"
	dubgo "github.com/dubinc/dub-go"
	"context"
	"github.com/dubinc/dub-go/models/operations"
	"log"
)

func main() {
    s := dubgo.New(
        dubgo.WithSecurity("DUB_API_KEY"),
        dubgo.WithWorkspaceID("<value>"),
    )

    ctx := context.Background()
    res, err := s.Analytics.TopLinks(ctx, operations.GetTopLinksByClicksDeprecatedRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.ResponseBodies != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.GetTopLinksByClicksDeprecatedRequest](../../models/operations/gettoplinksbyclicksdeprecatedrequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |


### Response

**[*operations.GetTopLinksByClicksDeprecatedResponse](../../models/operations/gettoplinksbyclicksdeprecatedresponse.md), error**
| Error Object                  | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.BadRequest          | 400                           | application/json              |
| sdkerrors.Unauthorized        | 401                           | application/json              |
| sdkerrors.Forbidden           | 403                           | application/json              |
| sdkerrors.NotFound            | 404                           | application/json              |
| sdkerrors.Conflict            | 409                           | application/json              |
| sdkerrors.InviteExpired       | 410                           | application/json              |
| sdkerrors.UnprocessableEntity | 422                           | application/json              |
| sdkerrors.RateLimitExceeded   | 429                           | application/json              |
| sdkerrors.InternalServerError | 500                           | application/json              |
| sdkerrors.SDKError            | 4xx-5xx                       | */*                           |

## ~~TopUrls~~

Retrieve the top URLs by number of clicks for a given short link.

> :warning: **DEPRECATED**: This method is deprecated. Use dub.analytics.clicks.topUrls instead.. Use `TopUrls` instead.

### Example Usage

```go
package main

import(
	"github.com/dubinc/dub-go/models/components"
	dubgo "github.com/dubinc/dub-go"
	"context"
	"github.com/dubinc/dub-go/models/operations"
	"log"
)

func main() {
    s := dubgo.New(
        dubgo.WithSecurity("DUB_API_KEY"),
        dubgo.WithWorkspaceID("<value>"),
    )

    ctx := context.Background()
    res, err := s.Analytics.TopUrls(ctx, operations.GetTopURLsByClicksDeprecatedRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.ResponseBodies != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.GetTopURLsByClicksDeprecatedRequest](../../models/operations/gettopurlsbyclicksdeprecatedrequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |


### Response

**[*operations.GetTopURLsByClicksDeprecatedResponse](../../models/operations/gettopurlsbyclicksdeprecatedresponse.md), error**
| Error Object                  | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.BadRequest          | 400                           | application/json              |
| sdkerrors.Unauthorized        | 401                           | application/json              |
| sdkerrors.Forbidden           | 403                           | application/json              |
| sdkerrors.NotFound            | 404                           | application/json              |
| sdkerrors.Conflict            | 409                           | application/json              |
| sdkerrors.InviteExpired       | 410                           | application/json              |
| sdkerrors.UnprocessableEntity | 422                           | application/json              |
| sdkerrors.RateLimitExceeded   | 429                           | application/json              |
| sdkerrors.InternalServerError | 500                           | application/json              |
| sdkerrors.SDKError            | 4xx-5xx                       | */*                           |
