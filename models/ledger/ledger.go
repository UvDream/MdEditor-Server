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

// LedgerCategory 分类
type LedgerCategory struct {
	models.Model
	// 分类名称
	Name string `json:"name"`
	// 分类缩略图
	Thumbnail string `json:"thumbnail"`
	// 分类描述
	Description string `json:"description"`
	// 分类状态
	Status string `json:"status"`
	// 分类创建者
	User system.User `json:"creator"`
	// 分类创建者ID
	UserID string `json:"creator_id"`
}

// LedgerTag 标签
type LedgerTag struct {
	models.Model
	// 标签名称
	Name string `json:"name"`
	// 标签缩略图
	Thumbnail string `json:"thumbnail"`
	// 标签描述
	Description string `json:"description"`
	// 标签状态
	Status string `json:"status"`
	// 标签创建者
	User system.User `json:"creator"`
	// 标签创建者ID
	UserID string `json:"creator_id"`
}

// LedgerBill 账单
type LedgerBill struct {
	models.Model `json:"models_._model"`
	//	账单名称
	Name string `json:"name"`
	//	账单金额
	Amount float64 `json:"amount" `
	//	账单类型
	Type string `json:"type"`
	//	账单状态
	Status string `json:"status"`
	//	账单创建者ID
	CreatorID string `json:"creator_id"`
	//	账单创建者
	Creator system.User `json:"creator"`
	//	账单分类ID
	CategoryID string `json:"category_id"`
	//	账单分类
	Category LedgerCategory `json:"ledger_category"`
	//	账单标签
	Tags []LedgerTag `json:"tags" gorm:"many2many:ledger_bill_tags;"`
}
