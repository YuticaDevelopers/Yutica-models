// models/accounts/account.go
package accounts

import (
	"time"

	"github.com/YuticaDevelopers/Yutica-models/models/common"
)

type Account struct {
	common.BaseModel

	Name        string     `gorm:"type:varchar(100);uniqueIndex" json:"name"`
	Email       string     `gorm:"type:varchar(100);uniqueIndex" json:"email"`
	Phone       string     `gorm:"type:varchar(20)" json:"phone"`
	Role        string     `gorm:"type:varchar(50);index" json:"role"` // e.g., admin, tenant, technician
	IsActive    bool       `json:"is_active"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	LastLoginAt *time.Time `json:"last_login_at,omitempty"`
}

// {
// 	"id": 404,
// 	"name": "Jane Doe",
// 	"email": "jane.doe@example.com",
// 	"phone": "+254712345678",
// 	"role": "tenant",
// 	"is_active": true,
// 	"created_at": "2025-03-15T10:30:00Z",
// 	"updated_at": "2025-04-20T14:45:00Z",
// 	"last_login_at": "2025-04-24T08:15:00Z"
//   }
