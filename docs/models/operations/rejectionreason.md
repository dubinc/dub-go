# RejectionReason

The reason for rejecting the submission.

## Example Usage

```go
import (
	"github.com/dubinc/dub-go/models/operations"
)

value := operations.RejectionReasonInvalidProof
```


## Values

| Name                                 | Value                                |
| ------------------------------------ | ------------------------------------ |
| `RejectionReasonInvalidProof`        | invalidProof                         |
| `RejectionReasonDuplicateSubmission` | duplicateSubmission                  |
| `RejectionReasonOutOfTimeWindow`     | outOfTimeWindow                      |
| `RejectionReasonDidNotMeetCriteria`  | didNotMeetCriteria                   |
| `RejectionReasonOther`               | other                                |