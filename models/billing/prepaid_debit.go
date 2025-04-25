// models/billing/prepaid_debit.go
package billing

import (
	"time"

	"github.com/YuticaDevelopers/Yutica-models/models/common"
)

type PrepaidDebit struct {
	common.BaseModel

	BillingAccountID uint                      `gorm:"index" json:"billing_account_id"`
	TenantID         uint                      `gorm:"index" json:"tenant_id"`
	Amount           float64                   `json:"amount"`
	Reason           common.PrepaidDebitReason `gorm:"type:varchar(20);index" json:"reason"`
	Reference        string                    `gorm:"index" json:"reference"` // linked invoice or reading
	DebitedAt        time.Time                 `json:"debited_at"`

	PreviousValue float64 `json:"previous_value"` // e.g., previous meter reading
	CurrentValue  float64 `json:"current_value"`  // current meter reading
	Usage         float64 `json:"usage"`          // derived usage = current - previous
	RatePerUnit   float64 `json:"rate_per_unit"`  // rate applied
	Unit          string  `json:"unit"`           // e.g., kWh, litres, KES
	Breakdown     string  `json:"breakdown"`      // optional JSON (e.g., {"unit_cost":150,"tax":5})

	Notes string `json:"notes,omitempty"`
}

// ⚠️ Grows with every utility usage, billing event, or manual deduction
// 🔁 Indexed by billing_account_id, tenant_id, reason, reference
// 🔐 Used to compute billing account balance and usage history
// 💡 If a debit was entered incorrectly, log a new entry with reason = "reversal" and the same amount to offset it.

/*
📦 Example Output (JSON):
{
  "id": 402,
  "billing_account_id": 101,
  "tenant_id": 5,
  "amount": 145.00,
  "reason": "usage",
  "reference": "METER-APR-20250425",
  "debited_at": "2025-04-25T07:00:00Z",
  "previous_value": 800.0,
  "current_value": 950.0,
  "usage": 150.0,
  "rate_per_unit": 1.0,
  "unit": "litres",
  "breakdown": "{\"unit_cost\": 150, \"tax\": 5}",
  "created_at": "2025-04-25T07:00:01Z",
  "updated_at": "2025-04-25T07:00:01Z",
  "notes": "Water usage for April"
}
*/
