# ClicksCities


## Fields

| Field                                                                            | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `City`                                                                           | *string*                                                                         | :heavy_check_mark:                                                               | The name of the city                                                             |
| `Country`                                                                        | [components.ClicksCitiesCountry](../../models/components/clickscitiescountry.md) | :heavy_check_mark:                                                               | The 2-letter country code of the city: https://d.to/geo                          |
| `Clicks`                                                                         | *float64*                                                                        | :heavy_check_mark:                                                               | The number of clicks from this city                                              |