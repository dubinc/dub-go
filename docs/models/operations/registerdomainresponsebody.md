# RegisterDomainResponseBody

The domain was registered.


## Fields

| Field                                                                    | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `Domain`                                                                 | *string*                                                                 | :heavy_check_mark:                                                       | The domain name.                                                         |
| `Status`                                                                 | *string*                                                                 | :heavy_check_mark:                                                       | The status of the domain registration.                                   |
| `Expiration`                                                             | *float64*                                                                | :heavy_check_mark:                                                       | The expiration timestamp of the domain (Unix timestamp in milliseconds). |