package meters

import (
	"time"

	"github.com/YuticaDevelopers/Yutica-models/models/common"
)

type MeterLog struct {
	common.BaseModel

	MeterID   uint                  `gorm:"index" json:"meter_id"`                    // Associated meter ID
	EventType common.MeterEventType `gorm:"type:varchar(50);index" json:"event_type"` // Type of event
	Value     *float64              `json:"value,omitempty"`                          // Optional value (e.g., reading)
	Unit      *string               `gorm:"type:varchar(10)" json:"unit,omitempty"`   // Unit of measurement
	Timestamp time.Time             `json:"timestamp"`                                // Time of the event
	Source    string                `gorm:"type:varchar(50)" json:"source"`           // Source of the event (e.g., system, user)
	IPAddress string                `gorm:"type:varchar(45)" json:"ip_address"`       // IP address of the source
	Hostname  string                `gorm:"type:varchar(255)" json:"hostname"`        // Hostname of the source
	Notes     string                `gorm:"type:text" json:"notes,omitempty"`         // Additional notes
}

// {
// 	"id": 1,
// 	"meter_id": 101,
// 	"event_type": "power_on",
// 	"timestamp": "2025-04-25T08:00:00Z",
// 	"source": "system",
// 	"ip_address": "192.168.1.10",
// 	"hostname": "meter-controller-01",
// 	"notes": "Scheduled activation"
//   }

// {
// 	"id": 2,
// 	"meter_id": 202,
// 	"event_type": "valve_close",
// 	"timestamp": "2025-04-25T09:30:00Z",
// 	"source": "user",
// 	"ip_address": "192.168.1.15",
// 	"hostname": "operator-terminal-02",
// 	"notes": "Emergency shutdown due to leak"
//   }

// {
// 	"id": 3,
// 	"meter_id": 303,
// 	"event_type": "reading",
// 	"value": 1234.56,
// 	"unit": "kWh",
// 	"timestamp": "2025-04-25T10:00:00Z",
// 	"source": "system",
// 	"ip_address": "192.168.1.20",
// 	"hostname": "auto-reader-03",
// 	"notes": "Monthly automated reading"
//   }
