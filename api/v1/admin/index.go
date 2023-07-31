package admin

import "server/service"

type ApiAdminGroup struct {
	UserAdminApi
	LedgerAdminApi
}

var (
	userAdminService   = service.ServicesGroupApp.AdminServiceGroup.UserService
	ledgerAdminService = service.ServicesGroupApp.AdminServiceGroup.LedgerAdminService
)
