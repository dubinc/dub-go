# Domains

The domains to search. We only support .link domains for now.


## Supported Types

### 

```go
domains := operations.CreateDomainsStr(string{/* values here */})
```

### 

```go
domains := operations.CreateDomainsArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch domains.Type {
	case operations.DomainsTypeStr:
		// domains.Str is populated
	case operations.DomainsTypeArrayOfStr:
		// domains.ArrayOfStr is populated
}
```
