package ledger

import (
	"server/models"
	"server/models/system"
)

// LedgerCategory 分类
type LedgerCategory struct {
	models.Model
	// 分类名称
	Name string `json:"name"`
	// 分类缩略图
	Thumbnail string `json:"thumbnail"`
	//分类图标类型 图片/icon/文字
	IconType string `json:"icon_type"`
	// 分类描述
	Description string `json:"description"`
	// 分类创建者
	User system.User `json:"creator"`
	// 分类创建者ID
	UserID string `json:"creator_id"`
	//父元素ID
	ParentID string `json:"parent_id"`
	//	类型 支出/收入....
	Type string `json:"type"`
}
