# RequestBodyPaymentProcessor

The payment processor via which the sale was made.

## Example Usage

```go
import (
	"github.com/dubinc/dub-go/models/operations"
)

value := operations.RequestBodyPaymentProcessorStripe
```


## Values

| Name                                      | Value                                     |
| ----------------------------------------- | ----------------------------------------- |
| `RequestBodyPaymentProcessorStripe`       | stripe                                    |
| `RequestBodyPaymentProcessorShopify`      | shopify                                   |
| `RequestBodyPaymentProcessorPolar`        | polar                                     |
| `RequestBodyPaymentProcessorPaddle`       | paddle                                    |
| `RequestBodyPaymentProcessorApple`        | apple                                     |
| `RequestBodyPaymentProcessorRevenuecat`   | revenuecat                                |
| `RequestBodyPaymentProcessorLemonsqueezy` | lemonsqueezy                              |
| `RequestBodyPaymentProcessorDub`          | dub                                       |
| `RequestBodyPaymentProcessorCustom`       | custom                                    |