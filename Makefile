teardown:
	${COMPOSE} down

pg:
	${COMPOSE} up -d pg

pg-migrate-up:
	${COMPOSE} run --rm pg-migrate sh -c 'sleep 5 && ./migrate -verbose -path /migrations -database $$PG_URL up'
pg-migrate-down:
	${COMPOSE} run --rm pg-migrate sh -c './migrate -verbose -path /migrations -database $$PG_URL down'
pg-migrate-redo: pg-migrate-down pg-migrate-up

go-api:
	${COMPOSE} run --service-ports --rm go sh -c 'go run cmd/serverd/*.go'
go-vendor:
	${COMPOSE} run --rm go sh -c 'go mod tidy && go mod vendor'
go-test:
	${COMPOSE} run --rm go sh -c 'go test -coverprofile=c.out -failfast -timeout 5m ./...'


COMPOSE_BIN := docker compose
COMPOSE := ${COMPOSE_BIN} -f build/docker-compose.yaml -p fursave
