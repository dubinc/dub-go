<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	dubgo "github.com/dubinc/dub-go"
	"github.com/dubinc/dub-go/models/operations"
	"log"
)

func main() {
	s := dubgo.New(
		dubgo.WithSecurity("DUB_API_KEY"),
	)
	var request *operations.CreateLinkRequestBody = &operations.CreateLinkRequestBody{
		URL:        "https://google/com",
		ExternalID: dubgo.String("123456"),
		TagIds: operations.CreateTagIdsArrayOfStr(
			[]string{
				"clux0rgak00011...",
			},
		),
	}
	ctx := context.Background()
	res, err := s.Links.Create(ctx, request)
	if err != nil {
		log.Fatal(err)
	}
	if res.LinkSchema != nil {
		// handle response
	}
}

```

```go
package main

import (
	"context"
	dubgo "github.com/dubinc/dub-go"
	"github.com/dubinc/dub-go/models/operations"
	"log"
)

func main() {
	s := dubgo.New(
		dubgo.WithSecurity("DUB_API_KEY"),
	)
	var request *operations.UpsertLinkRequestBody = &operations.UpsertLinkRequestBody{
		URL:        "https://google/com",
		ExternalID: dubgo.String("123456"),
		TagIds: operations.CreateUpsertLinkTagIdsArrayOfStr(
			[]string{
				"clux0rgak00011...",
			},
		),
	}
	ctx := context.Background()
	res, err := s.Links.Upsert(ctx, request)
	if err != nil {
		log.Fatal(err)
	}
	if res.LinkSchema != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->