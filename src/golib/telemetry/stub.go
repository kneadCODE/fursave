package telemetry

import (
	"github.com/kneadCODE/fursave/src/golib/internal/melt"
	"go.opentelemetry.io/otel/log/global"
)

var (
	newOTELLoggerProviderStub = melt.NewOTELLoggerProvider
	newZapStub                = melt.NewZap
	setOTELLoggerProviderStub = global.SetLoggerProvider
)
