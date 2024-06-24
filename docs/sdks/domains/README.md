# Domains
(*Domains*)

### Available Operations

* [List](#list) - Retrieve a list of domains
* [Create](#create) - Create a domain
* [Delete](#delete) - Delete a domain
* [Update](#update) - Update a domain
* [SetPrimary](#setprimary) - Set a domain as primary
* [Transfer](#transfer) - Transfer a domain

## List

Retrieve a list of domains associated with the authenticated workspace.

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
    res, err := s.Domains.List(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.DomainSchemas != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.ListDomainsResponse](../../models/operations/listdomainsresponse.md), error**
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

Create a domain for the authenticated workspace.

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
    var slug string = "acme.com"

    var expiredURL *string = dubgo.String("https://acme.com/expired")

    var archived *bool = dubgo.Bool(false)

    var placeholder *string = dubgo.String("https://dub.co/help/article/what-is-dub")
    ctx := context.Background()
    res, err := s.Domains.Create(ctx, slug, expiredURL, archived, placeholder)
    if err != nil {
        log.Fatal(err)
    }
    if res.DomainSchema != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        | Example                                                                                                            |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |                                                                                                                    |
| `slug`                                                                                                             | *string*                                                                                                           | :heavy_check_mark:                                                                                                 | Name of the domain.                                                                                                | acme.com                                                                                                           |
| `expiredURL`                                                                                                       | **string*                                                                                                          | :heavy_minus_sign:                                                                                                 | Redirect users to a specific URL when any link under this domain has expired.                                      | https://acme.com/expired                                                                                           |
| `archived`                                                                                                         | **bool*                                                                                                            | :heavy_minus_sign:                                                                                                 | Whether to archive this domain. `false` will unarchive a previously archived domain.                               | false                                                                                                              |
| `placeholder`                                                                                                      | **string*                                                                                                          | :heavy_minus_sign:                                                                                                 | Provide context to your teammates in the link creation modal by showing them an example of a link to be shortened. | https://dub.co/help/article/what-is-dub                                                                            |


### Response

**[*operations.CreateDomainResponse](../../models/operations/createdomainresponse.md), error**
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

Delete a domain from a workspace. It cannot be undone. This will also delete all the links associated with the domain.

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
    var slug string = "acme.com"
    ctx := context.Background()
    res, err := s.Domains.Delete(ctx, slug)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           | Example                                               |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |                                                       |
| `slug`                                                | *string*                                              | :heavy_check_mark:                                    | The domain name.                                      | acme.com                                              |


### Response

**[*operations.DeleteDomainResponse](../../models/operations/deletedomainresponse.md), error**
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

Update a domain for the authenticated workspace.

### Example Usage

```go
package main

import(
	dubgo "github.com/dubinc/dub-go"
	"github.com/dubinc/dub-go/models/operations"
	"context"
	"log"
)

func main() {
    s := dubgo.New(
        dubgo.WithSecurity("DUB_API_KEY"),
    )
    var slug string = "acme.com"

    var requestBody *operations.UpdateDomainRequestBody = &operations.UpdateDomainRequestBody{
        Slug: dubgo.String("acme.com"),
        ExpiredURL: dubgo.String("https://acme.com/expired"),
        Archived: dubgo.Bool(false),
        Placeholder: dubgo.String("https://dub.co/help/article/what-is-dub"),
    }
    ctx := context.Background()
    res, err := s.Domains.Update(ctx, slug, requestBody)
    if err != nil {
        log.Fatal(err)
    }
    if res.DomainSchema != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                 | Type                                                                                      | Required                                                                                  | Description                                                                               | Example                                                                                   |
| ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- |
| `ctx`                                                                                     | [context.Context](https://pkg.go.dev/context#Context)                                     | :heavy_check_mark:                                                                        | The context to use for the request.                                                       |                                                                                           |
| `slug`                                                                                    | *string*                                                                                  | :heavy_check_mark:                                                                        | The domain name.                                                                          | acme.com                                                                                  |
| `requestBody`                                                                             | [*operations.UpdateDomainRequestBody](../../models/operations/updatedomainrequestbody.md) | :heavy_minus_sign:                                                                        | N/A                                                                                       |                                                                                           |


### Response

**[*operations.UpdateDomainResponse](../../models/operations/updatedomainresponse.md), error**
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

## SetPrimary

Set a domain as primary for the authenticated workspace.

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
    var slug string = "acme.com"
    ctx := context.Background()
    res, err := s.Domains.SetPrimary(ctx, slug)
    if err != nil {
        log.Fatal(err)
    }
    if res.DomainSchema != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           | Example                                               |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |                                                       |
| `slug`                                                | *string*                                              | :heavy_check_mark:                                    | The domain name.                                      | acme.com                                              |


### Response

**[*operations.SetPrimaryDomainResponse](../../models/operations/setprimarydomainresponse.md), error**
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

## Transfer

Transfer a domain to another workspace within the authenticated account.

### Example Usage

```go
package main

import(
	dubgo "github.com/dubinc/dub-go"
	"github.com/dubinc/dub-go/models/operations"
	"context"
	"log"
)

func main() {
    s := dubgo.New(
        dubgo.WithSecurity("DUB_API_KEY"),
    )
    var slug string = "acme.com"

    var requestBody *operations.TransferDomainRequestBody = &operations.TransferDomainRequestBody{
        NewWorkspaceID: "<value>",
    }
    ctx := context.Background()
    res, err := s.Domains.Transfer(ctx, slug, requestBody)
    if err != nil {
        log.Fatal(err)
    }
    if res.DomainSchema != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                     | Type                                                                                          | Required                                                                                      | Description                                                                                   | Example                                                                                       |
| --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- |
| `ctx`                                                                                         | [context.Context](https://pkg.go.dev/context#Context)                                         | :heavy_check_mark:                                                                            | The context to use for the request.                                                           |                                                                                               |
| `slug`                                                                                        | *string*                                                                                      | :heavy_check_mark:                                                                            | The domain name.                                                                              | acme.com                                                                                      |
| `requestBody`                                                                                 | [*operations.TransferDomainRequestBody](../../models/operations/transferdomainrequestbody.md) | :heavy_minus_sign:                                                                            | N/A                                                                                           |                                                                                               |


### Response

**[*operations.TransferDomainResponse](../../models/operations/transferdomainresponse.md), error**
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
