# UpdateLinkTagIds

The unique IDs of the tags assigned to the short link.


## Supported Types

### 

```go
updateLinkTagIds := operations.CreateUpdateLinkTagIdsStr(string{/* values here */})
```

### 

```go
updateLinkTagIds := operations.CreateUpdateLinkTagIdsArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch updateLinkTagIds.Type {
	case operations.UpdateLinkTagIdsTypeStr:
		// updateLinkTagIds.Str is populated
	case operations.UpdateLinkTagIdsTypeArrayOfStr:
		// updateLinkTagIds.ArrayOfStr is populated
}
```
