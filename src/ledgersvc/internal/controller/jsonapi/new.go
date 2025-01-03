package jsonapi

// Controller handles API requests related to ledger management.
type Controller struct {
	ledgerMgmtUC LedgerMgmtUseCase // Use case for ledger management
}

// NewController creates a new Controller with the given ledger management use case.
func NewController(ledgerMgmtUC LedgerMgmtUseCase) Controller {
	return Controller{
		ledgerMgmtUC: ledgerMgmtUC,
	}
}
