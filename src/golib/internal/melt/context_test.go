package melt

import (
	"context"
	"testing"

	"github.com/kneadCODE/fursave/src/golib/internal/basic"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func TestContextKey_String(t *testing.T) {
	require.Equal(t, "golib:context_key:melt-zap", zapCtxKey.String())
	require.Equal(t, "golib:context_key:abc", basic.ContextKey{Name: "abc"}.String())
}

func TestZapFromContext(t *testing.T) {
	// Given:
	ctx := context.Background()

	// When:
	l := ZapFromContext(ctx)

	// Then:
	require.Nil(t, l)

	// When:
	newL := zap.NewExample().Sugar()
	ctx = context.WithValue(ctx, zapCtxKey, newL)

	// When:
	l = ZapFromContext(ctx)

	// Then:
	require.EqualValues(t, newL, l)
}

func TestSetZapInContext(t *testing.T) {
	// Given:
	ctx := context.Background()

	// When:
	l := ZapFromContext(ctx)

	// Then:
	require.Nil(t, l)

	// When:
	newL := zap.NewExample().Sugar()
	ctx = SetZapInContext(ctx, newL)

	// When:
	require.EqualValues(t, newL, ZapFromContext(ctx))
}
