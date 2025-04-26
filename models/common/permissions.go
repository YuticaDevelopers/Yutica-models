package common

type Permission string

const (
	// Account management
	PermCreateAccount Permission = "account:create"
	PermViewAccount   Permission = "account:view"
	PermEditAccount   Permission = "account:edit"

	// Tenant management
	PermCreateTenant Permission = "tenant:create"
	PermViewTenant   Permission = "tenant:view"

	// Meter management
	PermReadMeter   Permission = "meter:read"
	PermUpdateMeter Permission = "meter:update"

	// Billing
	PermViewBilling     Permission = "billing:view"
	PermGenerateInvoice Permission = "billing:invoice"

	// Support
	PermManageTickets Permission = "support:manage"

	// System
	PermAuditAccess Permission = "audit:read"
)

// AllPermissions returns a list of all available permissions
func AllPermissions() []Permission {
	return []Permission{
		// Account permissions
		PermCreateAccount,
		PermViewAccount,
		PermEditAccount,

		// Tenant permissions
		PermCreateTenant,
		PermViewTenant,

		// Meter permissions
		PermReadMeter,
		PermUpdateMeter,

		// Billing permissions
		PermViewBilling,
		PermGenerateInvoice,

		// Support permissions
		PermManageTickets,

		// System permissions
		PermAuditAccess,
	}
}
