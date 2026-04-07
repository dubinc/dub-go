# ResponseBodyPaymentProcessor

The payment processor via which the sale was made.

## Example Usage

```go
import (
	"github.com/dubinc/dub-go/models/operations"
)

value := operations.ResponseBodyPaymentProcessorStripe
```


## Values

| Name                                     | Value                                    |
| ---------------------------------------- | ---------------------------------------- |
| `ResponseBodyPaymentProcessorStripe`     | stripe                                   |
| `ResponseBodyPaymentProcessorShopify`    | shopify                                  |
| `ResponseBodyPaymentProcessorPolar`      | polar                                    |
| `ResponseBodyPaymentProcessorPaddle`     | paddle                                   |
| `ResponseBodyPaymentProcessorRevenuecat` | revenuecat                               |
| `ResponseBodyPaymentProcessorCustom`     | custom                                   |