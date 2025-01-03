package ledgermgmt

import (
    "testing"

    "github.com/stretchr/testify/require"
)

type mockRepository struct{}

func TestNewUseCase(t *testing.T) {
    repo := mockRepository{}
    useCase := New(repo)

    require.NotNil(t, useCase)
    require.Equal(t, repo, useCase.repo)
}
