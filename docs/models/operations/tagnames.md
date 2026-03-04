# TagNames

The unique name of the tags assigned to the short link (case insensitive).


## Supported Types

### 

```go
tagNames := operations.CreateTagNamesStr(string{/* values here */})
```

### 

```go
tagNames := operations.CreateTagNamesArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch tagNames.Type {
	case operations.TagNamesTypeStr:
		// tagNames.Str is populated
	case operations.TagNamesTypeArrayOfStr:
		// tagNames.ArrayOfStr is populated
}
```
