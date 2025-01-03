package entity

import "fmt"

// Payee represents a payee in the system.
type Payee struct {
	Common
	LedgerID  int64  // ID of the ledger this payee belongs to
	Name      string // Name of the payee
	PayeeType string // Type of the payee
}

// String returns a string representation of the Payee.
func (p Payee) String() string {
	return fmt.Sprintf("Payee[ID=%d, LedgerID=%d, Name=%s, PayeeType=%s, CreatedAt=%s, UpdatedAt=%s, DeletedAt=%s]",
		p.ID, p.LedgerID, p.Name, p.PayeeType, p.CreatedAt, p.UpdatedAt, p.DeletedAt)
}
