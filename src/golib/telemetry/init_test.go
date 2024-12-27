package telemetry

import (
	"context"
	"testing"

	"github.com/kneadCODE/fursave/src/golib/config"
	"github.com/stretchr/testify/require"
)

func TestInit(t *testing.T) {
	// Given:
	ctx := context.Background()
	app := config.App{Env: config.EnvDev}
	ctx = config.SetAppInContext(ctx, app)

	// When:
	ctx, shutdown, err := Init(ctx)

	// Then:
	require.NoError(t, err)
	require.NotNil(t, ctx)
	CaptureInfoEvent(ctx, "some msg")
	shutdown()

	// Given:
	app.Env = config.EnvProd
	ctx = config.SetAppInContext(ctx, app)

	// When:
	ctx, shutdown, err = Init(ctx)

	// Then:
	require.NoError(t, err)
	require.NotNil(t, ctx)
	CaptureInfoEvent(ctx, "some msg")
	shutdown()
}
