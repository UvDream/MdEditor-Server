package ledger

import "server/service"

type APILedgerGroup struct {
	ApiLedger
}

var ledgerService = service.ServicesGroupApp.LedgerServiceGroup
