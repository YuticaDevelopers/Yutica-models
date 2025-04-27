// models/accounts/user.go
package accounts

import (
	"time"

	"github.com/YuticaDevelopers/Yutica-models/models/common"
)

type User struct {
	common.BaseModel

	FirstName   string              `gorm:"type:varchar(50)" json:"first_name"`
	LastName    string              `gorm:"type:varchar(50)" json:"last_name"`
	Email       string              `gorm:"type:varchar(100);uniqueIndex" json:"email"`
	Phone       string              `gorm:"type:varchar(20);uniqueIndex" json:"phone"`
	Username    string              `gorm:"type:varchar(30);uniqueIndex" json:"username"`
	Password    string              `gorm:"type:varchar(255)" json:"password"` // hashed password
	Role        common.UserRole     `gorm:"type:varchar(20);index" json:"role"`
	Permissions []common.Permission `gorm:"-" json:"permissions,omitempty"` // dynamically loaded based on role
	IsActive    bool                `json:"is_active"`

	AccountIDs  []uint `gorm:"-" json:"account_ids,omitempty"`  // account_user may belong to multiple accounts
	PropertyIDs []uint `gorm:"-" json:"property_ids,omitempty"` // property_user may have access to one or more properties

	CurrentOTP        *string    `gorm:"type:varchar(10)" json:"current_otp,omitempty"`
	OTPExpiration     *time.Time `json:"otp_expiration,omitempty"`
	LastLogin         *time.Time `json:"last_login,omitempty"`
	LastPasswordReset *time.Time `json:"last_password_reset,omitempty"`
	Notes             string     `json:"notes,omitempty"`
}

// 🔐 Represents system-level user (admin, technician, account rep, etc.)
// 🔁 Indexed by email, phone, username
// 🛡️ Sensitive data like passwords should be hashed and not exposed in APIs
// 🧠 Permissions are assigned dynamically based on role from common.RolePermissions
// 📦 AccountUser may be assigned to multiple account IDs for billing access

/*
📦 Example Output (JSON):
{
  "id": 101,
  "first_name": "Daniel",
  "last_name": "Mburu",
  "email": "daniel.mburu@yutica.co.ke",
  "phone": "+254701234567",
  "username": "dmburu",
  "role": "admin",
  "permissions": [
    "account:create",
    "account:view",
    "account:edit",
    "tenant:create",
    "tenant:view",
    "meter:read",
    "meter:update",
    "billing:view",
    "billing:invoice",
    "support:manage",
    "audit:read"
  ],
  "is_active": true,
  "account_ids": [1, 2, 3],
  "current_otp": "928401",
  "otp_expiration": "2025-04-24T09:05:00Z",
  "last_login": "2025-04-24T08:00:00Z",
  "last_password_reset": "2025-04-01T00:00:00Z",
  "notes": "Main system administrator"
}
*/
