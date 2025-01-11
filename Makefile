# Variables
COMPOSE_BIN := docker compose
PROJECT_NAME := fursave

# Teardown
teardown-all: golib-teardown ledgersvc-teardown

# Golib Service
COMPOSE_GOLIB := ${COMPOSE_BIN} -f src/golib/build/docker-compose.yaml -p ${PROJECT_NAME}-golib
golib-teardown: COMPOSE_CMD=${COMPOSE_GOLIB}
golib-teardown: teardown

golib-clean-vendor: COMPOSE_CMD=${COMPOSE_GOLIB}
golib-clean-vendor: go-clean-vendor

golib-setup: COMPOSE_CMD=${COMPOSE_GOLIB}
golib-setup: go-vendor

golib-test: COMPOSE_CMD=${COMPOSE_GOLIB}
golib-test: go-test

# Ledgersvc Service
COMPOSE_LEDGERSVC := ${COMPOSE_BIN} -f src/ledgersvc/build/docker-compose.yaml -p ${PROJECT_NAME}-ledgersvc
ledgersvc-teardown: COMPOSE_CMD=${COMPOSE_LEDGERSVC}
ledgersvc-teardown: teardown

ledgersvc-clean-vendor: COMPOSE_CMD=${COMPOSE_LEDGERSVC}
ledgersvc-clean-vendor: go-clean-vendor

ledgersvc-pg:
	${COMPOSE_LEDGERSVC} up -d pg
ledgersvc-pg-migrate:
	${COMPOSE_LEDGERSVC} run --rm pg-migrate sh -c 'sleep 5 && ./migrate -verbose -path /migrations -database $$PG_URL up'
ledgersvc-pg-migrate-down:
	${COMPOSE_LEDGERSVC} run --rm pg-migrate sh -c 'sleep 5 && ./migrate -verbose -path /migrations -database $$PG_URL down'
ledgersvc-pg-migrate-redo: pg-migrate-down pg-migrate

ledgersvc-setup: COMPOSE_CMD=${COMPOSE_LEDGERSVC}
ledgersvc-setup: ledgersvc-pg ledgersvc-pg-migrate go-vendor go-gen

ledgersvc-test: COMPOSE_CMD=${COMPOSE_LEDGERSVC}
ledgersvc-test: go-test

ledgersvc-build-binaries: COMPOSE_CMD=${COMPOSE_LEDGERSVC}
ledgersvc-build-binaries: go-build-binaries

ledgersvc-serverd:
	${COMPOSE_LEDGERSVC} run --rm --service-ports go sh -c 'go run cmd/serverd/*.go'

# Reusable commands
go-vendor:
	${COMPOSE_CMD} run --rm go sh -c 'go mod vendor'
go-clean-vendor:
	${COMPOSE_CMD} run --rm go sh -c 'go mod tidy && go mod vendor'
go-gen:
	${COMPOSE_CMD} run --rm go sh -c 'go generate ./...'
go-test:
	${COMPOSE_CMD} run --rm go sh -c 'go test -coverprofile=coverage.out -failfast -timeout 5m ./...'
go-build-binaries:
	${COMPOSE_CMD} run --rm go sh -c 'for CMD in `ls cmd`; do (go build -v -o build/binaries/$$CMD ./cmd/$$CMD) done'

teardown:
	${COMPOSE_CMD} down --rmi local
