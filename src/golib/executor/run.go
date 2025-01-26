package executor

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/kneadCODE/fursave/src/golib/telemetry"
)

// Runnable represents a long-running service function that can be started
// and stopped via context cancellation.
type Runnable func(ctx context.Context) error

// Run executes multiple services concurrently and manages their lifecycle.
// It will:
// - Start all services in separate goroutines
// - Handle OS signals (SIGTERM, SIGINT) for graceful shutdown
// - Stop all services if any one of them encounters an error
//
// The function blocks until all services have completed or the context is cancelled.
func Run(ctx context.Context, runnables ...Runnable) {
	if len(runnables) == 0 {
		telemetry.CaptureWarnEvent(ctx, "No services to run")
		return
	}

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGTERM, os.Interrupt) // Unable to use signal.NotifyContext because there is no way to log what signal was received.
	defer signal.Stop(sigChan)

	// Add signal monitoring goroutine.
	go func() {
		sig := <-sigChan
		telemetry.CaptureInfoEvent(ctx, "Received shutdown signal: %v", sig)
		cancel()
	}()

	var wg sync.WaitGroup
	wg.Add(len(runnables))

	for idx := range runnables {
		r := runnables[idx] // https://golang.org/doc/faq#closures_and_goroutines
		go func() {
			defer wg.Done()

			if err := r(ctx); err != nil {
				telemetry.CaptureErrorEvent(ctx, err)
				telemetry.CaptureInfoEvent(ctx, "Canceling all services...")
				cancel() // if one service encounters error, terminate everything else.
			}
		}()
	}

	wg.Wait()
	telemetry.CaptureInfoEvent(ctx, "All services shut down")
}
