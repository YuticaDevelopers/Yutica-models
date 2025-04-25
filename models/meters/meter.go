// models/meters/meter.go
package meters

import (
	"time"

	"github.com/YuticaDevelopers/Yutica-models/models/common"
)

type Meter struct {
	common.BaseModel

	SerialNumber   string    `gorm:"type:varchar(100);uniqueIndex" json:"serial_number"`
	Type           string    `gorm:"type:varchar(20);index" json:"type"` // water or power
	ManufacturerID uint      `gorm:"index" json:"manufacturer_id"`
	PropertyID     uint      `gorm:"index" json:"property_id"`
	UnitID         uint      `gorm:"index" json:"unit_id"`
	CurrentReading float64   `json:"current_reading"`
	LastUpdated    time.Time `json:"last_updated"`
	IsActive       bool      `json:"is_active"`
	InstalledAt    time.Time `json:"installed_at"`
	Notes          string    `json:"notes,omitempty"`
}

// 🔌 Meter represents a utility device for measuring consumption
// 🔁 Indexed by serial, type, and property/unit for traceability

/*
📦 Example Output (JSON):
{
  "id": 501,
  "serial_number": "WMTR-1001-AZ",
  "type": "water",
  "manufacturer_id": 3,
  "property_id": 7,
  "unit_id": 15,
  "current_reading": 872.45,
  "last_updated": "2025-04-20T10:00:00Z",
  "is_active": true,
  "installed_at": "2024-11-01T08:00:00Z",
  "notes": "Outdoor ground-level installation"
}
*/
