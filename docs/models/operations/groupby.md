# GroupBy

The field to group the links by.


## Supported Types

### One

```go
groupBy := operations.CreateGroupByOne(operations.One{/* values here */})
```

### Two

```go
groupBy := operations.CreateGroupByTwo(operations.Two{/* values here */})
```

### Three

```go
groupBy := operations.CreateGroupByThree(operations.Three{/* values here */})
```

### Four

```go
groupBy := operations.CreateGroupByFour(operations.Four{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch groupBy.Type {
	case operations.GroupByTypeOne:
		// groupBy.One is populated
	case operations.GroupByTypeTwo:
		// groupBy.Two is populated
	case operations.GroupByTypeThree:
		// groupBy.Three is populated
	case operations.GroupByTypeFour:
		// groupBy.Four is populated
}
```
