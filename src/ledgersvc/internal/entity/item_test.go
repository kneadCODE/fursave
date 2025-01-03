package entity

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestTransactionItemCategory_String(t *testing.T) {
	category := TransactionItemCategory{
		Common: Common{
			ID:        1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		LedgerID: 123,
		Name:     "Food",
	}

	expected := "TransactionItemCategory[ID=1, LedgerID=123, Name=Food, CreatedAt=" + category.CreatedAt.String() + ", UpdatedAt=" + category.UpdatedAt.String() + ", DeletedAt=" + category.DeletedAt.String() + "]"
	require.Equal(t, expected, category.String())
}

func TestTransactionItem_String(t *testing.T) {
	item := TransactionItem{
		Common: Common{
			ID:        1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		LedgerID:   123,
		CategoryID: 456,
		Name:       "Burger",
		ItemType:   "Food",
	}

	expected := "TransactionItem[ID=1, LedgerID=123, CategoryID=456, Name=Burger, ItemType=Food, CreatedAt=" + item.CreatedAt.String() + ", UpdatedAt=" + item.UpdatedAt.String() + ", DeletedAt=" + item.DeletedAt.String() + "]"
	require.Equal(t, expected, item.String())
}

func TestItemType_String(t *testing.T) {
	itemType := ItemTypeExpense
	require.Equal(t, "EXPENSE", itemType.String())
}

func TestItemType_IsValid(t *testing.T) {
	tests := []struct {
		name     string
		itemType ItemType
		expected bool
	}{
		{"Valid Expense", ItemTypeExpense, true},
		{"Invalid Type", ItemType("INVALID"), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.expected, tt.itemType.IsValid())
		})
	}
}
