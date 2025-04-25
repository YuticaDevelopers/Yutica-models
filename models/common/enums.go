package common

type UtilityType string

const (
	WaterMeter UtilityType = "water"
	PowerMeter UtilityType = "power"
)

type UserRole string

const (
	AdminUser    UserRole = "admin"
	AccountUser  UserRole = "account_user"
	PropertyUser UserRole = "property_user"
	Technician   UserRole = "technician"
)

type TransactionType string

const (
	TransactionTypeMpesa  TransactionType = "mpesa"
	TransactionTypeWallet TransactionType = "wallet"
	TransactionTypeToken  TransactionType = "token"
)

type TransactionStatus string

const (
	TransactionStatusPending   TransactionStatus = "pending"
	TransactionStatusConfirmed TransactionStatus = "confirmed"
	TransactionStatusFailed    TransactionStatus = "failed"
)

type InvoiceStatus string

const (
	InvoicePending InvoiceStatus = "pending"
	InvoicePaid    InvoiceStatus = "paid"
	InvoiceOverdue InvoiceStatus = "overdue"
)

type BillingAccountType string

type BillingMode string

type BalanceUnit string

type AccountValueMode string

const (
	BillingTypeWater BillingAccountType = "water"
	BillingTypePower BillingAccountType = "power"

	BillingModePrepaid  BillingMode = "prepaid"
	BillingModePostpaid BillingMode = "postpaid"

	BalanceUnitKES    BalanceUnit = "KES"
	BalanceUnitKWH    BalanceUnit = "kWh"
	BalanceUnitLitres BalanceUnit = "litres"

	ValueModeWallet AccountValueMode = "wallet"
	ValueModeToken  AccountValueMode = "token"
)

type CreditSource string

const (
	CreditSourceMpesa    CreditSource = "mpesa"
	CreditSourceWallet   CreditSource = "wallet"
	CreditSourceReversal CreditSource = "reversal"
)

type PrepaidDebitReason string

const (
	DebitReasonUsage      PrepaidDebitReason = "usage"      // based on consumption
	DebitReasonInvoice    PrepaidDebitReason = "invoice"    // postpaid invoice deduction
	DebitReasonCorrection PrepaidDebitReason = "correction" // manual admin adjustment
	DebitReasonReversal   PrepaidDebitReason = "reversal"   // reversal of a previous incorrect debit
)

type NegativeStatus string

const (
	NegativeStatusUnresolved NegativeStatus = "unresolved"
	NegativeStatusResolved   NegativeStatus = "resolved"
	NegativeStatusFlagged    NegativeStatus = "flagged"
)

type RateType string

const (
	RateTypeFixed   RateType = "Fixed"
	RateTypeFormula RateType = "Formula"
)

type TicketStatus string

const (
	TicketOpen     TicketStatus = "open"
	TicketPending  TicketStatus = "pending"
	TicketResolved TicketStatus = "resolved"
	TicketClosed   TicketStatus = "closed"
)

type MeterEventType string

const (
	EventPowerOn    MeterEventType = "power_on"
	EventPowerOff   MeterEventType = "power_off"
	EventValveOpen  MeterEventType = "valve_open"
	EventValveClose MeterEventType = "valve_close"
	EventReading    MeterEventType = "reading"
	EventTamper     MeterEventType = "tamper"
	EventLeak       MeterEventType = "leak"
	EventLowBattery MeterEventType = "low_battery"
)
