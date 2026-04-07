# ListPayoutsDefaultPayoutMethod

The partner's default payout method. Connect: Bank account payouts via Stripe Connect; Stablecoin: USDC payouts directly to a crypto wallet; PayPal: Payouts via PayPal

## Example Usage

```go
import (
	"github.com/dubinc/dub-go/models/operations"
)

value := operations.ListPayoutsDefaultPayoutMethodConnect
```


## Values

| Name                                       | Value                                      |
| ------------------------------------------ | ------------------------------------------ |
| `ListPayoutsDefaultPayoutMethodConnect`    | connect                                    |
| `ListPayoutsDefaultPayoutMethodStablecoin` | stablecoin                                 |
| `ListPayoutsDefaultPayoutMethodPaypal`     | paypal                                     |