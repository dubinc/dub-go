# CreatePartnerLinkTagNames

The unique name of the tags assigned to the short link (case insensitive).


## Supported Types

### 

```go
createPartnerLinkTagNames := operations.CreateCreatePartnerLinkTagNamesStr(string{/* values here */})
```

### 

```go
createPartnerLinkTagNames := operations.CreateCreatePartnerLinkTagNamesArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch createPartnerLinkTagNames.Type {
	case operations.CreatePartnerLinkTagNamesTypeStr:
		// createPartnerLinkTagNames.Str is populated
	case operations.CreatePartnerLinkTagNamesTypeArrayOfStr:
		// createPartnerLinkTagNames.ArrayOfStr is populated
}
```
