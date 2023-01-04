package ledger

import (
	"server/models"
	"server/models/system"
	"time"
)

type Ledger struct {
	models.Model
	//  账本名称
	Name string `json:"name"`
	//	缩略图
	Thumbnail string `json:"thumbnail"`
	//  图标
	IconType  string `json:"icon_type" gorm:"type:varchar(255);comment:图标类型;default:'icon'"`
	ClassName string `json:"class_name" gorm:"type:varchar(255);comment:图标class名称"`
	Img       string `json:"img" gorm:"type:varchar(255);comment:图标图片"`
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
	Tags          []TagLedger `json:"tags"`
	ShareCode     string      `json:"share_code"`
	ShareCodeTime time.Time   `json:"share_code_time"`
}

// LedgerUser 关联表
type LedgerUser struct {
	LedgerID string `json:"ledger_id" gorm:"index;size:255;comment:账本ID"`
	UserID   string `json:"user_id" gorm:"index;size:255;comment:用户ID"`
	//	权限
	Permission string `json:"permission" gorm:"size:255;comment:权限"` // 0读写权限/1只读权限
}
type HomeStatisticsData struct {
	//	收入
	Income float64 `json:"income"`
	//	支出
	Expenditure float64 `json:"expenditure"`
	//	预算
	Budget float64 `json:"budget"`
}

type WeChatUserInfo struct {
	//UnionID
	OpenID string `json:"open_id"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
	//性别
	Gender string `json:"gender"`
	Phone  string `json:"phone"`
}

type DataSummary struct {
	//收入
	Income float64 `json:"income"`
	//支出
	Expenditure float64 `json:"expenditure"`
}

//LoopAccount 循环记账
type LoopAccount struct {
	models.Model
	//	类型
	Class string `json:"class"` //day/week/month/year/working-day/statutory-holiday/statutory-working-day
	//	状态
	Status string `json:"status"`
	//	分类ID
	CategoryID string `json:"category_id"`
	//	账本ID
	LedgerID string `json:"ledger_id"`
	Ledger   Ledger `json:"ledger"`
	//	金额
	Amount float64 `json:"amount" gorm:"type:decimal(10,2)"`
	//日期
	CreateTime time.Time `json:"create_time" gorm:"type:timestamp;comment:创建时间"`
	//	账单类型
	Type string `json:"type" gorm:"type:varchar(255)"`
	//	不计入预算
	NotBudget string `json:"not_budget" gorm:"type:varchar(255);comment:不计入预算;default:'0'"` //0 不计入 1 计入
	//	作者ID
	CreatorID string `json:"creator_id" gorm:"type:varchar(255);comment:账单创建者ID;"`
	//	账单创建者
	Creator system.User `json:"creator" gorm:"foreignKey:CreatorID;references:ID"`
}
