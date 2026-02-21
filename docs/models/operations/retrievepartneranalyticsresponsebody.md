# RetrievePartnerAnalyticsResponseBody

Partner analytics data


## Supported Types

### PartnerAnalyticsCount

```go
retrievePartnerAnalyticsResponseBody := operations.CreateRetrievePartnerAnalyticsResponseBodyPartnerAnalyticsCount(components.PartnerAnalyticsCount{/* values here */})
```

### 

```go
retrievePartnerAnalyticsResponseBody := operations.CreateRetrievePartnerAnalyticsResponseBodyArrayOfPartnerAnalyticsTimeseries([]components.PartnerAnalyticsTimeseries{/* values here */})
```

### 

```go
retrievePartnerAnalyticsResponseBody := operations.CreateRetrievePartnerAnalyticsResponseBodyArrayOfPartnerAnalyticsTopLinks([]components.PartnerAnalyticsTopLinks{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch retrievePartnerAnalyticsResponseBody.Type {
	case operations.RetrievePartnerAnalyticsResponseBodyTypePartnerAnalyticsCount:
		// retrievePartnerAnalyticsResponseBody.PartnerAnalyticsCount is populated
	case operations.RetrievePartnerAnalyticsResponseBodyTypeArrayOfPartnerAnalyticsTimeseries:
		// retrievePartnerAnalyticsResponseBody.ArrayOfPartnerAnalyticsTimeseries is populated
	case operations.RetrievePartnerAnalyticsResponseBodyTypeArrayOfPartnerAnalyticsTopLinks:
		// retrievePartnerAnalyticsResponseBody.ArrayOfPartnerAnalyticsTopLinks is populated
}
```
