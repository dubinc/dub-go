# Type

Filter the list of commissions by type. Supports advanced filtering: single value, multiple values (comma-separated), or exclusion (prefix with `-`). Examples: `sale`, `sale,lead`, `-click`.

## Example Usage

```go
import (
	"github.com/dubinc/dub-go/models/operations"
)

value := operations.TypeClick
```


## Values

| Name         | Value        |
| ------------ | ------------ |
| `TypeClick`  | click        |
| `TypeLead`   | lead         |
| `TypeSale`   | sale         |
| `TypeCustom` | custom       |