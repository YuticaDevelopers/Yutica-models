// models/properties/property_type.go
package properties

import (
	"github.com/YuticaDevelopers/Yutica-models/models/common"
)

type PropertyType struct {
	common.BaseModel

	Name  string `gorm:"type:varchar(100);uniqueIndex" json:"name"` // e.g. Residential, Commercial
	Notes string `json:"notes,omitempty"`
}

// 🏘 PropertyType defines the usage classification of a property
// 🔁 Used for tagging properties as Residential, Commercial, Industrial, etc.

/*
📦 Example Output (JSON):
{
  "id": 2,
  "name": "Residential",
  "notes": "Used for housing tenants in apartment-style units"
}
*/
