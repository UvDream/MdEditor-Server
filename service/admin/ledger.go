package admin

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
	"server/code"
	"server/global"
	"server/models/ledger"
	service "server/service/system"
	"server/utils"
)

type LedgerAdminService struct{}

func (*LedgerAdminService) GetLedgerListService(query ledger.LedgerRequest, c *gin.Context) (data []ledger.Ledger, total int64, cd int, err error) {
	userId := utils.FindUserID(c)
	var userService service.SysUserService
	roles, err := userService.FindUserRoles(userId)
	if err != nil {
		return nil, 0, code.ErrorMissingLedgerId, err
	}
	if len(roles) > 0 {
		for _, v := range roles {
			if v == "admin" || v == "root" {
				data, total, cd, err = getAllLedgerList(query, c)
				return data, total, cd, err
			}
		}
	}
	return data, 0, code.SUCCESS, err
}

func getAllLedgerList(query ledger.LedgerRequest, c *gin.Context) (data []ledger.Ledger, total int64, cd int, err error) {
	db := global.DB
	if query.Name != "" {
		db = db.Where("name LIKE ?", "%"+query.Name+"%")
	}
	if query.Type != "" {
		db = db.Where("type = ?", query.Type)
	}
	if query.Description != "" {
		db = db.Where("description LIKE ?", "%"+query.Description+"%")
	}
	//开始时间和结束时间
	if query.StartTime != "" && query.EndTime != "" {
		db = db.Where("created_at BETWEEN ? AND ?", query.StartTime, query.EndTime)
	}
	//查询总数
	if err := db.Model(&ledger.Ledger{}).Count(&total).Error; err != nil {
		return nil, 0, code.ErrorMissingLedgerId, err
	}
	//查询数据
	if err := db.Preload(clause.Associations).Scopes(utils.Paginator(c)).Find(&data).Error; err != nil {
		return nil, 0, code.ErrorLedgerNotExist, err
	}
	return data, total, code.SUCCESS, err
}
