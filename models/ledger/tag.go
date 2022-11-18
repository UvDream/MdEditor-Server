package ledger

import (
	"server/models"
	"server/models/system"
)

// TagLedger 标签
type TagLedger struct {
	models.Model
	// 标签名称
	Name string `json:"name"`
	// 标签描述
	Description string `json:"description"`
	// 标签创建者
	User system.User `json:"creator"`
	// 标签创建者ID
	UserID string `json:"creator_id"`
	//	账本ID
	LedgerID string `json:"ledger_id"`
}

type BillTag struct {
	BillID string `json:"bill_id" gorm:"index;size:255;comment:账单ID"`
	TagID  string `json:"tag_id" gorm:"index;size:255;comment:标签ID"`
}
