package support

import "github.com/YuticaDevelopers/Yutica-models/models/common"

type FAQ struct {
	common.BaseModel

	Question string `gorm:"type:text" json:"question"`
	Answer   string `gorm:"type:text" json:"answer"`
	Category string `gorm:"type:varchar(50);index" json:"category"`
	IsActive bool   `json:"is_active"`
}

// {
// 	"id": 101,
// 	"question": "How do I reset my password?",
// 	"answer": "Click on 'Forgot Password' at the login screen and follow the instructions sent to your email.",
// 	"category": "Account Management",
// 	"is_active": true
//   }
