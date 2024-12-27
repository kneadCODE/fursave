package melt

import (
	"go.opentelemetry.io/otel/exporters/stdout/stdoutlog"
	sdklog "go.opentelemetry.io/otel/sdk/log"
	"go.opentelemetry.io/otel/sdk/resource"
)

func NewOTELLoggerProvider(res *resource.Resource) (*sdklog.LoggerProvider, error) {
	logExporter, err := stdoutlog.New()
	if err != nil {
		return nil, err
	}

	loggerProvider := sdklog.NewLoggerProvider(
		sdklog.WithResource(res),
		// sdklog.WithProcessor(), TODO: Should we add redaction here?
		sdklog.WithProcessor(sdklog.NewBatchProcessor(logExporter)),
	)

	return loggerProvider, nil
}
