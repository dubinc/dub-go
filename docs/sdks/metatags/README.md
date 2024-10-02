# Metatags
(*Metatags*)

## Overview

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
	"github.com/dubinc/dub-go/models/operations"
	"log"
)

func main() {
    s := dubgo.New(
        dubgo.WithSecurity("DUB_API_KEY"),
    )

    ctx := context.Background()
    res, err := s.Metatags.Get(ctx, operations.GetMetatagsRequest{
        URL: "https://dub.co",
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
| `request`                                                                      | [operations.GetMetatagsRequest](../../models/operations/getmetatagsrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.GetMetatagsResponseBody](../../models/operations/getmetatagsresponsebody.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |