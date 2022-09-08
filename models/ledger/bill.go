package ledger

import (
	"server/models"
	"server/models/system"
)

// Bill 账单
type Bill struct {
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
