package entity

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestAccess_String(t *testing.T) {
	access := Access{
		Common: Common{
			ID:        1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		LedgerID: 123,
		UserID:   "user1",
		Type:     AccessTypeOwner,
	}

	expected := "Access[LedgerID=123, UserID=user1, Type=OWNER, CreatedAt=" + access.CreatedAt.String() + ", UpdatedAt=" + access.UpdatedAt.String() + ", DeletedAt=" + access.DeletedAt.String() + "]"
	require.Equal(t, expected, access.String())
}

func TestAccessType_String(t *testing.T) {
	tests := []struct {
		name string
		a    AccessType
		want string
	}{
		{"Owner", AccessTypeOwner, "OWNER"},
		{"Reader", AccessTypeReader, "READER"},
		{"Contributor", AccessTypeContributor, "CONTRIBUTOR"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, tt.a.String(), "AccessType.String() = %v, want %v", tt.a.String(), tt.want)
		})
	}
}

func TestAccessType_IsValid(t *testing.T) {
	tests := []struct {
		name string
		a    AccessType
		want bool
	}{
		{"ValidOwner", AccessTypeOwner, true},
		{"ValidReader", AccessTypeReader, true},
		{"ValidContributor", AccessTypeContributor, true},
		{"InvalidType", AccessType("INVALID"), false},
		{"EmptyType", AccessType(""), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, tt.a.IsValid(), "AccessType.IsValid() = %v, want %v", tt.a.IsValid(), tt.want)
		})
	}
}
