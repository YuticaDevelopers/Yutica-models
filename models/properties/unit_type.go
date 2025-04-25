// models/properties/unit_type.go
package properties

import "github.com/YuticaDevelopers/Yutica-models/models/common"

type UnitType struct {
	common.BaseModel
	PropertyID uint   `gorm:"index" json:"property_id"`
	Name       string `gorm:"type:varchar(50);uniqueIndex" json:"name"` // e.g., Apartment, Office
	Notes      string `json:"notes,omitempty"`
}

// 🧱 Represents standardized unit categories used across the system

/*
📦 Example Output (JSON):
{
  "id": 1,
  "name": "Apartment",
   "property_id":1,
  "notes": "Residential living space"
}
*/
