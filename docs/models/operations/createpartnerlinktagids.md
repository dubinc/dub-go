# CreatePartnerLinkTagIds

The unique IDs of the tags assigned to the short link.


## Supported Types

### 

```go
createPartnerLinkTagIds := operations.CreateCreatePartnerLinkTagIdsStr(string{/* values here */})
```

### 

```go
createPartnerLinkTagIds := operations.CreateCreatePartnerLinkTagIdsArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch createPartnerLinkTagIds.Type {
	case operations.CreatePartnerLinkTagIdsTypeStr:
		// createPartnerLinkTagIds.Str is populated
	case operations.CreatePartnerLinkTagIdsTypeArrayOfStr:
		// createPartnerLinkTagIds.ArrayOfStr is populated
}
```
