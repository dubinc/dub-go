# AnalyticsRefererUrls


## Fields

| Field                                                             | Type                                                              | Required                                                          | Description                                                       |
| ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- |
| `RefererURL`                                                      | *string*                                                          | :heavy_check_mark:                                                | The full URL of the referer. If unknown, this will be `(direct)`  |
| `Clicks`                                                          | **float64*                                                        | :heavy_minus_sign:                                                | The number of clicks from this referer to this URL                |
| `Leads`                                                           | **float64*                                                        | :heavy_minus_sign:                                                | The number of leads from this referer to this URL                 |
| `Sales`                                                           | **float64*                                                        | :heavy_minus_sign:                                                | The number of sales from this referer to this URL                 |
| `SaleAmount`                                                      | **float64*                                                        | :heavy_minus_sign:                                                | The total amount of sales from this referer to this URL, in cents |