// models/billing/prepaid_negative_balance.go
package billing

import (
	"time"

	"github.com/YuticaDevelopers/Yutica-models/models/common"
)

type PrepaidNegativeBalance struct {
	common.BaseModel

	BillingAccountID uint                  `gorm:"index" json:"billing_account_id"`
	TenantID         uint                  `gorm:"index" json:"tenant_id"`
	DebitID          uint                  `gorm:"index" json:"debit_id"` // Reference to the prepaid debit that triggered the overdraft
	Amount           float64               `json:"amount"`                // Always positive (e.g., 50 = -50 actual balance)
	OccurredAt       time.Time             `json:"occurred_at"`
	Status           common.NegativeStatus `json:"status"`
	ResolvedAt       *time.Time            `json:"resolved_at,omitempty"`
	Notes            string                `json:"notes,omitempty"`
}

// ⚠️ This table logs overdraft events without affecting BillingAccount balance directly
// 🔁 Tracked independently to maintain integrity and enable enforcement
// ✅ Calculate balance from credits - debits, then write here if < 0 triggered by a specific debit
// 🔄 On resolution, update this record — don't delete it

/*
📦 Example Output (JSON):
{
  "id": 708,
  "billing_account_id": 101,
  "tenant_id": 5,
  "debit_id": 402,
  "amount": 75.00,
  "occurred_at": "2025-04-25T06:00:00Z",
  "status": 'resolved',
  "resolved_at": null,
  "notes": "Exceeded balance from water usage deduction"
}
*/
