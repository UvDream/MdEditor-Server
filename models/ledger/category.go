package ledger

import (
	"server/models"
	"server/models/system"
)

// CategoryLedger 分类
type CategoryLedger struct {
	models.Model
	// 分类名称
	Name string `json:"name"`
	// 分类缩略图
	Thumbnail string `json:"thumbnail"`
	// 分类图标类型 1 图片/2 icon
	IconType  string `json:"icon_type" gorm:"default:'icon'"`
	Img       string `json:"img"`
	ClassName string `json:"class_name"`
	// 分类描述
	Description string `json:"description"`
	// 分类创建者
	User system.User `json:"user"`
	// 分类创建者ID
	UserID string `json:"user_id"`
	//父元素ID
	ParentID string `json:"parent_id"`
	//	类型  0 支出/ 1 收入....
	Type string `json:"type"`
	//账本ID
	LedgerID string           `json:"ledger_id"`
	Children []CategoryLedger `json:"children" gorm:"-"`
	//	颜色
	Color string `json:"color"`
	//	背景色
	BackgroundColor string `json:"background_color"`
}

// IconClassification 图标
type IconClassification struct {
	models.Model
	// 分类名称
	Name string `json:"name"`
	// 分类和图标一对多
	Icons []Icon `json:"icons"`
	////	用户ID
	UserID string `json:"user_id"`
	//	用户
	User system.User `json:"user" `
	// 账本
	LedgerID string `json:"ledger_id"`
	Ledger   Ledger `json:"ledger" gorm:"-"`
	//	是否公开
	IsPublic string `json:"is_public" gorm:"default:0"`
}
type Icon struct {
	models.Model
	// 图标类型
	Type      string `json:"type"`
	Img       string `json:"img"`
	Unicode   string `json:"unicode"`
	ClassName string `json:"class_name"`
	//	分类id
	IconClassificationID string `json:"icon_classification_id"`
	UserID               string `json:"user_id"`
	//	用户
	User system.User `json:"user"`
}
type Color struct {
	models.Model
	// 颜色
	Color string `json:"color"`
	//是否是背景
	IsBgColor string `json:"is_bg_color"`
	//	用户ID
	UserID string `json:"user_id"`
	//	用户
	User system.User `json:"user"`
	//	是否公开
	IsPublic string `json:"is_public" gorm:"default:0"`
}

var InitLedgerCategory = []CategoryLedger{
	// 支出
	{
		Name: "购物消费",
		Type: "0",
	},
	{
		Name: "餐饮消费",
		Type: "0",
	},
	{
		Name: "交通出行",
		Type: "0",
	},
	{
		Name: "居家生活",
		Type: "0",
	},
	{
		Name: "医疗保健",
		Type: "0",
	},
	{
		Name: "人情往来",
		Type: "0",
	},
	{
		Name: "出行交通",
		Type: "0",
	},
	{
		Name: "娱乐消费",
		Type: "0",
	},
	//	收入
	{
		Name: "工资",
		Type: "1",
	},
	{
		Name: "兼职",
		Type: "1",
	},
	{
		Name: "理财",
		Type: "1",
	},
	{
		Name: "中奖",
		Type: "1",
	},
	{
		Name: "二手闲置",
		Type: "1",
	},
	{
		Name: "补贴",
		Type: "1",
	},
	{
		Name: "奖金",
		Type: "1",
	},
}
