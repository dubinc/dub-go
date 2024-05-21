# Tags
(*Tags*)

### Available Operations

* [List](#list) - Retrieve a list of tags
* [Create](#create) - Create a new tag

## List

Retrieve a list of tags for the authenticated workspace.

### Example Usage

```go
package main

import(
	"github.com/dubinc/dub-go/models/components"
	dubgo "github.com/dubinc/dub-go"
	"context"
	"log"
)

func main() {
    s := dubgo.New(
        dubgo.WithSecurity("DUB_API_KEY"),
        dubgo.WithWorkspaceID("<value>"),
    )

    ctx := context.Background()
    res, err := s.Tags.List(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.TagSchemas != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetTagsResponse](../../models/operations/gettagsresponse.md), error**
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

Create a new tag for the authenticated workspace.

### Example Usage

```go
package main

import(
	"github.com/dubinc/dub-go/models/components"
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


    var tag string = "<value>"

    var color *operations.Color = operations.ColorBlue.ToPointer()

    ctx := context.Background()
    res, err := s.Tags.Create(ctx, tag, color)
    if err != nil {
        log.Fatal(err)
    }
    if res.TagSchema != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                        | Type                                                                                                                             | Required                                                                                                                         | Description                                                                                                                      |
| -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                            | :heavy_check_mark:                                                                                                               | The context to use for the request.                                                                                              |
| `tag`                                                                                                                            | *string*                                                                                                                         | :heavy_check_mark:                                                                                                               | The name of the tag to create.                                                                                                   |
| `color`                                                                                                                          | [*operations.Color](../../models/operations/color.md)                                                                            | :heavy_minus_sign:                                                                                                               | The color of the tag. If not provided, a random color will be used from the list: red, yellow, green, blue, purple, pink, brown. |


### Response

**[*operations.CreateTagResponse](../../models/operations/createtagresponse.md), error**
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
