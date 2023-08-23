package system

import "server/models"

type Feedback struct {
	models.Model
	//	类型
	Type string `json:"type" form:"type" gorm:"comment:类型"` // 1:bug,2:建议,3:其他
	//	内容
	Content string `json:"content" gorm:"comment:内容"`
	//	QQ联系方式
	QQ string `json:"qq" form:"qq" gorm:"comment:QQ联系方式"`
	//	邮箱联系方式
	Email string `json:"email" form:"email" gorm:"comment:邮箱联系方式"`
	//	用户ID
	UserID string `json:"user_id" gorm:"comment:用户ID"`
	//	用户
	User User `json:"user" gorm:"foreignKey:UserID"`
	//接受团队联系
	AcceptTeam string `json:"accept_team" gorm:"comment:接受团队联系"` // 1:接受,2:不接受
	//	处理状态
	Status string `json:"status" form:"status" gorm:"comment:处理状态;default:0"` // 0:未处理,1:已处理
	// 多图片附件
	Attachments []Attachment `json:"attachments" gorm:"foreignKey:FeedbackID"`
}
type Attachment struct {
	models.Model
	//	反馈ID
	FeedbackID string `json:"feedback_id" gorm:"comment:反馈ID"`
	//	反馈
	Feedback Feedback `json:"feedback" gorm:"foreignKey:FeedbackID"`
	//	附件地址
	Url string `json:"url" gorm:"comment:附件地址"`
}
