# GetCustomerResponseBody

The customer object.


## Fields

| Field                                                   | Type                                                    | Required                                                | Description                                             |
| ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- |
| `ID`                                                    | *string*                                                | :heavy_check_mark:                                      | The unique identifier of the customer in Dub.           |
| `ExternalID`                                            | *string*                                                | :heavy_check_mark:                                      | Unique identifier for the customer in the client's app. |
| `Name`                                                  | *string*                                                | :heavy_check_mark:                                      | Name of the customer.                                   |
| `Email`                                                 | **string*                                               | :heavy_minus_sign:                                      | Email of the customer.                                  |
| `Avatar`                                                | **string*                                               | :heavy_minus_sign:                                      | Avatar URL of the customer.                             |
| `CreatedAt`                                             | *string*                                                | :heavy_check_mark:                                      | The date the customer was created.                      |