# Sale


## Fields

| Field                                                                      | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `Amount`                                                                   | *int64*                                                                    | :heavy_check_mark:                                                         | The amount of the sale. Should be passed in cents.                         |
| `InvoiceID`                                                                | **string*                                                                  | :heavy_minus_sign:                                                         | The invoice ID of the sale.                                                |
| `PaymentProcessor`                                                         | [components.PaymentProcessor](../../models/components/paymentprocessor.md) | :heavy_check_mark:                                                         | The payment processor via which the sale was made.                         |