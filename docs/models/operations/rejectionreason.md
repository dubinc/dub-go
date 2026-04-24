# RejectionReason

The reason for rejecting the partner application. This will be shared with the partner via email.

## Example Usage

```go
import (
	"github.com/dubinc/dub-go/models/operations"
)

value := operations.RejectionReasonNeedsMoreDetail
```


## Values

| Name                                     | Value                                    |
| ---------------------------------------- | ---------------------------------------- |
| `RejectionReasonNeedsMoreDetail`         | needsMoreDetail                          |
| `RejectionReasonDoesNotMeetRequirements` | doesNotMeetRequirements                  |
| `RejectionReasonNotTheRightFit`          | notTheRightFit                           |
| `RejectionReasonOther`                   | other                                    |