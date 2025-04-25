package sessions

import (
	"time"

	"github.com/YuticaDevelopers/Yutica-models/models/common"
)

type Session struct {
	common.BaseModel

	UserID    uint      `gorm:"index" json:"user_id"`
	Token     string    `gorm:"uniqueIndex;not null" json:"token"`
	IPAddress string    `json:"ip_address"`
	UserAgent string    `json:"user_agent"`
	ExpiresAt time.Time `gorm:"index" json:"expires_at"`
	Revoked   bool      `gorm:"default:false" json:"revoked"`
}
