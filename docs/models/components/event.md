# Event


## Supported Types

### One

```go
event := components.CreateEventOne(components.One{/* values here */})
```

### Two

```go
event := components.CreateEventTwo(components.Two{/* values here */})
```

### Three

```go
event := components.CreateEventThree(components.Three{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch event.Type {
	case components.EventTypeOne:
		// event.One is populated
	case components.EventTypeTwo:
		// event.Two is populated
	case components.EventTypeThree:
		// event.Three is populated
}
```
