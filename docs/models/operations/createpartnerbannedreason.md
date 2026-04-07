# CreatePartnerBannedReason

If the partner was banned from the program, this is the reason for the ban.

## Example Usage

```go
import (
	"github.com/dubinc/dub-go/models/operations"
)

value := operations.CreatePartnerBannedReasonTosViolation
```


## Values

| Name                                            | Value                                           |
| ----------------------------------------------- | ----------------------------------------------- |
| `CreatePartnerBannedReasonTosViolation`         | tos_violation                                   |
| `CreatePartnerBannedReasonInappropriateContent` | inappropriate_content                           |
| `CreatePartnerBannedReasonFakeTraffic`          | fake_traffic                                    |
| `CreatePartnerBannedReasonFraud`                | fraud                                           |
| `CreatePartnerBannedReasonSpam`                 | spam                                            |
| `CreatePartnerBannedReasonBrandAbuse`           | brand_abuse                                     |