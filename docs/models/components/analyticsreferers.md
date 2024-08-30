# AnalyticsReferers


## Fields

| Field                                                        | Type                                                         | Required                                                     | Description                                                  |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `Referer`                                                    | *string*                                                     | :heavy_check_mark:                                           | The name of the referer. If unknown, this will be `(direct)` |
| `Clicks`                                                     | **float64*                                                   | :heavy_minus_sign:                                           | The number of clicks from this referer                       |
| `Leads`                                                      | **float64*                                                   | :heavy_minus_sign:                                           | The number of leads from this referer                        |
| `Sales`                                                      | **float64*                                                   | :heavy_minus_sign:                                           | The number of sales from this referer                        |
| `SaleAmount`                                                 | **float64*                                                   | :heavy_minus_sign:                                           | The total amount of sales from this referer                  |