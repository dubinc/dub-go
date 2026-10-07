# StripeInvoicesToImport

Import paid Stripe invoices for the customer and create a commission for each. Pass `all` to import every unimported, paid invoice, or an array of Stripe invoice IDs to import only those invoices. Refunded invoices are not imported. When not provided, create a single manual sale event using `sale.amount`


## Supported Types

### StripeInvoicesToImport1

```go
stripeInvoicesToImport := operations.CreateStripeInvoicesToImportStripeInvoicesToImport1(operations.StripeInvoicesToImport1{/* values here */})
```

### 

```go
stripeInvoicesToImport := operations.CreateStripeInvoicesToImportArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch stripeInvoicesToImport.Type {
	case operations.StripeInvoicesToImportTypeStripeInvoicesToImport1:
		// stripeInvoicesToImport.StripeInvoicesToImport1 is populated
	case operations.StripeInvoicesToImportTypeArrayOfStr:
		// stripeInvoicesToImport.ArrayOfStr is populated
}
```
