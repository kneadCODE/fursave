package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/kneadCODE/fursave/src/golib/config"
	"github.com/kneadCODE/fursave/src/golib/httpserver"
	"github.com/kneadCODE/fursave/src/golib/telemetry"
)

func main() {
	log.Println("Welcome to Fursave")
	ctx := context.Background()

	ctx, err := config.Init()
	if err != nil {
		log.Fatalf("Failed to initialize App: %v", err)
	}

	ctx, telShutdownF, err := telemetry.Init(ctx)
	defer telShutdownF()

	s, err := initServer(ctx)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := signal.NotifyContext(ctx, syscall.SIGTERM, os.Interrupt, os.Kill)
	defer cancel()
	s.Start(ctx)
}

func initServer(ctx context.Context) (*httpserver.Server, error) {
	s, err := httpserver.NewServer(
		ctx,
		httpserver.WithPort(4000),
		httpserver.WithProfilingHandler(),
		httpserver.WithReadinessHandler(func(w http.ResponseWriter, r *http.Request) {
			log.Println("readiness called")
		}),
		httpserver.WithRESTHandler(jsonAPIHandler),
	)
	if err != nil {
		return nil, err
	}

	return s, nil
}
