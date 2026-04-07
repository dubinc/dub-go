# QueryParamStatus

Filter the list of commissions by their corresponding status.

## Example Usage

```go
import (
	"github.com/dubinc/dub-go/models/operations"
)

value := operations.QueryParamStatusPending
```


## Values

| Name                        | Value                       |
| --------------------------- | --------------------------- |
| `QueryParamStatusPending`   | pending                     |
| `QueryParamStatusProcessed` | processed                   |
| `QueryParamStatusPaid`      | paid                        |
| `QueryParamStatusRefunded`  | refunded                    |
| `QueryParamStatusDuplicate` | duplicate                   |
| `QueryParamStatusFraud`     | fraud                       |
| `QueryParamStatusCanceled`  | canceled                    |