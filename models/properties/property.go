// models/properties/property.go
package properties

import (
	"time"

	"github.com/YuticaDevelopers/Yutica-models/models/common"
)

type Property struct {
	common.BaseModel

	Name          string    `gorm:"type:varchar(100);uniqueIndex" json:"name"`
	status        string    `gorm:"type:varchar(50);uniqueIndex" json:"status"` // internal status for property reference
	Type          string    `gorm:"type:varchar(50)" json:"type"`               // e.g., Residential, Commercial
	Location      string    `gorm:"type:varchar(255)" json:"location"`
	County        string    `gorm:"type:varchar(50)" json:"county"`
	Region        string    `gorm:"type:varchar(50)" json:"region"`
	ContactPhone  string    `gorm:"type:varchar(20)" json:"contact_phone"`
	ContactEmail  string    `gorm:"type:varchar(100)" json:"contact_email"`
	AccountID     uint      `gorm:"index" json:"account_id"` // Associated owning account
	ManagerUserID *uint     `gorm:"index" json:"manager_user_id,omitempty"`
	IsActive      bool      `json:"is_active"`
	RegisteredAt  time.Time `json:"registered_at"`
	Notes         string    `json:"notes,omitempty"`
}

// 🏠 Property is the top-level entity grouping units, tenants, meters, and billing accounts
// 🔁 Indexed by name, status, region, and account ID for fast access

/*
📦 Example Output (JSON):
{
  "id": 12,
  "name": "Greenhill Estate",
  "status": "GH-EST-001",
  "type": "Residential",
  "location": "Kiambu Road, Nairobi",
  "county": "Nairobi",
  "region": "Central",
  "contact_phone": "+254712345678",
  "contact_email": "admin@greenhill.co.ke",
  "account_id": 8,
  "manager_user_id": 101,
  "is_active": true,
  "registered_at": "2024-01-15T00:00:00Z",
  "notes": "Managed by Greenhill Group"
}
*/
