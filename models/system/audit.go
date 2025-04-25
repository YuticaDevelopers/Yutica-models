// models/system/audit.go
package system

import (
	"time"

	"github.com/YuticaDevelopers/Yutica-models/models/common"
)

type AuditLog struct {
	common.BaseModel

	ActorID   uint      `gorm:"index" json:"actor_id"`              // ID of the user who performed the action
	Action    string    `gorm:"type:varchar(50)" json:"action"`     // e.g., create, update, delete
	Entity    string    `gorm:"type:varchar(100)" json:"entity"`    // e.g., User, Invoice
	EntityID  uint      `gorm:"index" json:"entity_id"`             // ID of the entity affected
	Timestamp time.Time `json:"timestamp"`                          // Time of the action
	Details   string    `gorm:"type:text" json:"details"`           // Additional details about the action
	IPAddress string    `gorm:"type:varchar(45)" json:"ip_address"` // IP address of the client
	Hostname  string    `gorm:"type:varchar(255)" json:"hostname"`  // Hostname of the client
}

// {
// 	"id": 1,
// 	"actor_id": 42,
// 	"action": "update",
// 	"entity": "User",
// 	"entity_id": 101,
// 	"timestamp": "2025-04-25T10:15:00Z",
// 	"details": "Updated user email from old@example.com to new@example.com",
// 	"ip_address": "192.168.1.100",
// 	"hostname": "user-laptop.local"
//   }
