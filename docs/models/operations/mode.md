# Mode

The mode to use for tracking the lead event. `async` will not block the request; `wait` will block the request until the lead event is fully recorded in Dub; `deferred` will defer the lead event creation to a subsequent request.


## Values

| Name           | Value          |
| -------------- | -------------- |
| `ModeAsync`    | async          |
| `ModeWait`     | wait           |
| `ModeDeferred` | deferred       |