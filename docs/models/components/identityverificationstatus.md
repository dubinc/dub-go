# IdentityVerificationStatus

The partner's identity verification status. Null means not yet initiated.

## Example Usage

```go
import (
	"github.com/dubinc/dub-go/models/components"
)

value := components.IdentityVerificationStatusStarted
```


## Values

| Name                                              | Value                                             |
| ------------------------------------------------- | ------------------------------------------------- |
| `IdentityVerificationStatusStarted`               | started                                           |
| `IdentityVerificationStatusSubmitted`             | submitted                                         |
| `IdentityVerificationStatusApproved`              | approved                                          |
| `IdentityVerificationStatusDeclined`              | declined                                          |
| `IdentityVerificationStatusResubmissionRequested` | resubmissionRequested                             |
| `IdentityVerificationStatusExpired`               | expired                                           |
| `IdentityVerificationStatusAbandoned`             | abandoned                                         |
| `IdentityVerificationStatusReview`                | review                                            |