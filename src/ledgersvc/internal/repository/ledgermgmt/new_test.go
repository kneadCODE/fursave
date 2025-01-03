package ledgermgmt

import (
    "testing"

    "github.com/stretchr/testify/require"
)

func TestNewRepository(t *testing.T) {
    repo := New()

    require.NotNil(t, repo)
}
