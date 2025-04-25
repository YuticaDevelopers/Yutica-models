package billing

import (
	"time"

	"github.com/YuticaDevelopers/Yutica-models/models/common"
)

type Rate struct {
	common.BaseModel

	Name        string     `gorm:"type:varchar(100);index" json:"name"`
	AccountType string     `gorm:"type:varchar(20);index" json:"account_type"` // e.g., water, power
	Unit        string     `gorm:"type:varchar(10)" json:"unit"`               // e.g., kWh, litres, KES
	Value       float64    `json:"value"`                                      // price per unit
	Currency    string     `gorm:"type:varchar(10)" json:"currency"`           // e.g., KES
	EffectiveAt time.Time  `json:"effective_at"`
	ExpiresAt   *time.Time `json:"expires_at,omitempty"`
	IsActive    bool       `json:"is_active"`
	Notes       string     `json:"notes,omitempty"`
}

// {
// 	"id": 1,
// 	"name": "Standard Water Rate",
// 	"account_type": "water",
// 	"unit": "litres",
// 	"value": 0.05,
// 	"currency": "KES",
// 	"effective_at": "2025-01-01T00:00:00Z",
// 	"expires_at": null,
// 	"notes": "Applicable to residential properties"
//   }
