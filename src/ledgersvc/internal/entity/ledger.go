package entity

import (
	"fmt"
)

// Ledger represents a financial ledger with common fields and a name.
type Ledger struct {
	Common
	Name string // Name of the ledger
}

// String returns a string representation of the Ledger.
func (l Ledger) String() string {
	return fmt.Sprintf("Ledger[ID=%d, Name=%s, CreatedAt=%s, UpdatedAt=%s, DeletedAt=%s]",
		l.ID, l.Name, l.CreatedAt, l.UpdatedAt, l.DeletedAt)
}
