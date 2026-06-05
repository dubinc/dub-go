# RequestBody1


## Fields

| Field                                                                    | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `Type`                                                                   | [operations.RequestBodyType](../../models/operations/requestbodytype.md) | :heavy_check_mark:                                                       | N/A                                                                      |
| `PartnerID`                                                              | `string`                                                                 | :heavy_check_mark:                                                       | The ID of the partner to create the commission for.                      |
| `Amount`                                                                 | `float64`                                                                | :heavy_check_mark:                                                       | The commission amount in cents.                                          |
| `Date`                                                                   | `*string`                                                                | :heavy_minus_sign:                                                       | If not provided, the current date will be used.                          |
| `Description`                                                            | `*string`                                                                | :heavy_minus_sign:                                                       | The description of the commission.                                       |