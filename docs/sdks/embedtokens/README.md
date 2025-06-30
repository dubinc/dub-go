# EmbedTokens
(*EmbedTokens*)

## Overview

### Available Operations

* [Referrals](#referrals) - Create a referrals embed token

## Referrals

Create a referrals embed token for the given partner/tenant.

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

    res, err := s.EmbedTokens.Referrals(ctx, &operations.CreateReferralsEmbedTokenRequestBody{
        Partner: &operations.Partner{
            Email: "Letha_Wuckert2@yahoo.com",
            LinkProps: &operations.CreateReferralsEmbedTokenLinkProps{
                ExternalID: dubgo.String("123456"),
                TagIds: dubgo.Pointer(operations.CreateCreateReferralsEmbedTokenTagIdsArrayOfStr(
                    []string{
                        "clux0rgak00011...",
                    },
                )),
                TestVariants: []operations.CreateReferralsEmbedTokenTestVariants{
                    operations.CreateReferralsEmbedTokenTestVariants{
                        URL: "https://example.com/variant-1",
                        Percentage: 50,
                    },
                    operations.CreateReferralsEmbedTokenTestVariants{
                        URL: "https://example.com/variant-2",
                        Percentage: 50,
                    },
                },
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

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.CreateReferralsEmbedTokenRequestBody](../../models/operations/createreferralsembedtokenrequestbody.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |
| `opts`                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                           | :heavy_minus_sign:                                                                                                 | The options for this request.                                                                                      |

### Response

**[*operations.CreateReferralsEmbedTokenResponseBody](../../models/operations/createreferralsembedtokenresponsebody.md), error**

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