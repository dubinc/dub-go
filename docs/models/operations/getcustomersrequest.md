# GetCustomersRequest


## Fields

| Field                                                                                                       | Type                                                                                                        | Required                                                                                                    | Description                                                                                                 |
| ----------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------- |
| `Email`                                                                                                     | **string*                                                                                                   | :heavy_minus_sign:                                                                                          | A case-sensitive filter on the list based on the customer's `email` field. The value must be a string.      |
| `ExternalID`                                                                                                | **string*                                                                                                   | :heavy_minus_sign:                                                                                          | A case-sensitive filter on the list based on the customer's `externalId` field. The value must be a string. |
| `IncludeExpandedFields`                                                                                     | **bool*                                                                                                     | :heavy_minus_sign:                                                                                          | Whether to include expanded fields on the customer (`link`, `partner`, `discount`).                         |