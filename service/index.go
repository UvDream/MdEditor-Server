package service

import (
	"server/service/article"
	"server/service/file"
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
}

var ServicesGroupApp = new(ServicesGroup)
