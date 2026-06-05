# ReapplicationTimeframe

The mode for reapplying for the program. `instant`: The partner can reapply immediately. `standard`: The partner can reapply after 30 days. `never`: The partner can never reapply for the program. Defaults to `standard` if undefined.

## Example Usage

```go
import (
	"github.com/dubinc/dub-go/models/operations"
)

value := operations.ReapplicationTimeframeInstant
```


## Values

| Name                             | Value                            |
| -------------------------------- | -------------------------------- |
| `ReapplicationTimeframeInstant`  | instant                          |
| `ReapplicationTimeframeStandard` | standard                         |
| `ReapplicationTimeframeNever`    | never                            |