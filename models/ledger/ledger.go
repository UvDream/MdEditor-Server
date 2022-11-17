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
	//  图标
	Icon string `json:"icon"`
	//	账本描述
	Description string `json:"description"`
	//	账本类型
	Type string `json:"type"` // 0:个人账本 1:家庭账本 2:团队账本
	//	账本创建者
	Creator system.User `json:"creator"`
	//	账本创建者ID
	CreatorID string `json:"creator_id"`
	//	账本成员ID
	User []system.User `json:"user" gorm:"many2many:ledger_users;"`
	//	账本成员数量
	MemberCount int64 `json:"member_count" gorm:"-"`
	// 分类
	Categories []CategoryLedger `json:"categories"`
	//	标签
	Tags []TagLedger `json:"tags"`
}

// LedgerUser 关联表
type LedgerUser struct {
	LedgerID string `json:"ledger_id" gorm:"index;size:255;comment:账本ID"`
	UserID   string `json:"user_id" gorm:"index;size:255;comment:用户ID"`
	//	权限
	Permission string `json:"permission" gorm:"size:255;comment:权限"` // 读写权限/只读权限
}
