teardown:
	${COMPOSE} down

ledgersvc-pg:
	${COMPOSE} up -d ledgersvc-pg
ledgersvc-pg-migrate:
	${COMPOSE} run --rm ledgersvc-pg-migrate sh -c 'sleep 5 && ./migrate -verbose -path /migrations -database $$PG_URL up'
#ledgersvc-pg-migrate-down:
#	${COMPOSE} run --rm pg-migrate sh -c './migrate -verbose -path /migrations -database $$PG_URL down'
ledgersvc-pg-migrate-redo:
	${COMPOSE} run --rm ledgersvc-pg-migrate sh -c './migrate -verbose -path /migrations -database $$PG_URL down'
	pg-migrate-up
ledgersvc-setup: ledgersvc-pg ledgersvc-pg-migrate
ledgersvc-go-vendor:
	${COMPOSE} run --rm ledgersvc-go sh -c 'go mod tidy && go mod vendor'
ledgersvc-go-test:
	${COMPOSE} run --rm ledgersvc-go sh -c 'go test -coverprofile=c.out -failfast -timeout 5m ./...'
ledgersvc-go-api:
	${COMPOSE} run --rm --service-ports ledgersvc-go sh -c 'go run cmd/serverd/*.go'

COMPOSE_BIN := docker compose
COMPOSE := ${COMPOSE_BIN} -f build/docker-compose.yaml -p fursave
