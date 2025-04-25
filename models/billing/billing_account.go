// models/billing/billing_account.go
package billing

import (
	"time"

	"github.com/YuticaDevelopers/Yutica-models/models/common"
)

type BillingAccount struct {
	common.BaseModel

	TenantID   uint                      `gorm:"index" json:"tenant_id"`
	PropertyID uint                      `gorm:"index" json:"property_id"`
	UnitID     uint                      `gorm:"index" json:"unit_id"`
	Type       common.BillingAccountType `gorm:"type:varchar(10);index" json:"type"`       // water or power
	Mode       common.BillingMode        `gorm:"type:varchar(10);index" json:"mode"`       // prepaid or postpaid
	ValueMode  common.AccountValueMode   `gorm:"type:varchar(10);index" json:"value_mode"` // wallet or token based
	Unit       common.BalanceUnit        `gorm:"type:varchar(10);index" json:"unit"`       // KES, kWh, litres

	LastBilled *time.Time `json:"last_billed,omitempty"`
	IsActive   bool       `json:"is_active"`
	OpenedAt   time.Time  `json:"opened_at"`
	ClosedAt   *time.Time `json:"closed_at,omitempty"`
}

// ⚠️ Core financial record for tenants
// 🔁 Indexed by tenant_id, property_id, unit_id, type, mode, value_mode
// 🔐 Use this to associate transactions, credits, and invoices with logical billing buckets
// 💡 Best practice: DO NOT store balance directly in BillingAccount.
// 🔄 Instead, calculate it in real-time from related credit and debit transactions:
//     balance = SUM(credits) - SUM(debits)
// ✅ This ensures consistency, traceability, and supports multiple invoice settlements.

/*
📦 Example Output (JSON):
{
  "id": 101,
  "tenant_id": 5,
  "property_id": 3,
  "unit_id": 12,
  "type": "power",
  "mode": "prepaid",
  "value_mode": "wallet",
  "unit": "kWh",
  "last_billed": "2025-04-01T00:00:00Z",
  "is_active": true,
  "opened_at": "2024-10-01T00:00:00Z",
  "closed_at": null
}
*/
