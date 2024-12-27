package config

import (
	"context"

	"go.opentelemetry.io/otel/sdk/resource"
)

type App struct {
	// Env is the environment in which the application is running
	Env Environment
	res *resource.Resource
}

func newAppFromEnv(ctx context.Context) (App, error) {
	res, env, err := newOTELResourceFromEnvStub(ctx)
	if err != nil {
		return App{}, err
	}

	cfg := App{res: res, Env: Environment(env)}

	if err = cfg.Env.IsValid(); err != nil {
		return App{}, err
	}

	return cfg, nil
}
