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
	ctx := context.Background()

	s := dubgo.New(
		dubgo.WithSecurity("DUB_API_KEY"),
	)

	res, err := s.Links.Create(ctx, &operations.CreateLinkRequestBody{
		URL:        "https://google.com",
		ExternalID: dubgo.String("123456"),
		TagIds: dubgo.Pointer(operations.CreateTagIdsArrayOfStr(
			[]string{
				"clux0rgak00011...",
			},
		)),
		TestVariants: []operations.TestVariants{
			operations.TestVariants{
				URL:        "https://example.com/variant-1",
				Percentage: 50,
			},
			operations.TestVariants{
				URL:        "https://example.com/variant-2",
				Percentage: 50,
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

```go
package main

import (
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

	res, err := s.Links.Upsert(ctx, &operations.UpsertLinkRequestBody{
		URL:        "https://google.com",
		ExternalID: dubgo.String("123456"),
		TagIds: dubgo.Pointer(operations.CreateUpsertLinkTagIdsArrayOfStr(
			[]string{
				"clux0rgak00011...",
			},
		)),
		TestVariants: []operations.UpsertLinkTestVariants{
			operations.UpsertLinkTestVariants{
				URL:        "https://example.com/variant-1",
				Percentage: 50,
			},
			operations.UpsertLinkTestVariants{
				URL:        "https://example.com/variant-2",
				Percentage: 50,
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
<!-- End SDK Example Usage [usage] -->