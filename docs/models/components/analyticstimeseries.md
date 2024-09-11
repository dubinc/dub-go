# AnalyticsTimeseries


## Fields

| Field                                               | Type                                                | Required                                            | Description                                         |
| --------------------------------------------------- | --------------------------------------------------- | --------------------------------------------------- | --------------------------------------------------- |
| `Start`                                             | *string*                                            | :heavy_check_mark:                                  | The starting timestamp of the interval              |
| `Clicks`                                            | **float64*                                          | :heavy_minus_sign:                                  | The number of clicks in the interval                |
| `Leads`                                             | **float64*                                          | :heavy_minus_sign:                                  | The number of leads in the interval                 |
| `Sales`                                             | **float64*                                          | :heavy_minus_sign:                                  | The number of sales in the interval                 |
| `SaleAmount`                                        | **float64*                                          | :heavy_minus_sign:                                  | The total amount of sales in the interval, in cents |