# RetrieveAnalyticsResponseBody

Analytics data


## Supported Types

### AnalyticsCount

```go
retrieveAnalyticsResponseBody := operations.CreateRetrieveAnalyticsResponseBodyAnalyticsCount(components.AnalyticsCount{/* values here */})
```

### 

```go
retrieveAnalyticsResponseBody := operations.CreateRetrieveAnalyticsResponseBodyArrayOfAnalyticsTimeseries([]components.AnalyticsTimeseries{/* values here */})
```

### 

```go
retrieveAnalyticsResponseBody := operations.CreateRetrieveAnalyticsResponseBodyArrayOfAnalyticsContinents([]components.AnalyticsContinents{/* values here */})
```

### 

```go
retrieveAnalyticsResponseBody := operations.CreateRetrieveAnalyticsResponseBodyArrayOfAnalyticsCountries([]components.AnalyticsCountries{/* values here */})
```

### 

```go
retrieveAnalyticsResponseBody := operations.CreateRetrieveAnalyticsResponseBodyArrayOfAnalyticsRegions([]components.AnalyticsRegions{/* values here */})
```

### 

```go
retrieveAnalyticsResponseBody := operations.CreateRetrieveAnalyticsResponseBodyArrayOfAnalyticsCities([]components.AnalyticsCities{/* values here */})
```

### 

```go
retrieveAnalyticsResponseBody := operations.CreateRetrieveAnalyticsResponseBodyArrayOfAnalyticsDevices([]components.AnalyticsDevices{/* values here */})
```

### 

```go
retrieveAnalyticsResponseBody := operations.CreateRetrieveAnalyticsResponseBodyArrayOfAnalyticsBrowsers([]components.AnalyticsBrowsers{/* values here */})
```

### 

```go
retrieveAnalyticsResponseBody := operations.CreateRetrieveAnalyticsResponseBodyArrayOfAnalyticsOS([]components.AnalyticsOS{/* values here */})
```

### 

```go
retrieveAnalyticsResponseBody := operations.CreateRetrieveAnalyticsResponseBodyArrayOfAnalyticsTriggers([]components.AnalyticsTriggers{/* values here */})
```

### 

```go
retrieveAnalyticsResponseBody := operations.CreateRetrieveAnalyticsResponseBodyArrayOfAnalyticsReferers([]components.AnalyticsReferers{/* values here */})
```

### 

```go
retrieveAnalyticsResponseBody := operations.CreateRetrieveAnalyticsResponseBodyArrayOfAnalyticsRefererUrls([]components.AnalyticsRefererUrls{/* values here */})
```

### 

```go
retrieveAnalyticsResponseBody := operations.CreateRetrieveAnalyticsResponseBodyArrayOfAnalyticsTopLinks([]components.AnalyticsTopLinks{/* values here */})
```

### 

```go
retrieveAnalyticsResponseBody := operations.CreateRetrieveAnalyticsResponseBodyArrayOfAnalyticsTopUrls([]components.AnalyticsTopUrls{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch retrieveAnalyticsResponseBody.Type {
	case operations.RetrieveAnalyticsResponseBodyTypeAnalyticsCount:
		// retrieveAnalyticsResponseBody.AnalyticsCount is populated
	case operations.RetrieveAnalyticsResponseBodyTypeArrayOfAnalyticsTimeseries:
		// retrieveAnalyticsResponseBody.ArrayOfAnalyticsTimeseries is populated
	case operations.RetrieveAnalyticsResponseBodyTypeArrayOfAnalyticsContinents:
		// retrieveAnalyticsResponseBody.ArrayOfAnalyticsContinents is populated
	case operations.RetrieveAnalyticsResponseBodyTypeArrayOfAnalyticsCountries:
		// retrieveAnalyticsResponseBody.ArrayOfAnalyticsCountries is populated
	case operations.RetrieveAnalyticsResponseBodyTypeArrayOfAnalyticsRegions:
		// retrieveAnalyticsResponseBody.ArrayOfAnalyticsRegions is populated
	case operations.RetrieveAnalyticsResponseBodyTypeArrayOfAnalyticsCities:
		// retrieveAnalyticsResponseBody.ArrayOfAnalyticsCities is populated
	case operations.RetrieveAnalyticsResponseBodyTypeArrayOfAnalyticsDevices:
		// retrieveAnalyticsResponseBody.ArrayOfAnalyticsDevices is populated
	case operations.RetrieveAnalyticsResponseBodyTypeArrayOfAnalyticsBrowsers:
		// retrieveAnalyticsResponseBody.ArrayOfAnalyticsBrowsers is populated
	case operations.RetrieveAnalyticsResponseBodyTypeArrayOfAnalyticsOS:
		// retrieveAnalyticsResponseBody.ArrayOfAnalyticsOS is populated
	case operations.RetrieveAnalyticsResponseBodyTypeArrayOfAnalyticsTriggers:
		// retrieveAnalyticsResponseBody.ArrayOfAnalyticsTriggers is populated
	case operations.RetrieveAnalyticsResponseBodyTypeArrayOfAnalyticsReferers:
		// retrieveAnalyticsResponseBody.ArrayOfAnalyticsReferers is populated
	case operations.RetrieveAnalyticsResponseBodyTypeArrayOfAnalyticsRefererUrls:
		// retrieveAnalyticsResponseBody.ArrayOfAnalyticsRefererUrls is populated
	case operations.RetrieveAnalyticsResponseBodyTypeArrayOfAnalyticsTopLinks:
		// retrieveAnalyticsResponseBody.ArrayOfAnalyticsTopLinks is populated
	case operations.RetrieveAnalyticsResponseBodyTypeArrayOfAnalyticsTopUrls:
		// retrieveAnalyticsResponseBody.ArrayOfAnalyticsTopUrls is populated
}
```
