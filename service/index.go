package service

import (
	"server/service/admin"
	"server/service/article"
	"server/service/file"
	"server/service/halo"
	"server/service/ledger"
	"server/service/system"
	"server/service/theme"
)

type ServicesGroup struct {
	system.JWTService
	system.SysConfigService
	system.SysUserService
	article.ToArticleService
	article.TagService
	article.CategoryService
	file.FilesService
	theme.ThemesService
	ledger.LedgersService
	halo.ServiceHalo
	admin.UserService
	admin.LedgerAdminService
}

var ServicesGroupApp = new(ServicesGroup)
