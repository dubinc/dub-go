# GetCustomersResponseBody


## Fields

| Field                                                       | Type                                                        | Required                                                    | Description                                                 |
| ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- |
| `ID`                                                        | *string*                                                    | :heavy_check_mark:                                          | The unique identifier of the customer in Dub.               |
| `ExternalID`                                                | *string*                                                    | :heavy_check_mark:                                          | Unique identifier for the customer in the client's app.     |
| `Name`                                                      | *string*                                                    | :heavy_check_mark:                                          | Name of the customer.                                       |
| `Email`                                                     | **string*                                                   | :heavy_minus_sign:                                          | Email of the customer.                                      |
| `Avatar`                                                    | **string*                                                   | :heavy_minus_sign:                                          | Avatar URL of the customer.                                 |
| `Country`                                                   | **string*                                                   | :heavy_minus_sign:                                          | Country of the customer.                                    |
| `CreatedAt`                                                 | *string*                                                    | :heavy_check_mark:                                          | The date the customer was created.                          |
| `Link`                                                      | [*operations.Link](../../models/operations/link.md)         | :heavy_minus_sign:                                          | N/A                                                         |
| `Partner`                                                   | [*operations.Partner](../../models/operations/partner.md)   | :heavy_minus_sign:                                          | N/A                                                         |
| `Discount`                                                  | [*operations.Discount](../../models/operations/discount.md) | :heavy_minus_sign:                                          | N/A                                                         |