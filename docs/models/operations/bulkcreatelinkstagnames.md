# BulkCreateLinksTagNames

The unique name of the tags assigned to the short link (case insensitive).


## Supported Types

### 

```go
bulkCreateLinksTagNames := operations.CreateBulkCreateLinksTagNamesStr(string{/* values here */})
```

### 

```go
bulkCreateLinksTagNames := operations.CreateBulkCreateLinksTagNamesArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch bulkCreateLinksTagNames.Type {
	case operations.BulkCreateLinksTagNamesTypeStr:
		// bulkCreateLinksTagNames.Str is populated
	case operations.BulkCreateLinksTagNamesTypeArrayOfStr:
		// bulkCreateLinksTagNames.ArrayOfStr is populated
}
```
