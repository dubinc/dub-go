# QueryParamStatus

Filter applications by status. One of `pending`, `approved`, or `rejected`. Defaults to `pending`.

## Example Usage

```go
import (
	"github.com/dubinc/dub-go/models/operations"
)

value := operations.QueryParamStatusPending
```


## Values

| Name                       | Value                      |
| -------------------------- | -------------------------- |
| `QueryParamStatusPending`  | pending                    |
| `QueryParamStatusApproved` | approved                   |
| `QueryParamStatusRejected` | rejected                   |