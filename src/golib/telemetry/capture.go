package telemetry

import (
	"context"

	"github.com/kneadCODE/fursave/src/golib/internal/melt"
)

// CaptureDebugEvent captures the debug event and:-
// - writes it to log
func CaptureDebugEvent(ctx context.Context, msg string, args ...interface{}) {
	if z := melt.ZapFromContext(ctx); z != nil {
		z.Debugf(msg, args...)
	}
}

// CaptureInfoEvent captures the informational event and:-
// - writes it to log
func CaptureInfoEvent(ctx context.Context, msg string, args ...interface{}) {
	if z := melt.ZapFromContext(ctx); z != nil {
		z.Infof(msg, args...)
	}
}

// CaptureWarnEvent captures the warning event and:-
// - writes it to log
func CaptureWarnEvent(ctx context.Context, msg string, args ...interface{}) {
	if z := melt.ZapFromContext(ctx); z != nil {
		z.Warnf(msg, args...)
	}
}

// CaptureErrorEvent captures the error event and:-
// - writes it to log
func CaptureErrorEvent(ctx context.Context, err error) {
	if z := melt.ZapFromContext(ctx); z != nil {
		z.Error(err)
	}
}
