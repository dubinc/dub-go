# Links
(*Links*)

## Overview

### Available Operations

* [List](#list) - Retrieve a list of links
* [Create](#create) - Create a new link
* [Count](#count) - Retrieve links count
* [Get](#get) - Retrieve a link
* [Delete](#delete) - Delete a link
* [Update](#update) - Update a link
* [CreateMany](#createmany) - Bulk create links
* [DeleteMany](#deletemany) - Bulk delete links
* [UpdateMany](#updatemany) - Bulk update links
* [Upsert](#upsert) - Upsert a link

## List

Retrieve a paginated list of links for the authenticated workspace.

### Example Usage

```go
package main

import(
	dubgo "github.com/dubinc/dub-go"
	"context"
	"github.com/dubinc/dub-go/models/operations"
	"log"
)

func main() {
    s := dubgo.New(
        dubgo.WithSecurity("DUB_API_KEY"),
    )

    ctx := context.Background()
    res, err := s.Links.List(ctx, operations.GetLinksRequest{
        Page: dubgo.Float64(1),
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

### Parameters

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `request`                                                                | [operations.GetLinksRequest](../../models/operations/getlinksrequest.md) | :heavy_check_mark:                                                       | The request object to use for the request.                               |
| `opts`                                                                   | [][operations.Option](../../models/operations/option.md)                 | :heavy_minus_sign:                                                       | The options for this request.                                            |

### Response

**[*operations.GetLinksResponse](../../models/operations/getlinksresponse.md), error**

### Errors

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


## Create

Create a new link for the authenticated workspace.

### Example Usage

```go
package main

import(
	dubgo "github.com/dubinc/dub-go"
	"context"
	"github.com/dubinc/dub-go/models/operations"
	"log"
)

func main() {
    s := dubgo.New(
        dubgo.WithSecurity("DUB_API_KEY"),
    )

    ctx := context.Background()
    res, err := s.Links.Create(ctx, &operations.CreateLinkRequestBody{
        URL: "https://google.com",
        ExternalID: dubgo.String("123456"),
        TagIds: dubgo.Pointer(operations.CreateTagIdsArrayOfStr(
            []string{
                "clux0rgak00011...",
            },
        )),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.CreateLinkRequestBody](../../models/operations/createlinkrequestbody.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*components.LinkSchema](../../models/components/linkschema.md), error**

### Errors

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


## Count

Retrieve the number of links for the authenticated workspace.

### Example Usage

```go
package main

import(
	dubgo "github.com/dubinc/dub-go"
	"context"
	"github.com/dubinc/dub-go/models/operations"
	"log"
)

func main() {
    s := dubgo.New(
        dubgo.WithSecurity("DUB_API_KEY"),
    )

    ctx := context.Background()
    res, err := s.Links.Count(ctx, operations.GetLinksCountRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.GetLinksCountRequest](../../models/operations/getlinkscountrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*float64](../../.md), error**

### Errors

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


## Get

Retrieve the info for a link.

### Example Usage

```go
package main

import(
	dubgo "github.com/dubinc/dub-go"
	"context"
	"github.com/dubinc/dub-go/models/operations"
	"log"
)

func main() {
    s := dubgo.New(
        dubgo.WithSecurity("DUB_API_KEY"),
    )

    ctx := context.Background()
    res, err := s.Links.Get(ctx, operations.GetLinkInfoRequest{
        LinkID: dubgo.String("clux0rgak00011..."),
        ExternalID: dubgo.String("ext_123456"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.GetLinkInfoRequest](../../models/operations/getlinkinforequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*components.LinkSchema](../../models/components/linkschema.md), error**

### Errors

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


## Delete

Delete a link for the authenticated workspace.

### Example Usage

```go
package main

import(
	dubgo "github.com/dubinc/dub-go"
	"context"
	"log"
)

func main() {
    s := dubgo.New(
        dubgo.WithSecurity("DUB_API_KEY"),
    )

    ctx := context.Background()
    res, err := s.Links.Delete(ctx, "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                             | Type                                                                                                                                  | Required                                                                                                                              | Description                                                                                                                           |
| ------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                 | [context.Context](https://pkg.go.dev/context#Context)                                                                                 | :heavy_check_mark:                                                                                                                    | The context to use for the request.                                                                                                   |
| `linkID`                                                                                                                              | *string*                                                                                                                              | :heavy_check_mark:                                                                                                                    | The id of the link to delete. You may use either `linkId` (obtained via `/links/info` endpoint) or `externalId` prefixed with `ext_`. |
| `opts`                                                                                                                                | [][operations.Option](../../models/operations/option.md)                                                                              | :heavy_minus_sign:                                                                                                                    | The options for this request.                                                                                                         |

### Response

**[*operations.DeleteLinkResponseBody](../../models/operations/deletelinkresponsebody.md), error**

### Errors

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


## Update

Update a link for the authenticated workspace. If there's no change, returns it as it is.

### Example Usage

```go
package main

import(
	dubgo "github.com/dubinc/dub-go"
	"context"
	"github.com/dubinc/dub-go/models/operations"
	"log"
)

func main() {
    s := dubgo.New(
        dubgo.WithSecurity("DUB_API_KEY"),
    )

    ctx := context.Background()
    res, err := s.Links.Update(ctx, "<value>", &operations.UpdateLinkRequestBody{
        URL: dubgo.String("https://google.com"),
        ExternalID: dubgo.String("123456"),
        TagIds: dubgo.Pointer(operations.CreateUpdateLinkTagIdsArrayOfStr(
            []string{
                "clux0rgak00011...",
            },
        )),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                             | Type                                                                                                                                  | Required                                                                                                                              | Description                                                                                                                           |
| ------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                 | [context.Context](https://pkg.go.dev/context#Context)                                                                                 | :heavy_check_mark:                                                                                                                    | The context to use for the request.                                                                                                   |
| `linkID`                                                                                                                              | *string*                                                                                                                              | :heavy_check_mark:                                                                                                                    | The id of the link to update. You may use either `linkId` (obtained via `/links/info` endpoint) or `externalId` prefixed with `ext_`. |
| `requestBody`                                                                                                                         | [*operations.UpdateLinkRequestBody](../../models/operations/updatelinkrequestbody.md)                                                 | :heavy_minus_sign:                                                                                                                    | N/A                                                                                                                                   |
| `opts`                                                                                                                                | [][operations.Option](../../models/operations/option.md)                                                                              | :heavy_minus_sign:                                                                                                                    | The options for this request.                                                                                                         |

### Response

**[*components.LinkSchema](../../models/components/linkschema.md), error**

### Errors

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


## CreateMany

Bulk create up to 100 links for the authenticated workspace.

### Example Usage

```go
package main

import(
	dubgo "github.com/dubinc/dub-go"
	"context"
	"github.com/dubinc/dub-go/models/operations"
	"log"
)

func main() {
    s := dubgo.New(
        dubgo.WithSecurity("DUB_API_KEY"),
    )

    ctx := context.Background()
    res, err := s.Links.CreateMany(ctx, []operations.RequestBody{
        operations.RequestBody{
            URL: "https://google.com",
            ExternalID: dubgo.String("123456"),
            TagIds: dubgo.Pointer(operations.CreateBulkCreateLinksTagIdsArrayOfStr(
                []string{
                    "clux0rgak00011...",
                },
            )),
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

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `request`                                                | [[]operations.RequestBody](../../.md)                    | :heavy_check_mark:                                       | The request object to use for the request.               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[[]components.LinkSchema](../../.md), error**

### Errors

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


## DeleteMany

Bulk delete up to 100 links for the authenticated workspace.

### Example Usage

```go
package main

import(
	dubgo "github.com/dubinc/dub-go"
	"context"
	"github.com/dubinc/dub-go/models/operations"
	"log"
)

func main() {
    s := dubgo.New(
        dubgo.WithSecurity("DUB_API_KEY"),
    )

    ctx := context.Background()
    res, err := s.Links.DeleteMany(ctx, operations.BulkDeleteLinksRequest{
        LinkIds: []string{
            "clux0rgak00011...",
            "clux0rgak00022...",
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

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.BulkDeleteLinksRequest](../../models/operations/bulkdeletelinksrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.BulkDeleteLinksResponseBody](../../models/operations/bulkdeletelinksresponsebody.md), error**

### Errors

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


## UpdateMany

Bulk update up to 100 links with the same data for the authenticated workspace.

### Example Usage

```go
package main

import(
	dubgo "github.com/dubinc/dub-go"
	"context"
	"github.com/dubinc/dub-go/models/operations"
	"log"
)

func main() {
    s := dubgo.New(
        dubgo.WithSecurity("DUB_API_KEY"),
    )

    ctx := context.Background()
    res, err := s.Links.UpdateMany(ctx, &operations.BulkUpdateLinksRequestBody{
        LinkIds: []string{
            "<value>",
        },
        Data: operations.Data{
            URL: dubgo.String("https://google.com"),
            TagIds: dubgo.Pointer(operations.CreateBulkUpdateLinksTagIdsArrayOfStr(
                []string{
                    "clux0rgak00011...",
                },
            )),
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

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.BulkUpdateLinksRequestBody](../../models/operations/bulkupdatelinksrequestbody.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[[]components.LinkSchema](../../.md), error**

### Errors

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


## Upsert

Upsert a link for the authenticated workspace by its URL. If a link with the same URL already exists, return it (or update it if there are any changes). Otherwise, a new link will be created.

### Example Usage

```go
package main

import(
	dubgo "github.com/dubinc/dub-go"
	"context"
	"github.com/dubinc/dub-go/models/operations"
	"log"
)

func main() {
    s := dubgo.New(
        dubgo.WithSecurity("DUB_API_KEY"),
    )

    ctx := context.Background()
    res, err := s.Links.Upsert(ctx, &operations.UpsertLinkRequestBody{
        URL: "https://google.com",
        ExternalID: dubgo.String("123456"),
        TagIds: dubgo.Pointer(operations.CreateUpsertLinkTagIdsArrayOfStr(
            []string{
                "clux0rgak00011...",
            },
        )),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.UpsertLinkRequestBody](../../models/operations/upsertlinkrequestbody.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*components.LinkSchema](../../models/components/linkschema.md), error**

### Errors

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
