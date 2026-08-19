# DiscountCodeWebhookEventEvent


## Supported Types

### Event1

```go
discountCodeWebhookEventEvent := components.CreateDiscountCodeWebhookEventEventEvent1(components.Event1{/* values here */})
```

### Event2

```go
discountCodeWebhookEventEvent := components.CreateDiscountCodeWebhookEventEventEvent2(components.Event2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch discountCodeWebhookEventEvent.Type {
	case components.DiscountCodeWebhookEventEventTypeEvent1:
		// discountCodeWebhookEventEvent.Event1 is populated
	case components.DiscountCodeWebhookEventEventTypeEvent2:
		// discountCodeWebhookEventEvent.Event2 is populated
}
```
