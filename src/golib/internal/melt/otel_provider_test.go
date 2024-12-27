package melt

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/otel/log/global"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
)

func TestNewOTELLoggerProvider(t *testing.T) {
	// Given:
	// Given:
	res := resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.DeploymentEnvironment("development"),
	)

	// When:
	lp, err := NewOTELLoggerProvider(res)

	// Then:
	require.NoError(t, err)
	require.NotNil(t, lp)
	global.SetLoggerProvider(lp)
	require.NoError(t, lp.ForceFlush(context.Background()))
}
