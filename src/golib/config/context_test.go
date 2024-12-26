package config

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAppFromContext(t *testing.T) {
	// Given:
	ctx := context.Background()

	// When:
	cfg := AppFromContext(ctx)

	// Then:
	require.EqualValues(t, App{}, cfg)

	// When:
	newCfg := App{Env: EnvDev}
	ctx = context.WithValue(ctx, appCtxKey, newCfg)

	// When:
	cfg = AppFromContext(ctx)

	// Then:
	require.EqualValues(t, newCfg, cfg)
}

func Test_setAppInContext(t *testing.T) {
	// Given:
	ctx := context.Background()

	// When:
	cfg := AppFromContext(ctx)

	// Then:
	require.EqualValues(t, App{}, cfg)

	// When:
	newCfg := App{Env: EnvDev}
	ctx = setAppInContext(ctx, newCfg)

	// When:
	require.EqualValues(t, newCfg, AppFromContext(ctx))
}
