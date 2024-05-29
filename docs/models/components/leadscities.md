# LeadsCities


## Fields

| Field                                                                          | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `City`                                                                         | *string*                                                                       | :heavy_check_mark:                                                             | The name of the city                                                           |
| `Country`                                                                      | [components.LeadsCitiesCountry](../../models/components/leadscitiescountry.md) | :heavy_check_mark:                                                             | The 2-letter country code of the city: https://d.to/geo                        |
| `Leads`                                                                        | *float64*                                                                      | :heavy_check_mark:                                                             | The number of leads from this city                                             |