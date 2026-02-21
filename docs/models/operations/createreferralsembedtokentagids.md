# CreateReferralsEmbedTokenTagIds

The unique IDs of the tags assigned to the short link.


## Supported Types

### 

```go
createReferralsEmbedTokenTagIds := operations.CreateCreateReferralsEmbedTokenTagIdsStr(string{/* values here */})
```

### 

```go
createReferralsEmbedTokenTagIds := operations.CreateCreateReferralsEmbedTokenTagIdsArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch createReferralsEmbedTokenTagIds.Type {
	case operations.CreateReferralsEmbedTokenTagIdsTypeStr:
		// createReferralsEmbedTokenTagIds.Str is populated
	case operations.CreateReferralsEmbedTokenTagIdsTypeArrayOfStr:
		// createReferralsEmbedTokenTagIds.ArrayOfStr is populated
}
```
