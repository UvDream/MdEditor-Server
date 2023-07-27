package system

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
	"server/code"
	"server/global"
	"server/models/system"
	"server/utils"
)

func (*SysUserService) FeedbackService(feedback system.Feedback) (data system.Feedback, cd int, err error) {
	if err := global.DB.Create(&feedback).Error; err != nil {
		return feedback, code.ErrorFeedback, err
	}
	return feedback, code.SUCCESS, nil
}

func (*SysUserService) FeedbackListService(userId string, c *gin.Context) (data []system.Feedback, total int64, cd int, err error) {
	db := global.DB
	db = db.Where("user_id = ?", userId)
	//TODO 查询用户角色,管理员可以查看所有反馈
	//	分页查询
	if err := db.Model(&system.Feedback{}).Count(&total).Error; err != nil {
		return data, total, code.ErrorFeedbackList, err
	}
	keyword := c.Query("keyword")
	//	获取列表
	if err := db.Where("content LIKE ?", "%"+keyword+"%").Preload(clause.Associations).Scopes(utils.Paginator(c)).Order("create_time desc").Find(&data).Error; err != nil {
		return data, total, code.ErrorFeedbackList, err
	}
	return data, total, code.SUCCESS, nil
}
