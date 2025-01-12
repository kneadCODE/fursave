package telemetry

import (
	"context"

	"github.com/kneadCODE/fursave/src/golib/internal/melt"
)

// - writes it to log.
func CaptureDebugEvent(ctx context.Context, msg string, args ...any) {
	if z := melt.ZapFromContext(ctx); z != nil {
		z.Debugf(msg, args...)
	}
}

// - writes it to log.
func CaptureInfoEvent(ctx context.Context, msg string, args ...any) {
	if z := melt.ZapFromContext(ctx); z != nil {
		z.Infof(msg, args...)
	}
}

// - writes it to log.
func CaptureWarnEvent(ctx context.Context, msg string, args ...any) {
	if z := melt.ZapFromContext(ctx); z != nil {
		z.Warnf(msg, args...)
	}
}

// - writes it to log.
func CaptureErrorEvent(ctx context.Context, err error) {
	if z := melt.ZapFromContext(ctx); z != nil {
		z.Error(err)
	}
}
