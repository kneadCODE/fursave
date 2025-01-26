// Package main implements the main server daemon for the ledger service.
// It initializes configuration, telemetry, and starts an HTTP server.
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/kneadCODE/fursave/src/golib/config"
	"github.com/kneadCODE/fursave/src/golib/executor"
	"github.com/kneadCODE/fursave/src/golib/httpserver"
	"github.com/kneadCODE/fursave/src/golib/telemetry"
)

func main() {
	log.SetOutput(os.Stdout)
	log.Println("Welcome to Fursave")

	ctx, telemetryCleanupF, err := initBase()
	if err != nil {
		log.Panic(err)
	}
	defer telemetryCleanupF()

	s, err := initServer(ctx)
	if err != nil {
		telemetry.CaptureErrorEvent(ctx, err)
		return
	}

	executor.Run(ctx, s.Start)
}

// initBase initializes the basic application components including configuration and telemetry.
// It returns the initialized context, a cleanup function for telemetry, and any error that occurred.
func initBase() (context.Context, func(), error) {
	ctx, err := config.Init()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to initialize App: %w", err)
	}

	ctx, telShutdownF, err := telemetry.Init(ctx)
	if err != nil {
		return nil, nil, err
	}

	return ctx, telShutdownF, nil
}

// initServer creates and configures the HTTP server with the given context.
// It sets up profiling, readiness checks, and REST handlers.
func initServer(ctx context.Context) (*httpserver.Server, error) {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		return nil, fmt.Errorf("unable to parse PORT: %w", err)
	}
	s, err := httpserver.NewServer(
		ctx,
		httpserver.WithPort(port),
		httpserver.WithProfilingHandler(),
		httpserver.WithReadinessHandler(func(_ http.ResponseWriter, r *http.Request) {
			telemetry.CaptureInfoEvent(r.Context(), "readiness called")
		}),
		httpserver.WithRESTHandler(func(chi.Router) {
			// Add routes here.
		}),
	)
	if err != nil {
		return nil, err
	}

	return s, nil
}
