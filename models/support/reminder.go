package support

import (
	"time"

	"github.com/YuticaDevelopers/Yutica-models/models/common"
)

type Reminder struct {
	common.BaseModel

	UserID    uint       `gorm:"index" json:"user_id"`
	Title     string     `gorm:"type:varchar(100)" json:"title"`
	Message   string     `gorm:"type:text" json:"message"`
	TriggerAt time.Time  `json:"trigger_at"`
	Recurring bool       `json:"recurring"`
	Interval  string     `gorm:"type:varchar(20)" json:"interval,omitempty"` // e.g., daily, weekly
	IsSent    bool       `json:"is_sent"`
	SentAt    *time.Time `json:"sent_at,omitempty"`
}

// {
// 	"id": 202,
// 	"user_id": 15,
// 	"title": "Meter Reading Submission",
// 	"message": "Please submit your monthly meter reading.",
// 	"trigger_at": "2025-04-30T09:00:00Z",
// 	"recurring": true,
// 	"interval": "monthly",
// 	"is_sent": false,
// 	"sent_at": null
//   }
