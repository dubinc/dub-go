# ListPayoutsQueryParamStatus

Filter the list of payouts by their corresponding status.

## Example Usage

```go
import (
	"github.com/dubinc/dub-go/models/operations"
)

value := operations.ListPayoutsQueryParamStatusPending
```


## Values

| Name                                    | Value                                   |
| --------------------------------------- | --------------------------------------- |
| `ListPayoutsQueryParamStatusPending`    | pending                                 |
| `ListPayoutsQueryParamStatusProcessing` | processing                              |
| `ListPayoutsQueryParamStatusProcessed`  | processed                               |
| `ListPayoutsQueryParamStatusSent`       | sent                                    |
| `ListPayoutsQueryParamStatusCompleted`  | completed                               |
| `ListPayoutsQueryParamStatusFailed`     | failed                                  |
| `ListPayoutsQueryParamStatusCanceled`   | canceled                                |