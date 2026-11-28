# ListProgramApplicationsQueryParamStatus

Filter applications by status. One of `pending`, `approved`, or `rejected`. Defaults to `pending`.

## Example Usage

```go
import (
	"github.com/dubinc/dub-go/models/operations"
)

value := operations.ListProgramApplicationsQueryParamStatusPending
```


## Values

| Name                                              | Value                                             |
| ------------------------------------------------- | ------------------------------------------------- |
| `ListProgramApplicationsQueryParamStatusPending`  | pending                                           |
| `ListProgramApplicationsQueryParamStatusApproved` | approved                                          |
| `ListProgramApplicationsQueryParamStatusRejected` | rejected                                          |