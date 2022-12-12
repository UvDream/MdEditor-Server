package system

import "server/service"

type ApiSystemGroup struct {
	UserApi
	BaseApi
	SysApi
}

var (
	articleService = service.ServicesGroupApp.ArticleServiceGroup.ToArticleService
	userService    = service.ServicesGroupApp.SystemServiceGroup.SysUserService
	ledgerService  = service.ServicesGroupApp.LedgerServiceGroup
)
