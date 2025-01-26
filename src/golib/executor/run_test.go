package executor

import (
	"context"
	"errors"
	"sync/atomic"
	"syscall"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestRun(t *testing.T) {
	t.Run("empty runnables", func(*testing.T) {
		Run(context.Background())
		// Test passes if Run returns without panic.
	})

	t.Run("successful execution", func(t *testing.T) {
		var counter atomic.Int32
		r1 := func(ctx context.Context) error {
			counter.Add(1)
			<-ctx.Done()
			return nil
		}
		r2 := func(ctx context.Context) error {
			counter.Add(1)
			<-ctx.Done()
			return nil
		}

		ctx, cancel := context.WithCancel(context.Background())
		go func() {
			time.Sleep(100 * time.Millisecond)
			cancel()
		}()

		Run(ctx, r1, r2)
		require.Equal(t, int32(2), counter.Load(), "both runnables should execute")
	})

	t.Run("error propagation", func(t *testing.T) {
		expectedErr := errors.New("test error")
		var executed atomic.Bool

		r1 := func(context.Context) error {
			return expectedErr
		}
		r2 := func(ctx context.Context) error {
			<-ctx.Done()
			executed.Store(true)
			return nil
		}

		Run(context.Background(), r1, r2)
		require.True(t, executed.Load(), "second runnable should be cancelled")
	})

	t.Run("signal handling", func(t *testing.T) {
		var counter atomic.Int32
		r := func(ctx context.Context) error {
			<-ctx.Done()
			counter.Add(1)
			return nil
		}

		ctx := context.Background()
		go func() {
			time.Sleep(100 * time.Millisecond)
			_ = syscall.Kill(syscall.Getpid(), syscall.SIGTERM)
		}()

		Run(ctx, r, r)
		require.Equal(t, int32(2), counter.Load(), "both runnables should shut down")
	})

	t.Run("concurrent execution", func(t *testing.T) {
		start := time.Now()
		r := func(context.Context) error {
			time.Sleep(100 * time.Millisecond)
			return nil
		}

		Run(context.Background(), r, r, r)
		elapsed := time.Since(start)

		// All runnables should execute concurrently, taking ~100ms total.
		require.Less(t, elapsed, 150*time.Millisecond, "runnables should execute concurrently")
	})
}
