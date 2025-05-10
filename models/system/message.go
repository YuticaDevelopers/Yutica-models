// models/system/notification.go
package system

import (
	"time"

	"github.com/YuticaDevelopers/Yutica-models/models/common"
)

type Message struct {
	common.BaseModel

	UserID    uint       `gorm:"index" json:"user_id"`           // ID of the user to notify
	Title     string     `gorm:"type:varchar(100)" json:"title"` // Notification title
	Message   string     `gorm:"type:text" json:"message"`       // Notification message
	IsRead    bool       `json:"is_read"`                        // Read status
	CreatedAt time.Time  `json:"created_at"`                     // Time of creation
	ReadAt    *time.Time `json:"read_at,omitempty"`              // Time when read
}

// {
// 	"id": 1,
// 	"user_id": 42,
// 	"title": "Payment Due",
// 	"message": "Your payment is due on 2025-05-01.",
// 	"is_read": false,
// 	"created_at": "2025-04-25T09:00:00Z",
// 	"read_at": null
//   }
