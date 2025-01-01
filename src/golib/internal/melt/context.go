package melt

import (
	"context"

	"github.com/kneadCODE/fursave/src/golib/internal/basic"
	"go.uber.org/zap"
)

var zapCtxKey = basic.ContextKey{Name: "melt-zap"}

func SetZapInContext(ctx context.Context, l *zap.SugaredLogger) context.Context {
	return context.WithValue(ctx, zapCtxKey, l)
}

func ZapFromContext(ctx context.Context) *zap.SugaredLogger {
	if v, ok := ctx.Value(zapCtxKey).(*zap.SugaredLogger); ok {
		return v
	}
	return nil
}
