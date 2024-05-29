# SalesCities


## Fields

| Field                                                                          | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `City`                                                                         | *string*                                                                       | :heavy_check_mark:                                                             | The name of the city                                                           |
| `Country`                                                                      | [components.SalesCitiesCountry](../../models/components/salescitiescountry.md) | :heavy_check_mark:                                                             | The 2-letter country code of the city: https://d.to/geo                        |
| `Sales`                                                                        | *float64*                                                                      | :heavy_check_mark:                                                             | The number of sales from this city                                             |
| `Amount`                                                                       | *float64*                                                                      | :heavy_check_mark:                                                             | The total amount of sales from this city                                       |