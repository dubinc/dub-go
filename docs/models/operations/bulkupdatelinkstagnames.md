# BulkUpdateLinksTagNames

The unique name of the tags assigned to the short link (case insensitive).


## Supported Types

### 

```go
bulkUpdateLinksTagNames := operations.CreateBulkUpdateLinksTagNamesStr(string{/* values here */})
```

### 

```go
bulkUpdateLinksTagNames := operations.CreateBulkUpdateLinksTagNamesArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch bulkUpdateLinksTagNames.Type {
	case operations.BulkUpdateLinksTagNamesTypeStr:
		// bulkUpdateLinksTagNames.Str is populated
	case operations.BulkUpdateLinksTagNamesTypeArrayOfStr:
		// bulkUpdateLinksTagNames.ArrayOfStr is populated
}
```
