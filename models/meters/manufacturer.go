// models/meters/manufacturer.go
package meters

import (
	"github.com/YuticaDevelopers/Yutica-models/models/common"
)

type Manufacturer struct {
	common.BaseModel

	Name         string `gorm:"type:varchar(100);uniqueIndex" json:"name"`
	Country      string `gorm:"type:varchar(50)" json:"country"`
	ContactName  string `gorm:"type:varchar(100)" json:"contact_name"`
	ContactEmail string `gorm:"type:varchar(100)" json:"contact_email"`
	Phone        string `gorm:"type:varchar(20)" json:"phone"`
	Website      string `gorm:"type:varchar(255)" json:"website"`
	Notes        string `json:"notes,omitempty"`
}

// 🏭 Manufacturer represents the supplier or producer of utility meters
// 🔁 Used to associate each meter with a manufacturer for traceability

/*
📦 Example Output (JSON):
{
  "id": 3,
  "name": "HydroTech Industries",
  "country": "Germany",
  "contact_name": "Stefan Keller",
  "contact_email": "s.keller@hydrotech.de",
  "phone": "+49 30 12345678",
  "website": "https://hydrotech.de",
  "notes": "Approved vendor for smart water meters"
}
*/
