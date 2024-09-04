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

	ctx := context.Background()
	res, err := s.Links.Create(ctx, &operations.CreateLinkRequestBody{
		URL:        "https://google.com",
		ExternalID: dubgo.String("123456"),
		TagIds: operations.CreateTagIdsStr(
			"[\"clux0rgak00011...\"]",
		),
	})
	if err != nil {
		log.Fatal(err)
	}
	if res != nil {
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

	ctx := context.Background()
	res, err := s.Links.Upsert(ctx, &operations.UpsertLinkRequestBody{
		URL:        "https://google.com",
		ExternalID: dubgo.String("123456"),
		TagIds: operations.CreateUpsertLinkTagIdsArrayOfStr(
			[]string{
				"clux0rgak00011...",
			},
		),
	})
	if err != nil {
		log.Fatal(err)
	}
	if res != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->