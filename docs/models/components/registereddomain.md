# RegisteredDomain

The registered domain record.


## Fields

| Field                                       | Type                                        | Required                                    | Description                                 |
| ------------------------------------------- | ------------------------------------------- | ------------------------------------------- | ------------------------------------------- |
| `ID`                                        | *string*                                    | :heavy_check_mark:                          | The ID of the registered domain record.     |
| `AutoRenewalDisabledAt`                     | *string*                                    | :heavy_check_mark:                          | The date the domain auto-renew is disabled. |
| `CreatedAt`                                 | *string*                                    | :heavy_check_mark:                          | The date the domain was created.            |
| `ExpiresAt`                                 | *string*                                    | :heavy_check_mark:                          | The date the domain expires.                |
| `RenewalFee`                                | *float64*                                   | :heavy_check_mark:                          | The fee to renew the domain.                |