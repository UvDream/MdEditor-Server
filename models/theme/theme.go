package theme

import (
	"server/models"
	"server/models/system"
)

type Theme struct {
	models.Model
	//名字
	Name string `json:"name" gorm:"type:longblob;"`
	//描述
	Description string `json:"description" gorm:"type:varchar(100);"`
	//简略图
	Thumbnail string `json:"thumbnail" gorm:"type:varchar(100);"`
	//是否公开
	IsPublic bool `json:"is_public" gorm:"type:tinyint(1);"`
	//作者ID
	UserID string `json:"user_id" gorm:"type:varchar(100);comment:作者的UUID"`
	//作者
	User system.User `json:"user"`
	// 主题
	Theme string `json:"theme" gorm:"type:longblob" `
	//额外主题
	DarkTheme string `json:"dark_theme" gorm:"type:longblob"`
}
