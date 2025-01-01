package config

import (
	"context"
	"log"
	"os"
)

// Init initializes the application context and configuration.
// It sets up a basic logger, logs the initialization steps, and loads the application
// configuration from environment variables. If the configuration is successfully loaded,
// it is stored in the context.
//
// Returns:
// - ctx: The application context with the configuration set.
// - err: An error if the configuration could not be loaded.
func Init() (ctx context.Context, err error) {
	ctx = context.Background()
	basicLogger := log.New(os.Stdout, "", log.LstdFlags)

	basicLogger.Println("Starting App initialization...")

	basicLogger.Println("Initializing App from env...")
	cfg, err := newAppFromEnv(ctx)
	if err != nil {
		return
	}
	ctx = SetAppInContext(ctx, cfg)
	basicLogger.Println("App initialized")
	return
}
