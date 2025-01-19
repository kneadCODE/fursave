# ========================================================
# Fursave Project Makefile
# ========================================================

# ---------------- Global Variables ----------------------
COMPOSE_BIN := docker compose
PROJECT_NAME := fursave
TERRAFORM_BIN := terraform
K3D_BIN := k3d

# ---------------- Colors & Formatting -------------------
GREEN  := \033[0;32m
YELLOW := \033[0;33m
RESET  := \033[0m

# ========================================================
# Project-wide Configuration
# ========================================================
.PHONY: help
help:
	@echo -e "${GREEN}Fursave Project Makefile${RESET}"
	@echo -e "${YELLOW}Top-Level Targets:${RESET}"
	@echo -e "  ${BLUE}make help${RESET}         - Show this help message"
	@echo -e "  ${BLUE}make teardown-all${RESET} - Teardown all services"
	@echo ""
	@echo -e "${YELLOW}Go Library Service Targets:${RESET}"
	@echo -e "  ${BLUE}make golib-setup${RESET}       - Setup Go library service"
	@echo -e "  ${BLUE}make golib-teardown${RESET}    - Teardown Go library service"
	@echo -e "  ${BLUE}make golib-test${RESET}        - Run tests for Go library service"
	@echo ""
	@echo -e "${YELLOW}Ledger Service Targets:${RESET}"
	@echo -e "  ${BLUE}make ledgersvc-setup${RESET}         - Setup Ledger service and dependencies"
	@echo -e "  ${BLUE}make ledgersvc-teardown${RESET}      - Teardown Ledger service"
	@echo -e "  ${BLUE}make ledgersvc-test${RESET}          - Run tests for Ledger service"
	@echo -e "  ${BLUE}make ledgersvc-pg${RESET}            - Start PostgreSQL service"
	@echo -e "  ${BLUE}make ledgersvc-pg-migrate${RESET}    - Run database migrations"
	@echo -e "  ${BLUE}make ledgersvc-pg-migrate-down${RESET} - Rollback database migrations"
	@echo ""
	@echo -e "${YELLOW}Kubernetes Targets:${RESET}"
	@echo -e "  ${BLUE}make setup-k8s-cluster${RESET}  - Create local k3d cluster"
	@echo -e "  ${BLUE}make setup-k8s-ns${RESET} - Setup k8s namespaces"
	@echo -e "  ${BLUE}make teardown-k8s-cluster${RESET} - Delete local k3d cluster"
	@echo -e "  ${BLUE}make redo-k8s${RESET} - Delete and re-create the k3d cluster and the namespaces"

# ========================================================
# Service-Specific Compose Configurations
# ========================================================
# Go Library Service
COMPOSE_GOLIB := ${COMPOSE_BIN} -f src/golib/build/docker-compose.yaml -p ${PROJECT_NAME}-golib

# Ledger Service
COMPOSE_LEDGERSVC := ${COMPOSE_BIN} -f src/ledgersvc/build/docker-compose.yaml -p ${PROJECT_NAME}-ledgersvc

# ========================================================
# Top-Level Management Targets
# ========================================================
teardown-all: golib-teardown ledgersvc-teardown

# ========================================================
# Go Library Service Targets
# ========================================================
golib-teardown: COMPOSE_CMD=${COMPOSE_GOLIB}
golib-teardown: teardown

golib-clean-vendor: COMPOSE_CMD=${COMPOSE_GOLIB}
golib-clean-vendor: go-clean-vendor

golib-setup: COMPOSE_CMD=${COMPOSE_GOLIB}
golib-setup: go-vendor

golib-test: COMPOSE_CMD=${COMPOSE_GOLIB}
golib-test: go-test

# ========================================================
# Ledger Service Targets
# ========================================================
# PostgreSQL-specific targets
ledgersvc-pg:
	${COMPOSE_LEDGERSVC} up -d pg

ledgersvc-pg-migrate:
	${COMPOSE_LEDGERSVC} run --rm pg-migrate sh -c 'sleep 5 && ./migrate -verbose -path /migrations -database $$PG_URL up'

ledgersvc-pg-migrate-down:
	${COMPOSE_LEDGERSVC} run --rm pg-migrate sh -c 'sleep 5 && ./migrate -verbose -path /migrations -database $$PG_URL down'

ledgersvc-pg-migrate-redo: ledgersvc-pg-migrate-down ledgersvc-pg-migrate

# Service-level targets
ledgersvc-teardown: COMPOSE_CMD=${COMPOSE_LEDGERSVC}
ledgersvc-teardown: teardown

ledgersvc-clean-vendor: COMPOSE_CMD=${COMPOSE_LEDGERSVC}
ledgersvc-clean-vendor: go-clean-vendor

ledgersvc-setup: COMPOSE_CMD=${COMPOSE_LEDGERSVC}
ledgersvc-setup: ledgersvc-pg ledgersvc-pg-migrate go-vendor go-gen

ledgersvc-test: COMPOSE_CMD=${COMPOSE_LEDGERSVC}
ledgersvc-test: go-test

ledgersvc-build-binaries: COMPOSE_CMD=${COMPOSE_LEDGERSVC}
ledgersvc-build-binaries: go-build-binaries

ledgersvc-serverd:
	${COMPOSE_LEDGERSVC} run --rm --service-ports go sh -c 'go run cmd/serverd/*.go'

# ========================================================
# Reusable Go Commands
# ========================================================
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

# ========================================================
# Kubernetes Targets
# ========================================================
.PHONY: setup-k8s-cluster setup-k8s-ns setup-k8s-ns

setup-k8s-cluster:
	${K3D_BIN} cluster create ${PROJECT_NAME}

setup-k8s-ns:
	cd build/k8s/cluster-operations/namespaces/dev && \
	TF_VAR_k8s_config_context=k3d-${PROJECT_NAME} ${TERRAFORM_BIN} init && \
	TF_VAR_k8s_config_context=k3d-${PROJECT_NAME} ${TERRAFORM_BIN} validate && \
	TF_VAR_k8s_config_context=k3d-${PROJECT_NAME} ${TERRAFORM_BIN} apply -auto-approve

redo-k8s: teardown-k8s-cluster setup-k8s-cluster setup-k8s-ns

# ========================================================
# Cleanup Targets
# ========================================================
.PHONY: teardown-k8s-cluster

teardown:
	${COMPOSE_CMD} down --rmi local

teardown-k8s-cluster:
	${K3D_BIN} cluster delete ${PROJECT_NAME}

# ========================================================
# End of Makefile
# ========================================================
