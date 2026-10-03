# ProgramApplications

## Overview

### Available Operations

* [List](#list) - List all program applications
* [Approve](#approve) - Approve a partner application
* [Reject](#reject) - Reject a partner application

## List

Retrieve a paginated list of applications for your partner program. Filter by `status` to list pending, approved, or rejected applications.

### Example Usage

<!-- UsageSnippet language="go" operationID="listProgramApplications" method="get" path="/program-applications" -->
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

    res, err := s.ProgramApplications.List(ctx, operations.ListProgramApplicationsRequest{
        Country: dubgo.Pointer("US"),
        GroupID: dubgo.Pointer("grp_123"),
        Page: dubgo.Pointer[int64](1),
        PageSize: dubgo.Pointer[int64](50),
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

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.ListProgramApplicationsRequest](../../models/operations/listprogramapplicationsrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../models/operations/option.md)                                               | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[[]operations.ListProgramApplicationsResponseBody](../../.md), error**

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

## Approve

Approve a pending partner application to your program. The partner will be enrolled in the specified group and notified of the approval.

### Example Usage

<!-- UsageSnippet language="go" operationID="approveProgramApplication" method="post" path="/program-applications/approve" -->
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

    res, err := s.ProgramApplications.Approve(ctx, operations.ApproveProgramApplicationRequestBody{
        PartnerID: "<id>",
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
| `request`                                                                                                          | [operations.ApproveProgramApplicationRequestBody](../../models/operations/approveprogramapplicationrequestbody.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |
| `opts`                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                           | :heavy_minus_sign:                                                                                                 | The options for this request.                                                                                      |

### Response

**[*operations.ApproveProgramApplicationResponseBody](../../models/operations/approveprogramapplicationresponsebody.md), error**

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

## Reject

Reject a pending partner application to your program. The partner will be notified via email that their application was not approved.

### Example Usage

<!-- UsageSnippet language="go" operationID="rejectProgramApplication" method="post" path="/program-applications/reject" -->
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

    res, err := s.ProgramApplications.Reject(ctx, operations.RejectProgramApplicationRequestBody{
        PartnerID: "<id>",
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

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.RejectProgramApplicationRequestBody](../../models/operations/rejectprogramapplicationrequestbody.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `opts`                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                         | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |

### Response

**[*operations.RejectProgramApplicationResponseBody](../../models/operations/rejectprogramapplicationresponsebody.md), error**

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