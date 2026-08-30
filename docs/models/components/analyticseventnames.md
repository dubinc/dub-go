# AnalyticsEventNames


## Fields

| Field                                                    | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `EventName`                                              | `string`                                                 | :heavy_check_mark:                                       | The name of the conversion event (lead or sale)          |
| `Clicks`                                                 | `*float64`                                               | :heavy_minus_sign:                                       | The number of clicks from this event name                |
| `Leads`                                                  | `*float64`                                               | :heavy_minus_sign:                                       | The number of leads from this event name                 |
| `Sales`                                                  | `*float64`                                               | :heavy_minus_sign:                                       | The number of sales from this event name                 |
| `SaleAmount`                                             | `*float64`                                               | :heavy_minus_sign:                                       | The total amount of sales from this event name, in cents |