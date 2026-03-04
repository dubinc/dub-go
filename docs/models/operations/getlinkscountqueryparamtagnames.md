# GetLinksCountQueryParamTagNames

The unique name of the tags assigned to the short link (case insensitive).


## Supported Types

### 

```go
getLinksCountQueryParamTagNames := operations.CreateGetLinksCountQueryParamTagNamesStr(string{/* values here */})
```

### 

```go
getLinksCountQueryParamTagNames := operations.CreateGetLinksCountQueryParamTagNamesArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch getLinksCountQueryParamTagNames.Type {
	case operations.GetLinksCountQueryParamTagNamesTypeStr:
		// getLinksCountQueryParamTagNames.Str is populated
	case operations.GetLinksCountQueryParamTagNamesTypeArrayOfStr:
		// getLinksCountQueryParamTagNames.ArrayOfStr is populated
}
```
