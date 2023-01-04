package ledger

import (
	"server/models"
	"server/models/system"
)

type MoneyBudget struct {
	models.Model
	//	预算
	Budget float64 `json:"budget"`
	//	预算类型
	BudgetType string `json:"budget_type"` //0:月预算 1:年预算 2:自定义预算
	//	预算开始时间
	BudgetStartTime string `json:"budget_start_time"`
	//	预算结束时间
	BudgetEndTime string `json:"budget_end_time"`
	//是否每个月都相同
	IsSameMonth bool `json:"is_same_month"`
	//年份
	Year int `json:"year"`
	//	日期
	Date int `json:"date"` //月
	//	账本ID
	LedgerID string `json:"ledger_id"`
	//	账单创建者ID
	CreatorID string `json:"creator_id" gorm:"type:varchar(255);comment:账单创建者ID;"`
	//	账单创建者
	Creator system.User `json:"creator" gorm:"foreignKey:CreatorID;references:ID"`
	//	支出
	Expenditure float64 `json:"expenditure" gorm:"-"`
}

type BudgetDelete struct {
	Year       string `json:"year"`
	BudgetType string `json:"budget_type"`
	// 账本ID
	LedgerID string `json:"ledger_id"`
}
