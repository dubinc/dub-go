# WebhookEvent

Webhook event schema


## Supported Types

### LinkWebhookEvent

```go
webhookEvent := components.CreateWebhookEventLinkWebhookEvent(components.LinkWebhookEvent{/* values here */})
```

### LinkClickedEvent

```go
webhookEvent := components.CreateWebhookEventLinkClickedEvent(components.LinkClickedEvent{/* values here */})
```

### LeadCreatedEvent

```go
webhookEvent := components.CreateWebhookEventLeadCreatedEvent(components.LeadCreatedEvent{/* values here */})
```

### SaleCreatedEvent

```go
webhookEvent := components.CreateWebhookEventSaleCreatedEvent(components.SaleCreatedEvent{/* values here */})
```

### PartnerEnrolledEvent

```go
webhookEvent := components.CreateWebhookEventPartnerEnrolledEvent(components.PartnerEnrolledEvent{/* values here */})
```

### PartnerApplicationSubmittedEvent

```go
webhookEvent := components.CreateWebhookEventPartnerApplicationSubmittedEvent(components.PartnerApplicationSubmittedEvent{/* values here */})
```

### PartnerMergedEvent

```go
webhookEvent := components.CreateWebhookEventPartnerMergedEvent(components.PartnerMergedEvent{/* values here */})
```

### CommissionCreatedEvent

```go
webhookEvent := components.CreateWebhookEventCommissionCreatedEvent(components.CommissionCreatedEvent{/* values here */})
```

### DiscountCodeWebhookEvent

```go
webhookEvent := components.CreateWebhookEventDiscountCodeWebhookEvent(components.DiscountCodeWebhookEvent{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch webhookEvent.Type {
	case components.WebhookEventTypeLinkWebhookEvent:
		// webhookEvent.LinkWebhookEvent is populated
	case components.WebhookEventTypeLinkClickedEvent:
		// webhookEvent.LinkClickedEvent is populated
	case components.WebhookEventTypeLeadCreatedEvent:
		// webhookEvent.LeadCreatedEvent is populated
	case components.WebhookEventTypeSaleCreatedEvent:
		// webhookEvent.SaleCreatedEvent is populated
	case components.WebhookEventTypePartnerEnrolledEvent:
		// webhookEvent.PartnerEnrolledEvent is populated
	case components.WebhookEventTypePartnerApplicationSubmittedEvent:
		// webhookEvent.PartnerApplicationSubmittedEvent is populated
	case components.WebhookEventTypePartnerMergedEvent:
		// webhookEvent.PartnerMergedEvent is populated
	case components.WebhookEventTypeCommissionCreatedEvent:
		// webhookEvent.CommissionCreatedEvent is populated
	case components.WebhookEventTypeDiscountCodeWebhookEvent:
		// webhookEvent.DiscountCodeWebhookEvent is populated
}
```
