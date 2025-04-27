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
