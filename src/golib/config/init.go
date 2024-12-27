package config

import (
	"context"
	"log"
	"os"
)

// Init initializes the App and returns
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
