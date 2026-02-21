# UpdateLinkTagNames

The unique name of the tags assigned to the short link (case insensitive).


## Supported Types

### 

```go
updateLinkTagNames := operations.CreateUpdateLinkTagNamesStr(string{/* values here */})
```

### 

```go
updateLinkTagNames := operations.CreateUpdateLinkTagNamesArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch updateLinkTagNames.Type {
	case operations.UpdateLinkTagNamesTypeStr:
		// updateLinkTagNames.Str is populated
	case operations.UpdateLinkTagNamesTypeArrayOfStr:
		// updateLinkTagNames.ArrayOfStr is populated
}
```
