package system

import "server/service"

type ApiSystemGroup struct {
	UserApi
	BaseApi
	SysApi
}

var (
	articleService = service.ServicesGroupApp.ToArticleService
	userService    = service.ServicesGroupApp.SysUserService
	ledgerService  = service.ServicesGroupApp.LedgersService
)
