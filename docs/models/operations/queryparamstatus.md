# QueryParamStatus

Filter the list of partners by their status for your program. `discover` returns the partners that are not in your program and that you have not ignored, ranked by relevance.

## Example Usage

```go
import (
	"github.com/dubinc/dub-go/models/operations"
)

value := operations.QueryParamStatusDiscover
```


## Values

| Name                        | Value                       |
| --------------------------- | --------------------------- |
| `QueryParamStatusDiscover`  | discover                    |
| `QueryParamStatusInvited`   | invited                     |
| `QueryParamStatusRecruited` | recruited                   |
| `QueryParamStatusIgnored`   | ignored                     |