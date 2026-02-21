# UpsertPartnerLinkTagNames

The unique name of the tags assigned to the short link (case insensitive).


## Supported Types

### 

```go
upsertPartnerLinkTagNames := operations.CreateUpsertPartnerLinkTagNamesStr(string{/* values here */})
```

### 

```go
upsertPartnerLinkTagNames := operations.CreateUpsertPartnerLinkTagNamesArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch upsertPartnerLinkTagNames.Type {
	case operations.UpsertPartnerLinkTagNamesTypeStr:
		// upsertPartnerLinkTagNames.Str is populated
	case operations.UpsertPartnerLinkTagNamesTypeArrayOfStr:
		// upsertPartnerLinkTagNames.ArrayOfStr is populated
}
```
