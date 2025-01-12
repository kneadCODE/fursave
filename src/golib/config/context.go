package config

import (
	"context"

	"github.com/kneadCODE/fursave/src/golib/internal/basic"
)

// AppFromContext retrieves the App from context if exists else return a new App.
func AppFromContext(ctx context.Context) App {
	if v, ok := ctx.Value(appCtxKey).(App); ok {
		return v
	}
	return App{}
}

// SetAppInContext sets App in the given context.
func SetAppInContext(ctx context.Context, cfg App) context.Context {
	return context.WithValue(ctx, appCtxKey, cfg)
}

var appCtxKey = basic.ContextKey{Name: "app-config"}
