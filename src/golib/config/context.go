package config

import (
	"context"

	"github.com/kneadCODE/fursave/src/golib/internal/cfg"
)

// AppFromContext retrieves the App from context if exists else return a new App
func AppFromContext(ctx context.Context) App {
	if v, ok := ctx.Value(appCtxKey).(App); ok {
		return v
	}
	return App{}
}

var appCtxKey = cfg.ContextKey{Name: "app-config"}

func setAppInContext(ctx context.Context, cfg App) context.Context {
	return context.WithValue(ctx, appCtxKey, cfg)
}
