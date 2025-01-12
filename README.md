# fursave

## Technology Stack (NOTE: This is a draft)

### 1. Backend

- **Language:** [Go](https://go.dev/) <img height="12.5" src="https://cdn.simpleicons.org/go/00ADD8?viewbox=auto" />
- **Framework:** Go Standard Library
- **APIs:**
  - JSON: [go-chi](https://github.com/go-chi/chi) & [sonic](https://github.com/bytedance/sonic)
  - GraphQL: [gqlgen](https://gqlgen.com/)
  - gRPC: [grpc-go](https://github.com/grpc/grpc-go)
- **Libraries:**
  - Validation: [validator](https://github.com/go-playground/validator)
  - ORM: [sqlboiler](https://github.com/volatiletech/sqlboiler) (PostgreSQL)
  - Performance: [ants](https://github.com/panjf2000/ants)
- **Database Drivers:**
  - PostgreSQL: [pgx](https://github.com/jackc/pgx)
  - MongoDB: [mongo-go-driver](https://github.com/mongodb/mongo-go-driver)
  - Redis: [go-redis](https://github.com/redis/go-redis)
  - Elasticsearch: [go-elasticsearch](https://github.com/elastic/go-elasticsearch)
- **Observability:** zerolog & opentelemetry-go
- **Testing:**
  - Unit: [testify](https://github.com/stretchr/testify)
  - Leak Detection: [goleak](https://github.com/uber-go/goleak)
  - Mocking: [mockery](https://github.com/vektra/mockery)

### 2. Frontend

- Flutter Web & iOS <img height="20" src="https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/flutter/flutter-original.svg" />

### 3. Databases

- RDBMS: [PostgreSQL](https://www.postgresql.org/) <img height="20" src="https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/postgresql/postgresql-original.svg" />
- Document Store: [MongoDB](https://www.mongodb.com/) <img height="20" src="https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/mongodb/mongodb-original.svg" />

- Cache: [Redis](https://redis.io/) <img height="20" src="https://cdn.simpleicons.org/redis/FF4438?viewbox=auto" />
- Search: [Elasticsearch](https://www.elastic.co/elasticsearch) <img height="20" src="https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/elasticsearch/elasticsearch-original.svg" />

### 4. Communication

- **Service-to-Service:**
  - Sync: [gRPC](https://grpc.io/) <img height="25" src="https://raw.githubusercontent.com/marwin1991/profile-technology-icons/refs/heads/main/icons/grpc.png" />
  - Async: [Kafka](https://kafka.apache.org/) <img height="20" src="https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/apachekafka/apachekafka-original.svg" />

- **Frontend-to-Backend:**
  - File Uploads: REST <img height="25" src="https://raw.githubusercontent.com/marwin1991/profile-technology-icons/refs/heads/main/icons/rest.png" />
  - Realtime: WebSocket <img height="20" src="https://raw.githubusercontent.com/marwin1991/profile-technology-icons/refs/heads/main/icons/websocket.png" />
  - Everything else: [GraphQL](https://graphql.org/) <img height="20" src="https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/graphql/graphql-plain.svg" />

### 5. Security

- **Frontend-to-Backend:** OAuth2 <img height="20" src="https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/oauth/oauth-original.svg" /> + OIDC
- **Service-to-Service:** PKI mTLS

### 6. Environments

- **Development:** Local machine, unit tests
- **Integration:** CI pipeline, integration tests
- **Staging:** E2E tests, performance tests, load tests
- **Production:** Live deployment

### 7. Development Tools

- VCS: [Git](https://git-scm.com/) <img height="20" src="https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/git/git-original.svg" />
- VCS Provider: [GitHub](https://github.com) <img height="20" src="https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/github/github-original.svg" />
- Documentation: [Confluence](https://www.atlassian.com/software/confluence) <img height="20" src="https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/confluence/confluence-original.svg" />

- **Local Environment**:
  - Host OS: MacOS
  - IDE: [Visual Studio Code](https://code.visualstudio.com/) <img height="20" src="https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/vscode/vscode-original.svg" />

  - Containerization Host: [Colima](https://github.com/abiosoft/colima)
  - Containerization: [Docker](https://www.docker.com/) & Docker Compose <img height="20" src="https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/docker/docker-original.svg" />
  - Kubernetes Cluster: [k3d](https://k3d.io/)
  - API Client: [Insomnia](https://insomnia.rest/) <img height="20" src="https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/insomnia/insomnia-original.svg" />
  - Postgres GUI Client: [Beekeeper Studio](https://www.beekeeperstudio.io/)
  - MongoDB GUI Client: [MongoDB Compass](https://www.mongodb.com/products/tools/compass)
  - Redis GUI Client: [Redis Insight](https://redis.io/insight/)
  - Elasticsearch GUI Client: [Kibana](https://www.elastic.co/kibana)
  - Kafka GUI Client: [kafka-ui](https://github.com/provectus/kafka-ui)
  - Go:
    - [Go Workspace](https://go.dev/doc/tutorial/workspaces)
    - Hot Reload: [air](https://github.com/air-verse/air)
    - Linter: [golangci-lint](https://github.com/golangci/golangci-lint)

### 8. CI/CD(Delivery)

- **CI:** [GitHub Actions](https://github.com/features/actions) <img height="20" src="https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/githubactions/githubactions-original.svg" />
- **Quality Gates:**
  - Analysis: [SonarQube](https://www.sonarsource.com/products/sonarqube/) <img height="20" src="https://cdn.simpleicons.org/sonarcloud/F3702A?viewbox=auto" />
  - SAST: [Snyk](https://snyk.io/) <img height="20" src="https://cdn.simpleicons.org/snyk/4C4A73?viewbox=auto" />
  - DAST: [OWASP ZAP](https://www.zaproxy.org/) <img height="20" src="https://cdn.simpleicons.org/owasp/000000?viewbox=auto" />
  - Container Scanning: [Trivy](https://trivy.dev/) <img height="20" src="https://cdn.simpleicons.org/trivy/1904DA?viewbox=auto" />
  - SBOM: [Syft](https://github.com/anchore/syft)
- **Artifacts:** GitHub Packages
- **CD:** [GitHub Actions](https://github.com/features/actions) <img height="20" src="https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/githubactions/githubactions-original.svg" /> with [Kustomize](https://kustomize.io/)

### 9. Deployment Environment

#### Ideal Setup

- **Environment**:
  - Orchestration: [Kubernetes](https://kubernetes.io/) <img height="20" src="https://cdn.simpleicons.org/kubernetes/326CE5?viewbox=auto" /> on AKS
  - Service Mesh: Cilium <img height="20" src="https://cdn.simpleicons.org/cilium/F8C517?viewbox=auto" />  & Hubble
  - OTEL Collector: [Grafana Alloy](https://grafana.com/oss/alloy-opentelemetry-collector/)
  - Node Updates: [kured](https://kured.dev/)
- **Observability**:
  - Framework: [OpenTelemetry](https://opentelemetry.io/) <img height="20" src="https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/opentelemetry/opentelemetry-original.svg" />
  - Platform:
    - Visualization/Dashboard: [Grafana](https://grafana.com/oss/grafana/) <img height="20" src="https://raw.githubusercontent.com/marwin1991/profile-technology-icons/refs/heads/main/icons/grafana.png" />
    - Logs: [Grafana Loki](https://grafana.com/oss/loki/) <img height="20" src="https://raw.githubusercontent.com/marwin1991/profile-technology-icons/refs/heads/main/icons/loki.png" />
    - Traces: [Grafana Tempo](https://grafana.com/oss/tempo/)
    - Metrics: [Grafana Mimir](https://grafana.com/oss/mimir/)
    - Profiling: [Grafana Pyroscope](https://grafana.com/oss/pyroscope/)
    - RUM: [Grafana Faro](https://grafana.com/oss/faro/)
  - Incident Management: [Grafana Oncall](https://grafana.com/products/cloud/oncall/)

#### Budget Setup

- Bare metal with k3s

Rest of the tools ???

// go-discovery: ignore
// This is a private/experimental module
