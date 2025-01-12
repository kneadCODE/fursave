package melt

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
)

func Test_NewZap(t *testing.T) {
	// Given:.
	res := resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.DeploymentEnvironment("development"),
	)

	// When:.
	l, err := NewZap(true, nil)

	// Then:.
	require.NoError(t, err)
	require.NotNil(t, l)
	l.Info("testing")

	// When:.
	lp, err := NewOTELLoggerProvider(res)
	require.NoError(t, err)

	l, err = NewZap(false, lp)

	// Then:.
	require.NoError(t, err)
	require.NotNil(t, l)
	l.Info("testing")
	require.NoError(t, lp.ForceFlush(context.Background()))
}
