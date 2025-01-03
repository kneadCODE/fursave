package entity

import (
	"fmt"
)

// Access represents the access control for a ledger.
type Access struct {
	Common
	LedgerID int64      // ID of the ledger
	UserID   string     // ID of the user who has access
	Type     AccessType // Type of access the user has
}

// String returns a string representation of the Access.
func (a Access) String() string {
	return fmt.Sprintf("Access[LedgerID=%d, UserID=%s, Type=%s, CreatedAt=%s, UpdatedAt=%s, DeletedAt=%s]",
		a.LedgerID, a.UserID, a.Type, a.CreatedAt, a.UpdatedAt, a.DeletedAt)
}

// AccessType represents the type of access a user can have.
type AccessType string

// String returns the string representation of the AccessType.
func (a AccessType) String() string {
	return string(a)
}

// IsValid checks if the AccessType is valid.
func (a AccessType) IsValid() bool {
	switch a {
	case AccessTypeOwner, AccessTypeReader, AccessTypeContributor:
		return true
	}

	return false
}

const (
	// AccessTypeOwner represents owner access.
	AccessTypeOwner AccessType = "OWNER"
	// AccessTypeReader represents read-only access.
	AccessTypeReader AccessType = "READER"
	// AccessTypeContributor represents contributor access.
	AccessTypeContributor AccessType = "CONTRIBUTOR"
)
