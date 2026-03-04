# CreatePartnerFields


## Supported Types

### CreatePartnerFields1

```go
createPartnerFields := operations.CreateCreatePartnerFieldsCreatePartnerFields1(operations.CreatePartnerFields1{/* values here */})
```

### CreatePartnerFields2

```go
createPartnerFields := operations.CreateCreatePartnerFieldsCreatePartnerFields2(operations.CreatePartnerFields2{/* values here */})
```

### CreatePartnerFields3

```go
createPartnerFields := operations.CreateCreatePartnerFieldsCreatePartnerFields3(operations.CreatePartnerFields3{/* values here */})
```

### CreatePartnerFields4

```go
createPartnerFields := operations.CreateCreatePartnerFieldsCreatePartnerFields4(operations.CreatePartnerFields4{/* values here */})
```

### Fields5

```go
createPartnerFields := operations.CreateCreatePartnerFieldsFields5(operations.Fields5{/* values here */})
```

### Fields6

```go
createPartnerFields := operations.CreateCreatePartnerFieldsFields6(operations.Fields6{/* values here */})
```

### Fields7

```go
createPartnerFields := operations.CreateCreatePartnerFieldsFields7(operations.Fields7{/* values here */})
```

### Fields8

```go
createPartnerFields := operations.CreateCreatePartnerFieldsFields8(operations.Fields8{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch createPartnerFields.Type {
	case operations.CreatePartnerFieldsUnionTypeCreatePartnerFields1:
		// createPartnerFields.CreatePartnerFields1 is populated
	case operations.CreatePartnerFieldsUnionTypeCreatePartnerFields2:
		// createPartnerFields.CreatePartnerFields2 is populated
	case operations.CreatePartnerFieldsUnionTypeCreatePartnerFields3:
		// createPartnerFields.CreatePartnerFields3 is populated
	case operations.CreatePartnerFieldsUnionTypeCreatePartnerFields4:
		// createPartnerFields.CreatePartnerFields4 is populated
	case operations.CreatePartnerFieldsUnionTypeFields5:
		// createPartnerFields.Fields5 is populated
	case operations.CreatePartnerFieldsUnionTypeFields6:
		// createPartnerFields.Fields6 is populated
	case operations.CreatePartnerFieldsUnionTypeFields7:
		// createPartnerFields.Fields7 is populated
	case operations.CreatePartnerFieldsUnionTypeFields8:
		// createPartnerFields.Fields8 is populated
}
```
