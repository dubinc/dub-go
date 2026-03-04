# CreatePartnerTagIds

The unique IDs of the tags assigned to the short link.


## Supported Types

### 

```go
createPartnerTagIds := operations.CreateCreatePartnerTagIdsStr(string{/* values here */})
```

### 

```go
createPartnerTagIds := operations.CreateCreatePartnerTagIdsArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch createPartnerTagIds.Type {
	case operations.CreatePartnerTagIdsTypeStr:
		// createPartnerTagIds.Str is populated
	case operations.CreatePartnerTagIdsTypeArrayOfStr:
		// createPartnerTagIds.ArrayOfStr is populated
}
```
