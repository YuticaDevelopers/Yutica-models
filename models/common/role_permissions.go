package common

var RolePermissions = map[string][]Permission{
	"admin": {
		PermCreateAccount,
		PermViewAccount,
		PermEditAccount,
		PermCreateTenant,
		PermViewTenant,
		PermReadMeter,
		PermUpdateMeter,
		PermViewBilling,
		PermGenerateInvoice,
		PermManageTickets,
		PermAuditAccess,
	},
	"account_user": {
		PermViewAccount,
		PermViewTenant,
		PermReadMeter,
		PermViewBilling,
	},
	"property_user": {
		PermViewTenant,
		PermReadMeter,
	},
	"technician": {
		PermReadMeter,
		PermUpdateMeter,
	},
}
