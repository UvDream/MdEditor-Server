package ledger

import (
	"server/models"
	"server/models/system"
)

// LedgerTag 标签
type LedgerTag struct {
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
