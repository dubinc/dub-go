# ListEventsRequest


## Fields

| Field                                                                                                                    | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              | Example                                                                                                                  |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `Event`                                                                                                                  | [*operations.QueryParamEvent](../../models/operations/queryparamevent.md)                                                | :heavy_minus_sign:                                                                                                       | The type of event to retrieve analytics for. Defaults to 'clicks'.                                                       |                                                                                                                          |
| `Domain`                                                                                                                 | **string*                                                                                                                | :heavy_minus_sign:                                                                                                       | The domain to filter analytics for.                                                                                      |                                                                                                                          |
| `Key`                                                                                                                    | **string*                                                                                                                | :heavy_minus_sign:                                                                                                       | The short link slug.                                                                                                     |                                                                                                                          |
| `LinkID`                                                                                                                 | **string*                                                                                                                | :heavy_minus_sign:                                                                                                       | The unique ID of the short link on Dub.                                                                                  |                                                                                                                          |
| `ExternalID`                                                                                                             | **string*                                                                                                                | :heavy_minus_sign:                                                                                                       | This is the ID of the link in the your database. Must be prefixed with 'ext_' when passed as a query parameter.          |                                                                                                                          |
| `Interval`                                                                                                               | [*operations.QueryParamInterval](../../models/operations/queryparaminterval.md)                                          | :heavy_minus_sign:                                                                                                       | The interval to retrieve analytics for. Takes precedence over start and end. If undefined, defaults to 24h.              |                                                                                                                          |
| `Start`                                                                                                                  | **string*                                                                                                                | :heavy_minus_sign:                                                                                                       | The start date and time when to retrieve analytics from.                                                                 |                                                                                                                          |
| `End`                                                                                                                    | **string*                                                                                                                | :heavy_minus_sign:                                                                                                       | The end date and time when to retrieve analytics from. If not provided, defaults to the current date.                    |                                                                                                                          |
| `Timezone`                                                                                                               | **string*                                                                                                                | :heavy_minus_sign:                                                                                                       | The IANA time zone code for aligning timeseries granularity (e.g. America/New_York). Defaults to UTC.                    | America/New_York                                                                                                         |
| `Continent`                                                                                                              | [*components.ContinentCode](../../models/components/continentcode.md)                                                    | :heavy_minus_sign:                                                                                                       | The continent to retrieve analytics for.                                                                                 |                                                                                                                          |
| `Country`                                                                                                                | [*components.CountryCode](../../models/components/countrycode.md)                                                        | :heavy_minus_sign:                                                                                                       | The country to retrieve analytics for.                                                                                   |                                                                                                                          |
| `City`                                                                                                                   | **string*                                                                                                                | :heavy_minus_sign:                                                                                                       | The city to retrieve analytics for.                                                                                      | New York                                                                                                                 |
| `Device`                                                                                                                 | **string*                                                                                                                | :heavy_minus_sign:                                                                                                       | The device to retrieve analytics for.                                                                                    | Desktop                                                                                                                  |
| `Browser`                                                                                                                | **string*                                                                                                                | :heavy_minus_sign:                                                                                                       | The browser to retrieve analytics for.                                                                                   | Chrome                                                                                                                   |
| `Os`                                                                                                                     | **string*                                                                                                                | :heavy_minus_sign:                                                                                                       | The OS to retrieve analytics for.                                                                                        | Windows                                                                                                                  |
| `Referer`                                                                                                                | **string*                                                                                                                | :heavy_minus_sign:                                                                                                       | The referer to retrieve analytics for.                                                                                   | google.com                                                                                                               |
| `URL`                                                                                                                    | **string*                                                                                                                | :heavy_minus_sign:                                                                                                       | The URL to retrieve analytics for.                                                                                       |                                                                                                                          |
| `TagID`                                                                                                                  | **string*                                                                                                                | :heavy_minus_sign:                                                                                                       | The tag ID to retrieve analytics for.                                                                                    |                                                                                                                          |
| `Qr`                                                                                                                     | **bool*                                                                                                                  | :heavy_minus_sign:                                                                                                       | Filter for QR code scans. If true, filter for QR codes only. If false, filter for links only. If undefined, return both. |                                                                                                                          |
| `Root`                                                                                                                   | **bool*                                                                                                                  | :heavy_minus_sign:                                                                                                       | Filter for root domains. If true, filter for domains only. If false, filter for links only. If undefined, return both.   |                                                                                                                          |
| `Page`                                                                                                                   | **float64*                                                                                                               | :heavy_minus_sign:                                                                                                       | N/A                                                                                                                      |                                                                                                                          |
| `Limit`                                                                                                                  | **float64*                                                                                                               | :heavy_minus_sign:                                                                                                       | N/A                                                                                                                      |                                                                                                                          |
| `Order`                                                                                                                  | [*operations.Order](../../models/operations/order.md)                                                                    | :heavy_minus_sign:                                                                                                       | N/A                                                                                                                      |                                                                                                                          |
| `SortBy`                                                                                                                 | [*operations.SortBy](../../models/operations/sortby.md)                                                                  | :heavy_minus_sign:                                                                                                       | N/A                                                                                                                      |                                                                                                                          |