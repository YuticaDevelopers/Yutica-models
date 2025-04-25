// models/usage/monthly.go
package usage

import (
	"time"

	"github.com/YuticaDevelopers/Yutica-models/models/common"
)

type MonthlyUsage struct {
	common.BaseModel

	MeterID    uint               `gorm:"index" json:"meter_id"`
	TenantID   uint               `gorm:"index" json:"tenant_id"`
	PropertyID uint               `gorm:"index" json:"property_id"`
	UnitID     uint               `gorm:"index" json:"unit_id"`
	Month      time.Time          `gorm:"index" json:"month"`                 // the day of measurement
	VolumeUsed float64            `json:"volume_used"`                        // true if actual reading was unavailable
	Type       common.UtilityType `gorm:"type:varchar(10);index" json:"type"` // water or power
}

// ⚠️ Expected to grow significantly (daily rows per meter)
// 🔁 Indexed by meter_id, tenant_id, usage_date, type
// 📦 Recommend partitioning by usage_date monthly
// 📄 Archive entries older than 12 months
