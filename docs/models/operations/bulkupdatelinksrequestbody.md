# BulkUpdateLinksRequestBody


## Fields

| Field                                                                | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `LinkIds`                                                            | []*string*                                                           | :heavy_minus_sign:                                                   | The IDs of the links to update. Takes precedence over `externalIds`. |
| `ExternalIds`                                                        | []*string*                                                           | :heavy_minus_sign:                                                   | The external IDs of the links to update as stored in your database.  |
| `Data`                                                               | [operations.Data](../../models/operations/data.md)                   | :heavy_check_mark:                                                   | N/A                                                                  |