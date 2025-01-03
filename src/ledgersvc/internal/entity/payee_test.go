package entity

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestPayee_String(t *testing.T) {
	payee := Payee{
		Common: Common{
			ID:        1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		LedgerID:  123,
		Name:      "John Doe",
		PayeeType: "Individual",
	}

	expected := "Payee[ID=1, LedgerID=123, Name=John Doe, PayeeType=Individual, CreatedAt=" + payee.CreatedAt.String() + ", UpdatedAt=" + payee.UpdatedAt.String() + ", DeletedAt=" + payee.DeletedAt.String() + "]"
	require.Equal(t, expected, payee.String())
}
