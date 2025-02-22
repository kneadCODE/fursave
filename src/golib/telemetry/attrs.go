package telemetry

import (
	"context"

	"github.com/kneadCODE/fursave/src/golib/internal/melt"
)

func WithAttrs(ctx context.Context, args ...any) context.Context {
	return melt.SetZapInContext(
		ctx,
		melt.ZapFromContext(ctx).With(args...),
	)
}
