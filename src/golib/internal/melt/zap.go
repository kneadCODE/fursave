package melt

import (
	"fmt"

	"go.opentelemetry.io/contrib/bridges/otelzap"
	sdklog "go.opentelemetry.io/otel/sdk/log"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewZap(debugMode bool, loggerProvider *sdklog.LoggerProvider) (*zap.SugaredLogger, error) {
	if debugMode {
		z, err := zap.Config{
			Level: zap.NewAtomicLevelAt(zapcore.DebugLevel),
			// Development: true,
			Encoding: "console",
			EncoderConfig: zapcore.EncoderConfig{
				TimeKey:        "T",
				LevelKey:       "S",
				NameKey:        zapcore.OmitKey,
				CallerKey:      zapcore.OmitKey,
				FunctionKey:    zapcore.OmitKey,
				MessageKey:     "B",
				StacktraceKey:  zapcore.OmitKey,
				LineEnding:     zapcore.DefaultLineEnding,
				EncodeLevel:    zapcore.CapitalColorLevelEncoder,
				EncodeTime:     zapcore.ISO8601TimeEncoder,
				EncodeDuration: zapcore.StringDurationEncoder,
			},
			OutputPaths:      []string{"stdout"},
			ErrorOutputPaths: []string{"stderr"},
		}.Build()
		if err != nil {
			return nil, fmt.Errorf("golib:app:NewZap err initializing zap: %w", err)
		}
		return z.Sugar(), nil
	}

	return zap.New(
		otelzap.NewCore(
			"github.com/kneadCODE/fursave/src/golib/internal/melt",
			otelzap.WithLoggerProvider(loggerProvider),
		),
		zap.IncreaseLevel(zap.InfoLevel),
	).Sugar(), nil
}
