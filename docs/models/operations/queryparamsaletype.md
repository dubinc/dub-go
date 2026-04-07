# QueryParamSaleType

Filter sales by type: 'new' for first-time purchases, 'recurring' for repeat purchases. If undefined, returns both.

## Example Usage

```go
import (
	"github.com/dubinc/dub-go/models/operations"
)

value := operations.QueryParamSaleTypeNew
```


## Values

| Name                          | Value                         |
| ----------------------------- | ----------------------------- |
| `QueryParamSaleTypeNew`       | new                           |
| `QueryParamSaleTypeRecurring` | recurring                     |