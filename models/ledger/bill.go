package ledger

import (
	"server/models"
	"server/models/system"
)

// Bill 账单
type Bill struct {
	models.Model
	//	账单备注
	Remark string `json:"remark" gorm:"type:varchar(255);comment:账单备注"`
	//	账单金额
	Amount float64 `json:"amount" gorm:"type:decimal(10,2);comment:账单金额"`
	//	账单类型
	Type string `json:"type" gorm:"type:varchar(255);comment:账单类型"`
	//	账单状态
	Status string `json:"status" gorm:"type:varchar(255);comment:账单状态"`
	//	账单创建者ID
	CreatorID string `json:"creator_id" gorm:"type:varchar(255);comment:账单创建者ID;"`
	//	账单创建者
	Creator system.User `json:"creator" gorm:"foreignKey:CreatorID;references:ID"`
	//	账单分类ID
	CategoryID string `json:"category_id" gorm:"type:varchar(255);comment:账单分类ID"`
	//	账单分类
	Category CategoryLedger `json:"ledger_category"`
	//	账单标签
	Tags []TagLedger `json:"tags" gorm:"many2many:bill_tags;"`
	// 账本ID
	LedgerID string `json:"ledger_id" gorm:"type:varchar(255);comment:账本ID"`
	//	账单所属账本
	Ledger    Ledger `json:"ledger"`
	IconType  string `json:"icon_type" gorm:"type:varchar(255);comment:图标类型;default:'icon'"`
	ClassName string `json:"class_name" gorm:"type:varchar(255);comment:图标class名称"`
	Img       string `json:"img" gorm:"type:varchar(255);comment:图标图片"`
}
type BillRequest struct {
	//	账单名称
	Name string `form:"name" json:"name" `
	//	账单金额
	Amount string `form:"amount" json:"amount"`
	//	账单ID
	LedgerID string `form:"ledger_id" json:"ledger_id"`
}
