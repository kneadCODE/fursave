package main

import (
	"context"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/kneadCODE/fursave/src/golib/config"
	"github.com/kneadCODE/fursave/src/golib/httpserver"
)

func main() {
	log.Println("Welcome to Fursave")
	ctx := context.Background()

	_, err := config.Init()
	if err != nil {
		log.Fatalf("Failed to initialize App: %v", err)
	}

	s, err := initServer(ctx)
	if err != nil {
		log.Fatal(err)
	}

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
		httpserver.WithRESTHandler(func(rtr chi.Router) {
			rtr.Get("/abc", func(w http.ResponseWriter, r *http.Request) {
				log.Println("abc called")
			})
		}),
	)
	if err != nil {
		return nil, err
	}

	return s, nil
}
