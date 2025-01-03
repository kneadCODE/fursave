package ledgermgmt

// UseCase represents the use case for ledger management.
type UseCase struct {
	repo Repository // Repository for data storage and retrieval
}

// New creates a new instance of UseCase with the given repository.
func New(repo Repository) UseCase {
	return UseCase{repo: repo}
}
