# Track
(*Track*)

### Available Operations

* [Lead](#lead) - Track a lead
* [Sale](#sale) - Track a sale
* [Customer](#customer) - Track a customer

## Lead

Track a lead for a short link.

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
        dubgo.WithWorkspaceID("<value>"),
    )
    var request *operations.TrackLeadRequestBody = &operations.TrackLeadRequestBody{
        ClickID: "<value>",
        EventName: "Sign up",
        CustomerID: "<value>",
    }
    ctx := context.Background()
    res, err := s.Track.Lead(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.TrackLeadRequestBody](../../models/operations/trackleadrequestbody.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.TrackLeadResponse](../../models/operations/trackleadresponse.md), error**
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

## Sale

Track a sale for a short link.

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
        dubgo.WithWorkspaceID("<value>"),
    )
    var request *operations.TrackSaleRequestBody = &operations.TrackSaleRequestBody{
        CustomerID: "<value>",
        Amount: 996500,
        PaymentProcessor: operations.PaymentProcessorShopify,
        EventName: dubgo.String("Purchase"),
    }
    ctx := context.Background()
    res, err := s.Track.Sale(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.TrackSaleRequestBody](../../models/operations/tracksalerequestbody.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.TrackSaleResponse](../../models/operations/tracksaleresponse.md), error**
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

## Customer

Track a customer for an authenticated workspace.

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
        dubgo.WithWorkspaceID("<value>"),
    )
    var customerID string = "<value>"

    var customerName *string = dubgo.String("<value>")

    var customerEmail *string = dubgo.String("Wilson.Smith@gmail.com")

    var customerAvatar *string = dubgo.String("<value>")
    ctx := context.Background()
    res, err := s.Track.Customer(ctx, customerID, customerName, customerEmail, customerAvatar)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                         | Type                                                                                                              | Required                                                                                                          | Description                                                                                                       |
| ----------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                             | [context.Context](https://pkg.go.dev/context#Context)                                                             | :heavy_check_mark:                                                                                                | The context to use for the request.                                                                               |
| `customerID`                                                                                                      | *string*                                                                                                          | :heavy_check_mark:                                                                                                | This is the unique identifier for the customer in the client's app. This is used to track the customer's journey. |
| `customerName`                                                                                                    | **string*                                                                                                         | :heavy_minus_sign:                                                                                                | Name of the customer in the client's app.                                                                         |
| `customerEmail`                                                                                                   | **string*                                                                                                         | :heavy_minus_sign:                                                                                                | Email of the customer in the client's app.                                                                        |
| `customerAvatar`                                                                                                  | **string*                                                                                                         | :heavy_minus_sign:                                                                                                | Avatar of the customer in the client's app.                                                                       |


### Response

**[*operations.TrackCustomerResponse](../../models/operations/trackcustomerresponse.md), error**
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
