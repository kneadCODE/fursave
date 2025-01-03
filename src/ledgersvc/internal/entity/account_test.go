package entity

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestAccountString(t *testing.T) {
	account := Account{
		Common: Common{
			ID:        1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Name: "Test Account",
		Type: AccountTypeChecking,
	}

	expected := "Account[ID=1, LedgerID=0, Name=Test Account, Type=CHECKING, CreatedAt=" + account.CreatedAt.String() + ", UpdatedAt=" + account.UpdatedAt.String() + ", DeletedAt=" + account.DeletedAt.String() + "]"
	require.Equal(t, expected, account.String())
}

func TestAccountTypeString(t *testing.T) {
	tests := []struct {
		name string
		a    AccountType
		want string
	}{
		{"Checking", AccountTypeChecking, "CHECKING"},
		{"Savings", AccountTypeSavings, "SAVINGS"},
		{"Cash", AccountTypeCash, "CASH"},
		{"Digital Wallet", AccountTypeDigitalWallet, "DIGITAL_WALLET"},
		{"Fixed Deposit", AccountTypeFixedDeposit, "FIXED_DEPOSIT"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, tt.a.String())
		})
	}
}

func TestAccountTypeIsValid(t *testing.T) {
	tests := []struct {
		name        string
		accountType AccountType
		expected    bool
	}{
		{"Valid Checking", AccountTypeChecking, true},
		{"Valid Savings", AccountTypeSavings, true},
		{"Valid Cash", AccountTypeCash, true},
		{"Valid Digital Wallet", AccountTypeDigitalWallet, true},
		{"Valid Fixed Deposit", AccountTypeFixedDeposit, true},
		{"Invalid Type", AccountType("INVALID"), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.expected, tt.accountType.IsValid())
		})
	}
}
