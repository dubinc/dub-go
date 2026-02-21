# ListEventsResponseBody


## Supported Types

### ClickEvent

```go
listEventsResponseBody := operations.CreateListEventsResponseBodyClickEvent(operations.ClickEvent{/* values here */})
```

### LeadEvent

```go
listEventsResponseBody := operations.CreateListEventsResponseBodyLeadEvent(operations.LeadEvent{/* values here */})
```

### SaleEvent

```go
listEventsResponseBody := operations.CreateListEventsResponseBodySaleEvent(operations.SaleEvent{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch listEventsResponseBody.Type {
	case operations.ListEventsResponseBodyTypeClickEvent:
		// listEventsResponseBody.ClickEvent is populated
	case operations.ListEventsResponseBodyTypeLeadEvent:
		// listEventsResponseBody.LeadEvent is populated
	case operations.ListEventsResponseBodyTypeSaleEvent:
		// listEventsResponseBody.SaleEvent is populated
}
```
