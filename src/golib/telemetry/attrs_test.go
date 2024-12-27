package telemetry

import (
	"context"
	"testing"

	"github.com/kneadCODE/fursave/src/golib/internal/melt"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
)

func TestWithAttrs(t *testing.T) {
	// Given:
	ctx := context.Background()
	res := resource.NewWithAttributes(semconv.SchemaURL, semconv.DeploymentEnvironment("development"))
	lp, err := melt.NewOTELLoggerProvider(res)
	require.NoError(t, err)
	zapLogger, err := newZapStub(false, lp)
	require.NoError(t, err)
	ctx = melt.SetZapInContext(ctx, zapLogger)

	// When:
	ctx = WithAttrs(ctx,
		"string", "str",
		"int", 1,
		"float", float32(32),
		"bool", true,
	)

	// Then:
	CaptureInfoEvent(ctx, "test with attrs: %s", "hello world")
	require.NoError(t, lp.ForceFlush(ctx))
}
