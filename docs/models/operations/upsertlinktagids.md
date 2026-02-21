# UpsertLinkTagIds

The unique IDs of the tags assigned to the short link.


## Supported Types

### 

```go
upsertLinkTagIds := operations.CreateUpsertLinkTagIdsStr(string{/* values here */})
```

### 

```go
upsertLinkTagIds := operations.CreateUpsertLinkTagIdsArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch upsertLinkTagIds.Type {
	case operations.UpsertLinkTagIdsTypeStr:
		// upsertLinkTagIds.Str is populated
	case operations.UpsertLinkTagIdsTypeArrayOfStr:
		// upsertLinkTagIds.ArrayOfStr is populated
}
```
