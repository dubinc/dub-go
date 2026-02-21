# ResponseBody


## Supported Types

### LinkSchema

```go
responseBody := operations.CreateResponseBodyLinkSchema(components.LinkSchema{/* values here */})
```

### LinkErrorSchema

```go
responseBody := operations.CreateResponseBodyLinkErrorSchema(components.LinkErrorSchema{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch responseBody.Type {
	case operations.ResponseBodyTypeLinkSchema:
		// responseBody.LinkSchema is populated
	case operations.ResponseBodyTypeLinkErrorSchema:
		// responseBody.LinkErrorSchema is populated
}
```
