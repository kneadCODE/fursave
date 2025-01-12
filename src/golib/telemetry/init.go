package telemetry

import (
	"context"
	"log"
	"os"
	"sync"
	"time"

	"github.com/kneadCODE/fursave/src/golib/config"
	"github.com/kneadCODE/fursave/src/golib/internal/melt"
	sdklog "go.opentelemetry.io/otel/sdk/log"
	"go.uber.org/zap"
)

// - err if there was an error during initialization.
func Init(ctx context.Context) (newCtx context.Context, shutdown func(), err error) {
	basicLogger := log.New(os.Stdout, "", log.LstdFlags)
	newCtx = ctx
	app := config.AppFromContext(ctx)

	basicLogger.Println("Initializing OTEL Logger provider...")
	otelLoggerP, err := newOTELLoggerProviderStub(app.Res)
	if err != nil {
		return
	}
	setOTELLoggerProviderStub(otelLoggerP)
	basicLogger.Println("OTEL Logger provider initialized")

	basicLogger.Println("Initializing Zap...")
	zapLogger, err := newZapStub(app.Env == config.EnvDev, otelLoggerP)
	if err != nil {
		return
	}
	newCtx = melt.SetZapInContext(newCtx, zapLogger)
	zapLogger.Info("Zap initialized")

	shutdown = shutdownFunc(zapLogger, otelLoggerP)
	return
}

func shutdownFunc(
	zapLogger *zap.SugaredLogger,
	otelLoggerP *sdklog.LoggerProvider,
) func() {
	return func() {
		zapLogger.Info("Shutting down telemetry...")

		basicLogger := log.New(os.Stdout, "", log.LstdFlags)

		cancelCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()

		var wg sync.WaitGroup

		wg.Add(1)
		go func() {
			defer wg.Done()
			basicLogger.Println("Shutting down zap...")
			_ = zapLogger.Sync() // Intentionally ignoring err because we zap has a bug where it always returns err here.
			basicLogger.Println("Zap shutdown complete. Now shuting down OTEL Logger provider...")
			if err := otelLoggerP.Shutdown(cancelCtx); err != nil {
				basicLogger.Printf("OTEL Logger provider shutdown failed: %s", err.Error())
			} else {
				basicLogger.Println("OTEL Logger provider shutdown complete")
			}
		}()

		wg.Wait()

		basicLogger.Println("Telemetry shutdown complete")
	}
}
