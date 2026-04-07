# BulkUpdateCommissionsStatus

The status to apply to every commission in the batch.

## Example Usage

```go
import (
	"github.com/dubinc/dub-go/models/operations"
)

value := operations.BulkUpdateCommissionsStatusPending
```


## Values

| Name                                   | Value                                  |
| -------------------------------------- | -------------------------------------- |
| `BulkUpdateCommissionsStatusPending`   | pending                                |
| `BulkUpdateCommissionsStatusRefunded`  | refunded                               |
| `BulkUpdateCommissionsStatusDuplicate` | duplicate                              |
| `BulkUpdateCommissionsStatusCanceled`  | canceled                               |
| `BulkUpdateCommissionsStatusFraud`     | fraud                                  |