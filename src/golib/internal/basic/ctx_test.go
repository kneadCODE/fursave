package basic

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestContextKey_String(t *testing.T) {
	require.Equal(t, "golib:context_key:abc", ContextKey{"abc"}.String())
}
