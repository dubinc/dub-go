# Ids

IDs of tags to filter by.


## Supported Types

### 

```go
ids := operations.CreateIdsStr(string{/* values here */})
```

### 

```go
ids := operations.CreateIdsArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch ids.Type {
	case operations.IdsTypeStr:
		// ids.Str is populated
	case operations.IdsTypeArrayOfStr:
		// ids.ArrayOfStr is populated
}
```
