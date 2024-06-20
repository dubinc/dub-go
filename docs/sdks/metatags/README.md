# Metatags
(*Metatags*)

### Available Operations

* [Get](#get) - Retrieve the metatags for a URL

## Get

Retrieve the metatags for a URL.

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
    var url_ string = "https://dub.co"
    ctx := context.Background()
    res, err := s.Metatags.Get(ctx, url_)
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
| `url_`                                                | *string*                                              | :heavy_check_mark:                                    | The URL to retrieve metatags for.                     | https://dub.co                                        |


### Response

**[*operations.GetMetatagsResponse](../../models/operations/getmetatagsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
