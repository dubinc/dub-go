# AnalyticsTriggers


## Fields

| Field                                                        | Type                                                         | Required                                                     | Description                                                  |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `Trigger`                                                    | [components.Trigger](../../models/components/trigger.md)     | :heavy_check_mark:                                           | The type of trigger method: link click or QR scan            |
| `Clicks`                                                     | **float64*                                                   | :heavy_minus_sign:                                           | The number of clicks from this trigger method                |
| `Leads`                                                      | **float64*                                                   | :heavy_minus_sign:                                           | The number of leads from this trigger method                 |
| `Sales`                                                      | **float64*                                                   | :heavy_minus_sign:                                           | The number of sales from this trigger method                 |
| `SaleAmount`                                                 | **float64*                                                   | :heavy_minus_sign:                                           | The total amount of sales from this trigger method, in cents |