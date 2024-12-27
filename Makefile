teardown:
	${COMPOSE} down --rmi local

golib-go-vendor:
	${COMPOSE} run --rm golib-go sh -c 'go mod vendor'
golib-go-clean-vendor:
	${COMPOSE} run --rm golib-go sh -c 'go mod tidy && go mod vendor'
golib-setup: golib-go-vendor
golib-test:
	${COMPOSE} run --rm golib-go sh -c 'go test -coverprofile=coverage.out -failfast -timeout 5m ./...'

ledgersvc-pg:
	${COMPOSE} up -d ledgersvc-pg
ledgersvc-pg-migrate:
	${COMPOSE} run --rm ledgersvc-pg-migrate sh -c 'sleep 5 && ./migrate -verbose -path /migrations -database $$PG_URL up'
ledgersvc-pg-migrate-down:
	${COMPOSE} run --rm ledgersvc-pg-migrate sh -c 'sleep 5 && ./migrate -verbose -path /migrations -database $$PG_URL down'
ledgersvc-pg-migrate-redo: ledgersvc-pg-migrate-down ledgersvc-pg-migrate
ledgersvc-go-vendor:
	${COMPOSE} run --rm ledgersvc-go sh -c 'go mod vendor'
ledgersvc-go-clean-vendor:
	${COMPOSE} run --rm ledgersvc-go sh -c 'go mod tidy && go mod vendor'
ledgersvc-go-gen:
	${COMPOSE} run --rm ledgersvc-go sh -c 'go generate ./...'
ledgersvc-setup: ledgersvc-pg ledgersvc-pg-migrate ledgersvc-go-vendor ledgersvc-go-gen
ledgersvc-test:
	${COMPOSE} run --rm ledgersvc-go sh -c 'go test -coverprofile=coverage.out -failfast -timeout 5m ./...'
ledgersvc-serverd:
	${COMPOSE} run --rm --service-ports ledgersvc-go sh -c 'go run cmd/serverd/*.go'

COMPOSE_BIN := docker compose
COMPOSE := ${COMPOSE_BIN} -f build/docker-compose.yaml -p fursave
