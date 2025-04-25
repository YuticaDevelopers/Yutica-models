// models/properties/unit.go
package properties

import (
	"github.com/YuticaDevelopers/Yutica-models/models/common"
)

type Unit struct {
	common.BaseModel

	PropertyID uint   `gorm:"index" json:"property_id"`
	Name       string `gorm:"type:varchar(100);index" json:"name"` // e.g. Unit A1, Block 3F
	FloorNo    string `gorm:"type:varchar(20)" json:"floor_no"`
	Type       string `gorm:"type:varchar(50)" json:"type"` // optional: apartment, shop, office, etc.
	IsActive   bool   `json:"is_active"`
	Notes      string `json:"notes,omitempty"`
}

// 🏢 Unit is the physical sub-division within a property assigned to tenants
// 🔁 Indexed by property_id and name for fast lookup

/*
📦 Example Output (JSON):
{
  "id": 32,
  "property_id": 7,
  "name": "Block A4",
  "floor_no": "2nd Floor",
  "type": "apartment",
  "is_active": true,
  "created_at": "2025-04-01T10:00:00Z",
  "updated_at": "2025-04-01T10:00:00Z",
  "notes": "Corner unit with water meter installed"
}
*/
