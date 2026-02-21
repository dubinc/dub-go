# BulkCreateLinksTagIds

The unique IDs of the tags assigned to the short link.


## Supported Types

### 

```go
bulkCreateLinksTagIds := operations.CreateBulkCreateLinksTagIdsStr(string{/* values here */})
```

### 

```go
bulkCreateLinksTagIds := operations.CreateBulkCreateLinksTagIdsArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch bulkCreateLinksTagIds.Type {
	case operations.BulkCreateLinksTagIdsTypeStr:
		// bulkCreateLinksTagIds.Str is populated
	case operations.BulkCreateLinksTagIdsTypeArrayOfStr:
		// bulkCreateLinksTagIds.ArrayOfStr is populated
}
```
