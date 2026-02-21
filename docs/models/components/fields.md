# Fields


## Supported Types

### Fields1

```go
fields := components.CreateFieldsFields1(components.Fields1{/* values here */})
```

### Fields2

```go
fields := components.CreateFieldsFields2(components.Fields2{/* values here */})
```

### Fields3

```go
fields := components.CreateFieldsFields3(components.Fields3{/* values here */})
```

### Four

```go
fields := components.CreateFieldsFour(components.Four{/* values here */})
```

### Five

```go
fields := components.CreateFieldsFive(components.Five{/* values here */})
```

### Six

```go
fields := components.CreateFieldsSix(components.Six{/* values here */})
```

### Seven

```go
fields := components.CreateFieldsSeven(components.Seven{/* values here */})
```

### Eight

```go
fields := components.CreateFieldsEight(components.Eight{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch fields.Type {
	case components.FieldsUnionTypeFields1:
		// fields.Fields1 is populated
	case components.FieldsUnionTypeFields2:
		// fields.Fields2 is populated
	case components.FieldsUnionTypeFields3:
		// fields.Fields3 is populated
	case components.FieldsUnionTypeFour:
		// fields.Four is populated
	case components.FieldsUnionTypeFive:
		// fields.Five is populated
	case components.FieldsUnionTypeSix:
		// fields.Six is populated
	case components.FieldsUnionTypeSeven:
		// fields.Seven is populated
	case components.FieldsUnionTypeEight:
		// fields.Eight is populated
}
```
