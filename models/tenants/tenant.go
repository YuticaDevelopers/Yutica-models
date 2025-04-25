// models/tenants/tenant.go
package tenants

import (
	"time"

	"github.com/YuticaDevelopers/Yutica-models/models/common"
)

type Tenant struct {
	common.BaseModel

	FirstName   string     `gorm:"type:varchar(50)" json:"first_name"`
	LastName    string     `gorm:"type:varchar(50)" json:"last_name"`
	Email       string     `gorm:"type:varchar(100);uniqueIndex" json:"email"`
	Phone       string     `gorm:"type:varchar(20);uniqueIndex" json:"phone"`
	NationalID  string     `gorm:"type:varchar(20);uniqueIndex" json:"national_id"`
	Gender      string     `gorm:"type:varchar(10)" json:"gender"`
	DateOfBirth *time.Time `json:"date_of_birth,omitempty"`

	AccountID  uint       `gorm:"index" json:"account_id"`
	PropertyID uint       `gorm:"index" json:"property_id"`
	UnitID     uint       `gorm:"index" json:"unit_id"`
	JoinedAt   time.Time  `json:"joined_at"`
	LeftAt     *time.Time `json:"left_at,omitempty"`
	IsActive   bool       `json:"is_active"`
	Notes      string     `json:"notes,omitempty"`
}

// ⚠️ Primary customer record for unit-based utility billing
// 🔁 Indexed by email, phone, national_id, account, property, unit
// 🔐 Used to join with billing accounts, meters, usage, and invoices

/*
📦 Example Output (JSON):
{
  "id": 12,
  "first_name": "Jane",
  "last_name": "Doe",
  "email": "jane.doe@example.com",
  "phone": "+254712345678",
  "national_id": "12345678",
  "gender": "female",
  "date_of_birth": "1992-07-15",
  "account_id": 3,
  "property_id": 7,
  "unit_id": 22,
  "joined_at": "2024-11-01T00:00:00Z",
  "left_at": null,
  "is_active": true,
  "notes": "Moved in November"
}
*/
