# QueryParamTagIds

The tag IDs to filter the links by.


## Supported Types

### 

```go
queryParamTagIds := operations.CreateQueryParamTagIdsStr(string{/* values here */})
```

### 

```go
queryParamTagIds := operations.CreateQueryParamTagIdsArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch queryParamTagIds.Type {
	case operations.QueryParamTagIdsTypeStr:
		// queryParamTagIds.Str is populated
	case operations.QueryParamTagIdsTypeArrayOfStr:
		// queryParamTagIds.ArrayOfStr is populated
}
```
