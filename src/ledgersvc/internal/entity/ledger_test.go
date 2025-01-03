package entity

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestLedgerString(t *testing.T) {
	ledger := Ledger{
		Common: Common{
			ID:        1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Name: "Test Ledger",
	}

	expected := "Ledger[ID=1, Name=Test Ledger, CreatedAt=" + ledger.CreatedAt.String() + ", UpdatedAt=" + ledger.UpdatedAt.String() + ", DeletedAt=" + ledger.DeletedAt.String() + "]"
	require.Equal(t, expected, ledger.String())
}
