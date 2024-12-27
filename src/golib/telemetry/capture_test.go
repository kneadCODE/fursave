package telemetry

import (
	"context"
	"errors"
	"testing"

	"github.com/kneadCODE/fursave/src/golib/internal/melt"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
)

func TestCaptureDebugEvent(t *testing.T) {
	// Given:
	ctx := context.Background()

	// When && Then:
	CaptureDebugEvent(ctx, "some message")
	CaptureDebugEvent(ctx, "some message: %s, %d, %f, %t", "s", 1, float32(32), true)
	CaptureInfoEvent(ctx, "some message")
	CaptureInfoEvent(ctx, "some message: %s, %d, %f, %t", "s", 1, float32(32), true)
	CaptureWarnEvent(ctx, "some message")
	CaptureWarnEvent(ctx, "some message: %s, %d, %f, %t", "s", 1, float32(32), true)
	CaptureErrorEvent(ctx, errors.New("some err"))

	// Given:
	zapLogger, err := newZapStub(true, nil)
	require.NoError(t, err)
	ctx = melt.SetZapInContext(ctx, zapLogger)

	// When && Then:
	CaptureDebugEvent(ctx, "some message")
	CaptureDebugEvent(ctx, "some message: %s, %d, %f, %t", "s", 1, float32(32), true)
	CaptureInfoEvent(ctx, "some message")
	CaptureInfoEvent(ctx, "some message: %s, %d, %f, %t", "s", 1, float32(32), true)
	CaptureWarnEvent(ctx, "some message")
	CaptureWarnEvent(ctx, "some message: %s, %d, %f, %t", "s", 1, float32(32), true)
	CaptureErrorEvent(ctx, errors.New("some err"))

	// Given:
	res := resource.NewWithAttributes(semconv.SchemaURL, semconv.DeploymentEnvironment("development"))
	lp, err := melt.NewOTELLoggerProvider(res)
	require.NoError(t, err)
	zapLogger, err = newZapStub(false, lp)
	require.NoError(t, err)
	ctx = melt.SetZapInContext(ctx, zapLogger)

	// When && Then:
	CaptureDebugEvent(ctx, "some message")
	CaptureDebugEvent(ctx, "some message: %s, %d, %f, %t", "s", 1, float32(32), true)
	CaptureInfoEvent(ctx, "some message")
	CaptureInfoEvent(ctx, "some message: %s, %d, %f, %t", "s", 1, float32(32), true)
	CaptureWarnEvent(ctx, "some message")
	CaptureWarnEvent(ctx, "some message: %s, %d, %f, %t", "s", 1, float32(32), true)
	CaptureErrorEvent(ctx, errors.New("some err"))
	require.NoError(t, lp.ForceFlush(ctx))
}
