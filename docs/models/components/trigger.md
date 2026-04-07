# Trigger

The type of trigger method: link click or QR scan

## Example Usage

```go
import (
	"github.com/dubinc/dub-go/models/components"
)

value := components.TriggerQr
```


## Values

| Name              | Value             |
| ----------------- | ----------------- |
| `TriggerQr`       | qr                |
| `TriggerLink`     | link              |
| `TriggerPageview` | pageview          |
| `TriggerDeeplink` | deeplink          |