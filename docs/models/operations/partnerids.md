# PartnerIds

Only return the partners with these IDs (comma-separated, up to 100).


## Supported Types

### 

```go
partnerIds := operations.CreatePartnerIdsStr(string{/* values here */})
```

### 

```go
partnerIds := operations.CreatePartnerIdsArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch partnerIds.Type {
	case operations.PartnerIdsTypeStr:
		// partnerIds.Str is populated
	case operations.PartnerIdsTypeArrayOfStr:
		// partnerIds.ArrayOfStr is populated
}
```
