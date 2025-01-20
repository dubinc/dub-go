# QRCodes
(*QRCodes*)

## Overview

### Available Operations

* [Get](#get) - Retrieve a QR code

## Get

Retrieve a QR code for a link.

### Example Usage

```go
package main

import(
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

    res, err := s.QRCodes.Get(ctx, operations.GetQRCodeRequest{
        URL: "https://normal-making.name",
        Size: dubgo.Float64(600),
        Level: operations.LevelL.ToPointer(),
        FgColor: dubgo.String("#000000"),
        BgColor: dubgo.String("#FFFFFF"),
        HideLogo: dubgo.Bool(true),
        Margin: dubgo.Float64(2),
        IncludeMargin: dubgo.Bool(true),
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

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [operations.GetQRCodeRequest](../../models/operations/getqrcoderequest.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |
| `opts`                                                                     | [][operations.Option](../../models/operations/option.md)                   | :heavy_minus_sign:                                                         | The options for this request.                                              |

### Response

**[*string](../../.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
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
| sdkerrors.SDKError            | 4XX, 5XX                      | \*/\*                         |