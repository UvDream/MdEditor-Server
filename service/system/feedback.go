package system

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
	"server/code"
	"server/global"
	"server/models/system"
	"server/service/ledger"
	"server/utils"
)

func (*SysUserService) FeedbackService(feedback system.Feedback) (data system.Feedback, cd int, err error) {
	if err := global.DB.Create(&feedback).Error; err != nil {
		return feedback, code.ErrorFeedback, err
	}
	return feedback, code.SUCCESS, nil
}

func (*SysUserService) IsVipService(userId string) (data bool, cd int, err error) {
	data = ledger.FindIsMember(userId, global.DB)
	return data, code.SUCCESS, nil
}

func (*SysUserService) FeedbackListService(userId string, query system.Feedback, keyword string, c *gin.Context) (data []system.Feedback, total int64, cd int, err error) {
	db := global.DB
	db = db.Where("user_id = ?", userId)
	//TODO 查询用户角色,管理员可以查看所有反馈
	//关键字
	if keyword != "" {
		db = db.Where("content LIKE ?", "%"+keyword+"%")
	}
	//QQ搜索
	if query.QQ != "" {
		db = db.Where("qq LIKE ?", "%"+query.QQ+"%")
	}
	//邮箱搜索
	if query.Email != "" {
		db = db.Where("email LIKE ?", "%"+query.Email+"%")
	}
	//状态搜索
	if query.Status != "" {
		db = db.Where("status = ?", query.Status)
	}

	//	分页查询
	if err := db.Model(&system.Feedback{}).Count(&total).Error; err != nil {
		return data, total, code.ErrorFeedbackList, err
	}
	//	获取列表
	if err := db.Preload(clause.Associations).Scopes(utils.Paginator(c)).Order("create_time desc").Find(&data).Error; err != nil {
		return data, total, code.ErrorFeedbackList, err
	}
	return data, total, code.SUCCESS, nil
}
