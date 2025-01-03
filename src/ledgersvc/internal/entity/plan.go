package entity

import "fmt"

// BudgetPlan represents a budget plan.
type BudgetPlan struct {
	Common
	LedgerID int64  // ID of the ledger this budget plan belongs to
	Name     string // Name of the budget plan
}

func (b BudgetPlan) String() string {
	return fmt.Sprintf("BudgetPlan[ID=%d, LedgerID=%d, Name=%s, CreatedAt=%s, UpdatedAt=%s, DeletedAt=%s]",
		b.ID, b.LedgerID, b.Name, b.CreatedAt, b.UpdatedAt, b.DeletedAt)
}

// BudgetPlanItem represents an item in a budget plan.
type BudgetPlanItem struct {
	Common
	LedgerID int64 // ID of the ledger this budget plan item belongs to
	PlanID   int64 // ID of the budget plan this item belongs to
	ItemID   int64 // ID of the transaction item this budget plan item refers to
	Amount   int   // Amount allocated for this budget plan item
}

func (b BudgetPlanItem) String() string {
	return fmt.Sprintf("BudgetPlanItem[ID=%d, LedgerID=%d, PlanID=%d, ItemID=%d, Amount=%d, CreatedAt=%s, UpdatedAt=%s, DeletedAt=%s]",
		b.ID, b.LedgerID, b.PlanID, b.ItemID, b.Amount, b.CreatedAt, b.UpdatedAt, b.DeletedAt)
}
