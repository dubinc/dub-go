# UpsertPartnerLinkTagIds

The unique IDs of the tags assigned to the short link.


## Supported Types

### 

```go
upsertPartnerLinkTagIds := operations.CreateUpsertPartnerLinkTagIdsStr(string{/* values here */})
```

### 

```go
upsertPartnerLinkTagIds := operations.CreateUpsertPartnerLinkTagIdsArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch upsertPartnerLinkTagIds.Type {
	case operations.UpsertPartnerLinkTagIdsTypeStr:
		// upsertPartnerLinkTagIds.Str is populated
	case operations.UpsertPartnerLinkTagIdsTypeArrayOfStr:
		// upsertPartnerLinkTagIds.ArrayOfStr is populated
}
```
