package jsonapi

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type mockUseCase struct{}

func TestNewController(t *testing.T) {
	useCase := mockUseCase{}
	controller := NewController(useCase)

	require.NotNil(t, controller)
	require.Equal(t, useCase, controller.ledgerMgmtUC)
}
