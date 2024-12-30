teardown:
	${COMPOSE} down --rmi local

golib-setup: P=golib
golib-setup: go-vendor

golib-test: P=golib
golib-test: go-test

ledgersvc-pg-migrate:
	${COMPOSE} run --rm ledgersvc-pg-migrate sh -c 'sleep 5 && ./migrate -verbose -path /migrations -database $$PG_URL up'
ledgersvc-pg-migrate-down:
	${COMPOSE} run --rm ledgersvc-pg-migrate sh -c 'sleep 5 && ./migrate -verbose -path /migrations -database $$PG_URL down'
ledgersvc-pg-migrate-redo: ledgersvc-pg-migrate-down ledgersvc-pg-migrate

ledgersvc-setup: P=ledgersvc
ledgersvc-setup: pg ledgersvc-pg-migrate go-vendor go-gen

ledgersvc-test: P=ledgersvc
ledgersvc-test: go-test

ledgersvc-build-binaries: P=ledgersvc
ledgersvc-build-binaries: go-build-binaries

ledgersvc-serverd:
	${COMPOSE} run --rm --service-ports ledgersvc-go sh -c 'go run cmd/serverd/*.go'

# Reusable commands
go-vendor:
	${COMPOSE} run --rm ${P}-go sh -c 'go mod vendor'
go-clean-vendor:
	${COMPOSE} run --rm ${P}-go sh -c 'go mod tidy && go mod vendor'
go-gen:
	${COMPOSE} run --rm ${P}-go sh -c 'go generate ./...'
go-test:
	${COMPOSE} run --rm ${P}-go sh -c 'go test -coverprofile=coverage.out -failfast -timeout 5m ./...'
go-build-binaries:
	${COMPOSE} run --rm ${P}-go sh -c 'for CMD in `ls cmd`; do (go build -v -o build/binaries/$$CMD ./cmd/$$CMD) done'

pg:
	${COMPOSE} up -d ${P}-pg

# Binaries
COMPOSE_BIN := docker compose
COMPOSE := ${COMPOSE_BIN} -f build/docker-compose.yaml -p fursave
