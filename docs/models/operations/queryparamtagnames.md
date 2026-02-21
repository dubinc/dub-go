# QueryParamTagNames

The unique name of the tags assigned to the short link (case insensitive).


## Supported Types

### 

```go
queryParamTagNames := operations.CreateQueryParamTagNamesStr(string{/* values here */})
```

### 

```go
queryParamTagNames := operations.CreateQueryParamTagNamesArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch queryParamTagNames.Type {
	case operations.QueryParamTagNamesTypeStr:
		// queryParamTagNames.Str is populated
	case operations.QueryParamTagNamesTypeArrayOfStr:
		// queryParamTagNames.ArrayOfStr is populated
}
```
