// models/billing/prepaid_credit.go
package billing

import (
	"time"

	"github.com/YuticaDevelopers/Yutica-models/models/common"
)

type PrepaidCredit struct {
	common.BaseModel

	BillingAccountID uint                `gorm:"index" json:"billing_account_id"`
	TenantID         uint                `gorm:"index" json:"tenant_id"`
	Amount           float64             `json:"amount"`
	Source           common.CreditSource `gorm:"type:varchar(20);index" json:"source"`
	Reference        string              `gorm:"index" json:"reference"` // transaction ID, M-Pesa ref, etc.
	CreditedAt       time.Time           `json:"credited_at"`
	Notes            string              `json:"notes,omitempty"`
}

// ⚠️ Grows with every top-up, refund, or wallet credit
// 🔁 Indexed by billing_account_id, tenant_id, reference
// 🔐 Used to compute billing account balance

/*
📦 Example Output (JSON):
{
  "id": 302,
  "billing_account_id": 101,
  "tenant_id": 5,
  "amount": 500.00,
  "source": "mpesa",
  "reference": "MPESA-KDJ7238G3J",
  "credited_at": "2025-04-25T08:30:00Z",
  "created_at": "2025-04-25T08:30:01Z",
  "updated_at": "2025-04-25T08:30:01Z",
  "notes": "Top-up via M-Pesa"
}
*/
