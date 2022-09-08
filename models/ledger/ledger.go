package ledger

import (
	"server/models"
	"server/models/system"
)

type Ledger struct {
	models.Model
	//  账本名称
	Name string `json:"name"`
	//	缩略图
	Thumbnail string `json:"thumbnail"`
	//	账本描述
	Description string `json:"description"`
	//	账本类型
	Type string `json:"type"`
	//	账本状态
	Status string `json:"status"`
	//	账本创建者
	Creator system.User `json:"creator"`
	//	账本创建者ID
	CreatorID string `json:"creator_id"`
	//	账本成员ID
	User []system.User `json:"user" gorm:"many2many:ledger_users;"`
	//	账本成员数量
	MemberCount int `json:"member_count" gorm:"-"`
	// 分类
	Categories []LedgerCategory `json:"categories" gorm:"many2many:ledger_categories;"`
	//	标签
	Tags []LedgerTag `json:"tags" gorm:"many2many:ledger_tag_tags;"`
}

// LedgerUser 关联表
type LedgerUser struct {
	LedgerID string `json:"ledger_id" gorm:"index;size:255;comment:账本ID"`
	UserID   string `json:"user_id" gorm:"index;size:255;comment:用户ID"`
}
