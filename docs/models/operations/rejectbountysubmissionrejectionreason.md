# RejectBountySubmissionRejectionReason

The reason for rejecting the submission.

## Example Usage

```go
import (
	"github.com/dubinc/dub-go/models/operations"
)

value := operations.RejectBountySubmissionRejectionReasonInvalidProof
```


## Values

| Name                                                       | Value                                                      |
| ---------------------------------------------------------- | ---------------------------------------------------------- |
| `RejectBountySubmissionRejectionReasonInvalidProof`        | invalidProof                                               |
| `RejectBountySubmissionRejectionReasonDuplicateSubmission` | duplicateSubmission                                        |
| `RejectBountySubmissionRejectionReasonOutOfTimeWindow`     | outOfTimeWindow                                            |
| `RejectBountySubmissionRejectionReasonDidNotMeetCriteria`  | didNotMeetCriteria                                         |
| `RejectBountySubmissionRejectionReasonOther`               | other                                                      |