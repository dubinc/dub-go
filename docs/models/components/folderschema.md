# FolderSchema


## Fields

| Field                                                             | Type                                                              | Required                                                          | Description                                                       |
| ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- |
| `ID`                                                              | *string*                                                          | :heavy_check_mark:                                                | The unique ID of the folder.                                      |
| `Name`                                                            | *string*                                                          | :heavy_check_mark:                                                | The name of the folder.                                           |
| `Description`                                                     | *string*                                                          | :heavy_check_mark:                                                | The description of the folder.                                    |
| `Type`                                                            | [components.Type](../../models/components/type.md)                | :heavy_check_mark:                                                | N/A                                                               |
| `AccessLevel`                                                     | [*components.AccessLevel](../../models/components/accesslevel.md) | :heavy_minus_sign:                                                | The access level of the folder within the workspace.              |
| `CreatedAt`                                                       | *string*                                                          | :heavy_check_mark:                                                | The date the folder was created.                                  |
| `UpdatedAt`                                                       | *string*                                                          | :heavy_check_mark:                                                | The date the folder was updated.                                  |