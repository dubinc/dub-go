# CreatePartnerTagNames

The unique name of the tags assigned to the short link (case insensitive).


## Supported Types

### 

```go
createPartnerTagNames := operations.CreateCreatePartnerTagNamesStr(string{/* values here */})
```

### 

```go
createPartnerTagNames := operations.CreateCreatePartnerTagNamesArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch createPartnerTagNames.Type {
	case operations.CreatePartnerTagNamesTypeStr:
		// createPartnerTagNames.Str is populated
	case operations.CreatePartnerTagNamesTypeArrayOfStr:
		// createPartnerTagNames.ArrayOfStr is populated
}
```
