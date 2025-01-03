package entity

import "fmt"

// TransactionItemCategory represents a category for transaction items.
type TransactionItemCategory struct {
	Common
	LedgerID int64  // ID of the ledger this category belongs to
	Name     string // Name of the category
}

// String returns a string representation of the TransactionItemCategory.
func (c TransactionItemCategory) String() string {
	return fmt.Sprintf("TransactionItemCategory[ID=%d, LedgerID=%d, Name=%s, CreatedAt=%s, UpdatedAt=%s, DeletedAt=%s]",
		c.ID, c.LedgerID, c.Name, c.CreatedAt, c.UpdatedAt, c.DeletedAt)
}

// TransactionItem represents an item in a transaction.
type TransactionItem struct {
	Common
	LedgerID   int64  // ID of the ledger this item belongs to
	CategoryID int64  // ID of the category this item belongs to
	Name       string // Name of the item
	ItemType   string // Type of the item
}

// String returns a string representation of the TransactionItem.
func (i TransactionItem) String() string {
	return fmt.Sprintf("TransactionItem[ID=%d, LedgerID=%d, CategoryID=%d, Name=%s, ItemType=%s, CreatedAt=%s, UpdatedAt=%s, DeletedAt=%s]",
		i.ID, i.LedgerID, i.CategoryID, i.Name, i.ItemType, i.CreatedAt, i.UpdatedAt, i.DeletedAt)
}

// ItemType represents the type of a transaction item.
type ItemType string

// String returns the string representation of the ItemType.
func (i ItemType) String() string {
	return string(i)
}

// IsValid checks if the ItemType is valid.
func (i ItemType) IsValid() bool {
	switch i {
	case ItemTypeExpense:
		return true
	}
	return false
}

const (
	// ItemTypeExpense represents an expense item type.
	ItemTypeExpense ItemType = "EXPENSE"
)
