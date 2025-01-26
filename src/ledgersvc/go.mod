module github.com/kneadCODE/fursave/src/ledgersvc

// go-discovery: ignore
// This is a private/experimental moduleT

go 1.23.4

require (
	github.com/friendsofgo/errors v0.9.2
	github.com/go-chi/chi/v5 v5.2.0
	github.com/kneadCODE/fursave/src/golib v0.0.0
	github.com/stretchr/testify v1.10.0
	github.com/volatiletech/null/v8 v8.1.2
	github.com/volatiletech/sqlboiler/v4 v4.17.1
	github.com/volatiletech/strmangle v0.0.8
)

replace github.com/kneadCODE/fursave/src/golib => ../golib

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/ericlagergren/decimal v0.0.0-20190420051523-6335edbaa640 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/gofrs/uuid v4.2.0+incompatible // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/lib/pq v1.10.6 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/spf13/cast v1.5.0 // indirect
	github.com/volatiletech/inflect v0.0.1 // indirect
	github.com/volatiletech/randomize v0.0.1 // indirect
	go.opentelemetry.io/auto/sdk v1.1.0 // indirect
	go.opentelemetry.io/contrib/bridges/otelzap v0.8.0 // indirect
	go.opentelemetry.io/otel v1.33.0 // indirect
	go.opentelemetry.io/otel/exporters/stdout/stdoutlog v0.9.0 // indirect
	go.opentelemetry.io/otel/log v0.9.0 // indirect
	go.opentelemetry.io/otel/metric v1.33.0 // indirect
	go.opentelemetry.io/otel/sdk v1.33.0 // indirect
	go.opentelemetry.io/otel/sdk/log v0.9.0 // indirect
	go.opentelemetry.io/otel/trace v1.33.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	go.uber.org/zap v1.27.0 // indirect
	golang.org/x/sys v0.28.0 // indirect
	golang.org/x/xerrors v0.0.0-20220609144429-65e65417b02f // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
