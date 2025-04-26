// models/accounts/account.go
package accounts

import (
	"github.com/YuticaDevelopers/Yutica-models/models/common"
)

type Account struct {
	common.BaseModel

	AccountName string `json:"account_name"`
	IsActive    bool   `json:"is_active"`
}

// {
// 	"id": 404,
// 	"name": "Jane Doe",
// 	"email": "jane.doe@example.com",
// 	"phone": "+254712345678",
// 	"role": "tenant",
// 	"is_active": true,
// 	"created_at": "2025-03-15T10:30:00Z",
// 	"updated_at": "2025-04-20T14:45:00Z",
// 	"last_login_at": "2025-04-24T08:15:00Z"
//   }
