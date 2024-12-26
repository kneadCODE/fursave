package cfg

import (
	"go.opentelemetry.io/otel/sdk/resource"
)

var (
	withTelemetrySDKStub = resource.WithTelemetrySDK
	withOSStub           = resource.WithOS
	withHostStub         = resource.WithHost
	withContainerStub    = resource.WithContainer
	withProcessStub      = resource.WithProcess
)
