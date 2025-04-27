// models/sessions/archived_session.go
package sessions

import (
	"time"

	"github.com/YuticaDevelopers/Yutica-models/models/common"
)

type ArchivedSession struct {
	common.BaseModel

	OriginalSessionID uint      `gorm:"index" json:"original_session_id"`   // 🔍 Faster lookups for original session reference
	UserID            uint      `gorm:"index" json:"user_id"`               // 🔍 Fast querying by user
	Token             string    `gorm:"uniqueIndex" json:"token"`           // 🔒 Token must be unique even if archived
	IPAddress         string    `gorm:"type:varchar(45)" json:"ip_address"` // IPv6 supported
	UserAgent         string    `gorm:"type:text" json:"user_agent"`        // User device info
	ExpiresAt         time.Time `gorm:"index" json:"expires_at"`            // 🔍 Filter based on expiry time
	Revoked           bool      `gorm:"index" json:"revoked"`               // 🔍 Index revoked status
	ArchivedAt        time.Time `gorm:"autoCreateTime" json:"archived_at"`  // Date of archival
}
