# ProgramApplicationCreatedEventDefaultPayoutMethod

The partner's default payout method. Connect: Bank account payouts via Stripe Connect; Stablecoin: USDC payouts directly to a crypto wallet; PayPal: Payouts via PayPal

## Example Usage

```go
import (
	"github.com/dubinc/dub-go/models/components"
)

value := components.ProgramApplicationCreatedEventDefaultPayoutMethodConnect
```


## Values

| Name                                                          | Value                                                         |
| ------------------------------------------------------------- | ------------------------------------------------------------- |
| `ProgramApplicationCreatedEventDefaultPayoutMethodConnect`    | connect                                                       |
| `ProgramApplicationCreatedEventDefaultPayoutMethodStablecoin` | stablecoin                                                    |
| `ProgramApplicationCreatedEventDefaultPayoutMethodPaypal`     | paypal                                                        |
| `ProgramApplicationCreatedEventDefaultPayoutMethodTremendous` | tremendous                                                    |