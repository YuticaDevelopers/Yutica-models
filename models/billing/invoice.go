// models/billing/invoice.go
package billing

import (
	"time"

	"github.com/YuticaDevelopers/Yutica-models/models/common"
)

type Invoice struct {
	common.BaseModel

	TenantID    uint               `gorm:"index" json:"tenant_id"`
	PropertyID  uint               `gorm:"index" json:"property_id"`
	UnitID      uint               `gorm:"index" json:"unit_id"`
	Type        common.UtilityType `gorm:"type:varchar(10);index" json:"type"` // water or power
	InvoiceDate time.Time          `gorm:"index" json:"invoice_date"`
	DueDate     time.Time          `json:"due_date"`

	PreviousReading float64 `json:"previous_reading"`
	FinalReading    float64 `json:"final_reading"`
	Consumption     float64 `json:"consumption"` // meters consumed (liters or kWh)
	RatePerUnit     float64 `json:"rate_per_unit"`
	Tax             float64 `json:"tax"`
	Amount          float64 `json:"amount"`

	ConsumptionStart time.Time `json:"consumption_start_date"`
	ConsumptionEnd   time.Time `json:"consumption_end_date"`

	Breakdown string               `json:"breakdown"` // optional JSON or text detail
	Status    common.InvoiceStatus `gorm:"type:varchar(20);index" json:"status"`
	Reference string               `gorm:"uniqueIndex" json:"reference"`
	PaidAt    *time.Time           `json:"paid_at,omitempty"`
	Notes     string               `json:"notes,omitempty"`
}

// ⚠️ Grows monthly by number of tenants
// 🔁 Indexed by tenant_id, invoice_date, status
// 📦 Partitioning optional — consider for properties with 10K+ tenants
// 📄 Retain invoices for 7 years (regulatory)
// 💡 Best practice: balance brought forward and paid amounts should be calculated from debit and credit transactions, not stored in the invoice.

/*
📦 Example Output (JSON):
{
  "id": 202504251001,
  "tenant_id": 1,
  "property_id": 5,
  "unit_id": 12,
  "type": "water",
  "invoice_date": "2025-04-01T00:00:00Z",
  "due_date": "2025-04-15T00:00:00Z",
  "previous_reading": 3000,
  "final_reading": 4100,
  "consumption": 1100,
  "rate_per_unit": 1.2,
  "tax": 60.50,
  "amount": 1320.50,
  "consumption_start_date": "2025-03-01T00:00:00Z",
  "consumption_end_date": "2025-03-31T00:00:00Z",
  "breakdown": "{\"base_charge\": 1000, \"tax\": 60.5}",
  "status": "pending",
  "reference": "INV-APR-202504251001",
  "created_at": "2025-04-01T10:00:00Z",
  "updated_at": "2025-04-01T10:00:00Z",
  "notes": "First invoice for April 2025"
}
*/
