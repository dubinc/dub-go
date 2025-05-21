# Status

Useful for marking a commission as refunded, duplicate, canceled, or fraudulent. Takes precedence over `amount` and `modifyAmount`. When a commission is marked as refunded, duplicate, canceled, or fraudulent, it will be omitted from the payout, and the payout amount will be recalculated accordingly. Paid commissions cannot be updated.


## Values

| Name              | Value             |
| ----------------- | ----------------- |
| `StatusRefunded`  | refunded          |
| `StatusDuplicate` | duplicate         |
| `StatusCanceled`  | canceled          |
| `StatusFraud`     | fraud             |