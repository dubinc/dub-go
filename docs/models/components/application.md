# Application

Linked program application, including review outcome when applicable.


## Fields

| Field                                                                     | Type                                                                      | Required                                                                  | Description                                                               |
| ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- |
| `RejectionReason`                                                         | [*components.RejectionReason](../../models/components/rejectionreason.md) | :heavy_check_mark:                                                        | Preset reason when the application was rejected.                          |
| `RejectionNote`                                                           | `*string`                                                                 | :heavy_check_mark:                                                        | Free-form note when the application was rejected.                         |
| `ReviewedAt`                                                              | `*string`                                                                 | :heavy_check_mark:                                                        | When the application was approved or rejected.                            |