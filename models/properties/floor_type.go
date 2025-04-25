// models/properties/floor_type.go
package properties

import "github.com/YuticaDevelopers/Yutica-models/models/common"

type FloorType struct {
	common.BaseModel
	PropertyID uint   `gorm:"index" json:"property_id"`
	Name       string `gorm:"type:varchar(50);uniqueIndex" json:"name"` // e.g., Ground Floor, 1st Floor
	Position   int    `gorm:"index" json:"position"`                    // ordinal level for sorting
	Notes      string `json:"notes,omitempty"`
}

// 🏢 Represents predefined floor levels in a building for standardized referencing

/*
📦 Example Output (JSON):
{
  "id": 1,
  "name": "1st Floor",
  "property_id":1
  "position": 1,
  "notes": "Main lobby"
}
*/
