# CreateCommissionRequestBody


## Supported Types

### RequestBody1

```go
createCommissionRequestBody := operations.CreateCreateCommissionRequestBodyRequestBody1(operations.RequestBody1{/* values here */})
```

### RequestBody2

```go
createCommissionRequestBody := operations.CreateCreateCommissionRequestBodyRequestBody2(operations.RequestBody2{/* values here */})
```

### RequestBody3

```go
createCommissionRequestBody := operations.CreateCreateCommissionRequestBodyRequestBody3(operations.RequestBody3{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch createCommissionRequestBody.Type {
	case operations.CreateCommissionRequestBodyUnionTypeRequestBody1:
		// createCommissionRequestBody.RequestBody1 is populated
	case operations.CreateCommissionRequestBodyUnionTypeRequestBody2:
		// createCommissionRequestBody.RequestBody2 is populated
	case operations.CreateCommissionRequestBodyUnionTypeRequestBody3:
		// createCommissionRequestBody.RequestBody3 is populated
}
```
