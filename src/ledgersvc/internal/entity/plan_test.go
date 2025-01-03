package entity

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestBudgetPlan_String(t *testing.T) {
	plan := BudgetPlan{
		Common: Common{
			ID:        1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		LedgerID: 123,
		Name:     "Monthly Budget",
	}

	expected := "BudgetPlan[ID=1, LedgerID=123, Name=Monthly Budget, CreatedAt=" + plan.CreatedAt.String() + ", UpdatedAt=" + plan.UpdatedAt.String() + ", DeletedAt=" + plan.DeletedAt.String() + "]"
	require.Equal(t, expected, plan.String())
}

func TestBudgetPlanItem_String(t *testing.T) {
	item := BudgetPlanItem{
		Common: Common{
			ID:        1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		LedgerID: 123,
		PlanID:   456,
		ItemID:   789,
		Amount:   1000,
	}

	expected := "BudgetPlanItem[ID=1, LedgerID=123, PlanID=456, ItemID=789, Amount=1000, CreatedAt=" + item.CreatedAt.String() + ", UpdatedAt=" + item.UpdatedAt.String() + ", DeletedAt=" + item.DeletedAt.String() + "]"
	require.Equal(t, expected, item.String())
}
