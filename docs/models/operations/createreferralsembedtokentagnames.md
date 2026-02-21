# CreateReferralsEmbedTokenTagNames

The unique name of the tags assigned to the short link (case insensitive).


## Supported Types

### 

```go
createReferralsEmbedTokenTagNames := operations.CreateCreateReferralsEmbedTokenTagNamesStr(string{/* values here */})
```

### 

```go
createReferralsEmbedTokenTagNames := operations.CreateCreateReferralsEmbedTokenTagNamesArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch createReferralsEmbedTokenTagNames.Type {
	case operations.CreateReferralsEmbedTokenTagNamesTypeStr:
		// createReferralsEmbedTokenTagNames.Str is populated
	case operations.CreateReferralsEmbedTokenTagNamesTypeArrayOfStr:
		// createReferralsEmbedTokenTagNames.ArrayOfStr is populated
}
```
