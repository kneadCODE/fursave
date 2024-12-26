package cfg

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestContextKey_String(t *testing.T) {
	require.Equal(t, "app config context value abc", ContextKey{"abc"}.String())
}
