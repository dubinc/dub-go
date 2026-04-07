# CreatePartnerApplication

Linked program application, including review outcome when applicable.


## Fields

| Field                                                                                               | Type                                                                                                | Required                                                                                            | Description                                                                                         |
| --------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------- |
| `RejectionReason`                                                                                   | [*operations.CreatePartnerRejectionReason](../../models/operations/createpartnerrejectionreason.md) | :heavy_check_mark:                                                                                  | Preset reason when the application was rejected.                                                    |
| `RejectionNote`                                                                                     | `*string`                                                                                           | :heavy_check_mark:                                                                                  | Free-form note when the application was rejected.                                                   |
| `ReviewedAt`                                                                                        | `*string`                                                                                           | :heavy_check_mark:                                                                                  | When the application was approved or rejected.                                                      |