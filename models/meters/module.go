// models/meters/module.go
package meters

import "github.com/YuticaDevelopers/Yutica-models/models/common"

type Module struct {
	common.BaseModel

	Name        string `gorm:"type:varchar(100);uniqueIndex" json:"name"` // e.g., Smart Water Meter
	Description string `gorm:"type:text" json:"description,omitempty"`
	Type        string `gorm:"type:varchar(50);index" json:"type"` // water or power
	IsActive    bool   `json:"is_active"`
	Notes       string `json:"notes,omitempty"`
}

// ⚙️ Module represents the category of metering device logic for water or power systems
// 🔁 Used for defining the application layer context behind meters

/*
📦 Example Output (JSON):
{
  "id": 1,
  "name": "Smart Water Meter",
  "description": "LoRaWAN-based device for remote water monitoring",
  "type": "water",
  "is_active": true,
  "notes": "Used in high-density residential settings"
}
*/
