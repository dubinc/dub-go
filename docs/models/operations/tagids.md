# TagIds

The unique IDs of the tags assigned to the short link.


## Supported Types

### 

```go
tagIds := operations.CreateTagIdsStr(string{/* values here */})
```

### 

```go
tagIds := operations.CreateTagIdsArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch tagIds.Type {
	case operations.TagIdsTypeStr:
		// tagIds.Str is populated
	case operations.TagIdsTypeArrayOfStr:
		// tagIds.ArrayOfStr is populated
}
```
