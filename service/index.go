package service

import (
	"server/service/article"
	"server/service/file"
	"server/service/halo"
	"server/service/ledger"
	"server/service/system"
	"server/service/theme"
)

type ServicesGroup struct {
	SystemServiceGroup  system.SysServiceGroup
	ArticleServiceGroup article.ArticlesServiceGroup
	FileServiceGroup    file.FilesServiceGroup
	ThemeServiceGroup   theme.ThemesServiceGroup
	LedgerServiceGroup  ledger.LedgersServiceGroup
	HaloServiceGroup    halo.ServiceHaloGroup
}

var ServicesGroupApp = new(ServicesGroup)
