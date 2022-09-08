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
}

// LedgerTag
type LedgerTag struct {
	LedgerID string `json:"ledger_id" gorm:"index;size:255;comment:账本ID"`
	TagID    string `json:"tag_id" gorm:"index;size:255;comment:标签ID"`
}
