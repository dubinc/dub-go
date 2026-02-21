# BulkUpdateLinksTagIds

The unique IDs of the tags assigned to the short link.


## Supported Types

### 

```go
bulkUpdateLinksTagIds := operations.CreateBulkUpdateLinksTagIdsStr(string{/* values here */})
```

### 

```go
bulkUpdateLinksTagIds := operations.CreateBulkUpdateLinksTagIdsArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch bulkUpdateLinksTagIds.Type {
	case operations.BulkUpdateLinksTagIdsTypeStr:
		// bulkUpdateLinksTagIds.Str is populated
	case operations.BulkUpdateLinksTagIdsTypeArrayOfStr:
		// bulkUpdateLinksTagIds.ArrayOfStr is populated
}
```
