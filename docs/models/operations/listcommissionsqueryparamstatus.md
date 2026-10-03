# ListCommissionsQueryParamStatus

Filter the list of commissions by their corresponding status.

## Example Usage

```go
import (
	"github.com/dubinc/dub-go/models/operations"
)

value := operations.ListCommissionsQueryParamStatusPending
```


## Values

| Name                                       | Value                                      |
| ------------------------------------------ | ------------------------------------------ |
| `ListCommissionsQueryParamStatusPending`   | pending                                    |
| `ListCommissionsQueryParamStatusProcessed` | processed                                  |
| `ListCommissionsQueryParamStatusPaid`      | paid                                       |
| `ListCommissionsQueryParamStatusRefunded`  | refunded                                   |
| `ListCommissionsQueryParamStatusDuplicate` | duplicate                                  |
| `ListCommissionsQueryParamStatusFraud`     | fraud                                      |
| `ListCommissionsQueryParamStatusCanceled`  | canceled                                   |
| `ListCommissionsQueryParamStatusHold`      | hold                                       |