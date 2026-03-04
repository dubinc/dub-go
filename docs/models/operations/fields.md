# Fields


## Supported Types

### Fields1

```go
fields := operations.CreateFieldsFields1(operations.Fields1{/* values here */})
```

### Fields2

```go
fields := operations.CreateFieldsFields2(operations.Fields2{/* values here */})
```

### Fields3

```go
fields := operations.CreateFieldsFields3(operations.Fields3{/* values here */})
```

### Fields4

```go
fields := operations.CreateFieldsFields4(operations.Fields4{/* values here */})
```

### Five

```go
fields := operations.CreateFieldsFive(operations.Five{/* values here */})
```

### Six

```go
fields := operations.CreateFieldsSix(operations.Six{/* values here */})
```

### Seven

```go
fields := operations.CreateFieldsSeven(operations.Seven{/* values here */})
```

### Eight

```go
fields := operations.CreateFieldsEight(operations.Eight{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch fields.Type {
	case operations.FieldsUnionTypeFields1:
		// fields.Fields1 is populated
	case operations.FieldsUnionTypeFields2:
		// fields.Fields2 is populated
	case operations.FieldsUnionTypeFields3:
		// fields.Fields3 is populated
	case operations.FieldsUnionTypeFields4:
		// fields.Fields4 is populated
	case operations.FieldsUnionTypeFive:
		// fields.Five is populated
	case operations.FieldsUnionTypeSix:
		// fields.Six is populated
	case operations.FieldsUnionTypeSeven:
		// fields.Seven is populated
	case operations.FieldsUnionTypeEight:
		// fields.Eight is populated
}
```
