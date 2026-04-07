# Status

Useful for marking a commission as pending, refunded, duplicate, canceled, or fraudulent. Takes precedence over `saleAmount` and `modifySaleAmount`. When a commission is marked as pending, refunded, duplicate, canceled, or fraudulent, it will be omitted from the payout, and the payout amount will be recalculated accordingly. Paid commissions cannot be updated.

## Example Usage

```go
import (
	"github.com/dubinc/dub-go/models/operations"
)

value := operations.StatusPending
```


## Values

| Name              | Value             |
| ----------------- | ----------------- |
| `StatusPending`   | pending           |
| `StatusRefunded`  | refunded          |
| `StatusDuplicate` | duplicate         |
| `StatusCanceled`  | canceled          |
| `StatusFraud`     | fraud             |