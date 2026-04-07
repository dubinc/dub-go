# CreatePartnerDefaultPayoutMethod

The partner's default payout method. Connect: Bank account payouts via Stripe Connect; Stablecoin: USDC payouts directly to a crypto wallet; PayPal: Payouts via PayPal

## Example Usage

```go
import (
	"github.com/dubinc/dub-go/models/operations"
)

value := operations.CreatePartnerDefaultPayoutMethodConnect
```


## Values

| Name                                         | Value                                        |
| -------------------------------------------- | -------------------------------------------- |
| `CreatePartnerDefaultPayoutMethodConnect`    | connect                                      |
| `CreatePartnerDefaultPayoutMethodStablecoin` | stablecoin                                   |
| `CreatePartnerDefaultPayoutMethodPaypal`     | paypal                                       |