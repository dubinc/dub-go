<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	dubgo "github.com/dubinc/dub-go"
	"github.com/dubinc/dub-go/models/components"
	"github.com/dubinc/dub-go/models/operations"
	"log"
)

func main() {
	s := dubgo.New(
		dubgo.WithSecurity("DUB_API_KEY"),
		dubgo.WithWorkspaceID("<value>"),
	)

	ctx := context.Background()
	res, err := s.Links.Create(ctx, &operations.CreateLinkRequestBody{
		URL:        "https://google/com",
		ExternalID: dubgo.String("123456"),
		TagIds: operations.CreateTagIdsArrayOfstr(
			[]string{
				"clux0rgak00011...",
			},
		),
	})
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
	"github.com/dubinc/dub-go/models/components"
	"github.com/dubinc/dub-go/models/operations"
	"log"
)

func main() {
	s := dubgo.New(
		dubgo.WithSecurity("DUB_API_KEY"),
		dubgo.WithWorkspaceID("<value>"),
	)

	ctx := context.Background()
	res, err := s.Links.Upsert(ctx, &operations.UpsertLinkRequestBody{
		URL:        "https://google/com",
		ExternalID: dubgo.String("123456"),
		TagIds: operations.CreateUpsertLinkTagIdsArrayOfstr(
			[]string{
				"clux0rgak00011...",
			},
		),
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.LinkSchema != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->