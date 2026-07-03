# AccessLevel

The workspace-level access level settings for the folder. Default is `write` which allows full access to the folder for all team members. The other options are `read` (view-only access) and `null` (no access) and are only available on Business plans and above.

## Example Usage

```go
import (
	"github.com/dubinc/dub-go/models/components"
)

value := components.AccessLevelWrite
```


## Values

| Name               | Value              |
| ------------------ | ------------------ |
| `AccessLevelWrite` | write              |
| `AccessLevelRead`  | read               |