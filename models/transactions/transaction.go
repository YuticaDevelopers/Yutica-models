// models/transactions/transaction.go
package transactions

import (
	"github.com/YuticaDevelopers/Yutica-models/models/common"
)

type Transaction struct {
	common.BaseModel

	TenantID   uint                     `gorm:"index" json:"tenant_id"`
	MeterID    uint                     `gorm:"index" json:"meter_id"`
	PropertyID uint                     `gorm:"index" json:"property_id"`
	UnitID     uint                     `gorm:"index" json:"unit_id"`
	Amount     float64                  `json:"amount"`
	Type       common.TransactionType   `gorm:"type:varchar(20);index" json:"type"`
	Status     common.TransactionStatus `gorm:"type:varchar(20);index" json:"status"`
	Reference  string                   `gorm:"uniqueIndex" json:"reference"` // e.g. M-Pesa code

}

// ⚠️ This table is expected to grow rapidly (millions of rows)
// 🔁 Indexed by tenant_id, meter_id, created_at, status
// 📦 Recommend partitioning by created_at (monthly)
// 📄 Archive entries older than 24 months to cold storage

//  Partition Reminder
