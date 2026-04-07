# PaymentProcessor

The payment processor via which the sale was made.

## Example Usage

```go
import (
	"github.com/dubinc/dub-go/models/operations"
)

value := operations.PaymentProcessorStripe
```


## Values

| Name                         | Value                        |
| ---------------------------- | ---------------------------- |
| `PaymentProcessorStripe`     | stripe                       |
| `PaymentProcessorShopify`    | shopify                      |
| `PaymentProcessorPolar`      | polar                        |
| `PaymentProcessorPaddle`     | paddle                       |
| `PaymentProcessorRevenuecat` | revenuecat                   |
| `PaymentProcessorCustom`     | custom                       |