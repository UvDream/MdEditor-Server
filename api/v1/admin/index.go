package admin

import "server/service"

type ApiAdminGroup struct {
	UserAdminApi
	LedgerAdminApi
}

var (
	userAdminService   = service.ServicesGroupApp.UserService
	ledgerAdminService = service.ServicesGroupApp.LedgerAdminService
)
