# ListPartnerApplicationsRequest


## Fields

| Field                                                        | Type                                                         | Required                                                     | Description                                                  | Example                                                      |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `Country`                                                    | `*string`                                                    | :heavy_minus_sign:                                           | A filter on the list based on the partner's `country` field. | US                                                           |
| `GroupID`                                                    | `*string`                                                    | :heavy_minus_sign:                                           | A filter on the list based on the partner's `groupId` field. | grp_123                                                      |
| `Page`                                                       | `*float64`                                                   | :heavy_minus_sign:                                           | The page number for pagination.                              | 1                                                            |
| `PageSize`                                                   | `*float64`                                                   | :heavy_minus_sign:                                           | The number of items per page.                                | 50                                                           |