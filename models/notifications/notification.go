package notifications

import (
	"time"

	"github.com/YuticaDevelopers/Yutica-models/models/common"
)

type Notification struct {
	common.BaseModel

	To        string     `gorm:"type:varchar(100);not null" json:"to"`
	Subject   string     `gorm:"type:varchar(255)" json:"subject"`
	Message   string     `gorm:"type:text" json:"message"`
	Type      string     `gorm:"type:varchar(20);not null" json:"type"`             // email or sms
	CC        *string    `gorm:"type:varchar(255)" json:"cc,omitempty"`             // optional
	Status    string     `gorm:"type:varchar(20);default:'Not Sent'" json:"status"` // Not Sent, Sent, Failed
	SentAt    *time.Time `json:"sent_at,omitempty"`
	ChannelID *string    `gorm:"type:varchar(100)" json:"channel_id,omitempty"` // e.g., SMS vendor ID or email message ID
}
