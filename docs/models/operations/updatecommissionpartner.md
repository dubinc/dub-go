# UpdateCommissionPartner


## Fields

| Field                                                             | Type                                                              | Required                                                          | Description                                                       |
| ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- |
| `ID`                                                              | *string*                                                          | :heavy_check_mark:                                                | The partner's unique ID on Dub.                                   |
| `Name`                                                            | *string*                                                          | :heavy_check_mark:                                                | The partner's full legal name.                                    |
| `Email`                                                           | *string*                                                          | :heavy_check_mark:                                                | The partner's email address. Should be a unique value across Dub. |
| `Image`                                                           | *string*                                                          | :heavy_check_mark:                                                | The partner's avatar image.                                       |
| `PayoutsEnabledAt`                                                | *string*                                                          | :heavy_check_mark:                                                | The date when the partner enabled payouts.                        |
| `Country`                                                         | *string*                                                          | :heavy_check_mark:                                                | The partner's country (required for tax purposes).                |