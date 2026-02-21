# GetLinksCountQueryParamTagIds

The tag IDs to filter the links by.


## Supported Types

### 

```go
getLinksCountQueryParamTagIds := operations.CreateGetLinksCountQueryParamTagIdsStr(string{/* values here */})
```

### 

```go
getLinksCountQueryParamTagIds := operations.CreateGetLinksCountQueryParamTagIdsArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch getLinksCountQueryParamTagIds.Type {
	case operations.GetLinksCountQueryParamTagIdsTypeStr:
		// getLinksCountQueryParamTagIds.Str is populated
	case operations.GetLinksCountQueryParamTagIdsTypeArrayOfStr:
		// getLinksCountQueryParamTagIds.ArrayOfStr is populated
}
```
