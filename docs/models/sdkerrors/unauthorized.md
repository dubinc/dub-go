# Unauthorized

Although the HTTP standard specifies "unauthorized", semantically this response means "unauthenticated". That is, the client must authenticate itself to get the requested response.


## Fields

| Field                                                                      | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `Error`                                                                    | [sdkerrors.UnauthorizedError](../../models/sdkerrors/unauthorizederror.md) | :heavy_check_mark:                                                         | N/A                                                                        |