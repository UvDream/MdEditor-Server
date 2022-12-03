package ledger

import (
	"server/models"
	"server/models/system"
	"time"
)

// Bill 账单
type Bill struct {
	models.Model
	//	账单备注
	Remark string `json:"remark" gorm:"type:varchar(255);comment:账单备注"`
	//创建时间
	CreateTime time.Time `json:"create_time" gorm:"type:timestamp;comment:创建时间"`
	Date       string    `json:"date" gorm:"-"`
	//	账单金额
	Amount float64 `json:"amount" gorm:"type:decimal(10,2);comment:账单金额;default:0.00"`
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
	//	不计入预算
	NotBudget string `json:"not_budget" gorm:"type:varchar(255);comment:不计入预算;default:'0'"` //0 不计入 1 计入
}
type BillRequest struct {
	//	账单名称
	Name string `form:"name" json:"name" `
	//	账单金额
	Amount string `form:"amount" json:"amount"`
	//	账单ID
	LedgerID  string `form:"ledger_id" json:"ledger_id"`
	StartTime string `form:"start_time" json:"start_time"`
	EndTime   string `form:"end_time" json:"end_time"`
}

type BillChildren struct {
	// 	时间
	Time string `json:"time"`
	//	收入
	Income float64 `json:"income"`
	//	支出
	Expenditure float64 `json:"expenditure"`
	//	账单
	Bill []Bill `json:"bill"`
}

const (
	TimeFormat = "\"2006-01-02 15:04:05\""
)

type Time int64

func (p *Time) UnmarshalJSON(b []byte) error {
	parseTime, err := time.Parse(TimeFormat, string(b))
	if err != nil {
		return err
	}
	*p = Time(parseTime.Unix())
	return nil
}

func (p Time) MarshalJSON() ([]byte, error) {
	if p == Time(0) {
		return nil, nil
	}
	unix := time.Unix(int64(p), 0)
	format := unix.Format(TimeFormat)
	return []byte(format), nil
}
