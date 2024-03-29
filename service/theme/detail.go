package theme

import (
	code2 "server/code"
	"server/global"
	"server/models/theme"
)

// GetThemeDetailService 获取主题详情
func (*ThemesService) GetThemeDetailService(id string) (t theme.Theme, code int, err error) {
	db := global.DB
	if err = db.Where("id = ?", id).Preload("User").First(&t).Error; err != nil {
		return t, code2.ErrorGetTheme, err
	}
	return t, code2.SUCCESS, err
}
