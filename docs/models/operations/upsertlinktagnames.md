# UpsertLinkTagNames

The unique name of the tags assigned to the short link (case insensitive).


## Supported Types

### 

```go
upsertLinkTagNames := operations.CreateUpsertLinkTagNamesStr(string{/* values here */})
```

### 

```go
upsertLinkTagNames := operations.CreateUpsertLinkTagNamesArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch upsertLinkTagNames.Type {
	case operations.UpsertLinkTagNamesTypeStr:
		// upsertLinkTagNames.Str is populated
	case operations.UpsertLinkTagNamesTypeArrayOfStr:
		// upsertLinkTagNames.ArrayOfStr is populated
}
```
