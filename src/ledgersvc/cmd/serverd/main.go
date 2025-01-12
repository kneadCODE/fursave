package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi/v5"
	"github.com/kneadCODE/fursave/src/golib/config"
	"github.com/kneadCODE/fursave/src/golib/httpserver"
	"github.com/kneadCODE/fursave/src/golib/telemetry"
)

func main() {
	log.Println("Welcome to Fursave")

	ctx, err := config.Init()
	if err != nil {
		log.Panicf("Failed to initialize App: %v", err)
	}

	ctx, telShutdownF, err := telemetry.Init(ctx)
	if err != nil {
		log.Panic(err)
	}
	defer telShutdownF()

	s, err := initServer(ctx)
	if err != nil {
		log.Panic(err)
	}

	ctx, cancel := signal.NotifyContext(ctx, syscall.SIGTERM, os.Interrupt, os.Kill)
	defer cancel()
	_ = s.Start(ctx)
}

func initServer(ctx context.Context) (*httpserver.Server, error) {
	s, err := httpserver.NewServer(
		ctx,
		httpserver.WithPort(4000),
		httpserver.WithProfilingHandler(),
		httpserver.WithReadinessHandler(func(http.ResponseWriter, *http.Request) {
			log.Println("readiness called")
		}),
		httpserver.WithRESTHandler(func(rtr chi.Router) {
			rtr.Get("/abc", func(http.ResponseWriter, *http.Request) {
				log.Println("abc called")
			})
		}),
	)
	if err != nil {
		return nil, err
	}

	return s, nil
}
