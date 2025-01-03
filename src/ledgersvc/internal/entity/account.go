package entity

import "fmt"

// Account represents a financial account.
type Account struct {
	Common
	LedgerID int64       // ID of the ledger this account belongs to
	Name     string      // Name of the account
	Type     AccountType // Type of the account
}

// String returns a string representation of the Account.
func (a Account) String() string {
	return fmt.Sprintf("Account[ID=%d, LedgerID=%d, Name=%s, Type=%s, CreatedAt=%s, UpdatedAt=%s, DeletedAt=%s]",
		a.ID, a.LedgerID, a.Name, a.Type, a.CreatedAt, a.UpdatedAt, a.DeletedAt)
}

// AccountType represents the type of an account.
type AccountType string

// String returns the string representation of the AccountType.
func (a AccountType) String() string {
	return string(a)
}

// IsValid checks if the AccountType is valid.
func (a AccountType) IsValid() bool {
	switch a {
	case AccountTypeChecking,
		AccountTypeSavings,
		AccountTypeCash,
		AccountTypeDigitalWallet,
		AccountTypeFixedDeposit:
		return true
	}

	return false
}

const (
	// AccountTypeChecking represents a checking account.
	AccountTypeChecking AccountType = "CHECKING"
	// AccountTypeSavings represents a savings account.
	AccountTypeSavings AccountType = "SAVINGS"
	// AccountTypeCash represents a cash account.
	AccountTypeCash AccountType = "CASH"
	// AccountTypeDigitalWallet represents a digital wallet account.
	AccountTypeDigitalWallet AccountType = "DIGITAL_WALLET"
	// AccountTypeFixedDeposit represents a fixed deposit account.
	AccountTypeFixedDeposit AccountType = "FIXED_DEPOSIT"
)
