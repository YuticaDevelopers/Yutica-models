// models/sessions/session.go
package sessions

import (
	"time"

	"github.com/YuticaDevelopers/Yutica-models/models/common"
)

type Session struct {
	common.BaseModel

	UserID    uint      `gorm:"index" json:"user_id"`               // 🔍 Fast lookup by user
	Token     string    `gorm:"uniqueIndex;not null" json:"token"`  // 🔒 Ensure token uniqueness
	IPAddress string    `gorm:"type:varchar(45)" json:"ip_address"` // 🌐 IPv6 ready
	UserAgent string    `gorm:"type:text" json:"user_agent"`        // 🖥️ Client device info
	ExpiresAt time.Time `gorm:"index" json:"expires_at"`            // 🔍 Quick expiry lookups
	Revoked   bool      `gorm:"default:false;index" json:"revoked"` // 🛡️ Fast check if revoked
}
