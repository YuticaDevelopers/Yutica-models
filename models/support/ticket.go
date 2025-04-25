package support

import (
	"time"

	"github.com/YuticaDevelopers/Yutica-models/models/common"
)

type Ticket struct {
	common.BaseModel

	Subject      string              `gorm:"type:varchar(200)" json:"subject"`
	Description  string              `gorm:"type:text" json:"description"`
	Status       common.TicketStatus `gorm:"type:varchar(20);index" json:"status"`
	Priority     string              `gorm:"type:varchar(20)" json:"priority"` // e.g., low, medium, high
	CreatedByID  uint                `gorm:"index" json:"created_by_id"`
	AssignedToID *uint               `gorm:"index" json:"assigned_to_id,omitempty"`
	ResolvedAt   *time.Time          `json:"resolved_at,omitempty"`
	ClosedAt     *time.Time          `json:"closed_at,omitempty"`
}

// {
// 	"id": 303,
// 	"subject": "No Water Supply",
// 	"description": "There has been no water supply since yesterday evening.",
// 	"status": "open",
// 	"priority": "high",
// 	"created_by_id": 22,
// 	"assigned_to_id": 5,
// 	"resolved_at": null,
// 	"closed_at": null
//   }
